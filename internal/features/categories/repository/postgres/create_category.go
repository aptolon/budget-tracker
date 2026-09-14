package categories_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

func (r *CategoriesRepository) CreateCategory(
	ctx context.Context,
	category domain.Category,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO categories (id, version, user_id, title)
		VALUES ($1, $2, $3, $4);
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		category.ID,
		category.Version,
		category.UserID,
		category.Title,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return fmt.Errorf("create category: %w", core_errors.ErrConflict)
		}
		return fmt.Errorf("create category: %w", err)
	}

	return nil
}
