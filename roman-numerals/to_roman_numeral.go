package romannumerals

import "fmt"

// ToRomanNumeral function takes arabic int number and returns its roman numeral string.
func ToRomanNumeral(arabic int) (string, error) {

	if arabic > 3000 || arabic <= 0 {
		return "", fmt.Errorf("%d is bigger than 3000", arabic)
	}

	result := ""
	decimalValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanValues := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for index, decimal := range decimalValues {
		for decimal <= arabic {
			result += romanValues[index]
			arabic -= decimal
		}
	}

	return result, nil
}
