package xentity

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"reflect"
)

// A mappers container.
type EntityMappers struct {
	mappers []*EntityMapper
}

// MapFunc represents a model mapper's method, describe how to map model from `from` to `to`.
type MapFunc func(from interface{}, to interface{}) error

// An entity mapper.
type EntityMapper struct {
	// from instance.
	from interface{}

	// to instance.
	to interface{}

	// from type (struct's pointer).
	fromType reflect.Type

	// to type (struct's pointer).
	toType reflect.Type

	// to constructor.
	ctor func() interface{}

	// from -> to map function.
	mapFunc MapFunc
}

// Create a EntityMappers.
func New() *EntityMappers {
	return &EntityMappers{mappers: []*EntityMapper{}}
}

// Create a EntityMapper for EntityMappers.
func NewMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper {
	// check nil
	if from == nil || ctor == nil {
		panic("mapper's model could noe be nil")
	}
	to := ctor()
	if to == nil {
		panic("mapper's model could noe be nil")
	}

	// check pointer and struct
	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	if fromType.Kind() != reflect.Ptr || toType.Kind() != reflect.Ptr {
		panic("mapper's model could only be pointer")
	}

	fromTypeElem := fromType.Elem()
	toTypeElem := toType.Elem()
	if fromTypeElem.Kind() != reflect.Struct || toTypeElem.Kind() != reflect.Struct {
		panic("mapper's model could only be a pointer pointed to a struct")
	}

	// return
	return &EntityMapper{
		from:     from,
		to:       to,
		fromType: fromType,
		toType:   toType,
		ctor:     ctor,
		mapFunc:  mapFunc,
	}
}

// Add a EntityMapper to EntityMappers.
func (e *EntityMappers) AddMapper(m *EntityMapper) {
	for _, mapper := range e.mappers {
		if mapper.from == m.from && mapper.to == m.to {
			mapper.ctor = m.ctor
			mapper.mapFunc = m.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, m)
}

// Add some EntityMapper to EntityMappers.
func (e *EntityMappers) AddMappers(ms ...*EntityMapper) {
	for _, m := range ms {
		e.AddMapper(m)
	}
}

// Get a EntityMapper from EntityMappers.
func (e *EntityMappers) GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	for _, m := range e.mappers {
		if m.fromType == fromType && m.toType == toType {
			return m, nil
		}
	}
	return nil, fmt.Errorf("mapper is not found")
}

// Core implementation of EntityMappers.
func (e *EntityMappers) _map(mapper *EntityMapper, from interface{}, to interface{}, options ...MapFunc) error {
	err := mapper.mapFunc(from, to)
	if err != nil {
		return err
	}

	for _, option := range options {
		err = option(from, to)
		if err != nil {
			return err
		}
	}

	return nil
}

// Set `to` property from `from`.
func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	mapper, err := e.GetMapper(from, to)
	if err != nil {
		return err
	}

	return e._map(mapper, from, to, options...)
}

// Generate a `to` from `from`.
func (e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	mapper, err := e.GetMapper(from, toModel)
	if err != nil {
		return nil, err
	}

	to := mapper.ctor()
	err = e._map(mapper, from, to, options...)
	return to, err
}

// Generate a `to` slice from `from` slice. (No need to use xslice.Sti and xslice.Its)
func (e *EntityMappers) MapSlice(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	fromSlice := xslice.Sti(from)
	toType := reflect.SliceOf(reflect.TypeOf(toModel))
	toSlice := reflect.MakeSlice(toType, len(fromSlice), len(fromSlice))
	if len(fromSlice) == 0 {
		return toSlice, nil
	}

	mapper, err := e.GetMapper(fromSlice[0], toModel)
	if err != nil {
		return nil, err
	}

	for idx, item := range fromSlice {
		to := mapper.ctor()
		err = e._map(mapper, item, to, options...)
		if err != nil {
			return nil, err
		}

		val := reflect.ValueOf(to)
		toSlice.Index(idx).Set(val)
	}

	return toSlice.Interface(), nil
}

// Must version of EntityMappers.MapProp, panic if not found.
func (e *EntityMappers) MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	err := e.MapProp(from, to, options...)
	if err != nil {
		panic(err)
	}
}

// Must version of EntityMappers.Map, panic if not found.
func (e *EntityMappers) MustMap(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.Map(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

// Must version of EntityMappers.MapSlice, panic if not found.
func (e *EntityMappers) MustMapSlice(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.MapSlice(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

// Global EntityMappers.
var _mappers = New()

// Add a EntityMapper to global EntityMappers.
func AddMapper(mapper *EntityMapper) {
	_mappers.AddMapper(mapper)
}

// Add some EntityMapper to global EntityMappers.
func AddMappers(mappers ...*EntityMapper) {
	_mappers.AddMappers(mappers...)
}

// Get a EntityMapper from global EntityMappers.
func GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	return _mappers.GetMapper(from, to)
}

// Set `to` property from `from`.
func MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	return _mappers.MapProp(from, to, options...)
}

// Generate a `to` from `from`.
func Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.Map(from, to, options...)
}

// Generate a `to` slice from `from` slice.
func MapSlice(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.MapSlice(from, to, options...)
}

// Must version of MapProp, panic if not found.
func MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	_mappers.MustMapProp(from, to, options...)
}

// Must version of Map, panic if not found.
func MustMap(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMap(from, to, options...)
}

// Must version of MapSlice, panic if not found.
func MustMapSlice(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMapSlice(from, to, options...)
}
