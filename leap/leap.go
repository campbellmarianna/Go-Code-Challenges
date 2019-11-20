// Package leap reports if a given year is a leap year.
package leap

// IsLeapYear returns true if given year is a leap year and false if it is not.
func IsLeapYear(year int) bool {
	// Currently, checking if given year meets the qualifications of a leap year.
	if year % 4 == 0 {
		if year % 100 != 0 {
				return true
		}
		if year % 400 == 0 {
				return true
		}
	}
	return false
}