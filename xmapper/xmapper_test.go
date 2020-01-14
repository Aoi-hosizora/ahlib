package xmapper

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"testing"
	"time"
)

type Po struct {
	DefaultField float32
	FirstName    string
	LastName     string
	Birth        time.Time
}

type Dto struct {
	DefaultField float32
	Name         string
	Age          int
}

func TestEntitiesMapper_Map(t *testing.T) {
	_mapper := CreateMapper(&Po{}, &Dto{}).
		ForMember("Name", func(po interface{}) interface{} {
			return po.(Po).FirstName + " " + po.(Po).LastName
		}).
		ForMember("Age", func(po interface{}) interface{} {
			return time.Now().Year() - po.(Po).Birth.Year()
		}).
		Build().
		CreateMapper(&Po{}, &Dto{}).ForMember("Name", func(po interface{}) interface{} { return "" }).Build()

	po := &Po{
		DefaultField: 0.333,
		FirstName:    "First",
		LastName:     "Last",
		Birth:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	dto := xcondition.First(_mapper.Map(&Dto{}, po)).(*Dto)
	fmt.Println(dto)

	pos := [...]*Po{
		po, {
			DefaultField: 0.444,
			FirstName:    "fn",
			LastName:     "ln",
			Birth:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	dtos := xcondition.First(_mapper.Map([3]*Dto{}, pos)).([3]*Dto)
	fmt.Println(dtos[0], dtos[1])
}
