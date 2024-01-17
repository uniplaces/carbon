package carbon

import (
	"errors"
)

// CarbonInterval represents an interval between two carbons.
type CarbonInterval struct {
	Start *Carbon
	End   *Carbon
}

// NewCarbonInterval returns a pointer to a new CarbonInterval instance
func NewCarbonInterval(start, end *Carbon) (*CarbonInterval, error) {
	if start == nil {
		return nil, errors.New("start date cannot be nil")
	}
	if end == nil {
		return nil, errors.New("end date cannot be nil")
	}

	if start.Gte(end) {
		return nil, errors.New("the end date must be greater than or equal to the start date")
	}

	return &CarbonInterval{
		Start: start,
		End:   end,
	}, nil
}

// DiffInHours return the difference in hours between start and end date
func (ci *CarbonInterval) DiffInHours() int64 {
	return ci.End.DiffInHours(ci.Start, true)
}
