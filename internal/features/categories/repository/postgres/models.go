package categories_postgres_repository

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

type CategoryModel struct {
	ID      uuid.UUID
	Version int
	UserID  uuid.UUID
	Title   string
}

func (m *CategoryModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.UserID,
		&m.Title,
	)
}

func modelToDomain(model CategoryModel) domain.Category {
	return domain.NewCategory(
		model.ID,
		model.Version,
		model.UserID,
		model.Title,
	)
}
func modelsToDomains(models []CategoryModel) []domain.Category {
	domains := make([]domain.Category, len(models))
	for i, model := range models {
		domains[i] = modelToDomain(model)
	}
	return domains
}
