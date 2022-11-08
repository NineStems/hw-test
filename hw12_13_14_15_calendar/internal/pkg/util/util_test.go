package util

import (
	"testing"
	"time"
)

func TestCompareDateRange(t *testing.T) {
	type args struct {
		start     time.Time
		end       time.Time
		baseStart time.Time
		baseEnd   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "good check day",
			args: args{
				start:     time.Date(2022, 10, 30, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 10, 30, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 30, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: true,
		},
		{
			name: "bad check day",
			args: args{
				start:     time.Date(2022, 10, 29, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 10, 29, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 30, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: false,
		},
		{
			name: "good check week",
			args: args{
				start:     time.Date(2022, 10, 30, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 10, 30, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 24, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: true,
		},
		{
			name: "bad check week",
			args: args{
				start:     time.Date(2022, 10, 23, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 10, 23, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 24, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: false,
		},
		{
			name: "good check month",
			args: args{
				start:     time.Date(2022, 10, 30, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 10, 30, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: true,
		},
		{
			name: "bad check month",
			args: args{
				start:     time.Date(2022, 11, 30, 22, 0, 0, 0, time.Local),
				end:       time.Date(2022, 11, 30, 23, 0, 0, 0, time.Local),
				baseStart: time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local),
				baseEnd:   time.Date(2022, 10, 30, 23, 59, 59, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareDateRange(tt.args.start, tt.args.end, tt.args.baseStart, tt.args.baseEnd); got != tt.want {
				t.Errorf("CompareDateRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
