package xmapper

import (
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Po struct {
	Flag      float32
	FirstName string
	LastName  string
	Birth     time.Time
	Nest      *NestPo
}

type NestPo struct {
	T int
}

type Dto struct {
	Flag float32
	Name string
	Age  int
	Nest *NestDto
}

type NestDto struct {
	T int
}

func TestEntitiesMapper_Map(t *testing.T) {
	_mapper := NewEntitiesMapper().
		CreateMapper(&NestPo{}, &NestDto{}).
		ForMember("T", func(t interface{}) interface{} {
			return t.(NestPo).T + 1
		}).
		Build().
		CreateMapper(&Po{}, &Dto{}).
		ForMember("Name", func(po interface{}) interface{} {
			return po.(Po).FirstName + "=" + po.(Po).LastName
		}).
		ForMember("Age", func(po interface{}) interface{} {
			return time.Now().Year() - po.(Po).Birth.Year()
		}).
		ForNest("Nest", "Nest").
		Build()

	t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

	po1 := &Po{Flag: 0.333, FirstName: "First", LastName: "Last", Birth: t1, Nest: &NestPo{T: 1}}
	po2 := Po{Flag: 0.444, FirstName: "Third", LastName: "Fourth", Birth: t2}

	dtoTest1 := &Dto{Flag: 0.333, Name: "First=Last", Age: 20, Nest: &NestDto{T: 2}}
	dtoTest2 := Dto{Flag: 0.444, Name: "Third=Fourth", Age: 10, Nest: nil}

	pos1 := []*Po{po1, &po2}
	pos2 := [...]*Po{po1, &po2}
	pos3 := []Po{*po1, po2}

	dtosTest1 := []*Dto{dtoTest1, &dtoTest2}
	dtosTest2 := [...]*Dto{dtoTest1, &dtoTest2}
	dtosTest3 := []Dto{*dtoTest1, dtoTest2}

	dto1 := xcondition.First(_mapper.Map(&Dto{}, po1)).(*Dto)
	assert.Equal(t, dto1, dtoTest1)

	dto2 := xcondition.First(_mapper.Map(Dto{}, po2)).(Dto)
	assert.Equal(t, dto2, dtoTest2)

	dto3 := xcondition.First(_mapper.Map([]*Dto{}, pos1)).([]*Dto)
	assert.Equal(t, dto3, dtosTest1)

	dto4 := xcondition.First(_mapper.Map([2]*Dto{}, pos2)).([2]*Dto)
	assert.Equal(t, dto4, dtosTest2)

	dto5 := xcondition.First(_mapper.Map([]Dto{}, pos3)).([]Dto)
	assert.Equal(t, dto5, dtosTest3)
}
