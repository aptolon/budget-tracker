package users_postgres_repository

import (
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
	"github.com/google/uuid"
)

type UserModel struct {
	ID      uuid.UUID
	Version int
	Role    string
	Login   string
}

func (m *UserModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.Role,
		&m.Login,
	)
}
