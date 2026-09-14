package categories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *CategoriesService) DeleteCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
) error {
	err := s.categoriesRepository.DeleteCategory(
		ctx,
		categoryID,
		userID,
	)
	if err != nil {
		return fmt.Errorf(
			"delete category from repository: %w",
			err)
	}

	return nil
}
