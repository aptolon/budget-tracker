package auth_transport_http

import (
	"net/http"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	responseHandler.DeleteTokenCookie("access_token", h.secureCookies)
	responseHandler.DeleteTokenCookie("refresh_token", h.secureCookies)

	responseHandler.NoContentResponse()
}
