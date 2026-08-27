package auth_transport_http

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type UserDTOResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
	Role    string    `json:"role"`
	Login   string    `json:"login"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:      user.ID,
		Version: user.Version,
		Role:    user.Role,
		Login:   user.Login,
	}
}

type UserDTORequest struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}
