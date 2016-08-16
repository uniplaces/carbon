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

func TestAddWeeksPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeeks(2)

	expected := NewCarbon(time.Date(2016, time.August, 19, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestAddWeeksZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeeks(0)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestAddWeeksNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeeks(-2)

	expected := NewCarbon(time.Date(2016, time.July, 22, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 22")
}

func TestAddWeek(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddWeek()

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestAddHoursPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddHours(10)

	expected := NewCarbon(time.Date(2016, time.August, 5, 20, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 20")
}

func TestAddHoursZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddHours(0)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestAddHoursNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddHours(-11)

	expected := NewCarbon(time.Date(2016, time.August, 4, 23, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 23")
}

func TestAddHour(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddHour()

	expected := NewCarbon(time.Date(2016, time.August, 5, 11, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 11")

}

func TestAddMinutesZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddMinutes(0)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 0")
}

func TestAddMinutesPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddMinutes(50)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 50, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 50")
}

func TestAddMinutesNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddMinutes(-50)

	expected := NewCarbon(time.Date(2016, time.August, 5, 9, 10, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 50")
}

func TestAddMinute(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddMinute()

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 1, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 1")
}

func TestAddMonthsNoOverflowZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))

	d := c.AddMonthsNoOverflow(0)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to August")
}

func TestAddMonthsNoOverflowPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.AddMonthsNoOverflow(2)

	expected := NewCarbon(time.Date(2016, time.March, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestAddMonthsNoOverflowPositiveWithOverflow(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.AddMonthsNoOverflow(1)

	expected := NewCarbon(time.Date(2012, time.February, 29, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestAddMonthsNoOverflowNegative(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.February, 29, 1, 0, 0, 0, time.UTC))

	d := c.AddMonthsNoOverflow(-2)

	expected := NewCarbon(time.Date(2011, time.December, 29, 1, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestAddMonthsNoOverflowNegativeWithOverflow(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.March, 31, 1, 0, 0, 0, time.UTC))

	d := c.AddMonthsNoOverflow(-1)

	expected := NewCarbon(time.Date(2012, time.February, 29, 1, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestAddMonthNoOverflow(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.AddMonthNoOverflow()

	expected := NewCarbon(time.Date(2012, time.February, 29, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestPreviousMonthLastDay(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.March, 31, 10, 0, 0, 0, time.UTC))

	d := c.PreviousMonthLastDay()

	expected := NewCarbon(time.Date(2016, time.February, 29, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestSubYearsZero(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubYears(0)

	expected := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2012")
}

func TestSubYearsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubYears(2)

	expected := NewCarbon(time.Date(2010, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2010")
}

func TestSubYearsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubYears(-2)

	expected := NewCarbon(time.Date(2014, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2015")
}

func TestSubYear(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubYear()

	expected := NewCarbon(time.Date(2011, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2011")
}

func TestSubQuartersZero(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubQuarters(0)

	expected := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubQuartersPositive(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubQuarters(2)

	expected := NewCarbon(time.Date(2011, time.July, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to July")
}

func TestSubQuartersNegative(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubQuarters(-2)

	expected := NewCarbon(time.Date(2012, time.July, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to July")
}

func TestSubQuarter(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubQuarter()

	expected := NewCarbon(time.Date(2011, time.October, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to October")
}

func TestSubCenturiesZero(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubCenturies(0)

	expected := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2012")
}

func TestSubCenturiesPositive(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubCenturies(2)

	expected := NewCarbon(time.Date(1712, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 1712")
}

func TestSubCenturiesNegative(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubCenturies(-2)

	expected := NewCarbon(time.Date(2112, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 2112")
}

func TestSubCentury(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubCentury()

	expected := NewCarbon(time.Date(1812, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The year should be equal to 1812")
}

func TestSubMonthsZero(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonths(0)

	expected := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubMonthsPositive(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonths(2)

	expected := NewCarbon(time.Date(1911, time.November, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to November")
}

func TestSubMonthsNegative(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonths(-2)

	expected := NewCarbon(time.Date(1912, time.March, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestSubMonth(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonth()

	expected := NewCarbon(time.Date(1911, time.December, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestSubMonthsNoOverflowZero(t *testing.T) {
	c := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonthsNoOverflow(0)

	expected := NewCarbon(time.Date(2012, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubMonthsNoOverflowPositive(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonthsNoOverflow(2)

	expected := NewCarbon(time.Date(1911, time.November, 30, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to November")
}

func TestSubMonthsNoOverflowNegative(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonthsNoOverflow(-2)

	expected := NewCarbon(time.Date(1912, time.March, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestSubMonthNoOverflow(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubMonthNoOverflow()

	expected := NewCarbon(time.Date(1911, time.December, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestSubDaysZero(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubDays(0)

	expected := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 31")
}

func TestSubDaysPositive(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubDays(10)

	expected := NewCarbon(time.Date(1912, time.January, 21, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 21")
}

func TestSubDaysNegative(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubDays(-10)

	expected := NewCarbon(time.Date(1912, time.February, 10, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 10")
}

func TestSubDay(t *testing.T) {
	c := NewCarbon(time.Date(1912, time.January, 31, 10, 0, 0, 0, time.UTC))

	d := c.SubDay()

	expected := NewCarbon(time.Date(1912, time.January, 30, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 30")
}

func TestSubWeekdayZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeekdays(0)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeekdayPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeekdays(5)

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestSubWeekdayNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeekdays(-5)

	expected := NewCarbon(time.Date(2016, time.August, 19, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestSubWeekday(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 15, 10, 0, 0, 0, time.UTC))

	d := c.SubWeekday()

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeeksZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeeks(0)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeeksPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeeks(2)

	expected := NewCarbon(time.Date(2016, time.July, 29, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestSubWeeksNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeeks(-1)

	expected := NewCarbon(time.Date(2016, time.August, 19, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestSubWeek(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubWeek()

	expected := NewCarbon(time.Date(2016, time.August, 5, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestSubHoursZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubHours(0)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestSubHoursPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubHours(10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 0")
}

func TestSubHoursNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubHours(-10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 20, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 20")
}

func TestSubHour(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubHour()

	expected := NewCarbon(time.Date(2016, time.August, 12, 9, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 9")
}

func TestSubMinutesZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubMinutes(0)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestSubMinutesPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 20, 0, 0, time.UTC))

	d := c.SubMinutes(10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 10, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 10")
}

func TestSubMinutesNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 20, 0, 0, time.UTC))

	d := c.SubMinutes(-10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 30, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 30")
}

func TestSubMinute(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 20, 0, 0, time.UTC))

	d := c.SubMinute()

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 19, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The minutes should be equal to 19")
}

func TestSubSecondsZero(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	d := c.SubSeconds(0)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 0")
}

func TestSubSecondsPositive(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	d := c.SubSeconds(10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 20, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 20")
}

func TestSubSecondsNegative(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	d := c.SubSeconds(-10)

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 40, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 40")
}

func TestSubSecond(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	d := c.SubSecond()

	expected := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 29, 0, time.UTC))
	assert.Equal(t, expected, d, "The seconds should be equal to 29")
}

func TestSetters(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetYear(2010)
	c.SetMonth(time.May)
	c.SetDay(2)
	c.SetHour(5)
	c.SetMinute(10)
	c.SetSecond(10)

	expected := NewCarbon(time.Date(2010, time.May, 2, 5, 10, 10, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2010-05-02 5h 10m 10s")
}

func TestSetDate(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetDate(2015, time.May, 30)

	expected := NewCarbon(time.Date(2015, time.May, 30, 10, 0, 30, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2015-05-02")
}

func TestSetDateTime(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetDateTime(2010, time.May, 2, 5, 10, 10)

	expected := NewCarbon(time.Date(2010, time.May, 2, 5, 10, 10, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2010-05-02 5h 10m 10s")
}

func TestSetTimeFromTimeStringHour(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 0, 0, time.UTC))

	err := c.SetTimeFromTimeString("20")

	expected := NewCarbon(time.Date(2016, time.August, 12, 20, 0, 0, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 0m 0s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringMinute(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	err := c.SetTimeFromTimeString("20:20")

	expected := NewCarbon(time.Date(2016, time.August, 12, 20, 20, 30, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 30m 0s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringSecond(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	err := c.SetTimeFromTimeString("20:20:20")

	expected := NewCarbon(time.Date(2016, time.August, 12, 20, 20, 20, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 20m 20s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringEmpty(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	err := c.SetTimeFromTimeString("")

	assert.NotNil(t, err)
}

func TestSetTimeFromTimeStringInvalid(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	err := c.SetTimeFromTimeString("10-10-10")

	assert.NotNil(t, err)
}

func TestSetWeekEndsAt(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetWeekEndsAt(time.Monday)

	assert.Equal(t, time.Monday, c.WeekEndsAt(), "The end of the week should be Monday")
}

func TestSetWeekStartsAt(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetWeekStartsAt(time.Sunday)

	assert.Equal(t, time.Sunday, c.WeekStartsAt(), "The start of the week should be Sunday")
}

func TestSetWeekendDays(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))
	wds := []time.Weekday{
		time.Sunday,
		time.Monday,
	}

	c.SetWeekendDays(wds)

	assert.Equal(t, wds, c.WeekendDays(), "The start of the week should be Sunday")
}

func TestSetTimezone(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetTimezone("Europe/Lisbon")

	assert.Equal(t, "Europe/Lisbon", c.TimeZone(), "The start of the week should be Sunday")
}

func TestSetTimezoneError(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	err := c.SetTimezone("Mars/Wonderland")

	assert.NotNil(t, err)
}

func TestSetTimestamp(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetTimestamp(1171502725)

	expected := NewCarbon(time.Date(2007, time.February, 15, 1, 25, 25, 0, time.UTC))
	assert.Equal(t, expected, c, "The date should be 07-02-15 01:25:25")
}

func TestResetStringFormat(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	c.SetStringFormat(time.Kitchen)
	c.ResetStringFormat()

	assert.Equal(t, "2016-08-12 10:00:30", c.String())
}

func TestDateString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 12, 10, 0, 30, 0, time.UTC))

	assert.Equal(t, "2016-08-12", c.DateString())
}

func TestFormattedDateString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 2, 10, 0, 30, 0, time.UTC))

	assert.Equal(t, "Aug 2, 2016", c.FormattedDateString())
}

func TestTimeString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 2, 10, 0, 30, 0, time.UTC))

	assert.Equal(t, "10:00:30", c.TimeString())
}

func TestDateTimeString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 2, 10, 0, 30, 0, time.UTC))

	assert.Equal(t, "2016-08-02 10:00:30", c.DateTimeString())
}

func TestDayDateTimeString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 8, 1, 0, 30, 0, time.UTC))

	assert.Equal(t, "Mon, Aug 8, 2016 1:00 AM", c.DayDateTimeString())
}

func TestAtomString(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 1, 1, 0, 30, 0, time.UTC))

	assert.Equal(t, "2016-08-01T01:00:30+00:00", c.AtomString())
}

func TestCookieString(t *testing.T) {
	c := NewCarbon(time.Date(2005, time.August, 15, 15, 52, 1, 0, time.UTC))

	assert.Equal(t, "Monday, 15-Aug-2005 15:52:01 UTC", c.CookieString())
}

func TestIso8601String(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 1, 1, 0, 30, 0, time.UTC))

	assert.Equal(t, "2016-08-01T01:00:30+00:00", c.Iso8601String())
}

func TestRfc822String(t *testing.T) {
	c := NewCarbon(time.Date(2016, time.August, 1, 15, 28, 21, 0, time.UTC))

	assert.Equal(t, "Mon, 01 Aug 16 15:28:21 +0000", c.Rfc822String())
}
