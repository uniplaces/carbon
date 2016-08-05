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

func TestAddMonthsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddMonths(2)

	expected := NewCarbon(time.Date(2010, time.March, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to March")
}

func TestAddMonthsZero(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddMonths(0)

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal January")
}

func TestAddMonthsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddMonths(-3)

	expected := NewCarbon(time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to October")
}

func TestAddMonth(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddMonth()

	expected := NewCarbon(time.Date(2010, time.February, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The Month should be equal to February")
}

func TestAddSecondsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 22, 59, 0, 0, time.UTC))

	d := c.AddSeconds(70)

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 10, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 10")
}

func TestAddSecondsZero(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddSeconds(0)

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 0")
}

func TestAddSecondsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddSeconds(-20)

	expected := NewCarbon(time.Date(2010, time.January, 10, 22, 59, 40, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 40")
}

func TestAddSecond(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddSecond()

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 1, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 1")
}

func TestAddDaysPositive(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 22, 59, 0, 0, time.UTC))

	d := c.AddDays(10)

	expected := NewCarbon(time.Date(2010, time.January, 20, 22, 59, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 20")
}

func TestAddDaysZero(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddDays(0)

	expected := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 23")
}

func TestAddDaysNegative(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddDays(-10)

	expected := NewCarbon(time.Date(2009, time.December, 31, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 31")
}

func TestAddDay(t *testing.T) {
	c := NewCarbon(time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC))

	d := c.AddDay()

	expected := NewCarbon(time.Date(2010, time.January, 11, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 11")
}

func TestAddWeekdaysPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeekdays(5)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestAddWeekdaysZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeekdays(0)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestAddWeekdaysNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeekdays(-5)

	expected := NewCarbon(time.Date(2016, time.July, 29, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestAddWeekday(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeekday()

	expected := NewCarbon(time.Date(2016, time.August, 8, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 8")
}
