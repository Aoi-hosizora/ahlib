package xentity

import (
	"errors"
	"reflect"
)

// EntityMappers represents an entity mappers container, contains a list of EntityMapper.
type EntityMappers struct {
	mappers []*EntityMapper
}

// EntityMapper represents an entity mapper, contains the type information and map function.
type EntityMapper struct {
	// source represents mapping `source`.
	source interface{}

	// destination represents mapping `destination`.
	destination interface{}

	// srcType represents source type.
	srcType reflect.Type

	// destType represents destination type.
	destType reflect.Type

	// destCtor represents destination constructor.
	destCtor func() interface{}

	// mapFunc represents the mapping function from source to destination.
	mapFunc MapFunc
}

// MapFunc represents a mapping function, describes how to map from source to destination.
type MapFunc func(source interface{}, destination interface{}) error

// New creates an empty EntityMappers.
func New() *EntityMappers {
	return &EntityMappers{
		mappers: make([]*EntityMapper, 0),
	}
}

const (
	nilModelPanic          = "xentity: nil model"
	nilCtorPanic           = "xentity: nil constructor"
	nilMapFuncPanic        = "xentity: nil mapFunc"
	nonStructPtrModelPanic = "xentity: non-struct-pointer model"
	nonSliceModelPanic     = "xentity: non-slice model"
)

var (
	mapperNotFoundErr = errors.New("xentity: mapper not found")
)

// NewMapper creates an EntityMapper, panics when using nil model or invalid parameters.
func NewMapper(src interface{}, destCtor func() interface{}, mapFunc MapFunc) *EntityMapper {
	// check nil
	if src == nil {
		panic(nilModelPanic)
	}
	if destCtor == nil {
		panic(nilCtorPanic)
	}
	dest := destCtor()
	if dest == nil {
		panic(nilModelPanic)
	}
	if mapFunc == nil {
		panic(nilMapFuncPanic)
	}

	// check pointer
	srcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)
	if srcTyp.Kind() != reflect.Ptr || destTyp.Kind() != reflect.Ptr {
		panic(nonStructPtrModelPanic)
	}
	if srcTyp.Elem().Kind() != reflect.Struct || destTyp.Elem().Kind() != reflect.Struct {
		panic(nonStructPtrModelPanic)
	}

	// return
	return &EntityMapper{source: src, destination: dest, srcType: srcTyp, destType: destTyp, destCtor: destCtor, mapFunc: mapFunc}
}

// GetMapFunc returns the MapFunc from EntityMapper.
func (e *EntityMapper) GetMapFunc() MapFunc {
	return e.mapFunc
}

// AddMapper adds an EntityMapper to EntityMappers.
func (e *EntityMappers) AddMapper(mapper *EntityMapper) {
	for _, m := range e.mappers {
		if m.srcType == mapper.srcType && m.destType == mapper.destType {
			m.destCtor = mapper.destCtor
			m.mapFunc = mapper.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, mapper)
}

// AddMappers adds some EntityMapper-s to EntityMappers.
func (e *EntityMappers) AddMappers(mappers ...*EntityMapper) {
	for _, m := range mappers {
		e.AddMapper(m)
	}
}

// GetMapper returns the EntityMapper from EntityMappers, panics when using nil model, returns error when mapper not found.
func (e *EntityMappers) GetMapper(src interface{}, dest interface{}) (*EntityMapper, error) {
	if src == nil || dest == nil {
		panic(nilModelPanic)
	}

	srcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)
	for _, mapper := range e.mappers {
		if mapper.srcType == srcTyp && mapper.destType == destTyp {
			return mapper, nil
		}
	}
	return nil, mapperNotFoundErr
}

// ====
// core
// ====

// coreMap is the core implementation of EntityMapper, with EntityMapper.mapFunc and options.
func coreMap(mapper *EntityMapper, src interface{}, dest interface{}, opts ...MapFunc) error {
	err := mapper.mapFunc(src, dest)
	if err != nil {
		return err
	}

	for _, opt := range opts {
		if opt != nil {
			err = opt(src, dest)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// MapProp maps source properties into destination, panics when using nil model.
func (e *EntityMappers) MapProp(src interface{}, dest interface{}, opts ...MapFunc) error {
	mapper, err := e.GetMapper(src, dest)
	if err != nil {
		return err
	}

	return coreMap(mapper, src, dest, opts...)
}

// Map returns a destination instance mapped from source, panics when using nil model.
func (e *EntityMappers) Map(src interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error) {
	mapper, err := e.GetMapper(src, destModel)
	if err != nil {
		return nil, err
	}

	dest := mapper.destCtor()
	err = coreMap(mapper, src, dest, opts...)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

// MapSlice returns a destination slice mapped from source slice, panics when using nil model.
func (e *EntityMappers) MapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error) {
	if srcSlice == nil || destModel == nil {
		panic(nilModelPanic)
	}
	srcSliceVal := reflect.ValueOf(srcSlice)
	if srcSliceVal.Type().Kind() != reflect.Slice {
		panic(nonSliceModelPanic)
	}

	srcSliceLen := srcSliceVal.Len()
	srcItfSlice := make([]interface{}, srcSliceLen)
	for idx := 0; idx < srcSliceLen; idx++ {
		srcItfSlice[idx] = srcSliceVal.Index(idx).Interface()
	}
	destSliceTyp := reflect.SliceOf(reflect.TypeOf(destModel))
	destSliceVal := reflect.MakeSlice(destSliceTyp, len(srcItfSlice), len(srcItfSlice))
	if len(srcItfSlice) == 0 {
		return destSliceVal.Interface(), nil
	}

	mapper, err := e.GetMapper(srcItfSlice[0], destModel)
	if err != nil {
		return nil, err
	}

	for idx, src := range srcItfSlice {
		dest := mapper.destCtor()
		err = coreMap(mapper, src, dest, opts...)
		if err != nil {
			return nil, err
		}
		destSliceVal.Index(idx).Set(reflect.ValueOf(dest))
	}

	return destSliceVal.Interface(), nil
}

// MustMapProp is the must version of MapProp, panics when error.
func (e *EntityMappers) MustMapProp(src interface{}, dest interface{}, opts ...MapFunc) {
	err := e.MapProp(src, dest, opts...)
	if err != nil {
		panic(err)
	}
}

// MustMap is the must version of Map, panics when error.
func (e *EntityMappers) MustMap(src interface{}, destModel interface{}, opts ...MapFunc) interface{} {
	dest, err := e.Map(src, destModel, opts...)
	if err != nil {
		panic(err)
	}
	return dest
}

// MustMapSlice is the must version of MapSlice, panics when error.
func (e *EntityMappers) MustMapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) interface{} {
	destSlice, err := e.MapSlice(srcSlice, destModel, opts...)
	if err != nil {
		panic(err)
	}
	return destSlice
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

// GetMapper returns the EntityMapper from EntityMappers, panics when using nil model, returns error when mapper not found.
func GetMapper(src interface{}, dest interface{}) (*EntityMapper, error) {
	return _mappers.GetMapper(src, dest)
}

// MapProp maps source properties into destination, panics when using nil model.
func MapProp(src interface{}, dest interface{}, options ...MapFunc) error {
	return _mappers.MapProp(src, dest, options...)
}

// Map returns a destination instance mapped from source, panics when using nil model.
func Map(src interface{}, destModel interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.Map(src, destModel, options...)
}

// MapSlice returns a destination slice mapped from source slice, panics when using nil model.
func MapSlice(srcSlice interface{}, destModel interface{}, options ...MapFunc) (interface{}, error) {
	return _mappers.MapSlice(srcSlice, destModel, options...)
}

// MustMapProp is the must version of MapProp, panics when error.
func MustMapProp(src interface{}, dest interface{}, options ...MapFunc) {
	_mappers.MustMapProp(src, dest, options...)
}

// MustMap is the must version of Map, panics when error.
func MustMap(src interface{}, destModel interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMap(src, destModel, options...)
}

// MustMapSlice is the must version of MapSlice, panics when error.
func MustMapSlice(srcSlice interface{}, destModel interface{}, options ...MapFunc) interface{} {
	return _mappers.MustMapSlice(srcSlice, destModel, options...)
}
