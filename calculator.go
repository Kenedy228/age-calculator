package age_calculator

import (
	"fmt"
	"time"
)

func AgeAt(date, asOf time.Time) (int, error) {
	date = normalizeCalendarDate(date)
	asOf = normalizeCalendarDate(asOf)

	if date.After(asOf) {
		return 0, fmt.Errorf("%w", ErrDateAfterAsOf)
	}

	rawAge := asOf.Year() - date.Year()

	if !hasHadBirthdayAt(date, asOf) {
		rawAge--
	}

	return rawAge, nil
}
