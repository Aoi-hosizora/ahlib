package xreflect

import (
	"fmt"
	"reflect"
)

func ElemType(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func ElemValue(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func IsEqual(val1, val2 interface{}) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}
	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1, v2)
		} else {
			return reflect.DeepEqual(v1, v2.Interface())
		}
	} else {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1.Interface(), v2)
		} else {
			return reflect.DeepEqual(v1.Interface(), v2.Interface())
		}
	}
}

func GetInt(i interface{}) (int64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	}
	return 0, false
}

func GetUint(i interface{}) (uint64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	}
	return 0, false
}

func GetFloat(i interface{}) (float64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

func GetString(i interface{}) (string, bool) {
	s, ok := i.(string)
	if ok {
		return s, true
	}
	return "", false
}

func GetBool(i interface{}) (bool, bool) {
	s, ok := i.(bool)
	if ok {
		return s, true
	}
	return false, false
}

// ValueSize represents some different types of value size.
type ValueSize struct {
	fi int64
	fu uint64
	ff float64
}

func (v *ValueSize) Int() int64 {
	if v.fi != 0 {
		return v.fi
	}
	if v.fu != 0 {
		return int64(v.fu)
	}
	return int64(v.ff)
}

func (v *ValueSize) Uint() uint64 {
	if v.fu != 0 {
		return v.fu
	}
	if v.fi != 0 {
		return uint64(v.fi)
	}
	return uint64(v.ff)
}

func (v *ValueSize) Float() float64 {
	if v.ff != 0 {
		return v.ff
	}
	if v.fi != 0 {
		return float64(v.fi)
	}
	return float64(v.fu)
}

// Get value's size and return ValueSize.
//
// For numbers, it is the value.
// For strings, it is the number of characters.
// For slices, arrays, maps, it is the number of items.
func GetValueSize(i interface{}) (*ValueSize, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &ValueSize{fi: val.Int()}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return &ValueSize{fu: val.Uint()}, nil
	case reflect.Float32, reflect.Float64:
		return &ValueSize{ff: val.Float()}, nil
	case reflect.String:
		return &ValueSize{fi: int64(len([]rune(val.String())))}, nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return &ValueSize{fi: int64(val.Len())}, nil
	case reflect.Bool:
		v := 0
		if val.Bool() {
			v = 1
		}
		return &ValueSize{fi: int64(v)}, nil
	}
	return nil, fmt.Errorf("bad field type %T", val.Interface())
}
