package categories_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

func (r *CategoriesRepository) UpdateCategory(
	ctx context.Context,
	category domain.Category,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE categories
	SET 
		version = version + 1,
		title = $4
	WHERE id = $1 AND version = $2 AND user_id = $3;
	`

	result, err := r.pool.Exec(
		ctx,
		query,
		category.ID,
		category.Version,
		category.UserID,
		category.Title,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return fmt.Errorf("update category: %w", core_errors.ErrConflict)
		}

		return fmt.Errorf("update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return core_errors.ErrConflict
	}
	return nil
}
