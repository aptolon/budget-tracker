package subcategories_transport_http

import (
	"net/http"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
)

type CreateSubcategoryRequest SubcategoryRequest

type CreateSubcategoryResponse SubcategoryResponse

func (h *SubcategoriesHTTPHandler) CreateSubcategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	categoryID, err := core_http_request.GetUUIDPathValue(r, "categoryID")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get categoryID path param",
		)
		return
	}

	var request CreateSubcategoryRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	claims := crypto_token.ClaimsFromContext(ctx)
	subcategoryDomain, err := h.subcategoriesService.CreateSubcategory(
		ctx,
		categoryID,
		request.Title,
		claims.UserID,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create subcategory",
		)
		return
	}
	response := CreateSubcategoryResponse(subcategoryResponseFromDomain(subcategoryDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}
