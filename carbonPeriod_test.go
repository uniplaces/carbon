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

	want := make([]*Carbon, 0)
	want1a, _ := Create(2022, time.April, 1, 23, 0, 0, 0, "UTC")
	want1b, _ := Create(2022, time.April, 8, 23, 0, 0, 0, "UTC")
	want1c, _ := Create(2022, time.April, 15, 23, 0, 0, 0, "UTC")
	want1d, _ := Create(2022, time.April, 22, 23, 0, 0, 0, "UTC")
	want1e, _ := Create(2022, time.April, 29, 23, 0, 0, 0, "UTC")
	want = append(want, want1a)
	want = append(want, want1b)
	want = append(want, want1c)
	want = append(want, want1d)
	want = append(want, want1e)

	type args struct {
		start *Carbon
		days  int
		end   *Carbon
	}

	tests := []struct {
		name string
		args args
		want []*Carbon
	}{
		{
			name: "days",
			args: args{
				start: start1,
				days:  days,
				end:   end1,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Period(tt.args.start, tt.args.days, tt.args.end)
			assert.NoError(t, err)

			for key, val := range got {
				assert.Equal(t, tt.want[key], val)
			}
		})
	}
}
