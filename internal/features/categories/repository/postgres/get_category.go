package categories_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

func (r *CategoriesRepository) GetCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, user_id, title
		FROM categories
		WHERE id = $1 AND user_id = $2;
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		categoryID,
		userID,
	)

	var categoryModel CategoryModel
	if err := categoryModel.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Category{}, fmt.Errorf(
				"category with id='%s': %w",
				categoryID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Category{}, fmt.Errorf("scan error: %w", err)
	}

	categoryDomain := modelToDomain(categoryModel)

	return categoryDomain, nil
}
