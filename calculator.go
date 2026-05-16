package age_calculator

import (
	"fmt"
	"time"
)

// AgeAt вычисляет возраст в полных годах на дату asOf.
//
// Оба аргумента нормализуются до календарной даты (без учета времени и
// часового пояса), поэтому результат зависит только от года, месяца и дня.
//
// Если дата рождения позже даты расчета, возвращается ошибка ErrDateAfterAsOf.
//
// Для дат 29 февраля используется следующая логика:
//   - в невисокосный год днем рождения считается 1 марта;
//   - в високосный год — 29 февраля.
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
