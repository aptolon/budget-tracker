package categories_transport_http

import (
	"net/http"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
)

type CreateCategoryRequest CategoryRequest

type CreateCategoryResponse CategoryResponse

func (h *CategoriesHTTPHandler) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	var request CreateCategoryRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}
	claims := crypto_token.ClaimsFromContext(ctx)
	categoryDomain, err := h.categoriesService.CreateCategory(
		ctx,
		claims.UserID,
		request.Title,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create category",
		)
		return
	}
	response := CreateCategoryResponse(categoryResponseFromDomain(categoryDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}
