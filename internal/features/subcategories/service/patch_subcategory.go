package subcategories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubcategoriesService) PatchSubcategory(
	ctx context.Context,
	subcategoryID uuid.UUID,
	title string,
	userID uuid.UUID,
) (domain.Subcategory, error) {
	subcategory, err := s.subcategoriesRepository.GetSubcategory(ctx, subcategoryID, userID)
	if err != nil {
		return domain.Subcategory{}, fmt.Errorf("get subcategory: %w", err)
	}
	subcategory.Title = domain.NormalizeString(title)
	subcategory, err = s.subcategoriesRepository.UpdateSubcategory(ctx, subcategory, userID)
	if err != nil {
		return domain.Subcategory{}, fmt.Errorf("update subcategory: %w", err)
	}
	return subcategory, nil
}
