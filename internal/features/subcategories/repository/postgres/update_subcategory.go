package subcategories_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

func (r *SubcategoriesRepository) UpdateSubcategory(
	ctx context.Context,
	subcategory domain.Subcategory,
	userID uuid.UUID,
) (domain.Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE subcategories
	SET 
		version = version + 1,
		title = $3
	WHERE id = $1 AND version = $2 
		AND category_id IN (SELECT id FROM categories WHERE user_id = $4)
	RETURNING id, version, category_id, title;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		subcategory.ID,
		subcategory.Version,
		subcategory.Title,
		userID,
	)

	var subcategoryModel SubcategoryModel
	if err := subcategoryModel.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return domain.Subcategory{}, fmt.Errorf(
				"subcategory with title '%s' already exists: %w",
				subcategory.Title,
				core_errors.ErrConflict,
			)
		}
		return domain.Subcategory{}, fmt.Errorf("scan error: %w", err)
	}

	subcategoryDomain := modelToDomain(subcategoryModel)

	return subcategoryDomain, nil
}
