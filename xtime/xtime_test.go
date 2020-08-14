package xtime

import (
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	now := time.Now()
	t1 := NewJsonDate(now)
	t2 := NewJsonDateTime(now)
	assert.Equal(t, t1.Time(), ToDate(now))
	assert.Equal(t, t2.Time(), ToDateTime(now))

	log.Println(t1.String())
	log.Println(t2.String())
	assert.Equal(t, t1.String(), now.Format(RFC3339Date))
	assert.Equal(t, t2.String(), now.Format(RFC3339DateTime))

	bs1, _ := t1.MarshalJSON()
	bs2, _ := t2.MarshalJSON()
	assert.Equal(t, string(bs1), "\""+now.Format(RFC3339Date)+"\"")
	assert.Equal(t, string(bs2), "\""+now.Format(RFC3339DateTime)+"\"")
}

func TestParse(t *testing.T) {
	now := time.Now()
	t1 := NewJsonDate(now)
	t2 := NewJsonDateTime(now)

	tt1, _ := ParseRFC3339Date(now.Format(RFC3339Date))
	assert.Equal(t, tt1, t1)
	tt2, _ := ParseRFC3339DateTime(now.Format(RFC3339DateTime))
	assert.Equal(t, tt2, t2)
	tt3 := ParseRFC3339DateDefault(now.Format(RFC3339Date), t1)
	assert.Equal(t, tt3, t1)
	tt4 := ParseRFC3339DateTimeDefault(now.Format(RFC3339DateTime), t2)
	assert.Equal(t, tt4, t2)
}
