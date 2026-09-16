package categories_service

import (
	"context"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type CategoriesService struct {
	categoriesRepository CategoriesRepository
}
type CategoriesRepository interface {
	CreateCategory(
		ctx context.Context,
		category domain.Category,
	) (domain.Category, error)
	GetCategories(
		ctx context.Context,
		userID uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Category, error)
	GetCategory(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
	) (domain.Category, error)
	DeleteCategory(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
	) error
	UpdateCategory(
		ctx context.Context,
		category domain.Category,
	) (domain.Category, error)
}

func NewCategoriesService(
	categoriesRepository CategoriesRepository,
) *CategoriesService {
	return &CategoriesService{
		categoriesRepository: categoriesRepository,
	}
}
