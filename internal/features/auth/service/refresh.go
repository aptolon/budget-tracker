package auth_service

import (
	"context"
	"fmt"

	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
)

func (s *AuthService) Refresh(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	claims, err := s.tokenService.Validate(refreshToken)
	if err != nil {
		return "", fmt.Errorf(
			"refresh token: validate: %w",
			err,
		)
	}
	if claims.TokenType != crypto_token.TokenTypeRefresh {
		return "", fmt.Errorf(
			"refresh token: invalid token type: %q",
			claims.TokenType,
		)
	}
	accessToken, err := s.tokenService.Generate(
		claims.UserID,
		claims.Role,
		crypto_token.TokenTypeAccess,
	)
	if err != nil {
		return "", fmt.Errorf(
			"refresh token: generate access token: %w",
			err,
		)
	}

	return accessToken, nil
}
