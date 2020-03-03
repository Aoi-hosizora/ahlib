package xentity

import (
	"errors"
	"reflect"
)

type EntityMappers struct {
	mappers []*EntityMapper
}

type MapFunc func(from interface{}, to interface{}) error

type EntityMapper struct {
	from    interface{}
	to      interface{}
	mapFunc MapFunc
}

func NewEntityMappers() *EntityMappers {
	return &EntityMappers{mappers: []*EntityMapper{}}
}

func NewEntityMapper(from interface{}, to interface{}, mapFunc MapFunc) *EntityMapper {
	if reflect.TypeOf(from).Kind() != reflect.Ptr || reflect.TypeOf(to).Kind() != reflect.Ptr {
		panic(ErrNotPtr)
	}
	return &EntityMapper{from: from, to: to, mapFunc: mapFunc}
}

var (
	ErrMapperNotFound = errors.New("mapper type not found")
	ErrNotPtr         = errors.New("mapper type is not pointer")
)

func (e *EntityMappers) AddMapper(newMapper *EntityMapper) {
	for _, mapper := range e.mappers {
		if mapper.from == newMapper.from && mapper.to == newMapper.to {
			mapper.mapFunc = newMapper.mapFunc
			return
		}
	}
	e.mappers = append(e.mappers, newMapper)
}

func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	var mapper *EntityMapper
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

func (e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	cnt := 0
	toType := reflect.TypeOf(toModel)
	for toType.Kind() == reflect.Ptr {
		cnt += 1
		toType = toType.Elem()
	}
	toValue := reflect.New(toType).Elem()
	for idx := 0; idx < cnt; idx++ {
		toTmp := reflect.New(toValue.Type())
		toTmp.Elem().Set(toValue)
		toValue = toTmp
	}

	err := e.MapProp(from, toValue.Interface(), options...)
	if err != nil {
		return nil, err
	}
	return toValue.Interface(), nil
}

func (e *EntityMappers) MapSlice(from []interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
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
