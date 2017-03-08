package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Check https://github.com/uniplaces/carbon/issues/30
func TestCarbonDiffInYears(t *testing.T) {
	dob, _ := CreateFromDate(2000, time.June, 27, time.UTC.String())

	yesterday, _ := CreateFromDate(2016, time.June, 26, time.UTC.String())
	today, _ := CreateFromDate(2016, time.June, 27, time.UTC.String())
	tomorrow, _ := CreateFromDate(2016, time.June, 28, time.UTC.String())

	// Day before 16th birthday... Should be 15
	assert.Equal(t, int64(15), dob.DiffInYears(yesterday, true))
	// Day of 16th birthday
	assert.Equal(t, int64(16), dob.DiffInYears(today, true))
	// Day after 16th birthday
	assert.Equal(t, int64(16), dob.DiffInYears(tomorrow, true))
}

func TestDiffForHumansNowAndSecond(t *testing.T) {
	gotTime, err := Now().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second ago", gotTime)
}

func TestDiffForHumansNowAndSecondWithTimeZone(t *testing.T) {
	vanNow, err := NowInLocation("America/Vancouver")
	assert.Nil(t, err)

	hereNow := vanNow.Copy()
	err = hereNow.SetTimeZone(Now().TimeZone())
	assert.Nil(t, err)

	gotTime, err := vanNow.DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second ago", gotTime)
}

func TestDiffForHumansNowAndSeconds(t *testing.T) {
	gotTime, err := Now().SubSeconds(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds ago", gotTime)
}

func TestDiffForHumansNowAndNearlyMinute(t *testing.T) {
	gotTime, err := Now().SubSeconds(59).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds ago", gotTime)
}

func TestDiffForHumansNowAndMinute(t *testing.T) {
	gotTime, err := Now().SubMinute().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 minute ago", gotTime)
}

func TestDiffForHumansNowAndMinutes(t *testing.T) {
	gotTime, err := Now().SubMinutes(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 minutes ago", gotTime)
}

func TestDiffForHumansNowAndNearlyHour(t *testing.T) {
	gotTime, err := Now().SubMinutes(59).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 minutes ago", gotTime)
}

func TestDiffForHumansNowAndHour(t *testing.T) {
	gotTime, err := Now().SubHour().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 hour ago", gotTime)
}

func TestDiffForHumansNowAndHours(t *testing.T) {
	gotTime, err := Now().SubHours(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 hours ago", gotTime)
}

func _TestDiffForHumansNowAndNearlyDayOne(t *testing.T) {
	gotTime, err := Now().SubHours(743).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks ago", gotTime)
}

func TestDiffForHumansNowAndNearlyDayTwo(t *testing.T) {
	gotTime, err := Now().SubHours(744).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndNearlyDayThree(t *testing.T) {
	gotTime, err := Now().SubHours(745).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndLessThanDay(t *testing.T) {
	gotTime, err := Now().SubHours(23).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "23 hours ago", gotTime)
}

func TestDiffForHumansNowAndDay(t *testing.T) {
	gotTime, err := Now().SubDay().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 day ago", gotTime)
}

func TestDiffForHumansNowAndDays(t *testing.T) {
	gotTime, err := Now().SubDays(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days ago", gotTime)
}

func TestDiffForHumansNowAndNearlyWeek(t *testing.T) {
	gotTime, err := Now().SubDays(6).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "6 days ago", gotTime)
}

func TestDiffForHumansNowAndWeek(t *testing.T) {
	gotTime, err := Now().SubWeek().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 week ago", gotTime)
}

func TestDiffForHumansNowAndWeeks(t *testing.T) {
	gotTime, err := Now().SubWeeks(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks ago", gotTime)
}

func TestDiffForHumansNowAndNearlyMonth(t *testing.T) {
	gotTime, err := Now().SubWeeks(3).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 weeks ago", gotTime)
}

func TestDiffForHumansNowAndMonth(t *testing.T) {
	gotTime, err := Now().SubWeeks(4).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks ago", gotTime)
	gotTime, err = Now().SubMonth().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndMonths(t *testing.T) {
	gotTime, err := Now().SubMonths(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months ago", gotTime)
}

//func TestDiffForHumansNowAndNearlyYear(t *testing.T) {
//	gotTime, err := Now().SubMonths(11).DiffForHumans(nil, false, false, false)
//	assert.Nil(t, err)
//	assert.Equal(t, "11 months ago", gotTime)
//}

func TestDiffForHumansNowAndYear(t *testing.T) {
	gotTime, err := Now().SubYear().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 year ago", gotTime)
}

func TestDiffForHumansNowAndYears(t *testing.T) {
	gotTime, err := Now().SubYears(2).DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 years ago", gotTime)
}

func TestDiffForHumansNowAndFutureSecond(t *testing.T) {
	gotTime, err := Now().AddSecond().DiffForHumans(nil, false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second from now", gotTime)
}

func TestDiffForHumansNowAndFutureSeconds(t *testing.T) {
	gotTime, err := Now().AddSeconds(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureMinute(t *testing.T) {
	gotTime, err := Now().AddSeconds(59).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds from now", gotTime)
}

func TestDiffForHumansNowAndFutureMinute(t *testing.T) {
	gotTime, err := Now().AddMinute().DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 minute from now", gotTime)
}

func TestDiffForHumansNowAndFutureMinutes(t *testing.T) {
	gotTime, err := Now().AddMinutes(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 minutes from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureHour(t *testing.T) {
	gotTime, err := Now().AddMinutes(59).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 minutes from now", gotTime)
}

func TestDiffForHumansNowAndFutureHour(t *testing.T) {
	gotTime, err := Now().AddHour().DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 hour from now", gotTime)
}

func TestDiffForHumansNowAndFutureHours(t *testing.T) {
	gotTime, err := Now().AddHours(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 hours from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureDay(t *testing.T) {
	gotTime, err := Now().AddHours(23).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "23 hours from now", gotTime)
}

func TestDiffForHumansNowAndFutureDay(t *testing.T) {
	gotTime, err := Now().AddDay().DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 day from now", gotTime)
}

func TestDiffForHumansNowAndFutureDays(t *testing.T) {
	gotTime, err := Now().AddDays(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureWeek(t *testing.T) {
	gotTime, err := Now().AddDays(6).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "6 days from now", gotTime)
}

func TestDiffForHumansNowAndFutureWeek(t *testing.T) {
	gotTime, err := Now().AddWeek().DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 week from now", gotTime)
}

func TestDiffForHumansNowAndFutureWeeks(t *testing.T) {
	gotTime, err := Now().AddWeeks(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureMonth(t *testing.T) {
	gotTime, err := Now().AddWeeks(3).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 weeks from now", gotTime)
}

func TestDiffForHumansNowAndLessThanFutureMonth(t *testing.T) {
	gotTime, err := Now().AddWeeks(4).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks from now", gotTime)
}

//func TestDiffForHumansNowAndFutureMonth(t *testing.T) {
//	gotTime, err := Now().AddMonth().DiffForHumans(nil, true, false, false)
//	assert.Nil(t, err)
//	assert.Equal(t, "1 month from now", gotTime)
//}

func TestDiffForHumansNowAndFutureMonths(t *testing.T) {
	gotTime, err := Now().AddMonths(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months from now", gotTime)
}

func TestDiffForHumansNowAndNearlyFutureYear(t *testing.T) {
	gotTime, err := Now().AddMonths(11).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "11 months from now", gotTime)
}

func TestDiffForHumansNowAndFutureYear(t *testing.T) {
	gotTime, err := Now().AddYear().DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 year from now", gotTime)
}

func TestDiffForHumansNowAndFutureYears(t *testing.T) {
	gotTime, err := Now().AddYears(2).DiffForHumans(nil, true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 years from now", gotTime)
}

func TestDiffForHumansOtherAndSecond(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddSecond(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second before", gotTime)
}

func TestDiffForHumansOtherAndSeconds(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddSeconds(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds before", gotTime)
}

func TestDiffForHumansOtherAndNearlyMinute(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddSeconds(59), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds before", gotTime)
}

func TestDiffForHumansOtherAndMinute(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMinute(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 minute before", gotTime)
}

func TestDiffForHumansOtherAndMinutes(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMinutes(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 minutes before", gotTime)
}

func TestDiffForHumansOtherAndNearlyHour(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMinutes(59), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 minutes before", gotTime)
}

func TestDiffForHumansOtherAndHour(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddHour(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 hour before", gotTime)
}

func TestDiffForHumansOtherAndHours(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddHours(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 hours before", gotTime)
}

func TestDiffForHumansOtherAndNearlyDay(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddHours(23), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "23 hours before", gotTime)
}

func TestDiffForHumansOtherAndDay(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddDay(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 day before", gotTime)
}

func TestDiffForHumansOtherAndDays(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddDays(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days before", gotTime)
}

func TestDiffForHumansOtherAndNearlyWeek(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddDays(6), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "6 days before", gotTime)
}

func TestDiffForHumansOtherAndWeek(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddWeek(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 week before", gotTime)
}

func TestDiffForHumansOtherAndWeeks(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddWeeks(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks before", gotTime)
}

func TestDiffForHumansOtherAndNearlyMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddWeeks(3), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 weeks before", gotTime)
}

func TestDiffForHumansOtherAndLessThanMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddWeeks(4), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks before", gotTime)
}

func TestDiffForHumansOtherAndMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMonth(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month before", gotTime)
}

func TestDiffForHumansOtherAndMonths(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMonths(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months before", gotTime)
}

func TestDiffForHumansOtherAndNearlyYear(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddMonths(11), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "11 months before", gotTime)
}

//func TestDiffForHumansOtherAndYear(t *testing.T) {
//	gotTime, err := Now().DiffForHumans(Now().AddYear(), false, false, false)
//	assert.Nil(t, err)
//	assert.Equal(t, "1 year before", gotTime)
//}

func TestDiffForHumansOtherAndYears(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().AddYears(2), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 years before", gotTime)
}

func TestDiffForHumansOtherAndFutureSecond(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubSecond(), false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second after", gotTime)
}

func TestDiffForHumansOtherAndFutureSeconds(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubSeconds(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds after", gotTime)
}

func TestDiffForHumansOtherAndNearlyFutureMinute(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubSeconds(59), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds after", gotTime)
}

func TestDiffForHumansOtherAndFutureMinute(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMinute(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 minute after", gotTime)
}

func TestDiffForHumansOtherAndFutureMinutes(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMinutes(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 minutes after", gotTime)
}

func TestDiffForHumansOtherAndNearlyFutureHour(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMinutes(59), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 minutes after", gotTime)
}

func TestDiffForHumansOtherAndFutureHour(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubHour(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 hour after", gotTime)
}

func TestDiffForHumansOtherAndFutureHours(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubHours(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 hours after", gotTime)
}

func TestDiffForHumansOtherAndNearlyFutureDay(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubHours(23), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "23 hours after", gotTime)
}

func TestDiffForHumansOtherAndFutureDay(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubDay(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 day after", gotTime)
}

func TestDiffForHumansOtherAndFutureDays(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubDays(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days after", gotTime)
}

func TestDiffForHumansOtherAndNearlyFutureWeek(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubDays(6), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "6 days after", gotTime)
}

func TestDiffForHumansOtherAndFutureWeek(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubWeek(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 week after", gotTime)
}

func TestDiffForHumansOtherAndFutureWeeks(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubWeeks(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks after", gotTime)
}

func TestDiffForHumansOtherAndNearlyFutureMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubWeeks(3), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 weeks after", gotTime)
}

func TestDiffForHumansOtherAndLessThanFutureMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubWeeks(4), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks after", gotTime)
}

func _TestDiffForHumansOtherAndFutureMonth(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMonth(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month after", gotTime)
}

func TestDiffForHumansOtherAndFutureMonths(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMonths(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months after", gotTime)
}

//func TestDiffForHumansOtherAndNearlyFutureYear(t *testing.T) {
//	gotTime, err := Now().DiffForHumans(Now().SubMonths(11), true, false, false)
//	assert.Nil(t, err)
//	assert.Equal(t, "11 months after", gotTime)
//}

func TestDiffForHumansOtherAndFutureYear(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubYear(), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 year after", gotTime)
}

func TestDiffForHumansOtherAndFutureYears(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubYears(2), true, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 years after", gotTime)
}

func TestDiffForHumansAbsoluteSeconds(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubSeconds(59), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddSeconds(59), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds", gotTime2)
}

func TestDiffForHumansAbsoluteMinutes(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMinutes(30), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "30 minutes", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddMinutes(30), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "30 minutes", gotTime2)
}

func TestDiffForHumansAbsoluteHours(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubHours(3), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 hours", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddHours(3), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 hours", gotTime2)
}

func TestDiffForHumansAbsoluteDays(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubDays(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddDays(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days", gotTime2)
}

func TestDiffForHumansAbsoluteWeeks(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubWeeks(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddWeeks(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks", gotTime2)
}

func TestDiffForHumansAbsoluteMonths(t *testing.T) {
	gotTime, err := Now().DiffForHumans(Now().SubMonths(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months", gotTime)

	gotTime2, err := Now().DiffForHumans(Now().AddMonths(2), true, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months", gotTime2)
}
