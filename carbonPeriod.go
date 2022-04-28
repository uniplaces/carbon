package carbon

import "errors"

var (
	ErrEndMustBeAfterStart = errors.New("end date must be after start date")
	ErrDayAtLeast1         = errors.New("days must be at least 1")
)

// Period returns an array of Carbon dates by accepting start date, number of
// days, and end date. Useful for generating a recurring dates.
func Period(start *Carbon, days int, end *Carbon) ([]*Carbon, error) {
	if end.Before(start.Time) {
		return nil, ErrEndMustBeAfterStart
	}
	if days <= 0 {
		return nil, ErrDayAtLeast1
	}

	want := make([]*Carbon, 0)

	want = append(want, start)

	next := start
	for {
		try := next.AddDays(days)

		if try.Gte(end) {
			return want, nil
		}

		want = append(want, try)
		next = try
	}
}
