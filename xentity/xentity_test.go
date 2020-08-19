package xentity

import (
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
	"time"
)

type Param struct {
	FirstName string
	LastName  string
}

type Po struct {
	Id        int
	FirstName string
	LastName  string
	Score     float32
	Info      *InfoPo
}

type InfoPo struct {
	Count    int
	Birthday time.Time
}

type Dto struct {
	Id    int
	Name  string
	Score int
	Info  *InfoDto
}

type InfoDto struct {
	CountAddOne int
	Age         int
}

func TestEntityMapper(t *testing.T) {
	entityMapper := New()
	entityMapper.AddMapper(NewMapper(&InfoPo{}, func() interface{} { return &InfoDto{} }, func(from interface{}, to interface{}) error {
		po := from.(*InfoPo)
		dto := to.(*InfoDto)
		dto.CountAddOne = po.Count + 1
		dto.Age = time.Now().Year() - po.Birthday.Year()
		return nil
	}))
	entityMapper.AddMapper(NewMapper(&Po{}, func() interface{} { return &Dto{} }, func(from interface{}, to interface{}) error {
		po := from.(*Po)
		dto := to.(*Dto)
		dto.Id = po.Id
		dto.Name = po.LastName + " " + po.FirstName
		dto.Score = int(po.Score)
		dto.Info = xcondition.First(entityMapper.Map(po.Info, &InfoDto{})).(*InfoDto)
		return nil
	}))
	entityMapper.AddMapper(NewMapper(&Param{}, func() interface{} { return &Po{} }, func(from interface{}, to interface{}) error {
		param := from.(*Param)
		po := to.(*Po)
		po.FirstName = param.FirstName
		po.LastName = param.LastName
		return nil
	}))

	param := &Param{
		FirstName: "First",
		LastName:  "Last",
	}
	po := &Po{
		Id: 1, Score: 9.8,
		Info: &InfoPo{Count: 20, Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)},
	}

	po1 := &Po{
		Id: 1, FirstName: "First", LastName: "Last", Score: 9.8,
		Info: &InfoPo{Count: 20, Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)},
	}
	dto11 := &Dto{
		Id: 1, Name: "Last First Last", Score: 10,
		Info: &InfoDto{CountAddOne: 21, Age: 20},
	}
	dto1 := &Dto{
		Id: 1, Name: "Last First", Score: 20,
		Info: &InfoDto{CountAddOne: 21, Age: 20},
	}

	po2 := &Po{
		Id: 2, FirstName: "First2", LastName: "Last2", Score: 0.1,
		Info: &InfoPo{Count: 1, Birthday: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local)},
	}
	dto2 := &Dto{
		Id: 2, Name: "Last2 First2", Score: 20,
		Info: &InfoDto{CountAddOne: 2, Age: 1},
	}

	poArr := []*Po{po1, po2}
	dtoArr := []*Dto{dto1, dto2}

	err := entityMapper.MapProp(param, po)
	log.Println(po, err)
	xtesting.Equal(t, po, po1)

	dtoOut, err := entityMapper.Map(po1, &Dto{}, func(from interface{}, to interface{}) error {
		po := from.(*Po)
		dto := to.(*Dto)
		dto.Score = int(po.Score + 0.2)
		dto.Name += " " + po.LastName
		return nil
	})
	log.Println(dtoOut, err)
	xtesting.Equal(t, dtoOut.(*Dto), dto11)

	dtoArrOut, err := entityMapper.MapSlice(xslice.Sti(poArr), &Dto{}, func(from interface{}, to interface{}) error {
		dto := to.(*Dto)
		dto.Score = 20
		return nil
	})
	log.Println(dtoArrOut, err)
	xtesting.Equal(t, dtoArrOut.([]*Dto), dtoArr)
}

type po struct {
	Int    int64
	String string
	Float  float64
	Array  []int
}

type dto struct {
	Int    int32
	String []byte
	Float  float32
	Array  []int
}

func (d *dto) Source() interface{} {
	return &po{}
}

func (d *dto) Ctor() interface{} {
	return &dto{}
}

func (d *dto) MapFrom(source interface{}) error {
	p := source.(*po)
	d.Int = int32(p.Int)
	d.String = []byte(p.String)
	d.Float = float32(p.Float)
	d.Array = p.Array
	return nil
}

func TestMappable(t *testing.T) {
	AddMapper(NewMapperByMappable(&dto{})) // X
	AddMappers(NewMapperByMappable(&dto{}))

	p := &po{
		Int:    5,
		String: "テスト",
		Float:  5.,
		Array:  []int{1, 2, 3},
	}
	d, err := Map(p, &dto{})
	log.Println(d, err)

	d = &dto{}
	err = MapProp(p, d)
	log.Println(d, err)

	ds, err := MapSlice([]*po{p, p}, &dto{})
	log.Println(ds, err)

	log.Println(GetMapper(&po{}, &dto{}))

	d1 := MustMap(p, &dto{}).(*dto)
	xtesting.Equal(t, d1.Int, int32(5))
	xtesting.Equal(t, d1.String, []byte("テスト"))
	xtesting.Equal(t, d1.Float, float32(5.))
	xtesting.Equal(t, d1.Array, []int{1, 2, 3})

	d1 = &dto{}
	MustMapProp(p, d1, func(from interface{}, to interface{}) error {
		to.(*dto).Int++
		return nil
	})
	xtesting.Equal(t, d1.Int, int32(6))
	xtesting.Equal(t, d1.String, []byte("テスト"))
	xtesting.Equal(t, d1.Float, float32(5.))
	xtesting.Equal(t, d1.Array, []int{1, 2, 3})

	dd := MustMapSlice([]*po{p, p}, &dto{}).([]*dto)
	xtesting.Equal(t, len(dd), 2)
	xtesting.Equal(t, dd[0].Int, int32(5))
	xtesting.Equal(t, dd[0].String, []byte("テスト"))
	xtesting.Equal(t, dd[0].Float, float32(5.))
	xtesting.Equal(t, dd[0].Array, []int{1, 2, 3})
}
