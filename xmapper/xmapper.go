package xmapper

import (
	"errors"
	"reflect"
)

type MapFunc func(from interface{}, to interface{}) error

type EntityMapper struct {
	mappers []*Mapper
}

type Mapper struct {
	from    interface{}
	to      interface{}
	mapFunc MapFunc
}

func NewEntityMapper() *EntityMapper {
	return &EntityMapper{mappers: []*Mapper{}}
}

func NewMapper(from interface{}, to interface{}, mapFunc MapFunc) *Mapper {
	return &Mapper{from: from, to: to, mapFunc: mapFunc}
}

var (
	ErrNotPtr         = errors.New("model need to be pointer")
	ErrMapperNotFound = errors.New("mapper type not found")
)

func (e *EntityMapper) AddMapper(newMapper *Mapper) {
	for _, mapper := range e.mappers {
		if mapper.from == newMapper.from && mapper.to == newMapper.to {
			mapper.mapFunc = newMapper.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, newMapper)
}

func (e *EntityMapper) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	if reflect.TypeOf(from).Kind() != reflect.Ptr || reflect.TypeOf(to).Kind() != reflect.Ptr {
		return ErrNotPtr
	}
	var mapper *Mapper
	for _, m := range e.mappers {
		if reflect.TypeOf(m.from) == reflect.TypeOf(from) && reflect.TypeOf(m.to) == reflect.TypeOf(to) {
			mapper = m
		}
	}
	if mapper == nil {
		return ErrMapperNotFound
	}
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

func (e *EntityMapper) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	if reflect.TypeOf(toModel).Kind() != reflect.Ptr {
		return nil, ErrNotPtr
	}
	to := reflect.New(reflect.TypeOf(toModel).Elem()).Interface()
	err := e.MapProp(from, to, options...)
	if err != nil {
		return nil, err
	}
	return to, nil
}

func (e *EntityMapper) MapSlice(from []interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	to := reflect.New(reflect.SliceOf(reflect.TypeOf(toModel))).Elem()
	for idx := range from {
		val, err := e.Map(from[idx], toModel)
		if err != nil {
			return nil, err
		}
		to.Set(reflect.Append(to, reflect.ValueOf(val)))
	}
	return to.Interface(), nil
}
