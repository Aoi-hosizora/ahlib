package xtime

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	now := time.Date(2020, time.Month(9), 30, 23, 39, 18, 5, time.FixedZone("", 8*60*60))
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
	zero = SetNanosecond(zero, now.Nanosecond())
	xtesting.Equal(t, zero.Format(time.RFC3339Nano), "2020-09-30T23:39:18.000000005Z")
	zero = SetLocation(zero, now.Location())
	xtesting.Equal(t, zero.Format(time.RFC3339), "2020-09-30T23:39:18+08:00")

	xtesting.Equal(t, zero, now)
}

func TestTo(t *testing.T) {
	now := time.Now()

	date := ToDate(now)
	xtesting.Equal(t, date.Year(), now.Year())
	xtesting.Equal(t, date.Month(), now.Month())
	xtesting.Equal(t, date.Day(), now.Day())
	xtesting.Equal(t, date.Hour(), 0)
	xtesting.Equal(t, date.Minute(), 0)
	xtesting.Equal(t, date.Second(), 0)
	xtesting.Equal(t, date.Nanosecond(), 0)
	xtesting.Equal(t, date.Location().String(), time.UTC.String())

	datetime := ToDateTime(now)
	xtesting.Equal(t, datetime.Year(), now.Year())
	xtesting.Equal(t, datetime.Month(), now.Month())
	xtesting.Equal(t, datetime.Day(), now.Day())
	xtesting.Equal(t, datetime.Hour(), now.Hour())
	xtesting.Equal(t, datetime.Minute(), now.Minute())
	xtesting.Equal(t, datetime.Second(), now.Second())
	xtesting.Equal(t, datetime.Nanosecond(), 0)
	xtesting.Equal(t, datetime.Location().String(), "")
}

func TestNew(t *testing.T) {
	_ = RFC3339Date
	_ = LocalDate
	_ = RFC3339DateTime
	_ = LocalDateTime

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

func TestParse(t *testing.T) {
	now := time.Now()
	// now = now.In(time.UTC)

	date := NewJsonDate(now)
	dateTime := NewJsonDateTime(now)
	dateStr := date.String()
	dateTimeStr := dateTime.String()

	date1, _ := ParseRFC3339Date(dateStr)
	xtesting.Equal(t, date1, date)
	date2 := ParseRFC3339DateDefault(dateStr, date)
	xtesting.Equal(t, date2, date)
	date3 := ParseRFC3339DateDefault("", date)
	xtesting.Equal(t, date3, date)

	dateTime1, _ := ParseRFC3339DateTime(dateTimeStr)
	xtesting.Equal(t, dateTime1, dateTime)
	log.Println(dateTime1.String(), dateTime1.Time().Location())
	log.Println(dateTime.String(), dateTime.Time().Location())
	dateTime2 := ParseRFC3339DateTimeDefault(dateTimeStr, dateTime)
	xtesting.Equal(t, dateTime2, dateTime)
	dateTime3 := ParseRFC3339DateTimeDefault("", dateTime)
	xtesting.Equal(t, dateTime3, dateTime)
}

func TestGorm(t *testing.T) {
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
