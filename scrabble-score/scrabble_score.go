package scrabble

import "unicode"

// Score function takes a string and returns its int scrabble score
func Score(input string) int {
	result := 0

	for _, value := range input {
		result += letterScrabbleValue(unicode.ToUpper(value))
	}

	return result
}

// letterScrabbleValue function takes a byte letter and returns its int scrabble value
func letterScrabbleValue(letter rune) int {

	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}

	return 0
}
