package xtime

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
	"time"
)

func TestPadItoa(t *testing.T) {
	for _, tc := range []struct {
		giveNum   int
		giveDigit int
		giveSpace bool
		want      string
	}{
		{-1, 1, false, "???"},
		{-1, 2, false, "???"},
		{-1, 3, false, "???"},
		{-1, 4, false, "???"},
		{0, 1, false, "0"},
		{0, 2, false, "00"},
		{0, 3, false, "000"},
		{0, 4, false, "0000"},
		{10, 1, false, "10"},
		{10, 2, false, "10"},
		{10, 3, false, "010"},
		{10, 4, false, "0010"},
		{100, 1, false, "100"},
		{100, 2, false, "100"},
		{100, 3, false, "100"},
		{100, 4, false, "0100"},
		{1000, 1, false, "1000"},
		{1000, 2, false, "1000"},
		{1000, 3, false, "1000"},
		{1000, 4, false, "1000"},
		{1, 1, true, "1"},
		{1, 2, true, " 1"},
		{1, 3, true, "  1"},
		{1, 4, true, "   1"},
		{11, 1, true, "11"},
		{11, 2, true, "11"},
		{11, 3, true, " 11"},
		{11, 4, true, "  11"},
		{111, 1, true, "111"},
		{111, 2, true, "111"},
		{111, 3, true, "111"},
		{111, 4, true, " 111"},
		{1111, 1, true, "1111"},
		{1111, 2, true, "1111"},
		{1111, 3, true, "1111"},
		{1111, 4, true, "1111"},
	} {
		xtesting.Equal(t, _padItoa(tc.giveNum, tc.giveDigit, tc.giveSpace), tc.want)
	}
}

func TestInvalidStrftime(t *testing.T) {
	d := time.Date(2001, 1, 1, 1, 1, 1, 0, time.Local)
	for _, tc := range []struct {
		givePattern string
		wantError   bool
		wantResult  string
	}{
		{"", false, ""},
		{"test测试テスТест", false, "test测试テスТест"},
		{"%%", false, "%"},
		{"-%%", false, "-%"},
		{"%%%n%%%%", false, "%\n%%"},
		{"%n%%%%%t-%%-", false, "\n%%\t-%-"},
		{"%%%", true, ""},
		{"%%%-", true, ""},
		{"%S", false, "01"},
		{"%-S", false, "1"},
		{"%Y/%m/%d %H:%M:%S", false, "2001/01/01 01:01:01"},
		{"%-y年%-m月%-d日%h %-I时%-M分%-S秒%p", false, "1年1月1日Jan 1时1分1秒AM"},
		{"%f", true, ""},
		{"%-a", true, ""},
	} {
		t.Run(tc.givePattern, func(t *testing.T) {
			r, err := StrftimeInString(tc.givePattern, d)
			xtesting.Equal(t, err != nil, tc.wantError)
			if err == nil {
				xtesting.Equal(t, r, tc.wantResult)
			}
		})
	}
}

func TestStrftime(t *testing.T) {
	d1 := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2021, 12, 30, 23, 25, 39, 0, time.FixedZone("UTC+8", 8*60*60))
	d3 := time.Date(1901, 6, 9, 12, 1, 10, 0, time.FixedZone("UTC-9", -9*60*60))
	for _, tc := range []struct {
		giveDate    time.Time
		givePattern string
		wantResult  string
	}{
		{d1, "%Y %y %C %-Y %-y %-C", "2001 01 20 2001 1 20"},
		{d1, "%m %B %b %h %-m", "01 January Jan Jan 1"},
		{d1, "%d %e %A %a %-d", "01  1 Monday Mon 1"},
		{d1, "%H %k %I %l %-H %-I", "00  0 12 12 0 12"},
		{d1, "%p %P %M %S %-M %-S", "AM am 00 00 0 0"},
		{d1, "%Z %z %s %j %-j", "UTC +0000 978307200 001 1"},
		{d1, "%w %u %U %W %-U %-W", "1 1 01 01 1 1" /* "1 1 00 01 0 1" */},
		{d1, "%G %g %V %-G %-g %-V", "2001 01 01 2001 1 1"},
		{d1, "%c, %D, %F, %R, %r, %T, %v, %X, %x", "Mon Jan  1 00:00:00 2001, 01/01/01, 2001-01-01, 00:00, 12:00:00 AM, 00:00:00,  1-Jan-2001, 00:00:00, 01/01/01"},

		{d2, "%Y %y %C %-Y %-y %-C", "2021 21 20 2021 21 20"},
		{d2, "%m %B %b %h %-m", "12 December Dec Dec 12"},
		{d2, "%d %e %A %a %-d", "30 30 Thursday Thu 30"},
		{d2, "%H %k %I %l %-H %-I", "23 23 11 11 23 11"},
		{d2, "%p %P %M %S %-M %-S", "PM pm 25 39 25 39"},
		{d2, "%Z %z %s %j %-j", "UTC+8 +0800 1640877939 364 364"},
		{d2, "%w %u %U %W %-U %-W", "4 4 52 52 52 52"},
		{d2, "%G %g %V %-G %-g %-V", "2021 21 52 2021 21 52"},
		{d2, "%c, %D, %F, %R, %r, %T, %v, %X, %x", "Thu Dec 30 23:25:39 2021, 12/30/21, 2021-12-30, 23:25, 11:25:39 PM, 23:25:39, 30-Dec-2021, 23:25:39, 12/30/21"},

		{d3, "%Y %y %C %-Y %-y %-C", "1901 01 19 1901 1 19"},
		{d3, "%m %B %b %h %-m", "06 June Jun Jun 6"},
		{d3, "%d %e %A %a %-d", "09  9 Sunday Sun 9"},
		{d3, "%H %k %I %l %-H %-I", "12 12 12 12 12 12"},
		{d3, "%p %P %M %S %-M %-S", "PM pm 01 10 1 10"},
		{d3, "%Z %z %s %j %-j", "UTC-9 -0900 -2163639530 160 160"},
		{d3, "%w %u %U %W %-U %-W", "0 7 23 23 23 23" /* "0 7 23 22 23 22" */},
		{d3, "%G %g %V %-G %-g %-V", "1901 01 23 1901 1 23"},
		{d3, "%c, %D, %F, %R, %r, %T, %v, %X, %x", "Sun Jun  9 12:01:10 1901, 06/09/01, 1901-06-09, 12:01, 12:01:10 PM, 12:01:10,  9-Jun-1901, 12:01:10, 06/09/01"},
	} {
		t.Run(tc.givePattern, func(t *testing.T) {
			r, err := StrftimeInString(tc.givePattern, tc.giveDate)
			xtesting.Nil(t, err)
			xtesting.Equal(t, r, tc.wantResult)
		})
	}
}

func TestStrftimeToGlobPattern(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{"test测试テスТест", "test测试テスТест"},
		{"%%", "%"},
		{"-%%", "-%"},
		{"%%%n%%%%", "%*%%"},
		{"%n%%%%%t-%%-", "*%%*-%-"},
		{"%%%", "%%"},     // <- error
		{"%%-%-", "%-%-"}, // <- error
		{"%S", "*"},
		{"%A%f%F", "*"},
		{"-%S%-S-", "-*-"},
		{"%Y/%m/%d %H:%M:%S", "*/*/* *:*:*"},
		{"%Y%m%d %H%M%S", "* *"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, StrftimeToGlobPattern(tc.give), tc.want)
		})
	}
}
