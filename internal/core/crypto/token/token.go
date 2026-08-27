package crypto_token

import (
	"context"

	"github.com/google/uuid"
)

type claimsContextKey struct{}

var key = claimsContextKey{}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

func (t TokenType) IsValid() bool {
	switch t {
	case TokenTypeAccess, TokenTypeRefresh:
		return true
	default:
		return false
	}
}

type TokenService interface {
	Generate(userID uuid.UUID, role string, tokenType TokenType) (string, error)
	Validate(token string) (Claims, error)
}

type Claims struct {
	UserID    uuid.UUID
	Role      string
	TokenType TokenType
}

func ClaimsToContext(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, key, claims)
}

func ClaimsFromContext(ctx context.Context) Claims {
	claims, ok := ctx.Value(claimsContextKey{}).(Claims)
	if !ok {
		panic("claims not found in context")
	}

	return claims
}
