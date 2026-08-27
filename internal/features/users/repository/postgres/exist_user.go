package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) ExistUserByLogin(
	ctx context.Context,
	login string,
) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE login = $1
		);
	`

	var exists bool

	if err := r.pool.QueryRow(
		ctx,
		query,
		login,
	).Scan(&exists); err != nil {
		return false, fmt.Errorf("check user exists: %w", err)
	}

	return exists, nil
}
