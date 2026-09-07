package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_middleware "github.com/aptolon/budget-tracker/internal/core/transport/http/middleware"
	core_http_server "github.com/aptolon/budget-tracker/internal/core/transport/http/server"

	auth_service "github.com/aptolon/budget-tracker/internal/features/auth/service"
	auth_transport_http "github.com/aptolon/budget-tracker/internal/features/auth/transport/http"
	users_postgres_repository "github.com/aptolon/budget-tracker/internal/features/users/repository/postgres"
	users_service "github.com/aptolon/budget-tracker/internal/features/users/service"
	users_transport_http "github.com/aptolon/budget-tracker/internal/features/users/transport/http"

	crypto_hasher "github.com/aptolon/budget-tracker/internal/core/crypto/hasher"
	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"

	core_pgx_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool/pgx"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing PostgreSQL connection pool")

	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("Postgres connection pool init error", zap.Error(err))
	}
	defer pool.Close()

	httpConfig := core_http_server.NewConfigMust()

	logger.Debug("Initializing hasher")
	hasher := crypto_hasher.NewBcryptHasher(crypto_hasher.NewConfigMust())

	logger.Debug("Initializing token service")
	tokenService := crypto_token.NewJWT(crypto_token.NewConfigMust())

	logger.Debug("Initializing features", zap.String("features", "auth"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	authService := auth_service.NewAuthService(usersRepository, hasher, tokenService)
	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService, httpConfig.SecureCookies)

	logger.Debug("Initializing features", zap.String("features", "users"))
	usersService := users_service.NewUsersService(usersRepository, hasher)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService, httpConfig.SecureCookies)

	logger.Debug("Initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		httpConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	auth := core_http_middleware.Auth(tokenService)
	requireAdmin := core_http_middleware.RequireAdmin()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.AddRoutes(authTransportHTTP.Routes()...)
	apiVersionRouter.AddRoutes(
		usersTransportHTTP.Routes(
			auth,
			requireAdmin,
		)...,
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error: %w", zap.Error(err))
	}
}
