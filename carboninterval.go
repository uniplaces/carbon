package carbon

import (
	"errors"
)

// CarbonInterval represents an interval between two carbons.
type CarbonInterval struct {
	Start *Carbon
	End *Carbon
}

// NewCarbonInterval returns a pointer to a new CarbonInterval instance
func NewCarbonInterval(start, end *Carbon) (*CarbonInterval, error) {
	if start.Gte(end) {
		return nil, errors.New("The end date must be greater than the start date.")
	}

	return &CarbonInterval{
		Start: start,
		End: end,
	}, nil
}

// DiffInHours return the difference in hours between start and end date
func (ci *CarbonInterval) DiffInHours() int64 {
	return ci.End.DiffInHours(ci.Start, true)
}
