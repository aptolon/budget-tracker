package auth_service

import (
	"context"
	"errors"
	"fmt"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (string, string, error) {
	if err := domain.ValidateLogin(login); err != nil {
		return "", "", fmt.Errorf("login user: %w", err)
	}

	if err := domain.ValidatePassword(password); err != nil {
		return "", "", fmt.Errorf("login user: %w", err)
	}

	user, err := s.usersRepository.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return "", "", fmt.Errorf(
				"login user: %w",
				core_errors.ErrInvalidCredentials,
			)
		}

		return "", "", fmt.Errorf("login user: %w", err)
	}

	if err := s.hasher.Compare(password, user.PasswordHash); err != nil {
		return "", "", fmt.Errorf(
			"login user: %w",
			core_errors.ErrInvalidCredentials,
		)
	}

	accessToken, err := s.tokenService.Generate(
		user.ID,
		user.Role,
		crypto_token.TokenTypeAccess,
	)
	if err != nil {
		return "", "", fmt.Errorf(
			"login user: generate access token: %w",
			err,
		)
	}

	refreshToken, err := s.tokenService.Generate(
		user.ID,
		user.Role,
		crypto_token.TokenTypeRefresh,
	)
	if err != nil {
		return "", "", fmt.Errorf(
			"login user: generate refresh token: %w",
			err,
		)
	}

	return accessToken, refreshToken, nil
}
