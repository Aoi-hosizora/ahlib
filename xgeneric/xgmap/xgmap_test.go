//go:build go1.18
// +build go1.18

package xgmap

import (
	"github.com/Aoi-hosizora/ahlib/xgeneric/internal"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xsugar"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xtuple"
	"math"
	"sort"
	"strconv"
	"testing"
)

func TestKeysValues(t *testing.T) {
	internal.TestEqual(t, Keys(map[int]uint{}), []int{})
	internal.TestEqual(t, Values(map[int]uint{}), []uint{})
	internal.TestEqual(t, KeyValues(map[int]uint{}), []xtuple.Tuple[int, uint]{})
	internal.TestEqual(t, Keys(map[int]uint{1: 0}), []int{1})
	internal.TestEqual(t, Values(map[int]uint{0: 1}), []uint{1})
	internal.TestEqual(t, KeyValues(map[int]uint{0: 1}), []xtuple.Tuple[int, uint]{{0, 1}})
	internal.TestEqual(t, sorted(Keys(map[string]int{"a": 1, "b": 2, "c": 3})), []string{"a", "b", "c"})
	internal.TestEqual(t, sorted(Values(map[string]int{"a": 1, "b": 2, "c": 3})), []int{1, 2, 3})
	internal.TestEqual(t, sorted2[string, int, xtuple.Tuple[string, int]](KeyValues(map[string]int{"a": 1, "b": 2, "c": 3})), []xtuple.Tuple[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
	internal.TestEqual(t, sorted(Keys(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), []string{"Authorization", "X"})
	internal.TestEqual(t, sorted2[string, []string, []string](Values(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), [][]string{{"!", "@"}, {"Token"}})
	internal.TestEqual(t, sorted2[string, []string, xtuple.Tuple[string, []string]](KeyValues(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), []xtuple.Tuple[string, []string]{{"Authorization", []string{"Token"}}, {"X", []string{"!", "@"}}})
	internal.TestEqual(t, sorted(Keys(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []uint32{1, 2, 3, 4})
	internal.TestEqual(t, sorted(Values(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []float64{0.1, 0.2, 0.3, 0.4})
	internal.TestEqual(t, sorted2[uint32, float64, xtuple.Tuple[uint32, float64]](KeyValues(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []xtuple.Tuple[uint32, float64]{{1, 0.1}, {2, 0.2}, {3, 0.3}, {4, 0.4}})

	internal.TestPanic(t, true, func() { FromKeys[bool, bool]([]bool{}, nil) })
	internal.TestPanic(t, true, func() { FromValues[bool, bool]([]bool{}, nil) })
	internal.TestPanic(t, false, func() { FromKeyValues([]xtuple.Tuple[bool, bool]{}) })
	internal.TestEqual(t, FromKeys([]int{}, func(i int, t int) uint { return uint(i) }), map[int]uint{})
	internal.TestEqual(t, FromValues([]uint{}, func(i int, t uint) int { return i }), map[int]uint{})
	internal.TestEqual(t, FromKeyValues([]xtuple.Tuple[int, uint]{}), map[int]uint{})
	internal.TestEqual(t, FromKeys([]int{1}, func(i int, t int) uint { return uint(i) }), map[int]uint{1: 0})
	internal.TestEqual(t, FromValues([]uint{1}, func(i int, t uint) int { return i }), map[int]uint{0: 1})
	internal.TestEqual(t, FromKeyValues([]xtuple.Tuple[int, uint]{{0, 1}}), map[int]uint{0: 1})
	internal.TestEqual(t, FromKeys([]string{"a", "b", "c"}, func(i int, t string) int { return i + 1 }), map[string]int{"a": 1, "b": 2, "c": 3})
	internal.TestEqual(t, FromValues([]int{1, 2, 3}, func(i int, t int) byte { return byte(i) + 'a' }), map[byte]int{'a': 1, 'b': 2, 'c': 3})
	internal.TestEqual(t, FromKeyValues([]xtuple.Tuple[byte, int]{{'a', 1}, {'b', 2}, {'c', 3}}), map[byte]int{'a': 1, 'b': 2, 'c': 3})
	internal.TestEqual(t, FromKeys([]string{"!", "@", "#", "$", "%"}, func(i int, k string) bool { return false }), map[string]bool{"!": false, "@": false, "#": false, "$": false, "%": false})
	internal.TestEqual(t, FromValues([]string{"!", "@", "#", "$", "%"}, func(i int, k string) int { return i }), map[int]string{0: "!", 1: "@", 2: "#", 3: "$", 4: "%"})
	internal.TestEqual(t, FromKeyValues([]xtuple.Tuple[int, string]{{0, "!"}, {1, "@"}, {2, "#"}, {3, "$"}, {4, "%"}}), map[int]string{0: "!", 1: "@", 2: "#", 3: "$", 4: "%"})
}

func TestForeach(t *testing.T) {
	internal.TestPanic(t, true, func() { Foreach(map[int]bool{}, nil) })
	internal.TestPanic(t, false, func() { Foreach(map[int]bool{}, func(t int, t2 bool) {}) })

	test1 := 0
	Foreach(map[int]int{1: 2, 2: 3, 3: 4, 4: 5}, func(t int, t2 int) { test1 += t + t2 })
	internal.TestEqual(t, test1, 1+2+2+3+3+4+4+5)
	test2 := uint(1)
	Foreach(map[uint]uint{1: 1, 2: 2, 3: 3, 4: 4}, func(t uint, t2 uint) { test2 *= t * t2 })
	internal.TestEqual(t, test2, uint(1*1*2*2*3*3*4*4))
	test3 := float32(0)
	Foreach(map[float32]float32{1.0: 0.0, 2.0: 1.0, 3.0: 2.0}, func(t float32, t2 float32) { test3 -= t - t2 })
	internal.TestEqual(t, test3, float32(-1.0-1.0-1.0))
	test4 := make([]string, 0)
	Foreach(map[string]string{"1": "a", "2": "b", "3": "c"}, func(t string, t2 string) { test4 = append(test4, t+t2) })
	internal.TestEqual(t, sorted(test4), []string{"1a", "2b", "3c"})
}

func TestMapExpand(t *testing.T) {
	internal.TestPanic(t, true, func() { Map[bool, bool, bool, bool](map[bool]bool{}, nil) })
	internal.TestEqual(t, Map(map[bool]bool{}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{})
	internal.TestEqual(t, Map(map[bool]bool{true: false, false: true}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{false: false, true: true})
	internal.TestEqual(t, Map(map[string]int{"!": 1, "@": 2, "#": 3}, func(t string, t2 int) (int32, string) { return int32(t2) + 1, t }), map[int32]string{2: "!", 3: "@", 4: "#"})
	internal.TestEqual(t, Map(map[float64]bool{1.: true, 2.: true}, func(t float64, t2 bool) (string, bool) { return strconv.FormatFloat(t, 'f', 1, 64), !t2 }), map[string]bool{"1.0": false, "2.0": false})
	internal.TestEqual(t, Map(map[string]string{"1": "", "@": "", "3": "", "4": ""}, func(t string, t2 string) (int32, bool) {
		s, err := strconv.Atoi(t)
		return int32(s), err == nil
	}), map[int32]bool{1: true, 0: false, 3: true, 4: true})

	internal.TestPanic(t, true, func() { Expand[bool, bool, bool, bool](map[bool]bool{}, nil) })
	internal.TestEqual(t, Expand(map[bool]bool{}, func(t bool, t2 bool) []xtuple.Tuple[bool, bool] { return []xtuple.Tuple[bool, bool]{{!t, !t}} }), map[bool]bool{})
	internal.TestEqual(t, Expand(map[bool]bool{true: false, false: true}, func(t bool, t2 bool) []xtuple.Tuple[bool, bool] { return []xtuple.Tuple[bool, bool]{{!t, !t}, {t, t}} }), map[bool]bool{false: false, true: true})
	internal.TestEqual(t, Expand(map[string]int{"!": 1, "@": 2, "#": 3}, func(t string, t2 int) []xtuple.Tuple[int32, string] {
		return []xtuple.Tuple[int32, string]{{int32(t2), t}, {int32(t2) + 10, t + " "}}
	}), map[int32]string{1: "!", 11: "! ", 2: "@", 12: "@ ", 3: "#", 13: "# "})
	internal.TestEqual(t, Expand(map[float64]bool{1.: true, 2.: true}, func(t float64, t2 bool) []xtuple.Tuple[string, bool] { return nil }), map[string]bool{})
	internal.TestEqual(t, Expand(map[string]string{"1": "", "@": "", "3": "", "4": ""}, func(t string, t2 string) []xtuple.Tuple[int32, bool] {
		s, err := strconv.Atoi(t)
		return []xtuple.Tuple[int32, bool]{{int32(s), err == nil}, {int32(s) + 10, err != nil}}
	}), map[int32]bool{1: true, 0: false, 3: true, 4: true, 11: false, 10: true, 13: false, 14: false})
}

func TestReduce(t *testing.T) {
	internal.TestPanic(t, true, func() { Reduce(map[int]bool{}, true, nil) })
	internal.TestEqual(t, Reduce(map[int]bool{}, true, func(k bool, t int, t2 bool) bool { return false }), true)
	internal.TestEqual(t, sorted(Reduce(map[string]int{"1": 1, "2": 2, "3": 3}, []string{"0"}, func(k []string, t string, t2 int) []string { return append(k, t+strconv.Itoa(t2)) })), []string{"0", "11", "22", "33"})
	internal.TestEqual(t, Reduce(map[int]float64{9: 0.5, 8: 1.5, 7: 2.5, 6: 3.5}, 1., func(k float64, t int, t2 float64) float64 { return k + float64(t)*t2 }), 1+9*0.5+8*1.5+7*2.5+6*3.5)

	fractions := map[int][2]int{0: {5, 1}, 1: {3, 6}, 2: {2, 0}, 3: {3, 1}}
	results := Reduce(Map(fractions, func(t int, t2 [2]int) (int, *float64) {
		if t2[1] == 0 {
			return t, nil
		}
		r := float64(t2[0]) / float64(t2[1])
		return t, &r
	}), 0.0, func(k float64, t int, t2 *float64) float64 {
		if t2 == nil {
			return k
		}
		return k + float64(t) + *t2
	})
	internal.TestEqual(t, results, 0+5./1.+1+3./6.+3+3./1.)
}

type UintStringMap map[uint]string

func (u UintStringMap) xxx() bool { return true }

func TestFilterAnyAll(t *testing.T) {
	internal.TestPanic(t, true, func() { Filter(map[bool]bool{}, nil) })
	internal.TestEqual(t, Filter(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), map[bool]bool{})
	internal.TestEqual(t, Filter(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t2 }), map[int32]bool{0: true, 2: true})
	internal.TestEqual(t, Filter(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t >= 5 || t2 > 5 }), map[int32]int32{5: 9, 6: 8, 7: 7})
	internal.TestEqual(t, Filter(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), map[byte]string{'1': "1", '3': "3"})
	internal.TestEqual(t, Filter(UintStringMap{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t != 0 && len(t2) > 2 }), UintStringMap{1: "ccccc"})
	internal.TestEqual(t, Filter(UintStringMap{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return true }).xxx(), true)

	internal.TestPanic(t, true, func() { Any(map[bool]bool{}, nil) })
	internal.TestEqual(t, Any(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	internal.TestEqual(t, Any(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), true)
	internal.TestEqual(t, Any(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t < 0 && t2 }), false)
	internal.TestEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), true)
	internal.TestEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 && t2 < 5 }), false)
	internal.TestEqual(t, Any(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), true)
	internal.TestEqual(t, Any(map[byte]string{'1': "!", '@': "@", '3': "#"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), false)

	internal.TestPanic(t, true, func() { All(map[bool]bool{}, nil) })
	internal.TestEqual(t, All(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	internal.TestEqual(t, All(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), false)
	internal.TestEqual(t, All(map[int32]bool{0: true, 1: false, 2: true, 3: true}, func(t int32, t2 bool) bool { return t == 1 || t2 }), true)
	internal.TestEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), false)
	internal.TestEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 0 && t2 > 0 }), true)
	internal.TestEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return len(t2) > 2 }), false)
	internal.TestEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t == 2 || len(t2) >= 2 }), true)
}

func TestExpMaps(t *testing.T) {
	// From https://cs.opensource.google/go/x/exp/+/master:maps/maps_test.go

	var m1 = map[int]int{1: 2, 2: 4, 4: 8, 8: 16}
	var m2 = map[int]string{1: "2", 2: "4", 4: "8", 8: "16"}

	t.Run("TestEqual", func(t *testing.T) {
		internal.TestEqual(t, Equal(m1, m1), true)
		internal.TestEqual(t, Equal(m1, map[int]int(nil)), false)
		internal.TestEqual(t, Equal[map[int]int](nil, nil), true)
		internal.TestEqual(t, Equal(map[int]int(nil), m1), false)
		internal.TestEqual(t, Equal(m1, map[int]int{1: 2}), false)
		internal.TestEqual(t, Equal(map[int]string{1: "2"}, m2), false)
		internal.TestEqual(t, Equal(m2, m2), true)

		// Comparing NaN for equality is expected to fail.
		mf := map[int]float64{1: 0, 2: math.NaN()}
		internal.TestEqual(t, Equal(mf, mf), false)
	})

	t.Run("TestEqualWith", func(t *testing.T) {
		equal := func(v1, v2 int) bool { return v1 == v2 }
		equalIntStr := func(v1 int, v2 string) bool { return strconv.Itoa(v1) == v2 }
		equalFp := func(v1, v2 float64) bool { return v1 == v2 }
		equalNaN := func(v1, v2 float64) bool { return v1 == v2 || (math.IsNaN(v1) && math.IsNaN(v2)) }

		internal.TestEqual(t, EqualWith(m1, (map[int]int)(nil), equal), false)
		internal.TestEqual(t, EqualWith((map[int]int)(nil), m1, equal), false)
		internal.TestEqual(t, EqualWith[map[int]int, map[int]int](nil, nil, equal), true)
		internal.TestEqual(t, EqualWith(m1, m1, equal), true)
		internal.TestEqual(t, EqualWith(m1, map[int]int{1: 2}, equal), false)
		internal.TestEqual(t, EqualWith(m1, m2, equalIntStr), true)

		// Comparing NaN for equality is expected to fail, but it should succeed using equalNaN.
		mf := map[int]float64{1: 0, 2: math.NaN()}
		internal.TestEqual(t, EqualWith(mf, mf, equalFp), false)
		internal.TestEqual(t, EqualWith(mf, mf, equalNaN), true)
	})

	t.Run("TestClone", func(t *testing.T) {
		internal.TestEqual(t, Clone(map[int]int(nil)), map[int]int{})
		internal.TestEqual(t, Clone(map[int]int{}), map[int]int{})
		internal.TestEqual(t, Clone(map[int]string{0: "0"}), map[int]string{0: "0"})
		internal.TestEqual(t, Clone(m1), m1)
		internal.TestEqual(t, Clone(m2), m2)
	})

	t.Run("TestCopyTo", func(t *testing.T) {
		mc := make(map[int]int, 0)
		CopyTo(mc, m1)
		internal.TestEqual(t, mc, m1)
		mc = Clone(m1)
		CopyTo(mc, mc)
		internal.TestEqual(t, mc, m1)
		CopyTo(mc, map[int]int{16: 32})
		internal.TestEqual(t, mc, map[int]int{1: 2, 2: 4, 4: 8, 8: 16, 16: 32})
	})

	t.Run("TestClear", func(t *testing.T) {
		ml := map[int]int{1: 1, 2: 2, 3: 3}
		Clear(ml)
		internal.TestEqual(t, len(ml), 0)
		internal.TestEqual(t, ml, map[int]int{})
		internal.TestEqual(t, Equal(ml, map[int]int(nil)), true)
	})
}

func sorted[T xsugar.Ordered](slice []T) []T {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return slice
}

func sorted2[K xsugar.Ordered, V any, T []string | xtuple.Tuple[K, V]](slice []T) []T {
	sort.Slice(slice, func(i, j int) bool {
		var t T
		switch (any)(t).(type) {
		case []string:
			return (any)(slice[i]).([]string)[0] < (any)(slice[j]).([]string)[0]
		case xtuple.Tuple[K, V]:
			return (any)(slice[i]).(xtuple.Tuple[K, V]).Item1 < (any)(slice[j]).(xtuple.Tuple[K, V]).Item1
		}
		return false
	})
	return slice
}
