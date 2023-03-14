package xtesting

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"os"
	"reflect"
	"testing"
	"unicode"
)

// =================
// testing functions
// =================

// Equal asserts that two objects are deep equal.
func Equal(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreNotFunc(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if !reflect.DeepEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("Equal: expect to be `%#v`, but actually was `%#v`", want, give), msgAndArgs...)
	}

	return true
}

// NotEqual asserts that the specified values are not deep equal.
func NotEqual(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreNotFunc(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotEqual: invalid operation `%#v` != `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if reflect.DeepEqual(give, want) {
		return failTest(t, 1, fmt.Sprintf("NotEqual: expect not to be `%#v`, but actually equaled", want), msgAndArgs...)
	}

	return true
}

// EqualValue asserts that two objects are equal or convertible to the same types and equal.
func EqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreNotFunc(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("EqualValue: invalid operation `%#v` == `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if !xreflect.DeepEqualInValue(give, want) {
		return failTest(t, 1, fmt.Sprintf("EqualValue: expect to be `%#v`, but actually was `%#v`", want, give), msgAndArgs...)
	}

	return true
}

// NotEqualValue asserts that two objects are not equal even when converted to the same type.
func NotEqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreNotFunc(give, want); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotEqualValue: invalid operation `%#v` != `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if xreflect.DeepEqualInValue(give, want) {
		return failTest(t, 1, fmt.Sprintf("NotEqualValue: expect not to be `%#v`, but actually equaled", want), msgAndArgs...)
	}

	return true
}

// SamePointer asserts that two pointers have the same pointer type, and point to the same address.
func SamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameKind(give, want, reflect.Ptr); err != nil {
		return failTest(t, 1, fmt.Sprintf("SamePointer: invalid operation `%#v` == `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	giveType, wantType := reflect.TypeOf(give), reflect.TypeOf(want)
	if giveType != wantType {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expect to have the same pointer type, but actually differ (%T and %T)", want, give), msgAndArgs...)
	}
	if give != want {
		return failTest(t, 1, fmt.Sprintf("SamePointer: expect to point to `%#v`, but actually pointed to `%#v`", want, give), msgAndArgs...)
	}

	return true
}

// NotSamePointer asserts that two pointers have different pointer types, or do not point to the same address.
func NotSamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameKind(give, want, reflect.Ptr); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotSamePointer: invalid operation `%#v` != `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	giveType, wantType := reflect.TypeOf(give), reflect.TypeOf(want)
	if giveType == wantType && give == want {
		return failTest(t, 1, fmt.Sprintf("NotSamePointer: expect not to point to `%#v` with `%T` type, but actually pointed", want, want), msgAndArgs...)
	}

	return true
}

// SameFunction asserts that types and underlying pointers of two functions are the same.
func SameFunction(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameKind(give, want, reflect.Func); err != nil {
		return failTest(t, 1, fmt.Sprintf("SameFunction: invalid operation `%#v` == `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if !xreflect.SameUnderlyingPointerWithTypeAndKind(give, want, reflect.Func) {
		return failTest(t, 1, fmt.Sprintf("SameFunction: expect to be `%p` (%T), but actually `%p` (%T)", want, want, give, give), msgAndArgs...)
	}

	return true
}

// NotSameFunction asserts that types and underlying pointers of two functions are not the same.
func NotSameFunction(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameKind(give, want, reflect.Func); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotSameFunction: invalid operation `%#v` != `%#v` (%+v)", give, want, err), msgAndArgs...)
	}

	if xreflect.SameUnderlyingPointerWithTypeAndKind(give, want, reflect.Func) {
		return failTest(t, 1, fmt.Sprintf("NotSameFunction: expect not to be `%p` (%T), but actually the same", want, want), msgAndArgs...)
	}

	return true
}

// True asserts that the specified value is true.
func True(t testing.TB, value bool, msgAndArgs ...interface{}) bool {
	if !value {
		return failTest(t, 1, fmt.Sprintf("True: expect to be `true`, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// False asserts that the specified value is false.
func False(t testing.TB, value bool, msgAndArgs ...interface{}) bool {
	if value {
		return failTest(t, 1, fmt.Sprintf("False: expect to be `false`, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// Nil asserts that the specified object is nil.
func Nil(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if !xreflect.IsNilValue(value) {
		return failTest(t, 1, fmt.Sprintf("Nil: expect to be `<nil>`, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// NotNil asserts that the specified object is not nil.
func NotNil(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if xreflect.IsNilValue(value) {
		return failTest(t, 1, fmt.Sprintf("NotNil: expect not to be `<nil>`, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// Zero asserts that the specified object is the zero value for its type.
func Zero(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if !xreflect.IsZeroValue(value) {
		return failTest(t, 1, fmt.Sprintf("Zero: expect to be zero value, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// NotZero asserts that the specified object is not the zero value for its type.
func NotZero(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if xreflect.IsZeroValue(value) {
		return failTest(t, 1, fmt.Sprintf("NotZero: expect not to be zero value, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// BlankString asserts that the specified object is an empty or black string.
func BlankString(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	s, ok := value.(string)
	if !ok {
		return failTest(t, 1, fmt.Sprintf("BlankString: expect string, got `%#v` (%T)", value, value), msgAndArgs...)
	}

	for _, r := range s {
		if !unicode.IsSpace(r) {
			return failTest(t, 1, fmt.Sprintf("BlankString: expect to be blank string, but actually was `%#v`", value), msgAndArgs...)
		}
	}

	return true
}

// NotBlankString asserts that the specified object is not an empty or black string.
func NotBlankString(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	s, ok := value.(string)
	if !ok {
		return failTest(t, 1, fmt.Sprintf("NotBlankString: expect string, got `%#v` (%T)", value, value), msgAndArgs...)
	}

	hasNotBlank := false
	for _, r := range s {
		if !unicode.IsSpace(r) {
			hasNotBlank = true
			break
		}
	}
	if !hasNotBlank {
		return failTest(t, 1, fmt.Sprintf("NotBlankString: expect not to be blank string, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// EmptyCollection asserts that the specified object is empty collection value.
func EmptyCollection(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if !xreflect.IsEmptyCollection(value) {
		return failTest(t, 1, fmt.Sprintf("EmptyCollection: expect to be empty collection value, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// NotEmptyCollection asserts that the specified object is not empty collection value.
func NotEmptyCollection(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool {
	if xreflect.IsEmptyCollection(value) {
		return failTest(t, 1, fmt.Sprintf("NotEmptyCollection: expect not to be empty collection value, but actually was `%#v`", value), msgAndArgs...)
	}

	return true
}

// Error asserts that a function returned an error.
func Error(t testing.TB, err error, msgAndArgs ...interface{}) bool {
	if err == nil {
		return failTest(t, 1, fmt.Sprintf("Error: expect not to be nil error, but actually was `%#v`", err), msgAndArgs...)
	}

	return true
}

// NilError asserts that a function returned no error.
func NilError(t testing.TB, err error, msgAndArgs ...interface{}) bool {
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NilError: expect to be nil error, but actually was `%#v`", err), msgAndArgs...)
	}

	return true
}

// EqualError asserts that a function returned an error and that it is equal to the provided error.
func EqualError(t testing.TB, err error, wantString string, msgAndArgs ...interface{}) bool {
	if err == nil {
		return failTest(t, 1, fmt.Sprintf("EqualError: expect not to be nil error, but actually was `%#v`", err), msgAndArgs...)
	}

	if msg := err.Error(); msg != wantString {
		return failTest(t, 1, fmt.Sprintf("EqualError: expect to be error with message `%#v`, but actually with `%#v`", wantString, msg), msgAndArgs...)
	}

	return true
}

// NotEqualError asserts that a function returned an error and that it is not equal to the provided error.
func NotEqualError(t testing.TB, err error, wantString string, msgAndArgs ...interface{}) bool {
	if err == nil {
		return failTest(t, 1, fmt.Sprintf("NotEqualError: expect not to be nil error, but actually was `%#v`", err), msgAndArgs...)
	}

	if err.Error() == wantString {
		return failTest(t, 1, fmt.Sprintf("NotEqualError: expect error message not to be `%#v`, but actually equaled", wantString), msgAndArgs...)
	}

	return true
}

// MatchRegexp asserts that a specified regexp matches a string.
func MatchRegexp(t testing.TB, rx interface{}, str string, msgAndArgs ...interface{}) bool {
	match, re, err := matchRegexp(rx, str)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("MatchRegexp: invalid regular expression value `%#v` (%+v)", rx, err), msgAndArgs...)
	}

	if !match {
		return failTest(t, 1, fmt.Sprintf("MatchRegexp: expect `%#v` to match `%#v`, but actually not matched", str, re.String()), msgAndArgs...)
	}

	return true
}

// NotMatchRegexp asserts that a specified regexp does not match a string.
func NotMatchRegexp(t testing.TB, rx interface{}, str string, msgAndArgs ...interface{}) bool {
	match, re, err := matchRegexp(rx, str)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotMatchRegexp: invalid regular expression value `%#v` (%+v)", rx, err), msgAndArgs...)
	}

	if match {
		return failTest(t, 1, fmt.Sprintf("NotMatchRegexp: expect `%#v` not to match `%#v`, but actually matched", str, re.String()), msgAndArgs...)
	}

	return true
}

// InDelta asserts that the two numerics are within delta of each other.
func InDelta(t testing.TB, give, want interface{}, delta float64, msgAndArgs ...interface{}) bool {
	inDelta, actualDiff, err := calcDiffInDelta(give, want, delta)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("InDelta: invalid operation on `%#v`, `%#v` and `%#v` (%+v)", give, want, delta, err), msgAndArgs...)
	}

	if !inDelta {
		return failTest(t, 1, fmt.Sprintf("InDelta: expect difference between `%#v` and `%#v` to be less than or equal to `%#v`, but actually was `%#v`", give, want, delta, actualDiff), msgAndArgs...)
	}

	return true
}

// NotInDelta asserts that the two numerics are not within delta of each other.
func NotInDelta(t testing.TB, give, want interface{}, delta float64, msgAndArgs ...interface{}) bool {
	inDelta, actualDiff, err := calcDiffInDelta(give, want, delta)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: invalid operation on `%#v`, `%#v` and `%#v` (%+v)", err, give, want, delta), msgAndArgs...)
	}

	if inDelta {
		return failTest(t, 1, fmt.Sprintf("NotInDelta: expect difference between `%#v` and `%#v` to be greater than `%#v`, but actually was `%#v`", give, want, delta, actualDiff), msgAndArgs...)
	}

	return true
}

// InEpsilon asserts that two numerics have a relative error less than epsilon.
func InEpsilon(t testing.TB, give, want interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	inEps, actualRee, err := calcRelativeError(give, want, epsilon)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("InEpsilon: invalid operation on `%#v`, `%#v` and `%#v` (%+v)", err, give, want, epsilon), msgAndArgs...)
	}

	if !inEps {
		return failTest(t, 1, fmt.Sprintf("InEpsilon: expect relative error between `%#v` and `%#v` to be less than or equal to `%#v`, but actually was `%#v`", give, want, epsilon, actualRee), msgAndArgs...)
	}

	return true
}

// NotInEpsilon asserts that two numerics have a relative error greater than epsilon.
func NotInEpsilon(t testing.TB, give, want interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	inEps, actualRee, err := calcRelativeError(give, want, epsilon)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotInEpsilon: invalid operation on `%#v`, `%#v` and `%#v` (%+v)", err, give, want, epsilon), msgAndArgs...)
	}

	if inEps {
		return failTest(t, 1, fmt.Sprintf("NotInEpsilon: expect relative error between `%#v` and `%#v` to be greater than `%#v`, but actually was `%#v`", give, want, epsilon, actualRee), msgAndArgs...)
	}

	return true
}

// Contain asserts that the specified string, list(array, slice...) or map contains the specified substring or element.
func Contain(t testing.TB, container, value interface{}, msgAndArgs ...interface{}) bool {
	found, err := containElement(container, value)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("Contain: invalid operation on `%#v` and `%#v` (%+v)", container, value, err), msgAndArgs...)
	}

	if !found {
		return failTest(t, 1, fmt.Sprintf("Contain: expect `%#v` to contain `%#v`, but actually not contained", container, value), msgAndArgs...)
	}

	return true
}

// NotContain asserts that the specified string, list(array, slice...) or map does not contain the specified substring or element.
func NotContain(t testing.TB, container, value interface{}, msgAndArgs ...interface{}) bool {
	found, err := containElement(container, value)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("NotContain: invalid operation on `%#v` and `%#v` (%+v)", container, value, err), msgAndArgs...)
	}

	if found {
		return failTest(t, 1, fmt.Sprintf("NotContain: expect `%#v` not to contain `%#v`, but actually contained", container, value), msgAndArgs...)
	}

	return true
}

// Subset asserts that the specified list(array, slice...) contains all elements given in the specified subset(array, slice...).
func Subset(t testing.TB, list, subset interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameList(list, subset); err != nil {
		return failTest(t, 1, fmt.Sprintf("Subset: invalid operation on `%#v` and `%#v` (%+v)", list, subset, err), msgAndArgs...)
	}

	allFound, element := containAllElements(list, subset)
	if !allFound {
		return failTest(t, 1, fmt.Sprintf("Subset: expect `%#v` to contain `%#v`, but actually not contained", list, element), msgAndArgs...)
	}

	return true
}

// NotSubset asserts that the specified list(array, slice...) contains not all elements given in the specified subset(array, slice...).
func NotSubset(t testing.TB, list, subset interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameList(list, subset); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotSubset: invalid operation on `%#v` and `%#v` (%+v)", list, subset, err), msgAndArgs...)
	}

	allFound, _ := containAllElements(list, subset)
	if allFound {
		return failTest(t, 1, fmt.Sprintf("NotSubset: expect `%#v` not to be a subset of `%#v`, but actually was", subset, list), msgAndArgs...)
	}

	return true
}

// ElementMatch asserts that the specified listA(array, slice...) equals to specified listB(array, slice...) ignoring the order of the elements.
// If there are duplicate elements, the number of appearances of each of them in both lists should match.
func ElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameList(listA, listB); err != nil {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: invalid operation on `%#v` and `%#v` (%+v)", listA, listB, err), msgAndArgs...)
	}

	extraA, extraB := diffLists(listA, listB)
	if len(extraA) != 0 || len(extraB) != 0 {
		return failTest(t, 1, fmt.Sprintf("ElementMatch: expect `%#v` and `%#v` to match each other, but actually not matched", listA, listB), msgAndArgs...)
	}

	return true
}

// NotElementMatch asserts that the specified listA(array, slice...) does not equal to specified listB(array, slice...) ignoring the order of the elements.
func NotElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool {
	if err := validateArgsAreSameList(listA, listB); err != nil {
		return failTest(t, 1, fmt.Sprintf("NotElementMatch: invalid operation on `%#v` and `%#v` (%+v)", listA, listB, err), msgAndArgs...)
	}

	extraA, extraB := diffLists(listA, listB)
	if len(extraA) == 0 && len(extraB) == 0 {
		return failTest(t, 1, fmt.Sprintf("NotElementMatch: expect `%#v` and `%#v` not to match each other, but actually matched", listA, listB), msgAndArgs...)
	}

	return true
}

// SameType asserts that the specified objects are of the same type.
func SameType(t testing.TB, value, want interface{}, msgAndArgs ...interface{}) bool {
	valueType := reflect.TypeOf(value)
	wantType := reflect.TypeOf(want)

	if !reflect.DeepEqual(valueType, wantType) {
		return failTest(t, 1, fmt.Sprintf("SameType: expect `%#v` to be of type `%T`, but actually was `%T`", value, want, value), msgAndArgs...)
	}

	return true
}

// NotSameType asserts that the specified objects are of the different types.
func NotSameType(t testing.TB, value, want interface{}, msgAndArgs ...interface{}) bool {
	valueType := reflect.TypeOf(value)
	wantType := reflect.TypeOf(want)

	if reflect.DeepEqual(valueType, wantType) {
		return failTest(t, 1, fmt.Sprintf("NotSameType: expect `%#v` not to be of type `%T`, but actually was the same type", value, want), msgAndArgs...)
	}

	return true
}

// Implement asserts that an object implements the specified interface.
func Implement(t testing.TB, value, interfacePtr interface{}, msgAndArgs ...interface{}) bool {
	interfaceType, err := validateArgsForImplement(value, interfacePtr)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("Implement: invalid parameters (%+v)", err), msgAndArgs...)
	}

	if !reflect.TypeOf(value).Implements(interfaceType) {
		return failTest(t, 1, fmt.Sprintf("Implement: expect type `%T` to implement `%s`, but actually not implemented", value, interfaceType.String()), msgAndArgs...)
	}

	return true
}

// NotImplement asserts that an object does not implement the specified interface.
func NotImplement(t testing.TB, value, interfacePtr interface{}, msgAndArgs ...interface{}) bool {
	interfaceType, err := validateArgsForImplement(value, interfacePtr)
	if err != nil {
		return failTest(t, 1, fmt.Sprintf("Implement: invalid parameters (%+v)", err), msgAndArgs...)
	}

	if reflect.TypeOf(value).Implements(interfaceType) {
		return failTest(t, 1, fmt.Sprintf("Implement: expect type `%T` not to implement `%s`, but actually implemented", value, interfaceType.String()), msgAndArgs...)
	}

	return true
}

// Panic asserts that the code inside the specified function panics.
func Panic(t testing.TB, f func(), msgAndArgs ...interface{}) bool {
	funcDidPanic, _ := checkPanic(f)
	if !funcDidPanic {
		return failTest(t, 1, fmt.Sprintf("Panic: expect function `%#v` to panic, but actually did not panic", interface{}(f)), msgAndArgs...)
	}

	return true
}

// NotPanic asserts that the code inside the specified function does not panic.
func NotPanic(t testing.TB, f func(), msgAndArgs ...interface{}) bool {
	funcDidPanic, panicValue := checkPanic(f)
	if funcDidPanic {
		return failTest(t, 1, fmt.Sprintf("NotPanic: expect function `%#v` not to panic, but actually paniced with `%#v`", interface{}(f), panicValue), msgAndArgs...)
	}

	return true
}

// PanicWithValue asserts that the code inside the specified function panics, and that the recovered panic value equals the wanted panic value.
func PanicWithValue(t testing.TB, want interface{}, f func(), msgAndArgs ...interface{}) bool {
	funcDidPanic, panicValue := checkPanic(f)
	if !funcDidPanic {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: expect function `%#v` to panic, but actually did not panic", interface{}(f)), msgAndArgs...)
	}

	if !reflect.DeepEqual(panicValue, want) {
		return failTest(t, 1, fmt.Sprintf("PanicWithValue: expect function `%#v` to panic with `%#v`, but actually with `%#v`", interface{}(f), want, panicValue), msgAndArgs...)
	}

	return true
}

// PanicWithError asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.
func PanicWithError(t testing.TB, wantString string, f func(), msgAndArgs ...interface{}) bool {
	funcDidPanic, panicValue := checkPanic(f)
	if !funcDidPanic {
		return failTest(t, 1, fmt.Sprintf("PanicWithError: expect function `%#v` to panic, but actually did not panic", interface{}(f)), msgAndArgs...)
	}

	panicErr, ok := panicValue.(error)
	if !ok || panicErr.Error() != wantString {
		return failTest(t, 1, fmt.Sprintf("PanicWithError: expect function `%#v` to panic with error message `%#v`, but actually with `%#v`", interface{}(f), wantString, panicValue), msgAndArgs...)
	}

	return true
}

// FileExist checks whether a file exists in given path. It fails if the path points to a directory, or there is an error when checking whether it exists.
func FileExist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("FileExist: error when calling os.Stat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err != nil {
		return failTest(t, 1, fmt.Sprintf("FileExist: expect file `%s` to exist, but actually not existed", path), msgAndArgs...)
	}
	if info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("FileExist: expect `%s` to be a file, but actually was a directory", path), msgAndArgs...)
	}

	return true
}

// FileNotExist checks whether a file does not exist in given path. It fails if the path points to an existing file only, or there is an error when checking whether it exists.
func FileNotExist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("FileNotExist: error when calling os.Stat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err == nil && !info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("FileNotExist: expect file `%s` not to exist, but actually was an existing file", path), msgAndArgs...)
	}

	return true
}

// FileLexist checks whether a file lexists in given path. It fails if the path points to a directory, or there is an error when checking whether it exists.
func FileLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("FileLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err != nil {
		return failTest(t, 1, fmt.Sprintf("FileLexist: expect file `%s` to exist, but actually not existed", path), msgAndArgs...)
	}
	if info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("FileLexist: expect `%s` to be a file, but actually was a directory", path), msgAndArgs...)
	}

	return true
}

// FileNotLexist checks whether a file does not lexist in given path. It fails if the path points to an existing file only, or there is an error when checking whether it exists.
func FileNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("FileNotLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err == nil && !info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("FileNotLexist: expect file `%s` not to exist, but actually was an existing file", path), msgAndArgs...)
	}

	return true
}

// DirExist checks whether a directory exists in given path. It fails if the path is a file rather a directory, or there is an error checking whether it exists.
func DirExist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("DirExist: error when calling os.Stat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err != nil {
		return failTest(t, 1, fmt.Sprintf("DirExist: expect directory `%s` to exist, but actually not existed", path), msgAndArgs...)
	}
	if !info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("DirExist: expect `%s` to be a directory, but actually was a file", path), msgAndArgs...)
	}

	return true
}

// DirNotExist checks whether a directory does not exist in given path. It fails if the path points to an existing directory only, or there is an error when checking whether it exists.
func DirNotExist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("DirNotExist: error when calling os.Stat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err == nil && info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("DirNotExist: expect directory `%s` not to exist, but actually was an existing directory", path), msgAndArgs...)
	}

	return true
}

// DirLexist checks whether a directory lexists in given path. It fails if the path is a file rather a directory, or there is an error checking whether it exists.
func DirLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("DirLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err != nil {
		return failTest(t, 1, fmt.Sprintf("DirLexist: expect directory `%s` to exist, but actually not existed", path), msgAndArgs...)
	}
	if !info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("DirLexist: expect `%s` to be a directory, but actually was a file", path), msgAndArgs...)
	}

	return true
}

// DirNotLexist checks whether a directory does not lexist in given path. It fails if the path points to an existing directory only, or there is an error when checking whether it exists.
func DirNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("DirNotLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err == nil && info.IsDir() {
		return failTest(t, 1, fmt.Sprintf("DirNotLexist: expect directory `%s` not to exist, but actually was an existing directory", path), msgAndArgs...)
	}

	return true
}

// SymlinkLexist checks whether a symlink lexists in given path. It fails if the path does not point to an existing symlink, or there is an error checking whether it exists.
func SymlinkLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("SymlinkLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err != nil {
		return failTest(t, 1, fmt.Sprintf("SymlinkLexist: expect symlink `%s` to exist, but actually not existed", path), msgAndArgs...)
	}
	if (info.Mode() & os.ModeSymlink) == 0 {
		return failTest(t, 1, fmt.Sprintf("SymlinkLexist: expect `%s` to be a symlink, but actually was an existing file or directory", path), msgAndArgs...)
	}

	return true
}

// SymlinkNotLexist checks whether a symlink does not lexist in given path. It fails if the path points to an existing symlink only, or there is an error when checking whether it exist.
func SymlinkNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool {
	info, err := os.Lstat(path)
	if err != nil && !os.IsNotExist(err) {
		return failTest(t, 1, fmt.Sprintf("SymlinkNotLexist: error when calling os.Lstat on `%#v` (%+v)", path, err), msgAndArgs...)
	}

	if err == nil && (info.Mode()&os.ModeSymlink) == os.ModeSymlink {
		return failTest(t, 1, fmt.Sprintf("SymlinkNotLexist: expect symlink `%s` not to exist, but actually was an existing file or directory", path), msgAndArgs...)
	}

	return true
}
