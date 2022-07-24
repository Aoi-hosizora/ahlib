//go:build go1.18
// +build go1.18

package xtuple

import (
	"github.com/Aoi-hosizora/ahlib/xgeneric/internal"
	"testing"
)

func TestPair(t *testing.T) {
	t.Run("tuple", func(t *testing.T) {
		a := NewTuple(-1, uint(2))
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.String(), "[-1, 2]")
		b := NewTuplePtr("three", 0.4)
		internal.TestEqual(t, b.Item1, "three")
		internal.TestEqual(t, b.Item2, 0.4)
		internal.TestEqual(t, b.String(), "[three, 0.4]")
		c := NewTuple([5]int{}, []int{6})
		internal.TestEqual(t, c.Item1, [5]int{})
		internal.TestEqual(t, c.Item2, []int{6})
		internal.TestEqual(t, c.String(), "[[0 0 0 0 0], [6]]")
	})

	t.Run("triple", func(t *testing.T) {
		a := NewTriple(-1, uint(2), 3i)
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.Item3, 3i)
		internal.TestEqual(t, a.String(), "[-1, 2, (0+3i)]")
		b := NewTriplePtr("four", 0.5, byte(6))
		internal.TestEqual(t, b.Item1, "four")
		internal.TestEqual(t, b.Item2, 0.5)
		internal.TestEqual(t, b.Item3, byte(6))
		internal.TestEqual(t, b.String(), "[four, 0.5, 6]")
		c := NewTriple([7]int{}, []int{8}, rune(9))
		internal.TestEqual(t, c.Item1, [7]int{})
		internal.TestEqual(t, c.Item2, []int{8})
		internal.TestEqual(t, c.Item3, rune(9))
		internal.TestEqual(t, c.String(), "[[0 0 0 0 0 0 0], [8], 9]")
	})

	t.Run("quadruple", func(t *testing.T) {
		a := NewQuadruple(-1, uint(2), 3i, byte(4))
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.Item3, 3i)
		internal.TestEqual(t, a.Item4, byte(4))
		internal.TestEqual(t, a.String(), "[-1, 2, (0+3i), 4]")
		b := NewQuadruplePtr("five", 0.6, rune(7), uint8(8))
		internal.TestEqual(t, b.Item1, "five")
		internal.TestEqual(t, b.Item2, 0.6)
		internal.TestEqual(t, b.Item3, rune(7))
		internal.TestEqual(t, b.Item4, uint8(8))
		internal.TestEqual(t, b.String(), "[five, 0.6, 7, 8]")
	})

	t.Run("quintuple", func(t *testing.T) {
		a := NewQuintuple(-1, uint(2), 3i, byte(4), rune(5))
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.Item3, 3i)
		internal.TestEqual(t, a.Item4, byte(4))
		internal.TestEqual(t, a.Item5, rune(5))
		internal.TestEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5]")
		b := NewQuintuplePtr("six", 0.7, uint8(8), int64(9), true)
		internal.TestEqual(t, b.Item1, "six")
		internal.TestEqual(t, b.Item2, 0.7)
		internal.TestEqual(t, b.Item3, uint8(8))
		internal.TestEqual(t, b.Item4, int64(9))
		internal.TestEqual(t, b.Item5, true)
		internal.TestEqual(t, b.String(), "[six, 0.7, 8, 9, true]")
	})

	t.Run("sextuple", func(t *testing.T) {
		a := NewSextuple(-1, uint(2), 3i, byte(4), rune(5), 6.6)
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.Item3, 3i)
		internal.TestEqual(t, a.Item4, byte(4))
		internal.TestEqual(t, a.Item5, rune(5))
		internal.TestEqual(t, a.Item6, 6.6)
		internal.TestEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5, 6.6]")
		b := NewSextuplePtr("seven", 0.8, uint8(9), int64(10), 11+11i, true)
		internal.TestEqual(t, b.Item1, "seven")
		internal.TestEqual(t, b.Item2, 0.8)
		internal.TestEqual(t, b.Item3, uint8(9))
		internal.TestEqual(t, b.Item4, int64(10))
		internal.TestEqual(t, b.Item5, 11+11i)
		internal.TestEqual(t, b.Item6, true)
		internal.TestEqual(t, b.String(), "[seven, 0.8, 9, 10, (11+11i), true]")
	})

	t.Run("septuple", func(t *testing.T) {
		a := NewSeptuple(-1, uint(2), 3i, byte(4), rune(5), 6.6, "seven")
		internal.TestEqual(t, a.Item1, -1)
		internal.TestEqual(t, a.Item2, uint(2))
		internal.TestEqual(t, a.Item3, 3i)
		internal.TestEqual(t, a.Item4, byte(4))
		internal.TestEqual(t, a.Item5, rune(5))
		internal.TestEqual(t, a.Item6, 6.6)
		internal.TestEqual(t, a.Item7, "seven")
		internal.TestEqual(t, a.String(), "[-1, 2, (0+3i), 4, 5, 6.6, seven]")
		b := NewSeptuplePtr(0.8, uint8(9), int64(10), 11+11i, true, false, int8(-12))
		internal.TestEqual(t, b.Item1, 0.8)
		internal.TestEqual(t, b.Item2, uint8(9))
		internal.TestEqual(t, b.Item3, int64(10))
		internal.TestEqual(t, b.Item4, 11+11i)
		internal.TestEqual(t, b.Item5, true)
		internal.TestEqual(t, b.Item6, false)
		internal.TestEqual(t, b.Item7, int8(-12))
		internal.TestEqual(t, b.String(), "[0.8, 9, 10, (11+11i), true, false, -12]")
	})
}

func TestPairItem(t *testing.T) {
	_a := func() (int, uint) { return 1, 2 }
	_b := func() (int, uint, string) { return 1, 2, "3" }
	_c := func() (int, uint, string, float64) { return 1, 2, "3", 0.4 }
	_d := func() (int, uint, string, float64, complex128) { return 1, 2, "3", 0.4, 5 + 5i }
	_e := func() (int, uint, string, float64, complex128, rune) { return 1, 2, "3", 0.4, 5 + 5i, 6 }
	_f := func() (int, uint, string, float64, complex128, rune, bool) { return 1, 2, "3", 0.4, 5 + 5i, 6, true }

	internal.TestEqual(t, TupleItem1(_a()), 1)
	internal.TestEqual(t, TupleItem2(_a()), uint(2))
	internal.TestEqual(t, TripleItem1(_b()), 1)
	internal.TestEqual(t, TripleItem2(_b()), uint(2))
	internal.TestEqual(t, TripleItem3(_b()), "3")
	internal.TestEqual(t, QuadrupleItem1(_c()), 1)
	internal.TestEqual(t, QuadrupleItem2(_c()), uint(2))
	internal.TestEqual(t, QuadrupleItem3(_c()), "3")
	internal.TestEqual(t, QuadrupleItem4(_c()), 0.4)
	internal.TestEqual(t, QuintupleItem1(_d()), 1)
	internal.TestEqual(t, QuintupleItem2(_d()), uint(2))
	internal.TestEqual(t, QuintupleItem3(_d()), "3")
	internal.TestEqual(t, QuintupleItem4(_d()), 0.4)
	internal.TestEqual(t, QuintupleItem5(_d()), 5+5i)
	internal.TestEqual(t, SextupleItem1(_e()), 1)
	internal.TestEqual(t, SextupleItem2(_e()), uint(2))
	internal.TestEqual(t, SextupleItem3(_e()), "3")
	internal.TestEqual(t, SextupleItem4(_e()), 0.4)
	internal.TestEqual(t, SextupleItem5(_e()), 5+5i)
	internal.TestEqual(t, SextupleItem6(_e()), rune(6))
	internal.TestEqual(t, SeptupleItem1(_f()), 1)
	internal.TestEqual(t, SeptupleItem2(_f()), uint(2))
	internal.TestEqual(t, SeptupleItem3(_f()), "3")
	internal.TestEqual(t, SeptupleItem4(_f()), 0.4)
	internal.TestEqual(t, SeptupleItem5(_f()), 5+5i)
	internal.TestEqual(t, SeptupleItem6(_f()), rune(6))
	internal.TestEqual(t, SeptupleItem7(_f()), true)
}
