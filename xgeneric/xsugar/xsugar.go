//go:build go1.18
// +build go1.18

package xsugar

import (
	"encoding/json"
	"fmt"
)

// ==============
// mass functions
// ==============

// IfThen returns value if condition is true, otherwise returns the default value of type T.
func IfThen[T any](condition bool, value T) T {
	if condition {
		return value
	}
	var v T
	return v
}

// IfThenElse returns value1 if condition is true, otherwise returns value2.
func IfThenElse[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

// DefaultIfNil returns value if it is not nil, otherwise returns defaultValue.
func DefaultIfNil[T any](value, defaultValue T) T {
	if any(value) != nil { // TODO nil checker
		return value
	}
	return defaultValue
}

const (
	panicNilValue = "xcondition: nil value for %T"
)

// PanicIfNil returns value if it is not nil, otherwise panics with given v.
func PanicIfNil[T any](value T, v any) T {
	if any(value) != nil { // TODO nil checker
		return value
	}
	if v == nil {
		panic(fmt.Sprintf(panicNilValue, value))
	}
	panic(v)
}

// PanicIfErr returns value if given err is nil, otherwise panics with error message.
func PanicIfErr[T any](value T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return value
}

// PanicIfErr2 returns value1 and value2 if given err is nil, otherwise panics with error message.
func PanicIfErr2[T, K any](value1 T, value2 K, err error) (T, K) {
	if err != nil {
		panic(err.Error())
	}
	return value1, value2
}

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

// Incr increments the value of given Real and then returns it, this is the same as C "++n" expression.
func Incr[T Real](n *T) T {
	*n++
	return *n
}

// Decr decrements the value of given Real and then returns it, this is the same as C "--n" expression.
func Decr[T Real](n *T) T {
	*n--
	return *n
}

// RIncr returns the value of given Real and then increments it, this is the same as C "n++" expression.
func RIncr[T Real](n *T) T {
	v := *n
	*n++
	return v
}

// RDecr returns the value of given Real and then decrements it, this is the same as C "n--" expression.
func RDecr[T Real](n *T) T {
	v := *n
	*n--
	return v
}

// UnmarshalJson unmarshals given byte array to T, just like json.Unmarshal, but with T returned.
func UnmarshalJson[T any](bs []byte, t T) (T, error) {
	err := json.Unmarshal(bs, t)
	if err != nil {
		var tt T
		return tt, err
	}
	return t, err
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
