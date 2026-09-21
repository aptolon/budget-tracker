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

func (r *SubcategoriesRepository) CreateSubcategory(
	ctx context.Context,
	subcategory domain.Subcategory,
	userID uuid.UUID,
) (domain.Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO subcategories (id, version, category_id, title)
			SELECT $1,$2, c.id, $4
			FROM categories c
				WHERE c.id = $3
				AND c.user_id = $5
		RETURNING id, version, category_id, title;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		subcategory.ID,
		subcategory.Version,
		subcategory.CategoryID,
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
