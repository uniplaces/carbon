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
	assert.Error(t, err, "The end date must be greater than the start date.")
}

func TestDiffInHours(t *testing.T) {
	startDate, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	endDate, _ := Create(2009, time.November, 10, 24, 0, 0, 0, "UTC")

	interval, _ := NewCarbonInterval(startDate, endDate)
	assert.Equal(t, int64(1), interval.DiffInHours())
}
