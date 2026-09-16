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
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO categories (id, version, user_id, title)
		VALUES ($1, $2, $3, $4)
		RETURNING id, version, user_id, title;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		category.ID,
		category.Version,
		category.UserID,
		category.Title,
	)

	var categoryModel CategoryModel
	if err := categoryModel.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return domain.Category{}, fmt.Errorf(
				"category with title '%s' already exists: %w",
				category.Title,
				core_errors.ErrConflict,
			)
		}
		return domain.Category{}, fmt.Errorf("scan error: %w", err)
	}

	categoryDomain := modelToDomain(categoryModel)

	return categoryDomain, nil
}
