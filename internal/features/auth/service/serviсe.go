package auth_service

import (
	"context"

	crypto_hasher "github.com/aptolon/budget-tracker/internal/core/crypto/hasher"
	crypto_token "github.com/aptolon/budget-tracker/internal/core/crypto/token"
	"github.com/aptolon/budget-tracker/internal/core/domain"
)

type AuthService struct {
	usersRepository UsersRepository
	hasher          crypto_hasher.Hasher
	tokenService    crypto_token.TokenService
}
type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) error

	ExistUserByLogin(
		ctx context.Context,
		login string,
	) (bool, error)
	GetUserByLogin(
		ctx context.Context,
		login string,
	) (domain.User, error)
}

func NewAuthService(
	usersRepository UsersRepository,
	hasher crypto_hasher.Hasher,
	tokenService crypto_token.TokenService,
) *AuthService {
	return &AuthService{
		usersRepository: usersRepository,
		hasher:          hasher,
		tokenService:    tokenService,
	}
}
