package xzone

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestTimeZone(t *testing.T) {
	re, err := regexp.Compile(`^([+-])([0-9]{1,2})(?::([0-9]{1,2}))?$`)
	if err != nil {
		log.Fatalln(err)
	}

	xtesting.Equal(t, re.Match([]byte("+")), false)
	xtesting.Equal(t, re.Match([]byte("+0")), true)
	xtesting.Equal(t, re.Match([]byte("+09")), true)
	xtesting.Equal(t, re.Match([]byte("+09:")), false)
	xtesting.Equal(t, re.Match([]byte("-9:0")), true)
	xtesting.Equal(t, re.Match([]byte("+9:00")), true)
	xtesting.Equal(t, re.Match([]byte("-09:0")), true)
	xtesting.Equal(t, re.Match([]byte("+09:30")), true)

	log.Println(re.FindAllStringSubmatch("+0", 1)[0][1:])     // + 0
	log.Println(re.FindAllStringSubmatch("+09", 1)[0][1:])    // + 09
	log.Println(re.FindAllStringSubmatch("-9:0", 1)[0][1:])   // - 9 0
	log.Println(re.FindAllStringSubmatch("+9:00", 1)[0][1:])  // + 9 00
	log.Println(re.FindAllStringSubmatch("-09:0", 1)[0][1:])  // - 09 0
	log.Println(re.FindAllStringSubmatch("+09:30", 1)[0][1:]) // + 09 30
}

func TestParseTimeZone(t *testing.T) {
	_, err := ParseTimeZone("+")
	xtesting.NotEqual(t, err, nil)

	loc, err := ParseTimeZone("+0")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC+00:00")

	loc, err = ParseTimeZone("+09")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimeZone("+09:")
	xtesting.NotEqual(t, err, nil)

	loc, err = ParseTimeZone("-9:0")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimeZone("+9:00")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimeZone("-09:0")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimeZone("+09:30")
	xtesting.Equal(t, err, nil)
	xtesting.Equal(t, loc.String(), "UTC+09:30")
}

func TestMoveToZone(t *testing.T) {
	tt, _ := time.Parse(time.RFC3339, "2020-08-06T12:46:43+08:00")

	tt2, _ := MoveToZone(tt, "+8")
	xtesting.Equal(t, tt2.Hour(), 12)
	xtesting.Equal(t, tt2.Minute(), 46)

	tt2, _ = MoveToZone(tt, "+09:00")
	xtesting.Equal(t, tt2.Hour(), 13)
	xtesting.Equal(t, tt2.Minute(), 46)

	tt2, _ = MoveToZone(tt, "-00:30")
	xtesting.Equal(t, tt2.Hour(), 4)
	xtesting.Equal(t, tt2.Minute(), 16)
}
