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
	assert.Equal(t, expected, d, "The year should be equal to 2019")
}

func TestAddYearsZero(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(0)

	expected := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The year should be equal to 2009")
}

func TestAddYearsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(-10)

	expected := time.Date(1999, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The year should be equal to 1999")
}

func TestAddYear(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYear()

	expected := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The year should be equal to 2010")
}

func TestAddQuartersPositive(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(2)

	expected := time.Date(2010, time.May, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The Month should be equal to May of 2010")
}

func TestAddQuartersZero(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(0)

	expected := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The Month should be equal to November")
}


func TestAddQuartersNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(-2)

	expected := time.Date(2009, time.May, 10, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, d, "The Month should be equal to May")
}

