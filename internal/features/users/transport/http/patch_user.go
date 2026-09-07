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

type PatchUserRequest struct {
	Login    *string `json:"login"    validate:"omitempty,min=3,max=32"`
	Role     *string `json:"role"     validate:"omitempty,oneof=user admin"`
	Password *string `json:"password" validate:"omitempty,min=8,max=128"`
}

type PatchUserResponse UserResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHeader(log, rw)

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

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
	if request.Role != nil && claims.Role != domain.RoleAdmin {
		responseHandler.ErrorResponse(
			core_errors.ErrForbidden,
			"only admin can change user role",
		)
		return
	}

	userDomain, err := h.usersService.PatchUser(
		ctx,
		userID,
		request.Login,
		request.Role,
		request.Password,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}
	response := PatchUserResponse(userResponseFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}
