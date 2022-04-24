package xreflect

import (
	"math"
	"reflect"
	"unsafe"
)

// https://github.com/modern-go/reflect2
// reflect2.IFaceToEFace() // TODO
// TODO add to xreflect

// ====================
// kind checker related
// ====================

// IsIntKind checks whether given reflect.Kind represents integers or not.
func IsIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

// IsUintKind checks whether given reflect.Kind represents unsigned integers or not.
func IsUintKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Uintptr
}

// IsFloatKind checks whether given reflect.Kind represents floating points or not.
func IsFloatKind(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

// IsComplexKind checks whether given reflect.Kind represents complex numerics or not.
func IsComplexKind(kind reflect.Kind) bool {
	return kind == reflect.Complex64 || kind == reflect.Complex128
}

// IsNumericKind checks whether given reflect.Kind represents numerics or not.
//
// Numeric types: integers, unsigned integers, floating points, complex numerics.
func IsNumericKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 ||
		kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Uintptr ||
		kind == reflect.Float32 || kind == reflect.Float64 || kind == reflect.Complex64 || kind == reflect.Complex128
}

// IsCollectionKind checks whether given reflect.Kind represents collections or not, whose related reflect.Value's ``Len()'' method is invokable.
//
// Collection types: string, array, slice, map, chan.
func IsCollectionKind(kind reflect.Kind) bool {
	return kind == reflect.String || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// IsNillableKind checks whether given reflect.Kind represents nillable types or not, whose related reflect.Value's ``IsNil()'' method is invokable.
//
// Nillable types: ptr, func, interface, unsafePtr, slice, map, chan.
func IsNillableKind(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Func || kind == reflect.Interface || kind == reflect.UnsafePointer ||
		kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// =====================
// value checker related
// =====================

// IsNilValue checks whether given value is nil in its type or not. Note that this will also check the wrapped data of given interface{}.
func IsNilValue(v interface{}) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Ptr, reflect.Func, reflect.Interface, reflect.UnsafePointer, reflect.Slice, reflect.Map, reflect.Chan:
		return val.IsNil()
	}
	return false
}

// IsZeroValue checks whether given value is the zero value of its type or not. Note that all non-nil nillable collection values (such as empty []int{}
// and map[string]int{}) are regarded as not zero.
func IsZeroValue(v interface{}) bool {
	if v == nil {
		return true
	}
	zero := reflect.Zero(reflect.TypeOf(v)).Interface()
	return reflect.DeepEqual(v, zero)
	// return reflect.ValueOf(v).IsZero()
}

// IsEmptyCollection checks whether given collection value is empty or not. Note that empty means its value is nil or its length is zero.
func IsEmptyCollection(v interface{}) bool {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return val.Len() == 0 // call val.Len() directly without unnecessary val.IsNil()
	}
	return false // including reflect.Invalid
}

// IsEmptyValue checks whether given value is empty or not. Note that empty means zero value, nil value, zero item and zero field, and this works
// almost the same as json.isEmptyValue for "omitempty".
func IsEmptyValue(v interface{}) bool {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		f := val.Float()
		return math.Float64bits(f) == 0
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return val.Len() == 0 // => val.IsNil() || val.Len() == 0, do not check fields recursively
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer, reflect.Func:
		return val.IsNil()
	case reflect.Struct:
		return val.NumField() == 0 // do not check fields recursively (different with val.Zero())
	default:
		// reflect.Invalid, that is untyped "nil"
		return true
	}
}

// =====================
// numeric value related
// =====================

// Float64Value returns the float64 value for given numeric value, returns false if given value is not numeric value.
func Float64Value(v interface{}) (float64, bool) {
	switch vv := v.(type) {
	case int:
		return float64(vv), true
	case int8:
		return float64(vv), true
	case int16:
		return float64(vv), true
	case int32:
		return float64(vv), true
	case int64:
		return float64(vv), true
	case uint:
		return float64(vv), true
	case uint8:
		return float64(vv), true
	case uint16:
		return float64(vv), true
	case uint32:
		return float64(vv), true
	case uint64:
		return float64(vv), true
	case uintptr:
		return float64(vv), true
	case float32:
		return float64(vv), true
	case float64:
		return vv, true
	}
	return 0, false
}

// Uint64Value returns the uint64 value for given numeric value, returns false if given value is not numeric value.
func Uint64Value(v interface{}) (uint64, bool) {
	switch vv := v.(type) {
	case int:
		return uint64(vv), true
	case int8:
		return uint64(vv), true
	case int16:
		return uint64(vv), true
	case int32:
		return uint64(vv), true
	case int64:
		return uint64(vv), true
	case uint:
		return uint64(vv), true
	case uint8:
		return uint64(vv), true
	case uint16:
		return uint64(vv), true
	case uint32:
		return uint64(vv), true
	case uint64:
		return vv, true
	case uintptr:
		return uint64(vv), true
	case float32:
		return uint64(vv), true
	case float64:
		return uint64(vv), true
	}
	return 0, false
}

// ========================
// unexported field related
// ========================

// GetUnexportedField gets the reflect.Value of unexported struct field's. Note that this is an unsafe function.
//
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).Interface().(gin.HandlersChain)             // (*app).noMethod is a gin.HandlersChain
// 	GetUnexportedField(FieldValueOf(app, "noMethod")).Interface().(gin.HandlersChain)                                   // <- or in this way
// 	GetUnexportedField(reflect.ValueOf(trans).Elem().FieldByName("translations")).MapIndex(reflect.ValueOf("required")) // (*trans).translations is a map[string]xxx
// 	GetUnexportedField(FieldValueOf(trans, "translations")).Interface().(gin.HandlersChain)                             // <- or in this way
func GetUnexportedField(field reflect.Value) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
}

// SetUnexportedField sets reflect.Value to unexported struct field. Note that this is an unsafe function.
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
func FieldValueOf(v interface{}, name string) reflect.Value {
	if v == nil {
		panic(panicNilInterface)
	}
	val := reflect.ValueOf(v)
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

// Attention:
//
// 1. For searching function by name, please refer to https://github.com/alangpierce/go-forceexport/blob/8f1d6941cd/forceexport.go#L10-L22
// and http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions.
//
// 2. For calling unexported struct methods, please refer to https://github.com/spance/go-callprivate/blob/master/examples.go#L20-L22.

// ==============
// mass functions
// ==============

// eface keeps the same as runtime.eface, which is the internal representation of interface{}.
type eface struct {
	_type uintptr // *runtime._type
	data  unsafe.Pointer
}

// HasZeroEface checks whether given interface value has no type information or wrapped data (more than `== nil`). Note that this is an unsafe function.
func HasZeroEface(v interface{}) bool {
	e := (*eface)(unsafe.Pointer(&v))
	return e._type == 0 || uintptr(e.data) == 0
}

// DeepEqualInValue checks whether given two values are deeply equal without considering their types. Note that it checks by checking type convertable
// and comparing after type conversion.
func DeepEqualInValue(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	val1, val2 := reflect.ValueOf(v1), reflect.ValueOf(v2)
	if !val1.IsValid() || !val2.IsValid() {
		return false
	}

	// check convertable, and compare after type conversion
	type1, type2 := val1.Type(), val2.Type()
	if type1.ConvertibleTo(type2) {
		return reflect.DeepEqual(val1.Convert(type2).Interface(), v2)
	}
	if type2.ConvertibleTo(type1) {
		return reflect.DeepEqual(v1, val2.Convert(type1).Interface())
	}
	return false // not equal
}

// IsSamePointer checks whether given two values are the same pointer types, and whether they point to the same address.
func IsSamePointer(p1, p2 interface{}) bool {
	val1, val2 := reflect.ValueOf(p1), reflect.ValueOf(p2)
	if val1.Kind() != reflect.Ptr || val2.Kind() != reflect.Ptr || val1.Type() != val2.Type() {
		return false
	}

	// compare addresses which two pointers point to
	return p1 == p2
}

const (
	panicNonNilMap = "xreflect: not a non-nil map"
)

// hmap keeps the same as runtime.hmap, which represents a header for a Go map.
type hmap struct {
	count int
	flags uint8
	B     uint8 // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	// ...
}

// GetMapB returns the B value from given map value. Note that this is an unsafe function, and returned value may differ in different Go versions.
func GetMapB(m interface{}) (b uint8) {
	if m == nil {
		panic(panicNonNilMap)
	}
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map || val.IsNil() {
		panic(panicNonNilMap)
	}

	// https://hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh
	e := (*eface)(unsafe.Pointer(&m))
	h := (*hmap)(e.data)
	return h.B
}

// GetMapBuckets returns the B value and the buckets count from given map value. Note that this is an unsafe function, and returned value may
// differ in different Go versions.
func GetMapBuckets(m interface{}) (b uint8, buckets uint64) {
	b = GetMapB(m)
	buckets = uint64(math.Pow(2, float64(b))) // 2^B
	return b, buckets
}
