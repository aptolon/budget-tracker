package users_transport_http

import (
	"context"
	"net/http"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"

	core_http_middleware "github.com/aptolon/budget-tracker/internal/core/transport/http/middleware"
	core_http_server "github.com/aptolon/budget-tracker/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService  UsersService
	secureCookies bool
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		login string,
		role string,
		password string,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limin *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		userID uuid.UUID,
	) (domain.User, error)
	PatchUser(
		ctx context.Context,
		userID uuid.UUID,
		login *string,
		role *string,
		password *string,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		userID uuid.UUID,
	) error
}

func NewUsersHTTPHandler(
	usersService UsersService,
	secureCookies bool,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService:  usersService,
		secureCookies: secureCookies,
	}
}

func (h *UsersHTTPHandler) Routes(
	auth core_http_middleware.Middleware,
	requireAdmin core_http_middleware.Middleware,
) []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
			Middleware: []core_http_middleware.Middleware{
				auth,
				requireAdmin,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
			Middleware: []core_http_middleware.Middleware{
				auth,
				requireAdmin,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
	}
}
