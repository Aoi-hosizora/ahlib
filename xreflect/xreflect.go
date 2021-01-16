package xreflect

import (
	"reflect"
	"unsafe"
)

// GetUnexportedField gets the unexported field value.
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).(gin.HandlersChain)
func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// SetUnexportedField sets the unexported field to value.
// Example:
// 	SetUnexportedField(reflect.ValueOf(c).Elem().FieldByName("fullPath"), fullPath)
func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}

// IsEmptyValue checks if a value is a empty value. As for numeric values, it equals to check zero; as for
// collection values, it equals to check the size; as for pointer adn interface, it equals to check nil,
// Note that this is different from xtesting.IsObjectEmpty.
func IsEmptyValue(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Chan:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// GetInt returns the int64 value from int, int8, int32, int64 interface.
func GetInt(i interface{}) (int64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	}
	return 0, false
}

// GetUint returns the uint64 value from uint, uint8, uint16, uint32, uint64, uintptr interface.
func GetUint(i interface{}) (uint64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	}
	return 0, false
}

// GetFloat returns the float64 value from float32, float64 interface.
func GetFloat(i interface{}) (float64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

// GetComplex returns the complex128 value from complex64, complex128 interface.
func GetComplex(i interface{}) (complex128, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	}
	return 0, false
}

// GetBool returns a bool value from bool interface.
func GetBool(i interface{}) (bool, bool) {
	s, ok := i.(bool)
	if ok {
		return s, true
	}
	return false, false
}

// GetString returns a string value from string interface.
func GetString(i interface{}) (string, bool) {
	s, ok := i.(string)
	if ok {
		return s, true
	}
	return "", false
}
