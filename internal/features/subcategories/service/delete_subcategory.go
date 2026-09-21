package subcategories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *SubcategoriesService) DeleteSubcategory(
	ctx context.Context,
	subcategoryID uuid.UUID,
	userID uuid.UUID,
) error {
	err := s.subcategoriesRepository.DeleteSubcategory(
		ctx,
		subcategoryID,
		userID,
	)
	if err != nil {
		return fmt.Errorf(
			"delete subcategory from repository: %w",
			err)
	}

	return nil
}
