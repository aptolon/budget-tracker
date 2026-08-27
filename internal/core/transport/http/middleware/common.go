package core_http_middleware

import (
	"net/http"
	"time"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	domain "github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHeader(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HHTP got unexpected panic",
					)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", before.UTC()),
			)
			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetSasusCode()),
				zap.Duration("latency", time.Since(before)),
			)
		})
	}
}

func Auth(token crypto_token.TokenService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			cookie, err := r.Cookie("access_token")
			responseHandler := core_http_response.NewHTTPResponseHeader(log, w)
			if err != nil {
				if err == http.ErrNoCookie {
					responseHandler.ErrorResponse(
						core_errors.ErrInvalidCredentials,
						"access token not found",
					)
					return
				}

				responseHandler.ErrorResponse(
					err,
					"failed to get access token cookie",
				)
				return
			}
			claims, err := token.Validate(cookie.Value)

			if err != nil {
				responseHandler.ErrorResponse(core_errors.ErrInvalidCredentials, "invalid access token")
				return
			}
			ctx = crypto_token.ClaimsToContext(ctx, claims)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
func RequireAdmin() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			responseHandler := core_http_response.NewHTTPResponseHeader(log, w)
			claims := crypto_token.ClaimsFromContext(ctx)

			if claims.Role != domain.RoleAdmin {
				responseHandler.ErrorResponse(core_errors.ErrForbidden, "admin role required")
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
