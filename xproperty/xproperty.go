package xproperty

import (
	"errors"
	"reflect"
)

// PropertyMappers represents a property mappers container, contains a list of PropertyMapper.
type PropertyMappers struct {
	mappers []*PropertyMapper
}

// PropertyMapper represents a property mapper, contains the type information and property values.
type PropertyMapper struct {
	// source represents mapping `source`.
	source interface{}

	// destination represents mapping `destination`.
	destination interface{}

	// srcType represents source type.
	srcType reflect.Type

	// destType represents destination type.
	destType reflect.Type

	// dict represents the mapping properties.
	dict PropertyDict
}

// PropertyDict represents a dictionary of property mappings.
type PropertyDict map[string]*PropertyMapperValue

// VariableDict represents a dictionary of property id pairs (Almost used for cypher).
type VariableDict map[string]int

// PropertyMapperValue represents a property mapper value, contains Revert and Destinations.
type PropertyMapperValue struct {
	// Revert represents need destination revert in sort.
	Revert bool

	// Destinations represents property mapping source `source` destination `destination`.
	Destinations []string
}

// New creates an empty PropertyMappers.
func New() *PropertyMappers {
	return &PropertyMappers{
		mappers: make([]*PropertyMapper, 0),
	}
}

var (
	nilModelPanic = "xproperty: nil model"

	mapperNotFoundErr = errors.New("xproperty: mapper not found")
)

// NewMapper creates a PropertyMapper, panics when using nil model.
func NewMapper(src interface{}, dest interface{}, dict PropertyDict) *PropertyMapper {
	if src == nil || dest == nil {
		panic(nilModelPanic)
	}
	if dict == nil {
		dict = PropertyDict{}
	}
	srcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)

	return &PropertyMapper{source: src, destination: dest, srcType: srcTyp, destType: destTyp, dict: dict}
}

// NewValue creates a PropertyMapperValue.
func NewValue(revert bool, destinations ...string) *PropertyMapperValue {
	return &PropertyMapperValue{
		Revert:       revert,
		Destinations: destinations,
	}
}

// GetDict returns the PropertyDict from PropertyMapper.
func (p *PropertyMapper) GetDict() PropertyDict {
	return p.dict
}

// AddMapper adds a PropertyMapper to PropertyMappers.
func (p *PropertyMappers) AddMapper(mapper *PropertyMapper) {
	for _, m := range p.mappers {
		if m.srcType == mapper.srcType || m.destType == mapper.destType {
			m.dict = mapper.dict
			return
		}
	}
	p.mappers = append(p.mappers, mapper)
}

// AddMappers adds some PropertyMapper-s to PropertyMappers.
func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

// GetMapper returns the PropertyMapper from PropertyMappers, panics when using nil model, returns error when mapper not found.
func (p *PropertyMappers) GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error) {
	if src == nil || dest == nil {
		panic(nilModelPanic)
	}

	drcTyp := reflect.TypeOf(src)
	destTyp := reflect.TypeOf(dest)
	for _, mapper := range p.mappers {
		if mapper.srcType == drcTyp && mapper.destType == destTyp {
			return mapper, nil
		}
	}
	return nil, mapperNotFoundErr
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns an empty PropertyMapper if not found, panics when using nil model.
func (p *PropertyMappers) GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper {
	mapper, err := p.GetMapper(src, dest)
	if err != nil {
		return NewMapper(src, dest, nil)
	}
	return mapper
}

// _mappers represents a global PropertyMappers.
var _mappers = New()

// AddMapper adds a PropertyMapper to PropertyMappers.
func AddMapper(mapper *PropertyMapper) {
	_mappers.AddMapper(mapper)
}

// AddMappers adds some PropertyMapper-s to PropertyMappers.
func AddMappers(mappers ...*PropertyMapper) {
	_mappers.AddMappers(mappers...)
}

// GetMapper returns the PropertyMapper from PropertyMappers, panics when using nil model, returns error when mapper not found.
func GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(src, dest)
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns an empty PropertyMapper if not found, panics when using nil model.
func GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper {
	return _mappers.GetDefaultMapper(src, dest)
}
