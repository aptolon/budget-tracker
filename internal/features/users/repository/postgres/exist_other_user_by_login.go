package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *UsersRepository) ExistOtherUserByLogin(
	ctx context.Context,
	userID uuid.UUID,
	login string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE login = $1
			AND id != $2
		);
	`

	var exists bool

	if err := r.pool.QueryRow(
		ctx,
		query,
		login,
		userID,
	).Scan(&exists); err != nil {
		return false, fmt.Errorf("check user exists: %w", err)
	}

	return exists, nil
}
