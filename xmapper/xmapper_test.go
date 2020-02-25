package xmapper

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
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
	T      int
	Extra1 string
}

type Dto struct {
	Flag float32
	Name string
	Age  int
	Nest *NestDto
}

type NestDto struct {
	T      int
	Extra2 string
}

func TestEntitiesMapper_Map(t *testing.T) {
	_mapper := NewEntityMapper().
		CreateMapper(&NestPo{}, &NestDto{}).
		ForMember("T", func(t interface{}) interface{} { // map between member
			return t.(NestPo).T + 1
		}).
		ForCopy("Extra1", "Extra2"). // map from copy
		Build().
		CreateMapper(&Po{}, &Dto{}).
		ForMember("Name", func(po interface{}) interface{} {
			return po.(Po).FirstName + "=" + po.(Po).LastName
		}).
		ForMember("Age", func(po interface{}) interface{} {
			return time.Now().Year() - po.(Po).Birth.Year()
		}).
		ForNest("Nest", "Nest"). // map nest
		Build().
		CreateMapper(&NestPo{}, &NestDto{}).
		ForMember("Extra2", func(i interface{}) interface{} {
			return i.(NestPo).Extra1 + "000"
		}).
		ForExtra(func(i interface{}, j interface{}) interface{} { // map extra function
			poFrom := i.(NestPo)
			dtoTo := j.(NestDto)
			dtoTo.Extra2 = "before " + poFrom.Extra1 + " after " + dtoTo.Extra2
			return dtoTo
		}).
		Build()

	t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

	po1 := &Po{Flag: 0.333, FirstName: "First", LastName: "Last", Birth: t1, Nest: &NestPo{T: 1, Extra1: "extra"}}
	po2 := Po{Flag: 0.444, FirstName: "Third", LastName: "Fourth", Birth: t2}

	dtoTest1 := &Dto{Flag: 0.333, Name: "First=Last", Age: 20, Nest: &NestDto{T: 2, Extra2: "before extra after extra000"}}
	dtoTest2 := Dto{Flag: 0.444, Name: "Third=Fourth", Age: 10, Nest: nil}
	dtoTest3 := &Dto{Flag: 0.333, Name: "First=Last_Extra_Disposable", Age: 20, Nest: &NestDto{T: 2, Extra2: "before extra after extra000"}}

	pos1 := []*Po{po1, &po2}
	pos2 := [...]*Po{po1, &po2}
	pos3 := []Po{*po1, po2}

	dtosTest1 := []*Dto{dtoTest1, &dtoTest2}
	dtosTest2 := [...]*Dto{dtoTest1, &dtoTest2}
	dtosTest3 := []Dto{*dtoTest1, dtoTest2}

	// directly map pointer
	dto1 := xcondition.First(_mapper.Map(&Dto{}, po1)).(*Dto)
	fmt.Println(xstring.MarshalJson(dto1))
	assert.Equal(t, dto1, dtoTest1)

	// directly map struct
	dto2 := xcondition.First(_mapper.Map(Dto{}, po2)).(Dto)
	fmt.Println(xstring.MarshalJson(dto2))
	assert.Equal(t, dto2, dtoTest2)

	// map pointer with disposable option
	dto6 := xcondition.First(_mapper.Map(&Dto{}, po1, NewMapOption(&Po{}, &Dto{}, func(i interface{}, j interface{}) interface{} {
		to := j.(Dto)
		to.Name += "_Extra"
		return to
	}), NewMapOption(&Po{}, &Dto{}, func(i interface{}, j interface{}) interface{} {
		to := j.(Dto)
		to.Name += "_Disposable"
		return to
	}))).(*Dto)
	fmt.Println(xstring.MarshalJson(dto6))
	assert.Equal(t, dto6, dtoTest3)

	// map slice of pointer
	dto3 := xcondition.First(_mapper.Map([]*Dto{}, pos1)).([]*Dto)
	assert.Equal(t, dto3, dtosTest1)

	// map array
	dto4 := xcondition.First(_mapper.Map([2]*Dto{}, pos2)).([2]*Dto)
	assert.Equal(t, dto4, dtosTest2)

	// map slice of struct
	dto5 := xcondition.First(_mapper.Map([]Dto{}, pos3)).([]Dto)
	assert.Equal(t, dto5, dtosTest3)
}

type NestParam struct {
	A int
}

type NestParamPo struct {
	A int
}

type Param struct {
	Name string
	Id   int
	T    string
	Nest NestParam
}

type ParamPo struct {
	FirstName string
	LastName  string
	Id        int
	Other     float32
	T2        string
	Nest      NestParamPo
}

func TestEntityMapper_MapProp(t *testing.T) {
	_mapper := NewEntityMapper().
		CreateMapper(&Param{}, &ParamPo{}).
		ForMember("Id", func(param interface{}) interface{} {
			return param.(Param).Id + 1
		}).
		ForExtra(func(param interface{}, po interface{}) interface{} {
			from := param.(Param)
			to := po.(ParamPo)
			sp := strings.Split(from.Name, " ")
			if len(sp) >= 2 {
				to.FirstName = sp[0]
				to.LastName = sp[1]
			}
			return to
		}).
		ForCopy("T", "T2").
		ForNest("Nest", "Nest").
		Build().
		CreateMapper(&NestParam{}, &NestParamPo{}).
		ForMember("A", func(from interface{}) interface{} {
			return from.(NestParam).A + 1
		}).
		Build()

	param := &Param{Name: "Ab Cd", Id: 10, T: "123", Nest: NestParam{A: 2}}
	po := &ParamPo{Other: 0.5}
	poTest := &ParamPo{FirstName: "Ab", LastName: "Cd", Id: 11, Other: 0.5, T2: "123", Nest: NestParamPo{A: 3}}

	err := _mapper.MapProp(param, po)
	log.Println(po, err)

	assert.Equal(t, po, poTest)
}
