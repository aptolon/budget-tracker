package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) UpdateUser(
	ctx context.Context,
	user domain.User,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE users
	SET 
		version = version + 1,
		role = $3,
		login = $4,
		password_hash = $5
	WHERE id = $1 AND version = $2;
	`

	result, err := r.pool.Exec(
		ctx,
		query,
		user.ID,
		user.Version,
		user.Role,
		user.Login,
		user.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return fmt.Errorf("update user: %w", core_errors.ErrLoginTaken)
		}

		return fmt.Errorf("update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return core_errors.ErrConflict
	}
	return nil
}
