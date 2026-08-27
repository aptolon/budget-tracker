package core_pgx_pool

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

type pgxCommandTag struct {
	pgconn.CommandTag
}
type pgxRow struct {
	pgx.Row
}
type pgxRows struct {
	pgx.Rows
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKeyErrorCode = "23503"
		pgxUniqueViolationErrorCode    = "23505"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgxUniqueViolationErrorCode:
			return fmt.Errorf(
				"%w: %v",
				core_postgres_pool.ErrUniqueViolation,
				err,
			)

		case pgxViolatesForeignKeyErrorCode:
			return fmt.Errorf(
				"%w: %v",
				core_postgres_pool.ErrViolatesForeignKey,
				err,
			)
		}

	}
	return fmt.Errorf(
		"%w: %v",
		err,
		core_postgres_pool.ErrUnknown,
	)
}
