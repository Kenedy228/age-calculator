# age_calculator

[![Go Version](https://img.shields.io/badge/go-1.26.2-blue)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/Kenedy228/age-calculator.svg)](https://pkg.go.dev/github.com/Kenedy228/age-calculator)
[![CI](https://github.com/Kenedy228/age-calculator/actions/workflows/go.yml/badge.svg)](https://github.com/Kenedy228/age-calculator/actions)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

Библиотека на Go для вычисления возраста в полных годах по календарной дате.

Библиотека работает только с **датой (год, месяц, день)** и игнорирует время и часовые пояса.

---

## ✨ Возможности

- вычисление возраста в полных годах;
- корректная обработка високосных лет;
- поддержка даты рождения 29 февраля:
    - в високосный год — 29 февраля;
    - в невисокосный — 1 марта (как в РФ).
- независимость от времени и таймзоны;
- простой и предсказуемый API.

---

## 📦 Установка

```bash
go get github.com/Kenedy228/age-calculator
```

## 🚀 Пример использования

```go
package main

import (
    "fmt"
    "time"

    "github.com/Kenedy228/age-calculator"
)

func main() {
    dob := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
    asOf := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)

    age, err := age_calculator.AgeAt(dob, asOf)
    if err != nil {
        panic(err)
    }

    fmt.Println(age) // 25
}
```

---

## ⚠️ Особенности поведения

- учитывается только календарная дата (время отбрасывается);
- все даты внутри нормализуются;
- если дата рождения позже даты расчета — возвращается ошибка.

---

## 📖 API

```go
AgeAt(date, asOf time.Time) (int, error)
```

Возвращает возраст в полных годах на указанную дату.

---

## 🧪 Тестирование

Локально:

```bash
go test ./...
```

---

## 📄 Лицензия

MIT License.