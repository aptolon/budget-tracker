package subcategories_transport_http

import (
	"context"
	"net/http"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"

	core_http_middleware "github.com/aptolon/budget-tracker/internal/core/transport/http/middleware"
	core_http_server "github.com/aptolon/budget-tracker/internal/core/transport/http/server"
)

type SubcategoriesHTTPHandler struct {
	subcategoriesService SubcategoriesService
}

type SubcategoriesService interface {
	CreateSubcategory(
		ctx context.Context,
		categoryID uuid.UUID,
		title string,
		userID uuid.UUID,
	) (domain.Subcategory, error)

	GetSubcategories(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Subcategory, error)

	GetSubcategory(
		ctx context.Context,
		subcategoryID uuid.UUID,
		userID uuid.UUID,
	) (domain.Subcategory, error)
	DeleteSubcategory(
		ctx context.Context,
		subcategoryID uuid.UUID,
		userID uuid.UUID,
	) error
	PatchSubcategory(
		ctx context.Context,
		subcategoryID uuid.UUID,
		title string,
		userID uuid.UUID,
	) (domain.Subcategory, error)
}

func NewSubcategoriesHTTPHandler(
	subcategoriesService SubcategoriesService,
) *SubcategoriesHTTPHandler {
	return &SubcategoriesHTTPHandler{
		subcategoriesService: subcategoriesService,
	}
}

func (h *SubcategoriesHTTPHandler) Routes(
	auth core_http_middleware.Middleware,
) []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/categories/{categoryID}/subcategories",
			Handler: h.CreateSubcategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories/{categoryID}/subcategories",
			Handler: h.GetSubcategories,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/subcategories/{id}",
			Handler: h.GetSubcategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/subcategories/{id}",
			Handler: h.DeleteSubcategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/subcategories/{id}",
			Handler: h.PatchSubcategory,
			Middleware: []core_http_middleware.Middleware{
				auth,
			},
		},
	}
}
