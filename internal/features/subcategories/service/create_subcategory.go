package subcategories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubcategoriesService) CreateSubcategory(
	ctx context.Context,
	categoryID uuid.UUID,
	title string,
	userID uuid.UUID,
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
