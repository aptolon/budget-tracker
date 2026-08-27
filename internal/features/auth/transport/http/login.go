package auth_transport_http

import (
	"net/http"

	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	core_http_request "github.com/aptolon/budget-tracker/internal/core/transport/http/request"
	core_http_response "github.com/aptolon/budget-tracker/internal/core/transport/http/response"
)

type LoginRequest UserDTORequest

func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	accessToken, refreshToken, err := h.authService.Login(
		ctx,
		request.Login,
		request.Password,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login")
		return
	}

	responseHandler.SetAccessTokenCookie(
		accessToken,
		h.secureCookies,
	)

	responseHandler.SetRefreshTokenCookie(
		refreshToken,
		h.secureCookies,
	)

	responseHandler.NoContentResponse()
}
