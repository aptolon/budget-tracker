package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
)

func (r *UsersRepository) UpdateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE users
	SET 
		version = version + 1,
		role = $3,
		login = $4,
		password_hash = $5
	WHERE id = $1 AND version = $2
	RETURNING id, version, role, login, password_hash;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.ID,
		user.Version,
		user.Role,
		user.Login,
		user.PasswordHash,
	)
	var userModel UserModel
	if err := userModel.Scan(row); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := modelToDomain(userModel)

	return userDomain, nil
}
