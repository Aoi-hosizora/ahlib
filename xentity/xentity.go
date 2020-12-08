package xentity

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"reflect"
)

// EntityMappers represents an entity mappers container.
type EntityMappers struct {
	mappers []*EntityMapper
}

// EntityMapper represents an entity mapper.
type EntityMapper struct {
	// from represents mapping `from`.
	from interface{}

	// to represents mapping `to`.
	to interface{}

	// fromType represents from type.
	fromType reflect.Type

	// toType represents to type.
	toType reflect.Type

	// ctor represents to constructor.
	ctor func() interface{}

	// mapFunc represents the mapping function.
	mapFunc MapFunc
}

// MapFunc represents a model mapping method, describe how to map model from `from` to `to`.
type MapFunc func(from interface{}, to interface{}) error

// New creates an EntityMappers.
func New() *EntityMappers {
	return &EntityMappers{
		mappers: make([]*EntityMapper, 0),
	}
}

var (
	nilModelPanic         = "xentity: nil model"
	nilCtorPanic          = "xentity: nil constructor"
	nilMapFuncPanic       = "xentity: nil mapFunc"
	nonStructPointerModel = "xentity: non-struct-pointer model"

	nilModelErr       = errors.New("xentity: nil model")
	mapperNotFoundErr = errors.New("xentity: mapper not found")
)

// NewMapper creates an EntityMapper, panic when invalid arguments.
func NewMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper {
	// check nil
	if from == nil {
		panic(nilModelPanic)
	}
	if ctor == nil {
		panic(nilCtorPanic)
	}
	if mapFunc == nil {
		panic(nilMapFuncPanic)
	}
	to := ctor()
	if to == nil {
		panic(nilModelPanic)
	}

	// check pointer
	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	if fromType.Kind() != reflect.Ptr || toType.Kind() != reflect.Ptr {
		panic(nonStructPointerModel)
	}
	if fromType.Elem().Kind() != reflect.Struct || toType.Elem().Kind() != reflect.Struct {
		panic(nonStructPointerModel)
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

// GetMapFunc returns the mapFunc from EntityMapper.
func (e *EntityMapper) GetMapFunc() MapFunc {
	return e.mapFunc
}

// AddMapper adds an EntityMapper to EntityMappers.
func (e *EntityMappers) AddMapper(m *EntityMapper) {
	for _, mapper := range e.mappers {
		if mapper.fromType == m.fromType && mapper.toType == m.toType {
			mapper.ctor = m.ctor
			mapper.mapFunc = m.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, m)
}

// AddMappers adds some EntityMapper-s to EntityMappers.
func (e *EntityMappers) AddMappers(ms ...*EntityMapper) {
	for _, m := range ms {
		e.AddMapper(m)
	}
}

// GetMapper returns the EntityMapper from EntityMappers, error when nil model or mapper not found.
func (e *EntityMappers) GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	if from == nil || to == nil {
		return nil, nilModelErr
	}

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	for _, mapper := range e.mappers {
		if mapper.fromType == fromType && mapper.toType == toType {
			return mapper, nil
		}
	}
	return nil, mapperNotFoundErr
}

// doMap is the core implementation of EntityMapper, with EntityMapper.mapFunc and options.
func doMap(mapper *EntityMapper, from interface{}, to interface{}, options ...MapFunc) error {
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

// MapProp maps properties from `from` to `to`.
func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	mapper, err := e.GetMapper(from, to)
	if err != nil {
		return err
	}

	return doMap(mapper, from, to, options...)
}

// Map generates a `to` from `from`.
func (e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	mapper, err := e.GetMapper(from, toModel)
	if err != nil {
		return nil, err
	}

	to := mapper.ctor()
	err = doMap(mapper, from, to, options...)
	if err != nil {
		return nil, err
	}
	return to, nil
}

// MapSlice generates a `to` slice from `from` slice.
func (e *EntityMappers) MapSlice(fromInterface interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	fromSlice := xslice.Sti(fromInterface)
	toSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(toModel)), len(fromSlice), len(fromSlice))
	if len(fromSlice) == 0 {
		return toSlice.Interface(), nil
	}

	mapper, err := e.GetMapper(fromSlice[0], toModel)
	if err != nil {
		return nil, err
	}

	for idx, from := range fromSlice {
		to := mapper.ctor()
		err = doMap(mapper, from, to, options...)
		if err != nil {
			return nil, err
		}
		toSlice.Index(idx).Set(reflect.ValueOf(to))
	}

	return toSlice.Interface(), nil
}

// MustMapProp is the must version of MapProp, panic if error.
func (e *EntityMappers) MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	err := e.MapProp(from, to, options...)
	if err != nil {
		panic(err)
	}
}

// Map is the must version of MapProp, panic if error.
func (e *EntityMappers) MustMap(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.Map(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

// MapSlice is the must version of MapProp, panic if error.
func (e *EntityMappers) MustMapSlice(from interface{}, toModel interface{}, options ...MapFunc) interface{} {
	i, err := e.MapSlice(from, toModel, options...)
	if err != nil {
		panic(err)
	}
	return i
}

// _mappers represents a global EntityMappers.
var _mappers = New()

// AddMapper adds an EntityMapper to EntityMappers.
func AddMapper(mapper *EntityMapper) {
	_mappers.AddMapper(mapper)
}

// AddMappers adds some EntityMapper-s to EntityMappers.
func AddMappers(mappers ...*EntityMapper) {
	_mappers.AddMappers(mappers...)
}

// GetMapper returns the EntityMapper from EntityMappers, error when nil model or mapper not found.
func GetMapper(from interface{}, to interface{}) (*EntityMapper, error) {
	return _mappers.GetMapper(from, to)
}

// MapProp maps properties from `from` to `to`.
func MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	return _mappers.MapProp(from, to, options...)
}

// Map generates a `to` from `from`.
func Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.Map(from, to, options...)
}

// MapSlice generates a `to` slice from `from` slice.
func MapSlice(from interface{}, to interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.MapSlice(from, to, options...)
}

// MustMapProp is the must version of MapProp, panic if error.
func MustMapProp(from interface{}, to interface{}, options ...MapFunc) {
	_mappers.MustMapProp(from, to, options...)
}

// Map is the must version of MapProp, panic if error.
func MustMap(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMap(from, to, options...)
}

// MapSlice is the must version of MapProp, panic if error.
func MustMapSlice(from interface{}, to interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMapSlice(from, to, options...)
}
