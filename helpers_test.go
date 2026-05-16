package age_calculator

import (
	"testing"
	"time"
)

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2000 - високосный",
			args: args{
				year: 2000,
			},
			want: true,
		},
		{
			name: "2020 - високосный",
			args: args{
				year: 2020,
			},
			want: true,
		},
		{
			name: "2024 - високосный",
			args: args{
				year: 2024,
			},
			want: true,
		},
		{
			name: "2019 - невисокосный",
			args: args{
				year: 2019,
			},
			want: false,
		},
		{
			name: "2021 - невисокосный",
			args: args{
				year: 2021,
			},
			want: false,
		},
		{
			name: "1900 - невисокосный",
			args: args{
				year: 1900,
			},
			want: false,
		},
		{
			name: "2100 - невисокосный",
			args: args{
				year: 2100,
			},
			want: false,
		},
		{
			name: "2200 - невисокосный",
			args: args{
				year: 2200,
			},
			want: false,
		},
		{
			name: "1600 - високосный",
			args: args{
				year: 1600,
			},
			want: true,
		},
		{
			name: "2400 - високосный",
			args: args{
				year: 2400,
			},
			want: true,
		},
		{
			name: "4 - високосный",
			args: args{
				year: 4,
			},
			want: true,
		},
		{
			name: "100 - невисокосный",
			args: args{
				year: 100,
			},
			want: false,
		},
		{
			name: "400 - невисокосный",
			args: args{
				year: 400,
			},
			want: true,
		},
		{
			name: "1 - невисокосный",
			args: args{
				year: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasHadBirthdayAt(t *testing.T) {
	type args struct {
		date time.Time
		asOf time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ДР уже прошел",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 11, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "ДР на текущую дату",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "ДР еще не прошел",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 9, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "ДР был месяц назад",
			args: args{
				date: time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "до ДР еще месяц",
			args: args{
				date: time.Date(2000, 12, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 11, 30, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "ДР в високосный день 29 февраля, еще не наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "ДР 28 февраля в невисокосный год, считается, что еще не наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "ДР 1 марта в невисокосный год, уже наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "ДР 31 декабря, сейчас 1 января",
			args: args{
				date: time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasHadBirthdayAt(tt.args.date, tt.args.asOf); got != tt.want {
				t.Errorf("hasHadBirthdayAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
