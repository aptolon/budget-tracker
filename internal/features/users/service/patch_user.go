package users_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	userID uuid.UUID,
	login *string,
	role *string,
	password *string,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if login != nil {
		if err := domain.ValidateLogin(*login); err != nil {
			return domain.User{}, fmt.Errorf("patch user: %w", err)
		}
		exists, err := s.usersRepository.ExistOtherUserByLogin(ctx, userID, *login)
		if err != nil {
			return domain.User{}, fmt.Errorf("patch user: %w", err)
		}
		if exists {
			return domain.User{}, core_errors.ErrLoginTaken
		}
		user.Login = *login
	}
	if password != nil {
		if err := domain.ValidatePassword(*password); err != nil {
			return domain.User{}, fmt.Errorf("patch user: %w", err)
		}
		hash, err := s.hasher.Hash(*password)
		if err != nil {
			return domain.User{}, fmt.Errorf("hash password: %w", err)
		}
		user.PasswordHash = hash
	}
	if role != nil {

		if err := domain.ValidateRole(*role); err != nil {
			return domain.User{}, fmt.Errorf("patch user: %w", err)
		}
		user.Role = *role
	}
	err = s.usersRepository.UpdateUser(
		ctx,
		user,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}
	user.Version++
	return user, nil

}
