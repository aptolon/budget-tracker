package domain

import (
	"github.com/google/uuid"
)

type Subcategory struct {
	ID         uuid.UUID
	Version    int
	CategoryID uuid.UUID
	Title      string
}

func NewSubcategory(
	id uuid.UUID,
	version int,
	categoryID uuid.UUID,
	title string,
) Subcategory {
	return Subcategory{
		ID:         id,
		Version:    version,
		CategoryID: categoryID,
		Title:      NormalizeString(title),
	}
}

func CreateSubcategory(
	categoryID uuid.UUID,
	title string,
) Subcategory {
	var (
		id      = uuid.New()
		version = 1
	)
	return NewSubcategory(
		id,
		version,
		categoryID,
		title,
	)
}
