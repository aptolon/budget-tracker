package categories_service

import (
	"context"
	"fmt"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
	"github.com/google/uuid"
)

func (s *CategoriesService) GetCategories(
	ctx context.Context,
	userID uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Category, error) {
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

	categories, err := s.categoriesRepository.GetCategories(
		ctx,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get categories from repository: %w",
			err)
	}

	return categories, nil
}
