package subcategories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

func (s *SubcategoriesService) GetSubcategories(
	ctx context.Context,
	categoryID uuid.UUID,
	userID uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Subcategory, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument)
	}

	subcategories, err := s.subcategoriesRepository.GetSubcategories(
		ctx,
		categoryID,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get subcategories from repository: %w",
			err)
	}

	return subcategories, nil
}
