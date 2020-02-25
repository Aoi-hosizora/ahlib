package xmapper

import (
	"errors"
	"reflect"
)

var (
	ErrMapToNil        = errors.New("could not map to a nil model")
	ErrMapToDiffType   = errors.New("could not map to a different kind of type")
	ErrMapToSmallArray = errors.New("could not map to a smaller array")
	ErrNotSupportType  = errors.New("could not map a non-ptr/non-slice/non-array/non-struct model")
	ErrExtraToDiffType = errors.New("could not use extra function to map to a different type")

	ErrMapPropFromNil   = errors.New("could not map property from nil")
	ErrMapPropToNil     = errors.New("could not map property to nil")
	ErrMapPropNotPtr    = errors.New("could not map property from non-ptr data")
	ErrMapPropNotStruct = errors.New("could not map property from non-struct data")
)

// check element and struct type
func _checkEl(model interface{}) reflect.Type {
	t := reflect.TypeOf(model)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("createMapper: could not create mapper by non-ptr and non-struct model")
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

// Create MapOption for Map
func NewMapOption(fromModel interface{}, toModel interface{}, extraMapFunc ExtraMapFunc) *MapOption {
	_fromType := _checkEl(fromModel)
	_toType := _checkEl(toModel)

	return &MapOption{
		_fromType: _fromType,
		_toType:   _toType,
		_mapFunc:  extraMapFunc,
	}
}

// options: for disposable map options
func (e *EntityMapper) Map(toModel interface{}, from interface{}, options ...*MapOption) (interface{}, error) {
	if toModel == nil {
		return nil, ErrMapToNil
	}
	if from == nil {
		return nil, nil
	}

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(toModel)
	kind := fromType.Kind()
	if kind != toType.Kind() {
		return nil, ErrMapToDiffType
	}

	// handle kind of type,generate all struct type
	switch kind {
	case reflect.Ptr:
		// get elem of the from(ptr) -> map elem as toElem -> put toElem to to(ptr)
		toModelElem := reflect.New(toType.Elem()).Elem()
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
		toValue := reflect.New(toType.Elem())
		toValue.Elem().Set(reflect.ValueOf(toElem))
		return toValue.Interface(), nil
	case reflect.Slice, reflect.Array:
		// get elem of the from(slice) -> map all elem as toElem -> put toElem to to(slice)
		// check fromArr and toArr size -> get elem of the from(array) -> map all elem as toElem -> put toElem to to(array)
		fromValue := reflect.ValueOf(from)
		fromLen := fromValue.Len()
		var toValue reflect.Value
		if kind == reflect.Slice {
			toValue = reflect.MakeSlice(toType, fromLen, fromLen)
		} else {
			toValue = reflect.New(toType).Elem()
			if fromLen > toValue.Len() { // check length
				return nil, ErrMapToSmallArray
			}
		}
		if fromLen == 0 {
			return toValue.Interface(), nil
		}
		toModelElem := reflect.New(toType.Elem()).Elem()
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
		return nil, ErrNotSupportType
	}

	// -------------------- struct map --------------------

	// ptr || struct
	toPtrValue := reflect.New(toType)
	fromPtrValue := reflect.New(fromType)
	fromPtrValue.Elem().Set(reflect.ValueOf(from))
	err := e.MapProp(fromPtrValue.Interface(), toPtrValue.Interface(), options...)
	if err != nil {
		return nil, err
	}

	return toPtrValue.Elem().Interface(), nil
}

func (e *EntityMapper) MapProp(fromPtr interface{}, toPtr interface{}, options ...*MapOption) error {
	if fromPtr == nil {
		return ErrMapPropFromNil
	}
	if toPtr == nil {
		return ErrMapPropToNil
	}
	fromType := reflect.TypeOf(fromPtr)
	toType := reflect.TypeOf(toPtr)
	if fromType.Kind() != reflect.Ptr || toType.Kind() != reflect.Ptr {
		return ErrMapPropNotPtr
	}
	if fromType.Elem().Kind() != reflect.Struct || toType.Elem().Kind() != reflect.Struct {
		return ErrMapPropNotStruct
	}
	fromType = fromType.Elem()
	toType = toType.Elem()
	fromValue := reflect.ValueOf(fromPtr).Elem()
	toValue := reflect.ValueOf(toPtr).Elem()

	// -------------------- same field --------------------

	for idx := 0; idx < toType.NumField(); idx++ {
		toField := toType.Field(idx)
		if fromField, ok := fromType.FieldByName(toField.Name); ok && fromField.Type == toField.Type {
			toValue.FieldByIndex(toField.Index).Set(fromValue.FieldByIndex(fromField.Index))
		}
	}

	// -------------------- map rule --------------------

	// find the first map rule of (fromType, toType) tuple
	var mapEntity *entity
	for _, entity := range e._entities {
		// already non-ptr
		if entity._fromType == fromType && entity._toType == toType {
			mapEntity = entity
			break
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
				toValue.FieldByIndex(r._toField.Index).Set(reflect.ValueOf(r._mapFunc(fromValue.Interface()))) // <<<
			case *_fieldSelfMapRule:
				r := rule.(*_fieldSelfMapRule)
				if r._isNest { // nest, through e.Map
					toFieldValue := toValue.FieldByIndex(r._toField.Index)
					fromFieldValue := fromValue.FieldByIndex(r._fromField.Index)
					// from field not nil
					if fromFieldValue.Kind() != reflect.Ptr || !fromFieldValue.IsNil() {
						toFieldNew, err := e.Map(toFieldValue.Interface(), fromFieldValue.Interface())
						if err != nil {
							return err
						}
						toFieldValue.Set(reflect.ValueOf(toFieldNew)) // <<<
					}
				} else { // copy
					toValue.FieldByIndex(r._toField.Index).Set(fromValue.FieldByIndex(r._fromField.Index)) // <<<
				}
			case ExtraMapFunc: // _mapFunc
				r := rule.(ExtraMapFunc)
				extraValue := r(fromValue.Interface(), toValue.Interface())
				if reflect.TypeOf(extraValue) != toType {
					return ErrExtraToDiffType
				}
				toValue.Set(reflect.ValueOf(extraValue)) // <<<
			}
		}
	}

	// -------------------- map option --------------------

	// find all disposable map option for order
	var matchedOptions []*MapOption
	for _, option := range options {
		// already non-ptr
		if option._fromType == fromType && option._toType == toType {
			matchedOptions = append(matchedOptions, option)
		}
	}

	// map through the options
	for _, option := range matchedOptions {
		extraValue := option._mapFunc(fromValue.Interface(), toValue.Interface())
		if reflect.TypeOf(extraValue) != toType {
			return ErrExtraToDiffType
		}
		toValue.Set(reflect.ValueOf(extraValue)) // <<<
	}

	return nil
}
