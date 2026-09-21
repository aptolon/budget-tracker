package subcategories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubcategoriesService) CreateSubcategory(
	ctx context.Context,
	userID uuid.UUID,
	title string,
	categoryID uuid.UUID,
) (domain.Subcategory, error) {
	if err := domain.ValidateTitle(title); err != nil {
		return domain.Subcategory{}, fmt.Errorf("validate title: %w", err)
	}
	subcategory := domain.CreateSubcategory(
		categoryID,
		title,
	)
	subcategory, err := s.subcategoriesRepository.CreateSubcategory(ctx, subcategory, userID)
	if err != nil {
		return domain.Subcategory{}, fmt.Errorf("create subcategory: %w", err)
	}
	return subcategory, nil
}
