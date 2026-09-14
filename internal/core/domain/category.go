package domain

import (
	"fmt"
	"strings"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
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
		Title:   normalizeTitle(title),
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

func ValidateTitle(title string) error {
	titleLen := len([]rune(title))
	if titleLen < 3 || titleLen > 32 {
		return fmt.Errorf(
			"invalid `title` len %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func normalizeTitle(title string) string {
	if title == "" {
		return ""
	}

	s := strings.ToLower(title)
	return strings.ToUpper(s[:1]) + s[1:]
}
