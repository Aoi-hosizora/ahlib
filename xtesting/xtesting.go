package xtesting

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

// failTest outputs the error message and fails the test.
func failTest(t *testing.T, skip int, msg string) bool {
	flag := ""

	if skip >= 0 {
		_, file, line, _ := runtime.Caller(skip + 1)
		fmt.Printf("%s%s:%d %s\n", flag, path.Base(file), line, msg)
	}
	t.Fail()

	return false
}

// Equal asserts that two objects are equal.
func Equal(t *testing.T, expected, actual interface{}) bool {
	if err := validateEqualArgs(expected, actual); err != nil {
		return failTest(t, 1, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (%v)", expected, actual, err))
	}

	if !IsObjectEqual(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("Equal: expected `%#v`, actual `%#v`", expected, actual))
	}

	return true
}

// NotEqual asserts that the specified values are Not equal.
func NotEqual(t *testing.T, expected, actual interface{}) bool {
	if err := validateEqualArgs(expected, actual); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotEqual: invalid operation `%#v` != `%#v` (%v)", expected, actual, err))
	}

	if IsObjectEqual(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("NotEqual: expected not to be `%#v`", actual))
	}

	return true
}

// EqualValue asserts that two objects are equal or convertible to the same types and equal.
func EqualValue(t *testing.T, expected, actual interface{}) bool {
	if !IsObjectValueEqual(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("EqualValue: expected `%#v`, actual `%#v`", expected, actual))
	}

	return true
}

// NotEqualValue asserts that two objects are not equal even when converted to the same type.
func NotEqualValue(t *testing.T, expected, actual interface{}) bool {
	if IsObjectValueEqual(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("NotEqualValue: expected not to be `%#v`", actual))
	}

	return true
}

// SamePointer asserts that two pointers reference the same object.
func SamePointer(t *testing.T, expected, actual interface{}) bool {
	if !IsPointerSame(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expected `%#v` (%p), actual `%#v` (%p)", expected, expected, actual, actual))
	}

	return true
}

// NotSamePointer asserts that two pointers do not reference the same object.
func NotSamePointer(t *testing.T, expected, actual interface{}) bool {
	if IsPointerSame(expected, actual) {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expected not be `%#v` (%p)", actual, actual))
	}

	return true
}

// Nil asserts that the specified object is nil.
func Nil(t *testing.T, object interface{}) bool {
	if !IsObjectNil(object) {
		return failTest(t, 1, fmt.Sprintf("Nil: expected `nil`, actual `%#v`", object))
	}

	return true
}

// NotNil asserts that the specified object is not nil.
func NotNil(t *testing.T, object interface{}) bool {
	if IsObjectNil(object) {
		return failTest(t, 1, fmt.Sprintf("NotNil, expected not to be `nil`, actual `%#v`", object))
	}

	return true
}

// True asserts that the specified value is true.
func True(t *testing.T, value bool) bool {
	if !value {
		return failTest(t, 1, fmt.Sprintf("True: expected `true`, actual `%#v`", value))
	}

	return true
}

// False asserts that the specified value is false.
func False(t *testing.T, value bool) bool {
	if value {
		return failTest(t, 1, fmt.Sprintf("False: expected to be `false`, actual `%#v`", value))
	}

	return true
}

// Zero asserts that the specified object is the zero value for its type.
func Zero(t *testing.T, object interface{}) bool {
	if !IsObjectZero(object) {
		return failTest(t, 1, fmt.Sprintf("Zero: expected to be zero value, actual `%#v`", object))
	}

	return true
}

// NotEmpty asserts that the specified object is not the zero value for its type.
func NotZero(t *testing.T, object interface{}) bool {
	if IsObjectZero(object) {
		return failTest(t, 1, fmt.Sprintf("NotZero: expected not to be zero value, actual `%#v`", object))
	}

	return true
}

// Empty asserts that the specified object is empty.
func Empty(t *testing.T, object interface{}) bool {
	if !IsObjectEmpty(object) {
		return failTest(t, 1, fmt.Sprintf("Empty: expected to be empty value, actual `%#v`", object))
	}

	return true
}

// NotEmpty asserts that the specified object is not empty.
func NotEmpty(t *testing.T, object interface{}) bool {
	if IsObjectEmpty(object) {
		return failTest(t, 1, fmt.Sprintf("NotEmpty: expected not to be empty value, actual `%#v`", object))
	}

	return true
}

// Contain asserts that the specified container contains the specified substring or element.
// Support string, array, slice or map.
func Contain(t *testing.T, container, object interface{}) bool {
	ok, found := includeElement(container, object)

	if !ok {
		return failTest(t, 1, fmt.Sprintf("Contain: invalid operator len(`%#v`)", container))
	}
	if !found {
		return failTest(t, 1, fmt.Sprintf("Contain: `%#v` is expected to contain `%#v`", container, object))
	}

	return true
}

// NotContain asserts that the specified container does not contain the specified substring or element.
// Support string, array, slice or map.
func NotContain(t *testing.T, container, object interface{}) bool {
	ok, found := includeElement(container, object)

	if !ok {
		return failTest(t, 1, fmt.Sprintf("NotContain: invalid operator len(`%#v`)", container))
	}
	if found {
		return failTest(t, 1, fmt.Sprintf("NotContain: `%#v` is expected not to contain `%#v`", container, object))
	}

	return true
}

// ElementMatch asserts that the specified listA is equal to specified listB ignoring the order of the elements.
// If there are duplicate elements, the number of appearances of each of them in both lists should match.
func ElementMatch(t *testing.T, listA, listB interface{}) bool {
	if IsObjectEmpty(listA) && IsObjectEmpty(listB) {
		return true
	}

	if err := validateArgIsList(listA, listB); err != nil {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: invalid operator: `%#v` <-> `%#v` (%v)", listA, listB, err))
	}

	extraA, extraB := diffLists(listA, listB)
	if len(extraA) != 0 || len(extraB) != 0 {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: `%#v` and `%#v` are expected to match each other", listA, listB))
	}

	return true
}

// Implements asserts that an object is implemented by the specified interface.
func Implements(t *testing.T, interfaceObject interface{}, object interface{}) bool {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return failTest(t, 1, fmt.Sprintf("Implements: invalid operation for `nil`"))
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return failTest(t, 1, fmt.Sprintf("Implements: %T expected to implement `%v`, actual not implment.", object, interfaceObject))
	}

	return true
}

// IsType asserts that the specified objects are of the same type.
func IsType(t *testing.T, expected interface{}, object interface{}) bool {
	objectType := reflect.TypeOf(object)
	expectedType := reflect.TypeOf(expected)

	if objectType != expectedType {
		return failTest(t, 1, fmt.Sprintf("IsType: expected to be of type `%s`, actual was `%s`", expectedType.String(), objectType.String()))
	}

	return true
}

// Panic asserts that the code inside the specified function panics.
func Panic(t *testing.T, f func()) bool {
	isPanic, _ := didPanic(f)
	if !isPanic {
		return failTest(t, 1, fmt.Sprintf("Panic: function (%p) is expected to panic, actual does not panic", f))
	}

	return true
}

// NotPanic asserts that the code inside the specified function does not panic.
func NotPanic(t *testing.T, f func()) bool {
	isPanic, value := didPanic(f)
	if isPanic {
		return failTest(t, 1, fmt.Sprintf("NotPanic: function (%p) is expected not to panic, acutal panic with `%v`", f, value))
	}

	return true
}

// PanicWithValue asserts that the code inside the specified function panics, and the recovered value equals the expected value.
func PanicWithValue(t *testing.T, expected interface{}, f func()) bool {
	isPanic, value := didPanic(f)
	if !isPanic {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: function (%p) is expected to panic with `%#v`, actual does not panic", f, expected))
	}

	if !IsObjectEqual(value, expected) {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: function (%p) is expected to panic with `%#v`, actual with `%#v`", f, expected, value))
	}

	return true
}

// InDelta asserts that the two numerals are within delta of each other.
func InDelta(t *testing.T, expected, actual interface{}, eps float64) bool {
	in, actualEps, err := calcDeltaInEps(expected, actual, eps)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("InDelta: invalid operation (%v)", err))
	}

	if !in {
		return failTest(t, 1, fmt.Sprintf("InDelta: max difference between `%#v` and `%#v` allowed is `%#v`, but difference was `%#v`", expected, actual, eps, actualEps))
	}

	return true
}

// NotInDelta asserts that the two numerals are not within delta of each other.
func NotInDelta(t *testing.T, expected, actual interface{}, eps float64) bool {
	in, actualEps, err := calcDeltaInEps(expected, actual, eps)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: invalid operation (%v)", err))
	}

	if in {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: max difference between `%#v` and `%#v` is not allowed in `%#v`, but difference was `%#v`", expected, actual, eps, actualEps))
	}

	return true
}
