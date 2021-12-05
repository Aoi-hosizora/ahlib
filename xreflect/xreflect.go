package xreflect

import (
	"math"
	"reflect"
	"unsafe"
)

// GetUnexportedField gets the unexported struct field's reflect.Value.
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).Interface().(gin.HandlersChain)
// 	GetUnexportedField(reflect.ValueOf(trans).Elem().FieldByName("translations")).MapIndex(reflect.ValueOf("required"))
func GetUnexportedField(field reflect.Value) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
}

// SetUnexportedField sets reflect.Value to the unexported struct field, or you can also use GetUnexportedField's returned reflect.Value to set value.
// Example:
// 	SetUnexportedField(reflect.ValueOf(c).Elem().FieldByName("fullPath"), reflect.ValueOf(newFullPath))
// 	SetUnexportedField(reflect.ValueOf(v).Elem().FieldByName("tagNameFunc"), reflect.ValueOf(nilFunc))
func SetUnexportedField(field reflect.Value, value reflect.Value) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(value)
}

// GetInt returns the int64 value from int, int8, int16, int32 and int64 interface.
func GetInt(i interface{}) (int64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	}
	return 0, false
}

// GetUint returns the uint64 value from uint, uint8, uint16, uint32, uint64 and uintptr interface.
func GetUint(i interface{}) (uint64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	}
	return 0, false
}

// GetFloat returns the float64 value from float32 and float64 interface.
func GetFloat(i interface{}) (float64, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

// GetComplex returns the complex128 value from complex64 and complex128 interface.
func GetComplex(i interface{}) (complex128, bool) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	}
	return 0, false
}

// GetBool returns the bool value from bool interface.
func GetBool(i interface{}) (bool, bool) {
	s, ok := i.(bool)
	if ok {
		return s, true
	}
	return false, false
}

// GetString returns the string value from string interface.
func GetString(i interface{}) (string, bool) {
	s, ok := i.(string)
	if ok {
		return s, true
	}
	return "", false
}

// IsEmptyValue checks if a value is an empty value, this function do never panic for all parameters.
// Support types: (all types)
// 	1. numeric:    int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection: string, array, slice, map, chan.
// 	3. wrapper:    interface, ptr, unsafePtr.
// 	4. composite:  struct.
// 	5. function:   func.
func IsEmptyValue(i interface{}) bool {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return val.Complex() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return val.Len() == 0
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer, reflect.Func:
		return val.IsNil()
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if !IsEmptyValue(val.Field(i).Interface()) {
				return false
			}
		}
		return true
	}

	// invalid, that is (interface{})(nil)
	return true
}

const (
	panicNilMap = "xreflect: nil map"
	panicNonMap = "xreflect: not a map"
)

// GetMapB returns the B value from the inputted map value. Note that this is an unsafe function, and the returned value may change in different Go versions.
func GetMapB(m interface{}) uint8 {
	if m == nil {
		panic(panicNilMap)
	}
	typ := reflect.TypeOf(m)
	if typ.Kind() != reflect.Map {
		panic(panicNonMap)
	}

	type eface struct {
		_type unsafe.Pointer
		data  unsafe.Pointer
	}
	type hmap struct {
		count int
		flags uint8
		B     uint8
		// ...
	}

	// https://hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh
	ei := *(*eface)(unsafe.Pointer(&m))
	mobj := *(*hmap)(ei.data)
	return mobj.B
}

// GetMapBuckets returns the B value and the buckets count from the inputted map value. Note that this is an unsafe function, and the returned B value may
// change in different Go versions, while the buckets count will always equal to 2^B.
func GetMapBuckets(m interface{}) (b uint8, buckets uint64) {
	b = GetMapB(m)
	buckets = uint64(math.Pow(2, float64(b)))
	return b, buckets
}
