package categories_transport_http

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
	UserID  uuid.UUID `json:"user_id"`
	Title   string    `json:"title"`
}

type CategoryRequest struct {
	Title string `json:"title"   validate:"required,min=3,max=32"`
}

func categoryResponseFromDomain(category domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:      category.ID,
		Version: category.Version,
		UserID:  category.UserID,
		Title:   category.Title,
	}
}

func categoriesResponseFromDomains(categories []domain.Category) []CategoryResponse {
	categoriesResponse := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		categoriesResponse[i] = categoryResponseFromDomain(category)
	}
	return categoriesResponse
}
