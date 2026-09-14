package categories_postgres_repository

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (r *CategoriesRepository) GetCategories(
	ctx context.Context,
	userID uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, user_id, title
		FROM categories
		WHERE user_id = $3
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
	`
	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("select categories: %w", err)
	}
	defer rows.Close()

	var categoryModels []CategoryModel
	for rows.Next() {
		var categoryModel CategoryModel
		if err := categoryModel.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan categories: %w", err)
		}
		categoryModels = append(categoryModels, categoryModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	categoryDomains := modelsToDomains(categoryModels)

	return categoryDomains, nil
}
