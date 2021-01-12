package xtime

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"github.com/Aoi-hosizora/ahlib/xtime/zones"
	"log"
	"testing"
	"time"
)

func TestSetXXX(t *testing.T) {
	now := time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("", 8*60*60))
	zero := time.Time{}
	xtesting.Equal(t, zero.Format(time.RFC3339), "0001-01-01T00:00:00Z")

	zero = SetYear(zero, now.Year())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-01-01T00:00:00Z")
	zero = SetMonth(zero, int(now.Month()))
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-01T00:00:00Z")
	zero = SetDay(zero, now.Day())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T00:00:00Z")
	zero = SetHour(zero, now.Hour())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T23:00:00Z")
	zero = SetMinute(zero, now.Minute())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T23:39:00Z")
	zero = SetSecond(zero, now.Second())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T23:39:18Z")
	zero = SetMillisecond(zero, 123)
	xtesting.Equal(t, zero.Format(time.RFC3339Nano), "2020-09-30T23:39:18.123Z")
	zero = SetMicrosecond(zero, 123456)
	xtesting.Equal(t, zero.Format(time.RFC3339Nano), "2020-09-30T23:39:18.123456Z")
	zero = SetNanosecond(zero, now.Nanosecond())
	xtesting.Equal(t, zero.Format(time.RFC3339Nano), "2020-09-30T23:39:18.000000789Z")
	zero = SetLocation(zero, now.Location())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T23:39:18+08:00")

	xtesting.Equal(t, zero, now)
}

func TestToXXX(t *testing.T) {
	now := time.Now()

	date := ToDate(now)
	xtesting.Equal(t, date.Year(), now.Year())
	xtesting.Equal(t, date.Month(), now.Month())
	xtesting.Equal(t, date.Day(), now.Day())
	xtesting.Equal(t, date.Hour(), 0)
	xtesting.Equal(t, date.Minute(), 0)
	xtesting.Equal(t, date.Second(), 0)
	xtesting.Equal(t, date.Nanosecond(), 0)
	xtesting.Equal(t, date.Location().String(), "")
	xtesting.Equal(t, LocationDuration(date.Location()), LocationDuration(now.Location()))

	datetime := ToDateTime(now)
	xtesting.Equal(t, datetime.Year(), now.Year())
	xtesting.Equal(t, datetime.Month(), now.Month())
	xtesting.Equal(t, datetime.Day(), now.Day())
	xtesting.Equal(t, datetime.Hour(), now.Hour())
	xtesting.Equal(t, datetime.Minute(), now.Minute())
	xtesting.Equal(t, datetime.Second(), now.Second())
	xtesting.Equal(t, datetime.Nanosecond(), 0)
	xtesting.Equal(t, datetime.Location().String(), "")
	xtesting.Equal(t, LocationDuration(datetime.Location()), LocationDuration(now.Location()))
}

func TestLocationDuration(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52Z")      // UTC
	t2, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52-07:00") // ""
	t3, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52+08:00") // Local
	t4, _ := time.Parse(time.RFC3339, "2020-09-30T23:56:52+09:00") // ""

	xtesting.Equal(t, int(LocationDuration(t1.Location()).Seconds()), 0)
	xtesting.Equal(t, int(LocationDuration(t2.Location()).Seconds()), -7*3600)
	xtesting.Equal(t, int(LocationDuration(t3.Location()).Seconds()), 8*3600)
	xtesting.Equal(t, int(LocationDuration(t4.Location()).Seconds()), 9*3600)

	t1l := GetTimeLocation(t1) // "" +00:00
	t2l := GetTimeLocation(t2) // "" -07:00
	t3l := GetTimeLocation(t3) // "" +08:00
	t4l := GetTimeLocation(t4) // "" +09:00
	xtesting.Equal(t, t1l.String(), "")
	xtesting.Equal(t, t2l.String(), "")
	xtesting.Equal(t, t3l.String(), "")
	xtesting.Equal(t, t4l.String(), "")
	xtesting.Equal(t, int(LocationDuration(t1l).Seconds()), 0)
	xtesting.Equal(t, int(LocationDuration(t2l).Seconds()), -7*3600)
	xtesting.Equal(t, int(LocationDuration(t3l).Seconds()), 8*3600)
	xtesting.Equal(t, int(LocationDuration(t4l).Seconds()), 9*3600)
}

func TestTimezone(t *testing.T) {
	re := timezoneRegexp
	xtesting.Equal(t, re.Match([]byte("+")), false)
	xtesting.Equal(t, re.Match([]byte("+0")), true)
	xtesting.Equal(t, re.Match([]byte("+09")), true)
	xtesting.Equal(t, re.Match([]byte("+09:")), false)
	xtesting.Equal(t, re.Match([]byte("-9:0")), true)
	xtesting.Equal(t, re.Match([]byte("+9:00")), true)
	xtesting.Equal(t, re.Match([]byte("-09:0")), true)
	xtesting.Equal(t, re.Match([]byte("+09:30")), true)

	_, err := ParseTimezone("+")
	xtesting.NotNil(t, err)

	loc, err := ParseTimezone("+0")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC+00:00")

	loc, err = ParseTimezone("+09")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimezone("+09:")
	xtesting.NotNil(t, err)

	loc, err = ParseTimezone("-9:0")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimezone("+9:00")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimezone("-09:0")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimezone("+09:30")
	xtesting.Nil(t, err)
	xtesting.Equal(t, loc.String(), "UTC+09:30")

	_, err = MoveToTimezone(time.Now(), "")
	xtesting.NotNil(t, err)

	tt, _ := time.Parse(time.RFC3339, "2020-08-06T12:46:43+08:00")

	tt2, _ := MoveToTimezone(tt, "+8")
	xtesting.Equal(t, tt2.Hour(), 12)
	xtesting.Equal(t, tt2.Minute(), 46)

	tt3, _ := MoveToTimezone(tt, "+09:00")
	xtesting.Equal(t, tt3.Hour(), 13)
	xtesting.Equal(t, tt3.Minute(), 46)

	tt4, _ := MoveToTimezone(tt, "-00:30")
	xtesting.Equal(t, tt4.Hour(), 4)
	xtesting.Equal(t, tt4.Minute(), 16)

	_, err = MoveToLocation(time.Now(), "x")
	xtesting.NotNil(t, err)
	_, err = MoveToLocation(time.Now(), "Asia/Shangha")
	xtesting.NotNil(t, err)

	tt5, _ := MoveToLocation(tt, zones.Asia_Tokyo)
	xtesting.Equal(t, tt5.Hour(), 13)
	xtesting.Equal(t, tt5.Minute(), 46)

	tt6, _ := MoveToLocation(tt, zones.Asia_Shanghai)
	xtesting.Equal(t, tt6.Hour(), 12)
	xtesting.Equal(t, tt6.Minute(), 46)

	tt7, _ := MoveToLocation(tt, "")
	xtesting.Equal(t, tt7.Hour(), 4)
	xtesting.Equal(t, tt7.Minute(), 46)
}

func TestNewDateAndNewDateTime(t *testing.T) {
	_ = RFC3339Date
	_ = CJKDate
	_ = RFC3339DateTime
	_ = CJKDateTime

	now := time.Now()

	date := NewJsonDate(now)
	dateTime := NewJsonDateTime(now)
	xtesting.Equal(t, date.Time(), ToDate(now))
	xtesting.Equal(t, dateTime.Time(), ToDateTime(now))

	dateStr := now.Format(RFC3339Date)
	dateTimeStr := now.Format(RFC3339DateTime)
	xtesting.Equal(t, date.String(), dateStr)
	xtesting.Equal(t, dateTime.String(), dateTimeStr)

	bs1, _ := date.MarshalJSON()
	bs2, _ := dateTime.MarshalJSON()
	xtesting.Equal(t, string(bs1), "\""+dateStr+"\"")
	xtesting.Equal(t, string(bs2), "\""+dateTimeStr+"\"")
}

func TestParseDateAndParseDateTime(t *testing.T) {
	now := time.Now()
	// now = now.In(time.UTC)

	date := NewJsonDate(now)
	dateTime := NewJsonDateTime(now)
	dateStr := date.String()
	dateTimeStr := dateTime.String()

	date1, _ := ParseJsonDate(dateStr)
	xtesting.Equal(t, date1.Time(), date.Time())
	date2 := ParseJsonDateOr(dateStr, date)
	xtesting.Equal(t, date2.Time(), date.Time())
	date3 := ParseJsonDateOr("", date)
	xtesting.Equal(t, date3, date)

	dateTime1, _ := ParseJsonDateTime(dateTimeStr)
	xtesting.Equal(t, dateTime1, dateTime)
	log.Println(dateTime1.String(), dateTime1.Time().Location())
	log.Println(dateTime.String(), dateTime.Time().Location())
	dateTime2 := ParseJsonDateTimeOr(dateTimeStr, dateTime)
	xtesting.Equal(t, dateTime2, dateTime)
	dateTime3 := ParseJsonDateTimeOr("", dateTime)
	xtesting.Equal(t, dateTime3, dateTime)
}

func TestDateAndDateTimeScanValue(t *testing.T) {
	now := time.Now()
	now2 := time.Now()
	now2.Add(time.Hour * 24)

	date := NewJsonDate(now)
	xtesting.Equal(t, date.Time(), ToDate(now))
	val, err := date.Value()
	xtesting.Nil(t, err)
	xtesting.Equal(t, val, ToDate(now))

	err = date.Scan(ToDate(now2))
	xtesting.Nil(t, err)
	xtesting.Equal(t, date.Time(), ToDate(now2))
	err = date.Scan("")
	xtesting.NotNil(t, err)
	xtesting.Equal(t, date.Time(), ToDate(now2))
	err = date.Scan(nil)
	xtesting.Nil(t, err)

	dateTime := NewJsonDateTime(now)
	xtesting.Equal(t, dateTime.Time(), ToDateTime(now))
	val, err = dateTime.Value()
	xtesting.Nil(t, err)
	xtesting.Equal(t, val, ToDateTime(now))

	err = dateTime.Scan(ToDateTime(now2))
	xtesting.Nil(t, err)
	xtesting.Equal(t, dateTime.Time(), ToDateTime(now2))
	err = dateTime.Scan("")
	xtesting.NotNil(t, err)
	xtesting.Equal(t, dateTime.Time(), ToDateTime(now2))
	err = dateTime.Scan(nil)
	xtesting.Nil(t, err)
}

func TestTimeSpan(t *testing.T) {
	ts := NewTimeSpan(5*24*time.Hour + 5*time.Hour + 32*time.Minute + 24*time.Second + 123*time.Millisecond + 456*time.Microsecond + 789*time.Nanosecond)
	xtesting.Equal(t, ts, 5*Day+5*Hour+32*Minute+24*Second+123*Millisecond+456*Microsecond+789*Nanosecond)
	xtesting.Equal(t, ts.String(), "5d5h32m24.123456789s")

	xtesting.Equal(t, ts.Days(), 5)
	xtesting.Equal(t, ts.Hours(), 5)
	xtesting.Equal(t, ts.Minutes(), 32)
	xtesting.Equal(t, ts.Seconds(), 24)
	xtesting.Equal(t, ts.Milliseconds(), 123)
	xtesting.Equal(t, ts.Microseconds(), 456)
	xtesting.Equal(t, ts.Nanoseconds(), 789)

	xtesting.Equal(t, ts.TotalNanoseconds(), int64(451944123456789))
	xtesting.Equal(t, ts.TotalMicroseconds(), int64(451944123456))
	xtesting.Equal(t, ts.TotalMilliseconds(), int64(451944123))
	xtesting.InDelta(t, ts.TotalSeconds(), 451944.123456789, 1e-3)
	xtesting.InDelta(t, ts.TotalMinutes(), 7532.402057613, 1e-3)
	xtesting.InDelta(t, ts.TotalHours(), 125.540034293, 1e-3)
	xtesting.InDelta(t, ts.TotalDays(), 5.230834762, 1e-3)

	ts = -ts
	xtesting.Equal(t, ts, -(5*Day + 5*Hour + 32*Minute + 24*Second + 123*Millisecond + 456*Microsecond + 789*Nanosecond))
	xtesting.Equal(t, ts.String(), "-5d5h32m24.123456789s")

	xtesting.Equal(t, ts.Days(), -5)
	xtesting.Equal(t, ts.Hours(), -5)
	xtesting.Equal(t, ts.Minutes(), -32)
	xtesting.Equal(t, ts.Seconds(), -24)
	xtesting.Equal(t, ts.Milliseconds(), -123)
	xtesting.Equal(t, ts.Microseconds(), -456)
	xtesting.Equal(t, ts.Nanoseconds(), -789)

	xtesting.Equal(t, ts.TotalNanoseconds(), int64(-451944123456789))
	xtesting.Equal(t, ts.TotalMicroseconds(), int64(-451944123456))
	xtesting.Equal(t, ts.TotalMilliseconds(), int64(-451944123))
	xtesting.InDelta(t, ts.TotalSeconds(), -451944.123456789, 1e-3)
	xtesting.InDelta(t, ts.TotalMinutes(), -7532.402057613, 1e-3)
	xtesting.InDelta(t, ts.TotalHours(), -125.540034293, 1e-3)
	xtesting.InDelta(t, ts.TotalDays(), -5.230834762, 1e-3)

	ts = TimeSpan(0)
	xtesting.Equal(t, ts.Days(), 0)
	xtesting.Equal(t, ts.Hours(), 0)
	xtesting.Equal(t, ts.Minutes(), 0)
	xtesting.Equal(t, ts.Seconds(), 0)
	xtesting.Equal(t, ts.Milliseconds(), 0)
	xtesting.Equal(t, ts.Microseconds(), 0)
	xtesting.Equal(t, ts.Nanoseconds(), 0)
	xtesting.Equal(t, ts.TotalDays(), 0.)
	xtesting.Equal(t, ts.TotalHours(), 0.)
	xtesting.Equal(t, ts.TotalMinutes(), 0.)
	xtesting.Equal(t, ts.TotalSeconds(), 0.)
	xtesting.Equal(t, ts.TotalMilliseconds(), int64(0))
	xtesting.Equal(t, ts.TotalMicroseconds(), int64(0))
	xtesting.Equal(t, ts.TotalNanoseconds(), int64(0))
	xtesting.Equal(t, ts.String(), "0s")

	xtesting.Equal(t, (Day + 500*Nanosecond).String(), "1d0h0m0.000000500s")
	xtesting.Equal(t, (Day + 5*Second).String(), "1d0h0m5s")
	xtesting.Equal(t, (Hour + 5*Second).String(), "1h0m5s")
	xtesting.Equal(t, (-Second - 5*Nanosecond).String(), "-1.000000005s")

	ts = 5 * Nanosecond
	xtesting.Equal(t, ts.Add(Microsecond), Microsecond+5*Nanosecond)
	xtesting.Equal(t, ts.Sub(Microsecond), -995*Nanosecond)

	v, _ := ts.Value()
	xtesting.Equal(t, v, int64(5))
	err := ts.Scan(nil)
	xtesting.Nil(t, err)
	xtesting.Equal(t, ts.TotalNanoseconds(), int64(5))
	err = ts.Scan(int64(6))
	xtesting.Nil(t, err)
	xtesting.Equal(t, ts.TotalNanoseconds(), int64(6))
	err = ts.Scan("")
	xtesting.NotNil(t, err)
}
