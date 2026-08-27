package auth_transport_http

import (
	"context"
	"net/http"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_http_middleware "github.com/aptolon/budget-tracker/internal/core/transport/http/middleware"
	core_http_server "github.com/aptolon/budget-tracker/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService   AuthService
	secureCookies bool
}

type AuthService interface {
	Register(
		ctx context.Context,
		login string,
		password string,
	) (domain.User, error)

	Login(
		ctx context.Context,
		login string,
		password string,
	) (string, string, error)
	Refresh(
		ctx context.Context,
		refreshToken string,
	) (string, error)
}

func NewAuthHTTPHandler(
	authService AuthService,
	secureCookies bool,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService:   authService,
		secureCookies: secureCookies,
	}
}

func (h *AuthHTTPHandler) Routes(
	middleware ...core_http_middleware.Middleware,
) []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/register",
			Handler:    h.Register,
			Middleware: middleware,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/logout",
			Handler: h.Logout,
		},
		{
			Method:  http.MethodPost,
			Path:    "/refresh",
			Handler: h.Refresh,
		},
	}
}
