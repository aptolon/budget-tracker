package domain

import (
	"fmt"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID           uuid.UUID
	Version      int
	Role         string
	Login        string
	PasswordHash string
}

func NewUser(
	id uuid.UUID,
	version int,
	role string,
	login string,
	passwordHash string,
) User {
	return User{
		ID:           id,
		Version:      version,
		Role:         role,
		Login:        login,
		PasswordHash: passwordHash,
	}

}

func CreateUser(
	role string,
	login string,
	passwordHash string,
) User {
	var (
		id      = uuid.New()
		version = 1
	)
	return NewUser(
		id,
		version,
		role,
		login,
		passwordHash,
	)
}

func ValidateLogin(login string) error {
	loginLen := len([]rune(login))
	if loginLen < 3 || loginLen > 32 {
		return fmt.Errorf(
			"invalid `login` len %d: %w",
			loginLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordLen := len([]rune(password))
	if passwordLen < 8 || passwordLen > 128 {
		return fmt.Errorf(
			"invalid `password` len %d: %w",
			passwordLen,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
