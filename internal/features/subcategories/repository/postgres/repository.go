package subcategories_postgres_repository

import (
	core_postgres_pool "github.com/aptolon/budget-tracker/internal/core/repository/postgres/pool"
)

type SubcategoriesRepository struct {
	pool core_postgres_pool.Pool
}

func NewSubcategoriesRepository(
	pool core_postgres_pool.Pool,
) *SubcategoriesRepository {
	return &SubcategoriesRepository{
		pool: pool,
	}
}
