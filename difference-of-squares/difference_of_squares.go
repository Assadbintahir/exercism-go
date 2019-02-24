package diffsquares

// SquareOfSum function takes an integer and returns integer square of sum upto that number
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares function takes an integer and returns integer sum of squares upto that number
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

// Difference function takes an integer and returns difference of square of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
