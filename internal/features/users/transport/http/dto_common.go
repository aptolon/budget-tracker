package users_transport_http

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
	Role    string    `json:"role"`
	Login   string    `json:"login"`
}

type UserRequest struct {
	Login    string `json:"login"    validate:"required,min=3,max=32"`
	Role     string `json:"role"     validate:"required,oneof=user admin"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

func userResponseFromDomain(user domain.User) UserResponse {
	return UserResponse{
		ID:      user.ID,
		Version: user.Version,
		Role:    user.Role,
		Login:   user.Login,
	}
}

func usersResponseFromDomains(users []domain.User) []UserResponse {
	usersResponse := make([]UserResponse, len(users))
	for i, user := range users {
		usersResponse[i] = userResponseFromDomain(user)
	}
	return usersResponse
}
