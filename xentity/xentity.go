package xentity

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"reflect"
)

// A whole mappers container
type EntityMappers struct {
	mappers []*EntityMapper
}

type MapFunc func(from interface{}, to interface{}) error

// An entity mapper
type EntityMapper struct {
	from    interface{}
	to      interface{}
	ctor    func() interface{}
	mapFunc MapFunc
}

func New() *EntityMappers {
	return &EntityMappers{mappers: []*EntityMapper{}}
}

func NewMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper {
	to := ctor()
	if reflect.TypeOf(from).Kind() != reflect.Ptr || reflect.TypeOf(to).Kind() != reflect.Ptr {
		panic("mapper type is not pointer")
	}
	return &EntityMapper{
		from:    from,
		to:      to,
		ctor:    ctor,
		mapFunc: mapFunc,
	}
}

func (e *EntityMappers) AddMapper(mapper *EntityMapper) {
	for _, m := range e.mappers {
		if m.from == mapper.from && m.to == mapper.to {
			m.mapFunc = mapper.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, mapper)
}

func (e *EntityMappers) AddMappers(mappers ...*EntityMapper) {
	for _, m := range mappers {
		e.AddMapper(m)
	}
}

func (e *EntityMappers) GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	for _, m := range e.mappers {
		if reflect.TypeOf(m.from) == reflect.TypeOf(from) && reflect.TypeOf(m.to) == reflect.TypeOf(to) {
			return m, nil
		}
	}
	return nil, fmt.Errorf("mapper type not found")
}

func (e *EntityMappers) _map(mapper *EntityMapper, from interface{}, to interface{}, options ...MapFunc) error {
	err := mapper.mapFunc(from, to)
	if err != nil {
		return err
	}
	for _, option := range options {
		err := option(from, to)
		if err != nil {
			return err
		}
	}
	return nil
}

// Example:
//     mapper.Map(&Po{}, &Dto{})
func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	mapper, err := e.GetMapper(from, to)
	if err != nil {
		return err
	}

	return e._map(mapper, from, to, options...)
}

// Example:
//     mapper.Map(&Po{}, &Dto{})
func (e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	mapper, err := e.GetMapper(from, toModel)
	if err != nil {
		return nil, err
	}

	to := mapper.ctor()
	err = e._map(mapper, from, to, options...)
	return to, err
}

// Example:
//     mapper.Map([]*Po{}, &Dto{})
func (e *EntityMappers) MapSlice(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	fromSlice := xslice.Sti(from)
	toType := reflect.SliceOf(reflect.TypeOf(toModel))

	to := reflect.MakeSlice(toType, len(fromSlice), len(fromSlice))
	for idx := range fromSlice {
		val, err := e.Map(fromSlice[idx], toModel, options...)
		if err != nil {
			return nil, err
		}
		to.Index(idx).Set(reflect.ValueOf(val))
	}

	return to.Interface(), nil
}

func (e *EntityMappers) MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	err := e.MapProp(from, to, options...)
	if err != nil {
		panic(err)

	}
}

func (e *EntityMappers) MustMap(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.Map(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

func (e *EntityMappers) MustMapSlice(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.MapSlice(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

var _mappers = New()

func AddMapper(mapper *EntityMapper) {
	_mappers.AddMapper(mapper)
}

func AddMappers(mappers ...*EntityMapper) {
	_mappers.AddMappers(mappers...)
}

func GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	return _mappers.GetMapper(from, to)
}

func MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	return _mappers.MapProp(from, to, options...)
}

func Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.Map(from, to, options...)
}

func MapSlice(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.MapSlice(from, to, options...)
}

func MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	_mappers.MustMapProp(from, to, options...)
}

func MustMap(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMap(from, to, options...)
}

func MustMapSlice(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMapSlice(from, to, options...)
}
