package categories_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

func (r *CategoriesRepository) DeleteCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM categories
		WHERE id = $1 AND user_id = $2;
	`
	cmdTag, err := r.pool.Exec(ctx, query, categoryID, userID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("category with id='%s' and user_id='%s': %w", categoryID, userID, core_errors.ErrNotFound)
	}

	return nil
}
