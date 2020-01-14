package xmapper

import (
	"errors"
	"reflect"
)

// Save all mapper between entities
type EntitiesMapper struct {
	_entities []*Entity
}

// Save all map rule between specific _from and _to entity type
type Entity struct {
	_from    reflect.Type
	_to      reflect.Type
	_mapper  *EntitiesMapper
	_mapRule []*_mapRule
}

// Save a map rule between specific structField
type _mapRule struct {
	_to      reflect.StructField
	_mapFunc MapFunc
}

type MapFunc func(interface{}) interface{}

var (
	CreateMapperFromInvalidError = errors.New("could not create mapper by non-ptr and non-struct model")
	MapToModelNilError           = errors.New("could not map to a nil model")
	MapDifferentKindError        = errors.New("could not map to a different kind of type")
	MapSmallSieArrayError        = errors.New("could not map to a small size of array")
	MapToNotSupportKindError     = errors.New("could not map a non-ptr/non-slice/non-array/non-struct model")
)

// Create a entity from entitiesMapper, not add into entitiesMapper yet
func CreateMapper(fromModel interface{}, toModel interface{}) *Entity {
	m := new(EntitiesMapper)
	return m.CreateMapper(fromModel, toModel)
}

// CreateMapper of *EntitiesMapper
// panic when fromModel or toModel is non-ptr and non-struct
func (e *EntitiesMapper) CreateMapper(fromModel interface{}, toModel interface{}) *Entity {
	checkEl := func(model interface{}) reflect.Type {
		t := reflect.TypeOf(model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() != reflect.Struct {
			panic(CreateMapperFromInvalidError)
		}
		return t
	}

	return &Entity{
		_from:   checkEl(fromModel),
		_to:     checkEl(toModel),
		_mapper: e,
	}
}

// Add entity mapper rule to entitiesMapper
func (e *Entity) Build() *EntitiesMapper {
	e._mapper._entities = append(e._mapper._entities, e)
	return e._mapper
}

// Add user defined field mapper rule, ignore if this field not exist
func (e *Entity) ForMember(toFieldString string, mapFunc MapFunc) *Entity {
	toField, ok := e._to.FieldByName(toFieldString)
	if !ok {
		return e
	}
	rule := &_mapRule{
		_to:      toField,
		_mapFunc: mapFunc,
	}
	e._mapRule = append(e._mapRule, rule)
	return e
}

func (e *EntitiesMapper) Map(toModel interface{}, from interface{}) (interface{}, error) {
	if toModel == nil {
		return nil, MapToModelNilError
	}
	if from == nil {
		return nil, nil
	}

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(toModel)
	kind := fromType.Kind()
	if kind != toType.Kind() {
		return nil, MapDifferentKindError
	}

	switch kind {
	case reflect.Ptr:
		// toModelElem := reflect.ValueOf(toModel).Elem()
		toModelElem := reflect.New(reflect.TypeOf(toModel).Elem()).Elem()
		fromElem := reflect.ValueOf(from).Elem()
		toElem, err := e.Map(toModelElem.Interface(), fromElem.Interface())
		if err != nil {
			return nil, err
		}
		toValue := reflect.New(reflect.TypeOf(toModel).Elem())
		toValue.Elem().Set(reflect.ValueOf(toElem))
		return toValue.Interface(), nil
	case reflect.Slice:
		fromValue := reflect.ValueOf(from)
		fromLen := fromValue.Len()
		toValue := reflect.MakeSlice(toType, fromLen, fromLen)
		if fromLen == 0 {
			return toValue.Interface(), nil
		}
		// call of reflect.Value.Elem on slice Value
		// toModelElem := reflect.ValueOf(toModel).Elem()
		toModelElem := reflect.New(reflect.TypeOf(toModel).Elem()).Elem()
		for idx := 0; idx < fromLen; idx++ {
			toElem, err := e.Map(toModelElem.Interface(), fromValue.Index(idx).Interface())
			if err != nil {
				return nil, err
			}
			toValue.Index(idx).Set(reflect.ValueOf(toElem))
		}
		return toValue.Interface(), nil
	case reflect.Array:
		fromValue := reflect.ValueOf(from)
		fromLen := fromValue.Len()
		toValue := reflect.New(toType).Elem()
		if fromLen > toValue.Len() { // check length
			return nil, MapSmallSieArrayError
		}
		if fromLen == 0 {
			return toValue.Interface(), nil
		}
		toModelElem := reflect.New(reflect.TypeOf(toModel).Elem()).Elem()
		for idx := 0; idx < fromLen; idx++ {
			toElem, err := e.Map(toModelElem.Interface(), fromValue.Index(idx).Interface())
			if err != nil {
				return nil, err
			}
			toValue.Index(idx).Set(reflect.ValueOf(toElem))
		}
		return toValue.Interface(), nil
	case reflect.Struct:
	default:
		return nil, MapToNotSupportKindError
	}

	// !!!!!!
	// copy same field
	fromValue := reflect.ValueOf(from)
	toValue := reflect.New(reflect.TypeOf(toModel)).Elem()
	for idx := 0; idx < toType.NumField(); idx++ {
		toField := toType.Field(idx)
		// same name and same type
		if fromField, ok := fromType.FieldByName(toField.Name); ok && fromField.Type == toField.Type {
			toValue.FieldByIndex(toField.Index).Set(fromValue.FieldByIndex(fromField.Index))
		}
	}

	// find the first map rule
	var mapEntity *Entity
	for _, entity := range e._entities {
		if entity._from == fromType && entity._to == toType {
			mapEntity = entity
			break
		}
	}
	if mapEntity == nil {
		return toValue.Interface(), nil
	}

	// map through rule
	for _, rule := range mapEntity._mapRule {
		toValue.FieldByIndex(rule._to.Index).Set(reflect.ValueOf(rule._mapFunc(from)))
	}
	return toValue.Interface(), nil
}
