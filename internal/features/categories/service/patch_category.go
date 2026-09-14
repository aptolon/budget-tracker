package categories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *CategoriesService) PatchCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
	title string,
) (domain.Category, error) {
	category, err := s.categoriesRepository.GetCategory(ctx, categoryID, userID)
	if err != nil {
		return domain.Category{}, fmt.Errorf("get category: %w", err)
	}
	category.Title = title
	err = s.categoriesRepository.UpdateCategory(ctx, category)
	if err != nil {
		return domain.Category{}, fmt.Errorf("update category: %w", err)
	}
	category.Version++
	return category, nil
}
