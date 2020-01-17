package xmapper

import (
	"errors"
	"reflect"
)

var (
	createMapperFromInvalidPanic = errors.New("createMapper: could not create mapper by non-ptr and non-struct model")

	mapToModelNilError          = errors.New("could not map to a nil model")
	mapDifferentKindError       = errors.New("could not map to a different kind of type")
	mapSmallSieArrayError       = errors.New("could not map to a small size of array")
	mapToNotSupportKindError    = errors.New("could not map a non-ptr/non-slice/non-array/non-struct model")
	mapExtraFunctionReturnError = errors.New("could not map extra function to a different type")
)

// check element and struct type
func _checkEl(model interface{}) reflect.Type {
	t := reflect.TypeOf(model)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(createMapperFromInvalidPanic)
	}
	return t
}

// Create or modify mapper of *EntityMapper
// panic when fromModel or toModel is non-ptr and non-struct
func (e *EntityMapper) CreateMapper(fromModel interface{}, toModel interface{}) *entity {
	_fromType := _checkEl(fromModel)
	_toType := _checkEl(toModel)

	// Use exist entity
	for _, entity := range e._entities {
		if entity._fromType == _fromType && entity._toType == _toType {
			return entity
		}
	}

	return &entity{
		_fromType: _fromType,
		_toType:   _toType,
		_mapper:   e,
	}
}

// Add entity mapper rule to entitiesMapper
func (e *entity) Build() *EntityMapper {
	e._mapper._entities = append(e._mapper._entities, e)
	return e._mapper
}

// Create DisposableMapOption for Map
func NewMapOption(fromModel interface{}, toModel interface{}, extraMapFunc ExtraMapFunc) *DisposableMapOption {
	_fromType := _checkEl(fromModel)
	_toType := _checkEl(toModel)

	return &DisposableMapOption{
		_fromType: _fromType,
		_toType:   _toType,
		_mapFunc:  extraMapFunc,
	}
}

// options: for disposable map options
func (e *EntityMapper) Map(toModel interface{}, from interface{}, options ...*DisposableMapOption) (interface{}, error) {
	if toModel == nil {
		return nil, mapToModelNilError
	}
	if from == nil {
		return nil, nil
	}

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(toModel)
	kind := fromType.Kind()
	if kind != toType.Kind() {
		return nil, mapDifferentKindError
	}

	// handle kind of type,generate all struct type
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
		toElem, err := e.Map(toModelElem.Interface(), fromElem.Interface(), options...)
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
			toElem, err := e.Map(toModelElem.Interface(), fromValue.Index(idx).Interface(), options...)
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
			return nil, mapSmallSieArrayError
		}
		if fromLen == 0 {
			return toValue.Interface(), nil
		}
		toModelElem := reflect.New(reflect.TypeOf(toModel).Elem()).Elem()
		for idx := 0; idx < fromLen; idx++ {
			toElem, err := e.Map(toModelElem.Interface(), fromValue.Index(idx).Interface(), options...)
			if err != nil {
				return nil, err
			}
			toValue.Index(idx).Set(reflect.ValueOf(toElem))
		}
		return toValue.Interface(), nil
	case reflect.Struct:
		// after switch
	default:
		return nil, mapToNotSupportKindError
	}

	// --------- struct ---------

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

	// match map type and option

	// find the first map rule of (fromType, toType) tuple
	var mapEntity *entity
	for _, entity := range e._entities {
		// already non-ptr
		if entity._fromType == fromType && entity._toType == toType {
			mapEntity = entity
			break
		}
	}
	// find all disposable map option for order
	var matchedOptions []*DisposableMapOption
	for _, option := range options {
		// already non-ptr
		if option._fromType == fromType && option._toType == toType {
			matchedOptions = append(matchedOptions, option)
		}
	}

	// matched map entity type
	if mapEntity != nil {
		// map through the found rule
		// for functions register order
		for _, rule := range mapEntity._rules {
			switch rule.(type) {
			case *_fieldDirectMapRule: // direct, through _mapFunc
				r := rule.(*_fieldDirectMapRule)
				toValue.FieldByIndex(r._toField.Index).Set(reflect.ValueOf(r._mapFunc(from)))
			case *_fieldFromMapRule:
				r := rule.(*_fieldFromMapRule)
				if r._isNest { // nest, through e.Map
					toFieldValue := toValue.FieldByIndex(r._toField.Index)
					fromFieldValue := fromValue.FieldByIndex(r._fromField.Index)
					// from field not nil
					if fromFieldValue.Kind() != reflect.Ptr || !fromFieldValue.IsNil() {
						toFieldNew, err := e.Map(toFieldValue.Interface(), fromFieldValue.Interface(), options...)
						if err != nil {
							return nil, err
						}
						toFieldValue.Set(reflect.ValueOf(toFieldNew))
					}
				} else { // copy
					toValue.FieldByIndex(r._toField.Index).Set(fromValue.FieldByIndex(r._fromField.Index))
				}
			case ExtraMapFunc: // _mapFunc
				r := rule.(ExtraMapFunc)
				extraValue := r(fromValue.Interface(), toValue.Interface())
				if reflect.TypeOf(extraValue) != toType {
					return nil, mapExtraFunctionReturnError
				}
				toValue = reflect.ValueOf(extraValue)
			}
		}
	}

	// map through the options
	for _, option := range matchedOptions {
		extraValue := option._mapFunc(fromValue.Interface(), toValue.Interface())
		if reflect.TypeOf(extraValue) != toType {
			return nil, mapExtraFunctionReturnError
		}
		toValue = reflect.ValueOf(extraValue)
	}

	return toValue.Interface(), nil
}
