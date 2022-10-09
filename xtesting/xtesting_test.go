package xtesting

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func fail(t *testing.T) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("!!! testing on %s:%d is failed !!!\n", path.Base(file), line)
	t.Fail()
}

type testFlag uint8

const (
	positive testFlag = iota
	negative
	abnormal
)

/*
	// discard stderr
	defer func(stdout *os.File) {
		os.Stderr = stdout
	}(os.Stderr)
	os.Stderr = os.NewFile(uintptr(syscall.Stdin), os.DevNull)
*/

func TestEqualNotEqual(t *testing.T) {
	mockT := &testing.T{}

	type myType string
	var m map[string]interface{}

	for _, tc := range []struct {
		giveG, giveW interface{}
		want         testFlag
	}{
		// expect to Equal
		{"Hello World", "Hello World", positive},
		{123, 123, positive},
		{123.5, 123.5, positive},
		{[]byte("Hello World"), []byte("Hello World"), positive},
		{nil, nil, positive},
		{int32(123), int32(123), positive},
		{uint64(123), uint64(123), positive},
		{myType("1"), myType("1"), positive},
		{&struct{}{}, &struct{}{}, positive},

		// expect to NotEqual
		{"Hello World", "Hello World!", negative},
		{123, 1234, negative},
		{123.5, 123.55, negative},
		{[]byte("Hello World"), []byte("Hello World!"), negative},
		{nil, new(struct{}), negative},
		{10, uint(10), negative},
		{m["bar"], "something", negative},
		{myType("1"), myType("2"), negative},

		// expect to fail in all cases
		{func() {}, func() {}, abnormal},
		{func() int { return 23 }, func() int { return 42 }, abnormal},
	} {
		pos := Equal(mockT, tc.giveG, tc.giveW)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotEqual(mockT, tc.giveG, tc.giveW)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestEqualValueNotEqualValue(t *testing.T) {
	mockT := &testing.T{}

	type myType string

	for _, tc := range []struct {
		giveG, giveW interface{}
		want         testFlag
	}{
		// expect to EqualValue
		{nil, nil, positive},
		{"Hello World", "Hello World", positive},
		{123, 123, positive},
		{123.5, 123.5, positive},
		{[]byte("Hello World"), []byte("Hello World"), positive},
		{new(struct{}), new(struct{}), positive},
		{&struct{}{}, &struct{}{}, positive},
		{10, uint(10), positive},
		{struct{}{}, struct{}{}, positive},
		{myType("1"), "1", positive},

		// expect to NotEqualValue
		{"Hello World", "Hello World!", negative},
		{123, 1234, negative},
		{123.5, 123.55, negative},
		{[]byte("Hello World"), []byte("Hello World!"), negative},
		{myType("1"), myType("2"), negative},
		{"1", myType("2"), negative},

		// except to fail in all cases
		{func() {}, func() {}, abnormal},
		{func() int { return 23 }, func() int { return 42 }, abnormal},
	} {
		pos := EqualValue(mockT, tc.giveG, tc.giveW)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotEqualValue(mockT, tc.giveG, tc.giveW)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestSamePointer(t *testing.T) {
	mockT := &testing.T{}

	ptr := func(i int) *int { return &i }
	ptr2 := func(i int32) *int32 { return &i }
	p := ptr(2)
	p2 := ptr2(2)

	for _, tc := range []struct {
		giveG, giveW interface{}
		want         testFlag
	}{
		// expect to SamePointer
		{p, p, positive},
		{p2, p2, positive},

		// expect to NotSamePointer
		{ptr(1), ptr(1), negative},
		{new(uint), new(uint), negative},
		{1, 1, negative},
		{p, *p, negative},
		{p2, *p2, negative},
		{p, nil, negative},
		{nil, p2, negative},
		{p, p2, negative},
	} {
		pos := SamePointer(mockT, tc.giveG, tc.giveW)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotSamePointer(mockT, tc.giveG, tc.giveW)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
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

func TestNilNotNil(t *testing.T) {
	mockT := &testing.T{}

	nils := []interface{}{
		nil, (*struct{})(nil), (*int)(nil), (func())(nil), fmt.Stringer(nil), error(nil),
		[]int(nil), map[int]int(nil), (chan int)(nil), (chan<- int)(nil), (<-chan int)(nil),
	}

	nonNils := []interface{}{
		0, "", &struct{}{}, new(interface{}), new(int), func() {}, fmt.Stringer(&strings.Builder{}), errors.New(""),
		[]int{}, map[int]int{}, make(chan int), make(chan<- int), make(<-chan int),
	}

	// expect to Nil
	for _, nil_ := range nils {
		if !Nil(mockT, nil_) {
			fail(t)
		}
		if NotNil(mockT, nil_) {
			fail(t)
		}
	}

	// expect to NotNil
	for _, nonNil := range nonNils {
		if !NotNil(mockT, nonNil) {
			fail(t)
		}
		if Nil(mockT, nonNil) {
			fail(t)
		}
	}
}

func TestZeroNotZero(t *testing.T) {
	mockT := &testing.T{}

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
		[1]interface{}{1}, []interface{}{}, struct{ x int }{1}, new(interface{}), func() {}, interface{}(1),
		map[interface{}]interface{}{}, make(chan interface{}), (<-chan interface{})(make(chan interface{})), (chan<- interface{})(make(chan interface{})),
	}

	// expect to Zero
	for _, zero := range zeros {
		if !Zero(mockT, zero) {
			fail(t)
		}
		if NotZero(mockT, zero) {
			fail(t)
		}
	}

	// expect to NotZero
	for _, nonZero := range nonZeros {
		if !NotZero(mockT, nonZero) {
			fail(t)
		}
		if Zero(mockT, nonZero) {
			fail(t)
		}
	}
}

func TestBlankNotBlankString(t *testing.T) {
	mockT := &testing.T{}

	errs := []interface{}{
		nil, 1, uint64(0), 0.01, false, [0]interface{}{}, []interface{}(nil),
		struct{}{}, struct{ x int }{}, [0]string{}, []string{""}, map[string]string{"": ""},
	}

	blanks := []interface{}{
		"", " ", "\t", "\n", "\r", "\v", "\f", "　", string([]rune{0x85}), string([]rune{0xA0}), "  ", "\t 　", "\r\n", "\v\f",
		"\u2000", "\u2001", "\u2002", "\u2003", "\u2004", "\u2005", "\u2006", "\u2007",
		"\u2008", "\u2009", "\u200A", "\u2028", "\u2029", "\u202F", "\u205F", "\u3000",
	}

	nonBlanks := []interface{}{
		"1", "\n \t1", "测试", "テスト", "test", "　　　　1", "\u2000@\u205F",
	}

	// expect to error
	for _, err := range errs {
		if BlankString(mockT, err) {
			fail(t)
		}
		if NotBlankString(mockT, err) {
			fail(t)
		}
	}

	for _, zero := range blanks {
		if !BlankString(mockT, zero) {
			fail(t)
		}
		if NotBlankString(mockT, zero) {
			fail(t)
		}
	}

	// expect to NotBlankString
	for _, nonBlank := range nonBlanks {
		if !NotBlankString(mockT, nonBlank) {
			fail(t)
		}
		if BlankString(mockT, nonBlank) {
			fail(t)
		}
	}
}

func TestEmptyCollectionNotEmptyCollection(t *testing.T) {
	mockT := &testing.T{}

	empties := []interface{}{
		"", [0]interface{}{}, []interface{}(nil), []interface{}{}, map[interface{}]interface{}(nil),
		map[interface{}]interface{}{}, (chan interface{})(nil), (<-chan interface{})(nil), (chan<- interface{})(nil),
	}

	nonEmpties := []interface{}{
		nil, "1", false, complex64(0), complex128(0), float32(0), float64(0),
		0, int8(0), int16(0), int32(0), int64(0), byte(0), rune(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0),
		[1]interface{}{1}, []interface{}{1}, map[interface{}]interface{}{1: 1},
		struct{ x int }{}, (*interface{})(nil), (func())(nil),
	}

	// expect to EmptyCollection
	for _, empty := range empties {
		if !EmptyCollection(mockT, empty) {
			fail(t)
		}
		if NotEmptyCollection(mockT, empty) {
			fail(t)
		}
	}

	// expect to NotEmptyCollection
	for _, nonEmpty := range nonEmpties {
		if !NotEmptyCollection(mockT, nonEmpty) {
			fail(t)
		}
		if EmptyCollection(mockT, nonEmpty) {
			fail(t)
		}
	}
}

type customError struct{}

func (c *customError) Error() string {
	if c == nil {
		return "customError (nil)"
	}
	return "customError"
}

func TestErrorNilError(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		give error
		want testFlag
	}{
		// expect to Error
		{errors.New("some error"), positive},
		{func() error { return &customError{} }(), positive},
		{func() error { return (*customError)(nil) }(), positive},

		// expect to NilError
		{nil, negative},
		{func() error { return nil }(), negative},
	} {
		pos := Error(mockT, tc.give)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NilError(mockT, tc.give)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestEqualErrorNotEqualError(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveE error
		giveS string
		want  testFlag
	}{
		// expect to EqualError
		{errors.New("some error"), "some error", positive},
		{func() error { return &customError{} }(), "customError", positive},
		{func() error { return (*customError)(nil) }(), "customError (nil)", positive},

		// expect to NotEqualError
		{errors.New("some error"), "some errors", negative},
		{func() error { return (*customError)(nil) }(), "customError", negative},

		// expect to fail in all cases
		{nil, "", abnormal},
		{func() error { return nil }(), "", abnormal},
	} {
		pos := EqualError(mockT, tc.giveE, tc.giveS)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotEqualError(mockT, tc.giveE, tc.giveS)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestMatchRegexpNotMatchRegexp(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveR interface{}
		giveS string
		want  testFlag
	}{
		// expect to MatchRegexp
		{"^start", "start of the line", positive},
		{"end$", "in the end", positive},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12.34", positive},
		{regexp.MustCompile("^start"), "start of the line", positive},
		{regexp.MustCompile("end$"), "in the end", positive},
		{regexp.MustCompile("[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}"), "My phone number is 650.12.34", positive},

		// expect to NotMatchRegexp
		{"^asdfastart", "Not the start of the line", negative},
		{"end$", "in the end.", negative},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12a.34", negative},
		{regexp.MustCompile("^asdfastart"), "Not the start of the line", negative},
		{regexp.MustCompile("end$"), "in the end.", negative},
		{regexp.MustCompile("[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}"), "My phone number is 650.12a.34", negative},

		// expect to fail in all cases
		{0, "start of the line", abnormal},
		{regexp.Regexp{}, "-", abnormal},
		{"end[$", "in the end", abnormal},
	} {
		pos := MatchRegexp(mockT, tc.giveR, tc.giveS)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotMatchRegexp(mockT, tc.giveR, tc.giveS)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestInDeltaNotInDelta(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveA, giveB interface{}
		giveDelta    float64
		want         testFlag
	}{
		// expect to InDelta
		{1.001, 1, 0.01, positive},
		{1, 1.001, 0.01, positive},
		{1, 2, 1, positive},
		{2, 1, 1, positive},
		{uint(2), uint(1), 1, positive},
		{uint8(2), uint8(1), 1, positive},
		{uint16(2), uint16(1), 1, positive},
		{uint32(2), uint32(1), 1, positive},
		{uint64(2), uint64(1), 1, positive},
		{int8(2), int8(1), 1, positive},
		{int16(2), int16(1), 1, positive},
		{int32(2), int32(1), 1, positive},
		{int64(2), int64(1), 1, positive},
		{float32(2), float32(1), 1, positive},
		{float64(2), float64(1), 1, positive},

		// expect to NotInDelta
		{1.001, 1, 0.0001, negative},
		{1, 1.001, 0.0001, negative},
		{1, 2, 0.5, negative},
		{2, 1, 0.5, negative},
		{uint(2), uint(1), 0.5, negative},
		{uint8(2), uint8(1), 0.5, negative},
		{uint16(2), uint16(1), 0.5, negative},
		{uint32(2), uint32(1), 0.5, negative},
		{uint64(2), uint64(1), 0.5, negative},
		{int8(2), int8(1), 0.5, negative},
		{int16(2), int16(1), 0.5, negative},
		{int32(2), int32(1), 0.5, negative},
		{int64(2), int64(1), 0.5, negative},
		{float32(2), float32(1), 0.5, negative},
		{float64(2), float64(1), 0.5, negative},

		// expect to fail in all cases
		{"1", 1, 1, abnormal},
		{1, nil, 1, abnormal},
		{nil, "1", 1, abnormal},
		{42, math.NaN(), 0.01, abnormal},
		{math.NaN(), 42, 0.01, abnormal},
		{42, 42, math.NaN(), abnormal},
	} {
		pos := InDelta(mockT, tc.giveA, tc.giveB, tc.giveDelta)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotInDelta(mockT, tc.giveA, tc.giveB, tc.giveDelta)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestInEpsilonNotInEpsilon(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveA, giveB interface{}
		giveDelta    float64
		want         testFlag
	}{
		// expect to InEpsilon
		{uint8(2), uint16(2), .001, positive},
		{2.1, 2.2, 0.1, positive},
		{2.2, 2.1, 0.1, positive},
		{-2.1, -2.2, 0.1, positive},
		{-2.2, -2.1, 0.1, positive},
		{uint64(100), uint8(101), 0.01, positive},
		{0.1, -0.1, 2, positive},
		{0, 0.1, 2, positive},

		// expect to NotInEpsilon
		{uint8(2), int16(-2), .001, negative},
		{uint64(100), uint8(102), 0.01, negative},
		{2.1, 2.2, 0.001, negative},
		{2.2, 2.1, 0.001, negative},
		{2.1, -2.2, 1, negative},
		{0.1, -0.1, 1.99, negative},
		{0, 0.1, 0.001, negative},

		// expect to fail in all cases
		{0.1, 0, 2, abnormal},
		{0, math.NaN(), 1, abnormal},
		{math.NaN(), 1, 1, abnormal},
		{0, 1, math.NaN(), abnormal},
		{math.NaN(), math.NaN(), 1, abnormal},
		{"bla-bla", 2.1, 0, abnormal},
		{2.1, "bla-bla", 0, abnormal},
	} {
		pos := InEpsilon(mockT, tc.giveA, tc.giveB, tc.giveDelta)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotInEpsilon(mockT, tc.giveA, tc.giveB, tc.giveDelta)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestContainNotContain(t *testing.T) {
	mockT := &testing.T{}

	type A struct{ Name, Value string }
	str := "Hello World"
	array := [2]string{"Foo", "Bar"}
	slice := []string{"Foo", "Bar"}
	complexSlice := []*A{{"b", "c"}, {"d", "e"}, {"g", "h"}, {"j", "k"}}
	map_ := map[interface{}]interface{}{"Foo": "Bar"}
	var zeroMap map[interface{}]interface{}

	for _, tc := range []struct {
		giveL, giveE interface{}
		want         testFlag
	}{
		// expect to Contain
		{str, "Hello", positive},
		{array, "Foo", positive},
		{slice, "Bar", positive},
		{complexSlice, &A{"g", "h"}, positive},
		{map_, "Foo", positive},

		// expect to NotContain
		{str, "Salut", negative},
		{array, "Salut", negative},
		{slice, "Salut", negative},
		{complexSlice, &A{"g", "e"}, negative},
		{complexSlice, (*A)(nil), negative},
		{map_, "Bar", negative},
		{map_, 111, negative},
		{zeroMap, "Bar", negative},

		// expect to fail in all case
		{str, 0, abnormal},
		{array, 1, abnormal},
		{slice, 2, abnormal},
		{0, "Hello", abnormal},
		{nil, "Hello", abnormal},
		{func() {}, "Hello", abnormal},
		{struct{}{}, "Hello", abnormal},
	} {
		pos := Contain(mockT, tc.giveL, tc.giveE)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotContain(mockT, tc.giveL, tc.giveE)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestSubsetNotSubset(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveL, giveS interface{}
		want         testFlag
	}{
		// expect to Subset
		{[]int{}, []int{}, positive},
		{[]int{1, 2, 3}, []int{}, positive},
		{[]int{1, 2, 3}, []int{1, 2}, positive},
		{[]int8{1, 2, 3}, []int8{3, 2}, positive},
		{[]uint64{1, 2, 3}, []uint64{1, 3, 2}, positive},
		{[]float64{.1, .2, .3}, []float64{.1}, positive},
		{[]string{"hello", "world"}, []string{"hello"}, positive},

		// expect to NotSubset
		{[]int{1, 2, 3}, []int{4, 5}, negative},
		{[]int8{1, 2, 3}, []int8{0}, negative},
		{[]uint64{1, 2, 3}, []uint64{1, 5}, negative},
		{[]float64{.1, .2, .3}, []float64{.1, .2, .5}, negative},
		{[]string{"hello", "world"}, []string{"hello", "x"}, negative},

		// expect to fail in all case
		{[]int{1, 2, 3}, nil, abnormal},
		{[]int{1, 2, 3}, []int8{}, abnormal},
		{[]uint64{1, 2, 3}, []int32{3, 2}, abnormal},
		{[]float64{.1, .2, .3}, []float32{.1}, abnormal},
		{[]int8{1, 2, 3}, uint8(1), abnormal},
		{[]string{"hello", "world"}, []byte{'h'}, abnormal},
		{"hello", []byte{'h'}, abnormal},
	} {
		pos := Subset(mockT, tc.giveL, tc.giveS)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotSubset(mockT, tc.giveL, tc.giveS)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestElementMatchNotElementMatch(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveA, giveB interface{}
		want         testFlag
	}{
		// expect to Subset
		{[]int{}, []int{}, positive},
		{[]int{1}, []int{1}, positive},
		{[]int{1, 1}, []int{1, 1}, positive},
		{[]int{1, 2}, []int{1, 2}, positive},
		{[]int{1, 2}, []int{2, 1}, positive},
		{[2]int{1, 2}, [2]int{2, 1}, positive},
		{[]string{"hello", "world"}, []string{"world", "hello"}, positive},
		{[]string{"hello", "hello"}, []string{"hello", "hello"}, positive},
		{[]string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}, positive},
		{[3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}, positive},

		// expect to NotSubset
		{[]int{1}, []int{1, 1}, negative},
		{[]int{1, 2}, []int{2, 2}, negative},
		{[]int{1, 1, 1, 2}, []int{2, 1}, negative},
		{[]string{"hello", "hello"}, []string{"hello"}, negative},

		// expect to fail in all case
		{nil, nil, abnormal},
		{[]int{}, nil, abnormal},
		{[]int{1, 2, 3}, []int8{1, 2, 3}, abnormal},
		{[]uint64{1, 2, 3}, []int32{1, 2, 3}, abnormal},
		{[]float64{.1, .2, .3}, []float32{.1, .2, .3}, abnormal},
		{[]int8{1, 2, 3}, uint8(1), abnormal},
		{[]string{"hello", "world"}, []byte{'h'}, abnormal},
		{"hello", []byte{'h'}, abnormal},
	} {
		pos := ElementMatch(mockT, tc.giveA, tc.giveB)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotElementMatch(mockT, tc.giveA, tc.giveB)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

type testInterface interface {
	TestMethod()
}

type testStruct struct{}

func (a *testStruct) TestMethod() {}

type testStruct2 struct{}

func TestSameTypeNotSameType(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveV, giveW interface{}
		want         testFlag
	}{
		// expect to SameType
		{1, 2, positive},
		{uint8('a'), byte('b'), positive},
		{3.14, 3.15, positive},
		{struct{ I string }{}, struct{ I string }{}, positive},
		{&testStruct{}, &testStruct{}, positive},
		{nil, nil, positive},
		{new(testInterface), (*testInterface)(nil), positive},

		// expect to NotSameType
		{1, uint(2), negative},
		{'a', byte('b'), negative},
		{3.14, float32(3.15), negative},
		{struct{ I string }{}, struct{ J string }{}, negative},
		{&testStruct{}, &testStruct2{}, negative},
		{testStruct{}, &testStruct{}, negative},
		{new(testInterface), testInterface(nil), negative},
	} {
		pos := SameType(mockT, tc.giveV, tc.giveW)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotSameType(mockT, tc.giveV, tc.giveW)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestImplementNotImplement(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveV, giveW interface{}
		want         testFlag
	}{
		// expect to Implement
		{&testStruct{}, (*testInterface)(nil), positive},
		{&testStruct{}, new(testInterface), positive},
		{errors.New("test"), (*error)(nil), positive},
		{&customError{}, new(error), positive},
		{0, (*interface{})(nil), positive},
		{uint64(0), new(interface{}), positive},

		// expect to NotImplement
		{testStruct{}, (*testInterface)(nil), negative},
		{&testStruct2{}, (*testInterface)(nil), negative},
		{"test", (*error)(nil), negative},
		{struct{}{}, new(error), negative},

		// fail in all cases
		{"", nil, abnormal},
		{"", 0, abnormal},
		{"", new(int), abnormal},
		{nil, (*interface{})(nil), abnormal},
	} {
		pos := Implement(mockT, tc.giveV, tc.giveW)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotImplement(mockT, tc.giveV, tc.giveW)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestPanicNotPanic(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		give func()
		want testFlag
	}{
		// expect to Panic
		{func() { panic("Panic!") }, positive},
		{func() { panic(0) }, positive},
		{func() { panic(nil) }, positive},

		// expect to NotPanic
		{func() {}, negative},
	} {
		pos := Panic(mockT, tc.give)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
		neg := NotPanic(mockT, tc.give)
		if (tc.want == negative && !neg) || (tc.want != negative && neg) {
			fail(t)
		}
	}
}

func TestPanicWithValuePanicWithError(t *testing.T) {
	mockT := &testing.T{}

	for _, tc := range []struct {
		giveF func()
		giveW interface{}
		want  testFlag
	}{
		// expect to pass PanicWithValue
		{func() { panic("Panic!") }, "Panic!", positive},
		{func() { panic(0) }, 0, positive},
		{func() { panic(nil) }, nil, positive},
		{func() { panic(errors.New("panic")) }, errors.New("panic"), positive},

		// expect to fail PanicWithValue
		{func() {}, nil, negative},
		{func() { panic("Panic!") }, "Panic", negative},
		{func() { panic(uint8(0)) }, 0, negative},
		{func() { panic(errors.New("panic")) }, "panic", negative},
	} {
		pos := PanicWithValue(mockT, tc.giveW, tc.giveF)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
	}

	for _, tc := range []struct {
		giveF func()
		giveW string
		want  testFlag
	}{
		// expect to pass PanicWithError
		{func() { panic(errors.New("panic")) }, "panic", positive},
		{func() { panic((*customError)(nil)) }, "customError (nil)", positive},
		{func() { panic(&customError{}) }, "customError", positive},

		// expect to fail PanicWithError
		{func() {}, "", negative},
		{func() { panic("Panic!") }, "", negative},
		{func() { panic(nil) }, "", negative},
	} {
		pos := PanicWithError(mockT, tc.giveW, tc.giveF)
		if (tc.want == positive && !pos) || (tc.want != positive && pos) {
			fail(t)
		}
	}
}

func TestFileDirExistNotExist(t *testing.T) {
	mockT := &testing.T{}

	for _, pair := range [][2]string{
		{"xtesting.go", "xtesting.go_symlink"},
		{"xxx", "xxx_symlink"},
		{"../xtesting", "xtesting_symlink"},
	} {
		if err := os.Symlink(pair[0], pair[1]); err != nil {
			fail(t)
			t.FailNow()
		}
	}
	defer func() {
		matches, _ := filepath.Glob("*_symlink")
		for _, match := range matches {
			if os.Remove(match) != nil {
				fail(t)
				t.FailNow()
			}
		}
	}()

	for _, tc := range []struct {
		give      string
		wantFile  testFlag
		wantLfile testFlag
		wantDir   testFlag
		wantLdir  testFlag
		wantLsym  testFlag
	}{
		{"xtesting.go", positive, positive, negative, negative, negative},
		{"./xtesting.go/", abnormal, abnormal, abnormal, abnormal, abnormal},
		{"xtesting.go_symlink", positive, positive, negative, negative, positive},
		{"xxx", negative, negative, negative, negative, negative},
		{"xxx_symlink", negative, positive, negative, negative, positive},

		{"../xtesting", negative, negative, positive, positive, negative},
		{"../xtesting/", negative, negative, positive, positive, negative},
		{"xtesting_symlink", negative, positive, positive, negative, positive},
		{".", negative, negative, positive, positive, negative},
		{"..", negative, negative, positive, positive, negative},
	} {
		// File
		pos := FileExist(mockT, tc.give)
		if (tc.wantFile == positive && !pos) || (tc.wantFile != positive && pos) {
			fail(t)
		}
		neg := FileNotExist(mockT, tc.give)
		if (tc.wantFile == negative && !neg) || (tc.wantFile != negative && neg) {
			fail(t)
		}
		pos = FileLexist(mockT, tc.give)
		if (tc.wantLfile == positive && !pos) || (tc.wantLfile != positive && pos) {
			fail(t)
		}
		neg = FileNotLexist(mockT, tc.give)
		if (tc.wantLfile == negative && !neg) || (tc.wantLfile != negative && neg) {
			fail(t)
		}

		// Dir
		pos = DirExist(mockT, tc.give)
		if (tc.wantDir == positive && !pos) || (tc.wantDir != positive && pos) {
			fail(t)
		}
		neg = DirNotExist(mockT, tc.give)
		if (tc.wantDir == negative && !neg) || (tc.wantDir != negative && neg) {
			fail(t)
		}
		pos = DirLexist(mockT, tc.give)
		if (tc.wantLdir == positive && !pos) || (tc.wantLdir != positive && pos) {
			fail(t)
		}
		neg = DirNotLexist(mockT, tc.give)
		if (tc.wantLdir == negative && !neg) || (tc.wantLdir != negative && neg) {
			fail(t)
		}

		// Symlink
		pos = SymlinkLexist(mockT, tc.give)
		if (tc.wantLsym == positive && !pos) || (tc.wantLsym != positive && pos) {
			fail(t)
		}
		neg = SymlinkNotLexist(mockT, tc.give)
		if (tc.wantLsym == negative && !neg) || (tc.wantLsym != negative && neg) {
			fail(t)
		}
	}
}

// ==================================
// testings for helpers and internals
// ==================================

func TestCombineMsgAndArgs(t *testing.T) {
	for _, tc := range []struct {
		give []interface{}
		want string
	}{
		{nil, ""},
		{[]interface{}{"0"}, "0"},
		{[]interface{}{[]int{1, 2}}, "[1 2]"},
		{[]interface{}{nil}, "<nil>"},
		{[]interface{}{"a%sc", "b"}, "abc"},
	} {
		s := combineMsgAndArgs(tc.give...)
		if s != tc.want {
			fail(t)
		}
	}
}

type mockFinishFlagTestingT struct {
	testing.TB
	finished bool
}

func (m *mockFinishFlagTestingT) Fail()                { m.finished = false }
func (m *mockFinishFlagTestingT) FailNow()             { m.finished = true }
func (m *mockFinishFlagTestingT) Fatal(...interface{}) { m.finished = true }

func TestFailTestOptions(t *testing.T) {
	captureStderr := func(f func()) string {
		stderr := os.Stderr
		r, w, _ := os.Pipe()
		os.Stderr = w
		defer func() { os.Stderr = stderr }()
		f()
		w.Close()
		bs, _ := ioutil.ReadAll(r)
		return string(bs)
	}

	t.Run("SetExtraSkip", func(t *testing.T) {
		mockT := &testing.T{}

		// 1
		SetExtraSkip(1)
		result1 := captureStderr(func() {
			if failTest(mockT, -1, "a", "") != false {
				fail(t)
			}
		})
		if !strings.HasSuffix(result1, "a\n") {
			fail(t)
		}

		// 2
		result2 := captureStderr(func() {
			if failTest(mockT, 0, "a", ", %%a%s", "bbb") != false {
				fail(t)
			}
		})
		if !strings.HasSuffix(result2, "a, %abbb\n") {
			fail(t)
		}

		// 3
		SetExtraSkip(0)
		result3 := captureStderr(func() {
			if failTest(mockT, 0, "%s", ", %s%s%03d", "xx", "yy", 3) != false {
				fail(t)
			}
		})
		if !strings.HasSuffix(result3, "%s, xxyy003\n") {
			fail(t)
		}
	})

	t.Run("UseFailNow", func(t *testing.T) {
		mockT := &mockFinishFlagTestingT{}
		UseFailNow(true)
		result1 := captureStderr(func() {
			if failTest(mockT, 1, "") != false {
				fail(t)
			}
		})
		if !mockT.finished {
			fail(t)
		}
		if !strings.HasPrefix(result1, "xtesting_test.go") {
			fail(t)
		}

		mockT = &mockFinishFlagTestingT{}
		UseFailNow(false)
		result2 := captureStderr(func() {
			if failTest(mockT, 1, "") != false {
				fail(t)
			}
		})
		if mockT.finished {
			fail(t)
		}
		if !strings.HasPrefix(result2, "xtesting_test.go") {
			fail(t)
		}
	})
}

func TestAssert(t *testing.T) {
	funcDidPanic, _ := checkPanic(func() {
		Assert(true, "test %s", "test")
	})
	if funcDidPanic {
		fail(t)
	}

	funcDidPanic, panicValue := checkPanic(func() {
		Assert(false, "test %s", "test")
	})
	if !funcDidPanic {
		fail(t)
	}
	if panicValue != "test test" {
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
