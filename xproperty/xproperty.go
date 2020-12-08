package xproperty

import (
	"errors"
	"reflect"
)

// PropertyMappers represents a property mappers container.
type PropertyMappers struct {
	mappers []*PropertyMapper
}

// PropertyMapper represents a property mapper.
type PropertyMapper struct {
	// from represents mapping `from`.
	from interface{}

	// to represents mapping `to`.
	to interface{}

	// fromType represents from type.
	fromType reflect.Type

	// toType represents to type.
	toType reflect.Type

	// dict represents the mapping properties.
	dict PropertyDict
}

// PropertyMapperValue represents a property mapper value.
type PropertyMapperValue struct {
	// Revert represents need to revert sort.
	Revert bool

	// Destinations represents property mapping from `from` to `to`.
	Destinations []string
}

// PropertyDict represents a dictionary of property mappings.
type PropertyDict map[string]*PropertyMapperValue

// VariableDict represents a dictionary of property id pairs (Almost used for cypher).
type VariableDict map[string]int

// New creates a PropertyMappers.
func New() *PropertyMappers {
	return &PropertyMappers{mappers: make([]*PropertyMapper, 0)}
}

var (
	nilModelPanic = "xproperty: nil model"

	nilModelErr       = errors.New("xproperty: nil model")
	mapperNotFoundErr = errors.New("xproperty: mapper not found")
)

// NewMapper creates a PropertyMapper.
func NewMapper(from interface{}, to interface{}, dict PropertyDict) *PropertyMapper {
	if from == nil || to == nil {
		panic(nilModelPanic)
	}
	if dict == nil {
		dict = make(PropertyDict)
	}

	return &PropertyMapper{
		from:     from,
		to:       to,
		fromType: reflect.TypeOf(from),
		toType:   reflect.TypeOf(to),
		dict:     dict,
	}
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
func (p *PropertyMappers) AddMapper(m *PropertyMapper) {
	for _, mapper := range p.mappers {
		if mapper.fromType == m.fromType || mapper.toType == m.toType {
			mapper.dict = m.dict
			return
		}
	}
	p.mappers = append(p.mappers, m)
}

// AddMappers adds some PropertyMapper-s to PropertyMappers.
func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

// GetMapper returns the PropertyMapper from PropertyMappers.
func (p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	if from == nil || to == nil {
		return nil, nilModelErr
	}

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)
	for _, mapper := range p.mappers {
		if mapper.fromType == fromType && mapper.toType == toType {
			return mapper, nil
		}
	}
	return nil, mapperNotFoundErr
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns a empty PropertyMapper if not found.
func (p *PropertyMappers) GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper {
	mapper, err := p.GetMapper(from, to)
	if err != nil {
		return NewMapper(from, to, nil)
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

// GetMapper returns the PropertyMapper from PropertyMappers.
func GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(from, to)
}

// GetDefaultMapper returns the PropertyMapper from PropertyMappers, returns a empty PropertyMapper if not found.
func GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper {
	return _mappers.GetDefaultMapper(from, to)
}
