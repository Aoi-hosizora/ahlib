//go:build go1.18
// +build go1.18

package xslice

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	for _, tc := range []struct {
		give []int
	}{
		{[]int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
	} {
		me := make([]int, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			result := Shuffle(tc.give)
			xtestingEqual(t, tc.give, me)
			xtestingEqual(t, ElementMatch(result, me), true)
			fmt.Println(result)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			ShuffleSelf(tc.give)
			xtestingEqual(t, ElementMatch(tc.give, me), true)
			fmt.Println(tc.give)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		give []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	} {
		me := make([]int, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		result := Reverse(tc.give)
		xtestingEqual(t, tc.give, me)
		xtestingEqual(t, result, tc.want)
		ReverseSelf(tc.give)
		xtestingEqual(t, tc.give, tc.want)
	}
}

func TestSort(t *testing.T) {
	le := func(i, j int) bool { return i < j }
	for _, tc := range []struct {
		give     []int
		giveLess Lesser[int]
		want     []int
	}{
		{[]int{}, le, []int{}},
		{[]int{0}, le, []int{0}},
		{[]int{1, 1, 1}, le, []int{1, 1, 1}},
		{[]int{4, 3, 2, 1}, le, []int{1, 2, 3, 4}},
		{[]int{8, 1, 6, 8, 1, 2}, le, []int{1, 1, 2, 6, 8, 8}},
	} {
		me := make([]int, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		result := Sort(tc.give)
		xtestingEqual(t, tc.give, me)
		xtestingEqual(t, result, tc.want)
		SortSelf(tc.give)
		xtestingEqual(t, tc.give, tc.want)

		tc.give = me
		result = SortWith(tc.give, tc.giveLess)
		xtestingEqual(t, tc.give, me)
		xtestingEqual(t, result, tc.want)
		SortSelfWith(tc.give, tc.giveLess)
		xtestingEqual(t, tc.give, tc.want)
	}
}

func TestStableSort(t *testing.T) {
	le := func(i, j int) bool { return i < j }
	for _, tc := range []struct {
		give     []int
		giveLess Lesser[int]
		want     []int
	}{
		{[]int{}, le, []int{}},
		{[]int{0}, le, []int{0}},
		{[]int{1, 1, 1}, le, []int{1, 1, 1}},
		{[]int{4, 3, 2, 1}, le, []int{1, 2, 3, 4}},
		{[]int{8, 1, 6, 8, 1, 2}, le, []int{1, 1, 2, 6, 8, 8}},
	} {
		me := make([]int, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		result := StableSort(tc.give)
		xtestingEqual(t, tc.give, me)
		xtestingEqual(t, result, tc.want)
		StableSortSelf(tc.give)
		xtestingEqual(t, tc.give, tc.want)
	}

	type tuple = [2]int
	new1 := func(v int) tuple { return tuple{v, 0} }
	new2 := func(v, o int) tuple { return tuple{v, o} }
	le2 := func(i, j tuple) bool { return i[0] < j[0] }
	for _, tc := range []struct {
		give     []tuple
		giveLess Lesser[tuple]
		want     []tuple
	}{
		{[]tuple{}, le2, []tuple{}},
		{[]tuple{new1(0)}, le2, []tuple{new1(0)}},
		{[]tuple{new2(1, 3), new2(1, 2), new2(1, 1)}, le2,
			[]tuple{new2(1, 3), new2(1, 2), new2(1, 1)}},
		{[]tuple{new1(4), new1(3), new1(2), new1(1)}, le2,
			[]tuple{new1(1), new1(2), new1(3), new1(4)}},
		{[]tuple{new2(8, 2), new2(1, 2), new1(6), new2(8, 1), new2(1, 1), new1(2)}, le2,
			[]tuple{new2(1, 2), new2(1, 1), new1(2), new1(6), new2(8, 2), new2(8, 1)}},
	} {
		me := make([]tuple, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		result := StableSortWith(tc.give, tc.giveLess)
		xtestingEqual(t, tc.give, me)
		xtestingEqual(t, result, tc.want)
		StableSortSelfWith(tc.give, tc.giveLess)
		xtestingEqual(t, tc.give, tc.want)
	}
}

type testStruct struct {
	value int
}

func (t testStruct) String() string {
	return strconv.Itoa(t.value)
}

func newTestStruct(value int) testStruct {
	return testStruct{value: value}
}

func newTestSlice(s []int) []testStruct {
	newSlice := make([]testStruct, len(s))
	for idx, item := range s {
		newSlice[idx] = newTestStruct(item)
	}
	return newSlice
}

func testToIntSlice(s interface{}) []int {
	out := make([]int, len(s.([]testStruct)))
	for idx, item := range s.([]testStruct) {
		out[idx] = item.value
	}
	return out
}

func TestIndexOf(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 2, 3}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give      []int
		giveValue int
		want      int
	}{
		{[]int{}, 0, -1},
		{s2, -1, -1},
		{s2, 0, -1},
		{s2, 1, 0},
		{s2, 2, 2},
		{s2, 3, 5},
		{s2, 4, -1},
		{s2, 5, 1},
		{s2, 6, -1},
	} {
		xtestingEqual(t, IndexOf(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
		xtestingEqual(t, IndexOfWith(give, giveValue, eq), tc.want)
	}
}

func TestContains(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 2, 3}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give      []int
		giveValue int
		want      bool
	}{
		{[]int{}, 0, false},
		{s2, -1, false},
		{s2, 0, false},
		{s2, 1, true},
		{s2, 2, true},
		{s2, 3, true},
		{s2, 4, false},
		{s2, 5, true},
		{s2, 6, false},
	} {
		xtestingEqual(t, Contains(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
		xtestingEqual(t, ContainsWith(give, giveValue, eq), tc.want)
	}
}

func TestCount(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give      []int
		giveValue int
		want      int
	}{
		{[]int{}, 0, 0},
		{s2, -1, 0},
		{s2, 0, 0},
		{s2, 1, 2},
		{s2, 2, 3},
		{s2, 3, 1},
		{s2, 4, 0},
		{s2, 5, 2},
		{s2, 6, 1},
		{s2, 7, 0},
	} {
		xtestingEqual(t, Count(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
		xtestingEqual(t, CountWith(give, giveValue, eq), tc.want)
	}
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		give      []int
		giveValue int
		giveIndex int
		want      []int
	}{
		{[]int{}, 9, -2, []int{9}},
		{[]int{}, 9, -1, []int{9}},
		{[]int{}, 9, 0, []int{9}},
		{[]int{}, 9, 1, []int{9}},
		{[]int{1}, 9, -1, []int{9, 1}},
		{[]int{1}, 9, 0, []int{9, 1}},
		{[]int{1}, 9, 1, []int{1, 9}},
		{[]int{1}, 9, 2, []int{1, 9}},
		{[]int{1, 2}, 9, -1, []int{9, 1, 2}},
		{[]int{1, 2}, 9, 0, []int{9, 1, 2}},
		{[]int{1, 2}, 9, 1, []int{1, 9, 2}},
		{[]int{1, 2}, 9, 2, []int{1, 2, 9}},
		{[]int{1, 2}, 9, 3, []int{1, 2, 9}},
		{[]int{1, 2, 3}, 9, -1, []int{9, 1, 2, 3}},
		{[]int{1, 2, 3}, 9, 0, []int{9, 1, 2, 3}},
		{[]int{1, 2, 3}, 9, 1, []int{1, 9, 2, 3}},
		{[]int{1, 2, 3}, 9, 2, []int{1, 2, 9, 3}},
		{[]int{1, 2, 3}, 9, 3, []int{1, 2, 3, 9}},
		{[]int{1, 2, 3}, 9, 4, []int{1, 2, 3, 9}},
		{[]int{1, 2, 3, 4}, 9, -1, []int{9, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 9, 0, []int{9, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 9, 1, []int{1, 9, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 9, 2, []int{1, 2, 9, 3, 4}},
		{[]int{1, 2, 3, 4}, 9, 3, []int{1, 2, 3, 9, 4}},
		{[]int{1, 2, 3, 4}, 9, 4, []int{1, 2, 3, 4, 9}},
		{[]int{1, 2, 3, 4}, 9, 5, []int{1, 2, 3, 4, 9}},
	} {
		xtestingEqual(t, Insert(tc.give, tc.giveValue, tc.giveIndex), tc.want)
	}
}

func TestDelete(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		name      string
		give      []int
		giveValue int
		giveN     int
		want      []int
	}{
		{"", []int{}, 0, 1, []int{}},
		{"", s2, -1, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{"", s2, 0, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{"", s2, 1, 1, []int{5, 2, 1, 5, 2, 6, 3, 2}},
		{"", s2, 1, 2, []int{5, 2, 5, 2, 6, 3, 2}},
		{"", s2, 2, 1, []int{1, 5, 1, 5, 2, 6, 3, 2}},
		{"", s2, 2, 2, []int{1, 5, 1, 5, 6, 3, 2}},
		{"", s2, 2, -1, []int{1, 5, 1, 5, 6, 3}},
		{"", s2, 3, 1, []int{1, 5, 2, 1, 5, 2, 6, 2}},
		{"", s2, 4, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{"", s2, 5, 1, []int{1, 2, 1, 5, 2, 6, 3, 2}},
		{"", s2, 6, 1, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{"", s2, 7, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			xtestingEqual(t, Delete(tc.give, tc.giveValue, tc.giveN), tc.want)
			give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
			xtestingEqual(t, testToIntSlice(DeleteWith(give, giveValue, tc.giveN, eq)), tc.want)
		})
	}
}

func TestDeleteAll(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give      []int
		giveValue int
		want      []int
	}{
		{[]int{}, 0, []int{}},
		{s2, -1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 0, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 1, []int{5, 2, 5, 2, 6, 3, 2}},
		{s2, 2, []int{1, 5, 1, 5, 6, 3}},
		{s2, 3, []int{1, 5, 2, 1, 5, 2, 6, 2}},
		{s2, 4, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 5, []int{1, 2, 1, 2, 6, 3, 2}},
		{s2, 6, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{s2, 7, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtestingEqual(t, DeleteAll(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
		xtestingEqual(t, testToIntSlice(DeleteAllWith(give, giveValue, eq)), tc.want)
	}
}

func TestDiff(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  []int
	}{
		{[]int{}, []int{}, []int{}},
		{s2, []int{}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, []int{1}, []int{5, 2, 5, 2, 6, 3, 2}},
		{s2, []int{1, 2}, []int{5, 5, 6, 3}},
		{s2, []int{1, 2, 3}, []int{5, 5, 6}},
		{s2, []int{1, 2, 3, 4}, []int{5, 5, 6}},
		{s2, []int{1, 2, 3, 4, 5}, []int{6}},
		{s2, []int{1, 2, 3, 4, 5, 6}, []int{}},
		{s2, []int{6}, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{s2, []int{6, 5}, []int{1, 2, 1, 2, 3, 2}},
		{s2, []int{6, 5, 4}, []int{1, 2, 1, 2, 3, 2}},
		{s2, []int{6, 5, 4, 3}, []int{1, 2, 1, 2, 2}},
		{s2, []int{6, 5, 4, 3, 2}, []int{1, 1}},
		{s2, []int{6, 5, 4, 3, 2, 1}, []int{}},
	} {
		xtestingEqual(t, Diff(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice(tc.give1), newTestSlice(tc.give2)
		xtestingEqual(t, testToIntSlice(DiffWith(give1, give2, eq)), tc.want)
	}
}

func TestUnion(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  []int
	}{
		{[]int{}, []int{}, []int{}},
		{s2, []int{}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, []int{11}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11}},
		{s2, []int{11, 2}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11}},
		{s2, []int{11, 2, 13}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13}},
		{s2, []int{11, 2, 13, 14}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14}},
		{s2, []int{11, 2, 13, 14, 5}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14}},
		{s2, []int{11, 2, 13, 14, 5, 16}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14, 16}},
	} {
		xtestingEqual(t, Union(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice(tc.give1), newTestSlice(tc.give2)
		xtestingEqual(t, testToIntSlice(UnionWith(give1, give2, eq)), tc.want)
	}
}

func TestIntersect(t *testing.T) {
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  []int
	}{
		{[]int{}, []int{}, []int{}},
		{s2, []int{}, []int{}},
		{s2, []int{1}, []int{1, 1}},
		{s2, []int{1, 2}, []int{1, 2, 1, 2, 2}},
		{s2, []int{1, 2, 3}, []int{1, 2, 1, 2, 3, 2}},
		{s2, []int{1, 2, 3, 4}, []int{1, 2, 1, 2, 3, 2}},
		{s2, []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{s2, []int{1, 2, 3, 4, 5, 6}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtestingEqual(t, Intersect(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice(tc.give1), newTestSlice(tc.give2)
		xtestingEqual(t, testToIntSlice(IntersectWith(give1, give2, eq)), tc.want)
	}
}

func TestDeduplicate(t *testing.T) {
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{2, 1, 1}, []int{2, 1}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, []int{1, 5, 2, 6, 3}},
	} {
		xtestingEqual(t, Deduplicate(tc.give), tc.want)
		give := newTestSlice(tc.give)
		xtestingEqual(t, testToIntSlice(DeduplicateWith(give, eq)), tc.want)
	}
}

func TestElementMatch(t *testing.T) {
	eq := func(i, j testStruct) bool { return i.value == j.value }

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  bool
	}{
		{[]int{}, []int{}, true},
		{[]int{0}, []int{}, false},
		{[]int{0}, []int{0}, true},
		{[]int{1, 1, 1}, []int{1}, false},
		{[]int{1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{1, 2, 1}, []int{1, 2, 1}, true},
		{[]int{1, 2, 1}, []int{1, 2, 2}, false},
	} {
		xtestingEqual(t, ElementMatch(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice(tc.give1), newTestSlice(tc.give2)
		xtestingEqual(t, ElementMatchWith(give1, give2, eq), tc.want)
	}
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
