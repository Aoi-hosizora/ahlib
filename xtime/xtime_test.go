package xtime

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	now := time.Now()
	t1 := NewJsonDate(now)
	t2 := NewJsonDateTime(now)
	xtesting.Equal(t, t1.Time(), ToDate(now))
	xtesting.Equal(t, t2.Time(), ToDateTime(now))

	ds := now.Format(RFC3339Date)
	dts := now.Format(RFC3339DateTime)
	log.Println(ds, dts)
	log.Println(t1.Time(), t2.Time())

	log.Println(t1.String())
	log.Println(t2.String())
	xtesting.Equal(t, t1.String(), ds)
	xtesting.Equal(t, t2.String(), dts)

	bs1, _ := t1.MarshalJSON()
	bs2, _ := t2.MarshalJSON()
	xtesting.Equal(t, string(bs1), "\""+ds+"\"")
	xtesting.Equal(t, string(bs2), "\""+dts+"\"")
}

func TestParse(t *testing.T) {
	now := time.Now()
	now = now.In(time.UTC)
	t1 := NewJsonDate(now)
	t2 := NewJsonDateTime(now)

	ds := now.Format(RFC3339Date)
	dts := now.Format(RFC3339DateTime)
	log.Println(ds, dts)

	tt1, _ := ParseRFC3339Date(ds)
	xtesting.Equal(t, tt1, t1)
	log.Println(tt1.Time(), t1.Time())

	tt2, _ := ParseRFC3339DateTime(dts)
	xtesting.Equal(t, tt2, t2)
	log.Println(tt2.Time(), t2.Time())
	log.Println(t2.Time().Nanosecond(), tt2.Time().Nanosecond())

	tt3 := ParseRFC3339DateDefault(ds, t1)
	xtesting.Equal(t, tt3, t1)
	tt4 := ParseRFC3339DateTimeDefault(dts, t2)
	xtesting.Equal(t, tt4, t2)
}
