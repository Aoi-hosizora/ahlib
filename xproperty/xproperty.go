package xproperty

import (
	"reflect"
)

type PropertyMappers struct {
	mappings []*PropertyMapper
}

type PropertyMapper struct {
	from interface{}
	to   interface{}
	dict map[string]*PropertyMapperValue // dto -> po
}

type PropertyMapperValue struct {
	destProps []string // name -> first last
	revert    bool     // age -> birth
}

func NewPropertyMappers() *PropertyMappers {
	return &PropertyMappers{mappings: make([]*PropertyMapper, 0)}
}

func NewPropertyMapper(from interface{}, to interface{}, dict map[string]*PropertyMapperValue) *PropertyMapper {
	if dict == nil {
		dict = make(map[string]*PropertyMapperValue)
	}
	return &PropertyMapper{from: from, to: to, dict: dict}
}

func NewPropertyMapperValue(destProps []string, revert bool) *PropertyMapperValue {
	if destProps == nil {
		destProps = make([]string, 0)
	}
	return &PropertyMapperValue{destProps: destProps, revert: revert}
}

func (p *PropertyMappers) AddMapping(newMapping *PropertyMapper) {
	for _, mapping := range p.mappings {
		if reflect.TypeOf(mapping.from) == reflect.TypeOf(newMapping.from) || reflect.TypeOf(mapping.to) == reflect.TypeOf(newMapping.to) {
			mapping.dict = newMapping.dict
			return
		}
	}
	p.mappings = append(p.mappings, newMapping)
}

func (p *PropertyMappers) GetPropertyMapping(from interface{}, to interface{}) *PropertyMapper {
	for _, m := range p.mappings {
		if reflect.TypeOf(m.from) == reflect.TypeOf(from) && reflect.TypeOf(m.to) == reflect.TypeOf(to) {
			return m
		}
	}
	return NewPropertyMapper(from, to, map[string]*PropertyMapperValue{})
}
