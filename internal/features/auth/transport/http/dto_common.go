package auth_transport_http

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type AuthResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
	Role    string    `json:"role"`
	Login   string    `json:"login"`
}

func authResponseFromDomain(user domain.User) AuthResponse {
	return AuthResponse{
		ID:      user.ID,
		Version: user.Version,
		Role:    user.Role,
		Login:   user.Login,
	}
}

type AuthRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}
