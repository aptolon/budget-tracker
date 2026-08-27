package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUserByLogin(
	ctx context.Context,
	login string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, role, login, password_hash
		FROM users
		WHERE login = $1;
	`

	var user domain.User

	if err := r.pool.QueryRow(
		ctx,
		query,
		login,
	).Scan(
		&user.ID,
		&user.Version,
		&user.Role,
		&user.Login,
		&user.PasswordHash,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"get user by login: %w",
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf(
			"get user by login: %w",
			err,
		)
	}

	return user, nil
}
