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

	expected := NewCarbon(time.Date(2019, time.November, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2019")
}

func TestAddYearsZero(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(0)

	expected := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2009")
}

func TestAddYearsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYears(-10)

	expected := NewCarbon(time.Date(1999, time.November, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 1999")
}

func TestAddYear(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddYear()

	expected := NewCarbon(time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2010")
}

func TestAddQuartersPositive(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(2)

	expected := NewCarbon(time.Date(2010, time.May, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to May of 2010")
}

func TestAddQuartersZero(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(0)

	expected := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to November")
}

func TestAddQuartersNegative(t *testing.T) {
	c := NewCarbon(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarters(-2)

	expected := NewCarbon(time.Date(2009, time.May, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to May")
}

func TestAddQuarter(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddQuarter()

	expected := NewCarbon(time.Date(2010, time.April, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to June of 2010")
}

func TestAddCenturiesPositive(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddCenturies(3)

	expected := NewCarbon(time.Date(2310, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal 2110")
}

func TestAddCenturiesZero(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddCenturies(0)

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal 2010")
}

func TestAddCenturiesNegative(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddCenturies(-2)

	expected := NewCarbon(time.Date(1810, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal 1810")
}

func TestAddCentury(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddCentury()

	expected := NewCarbon(time.Date(2110, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal 2110")
}
