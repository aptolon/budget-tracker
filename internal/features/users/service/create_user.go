package users_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	login string,
	role string,
	password string,
) (domain.User, error) {
	if err := domain.ValidateLogin(login); err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	if err := domain.ValidatePassword(password); err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	if err := domain.ValidateRole(role); err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	exists, err := s.usersRepository.ExistUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	if exists {
		return domain.User{}, fmt.Errorf("login already taken: %w", core_errors.ErrConflict)
	}

	passwordHash, err := s.hasher.Hash(password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user := domain.CreateUser(
		role,
		login,
		passwordHash,
	)
	user, err = s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}
