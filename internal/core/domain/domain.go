package domain

import (
	"strings"
	"unicode"
)

func NormalizeString(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(strings.ToLower(str))
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
