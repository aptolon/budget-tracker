package domain

import (
	"github.com/google/uuid"
)

type Category struct {
	ID      uuid.UUID
	Version int
	UserID  uuid.UUID
	Title   string
}

func NewCategory(
	id uuid.UUID,
	version int,
	userID uuid.UUID,
	title string,

) Category {
	return Category{
		ID:      id,
		Version: version,
		UserID:  userID,
		Title:   NormalizeString(title),
	}

}

func CreateCategory(
	userID uuid.UUID,
	title string,
) Category {
	var (
		id      = uuid.New()
		version = 1
	)
	return NewCategory(
		id,
		version,
		userID,
		title,
	)
}
