package carbon

// A simple API extension for DateTime
// @property      int $year
// @property      int $yearIso
// @property      int $month
// @property      int $day
// @property      int $hour
// @property      int $minute
// @property      int $second
// @property      int $timestamp seconds since the Unix Epoch
// @property      \DateTimeZone $timezone the current timezone
// @property      \DateTimeZone $tz alias of timezone
// @property-read int $micro
// @property-read int $dayOfWeek 0 (for Sunday) through 6 (for Saturday)
// @property-read int $dayOfYear 0 through 365
// @property-read int $weekOfMonth 1 through 5
// @property-read int $weekOfYear ISO-8601 week number of year, weeks starting on Monday
// @property-read int $daysInMonth number of days in the given month
// @property-read int $age does a diffInYears() with default parameters
// @property-read int $quarter the quarter of this instance, 1 - 4
// @property-read int $offset the timezone offset in seconds from UTC
// @property-read int $offsetHours the timezone offset in hours from UTC
// @property-read bool $dst daylight savings time indicator, true if DST, false otherwise
// @property-read bool $local checks if the timezone is local, true if local, false otherwise
// @property-read bool $utc checks if the timezone is UTC, true if UTC, false otherwise
// @property-read string $timezoneName
// @property-read string $tzName

// Create a Carbon instance from a DateTime one.
func instance() {
}

// Create a carbon instance from a string.
// This is an alias for the constructor that allows better fluent syntax
// as it allows you to do Carbon::parse('Monday next week')->fn() rather
// than (new Carbon('Monday next week'))->fn().
func parse() {
}

// Get a Carbon instance for the current date and time.
func now() {
}

// Create a Carbon instance for today.
func today() {
}

// Create a Carbon instance for tomorrow.
func tomorrow() {
}

// Create a Carbon instance for yesterday.
func yesterday() {
}

// Create a Carbon instance for the greatest supported date.
func maxValue() {
}

// Create a Carbon instance for the lowest supported date.
func minValue() {
}

// Create a new Carbon instance from a specific date and time.
// If any of $year, $month or $day are set to null their now() values will
// be used.
// If $hour is null it will be set to its now() value and the default
// values for $minute and $second will be their now() values.
// If $hour is not null then the default values for $minute and $second
// will be 0.
func create() {
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
func createSafe() {
}

// Create a Carbon instance from just a date. The time portion is set to now.
func createFromDate() {
}

// Create a Carbon instance from just a time. The date portion is set to today.
func createFromTime() {
}

// Create a Carbon instance from a specific format.
// @throws \InvalidArgumentException
func createFromFormat() {
}

// Set last errors.
func setLastErrors() {
}

// {@inheritdoc}
func getLastErrors() {
}

// Create a Carbon instance from a timestamp.
func createFromTimestamp() {
}

// Create a Carbon instance from an UTC timestamp.
func createFromTimestampUTC() {
}

// Get a copy of the instance.
func copy() {
}

// Get a part of the Carbon object
// @throws \InvalidArgumentException
func __get() {
}

// Check if an attribute exists on the object
func __isset() {
}

// Set a part of the Carbon object
// @throws \InvalidArgumentException
func __set() {
}

// Set the instance's year
func year() {
}

// Set the instance's month
func month() {
}

// Set the instance's day
func day() {
}

// Set the instance's hour
func hour() {
}

// Set the instance's minute
func minute() {
}

// Set the instance's second
func second() {
}

// Sets the current date of the DateTime object to a different date.
// Calls modify as a workaround for a php bug
func setDate() {
}

// Set the date and time all together
func setDateTime() {
}

// Set the time by time string
func setTimeFromTimeString() {
}

// Set the instance's timestamp
func timestamp() {
}

// Alias for setTimezone()
func timezone() {
}

// Alias for setTimezone()
func tz() {
}

// Set the instance's timezone from a string or object
func setTimezone() {
}

// Get the first day of week
func getWeekStartsAt() {
}

// Set the first day of week
func setWeekStartsAt() {
}

// Get the last day of week
func getWeekEndsAt() {
}

// Set the last day of week
func setWeekEndsAt() {
}

// Get weekend days
func getWeekendDays() {
}

// Set weekend days
func setWeekendDays() {
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
func setTestNow() {
}

// Get the Carbon instance (real or mock) to be returned when a "now"
// instance is created.
func getTestNow() {
}

// Determine if there is a valid test instance set. A valid test instance
// is anything that is not null.
func hasTestNow() {
}

// Determine if there is a relative keyword in the time string, this is to
// create dates relative to now for test instances. e.g.: next tuesday
func hasRelativeKeywords() {
}

// Intialize the translator instance if necessary.
func translator() {
}

// Get the translator instance in use
func getTranslator() {
}

// Set the translator instance to use
func setTranslator() {
}

// Get the current translator locale
func getLocale() {
}

// Set the current translator locale and indicate if the source locale file exists
func setLocale() {
}

// Format the instance with the current locale.  You can set the current
func formatLocalized() {
}

// Reset the format used to the default when type juggling a Carbon instance to a string
func resetToStringFormat() {
}

// Set the default format used when type juggling a Carbon instance to a string
func setToStringFormat() {
}

// Format the instance as a string using the set format
func __toString() {
}

// Format the instance as date
func toDateString() {
}

// Format the instance as a readable date
func toFormattedDateString() {
}

// Format the instance as time
func toTimeString() {
}

// Format the instance as date and time
func toDateTimeString() {
}

// Format the instance with day, date and time
func toDayDateTimeString() {
}

// Format the instance as ATOM
func toAtomString() {
}

// Format the instance as COOKIE
func toCookieString() {
}

// Format the instance as ISO8601
func toIso8601String() {
}

// Format the instance as RFC822
func toRfc822String() {
}

// Format the instance as RFC850
func toRfc850String() {
}

// Format the instance as RFC1036
func toRfc1036String() {
}

// Format the instance as RFC1123
func toRfc1123String() {
}

// Format the instance as RFC2822
func toRfc2822String() {
}

// Format the instance as RFC3339
func toRfc3339String() {
}

// Format the instance as RSS
func toRssString() {
}

// Format the instance as W3C
func toW3cString() {
}

// Determines if the instance is equal to another
func eq() {
}

// Determines if the instance is equal to another
// @see eq()
func equalTo() {
}

// Determines if the instance is not equal to another
func ne() {
}

// Determines if the instance is not equal to another
// @see ne()
func notEqualTo() {
}

// Determines if the instance is greater (after) than another
func gt() {
}

// Determines if the instance is greater (after) than another
// @see gt()
func greaterThan() {
}

// Determines if the instance is greater (after) than or equal to another
func gte() {
}

// Determines if the instance is greater (after) than or equal to another
// @see gte()
func greaterThanOrEqualTo() {
}

// Determines if the instance is less (before) than another
func lt() {
}

// Determines if the instance is less (before) than another
// @see lt()
func lessThan() {
}

// Determines if the instance is less (before) or equal to another
func lte() {
}

// Determines if the instance is less (before) or equal to another
// @see lte()
func lessThanOrEqualTo() {
}

// Determines if the instance is between two others
func between() {
}

// Get the closest date from the instance.
func closest() {
}

// Get the farthest date from the instance.
func farthest() {
}

// Get the minimum instance between a given instance (default now) and the current instance.
func min() {
}

// Get the minimum instance between a given instance (default now) and the current instance.
// @see min()
func minimum() {
}

// Get the maximum instance between a given instance (default now) and the current instance.
func max() {
}

// Get the maximum instance between a given instance (default now) and the current instance.
// @see max()
func maximum() {
}

// Determines if the instance is a weekday
func isWeekday() {
}

// Determines if the instance is a weekend day
func isWeekend() {
}

// Determines if the instance is yesterday
func isYesterday() {
}

// Determines if the instance is today
func isToday() {
}

// Determines if the instance is tomorrow
func isTomorrow() {
}

// Determines if the instance is in the future, ie. greater (after) than now
func isFuture() {
}

// Determines if the instance is in the past, ie. less (before) than now
func isPast() {
}

// Determines if the instance is a leap year
func isLeapYear() {
}

// Determines if the instance is a long year
func isLongYear() {
}

// Compares the formatted values of the two dates.
func isSameAs() {
}

// Determines if the instance is in the current year
func isCurrentYear() {
}

// Checks if the passed in date is in the same year as the instance year.
func isSameYear() {
}

// Determines if the instance is in the current month
func isCurrentMonth() {
}

// Checks if the passed in date is in the same month as the instance month (and year if needed).
func isSameMonth() {
}

// Checks if the passed in date is the same day as the instance current day.
func isSameDay() {
}

// Checks if this day is a Sunday.
func isSunday() {
}

// Checks if this day is a Monday.
func isMonday() {
}

// Checks if this day is a Tuesday.
func isTuesday() {
}

// Checks if this day is a Wednesday.
func isWednesday() {
}

// Checks if this day is a Thursday.
func isThursday() {
}

// Checks if this day is a Friday.
func isFriday() {
}

// Checks if this day is a Saturday.
func isSaturday() {
}

// Add years to the instance. Positive $value travel forward while
// negative $value travel into the past.
func addYears() {
}

// Add a year to the instance
func addYear() {
}

// Remove a year from the instance
func subYear() {
}

// Remove years from the instance.
func subYears() {
}

// Add quarters to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addQuarters() {
}

// Add a quarter to the instance
func addQuarter() {
}

// Remove a quarter from the instance
func subQuarter() {
}

// Remove quarters from the instance
func subQuarters() {
}

// Add centuries to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addCenturies() {
}

// Add a century to the instance
func addCentury() {
}

// Remove a century from the instance
func subCentury() {
}

// Remove centuries from the instance
func subCenturies() {
}

// Add months to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addMonths() {
}

// Add a month to the instance
func addMonth() {
}

// Remove a month from the instance
func subMonth() {
}

// Remove months from the instance
func subMonths() {
}

// Add months without overflowing to the instance. Positive $value
// travels forward while negative $value travels into the past.
func addMonthsNoOverflow() {
}

// Add a month with no overflow to the instance
func addMonthNoOverflow() {
}

// Remove a month with no overflow from the instance
func subMonthNoOverflow() {
}

// Remove months with no overflow from the instance
func subMonthsNoOverflow() {
}

// Add days to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addDays() {
}

// Add a day to the instance
func addDay() {
}

// Remove a day from the instance
func subDay() {
}

// Remove days from the instance
func subDays() {
}

// Add weekdays to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addWeekdays() {
}

// Add a weekday to the instance
func addWeekday() {
}

// Remove a weekday from the instance
func subWeekday() {
}

// Remove weekdays from the instance
func subWeekdays() {
}

// Add weeks to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addWeeks() {
}

// Add a week to the instance
func addWeek() {
}

// Remove a week from the instance
func subWeek() {
}

// Remove weeks to the instance
func subWeeks() {
}

// Add hours to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addHours() {
}

// Add an hour to the instance
func addHour() {
}

// Remove an hour from the instance
func subHour() {
}

// Remove hours from the instance
func subHours() {
}

// Add minutes to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addMinutes() {
}

// Add a minute to the instance
func addMinute() {
}

// Remove a minute from the instance
func subMinute() {
}

// Remove minutes from the instance
func subMinutes() {
}

// Add seconds to the instance. Positive $value travels forward while
// negative $value travels into the past.
func addSeconds() {
}

// Add a second to the instance
func addSecond() {
}

// Remove a second from the instance
func subSecond() {
}

// Remove seconds from the instance
func subSeconds() {
}

// Get the difference in years
func diffInYears() {
}

// Get the difference in months
func diffInMonths() {
}

// Get the difference in weeks
func diffInWeeks() {
}

// Get the difference in days
func diffInDays() {
}

// Get the difference in days using a filter closure
func diffInDaysFiltered() {
}

// Get the difference in hours using a filter closure
func diffInHoursFiltered() {
}

// Get the difference by the given interval using a filter closure
func diffFiltered() {
}

// Get the difference in weekdays
func diffInWeekdays() {
}

// Get the difference in weekend days using a filter
func diffInWeekendDays() {
}

// Get the difference in hours
func diffInHours() {
}

// Get the difference in minutes
func diffInMinutes() {
}

// Get the difference in seconds
func diffInSeconds() {
}

// The number of seconds since midnight.
func secondsSinceMidnight() {
}

// The number of seconds until 23:23:59.
func secondsUntilEndOfDay() {
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
func diffForHumans() {
}

// Resets the time to 00:00:00
func startOfDay() {
}

// Resets the time to 23:59:59
func endOfDay() {
}

// Resets the date to the first day of the month and the time to 00:00:00
func startOfMonth() {
}

// Resets the date to end of the month and time to 23:59:59
func endOfMonth() {
}

// Resets the date to the first day of the quarter and the time to 00:00:00
func startOfQuarter() {
}

// Resets the date to end of the quarter and time to 23:59:59
func endOfQuarter() {
}

// Resets the date to the first day of the year and the time to 00:00:00
func startOfYear() {
}

// Resets the date to end of the year and time to 23:59:59
func endOfYear() {
}

// Resets the date to the first day of the decade and the time to 00:00:00
func startOfDecade() {
}

// Resets the date to end of the decade and time to 23:59:59
func endOfDecade() {
}

// Resets the date to the first day of the century and the time to 00:00:00
func startOfCentury() {
}

// Resets the date to end of the century and time to 23:59:59
func endOfCentury() {
}

// Resets the date to the first day of week (defined in $weekStartsAt) and the time to 00:00:00
func startOfWeek() {
}

// Resets the date to end of week (defined in $weekEndsAt) and time to 23:59:59
func endOfWeek() {
}

// Modify to the next occurrence of a given day of the week.
// If no dayOfWeek is provided, modify to the next occurrence
// of the current day of the week.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func next() {
}

// Go forward to the next weekday.
func nextWeekday() {
}

// Go backward to the previous weekday.
func previousWeekday() {
}

// Go forward to the next weekend day.
func nextWeekendDay() {
}

// Go backward to the previous weekend day.
func previousWeekendDay() {
}

// Modify to the previous occurrence of a given day of the week.
// If no dayOfWeek is provided, modify to the previous occurrence
// of the current day of the week.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func previous() {
}

// Modify to the first occurrence of a given day of the week
// in the current month. If no dayOfWeek is provided, modify to the
// first day of the current month.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func firstOfMonth() {
}

// Modify to the last occurrence of a given day of the week
// in the current month. If no dayOfWeek is provided, modify to the
// last day of the current month.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func lastOfMonth() {
}

// Modify to the given occurrence of a given day of the week
// in the current month. If the calculated occurrence is outside the scope
// of the current month, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func nthOfMonth() {
}

// Modify to the first occurrence of a given day of the week
// in the current quarter. If no dayOfWeek is provided, modify to the
// first day of the current quarter.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func firstOfQuarter() {
}

// Modify to the last occurrence of a given day of the week
// in the current quarter. If no dayOfWeek is provided, modify to the
// last day of the current quarter.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func lastOfQuarter() {
}

// Modify to the given occurrence of a given day of the week
// in the current quarter. If the calculated occurrence is outside the scope
// of the current quarter, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func nthOfQuarter() {
}

// Modify to the first occurrence of a given day of the week
// in the current year. If no dayOfWeek is provided, modify to the
// first day of the current year.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func firstOfYear() {
}

// Modify to the last occurrence of a given day of the week
// in the current year. If no dayOfWeek is provided, modify to the
// last day of the current year.  Use the supplied consts
// to indicate the desired dayOfWeek, ex. static::MONDAY.
func lastOfYear() {
}

// Modify to the given occurrence of a given day of the week
// in the current year. If the calculated occurrence is outside the scope
// of the current year, then return false and no modifications are made.
// Use the supplied consts to indicate the desired dayOfWeek, ex. static::MONDAY.
func nthOfYear() {
}

// Modify the current instance to the average of a given instance (default now) and the current instance.
func average() {
}

func isBirthday() {
}

// Consider the timezone when modifying the instance.
func modify() {
}
