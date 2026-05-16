package age_calculator

import "time"

// hasHadBirthdayAt определяет, наступил ли день рождения на момент asOf
// в текущем году.
//
// Ожидается, что входные значения уже нормализованы по дате (без времени).
//
// Для дат 29 февраля используется следующая логика:
//   - в невисокосный год днем рождения считается 1 марта;
//   - в високосный год — 29 февраля.
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

// isLeapYear возвращает true, если указанный год является високосным.
//
// Год считается високосным, если:
//   - делится на 4;
//   - но не делится на 100;
//   - либо делится на 400.
func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	return year%4 == 0
}
