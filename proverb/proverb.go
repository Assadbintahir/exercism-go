// Package proverb is an exercise from exercism Golang track
package proverb

import "fmt"

// Proverb function takes slice of strings and return appropriate slice of strings
func Proverb(rhyme []string) []string {
	lastPieceString := "And all for the want of a %s."
	pieceString := "For want of a %s the %s was lost."
	result := []string{}

	for index := range rhyme {
		if (index + 1) == len(rhyme) {
			result = append(result, fmt.Sprintf(lastPieceString, rhyme[0]))
		} else {
			result = append(result, fmt.Sprintf(pieceString, rhyme[index], rhyme[index+1]))
		}
	}

	return result
}
