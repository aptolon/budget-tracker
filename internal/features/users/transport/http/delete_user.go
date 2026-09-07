package users_transport_http

import (
	"net/http"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)
	userID, err := core_http_request.GetUUIDPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path param",
		)
		return
	}
	claims := crypto_token.ClaimsFromContext(ctx)
	if claims.Role != domain.RoleAdmin && claims.UserID != userID {
		responseHandler.ErrorResponse(
			core_errors.ErrForbidden,
			"admin or owner access required",
		)
		return
	}

	err = h.usersService.DeleteUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete users",
		)
		return
	}
	if claims.UserID == userID {
		responseHandler.DeleteTokenCookie("access_token", h.secureCookies)
		responseHandler.DeleteTokenCookie("refresh_token", h.secureCookies)
	}
	responseHandler.NoContentResponse()

}
