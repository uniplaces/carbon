package carbon

import (
	"errors"
	"strings"
	"time"
)

// Represent the number of elements in a given period
const (
	DaysPerWeek       = 7
	MonthsPerQuarter  = 3
	YearsPerCenturies = 100
)

// The Carbon type represents a Time instance.
// Provides a simple API extention for Time.
type Carbon struct {
	time.Time
	weekStartsAt time.Weekday
	weekEndsAt   time.Weekday
	weekendDays  []time.Weekday
}

// NewCarbon returns a pointer to a new Carbon instance
func NewCarbon(t time.Time) *Carbon {
	wds := []time.Weekday{
		time.Saturday,
		time.Sunday,
	}
	return &Carbon{
		Time:         t,
		weekStartsAt: time.Monday,
		weekEndsAt:   time.Sunday,
		weekendDays:  wds,
	}
}

// WeekStartsAt get the starting day of the week
func (c *Carbon) WeekStartsAt() time.Weekday {
	return c.weekStartsAt
}

// WeekEndsAt gets the ending day of the week
func (c *Carbon) WeekEndsAt() time.Weekday {
	return c.weekEndsAt
}

// WeekendDays gets the weekend days of the week
func (c *Carbon) WeekendDays() []time.Weekday {
	return c.weekendDays
}

// Timezone gets the current timezone
func (c *Carbon) TimeZone() string {
	return c.Location().String()
}

// AddYears adds a year to the current time
// Positive value travel forward while negative value travel into the past
func (c *Carbon) AddYears(y int) *Carbon {
	return NewCarbon(c.AddDate(y, 0, 0))
}

// AddYear adds a year to the current time
func (c *Carbon) AddYear() *Carbon {
	return c.AddYears(1)
}

// AddQuarters adds quarters to the current timePositive $value travels forward while
// Positive value travel forward while negative value travel into the past
func (c *Carbon) AddQuarters(q int) *Carbon {
	return NewCarbon(c.AddDate(0, MonthsPerQuarter*q, 0))
}

// AddQuarter adds a quarter to the current time
func (c *Carbon) AddQuarter() *Carbon {
	return c.AddQuarters(1)
}

// AddCenturies adds centuries to the time
// Positive value travels forward while negative value travels into the past
func (c *Carbon) AddCenturies(cent int) *Carbon {
	return NewCarbon(c.AddDate(YearsPerCenturies*cent, 0, 0))
}

// AddCentury adds a century to the current time
func (c *Carbon) AddCentury() *Carbon {
	return c.AddCenturies(1)
}

// AddMonths adds months to the current time
// Positive value travels forward while negative value travels into the past
func (c *Carbon) AddMonths(m int) *Carbon {
	return NewCarbon(c.AddDate(0, m, 0))
}

// AddMonth adds a month to the current time
func (c *Carbon) AddMonth() *Carbon {
	return c.AddMonths(1)
}

// AddSeconds adds seconds to the current time.
// Positive value travels forward while negative value travels into the past.
func (c *Carbon) AddSeconds(s int) *Carbon {
	d := time.Duration(s) * time.Second
	return NewCarbon(c.Add(d))
}

// AddSecond adds a second to the time
func (c *Carbon) AddSecond() *Carbon {
	return c.AddSeconds(1)
}

// AddDays adds a day to the current time.
// Positive value travels forward while negative value travels into the past
func (c *Carbon) AddDays(d int) *Carbon {
	return NewCarbon(c.AddDate(0, 0, d))
}

// AddDay adds a day to the current time
func (c *Carbon) AddDay() *Carbon {
	return c.AddDays(1)
}

// AddWeekdays adds a weekday to the current time
// Positive value travels forward while negative value travels into the past
func (c *Carbon) AddWeekdays(wd int) *Carbon {
	d := 1
	if wd < 0 {
		d = -1
		wd *= -1
	}
	t := c.Time
	for wd > 0 {
		t = t.AddDate(0, 0, d)
		if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
			wd--
		}
	}

	return NewCarbon(t)
}

// AddWeekday adds a weekday to the current time
func (c *Carbon) AddWeekday() *Carbon {
	return c.AddWeekdays(1)
}

// AddWeeks adds a week to the current time
// Positive value travels forward while negative value travels into the past.
func (c *Carbon) AddWeeks(w int) *Carbon {
	return NewCarbon(c.AddDate(0, 0, DaysPerWeek*w))
}

// AddWeek adds a week to the current time
func (c *Carbon) AddWeek() *Carbon {
	return c.AddWeeks(1)
}

// AddHours adds an hour to the current time
// Positive value travels forward while negative value travels into the past
func (c *Carbon) AddHours(h int) *Carbon {
	d := time.Duration(h) * time.Hour

	return NewCarbon(c.Add(d))
}

// AddHour adds an hour to the current time
func (c *Carbon) AddHour() *Carbon {
	return c.AddHours(1)
}

// AddMonthsNoOverflow adds a month to the current time, not overflowing in case the
// destination month has less days than the current one.
// Positive value travels forward while negative value travels into the past.
func (c *Carbon) AddMonthsNoOverflow(m int) *Carbon {
	addedDate := NewCarbon(c.AddDate(0, m, 0))
	if c.Day() != addedDate.Day() {
		return addedDate.PreviousMonthLastDay()
	}

	return addedDate
}

// PreviousMonthLastDay returns the last day of the previous month
func (c *Carbon) PreviousMonthLastDay() *Carbon {
	return NewCarbon(c.AddDate(0, 0, -c.Day()))
}

// AddMonthNoOverflow adds a month with no overflow to the current time
func (c *Carbon) AddMonthNoOverflow() *Carbon {
	return c.AddMonthsNoOverflow(1)
}

// AddMinutes adds minutes to the current time
// Positive value travels forward while negative value travels into the past.
func (c *Carbon) AddMinutes(m int) *Carbon {
	d := time.Duration(m) * time.Minute

	return NewCarbon(c.Add(d))
}

// AddMinute adds a minute to the current time
func (c *Carbon) AddMinute() *Carbon {
	return c.AddMinutes(1)
}

// SubYear removes a year from the current time
func (c *Carbon) SubYear() *Carbon {
	return c.SubYears(1)
}

// SubYears removes years from current time
func (c *Carbon) SubYears(y int) *Carbon {
	return c.AddYears(-1 * y)
}

// SubQuarter removes a quarter from the current time
func (c *Carbon) SubQuarter() *Carbon {
	return c.SubQuarters(1)
}

// SubQuarters removes quarters from current time
func (c *Carbon) SubQuarters(q int) *Carbon {
	return c.AddQuarters(-q)
}

// SubCentury removes a century from the current time
func (c *Carbon) SubCentury() *Carbon {
	return c.SubCenturies(1)
}

// SubCenturies removes centuries from the current time
func (c *Carbon) SubCenturies(cent int) *Carbon {
	return c.AddCenturies(-cent)
}

// SubMonth removes a month from the current time
func (c *Carbon) SubMonth() *Carbon {
	return c.SubMonths(1)
}

// SubMonths removes months from the current time
func (c *Carbon) SubMonths(m int) *Carbon {
	return c.AddMonths(-m)
}

// SubMonthNoOverflow remove a month with no overflow from the current time
func (c *Carbon) SubMonthNoOverflow() *Carbon {
	return c.SubMonthsNoOverflow(1)
}

// SubMonthsNoOverflow removes months with no overflow from the current time
func (c *Carbon) SubMonthsNoOverflow(m int) *Carbon {
	return c.AddMonthsNoOverflow(-m)
}

// SubDay removes a day from the current instance
func (c *Carbon) SubDay() *Carbon {
	return c.SubDays(1)
}

// SubDays removes days from the current time
func (c *Carbon) SubDays(d int) *Carbon {
	return c.AddDays(-d)
}

// SubWeekday removes a weekday from the current time
func (c *Carbon) SubWeekday() *Carbon {
	return c.SubWeekdays(1)
}

// SubWeekdays removes a weekday from the current time
func (c *Carbon) SubWeekdays(wd int) *Carbon {
	return c.AddWeekdays(-wd)
}

// SubWeek removes a week from the current time
func (c *Carbon) SubWeek() *Carbon {
	return c.SubWeeks(1)
}

// SubWeeks removes weeks to the current time
func (c *Carbon) SubWeeks(w int) *Carbon {
	return c.AddWeeks(-w)
}

// SubHour removes an hour from the current time
func (c *Carbon) SubHour() *Carbon {
	return c.SubHours(1)
}

// SubHours removes hours from the current time
func (c *Carbon) SubHours(h int) *Carbon {
	return c.AddHours(-h)
}

// SubMinute removes a minute from the current time
func (c *Carbon) SubMinute() *Carbon {
	return c.SubMinutes(1)
}

// SubMinutes removes minutes from the current time
func (c *Carbon) SubMinutes(m int) *Carbon {
	return c.AddMinutes(-m)
}

// SubSecond removes a second from the current time
func (c *Carbon) SubSecond() *Carbon {
	return c.SubSeconds(1)
}

// SubSeconds removes seconds from the current time
func (c *Carbon) SubSeconds(s int) *Carbon {
	return c.AddSeconds(-s)
}

// SetYear sets the year of the current time
func (c *Carbon) SetYear(y int) {
	c.Time = time.Date(y, c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMonth sets the month of the current time
func (c *Carbon) SetMonth(m time.Month) {
	c.Time = time.Date(c.Year(), m, c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDay sets the day of the current time
func (c *Carbon) SetDay(d int) {
	c.Time = time.Date(c.Year(), c.Month(), d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetHour sets the hour of the current time
func (c *Carbon) SetHour(h int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), h, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetMinute sets the minute of the current time
func (c *Carbon) SetMinute(m int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), m, c.Second(), c.Nanosecond(), c.Location())
}

// SetSecond sets the second of the current time
func (c *Carbon) SetSecond(s int) {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), s, c.Nanosecond(), c.Location())
}

// SetDate sets only the date of the current time
func (c *Carbon) SetDate(y int, m time.Month, d int) {
	c.Time = time.Date(y, m, d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// SetDateTime sets the date and the time
func (c *Carbon) SetDateTime(y int, mon time.Month, d, h, m, s int) {
	c.Time = time.Date(y, mon, d, h, m, s, c.Nanosecond(), c.Location())
}

// SetTimeFromTimeString receives a string and sets the current time
// It accepts the following formats: "hh:mm:ss", "hh:mm" and "hh"
func (c *Carbon) SetTimeFromTimeString(timeString string) error {
	layouts := []string{"15", "04", "05"}
	h, m, s, err := c.parse(layouts, timeString)
	if err != nil {
		return err
	}
	c.SetHour(h)
	c.SetMinute(m)
	c.SetSecond(s)

	return nil
}

func (c *Carbon) parse(layouts []string, timeString string) (int, int, int, error) {
	currLayout := strings.Join(layouts, ":")
	parsed, err := time.Parse(currLayout, timeString)
	size := len(layouts)
	if err != nil && size > 0 {
		l := layouts[:size-1]
		h, m, s, err := c.parse(l, timeString)
		return h, m, s, err
	}
	h, m, s := parsed.Clock()
	switch len(layouts) {
	case 3:
		return h, m, s, nil
	case 2:
		return h, m, c.Second(), nil
	case 1:
		return h, c.Minute(), c.Second(), nil
	}

	return 0, 0, 0, errors.New("only supports hh:mm:ss, hh:mm and hh formats")
}

// SetWeekEndsAt sets the last day of week
func (c *Carbon) SetWeekEndsAt(wd time.Weekday) {
	c.weekEndsAt = wd
}

// SetWeekStartsAt sets the first day of week
func (c *Carbon) SetWeekStartsAt(wd time.Weekday) {
	c.weekStartsAt = wd
}

// SetWeekendDays sets the weekend days
func (c *Carbon) SetWeekendDays(wds []time.Weekday) {
	c.weekendDays = wds
}

// SetTimestamp sets the current time given a timestamp
func (c *Carbon) SetTimestamp(sec int64) {
	t := time.Unix(sec, 0)
	c.Time = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), c.Location())
}

// SetTimezone the location from a string
func (c *Carbon) SetTimezone(name string) error {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return err
	}
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), loc)

	return nil
}

// Get the current translator locale
func GetLocale() {
}

// Set the current translator locale and indicate if the source locale file exists
func SetLocale() {
}

//-----------------------------------------------------------
// Create a Carbon instance from a DateTime one.
func Instance() {
}

// Create a carbon instance from a string.
func Parse() {
}

// Get a Carbon instance for the current date and time.
func Now() {
}

// Create a Carbon instance for today.
func Today() {
}

// Create a Carbon instance for tomorrow.
func Tomorrow() {
}

// Create a Carbon instance for yesterday.
func Yesterday() {
}

// Create a Carbon instance for the greatest supported date.
func MaxValue() {
}

// Create a Carbon instance for the lowest supported date.
func MinValue() {
}

// Create a new Carbon instance from a specific date and time.
// If any of $year, $month or $day are set to null their now() values will
// be used.
// If $hour is null it will be set to its now() value and the default
// values for $minute and $second will be their now() values.
// If $hour is not null then the default values for $minute and $second
// will be 0.
func Create() {
}

// Create a new safe Carbon instance from a specific date and time.
// If any of $year, $month or $day are set to null their now() values will
// be used.
// If $hour is null it will be set to its now() value and the default
// values for $minute and $second will be their now() values.
// If $hour is not null then the default values for $minute and $second
// will be 0.
// If one of the set values is not valid, an \InvalidArgumentException
// will be thrown.
// @throws \Carbon\Exceptions\InvalidDateException
func CreateSafe() {
}

// Create a Carbon instance from just a date. The time portion is set to now.
func CreateFromDate() {
}

// Create a Carbon instance from just a time. The date portion is set to today.
func CreateFromTime() {
}

// Create a Carbon instance from a specific format.
// @throws \InvalidArgumentException
func CreateFromFormat() {
}

// Create a Carbon instance from a timestamp.
func CreateFromTimestamp() {
}

// Create a Carbon instance from an UTC timestamp.
func CreateFromTimestampUTC() {
}

// Get a copy of the instance.
func Copy() {
}

// Determine if there is a relative keyword in the time string, this is to
// create dates relative to now for test instances. e.g.: next tuesday
func HasRelativeKeywords() {
}

// Intialize the translator instance if necessary.
func Translator() {
}

// Get the translator instance in use
func GetTranslator() {
}

// Set the translator instance to use
func SetTranslator() {
}

// Format the instance with the current locale.  You can set the current
func FormatLocalized() {
}

// Reset the format used to the default when type juggling a Carbon instance to a string
func ResetToStringFormat() {
}

// Set the default format used when type juggling a Carbon instance to a string
func SetToStringFormat() {
}

// Format the instance as date
func ToDateString() {
}

// Format the instance as a readable date
func ToFormattedDateString() {
}

// Format the instance as time
func ToTimeString() {
}

// Format the instance as date and time
func ToDateTimeString() {
}

// Format the instance with day, date and time
func ToDayDateTimeString() {
}

// Format the instance as ATOM
func ToAtomString() {
}

// Format the instance as COOKIE
func ToCookieString() {
}

// Format the instance as ISO8601
func ToIso8601String() {
}

// Format the instance as RFC822
func ToRfc822String() {
}

// Format the instance as RFC850
func ToRfc850String() {
}

// Format the instance as RFC1036
func ToRfc1036String() {
}

// Format the instance as RFC1123
func ToRfc1123String() {
}

// Format the instance as RFC2822
func ToRfc2822String() {
}

// Format the instance as RFC3339
func ToRfc3339String() {
}

// Format the instance as RSS
func ToRssString() {
}

// Format the instance as W3C
func ToW3cString() {
}

// Determines if the instance is equal to another
func Eq() {
}

// Determines if the instance is equal to another
// @see eq()
func EqualTo() {
}

// Determines if the instance is not equal to another
func Ne() {
}

// Determines if the instance is not equal to another
// @see ne()
func NotEqualTo() {
}

// Determines if the instance is greater (after) than another
func Gt() {
}

// Determines if the instance is greater (after) than another
// @see gt()
func GreaterThan() {
}

// Determines if the instance is greater (after) than or equal to another
func Gte() {
}

// Determines if the instance is greater (after) than or equal to another
// @see gte()
func GreaterThanOrEqualTo() {
}

// Determines if the instance is less (before) than another
func Lt() {
}

// Determines if the instance is less (before) than another
// @see lt()
func LessThan() {
}

// Determines if the instance is less (before) or equal to another
func Lte() {
}

// Determines if the instance is less (before) or equal to another
// @see lte()
func LessThanOrEqualTo() {
}

// Determines if the instance is between two others
func Between() {
}

// Get the closest date from the instance.
func Closest() {
}

// Get the farthest date from the instance.
func Farthest() {
}

// Get the minimum instance between a given instance (default now) and the current instance.
func Min() {
}

// Get the minimum instance between a given instance (default now) and the current instance.
// @see min()
func Minimum() {
}

// Get the maximum instance between a given instance (default now) and the current instance.
func Max() {
}

// Get the maximum instance between a given instance (default now) and the current instance.
// @see max()
func Maximum() {
}

// Determines if the instance is a weekday
func IsWeekday() {
}

// Determines if the instance is a weekend day
func IsWeekend() {
}

// Determines if the instance is yesterday
func IsYesterday() {
}

// Determines if the instance is today
func IsToday() {
}

// Determines if the instance is tomorrow
func IsTomorrow() {
}

// Determines if the instance is in the future, ie. greater (after) than now
func IsFuture() {
}

// Determines if the instance is in the past, ie. less (before) than now
func IsPast() {
}

// Determines if the instance is a leap year
func IsLeapYear() {
}

// Determines if the instance is a long year
func IsLongYear() {
}

// Compares the formatted values of the two dates.
func IsSameAs() {
}

// Determines if the instance is in the current year
func IsCurrentYear() {
}

// Checks if the passed in date is in the same year as the instance year.
func IsSameYear() {
}

// Determines if the instance is in the current month
func IsCurrentMonth() {
}

// Checks if the passed in date is in the same month as the instance month (and year if needed).
func IsSameMonth() {
}

// Checks if the passed in date is the same day as the instance current day.
func IsSameDay() {
}

// Checks if this day is a Sunday.
func IsSunday() {
}

// Checks if this day is a Monday.
func IsMonday() {
}

// Checks if this day is a Tuesday.
func IsTuesday() {
}

// Checks if this day is a Wednesday.
func IsWednesday() {
}

// Checks if this day is a Thursday.
func IsThursday() {
}

// Checks if this day is a Friday.
func IsFriday() {
}

// Checks if this day is a Saturday.
func IsSaturday() {
}

// Get the difference in years
func DiffInYears() {
}

// Get the difference in months
func DiffInMonths() {
}

// Get the difference in weeks
func DiffInWeeks() {
}

// Get the difference in days
func DiffInDays() {
}

// Get the difference in days using a filter closure
func DiffInDaysFiltered() {
}

// Get the difference in hours using a filter closure
func DiffInHoursFiltered() {
}

// Get the difference by the given interval using a filter closure
func DiffFiltered() {
}

// Get the difference in weekdays
func DiffInWeekdays() {
}

// Get the difference in weekend days using a filter
func DiffInWeekendDays() {
}

// Get the difference in hours
func DiffInHours() {
}

// Get the difference in minutes
func DiffInMinutes() {
}

// Get the difference in seconds
func DiffInSeconds() {
}

// The number of seconds since midnight.
func SecondsSinceMidnight() {
}

// The number of seconds until 23:23:59.
func SecondsUntilEndOfDay() {
}

// Get the difference in a human readable format in the current locale.
// When comparing a value in the past to default now:
// 1 hour ago
// 5 months ago
// When comparing a value in the future to default now:
// 1 hour from now
// 5 months from now
// When comparing a value in the past to another value:
// 1 hour before
// 5 months before
// When comparing a value in the future to another value:
// 1 hour after
// 5 months after
func DiffForHumans() {
}

// Resets the time to 00:00:00
func StartOfDay() {
}

// Resets the time to 23:59:59
func EndOfDay() {
}

// Resets the date to the first day of the month and the time to 00:00:00
func StartOfMonth() {
}

// Resets the date to end of the month and time to 23:59:59
func EndOfMonth() {
}

// Resets the date to the first day of the quarter and the time to 00:00:00
func StartOfQuarter() {
}

// Resets the date to end of the quarter and time to 23:59:59
func EndOfQuarter() {
}

// Resets the date to the first day of the year and the time to 00:00:00
func StartOfYear() {
}

// Resets the date to end of the year and time to 23:59:59
func EndOfYear() {
}

// Resets the date to the first day of the decade and the time to 00:00:00
func StartOfDecade() {
}

// Resets the date to end of the decade and time to 23:59:59
func EndOfDecade() {
}

// Resets the date to the first day of the century and the time to 00:00:00
func StartOfCentury() {
}

// Resets the date to end of the century and time to 23:59:59
func EndOfCentury() {
}

// Resets the date to the first day of week (defined in $weekStartsAt) and the time to 00:00:00
func StartOfWeek() {
}

// Resets the date to end of week (defined in $weekEndsAt) and time to 23:59:59
func EndOfWeek() {
}

// Modify to the next occurrence of a given day of the week.
// If no dayOfWeek is provided, modify to the next occurrence
// of the current day of the week.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func Next() {
}

// Go forward to the next weekday.
func NextWeekday() {
}

// Go backward to the previous weekday.
func PreviousWeekday() {
}

// Go forward to the next weekend day.
func NextWeekendDay() {
}

// Go backward to the previous weekend day.
func PreviousWeekendDay() {
}

// Modify to the previous occurrence of a given day of the week.
// If no dayOfWeek is provided, modify to the previous occurrence
// of the current day of the week.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func Previous() {
}

// Modify to the first occurrence of a given day of the week
// in the current month. If no dayOfWeek is provided, modify to the
// first day of the current month.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func FirstOfMonth() {
}

// Modify to the last occurrence of a given day of the week
// in the current month. If no dayOfWeek is provided, modify to the
// last day of the current month.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func LastOfMonth() {
}

// Modify to the given occurrence of a given day of the week
// in the current month. If the calculated occurrence is outside the scope
// of the current month, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func NthOfMonth() {
}

// Modify to the first occurrence of a given day of the week
// in the current quarter. If no dayOfWeek is provided, modify to the
// first day of the current quarter.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func FirstOfQuarter() {
}

// Modify to the last occurrence of a given day of the week
// in the current quarter. If no dayOfWeek is provided, modify to the
// last day of the current quarter.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func LastOfQuarter() {
}

// Modify to the given occurrence of a given day of the week
// in the current quarter. If the calculated occurrence is outside the scope
// of the current quarter, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func NthOfQuarter() {
}

// Modify to the first occurrence of a given day of the week
// in the current year. If no dayOfWeek is provided, modify to the
// first day of the current year.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func FirstOfYear() {
}

// Modify to the last occurrence of a given day of the week
// in the current year. If no dayOfWeek is provided, modify to the
// last day of the current year.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func LastOfYear() {
}

// Modify to the given occurrence of a given day of the week
// in the current year. If the calculated occurrence is outside the scope
// of the current year, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func NthOfYear() {
}

// Modify the current instance to the average of a given instance (default now) and the current instance.
func Average() {
}

func IsBirthday() {
}

// Consider the timezone when modifying the instance.
func Modify() {
}
