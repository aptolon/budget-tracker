package subcategories_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

func (r *SubcategoriesRepository) DeleteSubcategory(
	ctx context.Context,
	subcategoryID uuid.UUID,
	userID uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
        DELETE FROM subcategories s
        USING categories c
        WHERE s.category_id = c.id
          AND s.id = $1
          AND c.user_id = $2;
    `
	cmdTag, err := r.pool.Exec(ctx, query, subcategoryID, userID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("subcategory with id='%s' and user_id='%s': %w", subcategoryID, userID, core_errors.ErrNotFound)
	}

	return nil
}
