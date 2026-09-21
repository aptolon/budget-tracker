package subcategories_postgres_repository

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (r *SubcategoriesRepository) GetSubcategories(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT s.id, s.version, s.category_id, s.title
		FROM subcategories s
		JOIN categories c ON s.category_id = c.id
		WHERE c.user_id = $2 AND s.category_id = $1
		ORDER BY s.id ASC
		LIMIT $3
		OFFSET $4;
	`
	rows, err := r.pool.Query(
		ctx,
		query,
		categoryID,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select subcategories: %w", err)
	}
	defer rows.Close()

	var subcategoryModels []SubcategoryModel
	for rows.Next() {
		var subcategoryModel SubcategoryModel
		if err := subcategoryModel.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan subcategories: %w", err)
		}
		subcategoryModels = append(subcategoryModels, subcategoryModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	subcategoryDomains := modelsToDomains(subcategoryModels)

	return subcategoryDomains, nil
}
