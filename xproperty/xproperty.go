package xproperty

import (
	"fmt"
	"reflect"
)

// A mappers container.
type PropertyMappers struct {
	mappers []*PropertyMapper
}

// A set of properties mapper.
type PropertyMapper struct {
	// `from` instance.
	from interface{}

	// `to` instance.
	to interface{}

	// `from` type (struct's pointer).
	fromType reflect.Type

	// `to` type (struct's pointer).
	toType reflect.Type

	// `from` -> `to` properties mapping.
	dict PropertyDict
}

// PropertyDict represents a dictionary of property mapping.
type PropertyDict map[string]*PropertyMapperValue

// A property mapper.
type PropertyMapperValue struct {
	// Is need to revert sort.
	Revert bool

	// `from` -> `to` properties mapping.
	Destinations []string
}

// Create a PropertyMappers.
func New() *PropertyMappers {
	return &PropertyMappers{mappers: make([]*PropertyMapper, 0)}
}

// Create a PropertyMapper.
func NewMapper(from interface{}, to interface{}, dict PropertyDict) *PropertyMapper {
	if from == nil || to == nil {
		panic("Mapper's model type should not be nil")
	}
	if dict == nil {
		dict = make(map[string]*PropertyMapperValue)
	}

	return &PropertyMapper{
		from:     from,
		to:       to,
		fromType: reflect.TypeOf(from),
		toType:   reflect.TypeOf(to),
		dict:     dict,
	}
}

// Create a PropertyMapperValue.
func NewValue(revert bool, destinations ...string) *PropertyMapperValue {
	return &PropertyMapperValue{
		Revert:       revert,
		Destinations: destinations,
	}
}

// Add a PropertyMapper to PropertyMappers.
func (p *PropertyMappers) AddMapper(m *PropertyMapper) {
	fromType := reflect.TypeOf(m.from)
	toType := reflect.TypeOf(m.to)
	for _, mapper := range p.mappers {
		if mapper.fromType == fromType || mapper.toType == toType {
			mapper.dict = m.dict
			return
		}
	}
	p.mappers = append(p.mappers, m)
}

// Add some PropertyMapper to PropertyMappers.
func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

// Get a PropertyMapper from PropertyMappers.
func (p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	for _, mapper := range p.mappers {
		if mapper.fromType == fromType && mapper.toType == toType {
			return mapper, nil
		}
	}
	return nil, fmt.Errorf("mapper type not found")
}

// Get a default PropertyMapper from PropertyMappers.
func (p *PropertyMappers) GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper {
	mapper, err := p.GetMapper(from, to)
	if err != nil {
		return NewMapper(from, to, map[string]*PropertyMapperValue{})
	}
	return mapper
}

// Global PropertyMappers.
var _mappers = New()

// Add a PropertyMapper to global PropertyMappers.
func AddMapper(mapper *PropertyMapper) {
	_mappers.AddMapper(mapper)
}

// Add some PropertyMapper to global PropertyMappers.
func AddMappers(mappers ...*PropertyMapper) {
	_mappers.AddMappers(mappers...)
}

// Get a PropertyMapper from global PropertyMappers.
func GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(from, to)
}

// Get a default PropertyMapper from global PropertyMappers.
func GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper {
	return _mappers.GetDefaultMapper(from, to)
}
