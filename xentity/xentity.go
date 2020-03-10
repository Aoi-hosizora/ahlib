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
	ctor    func() interface{}
	mapFunc MapFunc
}

func NewEntityMappers() *EntityMappers {
	return &EntityMappers{mappers: []*EntityMapper{}}
}

func NewEntityMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper {
	to := ctor()
	if reflect.TypeOf(from).Kind() != reflect.Ptr || reflect.TypeOf(to).Kind() != reflect.Ptr {
		panic(ErrNotPtr)
	}
	return &EntityMapper{
		from:    from,
		to:      to,
		ctor:    ctor,
		mapFunc: mapFunc,
	}
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

func (e *EntityMappers) _find(from interface{}, to interface{}) (*EntityMapper, error) {
	var mapper *EntityMapper
	for _, m := range e.mappers {
		if reflect.TypeOf(m.from) == reflect.TypeOf(from) && reflect.TypeOf(m.to) == reflect.TypeOf(to) {
			mapper = m
		}
	}
	if mapper == nil {
		return nil, ErrMapperNotFound
	}
	return mapper, nil
}

func (e *EntityMappers) _map(mapper *EntityMapper, from interface{}, to interface{}, options ...MapFunc) error {
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

func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error {
	mapper, err := e._find(from, to)
	if err != nil {
		return err
	}
	return e._map(mapper, from, to, options...)
}

// Example:
//     mapper.Map(&Po{}, &Dto{})
func (e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error) {
	mapper, err := e._find(from, toModel)
	if err != nil {
		return nil, err
	}
	to := mapper.ctor()
	err = e._map(mapper, from, to, options...)
	return to, err
}

// Example:
//     mapper.Map([]*Po{}, &Dto{})
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
