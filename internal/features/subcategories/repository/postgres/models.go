package subcategories_postgres_repository

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

type SubcategoryModel struct {
	ID         uuid.UUID
	Version    int
	CategoryID uuid.UUID
	Title      string
}

func (m *SubcategoryModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.CategoryID,
		&m.Title,
	)
}

func modelToDomain(model SubcategoryModel) domain.Subcategory {
	return domain.NewSubcategory(
		model.ID,
		model.Version,
		model.CategoryID,
		model.Title,
	)
}
func modelsToDomains(models []SubcategoryModel) []domain.Subcategory {
	domains := make([]domain.Subcategory, len(models))
	for i, model := range models {
		domains[i] = modelToDomain(model)
	}
	return domains
}
