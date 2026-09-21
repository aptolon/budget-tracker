package subcategories_transport_http

import (
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

type PatchSubcategoryRequest SubcategoryRequest
type PatchSubcategoryResponse SubcategoryResponse

func (h *SubcategoriesHTTPHandler) PatchSubcategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	var request PatchSubcategoryRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	subcategoryID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get subcategoryID path param",
		)
		return
	}

	claims := crypto_token.ClaimsFromContext(ctx)

	subcategoryDomain, err := h.subcategoriesService.PatchSubcategory(
		ctx,
		subcategoryID,
		request.Title,
		claims.UserID,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch subcategory",
		)
		return
	}
	response := PatchSubcategoryResponse(subcategoryResponseFromDomain(subcategoryDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
