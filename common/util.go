package common

import (
	"strings"
)

func NormalizeString(value string) string {
	input := strings.ToLower(value)
	return strings.ReplaceAll(input, " ", "-")
}
