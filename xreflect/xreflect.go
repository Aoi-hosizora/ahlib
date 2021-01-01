package xreflect

import (
	"reflect"
	"unsafe"
)

// ElemType returns the actual reflect.Type of a reflect.Ptr kind value.
func ElemType(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// ElemValue returns the actual reflect.Value of a reflect.Ptr kind value.
func ElemValue(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

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

// GetStructFields gets a slice of reflect.StructField from the given struct value.
func GetStructFields(i interface{}) []reflect.StructField {
	typ := reflect.TypeOf(i)
	fnum := typ.NumField()
	fields := make([]reflect.StructField, fnum)
	for idx := 0; idx < fnum; idx++ {
		fields[idx] = typ.Field(idx)
	}
	return fields
}

// GetInt returns the int64 value from int, int8, int32, int64.
func GetInt(i interface{}) (int64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	}
	return 0, false
}

// GetUint returns the uint64 value from uint, uint8, uint16, uint32, uint64, uintptr.
func GetUint(i interface{}) (uint64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	}
	return 0, false
}

// GetFloat returns the float64 value from float32, float64.
func GetFloat(i interface{}) (float64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

// GetComplex returns the complex128 value from complex64, complex128.
func GetComplex(i interface{}) (complex128, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	}
	return 0, false
}

// GetBool returns a bool value from a bool interface.
func GetBool(i interface{}) (bool, bool) {
	s, ok := i.(bool)
	if ok {
		return s, true
	}
	return false, false
}

// GetString returns a string value from a string interface.
func GetString(i interface{}) (string, bool) {
	s, ok := i.(string)
	if ok {
		return s, true
	}
	return "", false
}
