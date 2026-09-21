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

func (r *SubcategoriesRepository) GetSubcategory(
	ctx context.Context,
	subcategoryID uuid.UUID,
	userID uuid.UUID,
) (domain.Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT s.id, s.version, s.category_id, s.title
		FROM subcategories s
		JOIN categories c ON s.category_id = c.id
		WHERE s.id = $1
		AND c.user_id = $2;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		subcategoryID,
		userID,
	)

	var subcategoryModel SubcategoryModel
	if err := subcategoryModel.Scan(row); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Subcategory{}, fmt.Errorf(
				"subcategory with id='%s': %w",
				subcategoryID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Subcategory{}, fmt.Errorf("scan error: %w", err)
	}

	subcategoryDomain := modelToDomain(subcategoryModel)

	return subcategoryDomain, nil
}
