package xmap

import (
	"constraints"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xtuple"
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
	xtestingEqual(t, Keys(map[int]uint{1: 0}), []int{1})
	xtestingEqual(t, Values(map[int]uint{0: 1}), []uint{1})
	xtestingEqual(t, sorted(Keys(map[string]int{"a": 1, "b": 2, "c": 3})), []string{"a", "b", "c"})
	xtestingEqual(t, sorted(Values(map[string]int{"a": 1, "b": 2, "c": 3})), []int{1, 2, 3})
	xtestingEqual(t, sorted(Keys(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}})), []string{"Authorization", "X"})
	xtestingEqual(t, Values(map[string][]string{"Authorization": {"Token"}, "X": {"!", "@"}}), [][]string{{"Token"}, {"!", "@"}})
	xtestingEqual(t, sorted(Keys(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []uint32{1, 2, 3, 4})
	xtestingEqual(t, sorted(Values(map[uint32]float64{1: 0.1, 2: 0.2, 3: 0.3, 4: 0.4})), []float64{0.1, 0.2, 0.3, 0.4})

	xtestingPanic(t, func() { FromKeys[bool, bool]([]bool{}, nil) }, true)
	xtestingPanic(t, func() { FromValues[bool, bool]([]bool{}, nil) }, true)
	xtestingEqual(t, FromKeys([]int{}, func(i int, t int) uint { return uint(i) }), map[int]uint{})
	xtestingEqual(t, FromValues([]uint{}, func(i int, t uint) int { return i }), map[int]uint{})
	xtestingEqual(t, FromKeys([]int{1}, func(i int, t int) uint { return uint(i) }), map[int]uint{1: 0})
	xtestingEqual(t, FromValues([]uint{1}, func(i int, t uint) int { return i }), map[int]uint{0: 1})
	xtestingEqual(t, FromKeys([]string{"a", "b", "c"}, func(i int, t string) int { return i + 1 }), map[string]int{"a": 1, "b": 2, "c": 3})
	xtestingEqual(t, FromValues([]int{1, 2, 3}, func(i int, t int) byte { return byte(i) + 'a' }), map[byte]int{'a': 1, 'b': 2, 'c': 3})
	xtestingEqual(t, FromKeys([]string{"!", "@", "#", "$", "%"}, func(i int, k string) bool { return false }), map[string]bool{"!": false, "@": false, "#": false, "$": false, "%": false})
	xtestingEqual(t, FromValues([]string{"!", "@", "#", "$", "%"}, func(i int, k string) int { return i }), map[int]string{0: "!", 1: "@", 2: "#", 3: "$", 4: "%"})
}

func TestForeach(t *testing.T) {
	xtestingPanic(t, func() { Foreach(map[int]bool{}, nil) }, true)
	xtestingPanic(t, func() { Foreach(map[int]bool{}, func(t int, t2 bool) {}) }, false)

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
	xtestingPanic(t, func() { Map[bool, bool, bool, bool](map[bool]bool{}, nil) }, true)
	xtestingEqual(t, Map(map[bool]bool{}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{})
	xtestingEqual(t, Map(map[bool]bool{true: false, false: true}, func(t bool, t2 bool) (bool, bool) { return !t, !t }), map[bool]bool{false: false, true: true})
	xtestingEqual(t, Map(map[string]int{"!": 1, "@": 2, "#": 3}, func(t string, t2 int) (int32, string) { return int32(t2) + 1, t }), map[int32]string{2: "!", 3: "@", 4: "#"})
	xtestingEqual(t, Map(map[float64]bool{1.: true, 2.: true}, func(t float64, t2 bool) (string, bool) { return strconv.FormatFloat(t, 'f', 1, 64), !t2 }), map[string]bool{"1.0": false, "2.0": false})
	xtestingEqual(t, Map(map[string]string{"1": "", "@": "", "3": "", "4": ""}, func(t string, t2 string) (int32, bool) {
		s, err := strconv.Atoi(t)
		return int32(s), err == nil
	}), map[int32]bool{1: true, 0: false, 3: true, 4: true})

	xtestingPanic(t, func() { Expand[bool, bool, bool, bool](map[bool]bool{}, nil) }, true)
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
	xtestingPanic(t, func() { Reduce(map[int]bool{}, true, nil) }, true)
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

func TestFilterAnyAll(t *testing.T) {
	xtestingPanic(t, func() { Filter(map[bool]bool{}, nil) }, true)
	xtestingEqual(t, Filter(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), map[bool]bool{})
	xtestingEqual(t, Filter(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t2 }), map[int32]bool{0: true, 2: true})
	xtestingEqual(t, Filter(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t >= 5 || t2 > 5 }), map[int32]int32{5: 9, 6: 8, 7: 7})
	xtestingEqual(t, Filter(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), map[byte]string{'1': "1", '3': "3"})
	xtestingEqual(t, Filter(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t != 0 && len(t2) > 2 }), map[uint]string{1: "ccccc"})

	xtestingPanic(t, func() { Any(map[bool]bool{}, nil) }, true)
	xtestingEqual(t, Any(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	xtestingEqual(t, Any(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), true)
	xtestingEqual(t, Any(map[int32]bool{0: true, 1: false, 2: true, 3: false}, func(t int32, t2 bool) bool { return t < 0 && t2 }), false)
	xtestingEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), true)
	xtestingEqual(t, Any(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 && t2 < 5 }), false)
	xtestingEqual(t, Any(map[byte]string{'1': "1", '@': "@", '3': "3"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), true)
	xtestingEqual(t, Any(map[byte]string{'1': "!", '@': "@", '3': "#"}, func(t byte, t2 string) bool { _, err := strconv.Atoi(t2); return err == nil }), false)

	xtestingPanic(t, func() { All(map[bool]bool{}, nil) }, true)
	xtestingEqual(t, All(map[bool]bool{}, func(t bool, t2 bool) bool { return t }), true)
	xtestingEqual(t, All(map[uint32]bool{0: true, 1: false, 2: true, 3: false}, func(t uint32, t2 bool) bool { return t2 }), false)
	xtestingEqual(t, All(map[int32]bool{0: true, 1: false, 2: true, 3: true}, func(t int32, t2 bool) bool { return t == 1 || t2 }), true)
	xtestingEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 5 }), false)
	xtestingEqual(t, All(map[int32]int32{5: 9, 4: 1, 6: 8, 3: 2, 7: 7}, func(t int32, t2 int32) bool { return t > 0 && t2 > 0 }), true)
	xtestingEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return len(t2) > 2 }), false)
	xtestingEqual(t, All(map[uint]string{0: "aaa", 2: "b", 1: "ccccc", 3: "dd"}, func(t uint, t2 string) bool { return t == 2 || len(t2) >= 2 }), true)
}

func sorted[T constraints.Ordered](slice []T) []T {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return slice
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

func xtestingPanic(t testing.TB, f func(), want bool) bool {
	didPanic := false
	func() { defer func() { didPanic = recover() != nil }(); f() }()
	if want && !didPanic {
		return failTest(t, fmt.Sprintf("Panic: function (%p) is expected to panic, actual does not panic", f))
	}
	if !want && didPanic {
		return failTest(t, fmt.Sprintf("NotPanic: function (%p) is expected not to panic, acutal panic", f))
	}
	return true
}
