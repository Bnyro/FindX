package utilities

import "strings"

func TakeN(text string, count int) string {
	if len(text) < count {
		return text
	}
	return text[:count] + " ..."
}

func IsBlank(text string) bool {
	return strings.TrimSpace(text) == ""
}
