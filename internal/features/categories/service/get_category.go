package categories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *CategoriesService) GetCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
) (domain.Category, error) {
	category, err := s.categoriesRepository.GetCategory(
		ctx,
		categoryID,
		userID,
	)
	if err != nil {
		return domain.Category{}, fmt.Errorf(
			"get category from repository: %w",
			err)
	}
	return category, nil
}
