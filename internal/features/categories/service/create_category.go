package categories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *CategoriesService) CreateCategory(
	ctx context.Context,
	userID uuid.UUID,
	title string,
) (domain.Category, error) {
	if err := domain.ValidateTitle(title); err != nil {
		return domain.Category{}, fmt.Errorf("validate title: %w", err)
	}
	category := domain.CreateCategory(
		userID,
		title,
	)
	if err := s.categoriesRepository.CreateCategory(ctx, category); err != nil {
		return domain.Category{}, fmt.Errorf("create category: %w", err)
	}
	return category, nil
}
