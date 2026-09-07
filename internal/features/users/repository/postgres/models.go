package users_postgres_repository

import (
	"github.com/aptolon/budget-tracker/internal/core/domain"
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

type UserModel struct {
	ID      uuid.UUID
	Version int
	Role    string
	Login   string

	PasswordHash string
}

func (m *UserModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.Role,
		&m.Login,
		&m.PasswordHash,
	)
}

func modelToDomain(model UserModel) domain.User {
	return domain.NewUser(
		model.ID,
		model.Version,
		model.Role,
		model.Login,
		model.PasswordHash,
	)
}
func modelsToDomains(models []UserModel) []domain.User {
	domains := make([]domain.User, len(models))
	for i, model := range models {
		domains[i] = modelToDomain(model)
	}
	return domains
}
