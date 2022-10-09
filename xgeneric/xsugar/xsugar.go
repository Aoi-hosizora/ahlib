//go:build go1.18
// +build go1.18

package xsugar

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// ========================
// xcondition compatibility
// ========================

// IfThen returns value if condition is true, otherwise returns the default value of type T.
func IfThen[T any](condition bool, value T) T {
	if condition {
		return value
	}
	var tt T
	return tt
}

// IfThenElse returns value1 if condition is true, otherwise returns value2.
func IfThenElse[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

// If is the short form of IfThenElse.
func If[T any](cond bool, v1, v2 T) T {
	return IfThenElse(cond, v1, v2)
}

// DefaultIfNil returns value if it is not nil, otherwise returns defaultValue. Note that this also checks the
// wrapped data of given value.
func DefaultIfNil[T any](value, defaultValue T) T {
	if !isNilValue(any(value)) {
		return value
	}
	return defaultValue
}

// PanicIfNil returns value if it is not nil, otherwise panics with given the first panicValue. Note that this
// also checks the wrapped data of given value.
func PanicIfNil[T any](value T, panicValue ...any) T {
	if !isNilValue(any(value)) {
		return value
	}
	if len(panicValue) == 0 || panicValue[0] == nil {
		panic(fmt.Sprintf("xcondition: nil value for %T", value))
	}
	panic(panicValue[0])
}

// Un is the short form of PanicIfNil without panicValue, which means "unwrap nil with builtin panic value".
func Un[T any](v T) T {
	return PanicIfNil(v)
}

// Unp is the short form of PanicIfNil with panicValue, which means "unwrap nil with custom panic value".
func Unp[T any](v T, panicV any) T {
	return PanicIfNil(v, panicV)
}

// PanicIfErr returns value if given err is nil, otherwise panics with given error message.
func PanicIfErr[T any](value T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return value
}

// PanicIfErr2 returns value1 and value2 if given err is nil, otherwise panics with given error message.
func PanicIfErr2[T1, T2 any](value1 T1, value2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err.Error())
	}
	return value1, value2
}

// PanicIfErr3 returns value1, value2 and value3 if given err is nil, otherwise panics with given error message.
func PanicIfErr3[T1, T2, T3 any](value1 T1, value2 T2, value3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err.Error())
	}
	return value1, value2, value3
}

// Ue is the short form of PanicIfErr, which means "unwrap error".
func Ue[T any](v T, err error) T {
	return PanicIfErr(v, err)
}

// Ue2 is the short form of PanicIfErr2, which means "unwrap error".
func Ue2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	return PanicIfErr2(v1, v2, err)
}

// Ue3 is the short form of PanicIfErr3, which means "unwrap error".
func Ue3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	return PanicIfErr3(v1, v2, v3, err)
}

// ==============
// mass functions
// ==============

// ValPtr returns a pointer pointed to given value.
func ValPtr[T any](t T) *T {
	return &t
}

// PtrVal returns a value from given pointer, returns the fallback value when pointer is nil.
func PtrVal[T any](t *T, o T) T {
	if t == nil {
		return o
	}
	return *t
}

// Incr increments the value of given Real first, and then returns it, this is the same as C "++n" expression.
func Incr[T Real](n *T) T {
	*n++
	return *n
}

// Decr decrements the value of given Real first, and then returns it, this is the same as C "--n" expression.
func Decr[T Real](n *T) T {
	*n--
	return *n
}

// RIncr returns the value of given Real first, and then increments it, this is the same as C "n++" expression.
func RIncr[T Real](n *T) T {
	v := *n
	*n++
	return v
}

// RDecr returns the value of given Real first, and then decrements it, this is the same as C "n--" expression.
func RDecr[T Real](n *T) T {
	v := *n
	*n--
	return v
}

// UnmarshalJson unmarshals given bytes to T, just like json.Unmarshal.
func UnmarshalJson[T any](bs []byte, t T) (T, error) {
	err := json.Unmarshal(bs, t)
	if err != nil {
		var tt T
		return tt, err
	}
	return t, err
}

// FastStoa casts slice to array pointer that has the same underlying data.
// Note that this is an unsafe function, and this function won't check any parameter type.
//
// Example:
// 	var _ *[2]int32 = FastStoa[[]int32, [2]int32]([]int32{3, 2, 1}) -> normal usage, got a 2-len array
// 	var _ *[12]int8 = FastStoa[[]int32, [12]int8]([]int32{3, 2, 1}) -> using with different types is also allowed
// 	var _ *[3]int64 = FastStoa[[]int32, [3]int64]([]int32{3, 2, 1}) -> out of length is also allowed, but in undefined behavior
func FastStoa[TSlice ~[]TItem, TArray, TItem any](slice TSlice) *TArray {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	return (*TArray)(unsafe.Pointer(sh.Data))
}

// FastAtos casts array pointer to slice that has the same underlying data but with given length.
// Note that this is an unsafe function, and this function won't check any parameter type.
//
// Example:
// 	var _ []int32 = FastAtos[[3]int32, []int32](&[...]int32{3, 2, 1}, 2) -> normal usage, got a 2-len and 2-cap slice
// 	var _ []int8  = FastAtos[[3]int32, []int8](&[...]int32{3, 2, 1}, 12) -> using with different types is also allowed
// 	var _ []int64 = FastAtos[[3]int32, []int64](&[...]int32{3, 2, 1}, 3) -> out of length is also allowed, but in undefined behavior
func FastAtos[TArray any, TSlice ~[]TItem, TItem any](array *TArray, length int) TSlice {
	sh := reflect.SliceHeader{Len: length, Cap: length, Data: uintptr(unsafe.Pointer(array))}
	return *(*TSlice)(unsafe.Pointer(&sh))
}

// isNilValue keeps the same as xreflect.IsNilValue, checks whether given value is nil in its type or not.
func isNilValue(v interface{}) bool {
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

// =====================
// constraint interfaces
// =====================

// Signed is a constraint that permits any signed integer type, that is: int / int8 / int16 / int32 / int64 / all types
// whose underlying type is one of these types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type, that is: uint / uint8 / uint16 / uint32 / uint64 /
// uintptr / all types whose underlying type is one of these types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint that permits any floating-point type, that is: float32 / float64 / all types whose underlying
// type is one of these types.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type, that is: complex64 / complex128 / all types whose
// underlying type is one of these types.
type Complex interface {
	~complex64 | ~complex128
}

// Integer is a constraint that permits any integer type, including Signed and Unsigned.
type Integer interface {
	Signed | Unsigned
}

// Real is a constraint that permits any real numeric type, including Integer and Float.
type Real interface {
	Integer | Float
}

// Numeric is a constraint that permits any numeric type, including Real and Complex.
type Numeric interface {
	Real | Complex
}

// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >, that is: Real
// (Signed / Unsigned / Float) / string / all types whose underlying type is one of these types.
type Ordered interface {
	Real | ~string
}
