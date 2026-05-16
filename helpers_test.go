package age_calculator

import "testing"

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
