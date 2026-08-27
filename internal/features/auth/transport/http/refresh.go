package auth_transport_http

import (
	"net/http"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func (h *AuthHTTPHandler) Refresh(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		if err == http.ErrNoCookie {
			responseHandler.ErrorResponse(
				core_errors.ErrInvalidCredentials,
				"refresh token not found",
			)
			return
		}
		responseHandler.ErrorResponse(
			err,
			"failed to get refresh token cookie",
		)
		return
	}

	accessToken, err := h.authService.Refresh(
		ctx,
		cookie.Value,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to refresh access token",
		)
		return
	}

	responseHandler.SetAccessTokenCookie(
		accessToken,
		h.secureCookies,
	)
	responseHandler.NoContentResponse()
}
