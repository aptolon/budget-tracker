package subcategories_transport_http

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type SubcategoryResponse struct {
	ID         uuid.UUID `json:"id"`
	Version    int       `json:"version"`
	CategoryID uuid.UUID `json:"category_id"`
	Title      string    `json:"title"`
}

type SubcategoryRequest struct {
	Title string `json:"title"   validate:"required,min=3,max=32"`
}

func subcategoryResponseFromDomain(subcategory domain.Subcategory) SubcategoryResponse {
	return SubcategoryResponse{
		ID:         subcategory.ID,
		Version:    subcategory.Version,
		CategoryID: subcategory.CategoryID,
		Title:      subcategory.Title,
	}
}

func subcategoriesResponseFromDomains(subcategories []domain.Subcategory) []SubcategoryResponse {
	subcategoriesResponse := make([]SubcategoryResponse, len(subcategories))
	for i, subcategory := range subcategories {
		subcategoriesResponse[i] = subcategoryResponseFromDomain(subcategory)
	}
	return subcategoriesResponse
}
