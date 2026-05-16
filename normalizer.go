package age_calculator

import "time"

// normalizeCalendarDate возвращает копию времени, в которой обнулены часы,
// минуты, секунды и наносекунды, а часовой пояс установлен в UTC.
//
// Функция сохраняет только календарную дату (год, месяц, день) и игнорирует
// исходное время и часовой пояс. Возвращаемое значение не представляет тот же
// момент времени, а используется для корректных сравнений по дате.
func normalizeCalendarDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}
