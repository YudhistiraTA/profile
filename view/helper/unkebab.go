package helper

import (
	"strings"
)

// Unkebab converts a kebab case string into a normal string using spaces
func Unkebab(input string) string {
	return strings.ReplaceAll(input, "-", " ")
}
