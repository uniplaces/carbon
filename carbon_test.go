package carbon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	. "github.com/uniplaces/carbon"
)

func TestAddYearsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(10)

	expected := time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d.Year(), "The year should be equal to 2019")
}

func TestAddYearsZero(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(0)

	expected := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d.Year(), "The year should be equal to 2009")
}

func TestAddYearsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(-10)

	expected := time.Date(1999, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d.Year(), "The year should be equal to 1999")
}

func TestAddYear(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYear()

	expected := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d.Year(), "The year should be equal to 2010")
}


