package xproperty

import (
	"fmt"
	"reflect"
)

// A whole mappers container
type PropertyMappers struct {
	mappers []*PropertyMapper
}

// An entity mapper
type PropertyMapper struct {
	from     interface{}
	to       interface{}
	fromType reflect.Type
	toType   reflect.Type
	dict     map[string]*PropertyMapperValue // dto -> po
}

// A destination mapping property
// Example:
//     name -> first last
//     age -> birth
type PropertyMapperValue struct {
	destProps []string
	revert    bool
}

func New() *PropertyMappers {
	return &PropertyMappers{mappers: make([]*PropertyMapper, 0)}
}

func NewMapper(from interface{}, to interface{}, dict map[string]*PropertyMapperValue) *PropertyMapper {
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

func NewValue(revert bool, destProps ...string) *PropertyMapperValue {
	return &PropertyMapperValue{destProps: destProps, revert: revert}
}

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

func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

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

func (p *PropertyMappers) GetMapperDefault(from interface{}, to interface{}) *PropertyMapper {
	mapper, err := p.GetMapper(from, to)
	if err != nil {
		return NewMapper(from, to, map[string]*PropertyMapperValue{})
	}
	return mapper
}

var _mappers = New()

func AddMapper(mapper *PropertyMapper) {
	_mappers.AddMapper(mapper)
}

func AddMappers(mappers ...*PropertyMapper) {
	_mappers.AddMappers(mappers...)
}

func GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(from, to)
}

func GetMapperDefault(from interface{}, to interface{}) *PropertyMapper {
	return _mappers.GetMapperDefault(from, to)
}
