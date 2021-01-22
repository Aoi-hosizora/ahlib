package xtime

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
	"time"
)

func TestNewJsonDateAndDateTime(t *testing.T) {
	_ = RFC3339Date
	_ = CJKDate
	_ = RFC3339DateTime
	_ = CJKDateTime

	for _, tc := range []struct {
		give         time.Time
		wantDate     time.Time
		wantDateTime time.Time
	}{
		{time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("test", 8*60*60)),
			time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, time.FixedZone("", 8*60*60)),
			time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*60*60))},

		{time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("", 0)),
			time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("", 0))},

		{time.Date(2, time.Month(2), 2, 0, 0, 1, 1, time.FixedZone("test_test", -9*3600-30*60)),
			time.Date(2, time.Month(2), 2, 0, 0, 0, 0, time.FixedZone("", -9*3600-30*60)),
			time.Date(2, time.Month(2), 2, 0, 0, 1, 0, time.FixedZone("", -9*3600-30*60))},
	} {
		date := NewJsonDate(tc.give)
		dateTime := NewJsonDateTime(tc.give)

		xtesting.Equal(t, date.Time(), tc.wantDate)
		xtesting.Equal(t, dateTime.Time(), tc.wantDateTime)
	}

	now := time.Now()
	xtesting.Equal(t, NewJsonDate(now), NewJsonDate(now))
	xtesting.Equal(t, NewJsonDateTime(now), NewJsonDateTime(now))
	xtesting.Equal(t, NewJsonDate(now).Time(), time.Time(NewJsonDate(now)))
	xtesting.Equal(t, NewJsonDateTime(now).Time(), time.Time(NewJsonDateTime(now)))
	xtesting.Equal(t, NewJsonDate(now).Time(), ToDate(now))
	xtesting.Equal(t, NewJsonDateTime(now).Time(), ToDateTime(now))
}

func TestMarshalJsonDateAndDateTime(t *testing.T) {
	for _, tc := range []struct {
		give           time.Time
		wantDate       string
		wantDateTime   string
		wantDateBs     []byte
		wantDateTimeBs []byte
	}{
		{time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("test", 8*60*60)),
			"2020-09-30", "2020-09-30T23:39:18+08:00", []byte(`"2020-09-30"`), []byte(`"2020-09-30T23:39:18+08:00"`)},

		{time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			"0001-01-01", "0001-01-01T00:00:00Z", []byte(`"0001-01-01"`), []byte(`"0001-01-01T00:00:00Z"`)},

		{time.Date(2, time.Month(2), 2, 0, 0, 1, 1, time.FixedZone("test_test", -9*3600-30*60)),
			"0002-02-02", "0002-02-02T00:00:01-09:30", []byte(`"0002-02-02"`), []byte(`"0002-02-02T00:00:01-09:30"`)},
	} {
		date := NewJsonDate(tc.give)
		dateTime := NewJsonDateTime(tc.give)

		xtesting.Equal(t, date.String(), tc.wantDate)
		xtesting.Equal(t, dateTime.String(), tc.wantDateTime)

		bs, err := date.MarshalJSON()
		xtesting.Equal(t, bs, tc.wantDateBs)
		xtesting.Nil(t, err, nil)
		bs, err = dateTime.MarshalJSON()
		xtesting.Equal(t, bs, tc.wantDateTimeBs)
		xtesting.Nil(t, err, nil)
	}
}

func TestUnmarshalJsonDateAndDateTime(t *testing.T) {
	for _, tc := range []struct {
		give           string
		wantDateOk     bool
		wantDate       JsonDate
		wantDateTimeOk bool
		wantDateTime   JsonDateTime
	}{
		{``, false, JsonDate{}, false, JsonDateTime{}},
		{`"null"`, false, JsonDate{}, false, JsonDateTime{}},
		{`2020-09-30`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09-30`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09"`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09-30Z"`, false, JsonDate{}, false, JsonDateTime{}},
		{`2020-09-30T23:39:+08:00`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09-30T23:39:18+08:00`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09-30T23:39"`, false, JsonDate{}, false, JsonDateTime{}},
		{`"2020-09-30T23:39:18"`, false, JsonDate{}, false, JsonDateTime{}},

		{`null`, true, JsonDate{}, true, JsonDateTime{}},
		{`"2020-09-30"`, true, JsonDate(time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, GetLocalLocation())),
			false, JsonDateTime{}},
		{`"2020-09-30T23:39:18Z"`, false, JsonDate{}, true,
			JsonDateTime(time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 0)))},
		{`"2020-09-30T23:39:18+08:00"`, false, JsonDate{}, true,
			JsonDateTime(time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*3600)))},
	} {
		d := JsonDate{}
		err := d.UnmarshalJSON([]byte(tc.give))
		if tc.wantDateOk {
			xtesting.Nil(t, err)
			xtesting.Equal(t, d, tc.wantDate)
		} else {
			xtesting.NotNil(t, err)
		}

		dt := JsonDateTime{}
		err = dt.UnmarshalJSON([]byte(tc.give))
		if tc.wantDateTimeOk {
			xtesting.Nil(t, err)
			xtesting.Equal(t, dt.Time(), tc.wantDateTime.Time())
		} else {
			xtesting.NotNil(t, err)
		}
	}
}

func TestScanValueJsonDateAndDateTime(t *testing.T) {
	for _, tc := range []struct {
		give              interface{}
		wantError         bool
		wantDate          JsonDate
		wantDateTime      JsonDateTime
		wantDateValue     time.Time
		wantDateTimeValue time.Time
	}{
		{nil, false, JsonDate{}, JsonDateTime{}, time.Time{}, time.Time{}},
		{0, true, JsonDate{}, JsonDateTime{}, time.Time{}, time.Time{}},
		{time.Date(2020, time.Month(9), 30, 23, 39, 18, 789, time.FixedZone("test", 8*60*60)), false,
			JsonDate(time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, time.FixedZone("", 8*60*60))),
			JsonDateTime(time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*60*60))),
			time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, time.FixedZone("", 8*60*60)),
			time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*60*60))},
		{time.Date(1, time.Month(1), 1, 1, 1, 1, 1, time.UTC), false,
			JsonDate(time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("", 0))),
			JsonDateTime(time.Date(1, time.Month(1), 1, 1, 1, 1, 0, time.FixedZone("", 0))),
			time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("", 0)),
			time.Date(1, time.Month(1), 1, 1, 1, 1, 0, time.FixedZone("", 0))},
	} {
		d := JsonDate{}
		err := d.Scan(tc.give)
		if tc.wantError {
			xtesting.NotNil(t, err)
		} else {
			xtesting.Equal(t, d, tc.wantDate)
			xtesting.Nil(t, err)

			val, err := d.Value()
			xtesting.Equal(t, val, tc.wantDateValue)
			xtesting.Nil(t, err)
		}

		dt := JsonDateTime{}
		err = dt.Scan(tc.give)
		if tc.wantError {
			xtesting.NotNil(t, err)
		} else {
			xtesting.Equal(t, dt, tc.wantDateTime)
			xtesting.Nil(t, err)

			val, err := dt.Value()
			xtesting.Equal(t, val, tc.wantDateTimeValue)
			xtesting.Nil(t, err)
		}
	}
}

func TestParseJsonDateAndDateTime(t *testing.T) {
	for _, tc := range []struct {
		give           string
		wantDateOk     bool
		wantDate       JsonDate
		wantDateTimeOk bool
		wantDateTime   JsonDateTime
	}{
		{"", false, JsonDate{}, false, JsonDateTime{}},
		{"2020-09", false, JsonDate{}, false, JsonDateTime{}},
		{"2020-09-30Z", false, JsonDate{}, false, JsonDateTime{}},
		{"2020-09-30T23:39", false, JsonDate{}, false, JsonDateTime{}},
		{"2020-09-30T23:39:18", false, JsonDate{}, false, JsonDateTime{}},
		{"2020-09-30", true, JsonDate(time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, GetLocalLocation())),
			false, JsonDateTime{}},
		{"2020-09-30T23:39:18Z", false, JsonDate{}, true,
			JsonDateTime(time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 0)))},
		{"2020-09-30T23:39:18+08:00", false, JsonDate{}, true,
			JsonDateTime(time.Date(2020, time.Month(9), 30, 23, 39, 18, 0, time.FixedZone("", 8*3600)))},
	} {
		d, err := ParseJsonDate(tc.give)
		if !tc.wantDateOk {
			xtesting.NotNil(t, err)
			now := NewJsonDate(time.Now())
			xtesting.Equal(t, ParseJsonDateOr(tc.give, now), now)
		} else {
			xtesting.Nil(t, err)
			xtesting.Equal(t, d, tc.wantDate)
			now := NewJsonDate(time.Now())
			xtesting.Equal(t, ParseJsonDateOr(tc.give, now), tc.wantDate)
		}

		dt, err := ParseJsonDateTime(tc.give)
		if !tc.wantDateTimeOk {
			xtesting.NotNil(t, err)
			now := NewJsonDateTime(time.Now())
			xtesting.Equal(t, ParseJsonDateTimeOr(tc.give, now), now)
		} else {
			xtesting.Nil(t, err)
			xtesting.Equal(t, dt, tc.wantDateTime)
			now := NewJsonDateTime(time.Now())
			xtesting.Equal(t, ParseJsonDateTimeOr(tc.give, now), tc.wantDateTime)
		}
	}
}
