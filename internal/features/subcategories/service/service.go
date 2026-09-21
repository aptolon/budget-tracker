package subcategories_service

import (
	"context"

	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type SubcategoriesService struct {
	subcategoriesRepository SubcategoriesRepository
}
type SubcategoriesRepository interface {
	CreateSubcategory(
		ctx context.Context,
		subcategory domain.Subcategory,
		userID uuid.UUID,
	) (domain.Subcategory, error)
	GetSubcategories(
		ctx context.Context,
		categoryID uuid.UUID,
		userID uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Subcategory, error)
	GetSubcategory(
		ctx context.Context,
		subcategoryID uuid.UUID,
		userID uuid.UUID,
	) (domain.Subcategory, error)
	DeleteSubcategory(
		ctx context.Context,
		subcategoryID uuid.UUID,
		userID uuid.UUID,
	) error
	UpdateSubcategory(
		ctx context.Context,
		subcategory domain.Subcategory,
		userID uuid.UUID,
	) (domain.Subcategory, error)
}

func NewSubcategoriesService(
	subcategoriesRepository SubcategoriesRepository,
) *SubcategoriesService {
	return &SubcategoriesService{
		subcategoriesRepository: subcategoriesRepository,
	}
}
