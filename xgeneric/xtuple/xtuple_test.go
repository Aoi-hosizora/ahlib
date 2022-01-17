//go:build go1.18
// +build go1.18

package xtuple

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func TestPairs(t *testing.T) {
	t.Run("tuple", func(t *testing.T) {
		a := NewTuple(-1, uint(2))
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.String(), "[-1, 2]")
		b := NewTuplePtr("three", 0.4)
		xtestingEqual(t, b.Item1, "three")
		xtestingEqual(t, b.Item2, 0.4)
		xtestingEqual(t, b.String(), "[three, 0.4]")
		c := NewTuple([5]int{}, []int{6})
		xtestingEqual(t, c.Item1, [5]int{})
		xtestingEqual(t, c.Item2, []int{6})
		xtestingEqual(t, c.String(), "[[0 0 0 0 0], [6]]")
	})

	t.Run("triple", func(t *testing.T) {
		a := NewTriple(-1, uint(2), 3i)
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.Item3, 3i)
		xtestingEqual(t, a.String(), "[-1, 2, (0+3i)]")
		b := NewTriplePtr("four", 0.5, byte(6))
		xtestingEqual(t, b.Item1, "four")
		xtestingEqual(t, b.Item2, 0.5)
		xtestingEqual(t, b.Item3, byte(6))
		xtestingEqual(t, b.String(), "[four, 0.5, 6]")
		c := NewTriple([7]int{}, []int{8}, rune(9))
		xtestingEqual(t, c.Item1, [7]int{})
		xtestingEqual(t, c.Item2, []int{8})
		xtestingEqual(t, c.Item3, rune(9))
		xtestingEqual(t, c.String(), "[[0 0 0 0 0 0 0], [8], 9]")
	})

	t.Run("quadruple", func(t *testing.T) {
		a := NewQuadruple(-1, uint(2), 3i, byte(4))
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.Item3, 3i)
		xtestingEqual(t, a.Item4, byte(4))
		xtestingEqual(t, a.String(), "[-1, 2, (0+3i), 4]")
		b := NewQuadruplePtr("five", 0.6, rune(7), uint8(8))
		xtestingEqual(t, b.Item1, "five")
		xtestingEqual(t, b.Item2, 0.6)
		xtestingEqual(t, b.Item3, rune(7))
		xtestingEqual(t, b.Item4, uint8(8))
		xtestingEqual(t, b.String(), "[five, 0.6, 7, 8]")
	})

	t.Run("quintuple", func(t *testing.T) {
		a := NewQuintuple(-1, uint(2), 3i, byte(4), rune(5))
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.Item3, 3i)
		xtestingEqual(t, a.Item4, byte(4))
		xtestingEqual(t, a.Item5, rune(5))
		xtestingEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5]")
		b := NewQuintuplePtr("six", 0.7, uint8(8), int64(9), true)
		xtestingEqual(t, b.Item1, "six")
		xtestingEqual(t, b.Item2, 0.7)
		xtestingEqual(t, b.Item3, uint8(8))
		xtestingEqual(t, b.Item4, int64(9))
		xtestingEqual(t, b.Item5, true)
		xtestingEqual(t, b.String(), "[six, 0.7, 8, 9, true]")
	})

	t.Run("sextuple", func(t *testing.T) {
		a := NewSextuple(-1, uint(2), 3i, byte(4), rune(5), 6.6)
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.Item3, 3i)
		xtestingEqual(t, a.Item4, byte(4))
		xtestingEqual(t, a.Item5, rune(5))
		xtestingEqual(t, a.Item6, 6.6)
		xtestingEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5, 6.6]")
		b := NewSextuplePtr("seven", 0.8, uint8(9), int64(10), 11+11i, true)
		xtestingEqual(t, b.Item1, "seven")
		xtestingEqual(t, b.Item2, 0.8)
		xtestingEqual(t, b.Item3, uint8(9))
		xtestingEqual(t, b.Item4, int64(10))
		xtestingEqual(t, b.Item5, 11+11i)
		xtestingEqual(t, b.Item6, true)
		xtestingEqual(t, b.String(), "[seven, 0.8, 9, 10, (11+11i), true]")
	})

	t.Run("septuple", func(t *testing.T) {
		a := NewSeptuple(-1, uint(2), 3i, byte(4), rune(5), 6.6, "seven")
		xtestingEqual(t, a.Item1, -1)
		xtestingEqual(t, a.Item2, uint(2))
		xtestingEqual(t, a.Item3, 3i)
		xtestingEqual(t, a.Item4, byte(4))
		xtestingEqual(t, a.Item5, rune(5))
		xtestingEqual(t, a.Item6, 6.6)
		xtestingEqual(t, a.Item7, "seven")
		xtestingEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5, 6.6, seven]")
		b := NewSeptuplePtr(0.8, uint8(9), int64(10), 11+11i, true, false, int8(-12))
		xtestingEqual(t, b.Item1, 0.8)
		xtestingEqual(t, b.Item2, uint8(9))
		xtestingEqual(t, b.Item3, int64(10))
		xtestingEqual(t, b.Item4, 11+11i)
		xtestingEqual(t, b.Item5, true)
		xtestingEqual(t, b.Item6, false)
		xtestingEqual(t, b.Item7, int8(-12))
		xtestingEqual(t, b.String(), "[0.8, 9, 10, (11+11i), true, false, -12]")
	})
}

func TestSugarIfThen(t *testing.T) {
	xtestingEqual(t, IfThen(true, "a"), "a")
	xtestingEqual(t, IfThen(false, "a"), "")
	xtestingEqual(t, IfThenElse(true, "x", "y"), "x")
	xtestingEqual(t, IfThenElse(false, "x", "y"), "y")

	xtestingEqual(t, IfThen(true, 1.1), 1.1)
	xtestingEqual(t, IfThen(false, 1.1), 0.0)
	xtestingEqual(t, IfThenElse(true, uint(1), uint(2)), uint(1))
	xtestingEqual(t, IfThenElse(false, 1+1i, 2+2i), 2+2i)
}

func TestSugarPtr(t *testing.T) {
	i := 1
	u := uint(1)
	a := [2]float64{1, 2}
	m := map[string]interface{}{"1": uint(1)}
	s := []string{"1", "1"}

	xtestingEqual(t, *ValPtr(i), i)
	xtestingEqual(t, *ValPtr(u), u)
	xtestingEqual(t, *ValPtr(a), a)
	xtestingEqual(t, *ValPtr(m), m)
	xtestingEqual(t, *ValPtr(s), s)
	xtestingEqual(t, *ValPtr(&i), &i)
	xtestingEqual(t, *ValPtr(&u), &u)
	xtestingEqual(t, *ValPtr(&a), &a)
	xtestingEqual(t, *ValPtr(&m), &m)
	xtestingEqual(t, *ValPtr(&s), &s)
	xtestingEqual(t, **ValPtr(ValPtr(&i)), &i)
	xtestingEqual(t, **ValPtr(ValPtr(&u)), &u)
	xtestingEqual(t, **ValPtr(ValPtr(&a)), &a)
	xtestingEqual(t, **ValPtr(ValPtr(&m)), &m)
	xtestingEqual(t, **ValPtr(ValPtr(&s)), &s)

	xtestingEqual(t, PtrVal[int](nil, i), i)
	xtestingEqual(t, PtrVal[uint](nil, u), u)
	xtestingEqual(t, PtrVal[[2]float64](nil, a), a)
	xtestingEqual(t, PtrVal[map[string]interface{}](nil, m), m)
	xtestingEqual(t, PtrVal[[]string](nil, s), s)
	xtestingEqual(t, PtrVal(&i, i), i)
	xtestingEqual(t, PtrVal(&u, u), u)
	xtestingEqual(t, PtrVal(&a, a), a)
	xtestingEqual(t, PtrVal(&m, m), m)
	xtestingEqual(t, PtrVal(&s, s), s)
	xtestingEqual(t, PtrVal(ValPtr(&i), nil), &i)
	xtestingEqual(t, PtrVal(ValPtr(&u), nil), &u)
	xtestingEqual(t, PtrVal(ValPtr(&a), nil), &a)
	xtestingEqual(t, PtrVal(ValPtr(&m), nil), &m)
	xtestingEqual(t, PtrVal(ValPtr(&s), nil), &s)
}

func TestSugarPairs(t *testing.T) {
	_a := func() (int, uint) { return 1, 2 }
	_b := func() (int, uint, string) { return 1, 2, "3" }
	_c := func() (int, uint, string, float64) { return 1, 2, "3", 0.4 }
	_d := func() (int, uint, string, float64, complex128) { return 1, 2, "3", 0.4, 5 + 5i }
	_e := func() (int, uint, string, float64, complex128, rune) { return 1, 2, "3", 0.4, 5 + 5i, 6 }
	_f := func() (int, uint, string, float64, complex128, rune, bool) { return 1, 2, "3", 0.4, 5 + 5i, 6, true }

	xtestingEqual(t, TupleItem1(_a()), 1)
	xtestingEqual(t, TupleItem2(_a()), uint(2))
	xtestingEqual(t, TripleItem1(_b()), 1)
	xtestingEqual(t, TripleItem2(_b()), uint(2))
	xtestingEqual(t, TripleItem3(_b()), "3")
	xtestingEqual(t, QuadrupleItem1(_c()), 1)
	xtestingEqual(t, QuadrupleItem2(_c()), uint(2))
	xtestingEqual(t, QuadrupleItem3(_c()), "3")
	xtestingEqual(t, QuadrupleItem4(_c()), 0.4)
	xtestingEqual(t, QuintupleItem1(_d()), 1)
	xtestingEqual(t, QuintupleItem2(_d()), uint(2))
	xtestingEqual(t, QuintupleItem3(_d()), "3")
	xtestingEqual(t, QuintupleItem4(_d()), 0.4)
	xtestingEqual(t, QuintupleItem5(_d()), 5+5i)
	xtestingEqual(t, SextupleItem1(_e()), 1)
	xtestingEqual(t, SextupleItem2(_e()), uint(2))
	xtestingEqual(t, SextupleItem3(_e()), "3")
	xtestingEqual(t, SextupleItem4(_e()), 0.4)
	xtestingEqual(t, SextupleItem5(_e()), 5+5i)
	xtestingEqual(t, SextupleItem6(_e()), rune(6))
	xtestingEqual(t, SeptupleItem1(_f()), 1)
	xtestingEqual(t, SeptupleItem2(_f()), uint(2))
	xtestingEqual(t, SeptupleItem3(_f()), "3")
	xtestingEqual(t, SeptupleItem4(_f()), 0.4)
	xtestingEqual(t, SeptupleItem5(_f()), 5+5i)
	xtestingEqual(t, SeptupleItem6(_f()), rune(6))
	xtestingEqual(t, SeptupleItem7(_f()), true)
}

func failTest(t testing.TB, msg string) bool {
	_, file, line, _ := runtime.Caller(2)
	fmt.Println(fmt.Sprintf("%s:%d %s", path.Base(file), line, msg))
	t.Fail()
	return false
}

func xtestingEqual(t testing.TB, give, want interface{}) bool {
	if give != nil && want != nil && (reflect.TypeOf(give).Kind() == reflect.Func || reflect.TypeOf(want).Kind() == reflect.Func) {
		return failTest(t, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (xtesting: cannot take func type as argument)", give, want))
	}
	if !reflect.DeepEqual(give, want) {
		return failTest(t, fmt.Sprintf("Equal: expected `%#v`, actual `%#v`", want, give))
	}
	return true
}
