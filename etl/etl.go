package etl

import "strings"

// Transform function takes map of strings with int keys and returns a map of int with string keys
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for key, value := range input {
		for _, element := range value {
			result[strings.ToLower(element)] = key
		}
	}
	return result
}
