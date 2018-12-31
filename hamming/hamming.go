// Package hamming is an exercise from exercism go track
package hamming

import "fmt"

// Distance Function takes two strings and returns hamming distance, or an error if lengths are not equal
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("%s and %s are not equal in length", a, b)
	}

	// it will be O(n)
	var count int
	for index := range a {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil
}
