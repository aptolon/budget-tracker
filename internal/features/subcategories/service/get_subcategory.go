package subcategories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

func (s *SubcategoriesService) GetSubcategory(
	ctx context.Context,
	subcategoryID uuid.UUID,
	userID uuid.UUID,
) (domain.Subcategory, error) {
	subcategory, err := s.subcategoriesRepository.GetSubcategory(
		ctx,
		subcategoryID,
		userID,
	)
	if err != nil {
		return domain.Subcategory{}, fmt.Errorf(
			"get subcategory from repository: %w",
			err)
	}
	return subcategory, nil
}
