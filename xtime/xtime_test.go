package xtime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
	"time"
)

func TestSetXXX(t *testing.T) {
	now := time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("", 8*60*60))
	zero := time.Time{}

	for _, tc := range []struct {
		giveFn1 func() time.Time
		giveFn2 func()
		want    string
	}{
		{nil, func() {}, "0001-01-01T00:00:00Z"},

		{func() time.Time { return SetYear(zero, now.Year()) }, nil, "2020-01-01T00:00:00Z"},
		{func() time.Time { return SetMonth(zero, int(now.Month())) }, nil, "0001-09-01T00:00:00Z"},
		{func() time.Time { return SetDay(zero, now.Day()) }, nil, "0001-01-30T00:00:00Z"},
		{func() time.Time { return SetHour(zero, now.Hour()) }, nil, "0001-01-01T23:00:00Z"},
		{func() time.Time { return SetMinute(zero, now.Minute()) }, nil, "0001-01-01T00:39:00Z"},
		{func() time.Time { return SetSecond(zero, now.Second()) }, nil, "0001-01-01T00:00:18Z"},
		{func() time.Time { return SetMillisecond(zero, 123) }, nil, "0001-01-01T00:00:00.123Z"},
		{func() time.Time { return SetMicrosecond(zero, 123456) }, nil, "0001-01-01T00:00:00.123456Z"},
		{func() time.Time { return SetNanosecond(zero, 123456789) }, nil, "0001-01-01T00:00:00.123456789Z"},
		{func() time.Time { return SetLocation(zero, now.Location()) }, nil, "0001-01-01T00:00:00+08:00"},

		{nil, func() { zero = SetYear(zero, now.Year()) }, "2020-01-01T00:00:00Z"},
		{nil, func() { zero = SetMonth(zero, int(now.Month())) }, "2020-09-01T00:00:00Z"},
		{nil, func() { zero = SetDay(zero, now.Day()) }, "2020-09-30T00:00:00Z"},
		{nil, func() { zero = SetHour(zero, now.Hour()) }, "2020-09-30T23:00:00Z"},
		{nil, func() { zero = SetMinute(zero, now.Minute()) }, "2020-09-30T23:39:00Z"},
		{nil, func() { zero = SetSecond(zero, now.Second()) }, "2020-09-30T23:39:18Z"},
		{nil, func() { zero = SetMillisecond(zero, 123) }, "2020-09-30T23:39:18.123Z"},
		{nil, func() { zero = SetMicrosecond(zero, 123456) }, "2020-09-30T23:39:18.123456Z"},
		{nil, func() { zero = SetNanosecond(zero, now.Nanosecond()) }, "2020-09-30T23:39:18.000000789Z"},
		{nil, func() { zero = SetLocation(zero, now.Location()) }, "2020-09-30T23:39:18.000000789+08:00"},
	} {
		if tc.giveFn1 != nil {
			newTime := tc.giveFn1()
			xtesting.Equal(t, newTime.Format(time.RFC3339Nano), tc.want)
		} else {
			tc.giveFn2()
			xtesting.Equal(t, zero.Format(time.RFC3339Nano), tc.want)
		}
	}

	xtesting.Equal(t, zero, now)
}

func TestToXXX(t *testing.T) {
	now := time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("test", 8*60*60))

	for _, tc := range []struct {
		giveFn   func() time.Time
		wantTime time.Time
	}{
		{func() time.Time { return ToDate(now) },
			time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, time.FixedZone("", 8*60*60))},
		{func() time.Time { return ToDateTime(now) },
			time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*60*60))},
		{func() time.Time { return ToDateTimeNS(now) },
			time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("", 8*60*60))},
	} {
		newTime := tc.giveFn()
		xtesting.Equal(t, newTime.Year(), tc.wantTime.Year())
		xtesting.Equal(t, newTime.Month(), tc.wantTime.Month())
		xtesting.Equal(t, newTime.Day(), tc.wantTime.Day())
		xtesting.Equal(t, newTime.Hour(), tc.wantTime.Hour())
		xtesting.Equal(t, newTime.Minute(), tc.wantTime.Minute())
		xtesting.Equal(t, newTime.Second(), tc.wantTime.Second())
		xtesting.Equal(t, newTime.Nanosecond(), tc.wantTime.Nanosecond())
		xtesting.Equal(t, newTime.Location(), tc.wantTime.Location())
	}
}

func TestLocationDurationAndGetTimeLocation(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52Z")      // UTC
	t2, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52-07:00") // ""
	t3, _ := time.Parse(CJKDateTime, "2020-09-30 23:56:52")        // UTC
	t4, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52+08:00") // Local
	t5, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52+09:00") // ""
	t6, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52-12:30") // ""

	for _, tc := range []struct {
		giveTime     time.Time
		wantDuration int
	}{
		{t1, 0},                // +00:00
		{t2, -7 * 3600},        // -07:00
		{t3, 0},                // +00:00
		{t4, 8 * 3600},         // +08:00
		{t5, 9 * 3600},         // +09:00
		{t6, -12*3600 - 30*60}, // -12:30
	} {
		duration := LocationDuration(tc.giveTime.Location())
		xtesting.Equal(t, int(duration.Seconds()), tc.wantDuration)

		newLocation := GetTimeLocation(tc.giveTime)
		xtesting.Equal(t, newLocation.String(), "")
		newDuration := LocationDuration(newLocation)
		xtesting.Equal(t, int(newDuration.Seconds()), tc.wantDuration)
	}

	loc := GetLocalLocation()
	xtesting.Equal(t, loc.String(), "")
	xtesting.Equal(t, loc, time.FixedZone("", int(LocationDuration(time.Local).Seconds())))
}

func TestParseTimezone(t *testing.T) {
	for _, tc := range []struct {
		give string
		want *time.Location
	}{
		{"", nil},
		{"0", nil},
		{"09", nil},
		{"+", nil},
		{"+0", time.FixedZone("UTC+00:00", 0)}, // +X
		{"+09", time.FixedZone("UTC+09:00", 9*3600)}, // +XX
		{"+009", nil},
		{"+09:", nil},
		{"-09", time.FixedZone("UTC-09:00", -9*3600)},          // -XX
		{"-9:0", time.FixedZone("UTC-09:00", -9*3600)},         // -X:X
		{"-9:00", time.FixedZone("UTC-09:00", -9*3600)},        // -X:XX
		{"-09:0", time.FixedZone("UTC-09:00", -9*3600)},        // -XX:X
		{"-09:30", time.FixedZone("UTC-09:30", -9*3600-30*60)}, // -XX:XX
		{"-09:300", nil},
	} {
		if tc.want == nil {
			_, err := ParseTimezone(tc.give)
			xtesting.NotNil(t, err)
		} else {
			loc, err := ParseTimezone(tc.give)
			xtesting.Equal(t, loc, tc.want)
			xtesting.Nil(t, err)
		}
	}
}

func TestTruncateTime(t *testing.T) {
	for _, loc := range []*time.Location{
		time.UTC,
		time.FixedZone("", 8*60*60),
		time.FixedZone("", -9*60*60),
		time.FixedZone("", -2*30*60),
	} {
		d1 := time.Date(2021, 12, 27, 23, 49, 57, 123456789, loc)
		d2 := time.Date(2018, 5, 1, 3, 4, 6, 999000000, loc)
		for _, tc := range []struct {
			giveTime     time.Time
			giveDuration time.Duration
			want         time.Time
		}{
			{d1, time.Nanosecond, time.Date(2021, 12, 27, 23, 49, 57, 123456789, loc)},
			{d1, time.Microsecond, time.Date(2021, 12, 27, 23, 49, 57, 123456000, loc)},
			{d1, time.Millisecond, time.Date(2021, 12, 27, 23, 49, 57, 123000000, loc)},
			{d1, time.Millisecond * 10, time.Date(2021, 12, 27, 23, 49, 57, 120000000, loc)},
			{d1, time.Second, time.Date(2021, 12, 27, 23, 49, 57, 0, loc)},
			{d1, time.Second * 2, time.Date(2021, 12, 27, 23, 49, 56, 0, loc)},
			{d1, time.Second * 5, time.Date(2021, 12, 27, 23, 49, 55, 0, loc)},
			{d1, time.Second * 10, time.Date(2021, 12, 27, 23, 49, 50, 0, loc)},
			{d1, time.Second * 20, time.Date(2021, 12, 27, 23, 49, 40, 0, loc)},
			{d1, time.Minute, time.Date(2021, 12, 27, 23, 49, 0, 0, loc)},
			{d1, time.Minute * 2, time.Date(2021, 12, 27, 23, 48, 0, 0, loc)},
			{d1, time.Minute * 5, time.Date(2021, 12, 27, 23, 45, 0, 0, loc)},
			{d1, time.Minute * 10, time.Date(2021, 12, 27, 23, 40, 0, 0, loc)},
			{d1, time.Minute * 20, time.Date(2021, 12, 27, 23, 40, 0, 0, loc)},
			//
			{d2, time.Millisecond, time.Date(2018, 5, 1, 3, 4, 6, 999000000, loc)},
			{d2, time.Millisecond * 10, time.Date(2018, 5, 1, 3, 4, 6, 990000000, loc)},
			{d2, time.Second, time.Date(2018, 5, 1, 3, 4, 6, 0, loc)},
			{d2, time.Second * 2, time.Date(2018, 5, 1, 3, 4, 6, 0, loc)},
			{d2, time.Second * 20, time.Date(2018, 5, 1, 3, 4, 0, 0, loc)},
			{d2, time.Minute, time.Date(2018, 5, 1, 3, 4, 0, 0, loc)},
			{d2, time.Minute * 2, time.Date(2018, 5, 1, 3, 4, 0, 0, loc)},
			{d2, time.Minute * 20, time.Date(2018, 5, 1, 3, 0, 0, 0, loc)},
			//
			{d1, time.Hour, time.Date(2021, 12, 27, 23, 0, 0, 0, loc)},
			{d1, time.Hour * 2, time.Date(2021, 12, 27, 22, 0, 0, 0, loc)},
			{d1, time.Hour * 3, time.Date(2021, 12, 27, 21, 0, 0, 0, loc)},
			{d1, time.Hour * 24, time.Date(2021, 12, 27, 0, 0, 0, 0, loc)},
			{d1, time.Hour * 24 * 2, time.Date(2021, 12, 27, 0, 0, 0, 0, loc)},
			{d1, time.Hour * 24 * 3, time.Date(2021, 12, 27, 0, 0, 0, 0, loc)},
			//
			{d2, time.Hour, time.Date(2018, 5, 1, 3, 0, 0, 0, loc)},
			{d2, time.Hour * 2, time.Date(2018, 5, 1, 2, 0, 0, 0, loc)},
			{d2, time.Hour * 3, time.Date(2018, 5, 1, 3, 0, 0, 0, loc)},
			{d2, time.Hour * 24, time.Date(2018, 5, 1, 0, 0, 0, 0, loc)},
			{d2, time.Hour * 24 * 2, time.Date(2018, 5, 1, 0, 0, 0, 0, loc)},
			{d2, time.Hour * 24 * 3, time.Date(2018, 4, 29, 0, 0, 0, 0, loc)},
		} {
			t.Run(fmt.Sprintf("%s_%s_%s", tc.giveTime.Format("20060102"), tc.giveDuration, LocationDuration(loc)), func(t *testing.T) {
				xtesting.Equal(t, TruncateTime(tc.giveTime, tc.giveDuration), tc.want)
			})
		}
	}
}

func TestDurationComponent(t *testing.T) {
	duration1 := 5*24*time.Hour + 5*time.Hour + 32*time.Minute + 24*time.Second + 123*time.Millisecond + 456*time.Microsecond + 789*time.Nanosecond
	duration2 := 105 * 24 * time.Hour
	duration3 := 24*time.Hour + 30*time.Minute + 1*time.Microsecond
	duration4 := 5*time.Hour + 30*time.Minute + 1*time.Microsecond
	duration5 := 0 * time.Second

	for _, tc := range []struct {
		give       time.Duration
		wantDay    int
		wantHour   int
		wantMinute int
		wantSecond int
		wantMs     int
		wantUs     int
		wantNs     int
	}{
		{duration1, 5, 5, 32, 24, 123, 456, 789},
		{duration2, 105, 0, 0, 0, 0, 0, 0},
		{duration3, 1, 0, 30, 0, 0, 1, 0},
		{duration4, 0, 5, 30, 0, 0, 1, 0},
		{duration5, 0, 0, 0, 0, 0, 0, 0},
	} {
		xtesting.Equal(t, DurationDayComponent(tc.give), tc.wantDay)
		xtesting.Equal(t, DurationHourComponent(tc.give), tc.wantHour)
		xtesting.Equal(t, DurationMinuteComponent(tc.give), tc.wantMinute)
		xtesting.Equal(t, DurationSecondComponent(tc.give), tc.wantSecond)
		xtesting.Equal(t, DurationMillisecondComponent(tc.give), tc.wantMs)
		xtesting.Equal(t, DurationMicrosecondComponent(tc.give), tc.wantUs)
		xtesting.Equal(t, DurationNanosecondComponent(tc.give), tc.wantNs)
	}
}

func TestDurationTotal(t *testing.T) {
	duration1 := 5*24*time.Hour + 5*time.Hour + 32*time.Minute + 24*time.Second + 123*time.Millisecond + 456*time.Microsecond + 789*time.Nanosecond
	duration2 := 105 * 24 * time.Hour
	duration3 := 24*time.Hour + 30*time.Minute + 1*time.Microsecond
	duration4 := 5*time.Hour + 30*time.Minute + 1*time.Microsecond
	duration5 := 0 * time.Second

	for _, tc := range []struct {
		give       time.Duration
		wantDay    float64
		wantHour   float64
		wantMinute float64
		wantSecond float64
		wantMs     int64
		wantUs     int64
		wantNs     int64
	}{
		{duration1, 5.2308, 125.5400, 7532.4020, 451944.123456789, 451944123, 451944123456, 451944123456789},
		{duration2, 105, 2520, 151200, 9072000, 9072000000, 9072000000000, 9072000000000000},
		{duration3, 1.0208, 24.5, 1470., 88200., 88200000, 88200000001, 88200000001000},
		{duration4, 0.2292, 5.5, 330., 19800, 19800000, 19800000001, 19800000001000},
		{duration5, 0, 0, 0, 0, 0, 0, 0},
	} {
		xtesting.Equal(t, DurationTotalNanoseconds(tc.give), tc.wantNs)
		xtesting.Equal(t, DurationTotalMicroseconds(tc.give), tc.wantUs)
		xtesting.Equal(t, DurationTotalMilliseconds(tc.give), tc.wantMs)
		xtesting.InDelta(t, DurationTotalSeconds(tc.give), tc.wantSecond, 1e-3)
		xtesting.InDelta(t, DurationTotalMinutes(tc.give), tc.wantMinute, 1e-3)
		xtesting.InDelta(t, DurationTotalHours(tc.give), tc.wantHour, 1e-3)
		xtesting.InDelta(t, DurationTotalDays(tc.give), tc.wantDay, 1e-3)
	}
}

func TestClock(t *testing.T) {
	// TODO
}
