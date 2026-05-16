package age_calculator

import (
	"testing"
	"time"
)

func TestAgeAt(t *testing.T) {
	type args struct {
		date time.Time
		asOf time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "ДР уже прошел",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 11, 0, 0, 0, 0, time.UTC),
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "ДР сегодня",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC),
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "ДР еще не наступил",
			args: args{
				date: time.Date(2000, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 1, 9, 0, 0, 0, 0, time.UTC),
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "ДР 29.02 в високосный год, еще не наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "ДР 29.02 в високосный год, уже прошел",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "ДР 29.02 в високосный год, еще не наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC),
			},
			want:    23,
			wantErr: false,
		},
		{
			name: "ДР 29.02 в високосный год, уже наступил",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "ДР 29.02 в високосный год, уже прошел",
			args: args{
				date: time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "время игнорируется в date",
			args: args{
				date: time.Date(2000, 5, 16, 23, 59, 59, 999999999, time.UTC),
				asOf: time.Date(2025, 5, 16, 0, 0, 0, 0, time.UTC),
			},
			want:    25,
			wantErr: false,
		},
		{
			name: "время игнорируется в asOf",
			args: args{
				date: time.Date(2000, 5, 17, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2025, 5, 16, 23, 59, 59, 999999999, time.UTC),
			},
			want:    24,
			wantErr: false,
		},
		// asOf раньше даты рождения
		{
			name: "asOf раньше date - ошибка",
			args: args{
				date: time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC),
				asOf: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AgeAt(tt.args.date, tt.args.asOf)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgeAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AgeAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
