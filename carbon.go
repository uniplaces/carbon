package carbon

import "time"

const (

	MonthsPerQuarter = 3
)

type Carbon struct {
	time.Time
}

func NewCarbon(t time.Time) Carbon {
	return Carbon{Time: t}
}

// AddYear adds a year to the current time
// Positive value travel forward while negative value travel into the past.
func (c Carbon) AddYears(y int) time.Time {
	return c.AddDate(y, 0, 0)
}

// AddYear adds a year to the current time
func (c Carbon) AddYear() time.Time {
	return c.AddYears(1)
}

// AddQuarters adds quarters to the current timePositive $value travels forward while
// Positive value travel forward while negative value travel into the past.
func (c Carbon) AddQuarters(q int) time.Time {
	return c.AddDate(0, MonthsPerQuarter * q, 0)
}

//-----------------------------------------------------------
// Create a Carbon instance from a DateTime one.
func Instance() {
}

// Create a carbon instance from a string.
// This is an alias for the constructor that allows better fluent syntax
// as it allows you to do Carbon::parse('Monday next week')->fn() rather
// than (new Carbon('Monday next week'))->fn().
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

// Set last errors.
func SetLastErrors() {
}

// {@inheritdoc}
func GetLastErrors() {
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

// Set the instance's year
func Year() {
}

// Set the instance's month
func Month() {
}

// Set the instance's day
func Day() {
}

// Set the instance's hour
func Hour() {
}

// Set the instance's minute
func Minute() {
}

// Set the instance's second
func Second() {
}

// Sets the current date of the DateTime object to a different date.
// Calls modify as a workaround for a php bug
func SetDate() {
}

// Set the date and time all together
func SetDateTime() {
}

// Set the time by time string
func SetTimeFromTimeString() {
}

// Set the instance's timestamp
func Timestamp() {
}

// Alias for setTimezone()
func Timezone() {
}

// Alias for setTimezone()
func Tz() {
}

// Set the instance's timezone from a string or object
func SetTimezone() {
}

// Get the first day of week
func GetWeekStartsAt() {
}

// Set the first day of week
func SetWeekStartsAt() {
}

// Get the last day of week
func GetWeekEndsAt() {
}

// Set the last day of week
func SetWeekEndsAt() {
}

// Get weekend days
func GetWeekendDays() {
}

// Set weekend days
func SetWeekendDays() {
}

// Set a Carbon instance (real or mock) to be returned when a "now"
// instance is created.  The provided instance will be returned
// specifically under the following conditions:
//   - A call to the static now() method, ex. Carbon::now()
//   - When a null (or blank string) is passed to the constructor or parse(), ex. new Carbon(null)
//   - When the string "now" is passed to the constructor or parse(), ex. new Carbon('now')
// Note the timezone parameter was left out of the examples above and
// has no affect as the mock value will be returned regardless of its value.
// To clear the test instance call this method using the default
// parameter of null.
func SetTestNow() {
}

// Get the Carbon instance (real or mock) to be returned when a "now"
// instance is created.
func GetTestNow() {
}

// Determine if there is a valid test instance set. A valid test instance
// is anything that is not null.
func HasTestNow() {
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

// Get the current translator locale
func GetLocale() {
}

// Set the current translator locale and indicate if the source locale file exists
func SetLocale() {
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

// Remove a year from the instance
func SubYear() {
}

// Remove years from the instance.
func SubYears() {
}



// Add a quarter to the instance
func AddQuarter() {
}

// Remove a quarter from the instance
func SubQuarter() {
}

// Remove quarters from the instance
func SubQuarters() {
}

// Add centuries to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddCenturies() {
}

// Add a century to the instance
func AddCentury() {
}

// Remove a century from the instance
func SubCentury() {
}

// Remove centuries from the instance
func SubCenturies() {
}

// Add months to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddMonths() {
}

// Add a month to the instance
func AddMonth() {
}

// Remove a month from the instance
func SubMonth() {
}

// Remove months from the instance
func SubMonths() {
}

// Add months without overflowing to the instance. Positive $value
// travels forward while negative $value travels into the past.
func AddMonthsNoOverflow() {
}

// Add a month with no overflow to the instance
func AddMonthNoOverflow() {
}

// Remove a month with no overflow from the instance
func SubMonthNoOverflow() {
}

// Remove months with no overflow from the instance
func SubMonthsNoOverflow() {
}

// Add days to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddDays() {
}

// Add a day to the instance
func AddDay() {
}

// Remove a day from the instance
func SubDay() {
}

// Remove days from the instance
func SubDays() {
}

// Add weekdays to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddWeekdays() {
}

// Add a weekday to the instance
func AddWeekday() {
}

// Remove a weekday from the instance
func SubWeekday() {
}

// Remove weekdays from the instance
func SubWeekdays() {
}

// Add weeks to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddWeeks() {
}

// Add a week to the instance
func AddWeek() {
}

// Remove a week from the instance
func SubWeek() {
}

// Remove weeks to the instance
func SubWeeks() {
}

// Add hours to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddHours() {
}

// Add an hour to the instance
func AddHour() {
}

// Remove an hour from the instance
func SubHour() {
}

// Remove hours from the instance
func SubHours() {
}

// Add minutes to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddMinutes() {
}

// Add a minute to the instance
func AddMinute() {
}

// Remove a minute from the instance
func SubMinute() {
}

// Remove minutes from the instance
func SubMinutes() {
}

// Add seconds to the instance. Positive $value travels forward while
// negative $value travels into the past.
func AddSeconds() {
}

// Add a second to the instance
func AddSecond() {
}

// Remove a second from the instance
func SubSecond() {
}

// Remove seconds from the instance
func SubSeconds() {
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
