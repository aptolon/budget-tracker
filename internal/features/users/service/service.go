package users_service

import (
	"context"

	crypto_hasher "github.com/aptolon/budget-tracker/internal/core/crypto/hasher"
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type UsersService struct {
	usersRepository UsersRepository
	hasher          crypto_hasher.Hasher
}
type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) error
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		userID uuid.UUID,
	) (domain.User, error)
	UpdateUser(
		ctx context.Context,
		user domain.User,
	) error
	DeleteUser(
		ctx context.Context,
		userID uuid.UUID,
	) error
	ExistUserByLogin(
		ctx context.Context,
		login string,
	) (bool, error)
	ExistOtherUserByLogin(
		ctx context.Context,
		userID uuid.UUID,
		login string,
	) (bool, error)
}

func NewUsersService(
	usersRepository UsersRepository,
	hasher crypto_hasher.Hasher,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		hasher:          hasher,
	}
}
