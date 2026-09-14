package categories_transport_http

import (
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

func (h *CategoriesHTTPHandler) DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)
	categoryID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get categoryID path param",
		)
		return
	}
	claims := crypto_token.ClaimsFromContext(ctx)

	err = h.categoriesService.DeleteCategory(ctx, categoryID, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete category",
		)
		return
	}
	responseHandler.NoContentResponse()

}
