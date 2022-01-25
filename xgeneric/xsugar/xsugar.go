//go:build go1.18
// +build go1.18

package xsugar

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

// PanicIfErr returns value if err is nil, otherwise panics with error message.
func PanicIfErr[T any](value T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return value
}

// ValPtr returns a pointer pointed to the given value.
func ValPtr[T any](t T) *T {
	return &t
}

// PtrVal returns a value from the given pointer, returns the fallback value when pointer is nil.
func PtrVal[T any](t *T, o T) T {
	if t == nil {
		return o
	}
	return *t
}
