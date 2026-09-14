package categories_transport_http

import (
	"context"
	"net/http"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"

	core_http_middleware "github.com/aptolon/budget-tracker/internal/core/transport/http/middleware"
	core_http_server "github.com/aptolon/budget-tracker/internal/core/transport/http/server"
)

type CategoriesHTTPHandler struct {
	categoriesService CategoriesService
}

type CategoriesService interface {
	CreateCategory(
		ctx context.Context,
		userID uuid.UUID,
		title string,
	) (domain.Category, error)

	GetCategories(
		ctx context.Context,
		userID uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Category, error)

	GetCategory(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
	) (domain.Category, error)
	DeleteCategory(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
	) error
	PatchCategory(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
		title string,
	) (domain.Category, error)
}

func NewCategoriesHTTPHandler(
	categoriesService CategoriesService,
) *CategoriesHTTPHandler {
	return &CategoriesHTTPHandler{
		categoriesService: categoriesService,
	}
}

func (h *CategoriesHTTPHandler) Routes(
	auth core_http_middleware.Middleware,
) []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: h.CreateCategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories",
			Handler: h.GetCategories,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories/{id}",
			Handler: h.GetCategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/categories/{id}",
			Handler: h.DeleteCategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/categories/{id}",
			Handler: h.PatchCategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
	}
}
