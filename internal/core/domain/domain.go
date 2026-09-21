package domain

import (
	"fmt"
	"strings"
	"unicode"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func NormalizeString(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(strings.ToLower(str))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
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
