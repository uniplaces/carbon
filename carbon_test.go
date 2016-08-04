package carbon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	. "github.com/uniplaces/carbon"
)

func TestAddYearsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(1)

	assert.Equal(t, 2010, d.Year(), "The year should be equal to 2010")
	assert.Equal(t, time.November, d.Month(), "The month should be equal to November")
}

func TestAddYearsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(-1)

	assert.Equal(t, 2008, d.Year(), "The year should be equal to 2010")
	assert.Equal(t, time.November, d.Month(), "The month should be equal to November")
}
