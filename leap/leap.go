// This is a side exercise in exercism go track
package leap

// IsLeapYear tells if the given year is a Leap year or not.
func IsLeapYear(year int) bool {
	// every year that is evenly divisible by 4
	if year % 4 == 0 {
		// except every year that is evenly divisible by 100
		if year % 100 == 0 {
			// unless the year is also evenly divisible by 400
			if year % 400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return false
}
