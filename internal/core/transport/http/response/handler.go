package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_logger "github.com/aptolon/budget-tracker/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHeader struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHeader(
	log *core_logger.Logger,
	rw http.ResponseWriter,
) *HTTPResponseHeader {
	return &HTTPResponseHeader{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHeader) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHeader) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}
func (h *HTTPResponseHeader) SetAccessTokenCookie(
	token string,
	secureCookies bool,
) {
	http.SetCookie(
		h.rw,
		&http.Cookie{
			Name:     "access_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   secureCookies,
			SameSite: http.SameSiteLaxMode,
		},
	)

}

func (h *HTTPResponseHeader) DeleteTokenCookie(
	name string,
	secureCookies bool,
) {
	http.SetCookie(
		h.rw,
		&http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   secureCookies,
			SameSite: http.SameSiteLaxMode,
		},
	)
}

func (h *HTTPResponseHeader) SetRefreshTokenCookie(
	token string,
	secureCookies bool,
) {
	http.SetCookie(
		h.rw,
		&http.Cookie{
			Name:     "refresh_token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   secureCookies,
			SameSite: http.SameSiteLaxMode,
		},
	)
}
func (h *HTTPResponseHeader) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)
	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrInvalidCredentials):
		statusCode = http.StatusUnauthorized
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrForbidden):
		statusCode = http.StatusForbidden
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrLoginTaken):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrInternal):
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg,
	)

}
func (h *HTTPResponseHeader) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic : %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg,
	)

}

func (h *HTTPResponseHeader) errorResponse(
	statusCode int,
	err error,
	msg string,
) {

	response := map[string]string{
		"message": msg,
		"err":     err.Error(),
	}
	h.JSONResponse(
		response,
		statusCode,
	)
}
