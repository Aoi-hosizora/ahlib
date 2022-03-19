package xtesting

import (
	"fmt"
	"math"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func fail(t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%d Failed <<<\n", path.Base(file), line)
	t.Fail()
}

func TestAssert(t *testing.T) {
	Assert(true, "test %s", "test")
	func() {
		defer func() {
			err := recover()
			if err != "test test" {
				fail(t)
			}
		}()
		Assert(false, "test %s", "test")
	}()
}

func TestEqual(t *testing.T) {
	mockT := &testing.T{}

	// Equal
	if !Equal(mockT, "Hello World", "Hello World") {
		fail(t)
	}
	if !Equal(mockT, 123, 123) {
		fail(t)
	}
	if !Equal(mockT, 123.5, 123.5) {
		fail(t)
	}
	if !Equal(mockT, []byte("Hello World"), []byte("Hello World")) {
		fail(t)
	}
	if !Equal(mockT, nil, nil) {
		fail(t)
	}
	if !Equal(mockT, int32(123), int32(123)) {
		fail(t)
	}
	if !Equal(mockT, uint64(123), uint64(123)) {
		fail(t)
	}
	type myType string
	if !Equal(mockT, myType("1"), myType("1")) {
		fail(t)
	}
	if !Equal(mockT, &struct{}{}, &struct{}{}) {
		fail(t)
	}
	var m map[string]interface{}
	if Equal(mockT, m["bar"], "something") {
		fail(t)
	}
	if Equal(mockT, myType("1"), myType("2")) {
		fail(t)
	}
	if Equal(mockT, 10, uint(10)) {
		fail(t)
	}
	if Equal(mockT, func() {}, func() {}) {
		fail(t)
	}

	// NotEqual
	if !NotEqual(mockT, "Hello World", "Hello World!") {
		fail(t)
	}
	if !NotEqual(mockT, 123, 1234) {
		fail(t)
	}
	if !NotEqual(mockT, 123.5, 123.55) {
		fail(t)
	}
	if !NotEqual(mockT, []byte("Hello World"), []byte("Hello World!")) {
		fail(t)
	}
	funcA := func() int { return 23 }
	funcB := func() int { return 42 }
	if NotEqual(mockT, funcA, funcB) {
		fail(t)
	}
	if NotEqual(mockT, nil, nil) {
		fail(t)
	}
	if NotEqual(mockT, "Hello World", "Hello World") {
		fail(t)
	}
	if NotEqual(mockT, 123, 123) {
		fail(t)
	}
	if NotEqual(mockT, 123.5, 123.5) {
		fail(t)
	}
	if NotEqual(mockT, []byte("Hello World"), []byte("Hello World")) {
		fail(t)
	}
	if NotEqual(mockT, &struct{}{}, &struct{}{}) {
		fail(t)
	}
	if !NotEqual(mockT, 10, uint(10)) {
		fail(t)
	}
}

func TestEqualValues(t *testing.T) {
	mockT := &testing.T{}

	// EqualValue
	if EqualValue(mockT, "Hello World", "Hello World!") {
		fail(t)
	}
	if EqualValue(mockT, 123, 1234) {
		fail(t)
	}
	if EqualValue(mockT, 123.5, 123.55) {
		fail(t)
	}
	if EqualValue(mockT, []byte("Hello World"), []byte("Hello World!")) {
		fail(t)
	}
	if !EqualValue(mockT, nil, nil) {
		fail(t)
	}
	if !EqualValue(mockT, "Hello World", "Hello World") {
		fail(t)
	}
	if !EqualValue(mockT, 123, 123) {
		fail(t)
	}
	if !EqualValue(mockT, 123.5, 123.5) {
		fail(t)
	}
	if !EqualValue(mockT, []byte("Hello World"), []byte("Hello World")) {
		fail(t)
	}
	if !EqualValue(mockT, &struct{}{}, &struct{}{}) {
		fail(t)
	}

	// NotEqualValue
	if !NotEqualValue(mockT, "Hello World", "Hello World!") {
		fail(t)
	}
	if !NotEqualValue(mockT, 123, 1234) {
		fail(t)
	}
	if !NotEqualValue(mockT, 123.5, 123.55) {
		fail(t)
	}
	if !NotEqualValue(mockT, []byte("Hello World"), []byte("Hello World!")) {
		fail(t)
	}
	if NotEqualValue(mockT, nil, nil) {
		fail(t)
	}
	if NotEqualValue(mockT, "Hello World", "Hello World") {
		fail(t)
	}
	if NotEqualValue(mockT, 123, 123) {
		fail(t)
	}
	if NotEqualValue(mockT, 123.5, 123.5) {
		fail(t)
	}
	if NotEqualValue(mockT, []byte("Hello World"), []byte("Hello World")) {
		fail(t)
	}
	if NotEqualValue(mockT, &struct{}{}, &struct{}{}) {
		fail(t)
	}
	funcA := func() int { return 23 }
	funcB := func() int { return 42 }
	if !NotEqualValue(mockT, funcA, funcB) {
		fail(t)
	}
	if !NotEqualValue(mockT, 10, 11) {
		fail(t)
	}
	if NotEqualValue(mockT, 10, uint(10)) {
		fail(t)
	}
	if NotEqualValue(mockT, struct{}{}, struct{}{}) {
		fail(t)
	}
}

func TestSamePointer(t *testing.T) {
	mockT := &testing.T{}
	ptr := func(i int) *int {
		return &i
	}
	ptr2 := func(i int32) *int32 {
		return &i
	}
	p := ptr(2)
	p2 := ptr2(2)

	// SamePointer
	if SamePointer(mockT, ptr(1), ptr(1)) {
		fail(t)
	}
	if SamePointer(mockT, 1, 1) {
		fail(t)
	}
	if SamePointer(mockT, p, *p) {
		fail(t)
	}
	if !SamePointer(mockT, p, p) {
		fail(t)
	}
	if SamePointer(mockT, p, p2) {
		fail(t)
	}

	// NotSamePointer
	if !NotSamePointer(mockT, ptr(1), ptr(1)) {
		fail(t)
	}
	if !NotSamePointer(mockT, 1, 1) {
		fail(t)
	}
	if !NotSamePointer(mockT, p, *p) {
		fail(t)
	}
	if NotSamePointer(mockT, p, p) {
		fail(t)
	}
	if !NotSamePointer(mockT, p, p2) {
		fail(t)
	}
}

func TestTrueFalse(t *testing.T) {
	mockT := &testing.T{}

	// True
	if !True(mockT, true) {
		fail(t)
	}
	if True(mockT, false) {
		fail(t)
	}

	// False
	if !False(mockT, false) {
		fail(t)
	}
	if False(mockT, true) {
		fail(t)
	}
}

func TestNil(t *testing.T) {
	mockT := &testing.T{}

	// Nil
	if Nil(mockT, new(interface{})) {
		fail(t)
	}
	if !Nil(mockT, nil) {
		fail(t)
	}
	if !Nil(mockT, (*struct{})(nil)) {
		fail(t)
	}
	if Nil(mockT, "") {
		fail(t)
	}
	if Nil(mockT, 12) {
		fail(t)
	}
	if Nil(mockT, &struct{}{}) {
		fail(t)
	}
	if Nil(mockT, func() {}) {
		fail(t)
	}

	// NotNil
	if !NotNil(mockT, new(interface{})) {
		fail(t)
	}
	if NotNil(mockT, nil) {
		fail(t)
	}
	if NotNil(mockT, (*struct{})(nil)) {
		fail(t)
	}
	if !NotNil(mockT, "") {
		fail(t)
	}
	if !NotNil(mockT, 12) {
		fail(t)
	}
	if !NotNil(mockT, &struct{}{}) {
		fail(t)
	}
	if !NotNil(mockT, func() {}) {
		fail(t)
	}
}

func TestZero(t *testing.T) {
	mockT := &testing.T{}

	var i interface{}
	zeros := []interface{}{
		nil, "", false, complex64(0), complex128(0), float32(0), float64(0),
		0, int8(0), int16(0), int32(0), int64(0), byte(0), rune(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0),
		[0]interface{}{}, []interface{}(nil), struct{ x int }{}, (*interface{})(nil), (func())(nil), interface{}(nil),
		map[interface{}]interface{}(nil), (chan interface{})(nil), (<-chan interface{})(nil), (chan<- interface{})(nil),
	}

	nonZeros := []interface{}{
		"s", true, complex64(1), complex128(1), float32(1), float64(1),
		1, int8(1), int16(1), int32(1), int64(1), byte(1), rune(1),
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1), uintptr(1),
		[1]interface{}{1}, []interface{}{}, struct{ x int }{1}, &i, func() {}, interface{}(1),
		map[interface{}]interface{}{}, make(chan interface{}), (<-chan interface{})(make(chan interface{})), (chan<- interface{})(make(chan interface{})),
	}

	// Zero
	for _, zero := range zeros {
		if !Zero(mockT, zero) {
			fail(t)
		}
	}
	for _, nonZero := range nonZeros {
		if Zero(mockT, nonZero) {
			fail(t)
		}
	}

	// NotZero
	for _, zero := range zeros {
		if NotZero(mockT, zero) {
			fail(t)
		}
	}
	for _, nonZero := range nonZeros {
		if !NotZero(mockT, nonZero) {
			fail(t)
		}
	}
}

func TestEmptyCollection(t *testing.T) {
	mockT := &testing.T{}

	lenZeros := []interface{}{
		nil, "", [0]interface{}{}, []interface{}(nil), []interface{}{}, map[interface{}]interface{}(nil),
		map[interface{}]interface{}{}, (chan interface{})(nil), (<-chan interface{})(nil), (chan<- interface{})(nil),
	}

	nonLenZeros := []interface{}{
		"1", false, complex64(0), complex128(0), float32(0), float64(0),
		0, int8(0), int16(0), int32(0), int64(0), byte(0), rune(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0),
		[1]interface{}{1}, []interface{}{1}, map[interface{}]interface{}{1: 1},
		struct{ x int }{}, (*interface{})(nil), (func())(nil),
	}

	// ZeroLen
	for _, zero := range lenZeros {
		if !EmptyCollection(mockT, zero) {
			fail(t)
		}
	}
	for _, nonZero := range nonLenZeros {
		if EmptyCollection(mockT, nonZero) {
			fail(t)
		}
	}

	// NotZeroLen
	for _, zero := range lenZeros {
		if NotEmptyCollection(mockT, zero) {
			fail(t)
		}
	}
	for _, nonZero := range nonLenZeros {
		if !NotEmptyCollection(mockT, nonZero) {
			fail(t)
		}
	}
}

func TestContain(t *testing.T) {
	mockT := new(testing.T)
	type A struct {
		Name, Value string
	}
	list := []string{"Foo", "Bar"}
	complexList := []*A{
		{"b", "c"},
		{"d", "e"},
		{"g", "h"},
		{"j", "k"},
	}
	simpleMap := map[interface{}]interface{}{"Foo": "Bar"}

	// Contain
	if !Contain(mockT, "Hello World", "Hello") {
		fail(t)
	}
	if Contain(mockT, "Hello World", "Salut") {
		fail(t)
	}
	if !Contain(mockT, list, "Bar") {
		fail(t)
	}
	if Contain(mockT, list, "Salut") {
		fail(t)
	}
	if !Contain(mockT, complexList, &A{"g", "h"}) {
		fail(t)
	}
	if Contain(mockT, complexList, &A{"g", "e"}) {
		fail(t)
	}
	if Contain(mockT, complexList, &A{"g", "e"}) {
		fail(t)
	}
	if !Contain(mockT, simpleMap, "Foo") {
		fail(t)
	}
	if Contain(mockT, simpleMap, "Bar") {
		fail(t)
	}
	if Contain(mockT, func() {}, "Hello") {
		fail(t)
	}

	// NotContain
	if !NotContain(mockT, "Hello World", "Hello!") {
		fail(t)
	}
	if NotContain(mockT, "Hello World", "Hello") {
		fail(t)
	}
	if !NotContain(mockT, list, "Foo!") {
		fail(t)
	}
	if NotContain(mockT, list, "Foo") {
		fail(t)
	}
	if NotContain(mockT, simpleMap, "Foo") {
		fail(t)
	}
	if !NotContain(mockT, simpleMap, "Bar") {
		fail(t)
	}
	if NotContain(mockT, func() {}, "Hello") {
		fail(t)
	}
}

func TestElementMatch(t *testing.T) {
	mockT := &testing.T{}

	// ElementMatch
	if !ElementMatch(mockT, nil, nil) {
		fail(t)
	}
	if !ElementMatch(mockT, []int{}, []int{}) {
		fail(t)
	}
	if !ElementMatch(mockT, []int{1}, []int{1}) {
		fail(t)
	}
	if !ElementMatch(mockT, []int{1, 1}, []int{1, 1}) {
		fail(t)
	}
	if !ElementMatch(mockT, []int{1, 2}, []int{1, 2}) {
		fail(t)
	}
	if !ElementMatch(mockT, []int{1, 2}, []int{2, 1}) {
		fail(t)
	}
	if !ElementMatch(mockT, [2]int{1, 2}, [2]int{2, 1}) {
		fail(t)
	}
	if !ElementMatch(mockT, []string{"hello", "world"}, []string{"world", "hello"}) {
		fail(t)
	}
	if !ElementMatch(mockT, []string{"hello", "hello"}, []string{"hello", "hello"}) {
		fail(t)
	}
	if !ElementMatch(mockT, []string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}) {
		t.Error("ElementsMatch should return true")
	}
	if !ElementMatch(mockT, [3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}) {
		t.Error("ElementsMatch should return true")
	}
	if !ElementMatch(mockT, []int{}, nil) {
		fail(t)
	}
	if ElementMatch(mockT, []int{1}, []int{1, 1}) {
		fail(t)
	}
	if ElementMatch(mockT, []int{1, 2}, []int{2, 2}) {
		fail(t)
	}
	if ElementMatch(mockT, []string{"hello", "hello"}, []string{"hello"}) {
		fail(t)
	}
	if ElementMatch(mockT, []string{}, func() {}) {
		fail(t)
	}
	if ElementMatch(mockT, func() {}, []string{}) {
		fail(t)
	}
	if ElementMatch(mockT, func() {}, func() {}) {
		fail(t)
	}
}

func TestInDelta(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		a, b   interface{}
		delta  float64
		result bool
	}{
		{1.001, 1, 0.01, true},
		{1, 1.001, 0.01, true},
		{1, 2, 1, true},
		{1, 2, 0.5, false},
		{2, 1, 0.5, false},
		{"", nil, 1, false},
		{42, math.NaN(), 0.01, false},
		{math.NaN(), 42, 0.01, false},
		{uint(2), uint(1), 1, true},
		{uint8(2), uint8(1), 1, true},
		{uint16(2), uint16(1), 1, true},
		{uint32(2), uint32(1), 1, true},
		{uint64(2), uint64(1), 1, true},
		{2, 1, 1, true},
		{int8(2), int8(1), 1, true},
		{int16(2), int16(1), 1, true},
		{int32(2), int32(1), 1, true},
		{int64(2), int64(1), 1, true},
		{float32(2), float32(1), 1, true},
		{float64(2), float64(1), 1, true},
	} {
		if InDelta(mockT, tc.a, tc.b, tc.delta) != tc.result {
			fail(t)
		}
	}

	for _, tc := range []struct {
		a, b   interface{}
		delta  float64
		result bool
	}{
		{1.001, 1, 0.01, false},
		{1, 1.001, 0.01, false},
		{1, 2, 1, false},
		{1, 2, 0.5, true},
		{2, 1, 0.5, true},
		{"", nil, 1, false},
		{42, math.NaN(), 0.01, false},
		{math.NaN(), 42, 0.01, false},
		{uint(2), uint(1), 1, false},
		{uint8(2), uint8(1), 1, false},
		{uint16(2), uint16(1), 1, false},
		{uint32(2), uint32(1), 1, false},
		{uint64(2), uint64(1), 1, false},
		{2, 1, 1, false},
		{int8(2), int8(1), 1, false},
		{int16(2), int16(1), 1, false},
		{int32(2), int32(1), 1, false},
		{int64(2), int64(1), 1, false},
		{float32(2), float32(1), 1, false},
		{float64(2), float64(1), 1, false},
	} {
		if NotInDelta(mockT, tc.a, tc.b, tc.delta) != tc.result {
			fail(t)
		}
	}
}

func TestIsType(t *testing.T) {
	mockT := &testing.T{}

	if !IsType(mockT, &TypeStruct{}, &TypeStruct{}) {
		fail(t)
	}
	if IsType(mockT, &TypeStruct{}, &TypeStruct2{}) {
		fail(t)
	}
}

type TypeInterface interface {
	TestMethod()
}

type TypeStruct struct{}

func (a *TypeStruct) TestMethod() {}

type TypeStruct2 struct{}

func TestImplements(t *testing.T) {
	mockT := &testing.T{}

	if !Implement(mockT, &TypeStruct{}, (*TypeInterface)(nil)) {
		fail(t)
	}
	if Implement(mockT, &TypeStruct2{}, (*TypeInterface)(nil)) {
		fail(t)
	}
	if Implement(mockT, nil, (*TypeInterface)(nil)) {
		fail(t)
	}
}

func TestPanics(t *testing.T) {
	mockT := &testing.T{}

	if !NotPanic(mockT, func() {
	}) {
		fail(t)
	}
	if NotPanic(mockT, func() {
		panic("Panic!")
	}) {
		fail(t)
	}

	if !Panic(mockT, func() {
		panic("Panic!")
	}) {
		fail(t)
	}
	if Panic(mockT, func() {
	}) {
		fail(t)
	}

	if !PanicWithValue(mockT, "Panic!", func() {
		panic("Panic!")
	}) {
		fail(t)
	}
	if PanicWithValue(mockT, "Panic!", func() {
	}) {
		fail(t)
	}
	if PanicWithValue(mockT, "at the disco", func() {
		panic("Panic!")
	}) {
		fail(t)
	}
}

func TestMsgAndArgs(t *testing.T) {
	s := combineMsgAndArgs()
	if s != "" {
		fail(t)
	}

	s = combineMsgAndArgs("0")
	if s != "0" {
		fail(t)
	}

	s = combineMsgAndArgs([]int{1, 2})
	if s != "[1 2]" {
		fail(t)
	}

	s = combineMsgAndArgs(nil)
	if s != "<nil>" {
		fail(t)
	}

	s = combineMsgAndArgs("a%sc", "b")
	if s != "abc" {
		fail(t)
	}
}

type mockFinishFlagTestingT struct {
	testing.TB
	finished bool
}

func (m *mockFinishFlagTestingT) FailNow()             { m.finished = true }
func (m *mockFinishFlagTestingT) Fatal(...interface{}) { m.finished = true }

func TestOptions(t *testing.T) {
	mockT := &testing.T{}
	SetExtraSkip(1)
	if failTest(mockT, 0, "a", "") != false {
		fail(t)
	}
	if failTest(mockT, -1, "a", "%%a%s", "b") != false {
		fail(t)
	}
	SetExtraSkip(0)

	mockT2 := &mockFinishFlagTestingT{}
	UseFailNow(true)
	if failTest(mockT2, 1, "a", "") != false {
		fail(t)
	}
	if !mockT2.finished {
		fail(t)
	}

	mockT2 = &mockFinishFlagTestingT{}
	UseFailNow(false)
	if failTest(mockT, 1, "a", "") != false {
		fail(t)
	}
	if mockT2.finished {
		fail(t)
	}
}

func TestGoTool(t *testing.T) {
	defer _testGoToolFlag.Store(false)

	_testGoToolFlag.Store(false)
	p, err := GoCommand()
	if err != nil {
		fail(t)
	}
	if !strings.HasPrefix(p, filepath.Join(runtime.GOROOT(), "bin")) {
		fail(t)
	}

	_testGoToolFlag.Store(true)
	p, err = GoCommand()
	if err == nil {
		fail(t)
	}
	if p != "" {
		fail(t)
	}
}
