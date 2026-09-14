package categories_transport_http

import (
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

type PatchCategoryRequest CategoryRequest
type PatchCategoryResponse CategoryResponse

func (h *CategoriesHTTPHandler) PatchCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	var request PatchCategoryRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	categoryID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get categoryID path param",
		)
		return
	}

	claims := crypto_token.ClaimsFromContext(ctx)

	categoryDomain, err := h.categoriesService.PatchCategory(
		ctx,
		categoryID,
		claims.UserID,
		request.Title,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch category",
		)
		return
	}
	response := PatchCategoryResponse(categoryResponseFromDomain(categoryDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
