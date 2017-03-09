package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddYearsPositive(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddYears(10)

	expected, _ := Create(2019, time.November, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2019")
}

func TestAddYearsZero(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddYears(0)

	expected, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2009")
}

func TestAddYearsNegative(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddYears(-10)

	expected, _ := Create(1999, time.November, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 1999")
}

func TestAddYear(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddYear()

	expected, _ := Create(2010, time.November, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2010")
}

func TestAddQuartersPositive(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddQuarters(2)

	expected, _ := Create(2010, time.May, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to May of 2010")
}

func TestAddQuartersZero(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddQuarters(0)

	expected, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to November")
}

func TestAddQuartersNegative(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	d := c.AddQuarters(-2)

	expected, _ := Create(2009, time.May, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to May")
}

func TestAddQuarter(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddQuarter()

	expected, _ := Create(2010, time.April, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to June of 2010")
}

func TestAddCenturiesPositive(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddCenturies(3)

	expected, _ := Create(2310, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal 2110")
}

func TestAddCenturiesZero(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddCenturies(0)

	expected, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal 2010")
}

func TestAddCenturiesNegative(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddCenturies(-2)

	expected, _ := Create(1810, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal 1810")
}

func TestAddCentury(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddCentury()

	expected, _ := Create(2110, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal 2110")
}

func TestAddMonthsPositive(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddMonths(2)

	expected, _ := Create(2010, time.March, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to March")
}

func TestAddMonthsZero(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddMonths(0)

	expected, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal January")
}

func TestAddMonthsNegative(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddMonths(-3)

	expected, _ := Create(2009, time.October, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to October")
}

func TestAddMonth(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddMonth()

	expected, _ := Create(2010, time.February, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The Month should be equal to February")
}

func TestAddSecondsPositive(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 22, 59, 0, 0, "UTC")

	d := c.AddSeconds(70)

	expected, _ := Create(2010, time.January, 10, 23, 0, 10, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 10")
}

func TestAddSecondsZero(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddSeconds(0)

	expected, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 0")
}

func TestAddSecondsNegative(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddSeconds(-20)

	expected, _ := Create(2010, time.January, 10, 22, 59, 40, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 40")
}

func TestAddSecond(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddSecond()

	expected, _ := Create(2010, time.January, 10, 23, 0, 1, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 1")
}

func TestAddDaysPositive(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 22, 59, 0, 0, "UTC")

	d := c.AddDays(10)

	expected, _ := Create(2010, time.January, 20, 22, 59, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 20")
}

func TestAddDaysZero(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddDays(0)

	expected, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 23")
}

func TestAddDaysNegative(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddDays(-10)

	expected, _ := Create(2009, time.December, 31, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 31")
}

func TestAddDay(t *testing.T) {
	c, _ := Create(2010, time.January, 10, 23, 0, 0, 0, "UTC")

	d := c.AddDay()

	expected, _ := Create(2010, time.January, 11, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 11")
}

func TestAddWeekdaysPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeekdays(5)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestAddWeekdaysZero(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeekdays(0)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestAddWeekdaysNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeekdays(-5)

	expected, _ := Create(2016, time.July, 29, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestAddWeekday(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeekday()

	expected, _ := Create(2016, time.August, 8, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 8")
}

func TestAddWeeksPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeeks(2)

	expected, _ := Create(2016, time.August, 19, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestAddWeeksZero(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeeks(0)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestAddWeeksNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeeks(-2)

	expected, _ := Create(2016, time.July, 22, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 22")
}

func TestAddWeek(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddWeek()

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestAddHoursPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddHours(10)

	expected, _ := Create(2016, time.August, 5, 20, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 20")
}

func TestAddHoursZero(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddHours(0)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestAddHoursNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddHours(-11)

	expected, _ := Create(2016, time.August, 4, 23, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 23")
}

func TestAddHour(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddHour()

	expected, _ := Create(2016, time.August, 5, 11, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 11")

}

func TestAddMinutesZero(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddMinutes(0)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 0")
}

func TestAddMinutesPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddMinutes(50)

	expected, _ := Create(2016, time.August, 5, 10, 50, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 50")
}

func TestAddMinutesNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddMinutes(-50)

	expected, _ := Create(2016, time.August, 5, 9, 10, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 50")
}

func TestAddMinute(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddMinute()

	expected, _ := Create(2016, time.August, 5, 10, 1, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 1")
}

func TestAddMonthsNoOverflowZero(t *testing.T) {
	c, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")

	d := c.AddMonthsNoOverflow(0)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to August")
}

func TestAddMonthsNoOverflowPositive(t *testing.T) {
	c, _ := Create(2016, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.AddMonthsNoOverflow(2)

	expected, _ := Create(2016, time.March, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestAddMonthsNoOverflowPositiveWithOverflow(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.AddMonthsNoOverflow(1)

	expected, _ := Create(2012, time.February, 29, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestAddMonthsNoOverflowNegative(t *testing.T) {
	c, _ := Create(2012, time.February, 29, 1, 0, 0, 0, "UTC")

	d := c.AddMonthsNoOverflow(-2)

	expected, _ := Create(2011, time.December, 29, 1, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestAddMonthsNoOverflowNegativeWithOverflow(t *testing.T) {
	c, _ := Create(2012, time.March, 31, 1, 0, 0, 0, "UTC")

	d := c.AddMonthsNoOverflow(-1)

	expected, _ := Create(2012, time.February, 29, 1, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestAddMonthNoOverflow(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.AddMonthNoOverflow()

	expected, _ := Create(2012, time.February, 29, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to February")
}

func TestPreviousMonthLastDay(t *testing.T) {
	c, _ := Create(2016, time.March, 31, 10, 0, 0, 0, "UTC")

	d := c.PreviousMonthLastDay()

	expected, _ := Create(2016, time.February, 29, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestSubYearsZero(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubYears(0)

	expected, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2012")
}

func TestSubYearsPositive(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubYears(2)

	expected, _ := Create(2010, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2010")
}

func TestSubYearsNegative(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubYears(-2)

	expected, _ := Create(2014, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2015")
}

func TestSubYear(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubYear()

	expected, _ := Create(2011, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2011")
}

func TestSubQuartersZero(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubQuarters(0)

	expected, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubQuartersPositive(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubQuarters(2)

	expected, _ := Create(2011, time.July, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to July")
}

func TestSubQuartersNegative(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubQuarters(-2)

	expected, _ := Create(2012, time.July, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to July")
}

func TestSubQuarter(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubQuarter()

	expected, _ := Create(2011, time.October, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to October")
}

func TestSubCenturiesZero(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubCenturies(0)

	expected, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2012")
}

func TestSubCenturiesPositive(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubCenturies(2)

	expected, _ := Create(1712, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 1712")
}

func TestSubCenturiesNegative(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubCenturies(-2)

	expected, _ := Create(2112, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 2112")
}

func TestSubCentury(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubCentury()

	expected, _ := Create(1812, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The year should be equal to 1812")
}

func TestSubMonthsZero(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonths(0)

	expected, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubMonthsPositive(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonths(2)

	expected, _ := Create(1911, time.November, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to November")
}

func TestSubMonthsNegative(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonths(-2)

	expected, _ := Create(1912, time.March, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestSubMonth(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonth()

	expected, _ := Create(1911, time.December, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestSubMonthsNoOverflowZero(t *testing.T) {
	c, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonthsNoOverflow(0)

	expected, _ := Create(2012, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to January")
}

func TestSubMonthsNoOverflowPositive(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonthsNoOverflow(2)

	expected, _ := Create(1911, time.November, 30, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to November")
}

func TestSubMonthsNoOverflowNegative(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonthsNoOverflow(-2)

	expected, _ := Create(1912, time.March, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to March")
}

func TestSubMonthNoOverflow(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubMonthNoOverflow()

	expected, _ := Create(1911, time.December, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The month should be equal to December")
}

func TestSubDaysZero(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubDays(0)

	expected, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 31")
}

func TestSubDaysPositive(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubDays(10)

	expected, _ := Create(1912, time.January, 21, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 21")
}

func TestSubDaysNegative(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubDays(-10)

	expected, _ := Create(1912, time.February, 10, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 10")
}

func TestSubDay(t *testing.T) {
	c, _ := Create(1912, time.January, 31, 10, 0, 0, 0, "UTC")

	d := c.SubDay()

	expected, _ := Create(1912, time.January, 30, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 30")
}

func TestSubWeekdayZero(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeekdays(0)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeekdayPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeekdays(5)

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestSubWeekdayNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeekdays(-5)

	expected, _ := Create(2016, time.August, 19, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestSubWeekday(t *testing.T) {
	c, _ := Create(2016, time.August, 15, 10, 0, 0, 0, "UTC")

	d := c.SubWeekday()

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeeksZero(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeeks(0)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 12")
}

func TestSubWeeksPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeeks(2)

	expected, _ := Create(2016, time.July, 29, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestSubWeeksNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeeks(-1)

	expected, _ := Create(2016, time.August, 19, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 19")
}

func TestSubWeek(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubWeek()

	expected, _ := Create(2016, time.August, 5, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The day should be equal to 5")
}

func TestSubHoursZero(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubHours(0)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestSubHoursPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubHours(10)

	expected, _ := Create(2016, time.August, 12, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 0")
}

func TestSubHoursNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubHours(-10)

	expected, _ := Create(2016, time.August, 12, 20, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 20")
}

func TestSubHour(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubHour()

	expected, _ := Create(2016, time.August, 12, 9, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 9")
}

func TestSubMinutesZero(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubMinutes(0)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The hour should be equal to 10")
}

func TestSubMinutesPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 20, 0, 0, "UTC")

	d := c.SubMinutes(10)

	expected, _ := Create(2016, time.August, 12, 10, 10, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 10")
}

func TestSubMinutesNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 20, 0, 0, "UTC")

	d := c.SubMinutes(-10)

	expected, _ := Create(2016, time.August, 12, 10, 30, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 30")
}

func TestSubMinute(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 20, 0, 0, "UTC")

	d := c.SubMinute()

	expected, _ := Create(2016, time.August, 12, 10, 19, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The minutes should be equal to 19")
}

func TestSubSecondsZero(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	d := c.SubSeconds(0)

	expected, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 0")
}

func TestSubSecondsPositive(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	d := c.SubSeconds(10)

	expected, _ := Create(2016, time.August, 12, 10, 0, 20, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 20")
}

func TestSubSecondsNegative(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	d := c.SubSeconds(-10)

	expected, _ := Create(2016, time.August, 12, 10, 0, 40, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 40")
}

func TestSubSecond(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	d := c.SubSecond()

	expected, _ := Create(2016, time.August, 12, 10, 0, 29, 0, "UTC")
	assert.Equal(t, expected, d, "The seconds should be equal to 29")
}

func TestSetters(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetYear(2010)
	c.SetMonth(time.May)
	c.SetDay(2)
	c.SetHour(5)
	c.SetMinute(10)
	c.SetSecond(10)

	expected, _ := Create(2010, time.May, 2, 5, 10, 10, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2010-05-02 5h 10m 10s")
}

func TestSetDate(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetDate(2015, time.May, 30)

	expected, _ := Create(2015, time.May, 30, 10, 0, 30, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2015-05-02")
}

func TestSetDateTime(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetDateTime(2010, time.May, 2, 5, 10, 10)

	expected, _ := Create(2010, time.May, 2, 5, 10, 10, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2010-05-02 5h 10m 10s")
}

func TestSetTimeFromTimeStringHour(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 0, 0, "UTC")

	err := c.SetTimeFromTimeString("20")

	expected, _ := Create(2016, time.August, 12, 20, 0, 0, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 0m 0s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringMinute(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	err := c.SetTimeFromTimeString("20:20")

	expected, _ := Create(2016, time.August, 12, 20, 20, 30, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 30m 0s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringSecond(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	err := c.SetTimeFromTimeString("20:20:20")

	expected, _ := Create(2016, time.August, 12, 20, 20, 20, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 2010-05-02 20h 20m 20s")
	assert.Nil(t, err)
}

func TestSetTimeFromTimeStringEmpty(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	err := c.SetTimeFromTimeString("")

	assert.NotNil(t, err)
}

func TestSetTimeFromTimeStringInvalid(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	err := c.SetTimeFromTimeString("10-10-10")

	assert.NotNil(t, err)
}

func TestSetWeekEndsAt(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetWeekEndsAt(time.Monday)

	assert.Equal(t, time.Monday, c.WeekEndsAt(), "The end of the week should be Monday")
}

func TestSetWeekStartsAt(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetWeekStartsAt(time.Sunday)

	assert.Equal(t, time.Sunday, c.WeekStartsAt(), "The start of the week should be Sunday")
}

func TestSetWeekendDays(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")
	wds := []time.Weekday{
		time.Sunday,
		time.Monday,
	}

	c.SetWeekendDays(wds)

	assert.Equal(t, wds, c.WeekendDays(), "The start of the week should be Sunday")
}

func TestSetTimezone(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetTimeZone("Europe/Lisbon")

	assert.Equal(t, "Europe/Lisbon", c.TimeZone(), "The start of the week should be Sunday")
}

func TestSetTimezoneError(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	err := c.SetTimeZone("Mars/Wonderland")

	assert.NotNil(t, err)
}

func TestSetTimestamp(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetTimestamp(1171502725)

	expected, _ := Create(2007, time.February, 15, 1, 25, 25, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 07-02-15 01:25:25")
}

func TestResetStringFormat(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	c.SetStringFormat(time.Kitchen)
	c.ResetStringFormat()

	assert.Equal(t, "2016-08-12 10:00:30", c.String())
}

func TestDateString(t *testing.T) {
	c, _ := Create(2016, time.August, 12, 10, 0, 30, 0, "UTC")

	assert.Equal(t, "2016-08-12", c.DateString())
}

func TestFormattedDateString(t *testing.T) {
	c, _ := Create(2016, time.August, 2, 10, 0, 30, 0, "UTC")

	assert.Equal(t, "Aug 2, 2016", c.FormattedDateString())
}

func TestTimeString(t *testing.T) {
	c, _ := Create(2016, time.August, 2, 10, 0, 30, 0, "UTC")

	assert.Equal(t, "10:00:30", c.TimeString())
}

func TestDateTimeString(t *testing.T) {
	c, _ := Create(2016, time.August, 2, 10, 0, 30, 0, "UTC")

	assert.Equal(t, "2016-08-02 10:00:30", c.DateTimeString())
}

func TestDayDateTimeString(t *testing.T) {
	c, _ := Create(2016, time.August, 8, 1, 0, 30, 0, "UTC")

	assert.Equal(t, "Mon, Aug 8, 2016 1:00 AM", c.DayDateTimeString())
}

func TestAtomString(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 1, 0, 30, 0, "UTC")

	assert.Equal(t, "2016-08-01T01:00:30+00:00", c.AtomString())
}

func TestCookieString(t *testing.T) {
	c, _ := Create(2005, time.August, 15, 15, 52, 1, 0, "UTC")

	assert.Equal(t, "Monday, 15-Aug-2005 15:52:01 UTC", c.CookieString())
}

func TestIso8601String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 1, 0, 30, 0, "UTC")

	assert.Equal(t, "2016-08-01T01:00:30+00:00", c.ISO8601String())
}

func TestRfc822String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Mon, 01 Aug 16 15:28:21 +0000", c.RFC822String())
}

func TestRfc850String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Monday, 01-Aug-16 15:28:21 UTC", c.RFC850String())
}

func TestRfc1036String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Mon, 01 Aug 16 15:28:21 +0000", c.RFC1036String())
}

func TestRfc1123String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Mon, 01 Aug 2016 15:28:21 +0000", c.RFC1123String())
}

func TestRfc2822String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Mon, 01 Aug 2016 15:28:21 +0000", c.RFC2822String())
}

func TestRfc3339String(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "2016-08-01T15:28:21+00:00", c.RFC3339String())
}

func TestRssString(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "Mon, 01 Aug 2016 15:28:21 +0000", c.RSSString())
}

func TestW3cString(t *testing.T) {
	c, _ := Create(2016, time.August, 1, 15, 28, 21, 0, "UTC")

	assert.Equal(t, "2016-08-01T15:28:21+00:00", c.W3CString())
}

func TestIsWeekendTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 14, 15, 28, 21, 0, "UTC")

	assert.True(t, c.IsWeekend())
}

func TestIsWeekendFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 17, 15, 28, 21, 0, "UTC")

	assert.False(t, c.IsWeekend())
}

func TestIsWeekdayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 14, 15, 28, 21, 0, "UTC")

	assert.False(t, c.IsWeekday())
}

func TestIsWeekdayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 17, 15, 28, 21, 0, "UTC")

	assert.True(t, c.IsWeekday())
}

func TestIsYesterdayTrue(t *testing.T) {
	n := Now().SubDay()

	assert.True(t, n.IsYesterday())
}

func TestIsYesterdayFalse(t *testing.T) {
	n := Now()

	assert.False(t, n.IsYesterday())
}

func TestIsTodayTrue(t *testing.T) {
	n := Now()

	assert.True(t, n.IsToday())
}

func TestIsTodayFalse(t *testing.T) {
	n := Now().SubDay()

	assert.False(t, n.IsToday())
}

func TestIsTomorrowTrue(t *testing.T) {
	n := Now().AddDay()

	assert.True(t, n.IsTomorrow())
}

func TestIsTomorrowFalse(t *testing.T) {
	n := Now()

	assert.False(t, n.IsTomorrow())
}

func TestIsFutureTrue(t *testing.T) {
	n := Now().AddSecond()

	assert.True(t, n.IsFuture())
}

func TestIsFutureFalse(t *testing.T) {
	n := Now()

	assert.False(t, n.IsFuture())
}

func TestIsPastTrue(t *testing.T) {
	n := Now().SubSecond()

	assert.True(t, n.IsPast())
}

func TestIsPastFalse(t *testing.T) {
	n := Now()
	n.SetYear(n.Year() + 1)

	assert.False(t, n.IsPast())
}

func TestIsLeapYearTrue(t *testing.T) {
	n := Now()
	n.SetYear(2016)

	assert.True(t, n.IsLeapYear())
}

func TestIsLeapYearFalse(t *testing.T) {
	n := Now()
	n.SetYear(2015)

	assert.False(t, n.IsLeapYear())
}

func TestIsLongYearTrue(t *testing.T) {
	n := Now()
	n.SetYear(2015)

	assert.True(t, n.IsLongYear())
}

func TestIsLongYearFalse(t *testing.T) {
	n := Now()
	n.SetYear(2016)

	assert.False(t, n.IsLongYear())
}

func TestIsSameAsTrue(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSameAs(DefaultFormat, d))
}

func TestIsSameAsFalse(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.November, 11, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameAs(DefaultFormat, d))
}

func TestIsSameAsNil(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameAs(DefaultFormat, nil))
}

func TestIsCurrentYearTrue(t *testing.T) {
	c := Now()

	assert.True(t, c.IsCurrentYear())
}

func TestIsCurrentYearFalse(t *testing.T) {
	c := Now().SubYear()

	assert.False(t, c.IsCurrentYear())
}

func TestIsCurrentMonthTrue(t *testing.T) {
	c := Now()

	assert.True(t, c.IsCurrentMonth())
}

func TestIsCurrentMonthFalse(t *testing.T) {
	c := Now().SubMonth()

	assert.False(t, c.IsCurrentMonth())
}

func TestIsSameYearTrue(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.January, 10, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSameYear(d))
}

func TestIsSameYearFalse(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2019, time.January, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameYear(d))
}

func TestIsSameYearNil(t *testing.T) {
	c, _ := Create(1970, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameYear(nil))
}

func TestIsSameMonthNotYearTrue(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSameMonth(d, false))
}

func TestIsSameMonthNotYearFalse(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2019, time.January, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameMonth(d, false))
}

func TestIsSameMonthNotYearNil(t *testing.T) {
	c := Now().AddMonth()

	assert.False(t, c.IsSameMonth(nil, false))
}

func TestIsSameMonthSameYearTrue(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSameMonth(d, true))
}

func TestIsSameMonthSameYearFalse(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2019, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameMonth(d, true))
}

func TestIsSameMonthSameYearNil(t *testing.T) {
	c := Now().AddMonth()

	assert.False(t, c.IsSameMonth(nil, true))
}

func TestIsSameDayTrue(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSameDay(d))
}

func TestIsSameDayFalse(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	d, _ := Create(2019, time.November, 10, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSameDay(d))
}

func TestIsSameDayNil(t *testing.T) {
	c := Now().AddMonth()

	assert.False(t, c.IsSameDay(nil))
}

func TestIsSundayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 14, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSunday())
}

func TestIsSundayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 15, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSunday())
}

func TestIsMondayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 15, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsMonday())
}

func TestIsMondayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 16, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsMonday())
}

func TestIsTuesdayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 16, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsTuesday())
}

func TestIsTuesdayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 17, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsTuesday())
}

func TestIsWednesdayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 17, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsWednesday())
}

func TestIsWednesdayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 18, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsWednesday())
}

func TestIsThursdayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 18, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsThursday())
}

func TestIsThursdayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 19, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsThursday())
}

func TestIsFridayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 19, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsFriday())
}

func TestIsFridayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 20, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsFriday())
}

func TestIsSaturdayTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 20, 23, 0, 0, 0, "UTC")

	assert.True(t, c.IsSaturday())
}

func TestIsSaturdayFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.False(t, c.IsSaturday())
}

func TestEqTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.True(t, c.Eq(d))
	assert.True(t, c.EqualTo(d))
}

func TestEqFalse(t *testing.T) {
	c, _ := Create(2015, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.False(t, c.Eq(d))
	assert.False(t, c.EqualTo(d))
}

func TestNeTrue(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2015, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.True(t, c.Ne(d))
	assert.True(t, c.NotEqualTo(d))
}

func TestNeFalse(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.False(t, c.Ne(d))
	assert.False(t, c.NotEqualTo(d))
}

func TestGtTrue(t *testing.T) {
	c := Now()
	d := Now().SubMonth()

	assert.True(t, c.Gt(d))
	assert.True(t, c.GreaterThan(d))
}

func TestGtFalseEqual(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.False(t, c.Gt(d))
	assert.False(t, c.GreaterThan(d))
}

func TestGtFalseLesser(t *testing.T) {
	c := Now()
	d := Now().AddMonth()

	assert.False(t, c.Gt(d))
	assert.False(t, c.GreaterThan(d))
}

func TestGteTrue(t *testing.T) {
	c := Now()
	d := Now().SubMonth()

	assert.True(t, c.Gte(d))
	assert.True(t, c.GreaterThanOrEqualTo(d))
}

func TestGteTrueEqual(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.True(t, c.Gte(d))
	assert.True(t, c.GreaterThanOrEqualTo(d))
}

func TestGteFalseLesser(t *testing.T) {
	c := Now()
	d := Now().AddMonth()

	assert.False(t, c.Gte(d))
	assert.False(t, c.GreaterThanOrEqualTo(d))
}

func TestLtTrue(t *testing.T) {
	c := Now()
	d := Now().AddMonth()

	assert.True(t, c.Lt(d))
	assert.True(t, c.LessThan(d))
}

func TestLtFalseEqual(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.False(t, c.Lt(d))
	assert.False(t, c.LessThan(d))
}

func TestLtFalseLesser(t *testing.T) {
	c := Now()
	d := Now().SubMonth()

	assert.False(t, c.Lt(d))
	assert.False(t, c.LessThan(d))
}

func TestLteTrue(t *testing.T) {
	c := Now()
	d := Now().AddMonth()

	assert.True(t, c.Lte(d))
	assert.True(t, c.LessThanOrEqualTo(d))
}

func TestLteTrueEqual(t *testing.T) {
	c, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")
	d, _ := Create(2016, time.August, 21, 23, 0, 0, 0, "UTC")

	assert.True(t, c.Lte(d))
	assert.True(t, c.LessThanOrEqualTo(d))
}

func TestLteFalseLesser(t *testing.T) {
	c := Now()
	d := Now().SubMonth()

	assert.False(t, c.Lte(d))
	assert.False(t, c.LessThanOrEqualTo(d))
}

func TestBetweenTrue(t *testing.T) {
	c := Now()
	d := c.SubMonth()
	b := c.AddMonth()

	assert.True(t, c.Between(d, b, false))
}

func TestBetweenEqualTrue(t *testing.T) {
	c := Now()
	d := c.SubMonth()
	b := c.AddMonth()
	c = c.SubMonth()

	assert.True(t, c.Between(d, b, true))
}

func TestBetweenSwapedTrue(t *testing.T) {
	c := Now()
	a := c.SubMonth()
	b := c.AddMonth()

	assert.True(t, c.Between(b, a, false))
}

func TestBetweenFalse(t *testing.T) {
	c := Now()
	d := c.SubMonth()
	b := c.AddMonth()
	c = c.SubMonth()

	assert.False(t, c.Between(d, b, false))
}

func TestClosestFirstArgument(t *testing.T) {
	c := Now()
	a := c.AddHour()
	b := c.AddMonth()

	assert.Equal(t, a, c.Closest(a, b))
}

func TestClosestLastArgument(t *testing.T) {
	c := Now()
	a := c.AddHour()
	b := c.AddSecond()

	assert.Equal(t, b, c.Closest(a, b))
}

func TestClosestBeforeAndAfterDates(t *testing.T) {
	c := Now()
	a := c.SubHour()
	b := c.AddMinute()

	assert.Equal(t, b, c.Closest(a, b))
}

func TestFarthestFirstArgument(t *testing.T) {
	c := Now()
	a := c.AddYear()
	b := c.AddMonth()

	assert.Equal(t, a, c.Farthest(a, b))
}

func TestFarthestLastArgument(t *testing.T) {
	c := Now()
	a := c.AddSecond()
	b := c.AddMinute()

	assert.Equal(t, b, c.Farthest(a, b))
}

func TestFarthestBeforeAndAfterDates(t *testing.T) {
	c := Now()
	a := c.SubHour()
	b := c.AddMinute()

	assert.Equal(t, a, c.Farthest(a, b))
}

func TestMinLesser(t *testing.T) {
	a := Now()
	b := a.AddMinute()

	assert.Equal(t, a, a.Min(b))
	assert.Equal(t, a, a.Minimum(b))
}

func TestMinGreater(t *testing.T) {
	a := Now()
	b := a.SubMinute()

	assert.Equal(t, b, a.Min(b))
	assert.Equal(t, b, a.Minimum(b))
}

func TestMinNil(t *testing.T) {
	a := Now().SubYear()

	assert.Equal(t, a, a.Min(nil))
	assert.Equal(t, a, a.Minimum(nil))
}

func TestMaxLesser(t *testing.T) {
	a := Now()
	b := a.SubMinute()

	assert.Equal(t, a, a.Max(b))
	assert.Equal(t, a, a.Maximum(b))
}

func TestMaxGreater(t *testing.T) {
	a := Now()
	b := a.AddMinute()

	assert.Equal(t, b, a.Max(b))
	assert.Equal(t, b, a.Maximum(b))
}

func TestMaxNil(t *testing.T) {
	a := Now().AddYear()

	assert.Equal(t, a, a.Max(nil))
	assert.Equal(t, a, a.Maximum(nil))
}

func TestTodayEET(t *testing.T) {
	today, _ := Today("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", today.TimeZone())
}

func TestTodayUnknown(t *testing.T) {
	_, err := Today("Jupiter/Wondertown")

	assert.NotNil(t, err)
}

func TestTomorrowEET(t *testing.T) {
	today, _ := Today("Africa/Cairo")
	tomorrow, _ := Tomorrow("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", tomorrow.TimeZone())
	assert.Equal(t, today.AddDay().Day(), tomorrow.Day())
}

func TestTomorrowUnknown(t *testing.T) {
	_, err := Tomorrow("Jupiter/Wondertown")

	assert.NotNil(t, err)
}

func TestYesterdayEET(t *testing.T) {
	today, _ := Today("Africa/Cairo")
	yesterday, _ := Yesterday("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", yesterday.TimeZone())
	assert.Equal(t, today.SubDay().Day(), yesterday.Day())
}

func TestYesterdayUnknown(t *testing.T) {
	_, err := Yesterday("Jupiter/Wondertown")

	assert.NotNil(t, err)
}

func TestParse(t *testing.T) {
	d, _ := Parse(DefaultFormat, "2015-11-02 16:10:22", "Africa/Cairo")

	expected, _ := Create(2015, time.November, 2, 16, 10, 22, 0, "Africa/Cairo")
	assert.Equal(t, expected, d)
}

func TestParseUnknownLocation(t *testing.T) {
	_, err := Parse(DefaultFormat, "2015-11-02 16:10:22", "WonderLand")

	assert.NotNil(t, err)
}

func TestParseInvalidFormat(t *testing.T) {
	_, err := Parse(DefaultFormat, "Clearly an invalid date format", "Africa/Cairo")

	assert.NotNil(t, err)
}

func TestDiffInSecondsNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInSeconds(t1, true))
}
func TestDiffInSecondsTimeZoneSameTime(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInSeconds(t2, true))
}

func TestDiffInSecondsTimeZoneDifferentTime(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 13, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(3600), t1.DiffInSeconds(t2, true))
}

func TestDiffInSecondsAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddSecond()

	assert.Equal(t, int64(1), t1.DiffInSeconds(t2, true))
}

func TestDiffInSecondsNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddSecond()

	assert.Equal(t, int64(-1), t2.DiffInSeconds(t1, false))
}

func TestDiffInMinutesTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInMinutes(t2, true))
}

func TestDiffInMinutesTimeZone2(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 13, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(60), t1.DiffInMinutes(t2, true))
}

func TestDiffInMinutesAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddMinute()

	assert.Equal(t, int64(1), t1.DiffInMinutes(t2, true))
}

func TestDiffInMinutesNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddMinute()

	assert.Equal(t, int64(-1), t2.DiffInMinutes(t1, false))
}

func TestDiffInHoursTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInHours(t2, true))
}

func TestDiffInHoursTimeZone2(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 13, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(1), t1.DiffInHours(t2, true))
}

func TestDiffInHoursAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddHour()

	assert.Equal(t, int64(1), t1.DiffInHours(t2, true))
}

func TestDiffInHoursNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddHour()

	assert.Equal(t, int64(-1), t2.DiffInHours(t1, false))
}

func TestDiffInDaysTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInDays(t2, true))
}

func TestDiffInDaysTimeZone2(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 13, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(5), t1.DiffInDays(t2, true))
}

func TestDiffInDaysAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddDay()

	assert.Equal(t, int64(1), t1.DiffInDays(t2, true))
}

func TestDiffInDaysNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddDay()

	assert.Equal(t, int64(-1), t2.DiffInDays(t1, false))
}

func TestDiffInDaysNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInDays(nil, true))
}

func TestDiffInNightsTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInNights(t2, true))
}

func TestDiffInNightsTimeZone2(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 13, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(5), t1.DiffInNights(t2, true))
}

func TestDiffInNightsAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddDay()

	assert.Equal(t, int64(1), t1.DiffInNights(t2, true))
}

func TestDiffInNightsNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddDay()

	assert.Equal(t, int64(-1), t2.DiffInNights(t1, false))
}

func TestDiffInNightsNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInNights(nil, true))
}

func TestDiffInWeeksNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInWeeks(nil, true))
}

func TestDiffInWeeksTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(0), t1.DiffInWeeks(t2, true))
}

func TestDiffInWeeksTimeZone2(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.July, 29, 12, 0, 0, 0, "Europe/Madrid")

	assert.Equal(t, int64(2), t1.DiffInWeeks(t2, true))
}

func TestDiffInWeeksAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddWeek()

	assert.Equal(t, int64(1), t1.DiffInWeeks(t2, true))
}

func TestDiffInWeeksNoAbs(t *testing.T) {
	t1 := Now()
	t2 := t1.AddWeek()

	assert.Equal(t, int64(-1), t2.DiffInWeeks(t1, false))
}

func TestDiffInYearsNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInYears(nil, true))
}

func TestDiffInYearsTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.EqualValues(t, 0, t1.DiffInYears(t2, true))
}

func TestDiffInYearsOneYearDifferenceNoLeapYear(t *testing.T) {
	t1, _ := Create(2014, time.August, 10, 14, 0, 0, 0, "UTC")
	t2, _ := Create(2015, time.August, 10, 14, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.DiffInYears(t2, true))
}

func TestDiffInYearsOneHourLessInYear(t *testing.T) {
	t1, _ := Create(2014, time.August, 10, 15, 0, 0, 0, "UTC")
	t2, _ := Create(2015, time.August, 10, 14, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInYears(t2, true))
}

func TestDiffInYearsOneYearDifferenceForLeapYear(t *testing.T) {
	t1, _ := Create(2015, time.August, 10, 15, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 10, 15, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.DiffInYears(t2, true))
}

func TestDiffInMonthsNil(t *testing.T) {
	t1 := Now()

	assert.EqualValues(t, 0, t1.DiffInMonths(nil, true))
}

func TestDiffInMonthsTimeZone(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.January, 18, 12, 0, 0, 0, "Europe/Madrid")

	assert.EqualValues(t, 7, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsLessThanOneYear(t *testing.T) {
	t1, _ := Create(2016, time.July, 3, 1, 0, 0, 0, "UTC")
	t2, _ := Create(2015, time.August, 1, 1, 0, 0, 0, "UTC")

	assert.EqualValues(t, 11, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsAroundOneYear(t *testing.T) {
	t1, _ := Create(2016, time.July, 4, 1, 0, 0, 0, "UTC")
	t2, _ := Create(2015, time.July, 3, 1, 0, 0, 0, "UTC")

	assert.EqualValues(t, 12, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsMoreThanOneYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 4, 1, 0, 0, 0, "UTC")
	t2, _ := Create(2015, time.July, 3, 1, 0, 0, 0, "UTC")

	assert.EqualValues(t, 13, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsOneDayDifference(t *testing.T) {
	t1, _ := Create(2016, time.September, 1, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 31, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsOneMonthDifference(t *testing.T) {
	t1, _ := Create(2016, time.September, 1, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 1, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsOneHourLessInMonthDifference(t *testing.T) {
	t1, _ := Create(2016, time.August, 1, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.September, 1, 9, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInMonths(t2, true))
}

func TestDiffInMonthsSameMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 1, 12, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 10, 6, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInMonths(t2, true))
}

func TestDiffInString(t *testing.T) {
	t1, _ := Create(2016, time.August, 10, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 1, 23, 0, 0, 0, "UTC")

	assert.EqualValues(t, "203h0m0s", t1.DiffDurationInString(t2))
}

func TestSecondsSinceMidnight(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	assert.Equal(t, 10*60*60, t1.SecondsSinceMidnight())
}

func TestSecondsUntildEndOfDay(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 0, 0, 0, 0, "UTC")
	assert.Equal(t, 24*3600-1, t1.SecondsUntilEndOfDay())
}

func TestDiffInWeekdaysSameDay(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInWeekdays(t2, true))
}

func TestDiffInWeekdaysOverWeekend(t *testing.T) {
	t1, _ := Create(2016, time.August, 19, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 22, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.DiffInWeekdays(t2, true))
}

func TestDiffInWeekdaysWithoutAbs(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, -2, t1.DiffInWeekdays(t2, false))
}

func TestDiffInWeekdaysWithAbs(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 18, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 2, t1.DiffInWeekdays(t2, true))
}

func TestDiffInWeekdaysOnWeekend(t *testing.T) {
	t1, _ := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 21, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInWeekdays(t2, true))
}

func TestDiffInWeekendDays(t *testing.T) {
	t1, _ := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 0, t1.DiffInWeekendDays(t2, true))
}

func TestDiffInWeekendDaysOnWeekend(t *testing.T) {
	t1, _ := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 21, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.DiffInWeekendDays(t2, true))
}

func TestDiffInWeekendDaysOnWeekendNegative(t *testing.T) {
	t1, _ := Create(2016, time.August, 21, 10, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")

	assert.EqualValues(t, -1, t1.DiffInWeekendDays(t2, false))
}

func TestDiffInDaysFilteredNil(t *testing.T) {
	t1 := Now()
	f := func(c *Carbon) bool {
		return c.Weekday() == time.Sunday
	}

	assert.EqualValues(t, 0, t1.DiffInDaysFiltered(f, nil, true))
}

func TestDiffInDaysFiltered(t *testing.T) {
	t1, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	t2, _ := Create(2000, time.January, 31, 0, 0, 0, 0, "UTC")
	f := func(c *Carbon) bool {
		return c.Weekday() == time.Sunday
	}

	assert.EqualValues(t, 5, t1.DiffInDaysFiltered(f, t2, true))
}

func TestDiffInDaysFilteredNegative(t *testing.T) {
	t1, _ := Create(2000, time.January, 31, 0, 0, 0, 0, "UTC")
	t2, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	f := func(c *Carbon) bool {
		return c.Weekday() == time.Sunday
	}

	assert.EqualValues(t, -5, t1.DiffInDaysFiltered(f, t2, false))
}

func TestDiffInHoursFiltered(t *testing.T) {
	t1, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	t2, _ := Create(2000, time.January, 31, 23, 59, 59, 0, "UTC")
	f := func(c *Carbon) bool {
		return c.Hour() == 9
	}

	assert.EqualValues(t, 31, t1.DiffInHoursFiltered(f, t2, true))
}

func TestDiffInHoursFilteredNegative(t *testing.T) {
	t1, _ := Create(2000, time.January, 31, 23, 59, 59, 0, "UTC")
	t2, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	f := func(c *Carbon) bool {
		return c.Hour() == 9
	}

	assert.EqualValues(t, -31, t1.DiffInHoursFiltered(f, t2, false))
}

func TestStartOfDay(t *testing.T) {
	t1, _ := Create(2009, time.August, 31, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfDay()

	expected, _ := Create(2009, time.August, 31, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfDay(t *testing.T) {
	t1, _ := Create(2009, time.August, 31, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfDay()

	expected, _ := Create(2009, time.August, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfMonth(t *testing.T) {
	t1, _ := Create(2009, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfMonth()

	expected, _ := Create(2009, time.August, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfMonth(t *testing.T) {
	t1, _ := Create(2009, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfMonth()

	expected, _ := Create(2009, time.August, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfQuarter(t *testing.T) {
	t1, _ := Create(2009, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfQuarter()

	expected, _ := Create(2009, time.July, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfQuarter(t *testing.T) {
	t1, _ := Create(2009, time.January, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfQuarter()

	expected, _ := Create(2009, time.March, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfYear(t *testing.T) {
	t1, _ := Create(2009, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfYear()

	expected, _ := Create(2009, time.January, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfYear(t *testing.T) {
	t1, _ := Create(2009, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfYear()

	expected, _ := Create(2009, time.December, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfDecade(t *testing.T) {
	t1, _ := Create(2006, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfDecade()

	expected, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfDecade(t *testing.T) {
	t1, _ := Create(2006, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfDecade()

	expected, _ := Create(2009, time.December, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfCentury(t *testing.T) {
	t1, _ := Create(2006, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfCentury()

	expected, _ := Create(2000, time.January, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfCentury(t *testing.T) {
	t1, _ := Create(2006, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfCentury()

	expected, _ := Create(2099, time.December, 31, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestStartOfWeek(t *testing.T) {
	t1, _ := Create(2016, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.StartOfWeek()

	expected, _ := Create(2016, time.August, 22, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestEndOfWeek(t *testing.T) {
	t1, _ := Create(2016, time.August, 24, 13, 0, 0, 0, "UTC")
	t2 := t1.EndOfWeek()

	expected, _ := Create(2016, time.August, 28, 23, 59, 59, 999999999, "UTC")
	assert.Equal(t, expected, t2)
}

func TestNextWeekday(t *testing.T) {
	t1, _ := Create(2016, time.August, 19, 13, 0, 0, 0, "UTC")
	t2 := t1.NextWeekday()

	expected, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestPreviousWeekday(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.PreviousWeekday()

	expected, _ := Create(2016, time.August, 19, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestNextWeekendDay(t *testing.T) {
	t1, _ := Create(2016, time.August, 18, 13, 0, 0, 0, "UTC")
	t2 := t1.NextWeekendDay()

	expected, _ := Create(2016, time.August, 20, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestNextWeekendDayOnWeekend(t *testing.T) {
	t1, _ := Create(2016, time.August, 20, 13, 0, 0, 0, "UTC")
	t2 := t1.NextWeekendDay()

	expected, _ := Create(2016, time.August, 21, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestPreviousWeekendDay(t *testing.T) {
	t1, _ := Create(2016, time.August, 23, 13, 0, 0, 0, "UTC")
	t2 := t1.PreviousWeekendDay()

	expected, _ := Create(2016, time.August, 21, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstWednesdayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfMonth(time.Wednesday)

	expected, _ := Create(2016, time.August, 3, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstMondayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfMonth(time.Monday)

	expected, _ := Create(2016, time.August, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastWednesdayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfMonth(time.Wednesday)

	expected, _ := Create(2016, time.August, 31, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastSundayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfMonth(time.Sunday)

	expected, _ := Create(2016, time.August, 28, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestSecondModayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfMonth(2, time.Monday)

	expected, _ := Create(2016, time.August, 8, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestSixthModayOfMonth(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfMonth(6, time.Monday)

	expected, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstModayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfQuarter(time.Monday)

	expected, _ := Create(2016, time.July, 4, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstDayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfQuarter(time.Friday)

	expected, _ := Create(2016, time.July, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastDayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfQuarter(time.Friday)

	expected, _ := Create(2016, time.September, 30, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastMondayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfQuarter(time.Monday)

	expected, _ := Create(2016, time.September, 26, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestThirdFridayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfQuarter(3, time.Friday)

	expected, _ := Create(2016, time.July, 15, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestSixthModayOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfQuarter(6, time.Monday)

	expected, _ := Create(2016, time.August, 8, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestDayOutOfQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfQuarter(20, time.Monday)

	expected, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstDayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfYear(time.Friday)

	expected, _ := Create(2016, time.January, 1, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstThursdayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.FirstOfYear(time.Thursday)

	expected, _ := Create(2016, time.January, 7, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastFridayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfYear(time.Friday)

	expected, _ := Create(2016, time.December, 30, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastMondayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfYear(time.Monday)

	expected, _ := Create(2016, time.December, 26, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestLastDayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.LastOfYear(time.Saturday)

	expected, _ := Create(2016, time.December, 31, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestThirdModayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfYear(3, time.Monday)

	expected, _ := Create(2016, time.January, 18, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestTenthFridayOfYear(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfYear(10, time.Friday)

	expected, _ := Create(2016, time.March, 4, 0, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestNthOfYearOutOfBounds(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	t2 := t1.NthOfYear(90, time.Monday)

	expected, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestFirstQuarter(t *testing.T) {
	t1, _ := Create(2016, time.March, 22, 13, 0, 0, 0, "UTC")

	assert.EqualValues(t, 1, t1.Quarter())
}

func TestSecondQuarter(t *testing.T) {
	t1, _ := Create(2016, time.May, 22, 13, 0, 0, 0, "UTC")

	assert.EqualValues(t, 2, t1.Quarter())
}

func TestThirdQuarter(t *testing.T) {
	t1, _ := Create(2016, time.August, 22, 13, 0, 0, 0, "UTC")

	assert.EqualValues(t, 3, t1.Quarter())
}

func TestLastQuarter(t *testing.T) {
	t1, _ := Create(2016, time.December, 22, 13, 0, 0, 0, "UTC")

	assert.EqualValues(t, 4, t1.Quarter())
}

func TestAverageNil(t *testing.T) {
	t1 := Now()

	assert.Equal(t, t1, t1.Average(nil))
}

func TestAverageSame(t *testing.T) {
	t1, _ := Create(2016, time.December, 22, 13, 0, 0, 0, "UTC")
	t2, _ := Create(2016, time.December, 22, 13, 0, 0, 0, "UTC")
	t2 = t2.Average(t1)

	assert.Equal(t, t1, t2)
}

func TestAverageGreater(t *testing.T) {
	t1, _ := Create(2000, time.January, 1, 1, 1, 1, 0, "UTC")
	t2, _ := Create(2009, time.December, 31, 23, 59, 59, 0, "UTC")
	t2 = t2.Average(t1)

	expected, _ := Create(2004, time.December, 31, 12, 30, 30, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestAverageLower(t *testing.T) {
	t1, _ := Create(2009, time.December, 31, 23, 59, 59, 0, "UTC")
	t2, _ := Create(2000, time.January, 1, 1, 1, 1, 0, "UTC")

	t2 = t2.Average(t1)

	expected, _ := Create(2004, time.December, 31, 12, 30, 30, 0, "UTC")
	assert.Equal(t, expected, t2)
}

func TestCreateFromFormat(t *testing.T) {
	d, _ := CreateFromFormat(DefaultFormat, "2015-11-02 16:10:22", "Africa/Cairo")

	expected, _ := Create(2015, time.November, 2, 16, 10, 22, 0, "Africa/Cairo")
	assert.Equal(t, expected, d)
}

func TestCreateFromFormatInvalidFormat(t *testing.T) {
	_, err := CreateFromFormat(DefaultFormat, "", "Africa/Cairo")

	assert.NotNil(t, err)
}

func TestCreateFromTimestampUTC(t *testing.T) {
	c, _ := CreateFromTimestampUTC(1171502725)

	expected, _ := Create(2007, time.February, 15, 1, 25, 25, 0, "UTC")
	assert.Equal(t, expected, c, "The date should be 07-02-15 01:25:25")
}

func TestCreateFromTimestamp(t *testing.T) {
	c, _ := CreateFromTimestamp(1171502725, "Africa/Cairo")

	expected, _ := Create(2007, time.February, 15, 1, 25, 25, 0, "Africa/Cairo")
	assert.Equal(t, expected, c, "The date should be 07-02-15 01:25:25")
}

func TestCreateFromTime(t *testing.T) {
	c, _ := CreateFromTime(10, 10, 10, 10, "UTC")

	assert.EqualValues(t, 10, c.Hour())
	assert.EqualValues(t, 10, c.Minute())
	assert.EqualValues(t, 10, c.Second())
	assert.EqualValues(t, 10, c.Nanosecond())
}

func TestCreateFromDate(t *testing.T) {
	c, _ := CreateFromDate(2011, time.August, 10, "UTC")

	assert.EqualValues(t, 2011, c.Year())
	assert.EqualValues(t, time.August, c.Month())
	assert.EqualValues(t, 10, c.Day())
}

func TestAge(t *testing.T) {
	c, _ := CreateFromDate(2011, time.August, 10, "UTC")
	y, _ := CreateFromDate(2017, time.August, 10, "UTC")
	Freeze(y.Time)

	assert.EqualValues(t, 6, c.Age())
}

func TestWeekOfMonth(t *testing.T) {
	c, _ := CreateFromDate(2016, time.January, 1, "UTC")

	assert.EqualValues(t, 1, c.WeekOfMonth())
}

func TestWeekOfMonthLastWeek(t *testing.T) {
	c, _ := CreateFromDate(2016, time.January, 31, "UTC")

	assert.EqualValues(t, 5, c.WeekOfMonth())
}

func TestWeekOfMonthLastWeekOfYear(t *testing.T) {
	c, _ := CreateFromDate(2016, time.December, 31, "UTC")

	assert.EqualValues(t, 5, c.WeekOfMonth())
}

func TestDaysInMonth(t *testing.T) {
	c, _ := CreateFromDate(2016, time.February, 1, "UTC")

	assert.EqualValues(t, 29, c.DaysInMonth())
}

func TestDaysInYearNormal(t *testing.T) {
	c, _ := CreateFromDate(2015, time.February, 1, "UTC")

	assert.EqualValues(t, 365, c.DaysInYear())
}

func TestDaysInYearLeap(t *testing.T) {
	c, _ := CreateFromDate(2016, time.February, 1, "UTC")

	assert.EqualValues(t, 366, c.DaysInYear())
}

func TestNowInLocation(t *testing.T) {
	today, _ := NowInLocation("America/Vancouver")

	assert.NotNil(t, today)
}

func TestNowInLocationInvalid(t *testing.T) {
	_, err := NowInLocation("Wonderland/Vancouver")

	assert.NotNil(t, err)
}

func TestCreateInvalidLocation(t *testing.T) {
	_, err := Create(2009, time.November, 10, 23, 0, 0, 0, "Wonderland")

	assert.NotNil(t, err)
}

func TestCreateFromFormatInvalidLocation(t *testing.T) {
	_, err := CreateFromFormat(DefaultFormat, "2015-11-02 16:10:22", "Wonderland")

	assert.NotNil(t, err)
}

func TestCreateFromTimestampInvalidLocation(t *testing.T) {
	_, err := CreateFromTimestamp(123123123, "Wonderland")

	assert.NotNil(t, err)
}

func TestLastDayOfMonth(t *testing.T) {
	c, err := Create(2016, time.August, 20, 10, 0, 0, 0, "UTC")
	assert.Nil(t, err)
	d := c.LastDayOfMonth()
	assert.Equal(t, 31, d.Day())

	c = c.AddMonth()
	d = c.LastDayOfMonth()
	assert.Equal(t, 30, d.Day())
}

func TestFirstDayOfMonth(t *testing.T) {
	c, err := Create(2016, time.May, 20, 10, 0, 0, 0, "UTC")
	assert.Nil(t, err)
	d := c.FirstDayOfMonth()
	assert.Equal(t, 1, d.Day())

	c = c.AddMonth()
	d = c.FirstDayOfMonth()
	assert.Equal(t, 1, d.Day())
}

func TestIsLastWeekTrue(t *testing.T) {
	t.Skip("IsLastWeek is a little flacky, there is no definition of when a week starts...")

	c := Now().SubDays(3)
	assert.True(t, c.IsLastWeek())
}

func TestIsLastWeekFalse(t *testing.T) {
	c, err := Create(2016, time.May, 20, 10, 0, 0, 0, "UTC")
	assert.Nil(t, err)
	assert.False(t, c.IsLastWeek())
}

func TestIsLastMonthTrue(t *testing.T) {
	c, _ := Create(2016, time.May, 5, 10, 0, 0, 0, "UTC")
	Freeze(c.Time)

	c = Now().SubWeeks(2)
	assert.True(t, c.IsLastMonth())
}

func TestIsLastMonthFalse(t *testing.T) {
	c, err := Create(2016, time.May, 20, 10, 0, 0, 0, "UTC")
	Freeze(c.Time)

	c = Now().SubWeek()
	assert.Nil(t, err)
	assert.False(t, c.IsLastMonth())
}
