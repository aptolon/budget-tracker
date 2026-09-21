package subcategories_transport_http

import (
	"fmt"
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

type GetSubcategoriesResponse []SubcategoryResponse

func (h *SubcategoriesHTTPHandler) GetSubcategories(rw http.ResponseWriter, r *http.Request) {
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
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query param",
		)
		return
	}

	claims := crypto_token.ClaimsFromContext(ctx)
	subcategoryDomains, err := h.subcategoriesService.GetSubcategories(ctx, categoryID, claims.UserID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get subcategories",
		)
		return
	}

	response := GetSubcategoriesResponse(subcategoriesResponseFromDomains(subcategoryDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
