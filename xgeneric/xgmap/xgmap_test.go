//go:build go1.18
// +build go1.18

package xgmap

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xsugar"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xtuple"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"testing"
)

func TestKeysValues(t *testing.T) {
	xtestingEqual(t, Keys(map[int]uint{}), []int{})
	xtestingEqual(t, Values(map[int]uint{}), []uint{})
	xtestingEqual(t, KeyValues(map[int]uint{}), []xtuple.Tuple[int, uint]{})
	xtestingEqual(t, Keys(map[int]uint{1: 0}), []int{1})
	xtestingEqual(t, Values(map[int]uint{0: 1}), []uint{1})
	xtestingEqual(t, KeyValues(map[int]uint{0: 1}), []xtuple.Tuple[int, uint]{{0, 1}})
	xtestingEqual(t, sorted(Keys(map[string]int{"a": 1, "b": 2, "c": 3})), []string{"a", "b", "c"})
	xtestingEqual(t, sorted(Values(map[string]int{"a": 1, "b": 2, "c": 3})), []int{1, 2, 3})
	xtestingEqual(t, sorted2[string, int, xtuple.Tuple[string, int]](KeyValues(map[string]int{"a": 1, "b": 2, "c": 3})), []xtuple.Tuple[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
	xtestingEqual(t, sorted(Keys(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), []string{"Authorization", "X"})
	xtestingEqual(t, sorted2[string, []string, []string](Values(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), [][]string{{"!", "@"}, {"Token"}})
	xtestingEqual(t, sorted2[string, []string, xtuple.Tuple[string, []string]](KeyValues(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), []xtuple.Tuple[string, []string]{{"Authorization", []string{"Token"}}, {"X", []string{"!", "@"}}})
	xtestingEqual(t, sorted(Keys(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []uint32{1, 2, 3, 4})
	xtestingEqual(t, sorted(Values(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []float64{0.1, 0.2, 0.3, 0.4})
	xtestingEqual(t, sorted2[uint32, float64, xtuple.Tuple[uint32, float64]](KeyValues(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []xtuple.Tuple[uint32, float64]{{1, 0.1}, {2, 0.2}, {3, 0.3}, {4, 0.4}})

	xtestingPanic(t, true, func() { FromKeys[bool, bool]([]bool{}, nil) })
	xtestingPanic(t, true, func() { FromValues[bool, bool]([]bool{}, nil) })
	xtestingPanic(t, false, func() { FromKeyValues([]xtuple.Tuple[bool, bool]{}) })
	xtestingEqual(t, FromKeys([]int{}, func(i int, t int) uint { return uint(i) }), map[int]uint{})
	xtestingEqual(t, FromValues([]uint{}, func(i int, t uint) int { return i }), map[int]uint{})
	xtestingEqual(t, FromKeyValues([]xtuple.Tuple[int, uint]{}), map[int]uint{})
	xtestingEqual(t, FromKeys([]int{1}, func(i int, t int) uint { return uint(i) }), map[int]uint{1: 0})
	xtestingEqual(t, FromValues([]uint{1}, func(i int, t uint) int { return i }), map[int]uint{0: 1})
	xtestingEqual(t, FromKeyValues([]xtuple.Tuple[int, uint]{{0, 1}}), map[int]uint{0: 1})
	xtestingEqual(t, FromKeys([]string{"a", "b", "c"}, func(i int, t string) int { return i + 1 }), map[string]int{"a": 1, "b": 2, "c": 3})
	xtestingEqual(t, FromValues([]int{1, 2, 3}, func(i int, t int) byte { return byte(i) + 'a' }), map[byte]int{'a': 1, 'b': 2, 'c': 3})
	xtestingEqual(t, FromKeyValues([]xtuple.Tuple[byte, int]{{'a', 1}, {'b', 2}, {'c', 3}}), map[byte]int{'a': 1, 'b': 2, 'c': 3})
	xtestingEqual(t, FromKeys([]string{"!", "@", "#", "$", "%"}, func(i int, k string) bool { return false }), map[string]bool{"!": false, "@": false, "#": false, "$": false, "%": false})
	xtestingEqual(t, FromValues([]string{"!", "@", "#", "$", "%"}, func(i int, k string) int { return i }), map[int]string{0: "!", 1: "@", 2: "#", 3: "$", 4: "%"})
	xtestingEqual(t, FromKeyValues([]xtuple.Tuple[int, string]{{0, "!"}, {1, "@"}, {2, "#"}, {3, "$"}, {4, "%"}}), map[int]string{0: "!", 1: "@", 2: "#", 3: "$", 4: "%"})
}

func TestForeach(t *testing.T) {
	xtestingPanic(t, true, func() { Foreach(map[int]bool{}, nil) })
	xtestingPanic(t, false, func() { Foreach(map[int]bool{}, func(t int, t2 bool) {}) })

	test1 := 0
	Foreach(map[int]int{1: 2, 2: 3, 3: 4, 4: 5}, func(t int, t2 int) { test1 += t + t2 })
	xtestingEqual(t, test1, 1+2+2+3+3+4+4+5)
	test2 := uint(1)
	Foreach(map[uint]uint{1: 1, 2: 2, 3: 3, 4: 4}, func(t uint, t2 uint) { test2 *= t * t2 })
	xtestingEqual(t, test2, uint(1*1*2*2*3*3*4*4))
	test3 := float32(0)
	Foreach(map[float32]float32{1.0: 0.0, 2.0: 1.0, 3.0: 2.0}, func(t float32, t2 float32) { test3 -= t - t2 })
	xtestingEqual(t, test3, float32(-1.0-1.0-1.0))
	test4 := make([]string, 0)
	Foreach(map[string]string{"1": "a", "2": "b", "3": "c"}, func(t string, t2 string) { test4 = append(test4, t+t2) })
	xtestingEqual(t, sorted(test4), []string{"1a", "2b", "3c"})
}

func TestMapExpand(t *testing.T) {
	xtestingPanic(t, true, func() { Map[bool, bool, bool, bool](map[bool]bool{}, nil) })
	xtestingEqual(t, Map(map[bool]bool{}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{})
	xtestingEqual(t, Map(map[bool]bool{true: false, false: true}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{false: false, true: true})
	xtestingEqual(t, Map(map[string]int{"!": 1, "@": 2, "#": 3}, func(t string, t2 int) (int32, string) { return int32(t2) + 1, t }), map[int32]string{2: "!", 3: "@", 4: "#"})
	xtestingEqual(t, Map(map[float64]bool{1.: true, 2.: true}, func(t float64, t2 bool) (string, bool) { return strconv.FormatFloat(t, 'f', 1, 64), !t2 }), map[string]bool{"1.0": false, "2.0": false})
	xtestingEqual(t, Map(map[string]string{"1": "", "@": "", "3": "", "4": ""}, func(t string, t2 string) (int32, bool) {
		s, err := strconv.Atoi(t)
		return int32(s), err == nil
	}), map[int32]bool{1: true, 0: false, 3: true, 4: true})

	xtestingPanic(t, true, func() { Expand[bool, bool, bool, bool](map[bool]bool{}, nil) })
	xtestingEqual(t, Expand(map[bool]bool{}, func(t bool, t2 bool) []xtuple.Tuple[bool, bool] { return []xtuple.Tuple[bool, bool]{{!t, !t}} }), map[bool]bool{})
	xtestingEqual(t, Expand(map[bool]bool{true: false, false: true}, func(t bool, t2 bool) []xtuple.Tuple[bool, bool] { return []xtuple.Tuple[bool, bool]{{!t, !t}, {t, t}} }), map[bool]bool{false: false, true: true})
	xtestingEqual(t, Expand(map[string]int{"!": 1, "@": 2, "#": 3}, func(t string, t2 int) []xtuple.Tuple[int32, string] {
		return []xtuple.Tuple[int32, string]{{int32(t2), t}, {int32(t2) + 10, t + " "}}
	}), map[int32]string{1: "!", 11: "! ", 2: "@", 12: "@ ", 3: "#", 13: "# "})
	xtestingEqual(t, Expand(map[float64]bool{1.: true, 2.: true}, func(t float64, t2 bool) []xtuple.Tuple[string, bool] { return nil }), map[string]bool{})
	xtestingEqual(t, Expand(map[string]string{"1": "", "@": "", "3": "", "4": ""}, func(t string, t2 string) []xtuple.Tuple[int32, bool] {
		s, err := strconv.Atoi(t)
		return []xtuple.Tuple[int32, bool]{{int32(s), err == nil}, {int32(s) + 10, err != nil}}
	}), map[int32]bool{1: true, 0: false, 3: true, 4: true, 11: false, 10: true, 13: false, 14: false})
}

func TestReduce(t *testing.T) {
	xtestingPanic(t, true, func() { Reduce(map[int]bool{}, true, nil) })
	xtestingEqual(t, Reduce(map[int]bool{}, true, func(k bool, t int, t2 bool) bool { return false }), true)
	xtestingEqual(t, sorted(Reduce(map[string]int{"1": 1, "2": 2, "3": 3}, []string{"0"}, func(k []string, t string, t2 int) []string { return append(k, t+strconv.Itoa(t2)) })), []string{"0", "11", "22", "33"})
	xtestingEqual(t, Reduce(map[int]float64{9: 0.5, 8: 1.5, 7: 2.5, 6: 3.5}, 1., func(k float64, t int, t2 float64) float64 { return k + float64(t)*t2 }), 1+9*0.5+8*1.5+7*2.5+6*3.5)

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
	xtestingEqual(t, results, 0+5./1.+1+3./6.+3+3./1.)
}

type UintStringMap map[uint]string

func (u UintStringMap) xxx() bool { return true }

func TestFilterAnyAll(t *testing.T) {
	xtestingPanic(t, true, func() { Filter(map[bool]bool{}, nil) })
	xtestingEqual(t, Filter(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), map[bool]bool{})
	xtestingEqual(t, Filter(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t2 }), map[int32]bool{0: true, 2: true})
	xtestingEqual(t, Filter(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t >= 5 || t2 > 5 }), map[int32]int32{5: 9, 6: 8, 7: 7})
	xtestingEqual(t, Filter(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), map[byte]string{'1': "1", '3': "3"})
	xtestingEqual(t, Filter(UintStringMap{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t != 0 && len(t2) > 2 }), UintStringMap{1: "ccccc"})
	xtestingEqual(t, Filter(UintStringMap{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return true }).xxx(), true)

	xtestingPanic(t, true, func() { Any(map[bool]bool{}, nil) })
	xtestingEqual(t, Any(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	xtestingEqual(t, Any(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), true)
	xtestingEqual(t, Any(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t < 0 && t2 }), false)
	xtestingEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), true)
	xtestingEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 && t2 < 5 }), false)
	xtestingEqual(t, Any(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), true)
	xtestingEqual(t, Any(map[byte]string{'1': "!", '@': "@", '3': "#"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), false)

	xtestingPanic(t, true, func() { All(map[bool]bool{}, nil) })
	xtestingEqual(t, All(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	xtestingEqual(t, All(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), false)
	xtestingEqual(t, All(map[int32]bool{0: true, 1: false, 2: true, 3: true}, func(t int32, t2 bool) bool { return t == 1 || t2 }), true)
	xtestingEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), false)
	xtestingEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 0 && t2 > 0 }), true)
	xtestingEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return len(t2) > 2 }), false)
	xtestingEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t == 2 || len(t2) >= 2 }), true)
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

// =============================
// simplified xtesting functions
// =============================

func failTest(t testing.TB, failureMessage string) bool {
	_, file, line, _ := runtime.Caller(2)
	_, _ = fmt.Fprintf(os.Stderr, "%s:%d %s\n", path.Base(file), line, failureMessage)
	t.Fail()
	return false
}

func xtestingEqual(t testing.TB, give, want interface{}) bool {
	if give != nil && want != nil && (reflect.TypeOf(give).Kind() == reflect.Func || reflect.TypeOf(want).Kind() == reflect.Func) {
		return failTest(t, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (cannot take func type as argument)", give, want))
	}
	if !reflect.DeepEqual(give, want) {
		return failTest(t, fmt.Sprintf("Equal: expect to be `%#v`, but actually was `%#v`", want, give))
	}
	return true
}

func xtestingPanic(t *testing.T, want bool, f func(), v ...any) bool {
	isPanic, value := false, interface{}(nil)
	func() { defer func() { value = recover(); isPanic = value != nil }(); f() }()
	if want && !isPanic {
		return failTest(t, fmt.Sprintf("Panic: expect function `%#v` to panic, but actually did not panic", interface{}(f)))
	}
	if want && isPanic && len(v) > 0 && v[0] != nil && !reflect.DeepEqual(value, v[0]) {
		return failTest(t, fmt.Sprintf("PanicWithValue: expect function `%#v` to panic with `%#v`, but actually with `%#v`", interface{}(f), want, value))
	}
	if !want && isPanic {
		return failTest(t, fmt.Sprintf("NotPanic: expect function `%#v` not to panic, but actually panicked with `%v`", interface{}(f), value))
	}
	return true
}
