package age_calculator

import "errors"

// ErrDateAfterAsOf возвращается функцией AgeAt,
// если дата рождения находится позже даты расчета возраста.
var ErrDateAfterAsOf = errors.New("дата рождения позже даты расчета возраста")
