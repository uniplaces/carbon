package carbon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/carbon"
)

func TestFreeze(t *testing.T) {
	sleepDuration := 2 * time.Second

	// 2017-03-08T01:24:34+00:00
	expectedTimeStamp := int64(1488936274)
	expectedFrozenSleepTimeStamp := expectedTimeStamp + int64(sleepDuration.Seconds())

	timeToFreeze, err := carbon.CreateFromTimestampUTC(expectedTimeStamp)
	assert.Nil(t, err)

	carbon.Freeze(timeToFreeze.Time)
	defer carbon.UnFreeze()

	realNowBefore := time.Now()
	frozenNowBefore := carbon.Now()

	// Really sleep
	time.Sleep(sleepDuration)

	realNowAfter := time.Now()
	frozenNowAfter := carbon.Now()

	assert.NotEqual(t, realNowBefore, realNowAfter)
	assert.True(t, time.Since(realNowBefore).Seconds() >= sleepDuration.Seconds())
	assert.Equal(t, frozenNowBefore, frozenNowAfter)
	assert.Equal(t, expectedTimeStamp, frozenNowBefore.Unix())
	assert.Equal(t, expectedTimeStamp, frozenNowAfter.Unix())

	carbon.Sleep(sleepDuration)

	frozenAfterSleep := carbon.Now()
	assert.NotEqual(t, frozenNowAfter, frozenAfterSleep)
	assert.Equal(t, expectedFrozenSleepTimeStamp, frozenAfterSleep.Unix())
}

func TestUnFreeze(t *testing.T) {
	sleepDuration := 1 * time.Second

	// 2017-03-08T01:24:34+00:00
	expectedTimeStamp := int64(1488936274)

	timeToFreeze, err := carbon.CreateFromTimestampUTC(expectedTimeStamp)
	assert.Nil(t, err)

	carbon.Freeze(timeToFreeze.Time)
	defer carbon.UnFreeze()

	frozenNow := carbon.Now()

	// Really sleep
	time.Sleep(sleepDuration)

	carbon.UnFreeze()

	realNow := time.Now()
	unfrozenNow := carbon.Now()

	assert.NotEqual(t, frozenNow, unfrozenNow)
	assert.Equal(t, expectedTimeStamp, frozenNow.Unix())
	assert.Equal(t, realNow.Unix(), unfrozenNow.Time.Unix())
}

func TestNotSharing(t *testing.T) {
	// 2017-03-08T01:24:34+00:00
	expectedTimeStamp1 := int64(1488936274)
	timeToFreeze1, _ := carbon.CreateFromTimestampUTC(expectedTimeStamp1)
	carbon.Freeze(timeToFreeze1.Time)

	carbon1 := carbon.Now()

	time.Sleep(2 * time.Second)

	carbon11 := carbon.Now()

	// 2017-03-09T10:12:01+00:00
	expectedTimeStamp2 := int64(1489054321)
	timeToFreeze2, _ := carbon.CreateFromTimestampUTC(expectedTimeStamp2)
	carbon.Freeze(timeToFreeze2.Time)

	carbon2 := carbon.Now()

	assert.NotEqual(t, carbon1.UnixNano(), carbon2.UnixNano())
	assert.Equal(t, carbon1.UnixNano(), carbon11.UnixNano())
	assert.Equal(t, expectedTimeStamp1, carbon1.Unix())
	assert.Equal(t, expectedTimeStamp1, carbon11.Unix())
	assert.Equal(t, expectedTimeStamp2, carbon2.Unix())
}
