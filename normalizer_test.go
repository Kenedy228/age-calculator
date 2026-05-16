package age_calculator

import (
	"reflect"
	"testing"
	"time"
)

func Test_normalizeCalendarDate(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "дата со временем",
			args: args{
				date: time.Date(2025, 5, 16, 15, 30, 45, 123, time.UTC),
			},
			want: time.Date(2025, 5, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "уже нормализованная дата",
			args: args{
				date: time.Date(2025, 5, 16, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2025, 5, 16, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeCalendarDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeCalendarDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
