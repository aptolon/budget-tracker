package users_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *UsersService) DeleteUser(
	ctx context.Context,
	userID uuid.UUID,
) error {
	err := s.usersRepository.DeleteUser(
		ctx,
		userID,
	)
	if err != nil {
		return fmt.Errorf(
			"delete user from repository: %w",
			err)
	}

	return nil
}
