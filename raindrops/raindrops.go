package raindrops

import "strconv"

// Convert function takes an integer and return a string on the basis conditions
func Convert(input int) string {
	result := ""

	if isFactorOf(input, 3) {
		result += "Pling"
	}

	if isFactorOf(input, 5) {
		result += "Plang"
	}

	if isFactorOf(input, 7) {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(input)
	}

	return result
}

// isFactorOf function takes two integers and returns boolean weather the second input is factor or not
func isFactorOf(dividend, divisor int) bool {
	return dividend%divisor == 0
}
