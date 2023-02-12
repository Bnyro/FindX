package utilities

import (
	"regexp"
	"strings"
)

var IsAlphabetic = regexp.MustCompile(`^[A-Za-z]+$`).MatchString

func TakeN(text string, count int) string {
	if len(text) < count {
		return text
	}
	return text[:count] + " ..."
}

func IsBlank(text string) bool {
	return strings.TrimSpace(text) == ""
}
