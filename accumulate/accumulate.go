package accumulate

// Accumulate function takes an string array and an operation function, thus returns result array
func Accumulate(input []string, operation func(string) string) []string {
	result := make([]string, len(input))
	for index, value := range input {
		result[index] = operation(value)
	}

	return result
}
