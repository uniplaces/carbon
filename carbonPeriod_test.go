package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	start1, _ := Create(2022, time.April, 1, 23, 0, 0, 0, "UTC")
	end1, _ := Create(2022, time.April, 30, 23, 0, 0, 0, "UTC")

	days := 7

	wantTime := make([]*Carbon, 0)
	want1a, _ := Create(2022, time.April, 1, 23, 0, 0, 0, "UTC")
	want1b, _ := Create(2022, time.April, 8, 23, 0, 0, 0, "UTC")
	want1c, _ := Create(2022, time.April, 15, 23, 0, 0, 0, "UTC")
	want1d, _ := Create(2022, time.April, 22, 23, 0, 0, 0, "UTC")
	want1e, _ := Create(2022, time.April, 29, 23, 0, 0, 0, "UTC")
	wantTime = append(wantTime, want1a)
	wantTime = append(wantTime, want1b)
	wantTime = append(wantTime, want1c)
	wantTime = append(wantTime, want1d)
	wantTime = append(wantTime, want1e)

	type args struct {
		start *Carbon
		days  int
		end   *Carbon
	}

	type want struct {
		wantTime []*Carbon
		error
	}

	tests := []struct {
		name string
		args args
		//want []*Carbon
		want
	}{
		{
			name: "normal",
			args: args{
				start: start1,
				days:  days,
				end:   end1,
			},
			want: want{
				wantTime: wantTime,
				error:    nil,
			},
		},
		{
			name: "date to is before date start",
			args: args{
				start: end1,
				days:  1,
				end:   start1,
			},
			want: want{
				wantTime: nil,
				error:    ErrEndMustBeAfterStart,
			},
		},
		{
			name: "day is zero",
			args: args{
				start: start1,
				days:  0,
				end:   end1,
			},
			want: want{
				wantTime: nil,
				error:    ErrDayAtLeast1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Period(tt.args.start, tt.args.days, tt.args.end)
			assert.Equal(t, tt.want.error, err)

			for key, val := range got {
				assert.Equal(t, tt.want.wantTime[key], val)
			}
		})
	}
}
