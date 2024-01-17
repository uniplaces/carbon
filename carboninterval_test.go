package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrorOnConstruction(t *testing.T) {
	startDate, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	endDate, _ := Create(2009, time.November, 9, 23, 0, 0, 0, "UTC")

	_, err := NewCarbonInterval(startDate, endDate)
	assert.Error(t, err, "the end date must be greater than or equal to the start date")
}

func TestDiffInHours(t *testing.T) {
	startDate, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	endDate, _ := Create(2009, time.November, 10, 24, 0, 0, 0, "UTC")

	interval, _ := NewCarbonInterval(startDate, endDate)
	assert.Equal(t, int64(1), interval.DiffInHours())
}

func TestErrorOnNilStartDate(t *testing.T) {
	endDate, _ := Create(2009, time.November, 10, 24, 0, 0, 0, "UTC")

	_, err := NewCarbonInterval(nil, endDate)
	assert.Error(t, err, "start date cannot be nil")
}

func TestErrorOnNilEndDate(t *testing.T) {
	startDate, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	_, err := NewCarbonInterval(startDate, nil)
	assert.Error(t, err, "end date cannot be nil")
}

func TestErrorOnNilDates(t *testing.T) {
	_, err := NewCarbonInterval(nil, nil)
	assert.Error(t, err, "start date cannot be nil")
}
