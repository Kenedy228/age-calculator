package age_calculator

import "time"

func hasHadBirthdayAt(date, asOf time.Time) bool {
	if date.Month() == 2 && date.Day() == 29 && !isLeapYear(asOf.Year()) {
		pseudoDob := time.Date(asOf.Year(), 2, 28, 0, 0, 0, 0, time.UTC)
		return asOf.After(pseudoDob)
	}

	if asOf.Month() > date.Month() {
		return true
	}
	if asOf.Month() == date.Month() && asOf.Day() >= date.Day() {
		return true
	}

	return false
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	return year%4 == 0
}
