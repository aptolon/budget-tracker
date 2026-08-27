package auth_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func (s *AuthService) Register(
	ctx context.Context,
	login string,
	password string,
) (domain.User, error) {
	if err := domain.ValidateLogin(login); err != nil {
		return domain.User{}, fmt.Errorf("register user: %w", err)
	}

	if err := domain.ValidatePassword(password); err != nil {
		return domain.User{}, fmt.Errorf("register user: %w", err)
	}
	exists, err := s.usersRepository.ExistUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("register user: %w", err)
	}

	if exists {
		return domain.User{}, core_errors.ErrLoginTaken
	}
	passwordHash, err := s.hasher.Hash(password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	user := domain.CreateUser(
		domain.RoleUser,
		login,
		passwordHash,
	)
	if err = s.usersRepository.CreateUser(ctx, user); err != nil {
		return domain.User{}, fmt.Errorf("register user: %w", err)
	}
	return user, nil

}
