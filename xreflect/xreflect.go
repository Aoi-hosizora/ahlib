package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"unsafe"
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

// Get the unexported field value
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).(gin.HandlersChain)
func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// Set the unexported field to value
// Example:
// 	SetUnexportedField(reflect.ValueOf(c).Elem().FieldByName("fullPath"), fullPath)
func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}

// IsEqual is the copy of xtesting.IsEqual.
func IsEqual(val1, val2 interface{}) bool {
	return xtesting.IsEqual(val1, val2)
}

// BoolVal is the same with xnumber.Bool.
func BoolVal(b bool) int {
	if b {
		return 1
	}
	return 0
}

func GetStructFields(i interface{}) []reflect.StructField {
	typ := reflect.TypeOf(i)
	fnum := typ.NumField()
	fields := make([]reflect.StructField, fnum)
	for idx := 0; idx < fnum; idx++ {
		fields[idx] = typ.Field(idx)
	}
	return fields
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
