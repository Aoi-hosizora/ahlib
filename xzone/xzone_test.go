package xzone

import (
	"github.com/go-playground/assert/v2"
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

	assert.Equal(t, re.Match([]byte("+")), false)
	assert.Equal(t, re.Match([]byte("+0")), true)
	assert.Equal(t, re.Match([]byte("+09")), true)
	assert.Equal(t, re.Match([]byte("+09:")), false)
	assert.Equal(t, re.Match([]byte("-9:0")), true)
	assert.Equal(t, re.Match([]byte("+9:00")), true)
	assert.Equal(t, re.Match([]byte("-09:0")), true)
	assert.Equal(t, re.Match([]byte("+09:30")), true)

	log.Println(re.FindAllStringSubmatch("+0", 1)[0][1:])     // + 0
	log.Println(re.FindAllStringSubmatch("+09", 1)[0][1:])    // + 09
	log.Println(re.FindAllStringSubmatch("-9:0", 1)[0][1:])   // - 9 0
	log.Println(re.FindAllStringSubmatch("+9:00", 1)[0][1:])  // + 9 00
	log.Println(re.FindAllStringSubmatch("-09:0", 1)[0][1:])  // - 09 0
	log.Println(re.FindAllStringSubmatch("+09:30", 1)[0][1:]) // + 09 30
}

func TestParseTimeZone(t *testing.T) {
	_, err := ParseTimeZone("+")
	assert.NotEqual(t, err, nil)

	loc, err := ParseTimeZone("+0")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC+00:00")

	loc, err = ParseTimeZone("+09")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimeZone("+09:")
	assert.NotEqual(t, err, nil)

	loc, err = ParseTimeZone("-9:0")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimeZone("+9:00")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC+09:00")

	loc, err = ParseTimeZone("-09:0")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC-09:00")

	loc, err = ParseTimeZone("+09:30")
	assert.Equal(t, err, nil)
	assert.Equal(t, loc.String(), "UTC+09:30")
}

func TestMoveToZone(t *testing.T) {
	tt, _ := time.Parse(time.RFC3339, "2020-08-06T12:46:43+08:00")

	tt2, _ := MoveToZone(tt, "+8")
	assert.Equal(t, tt2.Hour(), 12)
	assert.Equal(t, tt2.Minute(), 46)

	tt2, _ = MoveToZone(tt, "+09:00")
	assert.Equal(t, tt2.Hour(), 13)
	assert.Equal(t, tt2.Minute(), 46)

	tt2, _ = MoveToZone(tt, "-00:30")
	assert.Equal(t, tt2.Hour(), 4)
	assert.Equal(t, tt2.Minute(), 16)
}
