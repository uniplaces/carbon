package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadResource(t *testing.T) {
	tr := NewTranslator()
	err := tr.LoadResource("pt_BR")
	assert.Nil(t, err)
}

func TestTrans(t *testing.T) {
	c, _ := Create(2009, time.November, 10, 23, 0, 0, 0, "UTC")
	c.SetLocale("pt")
	assert.Equal(t, "pt", c.GetLocale())
}

// TODO: These are DiffForHumans Test 496-1214, these needs to be moved to diff_test

func TestDiffForHumansNowAndSecond(t *testing.T) {
	gotTime, err := Now().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second ago", gotTime)
}

func TestDiffForHumansNowAndSecondWithTimeZone(t *testing.T) {
	vanNow, err := NowInLocation("America/Vancouver")
	assert.Nil(t, err)

	hereNow := vanNow.Copy()
	err = hereNow.SetTimeZone(Now().TimeZone())
	assert.Nil(t, err)

	gotTime, err := vanNow.DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second ago", gotTime)
}

func TestDiffForHumansNowAndSeconds(t *testing.T) {
	gotTime, err := Now().SubSeconds(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds ago", gotTime)
}

func TestDiffForHumansNowAndNearlyMinute(t *testing.T) {
	gotTime, err := Now().SubSeconds(59).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 seconds ago", gotTime)
}

func TestDiffForHumansNowAndMinute(t *testing.T) {
	gotTime, err := Now().SubMinute().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 minute ago", gotTime)
}

func TestDiffForHumansNowAndMinutes(t *testing.T) {
	gotTime, err := Now().SubMinutes(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 minutes ago", gotTime)
}

func TestDiffForHumansNowAndNearlyHour(t *testing.T) {
	gotTime, err := Now().SubMinutes(59).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "59 minutes ago", gotTime)
}

func TestDiffForHumansNowAndHour(t *testing.T) {
	gotTime, err := Now().SubHour().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 hour ago", gotTime)
}

func TestDiffForHumansNowAndHours(t *testing.T) {
	gotTime, err := Now().SubHours(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 hours ago", gotTime)
}

func TestDiffForHumansNowAndNearlyDayOne(t *testing.T) {
	gotTime, err := Now().SubHours(743).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks ago", gotTime)
}

func TestDiffForHumansNowAndNearlyDayTwo(t *testing.T) {
	gotTime, err := Now().SubHours(744).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndNearlyDayThree(t *testing.T) {
	gotTime, err := Now().SubHours(745).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndLessThanDay(t *testing.T) {
	gotTime, err := Now().SubHours(23).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "23 hours ago", gotTime)
}

func TestDiffForHumansNowAndDay(t *testing.T) {
	gotTime, err := Now().SubDay().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 day ago", gotTime)
}

func TestDiffForHumansNowAndDays(t *testing.T) {
	gotTime, err := Now().SubDays(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 days ago", gotTime)
}

func TestDiffForHumansNowAndNearlyWeek(t *testing.T) {
	gotTime, err := Now().SubDays(6).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "6 days ago", gotTime)
}

func TestDiffForHumansNowAndWeek(t *testing.T) {
	gotTime, err := Now().SubWeek().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 week ago", gotTime)
}

func TestDiffForHumansNowAndWeeks(t *testing.T) {
	gotTime, err := Now().SubWeeks(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 weeks ago", gotTime)
}

func TestDiffForHumansNowAndNearlyMonth(t *testing.T) {
	gotTime, err := Now().SubWeeks(3).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "3 weeks ago", gotTime)
}

func TestDiffForHumansNowAndMonth(t *testing.T) {
	gotTime, err := Now().SubWeeks(4).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "4 weeks ago", gotTime)
	gotTime, err = Now().SubMonth().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 month ago", gotTime)
}

func TestDiffForHumansNowAndMonths(t *testing.T) {
	gotTime, err := Now().SubMonths(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 months ago", gotTime)
}

func TestDiffForHumansNowAndNearlyYear(t *testing.T) {
	gotTime, err := Now().SubMonths(11).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "11 months ago", gotTime)
}

func TestDiffForHumansNowAndYear(t *testing.T) {
	gotTime, err := Now().SubYear().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 year ago", gotTime)
}

func TestDiffForHumansNowAndYears(t *testing.T) {
	gotTime, err := Now().SubYears(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 years ago", gotTime)
}

func TestDiffForHumansNowAndFutureSecond(t *testing.T) {
	gotTime, err := Now().AddSecond().DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "1 second from now", gotTime)
}

/* TODO Start fixing from here, as this test case is failing
func TestDiffForHumansNowAndFutureSeconds(t *testing.T) {
	gotTime, err := Now().AddSeconds(2).DiffForHumans(nil, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "2 seconds from now", gotTime)
}

   public function testDiffForHumansNowAndFutureSeconds()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 seconds from now', Carbon::now()->addSeconds(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 seconds from now', Carbon::now()->addSeconds(59)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 minute from now', Carbon::now()->addMinute()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureMinutes()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 minutes from now', Carbon::now()->addMinutes(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 minutes from now', Carbon::now()->addMinutes(59)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 hour from now', Carbon::now()->addHour()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureHours()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 hours from now', Carbon::now()->addHours(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('23 hours from now', Carbon::now()->addHours(23)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 day from now', Carbon::now()->addDay()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureDays()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 days from now', Carbon::now()->addDays(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('6 days from now', Carbon::now()->addDays(6)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 week from now', Carbon::now()->addWeek()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureWeeks()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 weeks from now', Carbon::now()->addWeeks(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('3 weeks from now', Carbon::now()->addWeeks(3)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('4 weeks from now', Carbon::now()->addWeeks(4)->diffForHumans());
           $scope->assertSame('1 month from now', Carbon::now()->addMonth()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureMonths()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 months from now', Carbon::now()->addMonths(2)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndNearlyFutureYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('11 months from now', Carbon::now()->addMonths(11)->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 year from now', Carbon::now()->addYear()->diffForHumans());
       });
   }

   public function testDiffForHumansNowAndFutureYears()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 years from now', Carbon::now()->addYears(2)->diffForHumans());
       });
   }

   public function testDiffForHumansOtherAndSecond()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 second before', Carbon::now()->diffForHumans(Carbon::now()->addSecond()));
       });
   }

   public function testDiffForHumansOtherAndSeconds()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 seconds before', Carbon::now()->diffForHumans(Carbon::now()->addSeconds(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 seconds before', Carbon::now()->diffForHumans(Carbon::now()->addSeconds(59)));
       });
   }

   public function testDiffForHumansOtherAndMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 minute before', Carbon::now()->diffForHumans(Carbon::now()->addMinute()));
       });
   }

   public function testDiffForHumansOtherAndMinutes()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 minutes before', Carbon::now()->diffForHumans(Carbon::now()->addMinutes(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 minutes before', Carbon::now()->diffForHumans(Carbon::now()->addMinutes(59)));
       });
   }

   public function testDiffForHumansOtherAndHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 hour before', Carbon::now()->diffForHumans(Carbon::now()->addHour()));
       });
   }

   public function testDiffForHumansOtherAndHours()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 hours before', Carbon::now()->diffForHumans(Carbon::now()->addHours(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('23 hours before', Carbon::now()->diffForHumans(Carbon::now()->addHours(23)));
       });
   }

   public function testDiffForHumansOtherAndDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 day before', Carbon::now()->diffForHumans(Carbon::now()->addDay()));
       });
   }

   public function testDiffForHumansOtherAndDays()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 days before', Carbon::now()->diffForHumans(Carbon::now()->addDays(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('6 days before', Carbon::now()->diffForHumans(Carbon::now()->addDays(6)));
       });
   }

   public function testDiffForHumansOtherAndWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 week before', Carbon::now()->diffForHumans(Carbon::now()->addWeek()));
       });
   }

   public function testDiffForHumansOtherAndWeeks()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 weeks before', Carbon::now()->diffForHumans(Carbon::now()->addWeeks(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('3 weeks before', Carbon::now()->diffForHumans(Carbon::now()->addWeeks(3)));
       });
   }

   public function testDiffForHumansOtherAndMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('4 weeks before', Carbon::now()->diffForHumans(Carbon::now()->addWeeks(4)));
           $scope->assertSame('1 month before', Carbon::now()->diffForHumans(Carbon::now()->addMonth()));
       });
   }

   public function testDiffForHumansOtherAndMonths()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 months before', Carbon::now()->diffForHumans(Carbon::now()->addMonths(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('11 months before', Carbon::now()->diffForHumans(Carbon::now()->addMonths(11)));
       });
   }

   public function testDiffForHumansOtherAndYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 year before', Carbon::now()->diffForHumans(Carbon::now()->addYear()));
       });
   }

   public function testDiffForHumansOtherAndYears()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 years before', Carbon::now()->diffForHumans(Carbon::now()->addYears(2)));
       });
   }

   public function testDiffForHumansOtherAndFutureSecond()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 second after', Carbon::now()->diffForHumans(Carbon::now()->subSecond()));
       });
   }

   public function testDiffForHumansOtherAndFutureSeconds()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 seconds after', Carbon::now()->diffForHumans(Carbon::now()->subSeconds(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 seconds after', Carbon::now()->diffForHumans(Carbon::now()->subSeconds(59)));
       });
   }

   public function testDiffForHumansOtherAndFutureMinute()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 minute after', Carbon::now()->diffForHumans(Carbon::now()->subMinute()));
       });
   }

   public function testDiffForHumansOtherAndFutureMinutes()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 minutes after', Carbon::now()->diffForHumans(Carbon::now()->subMinutes(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 minutes after', Carbon::now()->diffForHumans(Carbon::now()->subMinutes(59)));
       });
   }

   public function testDiffForHumansOtherAndFutureHour()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 hour after', Carbon::now()->diffForHumans(Carbon::now()->subHour()));
       });
   }

   public function testDiffForHumansOtherAndFutureHours()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 hours after', Carbon::now()->diffForHumans(Carbon::now()->subHours(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('23 hours after', Carbon::now()->diffForHumans(Carbon::now()->subHours(23)));
       });
   }

   public function testDiffForHumansOtherAndFutureDay()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 day after', Carbon::now()->diffForHumans(Carbon::now()->subDay()));
       });
   }

   public function testDiffForHumansOtherAndFutureDays()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 days after', Carbon::now()->diffForHumans(Carbon::now()->subDays(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('6 days after', Carbon::now()->diffForHumans(Carbon::now()->subDays(6)));
       });
   }

   public function testDiffForHumansOtherAndFutureWeek()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 week after', Carbon::now()->diffForHumans(Carbon::now()->subWeek()));
       });
   }

   public function testDiffForHumansOtherAndFutureWeeks()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 weeks after', Carbon::now()->diffForHumans(Carbon::now()->subWeeks(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('3 weeks after', Carbon::now()->diffForHumans(Carbon::now()->subWeeks(3)));
       });
   }

   public function testDiffForHumansOtherAndFutureMonth()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('4 weeks after', Carbon::now()->diffForHumans(Carbon::now()->subWeeks(4)));
           $scope->assertSame('1 month after', Carbon::now()->diffForHumans(Carbon::now()->subMonth()));
       });
   }

   public function testDiffForHumansOtherAndFutureMonths()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 months after', Carbon::now()->diffForHumans(Carbon::now()->subMonths(2)));
       });
   }

   public function testDiffForHumansOtherAndNearlyFutureYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('11 months after', Carbon::now()->diffForHumans(Carbon::now()->subMonths(11)));
       });
   }

   public function testDiffForHumansOtherAndFutureYear()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 year after', Carbon::now()->diffForHumans(Carbon::now()->subYear()));
       });
   }

   public function testDiffForHumansOtherAndFutureYears()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 years after', Carbon::now()->diffForHumans(Carbon::now()->subYears(2)));
       });
   }

   public function testDiffForHumansAbsoluteSeconds()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('59 seconds', Carbon::now()->diffForHumans(Carbon::now()->subSeconds(59), true));
           $scope->assertSame('59 seconds', Carbon::now()->diffForHumans(Carbon::now()->addSeconds(59), true));
       });
   }

   public function testDiffForHumansAbsoluteMinutes()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('30 minutes', Carbon::now()->diffForHumans(Carbon::now()->subMinutes(30), true));
           $scope->assertSame('30 minutes', Carbon::now()->diffForHumans(Carbon::now()->addMinutes(30), true));
       });
   }

   public function testDiffForHumansAbsoluteHours()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('3 hours', Carbon::now()->diffForHumans(Carbon::now()->subHours(3), true));
           $scope->assertSame('3 hours', Carbon::now()->diffForHumans(Carbon::now()->addHours(3), true));
       });
   }

   public function testDiffForHumansAbsoluteDays()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 days', Carbon::now()->diffForHumans(Carbon::now()->subDays(2), true));
           $scope->assertSame('2 days', Carbon::now()->diffForHumans(Carbon::now()->addDays(2), true));
       });
   }

   public function testDiffForHumansAbsoluteWeeks()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 weeks', Carbon::now()->diffForHumans(Carbon::now()->subWeeks(2), true));
           $scope->assertSame('2 weeks', Carbon::now()->diffForHumans(Carbon::now()->addWeeks(2), true));
       });
   }

   public function testDiffForHumansAbsoluteMonths()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('2 months', Carbon::now()->diffForHumans(Carbon::now()->subMonths(2), true));
           $scope->assertSame('2 months', Carbon::now()->diffForHumans(Carbon::now()->addMonths(2), true));
       });
   }

   public function testDiffForHumansAbsoluteYears()
   {
       $scope = $this;
       $this->wrapWithTestNow(function () use ($scope) {
           $scope->assertSame('1 year', Carbon::now()->diffForHumans(Carbon::now()->subYears(1), true));
           $scope->assertSame('1 year', Carbon::now()->diffForHumans(Carbon::now()->addYears(1), true));
       });
   }

   public function testDiffForHumansWithShorterMonthShouldStillBeAMonth()
   {
       $feb15 = Carbon::parse('2015-02-15');
       $mar15 = Carbon::parse('2015-03-15');
       $this->assertSame('1 month after', $mar15->diffForHumans($feb15));
   }
*/

// TODO: This is Localization all test, this needs to be moved to local_test

/*
 */
