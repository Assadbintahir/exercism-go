package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram function takes a string and return a boolean if its isogram or not
func IsIsogram(input string) bool {
	input = strings.ToUpper(input)
	var seen = map[rune]bool{}
	for _, value := range input {
		if !unicode.IsLetter(value) {
			continue
		}
		if seen[value] {
			return false
		}
		seen[value] = true
	}
	return true
}
