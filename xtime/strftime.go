package xtime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	fsmNormal       = 0
	fsmPercent      = 1
	fsmPercentMinus = 2
)

// StrftimeInBytes formats given time.Time and strftime pattern to bytes, returns error when using invalid pattern, such as trailing '%'
// or wrong character after '%' or '%-', please visit https://docs.python.org/3/library/datetime.html#strftime-and-strptime-format-codes
// and https://man7.org/linux/man-pages/man3/strftime.3.html#DESCRIPTION and https://github.com/lestrrat-go/strftime for more details.
func StrftimeInBytes(pattern []byte, t time.Time) ([]byte, error) {
	bs := make([]byte, 0, len(pattern))
	fsm := fsmNormal
	for _, b := range pattern {
		if fsm == fsmNormal {
			if b == '%' {
				fsm = fsmPercent
			} else {
				bs = append(bs, b)
			}
		} else {
			var result string
			var err error
			if fsm == fsmPercent {
				if b == '-' {
					fsm = fsmPercentMinus
					continue
				} else {
					fsm = fsmNormal
					result, err = strftimeFormatVerbatim(b, t, false)
				}
			} else if fsm == fsmPercentMinus {
				fsm = fsmNormal
				result, err = strftimeFormatVerbatim(b, t, true)
			}
			if err != nil {
				return nil, err
			}
			bs = append(bs, result...)
		}
	}

	if fsm != fsmNormal {
		return nil, fmt.Errorf("invalid pattern '%s'", pattern)
	}
	return bs, nil
}

// StrftimeInString formats given time.Time and strftime pattern to string, returns error when using invalid pattern, such as trailing '%'
// or wrong character after '%' or '%-', please visit https://docs.python.org/3/library/datetime.html#strftime-and-strptime-format-codes
// and https://man7.org/linux/man-pages/man3/strftime.3.html#DESCRIPTION and https://github.com/lestrrat-go/strftime for more details.
func StrftimeInString(pattern string, t time.Time) (string, error) {
	bs, err := StrftimeInBytes(xstring.FastStob(pattern), t)
	if err != nil {
		return "", err
	}
	return xstring.FastBtos(bs), nil
}

// strftimeFormatVerbatim formats given character to result string in verbatim.
func strftimeFormatVerbatim(ch byte, t time.Time, nopad bool) (result string, err error) {
	var formatted string
	if !nopad {
		switch ch {
		case '%':
			formatted = "%"
		case 'n':
			formatted = "\n"
		case 't':
			formatted = "\t"
		case 'Y': // fourDigitYearZeroPad
			formatted = t.Format("2006")
		case 'y': // twoDigitYearZeroPad, 00-99
			formatted = t.Format("06")
		case 'C': // centuryDecimalZeroPad
			formatted = _padItoa(t.Year()/100, 2, false)
		case 'm': // monthNumberZeroPad, 01-12
			formatted = t.Format("01")
		case 'B': // fullMonthName
			formatted = t.Format("January")
		case 'b', 'h': // abbrMonthName
			formatted = t.Format("Jan")
		case 'd': // dayOfMonthZeroPad, 01-31
			formatted = t.Format("02")
		case 'e': // dayOfMonthSpacePad, _1-31
			formatted = t.Format("_2")
		case 'A': // fullWeekDayName
			formatted = t.Format("Monday")
		case 'a': // abbrWeekDayName
			formatted = t.Format("Mon")
		case 'H': // twentyFourHourClockZeroPad, 00-23
			formatted = t.Format("15")
		case 'k': // twentyFourHourClockSpacePad, _0-23
			formatted = _padItoa(t.Hour(), 2, true)
		case 'I': // twelveHourClockZeroPad, 01-12
			formatted = t.Format("03")
		case 'l': // twelveHourClockSpacePad, _1-12
			formatted = _padItoa(_twelveHour(t), 2, true)
		case 'p': // capitalAmpm
			formatted = t.Format("PM")
		case 'P': // lowercaseAmpm
			formatted = _lowercaseAmpm(t.Format("PM"))
		case 'M': // minutesZeroPad, 00-59
			formatted = t.Format("04")
		case 'S': // secondsNumberZeroPad, 00-60
			formatted = t.Format("05")
		case 'Z': // timezone
			formatted = t.Format("MST")
		case 'z': // timezoneOffset
			formatted = t.Format("-0700")
		case 's': // secondsSinceEpoch
			formatted = strconv.FormatInt(t.Unix(), 10)
		case 'j': // dayOfYearZeroPad, 001-366
			formatted = _padItoa(t.YearDay(), 3, false)
		case 'w': // weekdaySundayOrigin, 0-6
			formatted = strconv.Itoa(int(t.Weekday()))
		case 'u': // weekdayMondayOrigin, 1-7
			formatted = strconv.Itoa(_weekDayOffset(t, 1))
		case 'U': // weekNumberSundayOriginZeroPad, 00-53
			formatted = _padItoa(_weekNumberOffset(t, false), 2, false)
		case 'W': // weekNumberMondayOriginZeroPad, 00-53
			formatted = _padItoa(_weekNumberOffset(t, true), 2, false)
		case 'G': // fourDigitISO8601YearZeroPad
			formatted = _padItoa(_isoYear(t), 4, false)
		case 'g': // twoDigitISO8601YearZeroPad, 01-99
			formatted = _padItoa(_isoYear(t)%100, 2, false)
		case 'V': // iso8601WeekNumberZeroPad, 01-53
			formatted = _padItoa(_isoWeek(t), 2, false)
		case 'c': // timeAndDate
			formatted = t.Format("Mon Jan _2 15:04:05 2006")
		case 'D': // mdy
			formatted = t.Format("01/02/06")
		case 'F': // ymd
			formatted = t.Format("2006-01-02")
		case 'R': // hm
			formatted = t.Format("15:04")
		case 'r': // imsp
			formatted = t.Format("03:04:05 PM")
		case 'T': // hms
			formatted = t.Format("15:04:05")
		case 'v': // eby
			formatted = t.Format("_2-Jan-2006")
		case 'X': // natReprTime
			formatted = t.Format("15:04:05")
		case 'x': // natReprDate
			formatted = t.Format("01/02/06")
		default:
			err = fmt.Errorf("invalid pattern '%%%s'", string(ch))
		}
	} else {
		switch ch {
		case 'Y': // fourDigitYearNoPad
			formatted = strconv.Itoa(t.Year())
		case 'y': // twoDigitYearNoPad, 0-99
			formatted = strconv.Itoa(t.Year() % 100)
		case 'C': // centuryDecimalNoPad
			formatted = strconv.Itoa(t.Year() / 100)
		case 'm': // monthNumberNoPad, 1-12
			formatted = t.Format("1")
		case 'd': // dayOfMonthNoPad， 1-31
			formatted = t.Format("2")
		case 'H': // twentyFourHourClockNoPad, 0-23
			formatted = strconv.Itoa(t.Hour())
		case 'I': // twelveHourClockNoPad, 1-12
			formatted = t.Format("3")
		case 'M': // minutesNoPad, 0-59
			formatted = t.Format("4")
		case 'S': // secondsNumberNoPad, 0-60
			formatted = t.Format("5")
		case 'j': // dayOfYearNoPad, 1-366
			formatted = strconv.Itoa(t.YearDay())
		case 'U': // weekNumberSundayOriginNoPad, 0-53
			formatted = strconv.Itoa(_weekNumberOffset(t, false))
		case 'W': // weekNumberMondayOriginNoPad, 0-53
			formatted = strconv.Itoa(_weekNumberOffset(t, true))
		case 'G': // fourDigitISO8601YearNoPad
			formatted = strconv.Itoa(_isoYear(t))
		case 'g': // twoDigitISO8601YearNoPad, 1-99
			formatted = strconv.Itoa(_isoYear(t) % 100)
		case 'V': // iso8601WeekNumberNoPad, 1-53
			formatted = strconv.Itoa(_isoWeek(t))
		default:
			err = fmt.Errorf("invalid pattern '%%-%s'", string(ch))
		}
	}
	if err != nil {
		return "", err
	}
	return formatted, nil
}

func _padItoa(num int, digit int, space bool) string {
	s := strconv.Itoa(num)
	switch {
	case num >= 1000, num >= 100 && digit < 4, num >= 10 && digit < 3, num >= 0 && digit < 2:
		return s
	case num >= 100 && digit >= 4, num >= 10 && digit == 3, num >= 0 && digit == 2:
		if space {
			return " " + s
		}
		return "0" + s
	case num >= 10 && digit >= 4, num >= 0 && digit == 3:
		if space {
			return "  " + s
		}
		return "00" + s
	case num >= 0 && digit >= 4:
		if space {
			return "   " + s
		}
		return "000" + s
	default:
		return "???"
	}
}

func _twelveHour(t time.Time) int {
	h := t.Hour() % 12
	if h == 0 {
		h = 12
	}
	return h
}

func _lowercaseAmpm(s string) string {
	if s == "AM" {
		return "am"
	}
	return "pm"
}

// Note: some following code are referred from https://github.com/lestrrat-go/strftime.

func _weekDayOffset(t time.Time, offset int) int {
	wd := int(t.Weekday())
	if wd < offset {
		wd += 7
	}
	return wd
}

func _weekNumberOffset(t time.Time, mondayFirst bool) int {
	// https://github.com/lestrrat-go/strftime/blob/master/appenders.go#L248
	// https://opensource.apple.com/source/Libc/Libc-167/string.subproj/strftime.c.auto.html
	// https://github.com/arnoldrobbins/strftime/blob/master/strftime.c
	yd := t.YearDay()
	wd := int(t.Weekday())
	if !mondayFirst {
		return (yd + 7 - wd) / 7
	}
	if wd == 0 {
		return (yd + 7 - 6) / 7
	}
	return (yd + 7 - (wd - 1)) / 7
}

func _isoYear(t time.Time) int {
	n, _ := t.ISOWeek()
	return n
}

func _isoWeek(t time.Time) int {
	_, n := t.ISOWeek()
	return n
}

var (
	strftime2GlobPatternRe1 = regexp.MustCompile(`%-?[A-Za-z]`)
	strftime2GlobPatternRe2 = regexp.MustCompile(`\*+`)
)

// StrftimeToGlobPattern returns a corresponding glob pattern from strftime pattern.
func StrftimeToGlobPattern(pattern string) string {
	pattern = strings.ReplaceAll(pattern, "%%", "%")
	s := strftime2GlobPatternRe1.ReplaceAllString(pattern, "*")
	s = strftime2GlobPatternRe2.ReplaceAllString(s, "*")
	return s
}
