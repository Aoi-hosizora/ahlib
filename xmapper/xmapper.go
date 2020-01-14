package xmapper

import (
	"errors"
	"reflect"
)

// Save all mapper between entities
type EntitiesMapper struct {
	_entities []*EntityMapper
}

// Save all map rule between specific _fromType and _toType entity type
type EntityMapper struct {
	_mapper *EntitiesMapper

	_fromType reflect.Type
	_toType   reflect.Type

	_directRule []*_fieldDirectMapRule
	_nestRule   []*_fieldNestMapRule
}

// Save the direct map rule between specific structField
type _fieldDirectMapRule struct {
	_toField reflect.StructField
	_mapFunc MapFunc
}

// Save the nest map rule between specific structField
type _fieldNestMapRule struct {
	_fromField reflect.StructField
	_toField   reflect.StructField
}

// Map Function from interface{} (is fromModel type) to interface{} (is toModel field type)
type MapFunc func(interface{}) interface{}

var (
	_createMapperFromInvalidPanic = errors.New("createMapper: could not create mapper by non-ptr and non-struct model")

	MapToModelNilError       = errors.New("could not map to a nil model")
	MapDifferentKindError    = errors.New("could not map to a different kind of type")
	MapSmallSieArrayError    = errors.New("could not map to a small size of array")
	MapToNotSupportKindError = errors.New("could not map a non-ptr/non-slice/non-array/non-struct model")
)

// Create a entity from entitiesMapper
func NewEntitiesMapper() *EntitiesMapper {
	return new(EntitiesMapper)
}

// CreateMapper of *EntitiesMapper
// panic when fromModel or toModel is non-ptr and non-struct
func (e *EntitiesMapper) CreateMapper(fromModel interface{}, toModel interface{}) *EntityMapper {
	checkEl := func(model interface{}) reflect.Type {
		t := reflect.TypeOf(model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() != reflect.Struct {
			panic(_createMapperFromInvalidPanic)
		}
		return t
	}

	return &EntityMapper{
		_fromType: checkEl(fromModel),
		_toType:   checkEl(toModel),
		_mapper:   e,
	}
}

// Add entity mapper rule to entitiesMapper
func (e *EntityMapper) Build() *EntitiesMapper {
	e._mapper._entities = append(e._mapper._entities, e)
	return e._mapper
}

// Add direct field mapper rule, ignore if toField is not exist
func (e *EntityMapper) ForMember(toFieldString string, mapFunc MapFunc) *EntityMapper {
	toField, ok := e._toType.FieldByName(toFieldString)
	if !ok {
		return e
	}
	rule := &_fieldDirectMapRule{
		_toField: toField,
		_mapFunc: mapFunc,
	}
	e._directRule = append(e._directRule, rule)
	return e
}

// Add nest field mapper rule, ignore if fromField or toField is not exist
func (e *EntityMapper) ForNest(fromFieldString string, toFieldString string) *EntityMapper {
	fromField, ok := e._fromType.FieldByName(fromFieldString)
	if !ok {
		return e
	}
	toField, ok := e._toType.FieldByName(toFieldString)
	if !ok {
		return e
	}
	rule := &_fieldNestMapRule{
		_fromField: fromField,
		_toField:   toField,
	}
	e._nestRule = append(e._nestRule, rule)
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
		// get elem of the from(ptr) -> map elem as toElem -> put toElem to to(ptr)
		toModelElem := reflect.New(reflect.TypeOf(toModel).Elem()).Elem()
		fromValue := reflect.ValueOf(from)
		if fromValue.IsNil() {
			return nil, nil
		}
		// from field not nil
		fromElem := fromValue.Elem()
		toElem, err := e.Map(toModelElem.Interface(), fromElem.Interface())
		if err != nil {
			return nil, err
		}
		toValue := reflect.New(reflect.TypeOf(toModel).Elem())
		toValue.Elem().Set(reflect.ValueOf(toElem))
		return toValue.Interface(), nil
	case reflect.Slice:
		// get elem of the from(slice) -> map all elem as toElem -> put toElem to to(slice)
		fromValue := reflect.ValueOf(from)
		fromLen := fromValue.Len()
		toValue := reflect.MakeSlice(toType, fromLen, fromLen)
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
	case reflect.Array:
		// check fromArr and toArr size -> get elem of the from(array) -> map all elem as toElem -> put toElem to to(array)
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
		// after switch
	default:
		return nil, MapToNotSupportKindError
	}

	// copy same name and same type of field in struct
	fromValue := reflect.ValueOf(from)
	toValue := reflect.New(reflect.TypeOf(toModel)).Elem()
	for idx := 0; idx < toType.NumField(); idx++ {
		toField := toType.Field(idx)
		// same name and same type
		if fromField, ok := fromType.FieldByName(toField.Name); ok && fromField.Type == toField.Type {
			toValue.FieldByIndex(toField.Index).Set(fromValue.FieldByIndex(fromField.Index))
		}
	}

	// find the first map rule of (fromType, toType) tuple
	var mapEntity *EntityMapper
	for _, entity := range e._entities {
		if entity._fromType == fromType && entity._toType == toType {
			mapEntity = entity
			break
		}
	}
	if mapEntity == nil {
		return toValue.Interface(), nil
	}

	// map through the found rule

	// direct, through _mapFunc
	for _, rule := range mapEntity._directRule {
		toValue.FieldByIndex(rule._toField.Index).Set(reflect.ValueOf(rule._mapFunc(from)))
	}

	// nest, through e.Map
	for _, rule := range mapEntity._nestRule {
		toFieldValue := toValue.FieldByIndex(rule._toField.Index)
		fromFieldValue := fromValue.FieldByIndex(rule._fromField.Index)
		// from field not nil
		if fromFieldValue.Kind() != reflect.Ptr || !fromFieldValue.IsNil() {
			toFieldNew, err := e.Map(toFieldValue.Interface(), fromFieldValue.Interface())
			if err != nil {
				return nil, err
			}
			toFieldValue.Set(reflect.ValueOf(toFieldNew))
		}
	}

	return toValue.Interface(), nil
}
