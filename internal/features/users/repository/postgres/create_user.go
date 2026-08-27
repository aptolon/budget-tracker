package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO users (id, version, role, login, password_hash)
		VALUES ($1, $2, $3, $4, $5);
	`

	_, err := r.pool.Exec(
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
			return fmt.Errorf("create user: %w", core_errors.ErrLoginTaken)
		}

		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
