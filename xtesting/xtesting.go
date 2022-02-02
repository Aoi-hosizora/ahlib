package xtesting

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"sync/atomic"
	"testing"
)

// =========
// fail test
// =========

// failTest outputs the error message and fails the test.
func failTest(t testing.TB, skip int, msg string, msgAndArgs ...interface{}) bool {
	if skip < 0 {
		skip = 0
	}
	exSkip := int(atomic.LoadInt32(&_extraSkip))

	_, file, line, _ := runtime.Caller(skip + 1 + exSkip)
	m := fmt.Sprintf("%s:%d %s", path.Base(file), line, msg)
	if exMsg := combineMsgAndArgs(msgAndArgs...); len(exMsg) > 0 {
		m += exMsg
	}
	fmt.Println(m)

	if failNow := atomic.LoadInt32(&_useFailNow) == 1; !failNow {
		t.Fail()
	} else {
		t.FailNow()
	}
	return false
}

var (
	// _extraSkip is the extra skip, and this value cannot be less than zero, defaults to zero.
	_extraSkip int32 = 0

	// _useFailNow is a flag for using `FailNow` (if set to 1) rather than `Fail` (if set to 0), defaults to 0.
	_useFailNow int32 = 0
)

// SetExtraSkip sets extra skip for test functions, and it will be used when printing the test failed message, defaults to zero.
func SetExtraSkip(skip int32) {
	if skip >= 0 {
		atomic.StoreInt32(&_extraSkip, skip)
	}
}

// UseFailNow makes test functions to fail now when test failed, defaults to false, that means to use `Fail` rather than `FailNow`.
func UseFailNow(failNow bool) {
	if failNow {
		atomic.StoreInt32(&_useFailNow, 1)
	} else {
		atomic.StoreInt32(&_useFailNow, 0)
	}
}

// =================
// testing functions
// =================

// Equal asserts that two objects are deep equal.
func Equal(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateEqualArgs(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (%v)", give, want, err), msgAndArgs...)
	}

	if !IsObjectDeepEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("Equal: expected `%#v`, actual `%#v`", want, give), msgAndArgs...)
	}

	return true
}

// NotEqual asserts that the specified values are not deep equal.
func NotEqual(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateEqualArgs(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotEqual: invalid operation `%#v` != `%#v` (%v)", give, want, err), msgAndArgs...)
	}

	if IsObjectDeepEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("NotEqual: expected not to be `%#v`", want), msgAndArgs...)
	}

	return true
}

// EqualValue asserts that two objects are equal or convertible to the same types and equal.
func EqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if !IsObjectValueEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("EqualValue: expected `%#v`, actual `%#v`", want, give), msgAndArgs...)
	}

	return true
}

// NotEqualValue asserts that two objects are not equal even when converted to the same type.
func NotEqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if IsObjectValueEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("NotEqualValue: expected not to be `%#v`", want), msgAndArgs...)
	}

	return true
}

// SamePointer asserts that two pointers reference the same object.
func SamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if !IsPointerSame(give, want) {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expected `%#v` (%p), actual `%#v` (%p)", want, want, give, give), msgAndArgs...)
	}

	return true
}

// NotSamePointer asserts that two pointers do not reference the same object.
func NotSamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if IsPointerSame(give, want) {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expected not be `%#v` (%p)", want, want), msgAndArgs...)
	}

	return true
}

// True asserts that the specified value is true.
func True(t testing.TB, value bool, msgAndArgs ...interface{}) bool {
	if !value {
		return failTest(t, 1, fmt.Sprintf("True: expected `true`, actual `%#v`", value), msgAndArgs...)
	}

	return true
}

// False asserts that the specified value is false.
func False(t testing.TB, value bool, msgAndArgs ...interface{}) bool {
	if value {
		return failTest(t, 1, fmt.Sprintf("False: expected to be `false`, actual `%#v`", value), msgAndArgs...)
	}

	return true
}

// Nil asserts that the specified object is nil.
func Nil(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if !IsObjectNil(object) {
		return failTest(t, 1, fmt.Sprintf("Nil: expected `nil`, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// NotNil asserts that the specified object is not nil.
func NotNil(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if IsObjectNil(object) {
		return failTest(t, 1, fmt.Sprintf("NotNil, expected not to be `nil`, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// Zero asserts that the specified object is the zero value for its type.
func Zero(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if !IsObjectZero(object) {
		return failTest(t, 1, fmt.Sprintf("Zero: expected to be zero value, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// NotZero asserts that the specified object is not the zero value for its type.
func NotZero(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if IsObjectZero(object) {
		return failTest(t, 1, fmt.Sprintf("NotZero: expected not to be zero value, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// ZeroLen asserts that the length of specified object is zero.
func ZeroLen(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if !IsObjectZeroLen(object) {
		return failTest(t, 1, fmt.Sprintf("ZeroLen: expected to be zero length, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// NotZeroLen asserts that the length of specified object is not zero.
func NotZeroLen(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool {
	if IsObjectZeroLen(object) {
		return failTest(t, 1, fmt.Sprintf("NotZeroLen: expected not to be zero length, actual `%#v`", object), msgAndArgs...)
	}

	return true
}

// InDelta asserts that the two numerals are within delta of each other.
func InDelta(t testing.TB, give, want interface{}, eps float64, msgAndArgs ...interface{}) bool {
	in, actualEps, err := calcDeltaInEps(give, want, eps)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("InDelta: invalid operation (%v)", err), msgAndArgs...)
	}

	if !in {
		return failTest(t, 1, fmt.Sprintf("InDelta: max difference between `%#v` and `%#v` allowed is `%#v`, but difference was `%#v`", give, want, eps, actualEps), msgAndArgs...)
	}

	return true
}

// NotInDelta asserts that the two numerals are not within delta of each other.
func NotInDelta(t testing.TB, give, want interface{}, eps float64, msgAndArgs ...interface{}) bool {
	in, actualEps, err := calcDeltaInEps(give, want, eps)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: invalid operation (%v)", err), msgAndArgs...)
	}

	if in {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: max difference between `%#v` and `%#v` is not allowed in `%#v`, but difference was `%#v`", give, want, eps, actualEps), msgAndArgs...)
	}

	return true
}

// Contain asserts that the specified container contains the specified substring or element.
// Support string, array, slice or map.
func Contain(t testing.TB, container, object interface{}, msgAndArgs ...interface{}) bool {
	ok, found := includeElement(container, object)

	if !ok {
		return failTest(t, 1, fmt.Sprintf("Contain: invalid operator len(`%#v`)", container), msgAndArgs...)
	}
	if !found {
		return failTest(t, 1, fmt.Sprintf("Contain: `%#v` is expected to contain `%#v`", container, object), msgAndArgs...)
	}

	return true
}

// NotContain asserts that the specified container does not contain the specified substring or element.
// Support string, array, slice or map.
func NotContain(t testing.TB, container, object interface{}, msgAndArgs ...interface{}) bool {
	ok, found := includeElement(container, object)

	if !ok {
		return failTest(t, 1, fmt.Sprintf("NotContain: invalid operator len(`%#v`)", container), msgAndArgs...)
	}
	if found {
		return failTest(t, 1, fmt.Sprintf("NotContain: `%#v` is expected not to contain `%#v`", container, object), msgAndArgs...)
	}

	return true
}

// ElementMatch asserts that the specified listA is equal to specified listB ignoring the order of the elements.
// If there are duplicate elements, the number of appearances of each of them in both lists should match.
func ElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool {
	if IsObjectZeroLen(listA) && IsObjectZeroLen(listB) {
		return true
	}

	if err := validateArgIsList(listA, listB); err != nil {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: invalid operator: `%#v` <-> `%#v` (%v)", listA, listB, err), msgAndArgs...)
	}

	extraA, extraB := diffLists(listA, listB)
	if len(extraA) != 0 || len(extraB) != 0 {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: `%#v` and `%#v` are expected to match each other", listA, listB), msgAndArgs...)
	}

	return true
}

// IsType asserts that the specified objects are of the same type.
func IsType(t testing.TB, object, want interface{}, msgAndArgs ...interface{}) bool {
	objectType := reflect.TypeOf(object)
	wantType := reflect.TypeOf(want)

	if objectType != wantType {
		return failTest(t, 1, fmt.Sprintf("IsType: expected to be of type `%s`, actual was `%s`", wantType.String(), objectType.String()), msgAndArgs...)
	}

	return true
}

// Implements asserts that an object is implemented by the specified interface.
func Implements(t testing.TB, object, interfacePtr interface{}, msgAndArgs ...interface{}) bool {
	interfaceType := reflect.TypeOf(interfacePtr).Elem()

	if object == nil {
		return failTest(t, 1, fmt.Sprintf("Implements: invalid operation for `nil`"), msgAndArgs...)
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return failTest(t, 1, fmt.Sprintf("Implements: %T expected to implement `%v`, actual not implment.", object, interfacePtr), msgAndArgs...)
	}

	return true
}

// Panic asserts that the code inside the specified function panics.
func Panic(t testing.TB, f func(), msgAndArgs ...interface{}) bool {
	isPanic, _ := didPanic(f)
	if !isPanic {
		return failTest(t, 1, fmt.Sprintf("Panic: function (%p) is expected to panic, actual does not panic", f), msgAndArgs...)
	}

	return true
}

// NotPanic asserts that the code inside the specified function does not panic.
func NotPanic(t testing.TB, f func(), msgAndArgs ...interface{}) bool {
	isPanic, value := didPanic(f)
	if isPanic {
		return failTest(t, 1, fmt.Sprintf("NotPanic: function (%p) is expected not to panic, acutal panic with `%v`", f, value), msgAndArgs...)
	}

	return true
}

// PanicWithValue asserts that the code inside the specified function panics, and the recovered value equals the expected value.
func PanicWithValue(t testing.TB, want interface{}, f func(), msgAndArgs ...interface{}) bool {
	isPanic, value := didPanic(f)
	if !isPanic {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: function (%p) is expected to panic with `%#v`, actual does not panic", f, want), msgAndArgs...)
	}

	if !IsObjectDeepEqual(value, want) {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: function (%p) is expected to panic with `%#v`, actual with `%#v`", f, want, value), msgAndArgs...)
	}

	return true
}

/*

// Exit asserts that the code inside the specified function exits.
func Exit(t testing.TB, f func(), msgAndArgs ...interface{}) bool {
	// 1. Create a temp code file, use exec.Command to run and get exit code => need to write code file manually
	// https://github.com/sirupsen/logrus/blob/master/alt_exit_test.go#L75
	// https://github.com/sirupsen/logrus/blob/master/alt_exit.go#L49
	// https://stackoverflow.com/questions/10385551/get-exit-code-go

	// 2. Use a stub function and replace os.Exit when test => need to replace all os.Exit and only for internal
	// https://github.com/uber-go/zap/blob/a68efdbdd15b7816de33cdbe7e6def2a559bdf64/internal/exit/exit.go#L44
	// https://github.com/uber-go/zap/blob/a68efdbdd1/zapcore/entry_test.go#L124
	// https://github.com/uber-go/zap/blob/a68efdbdd15b7816de33cdbe7e6def2a559bdf64/zapcore/entry.go#L236

	// 3. Use exec.Command and rerun the test with an argument => gracefullest and recommend
	// https://talks.golang.org/2014/testing.slide#23

	// 4. Replace os.Exit to other function (patch), and restore it later => unsafe when run os.Exec in concurrency and difficult
	// https://stackoverflow.com/questions/26225513/how-to-test-os-exit-scenarios-in-go
	// https://github.com/bouk/monkey/blob/master/monkey.go#L67
	// https://github.com/bouk/monkey/blob/master/monkey.go#L119

	return true
}

// ExitWithCode asserts that the code inside the specified function exits with a code which not equals the expected code.
func ExitWithCode(t testing.TB, code int, f func(), msgAndArgs ...interface{}) bool {
	return true
}

*/
