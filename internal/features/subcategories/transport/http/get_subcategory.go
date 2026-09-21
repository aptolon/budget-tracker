package subcategories_transport_http

import (
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

type GetSubcategoryResponse SubcategoryResponse

func (h *SubcategoriesHTTPHandler) GetSubcategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)
	subcategoryID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get subcategoryID path param",
		)
		return
	}
	claims := crypto_token.ClaimsFromContext(ctx)

	subcategoryDomain, err := h.subcategoriesService.GetSubcategory(ctx, subcategoryID, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get subcategory",
		)
		return
	}

	response := GetSubcategoryResponse(subcategoryResponseFromDomain(subcategoryDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
