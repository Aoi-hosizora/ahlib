package xreflect

import (
	"math"
	"reflect"
	"unsafe"
)

// =================
// IsXXXKind related
// =================

// IsIntKind checks if given reflect.Kind is int kinds or not.
func IsIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

// IsUintKind checks if given reflect.Kind is uint kinds or not.
func IsUintKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Uintptr
}

// IsFloatKind checks if given reflect.Kind is float kinds or not.
func IsFloatKind(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

// IsComplexKind checks if given reflect.Kind is complex kinds or not.
func IsComplexKind(kind reflect.Kind) bool {
	return kind == reflect.Complex64 || kind == reflect.Complex128
}

// IsLenGettableKind checks if given reflect.Kind's related reflect.Value can use Len() method or not.
func IsLenGettableKind(kind reflect.Kind) bool {
	return kind == reflect.String || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// IsNillableKind checks if given reflect.Kind's related reflect.Value can use IsNil() method or not.
func IsNillableKind(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Func || kind == reflect.Interface || kind == reflect.UnsafePointer ||
		kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// ========================
// unexported field related
// ========================

// GetUnexportedField gets the reflect.Value of unexported struct field's.
//
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).Interface().(gin.HandlersChain)             // (*app).noMethod is a gin.HandlersChain
// 	GetUnexportedField(FieldValueOf(app, "noMethod")).Interface().(gin.HandlersChain)                                   // <- or in this way
// 	GetUnexportedField(reflect.ValueOf(trans).Elem().FieldByName("translations")).MapIndex(reflect.ValueOf("required")) // (*trans).translations is a map[string]xxx
// 	GetUnexportedField(FieldValueOf(trans, "translations")).Interface().(gin.HandlersChain)                             // <- or in this way
func GetUnexportedField(field reflect.Value) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
}

// SetUnexportedField sets reflect.Value to unexported struct field, this can also be implemented by using the reflect.Value returned from GetUnexportedField.
//
// Example:
// 	SetUnexportedField(reflect.ValueOf(ctx).Elem().FieldByName("fullPath"), reflect.ValueOf(newFullPath)) // (*ctx).fullPath and newFullPath is a string
// 	SetUnexportedField(FieldValueOf(ctx, "fullPath"), reflect.ValueOf(newFullPath))                       // <- or in this way
// 	SetUnexportedField(reflect.ValueOf(val).Elem().FieldByName("tagNameFunc"), reflect.ValueOf(nilFunc))  // (*val).tagNameFunc and nilFunc is a func
// 	SetUnexportedField(FieldValueOf(val, "tagNameFunc"), reflect.ValueOf(newFullPath))                    // <- or in this way
func SetUnexportedField(field reflect.Value, value reflect.Value) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(value)
}

const (
	panicNilInterface           = "xreflect: nil interface"
	panicNilPtr                 = "xreflect: nil pointer"
	panicNonStructOrPtrOfStruct = "xreflect: not a struct or pointers of struct"
	panicNonexistentField       = "xreflect: nonexistent struct field"
)

// FieldValueOf returns the reflect.Value of specific struct field from given struct or pointers of struct.
//
// Example:
// 	FieldValueOf(app, "noMethod")       // equals to reflect.ValueOf(app)[.Elem()*].FieldByName("noMethod")
// 	FieldValueOf(trans, "translations") // equals to reflect.ValueOf(trans)[.Elem()*].FieldByName("translations")
func FieldValueOf(i interface{}, name string) reflect.Value {
	if i == nil {
		panic(panicNilInterface)
	}
	val := reflect.ValueOf(i)
	for val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr && val.IsNil() {
		panic(panicNilPtr)
	}
	if val.Kind() != reflect.Struct {
		panic(panicNonStructOrPtrOfStruct)
	}
	fval := val.FieldByName(name)
	if !fval.IsValid() { // not existed
		panic(panicNonexistentField)
	}
	return fval
}

// ==============
// mass functions
// ==============

// IsEmptyValue checks if a value is an empty value, this function do never panic for all parameters.
// Support types: (all types)
// 	1. numeric:    int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection: string, array, slice, map, chan.
// 	3. wrapper:    interface, ptr, unsafePtr.
// 	4. composite:  struct.
// 	5. function:   func.
func IsEmptyValue(i interface{}) bool {
	return isEmptyValueInternal(reflect.ValueOf(i))
}

// isEmptyValueInternal is the internal implementation of IsEmptyValue.
func isEmptyValueInternal(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(val.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.String, reflect.Array:
		return val.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Chan:
		return val.IsNil() || val.Len() == 0
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer, reflect.Func:
		return val.IsNil()
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if !isEmptyValueInternal(val.Field(i)) {
				return false
			}
		}
		return true
	default:
		// reflect.Invalid, that is (SomeInterface)(nil)
		return true
	}
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
