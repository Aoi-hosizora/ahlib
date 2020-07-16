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
	from interface{}
	to   interface{}
	dict map[string]*PropertyMapperValue // dto -> po
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
	if dict == nil {
		dict = make(map[string]*PropertyMapperValue)
	}
	return &PropertyMapper{from: from, to: to, dict: dict}
}

func NewValue(revert bool, destProps ...string) *PropertyMapperValue {
	return &PropertyMapperValue{destProps: destProps, revert: revert}
}

func (p *PropertyMappers) AddMapper(mapper *PropertyMapper) {
	for _, m := range p.mappers {
		if reflect.TypeOf(m.from) == reflect.TypeOf(mapper.from) || reflect.TypeOf(m.to) == reflect.TypeOf(mapper.to) {
			m.dict = mapper.dict
			return
		}
	}
	p.mappers = append(p.mappers, mapper)
}

func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper) {
	for _, m := range mappers {
		p.AddMapper(m)
	}
}

func (p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	for _, m := range p.mappers {
		if reflect.TypeOf(m.from) == reflect.TypeOf(from) && reflect.TypeOf(m.to) == reflect.TypeOf(to) {
			return m, nil
		}
	}
	return nil, fmt.Errorf("mapper type not found")
}

func (p *PropertyMappers) GetMapperDefault(from interface{}, to interface{}) *PropertyMapper {
	m, err := p.GetMapper(from, to)
	if err != nil {
		return NewMapper(from, to, map[string]*PropertyMapperValue{})
	}
	return m
}

var _mappers = New()

// noinspection GoUnusedExportedFunction
func AddMapper(mapper *PropertyMapper) {
	_mappers.AddMapper(mapper)
}

// noinspection GoUnusedExportedFunction
func AddMappers(mappers ...*PropertyMapper) {
	_mappers.AddMappers(mappers...)
}

// noinspection GoUnusedExportedFunction
func GetMapper(from interface{}, to interface{}) (*PropertyMapper, error) {
	return _mappers.GetMapper(from, to)
}

// noinspection GoUnusedExportedFunction
func GetMapperDefault(from interface{}, to interface{}) *PropertyMapper {
	return _mappers.GetMapperDefault(from, to)
}
