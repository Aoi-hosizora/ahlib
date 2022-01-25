//go:build go1.18
// +build go1.18

package xslice

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xtuple"
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

func TestLastIndexOf(t *testing.T) {
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
		{s2, 1, 3},
		{s2, 2, 4},
		{s2, 3, 5},
		{s2, 4, -1},
		{s2, 5, 1},
		{s2, 6, -1},
	} {
		xtestingEqual(t, LastIndexOf(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice(tc.give), newTestStruct(tc.giveValue)
		xtestingEqual(t, LastIndexOfWith(give, giveValue, eq), tc.want)
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

type testError struct{ m string }

func (t *testError) Error() string { return t.m }

func TestRepeat(t *testing.T) {
	for _, tc := range []struct {
		give interface{}
		want interface{}
	}{
		{Repeat[interface{}](nil, 0), []interface{}{}},
		{Repeat[interface{}](nil, 2), []interface{}{nil, nil}},
		{Repeat(true, 0), []bool{}},
		{Repeat(true, 1), []bool{true}},
		{Repeat(5, 5), []int{5, 5, 5, 5, 5}},
		{Repeat("", 5), []string{"", "", "", "", ""}},
		{Repeat(uint(0), 2), []uint{uint(0), uint(0)}},
		{Repeat([]float64{1.1, 2.2}, 3), [][]float64{{1.1, 2.2}, {1.1, 2.2}, {1.1, 2.2}}},
		{Repeat(error(nil), 2), []error{nil, nil}},                                           // <<<
		{Repeat(error((*testError)(nil)), 2), []error{(*testError)(nil), (*testError)(nil)}}, // <<<
		{Repeat((*testError)(nil), 2), []*testError{(*testError)(nil), (*testError)(nil)}},
		{Repeat(&testError{"test"}, 2), []*testError{{"test"}, {"test"}}},
	} {
		xtestingEqual(t, tc.give, tc.want)
	}
}

func TestForeach(t *testing.T) {
	xtestingPanic(t, func() { Foreach([]bool{}, nil) }, true)
	xtestingPanic(t, func() { Foreach([]bool{}, func(t bool) {}) }, false)

	test1 := 0
	Foreach([]int{1, 2, 3, 4}, func(t int) { test1 += t })
	xtestingEqual(t, test1, 1+2+3+4)
	test2 := uint(1)
	Foreach([]uint{1, 2, 3, 4, 5}, func(t uint) { test2 *= t })
	xtestingEqual(t, test2, uint(1*2*3*4*5))
	test3 := float32(0)
	Foreach([]float32{1.0, 2.0, 3.0}, func(t float32) { test3 -= t })
	xtestingEqual(t, test3, float32(-1.0-2.0-3.0))
	test4 := ""
	Foreach([]string{"1", "2", "3", "a", "b", "c"}, func(t string) { test4 += t })
	xtestingEqual(t, test4, "123abc")
}

func TestMapExpand(t *testing.T) {
	xtestingPanic(t, func() { Map[bool, bool]([]bool{}, nil) }, true)
	xtestingEqual(t, Map([]bool{}, func(t bool) bool { return false }), []bool{})
	xtestingEqual(t, Map([]bool{true, true, true, true}, func(t bool) bool { return false }), []bool{false, false, false, false})
	xtestingEqual(t, Map([]int{1, 2, 3}, func(t int) int32 { return int32(t) + 1 }), []int32{2, 3, 4})
	xtestingEqual(t, Map([]float64{1., 2.}, func(t float64) string { return strconv.FormatFloat(t, 'f', 1, 64) }), []string{"1.0", "2.0"})
	xtestingEqual(t, Map([]string{"1", "@", "3", "4"}, func(t string) int32 {
		s, _ := strconv.Atoi(t)
		return int32(s)
	}), []int32{1, 0, 3, 4})

	xtestingPanic(t, func() { Expand[bool, bool]([]bool{}, nil) }, true)
	xtestingEqual(t, Expand([]bool{}, func(t bool) []bool { return []bool{} }), []bool{})
	xtestingEqual(t, Expand([]bool{true, false}, func(t bool) []bool { return []bool{t, !t} }), []bool{true, false, false, true})
	xtestingEqual(t, Expand([]int{1, 2, 3}, func(t int) []int32 { return []int32{int32(t), int32(t) + 1} }), []int32{1, 2, 2, 3, 3, 4})
	xtestingEqual(t, Expand([]float64{1., 2.}, func(t float64) []string { return []string{"", strconv.FormatFloat(t, 'f', 1, 64), "|"} }), []string{"", "1.0", "|", "", "2.0", "|"})
	xtestingEqual(t, Expand([]string{"1", "@", "3", "4"}, func(t string) []int32 {
		s, _ := strconv.Atoi(t)
		return []int32{int32(s), int32(s) * 2}
	}), []int32{1, 2, 0, 0, 3, 6, 4, 8})
}

func TestReduce(t *testing.T) {
	xtestingPanic(t, func() { Reduce([]bool{}, true, nil) }, true)
	xtestingEqual(t, Reduce([]bool{}, true, func(k bool, t bool) bool { return false }), true)
	xtestingEqual(t, Reduce([]int{1, 2, 3}, 0.0, func(k float64, t int) float64 { return k + float64(t) }), 1.+2.+3.)
	xtestingEqual(t, Reduce([]string{"a", "b", "c", "d"}, "0", func(k string, t string) string { return k + t }), "0abcd")
	xtestingEqual(t, Reduce([]string{"a", "b", "c", "d"}, []string{}, func(k []string, t string) []string { return append(k, string(t[0]+1)) }), []string{"b", "c", "d", "e"})

	fractions := [][2]int{{5, 1}, {3, 6}, {2, 0}, {3, 1}}
	results := Reduce(Map(fractions, func(t [2]int) *float64 {
		if t[1] == 0 {
			return nil
		}
		r := float64(t[0]) / float64(t[1])
		return &r
	}), 0.0, func(k float64, t *float64) float64 {
		if t == nil {
			return k
		}
		return k + *t
	})
	xtestingEqual(t, results, 5./1.+3./6.+3./1.)
}

func TestFilterAnyAll(t *testing.T) {
	xtestingPanic(t, func() { Filter([]bool{}, nil) }, true)
	xtestingEqual(t, Filter([]bool{}, func(t bool) bool { return t }), []bool{})
	xtestingEqual(t, Filter([]bool{true, false, true, false}, func(t bool) bool { return t }), []bool{true, true})
	xtestingEqual(t, Filter([]int32{9, 1, 8, 2, 7, 3, 6, 4, 5}, func(t int32) bool { return t > 5 }), []int32{9, 8, 7, 6})
	xtestingEqual(t, Filter([]string{"1", "@", "3"}, func(t string) bool { _, err := strconv.Atoi(t); return err == nil }), []string{"1", "3"})
	xtestingEqual(t, Filter([]string{"aaa", "b", "ccccc", "dd"}, func(t string) bool { return len(t) > 2 }), []string{"aaa", "ccccc"})

	xtestingPanic(t, func() { Any([]bool{}, nil) }, true)
	xtestingEqual(t, Any([]bool{}, func(t bool) bool { return t }), true)
	xtestingEqual(t, Any([]bool{true, false, true, false}, func(t bool) bool { return t }), true)
	xtestingEqual(t, Any([]bool{false, false}, func(t bool) bool { return t }), false)
	xtestingEqual(t, Any([]int32{9, 1, 8, 2, 7, 3, 6, 4, 5}, func(t int32) bool { return t > 5 }), true)
	xtestingEqual(t, Any([]int32{9, 1, 8, 2, 7, 3, 6, 4, 5}, func(t int32) bool { return t > 10 }), false)
	xtestingEqual(t, Any([]string{"1", "@", "3"}, func(t string) bool { _, err := strconv.Atoi(t); return err == nil }), true)
	xtestingEqual(t, Any([]string{"!", "@", "#"}, func(t string) bool { _, err := strconv.Atoi(t); return err == nil }), false)

	xtestingPanic(t, func() { All([]bool{}, nil) }, true)
	xtestingEqual(t, All([]bool{}, func(t bool) bool { return t }), true)
	xtestingEqual(t, All([]bool{true, false, true, false}, func(t bool) bool { return t }), false)
	xtestingEqual(t, All([]bool{true, true}, func(t bool) bool { return t }), true)
	xtestingEqual(t, All([]int32{9, 1, 8, 2, 7, 3, 6, 4, 5}, func(t int32) bool { return t > 5 }), false)
	xtestingEqual(t, All([]int32{9, 1, 8, 2, 7, 3, 6, 4, 5}, func(t int32) bool { return t > 0 }), true)
	xtestingEqual(t, All([]string{"aaa", "b", "ccccc", "dd"}, func(t string) bool { return len(t) > 2 }), false)
	xtestingEqual(t, All([]string{"aaa", "bb", "ccccc", "dd"}, func(t string) bool { return len(t) >= 2 }), true)
}

func TestZipUnzip(t *testing.T) {
	xtestingEqual(t, Zip[bool, bool](nil, nil), []xtuple.Tuple[bool, bool]{})
	xtestingEqual(t, Zip([]bool{}, []bool{}), []xtuple.Tuple[bool, bool]{})
	xtestingEqual(t, Zip([]bool{true}, []bool{}), []xtuple.Tuple[bool, bool]{})
	xtestingEqual(t, Zip([]bool{true}, []bool{false, false}), []xtuple.Tuple[bool, bool]{{true, false}})
	xtestingEqual(t, Zip([]int{1, 2, 3}, []string{"1", "2", "3"}), []xtuple.Tuple[int, string]{{1, "1"}, {2, "2"}, {3, "3"}})
	xtestingEqual(t, Zip([]string{")", "(", "*", "&", "^"}, []uint32{0, 9, 8, 7, 6}), []xtuple.Tuple[string, uint32]{{")", 0}, {"(", 9}, {"*", 8}, {"&", 7}, {"^", 6}})
	xtestingEqual(t, Zip([]float64{1 / 2, 3 / 4, 5 / 8, 6 / 3}, [][2]int{{1, 2}, {3, 4}, {5, 8}, {6, 3}}), []xtuple.Tuple[float64, [2]int]{{1 / 2, [2]int{1, 2}}, {3 / 4, [2]int{3, 4}}, {5 / 8, [2]int{5, 8}}, {6 / 3, [2]int{6, 3}}})

	xtestingEqual(t, xtuple.NewTuple(Unzip[bool, bool](nil)), xtuple.Tuple[[]bool, []bool]{Item1: []bool{}, Item2: []bool{}})
	xtestingEqual(t, xtuple.NewTuple(Unzip([]xtuple.Tuple[bool, bool]{})), xtuple.Tuple[[]bool, []bool]{Item1: []bool{}, Item2: []bool{}})
	xtestingEqual(t, xtuple.NewTuple(Unzip([]xtuple.Tuple[bool, bool]{{true, false}})), xtuple.Tuple[[]bool, []bool]{Item1: []bool{true}, Item2: []bool{false}})
	xtestingEqual(t, xtuple.NewTuple(Unzip([]xtuple.Tuple[bool, bool]{{true, false}, {false, true}})), xtuple.Tuple[[]bool, []bool]{Item1: []bool{true, false}, Item2: []bool{false, true}})
	xtestingEqual(t, xtuple.NewTuple(Unzip([]xtuple.Tuple[int, string]{{1, "1"}, {2, "2"}, {3, "3"}})), xtuple.Tuple[[]int, []string]{Item1: []int{1, 2, 3}, Item2: []string{"1", "2", "3"}})
	xtestingEqual(t, xtuple.NewTuple(Unzip([]xtuple.Tuple[string, uint32]{{")", 0}, {"(", 9}, {"*", 8}})), xtuple.Tuple[[]string, []uint32]{Item1: []string{")", "(", "*"}, Item2: []uint32{0, 9, 8}})

	xtestingEqual(t, Zip3[bool, bool, bool](nil, nil, nil), []xtuple.Triple[bool, bool, bool]{})
	xtestingEqual(t, Zip3([]bool{}, []bool{}, []bool{}), []xtuple.Triple[bool, bool, bool]{})
	xtestingEqual(t, Zip3([]bool{true, true, true}, []bool{true, true}, []bool{}), []xtuple.Triple[bool, bool, bool]{})
	xtestingEqual(t, Zip3([]bool{true}, []bool{true, true}, []bool{false, false, true}), []xtuple.Triple[bool, bool, bool]{{true, true, false}})
	xtestingEqual(t, Zip3([]int{1, 2, 3}, []string{"1", "2", "3"}, []uint32{2, 3, 4}), []xtuple.Triple[int, string, uint32]{{1, "1", 2}, {2, "2", 3}, {3, "3", 4}})
	xtestingEqual(t, Zip3([]string{")", "(", "*"}, []uint32{0, 9, 8}, []byte{'0', '9', '8'}), []xtuple.Triple[string, uint32, byte]{{")", 0, '0'}, {"(", 9, '9'}, {"*", 8, '8'}})

	xtestingEqual(t, xtuple.NewTriple(Unzip3[bool, bool, bool](nil)), xtuple.Triple[[]bool, []bool, []bool]{Item1: []bool{}, Item2: []bool{}, Item3: []bool{}})
	xtestingEqual(t, xtuple.NewTriple(Unzip3([]xtuple.Triple[bool, bool, bool]{})), xtuple.Triple[[]bool, []bool, []bool]{Item1: []bool{}, Item2: []bool{}, Item3: []bool{}})
	xtestingEqual(t, xtuple.NewTriple(Unzip3([]xtuple.Triple[bool, bool, int]{{true, false, 0}, {false, true, 1}})), xtuple.Triple[[]bool, []bool, []int]{Item1: []bool{true, false}, Item2: []bool{false, true}, Item3: []int{0, 1}})
	xtestingEqual(t, xtuple.NewTriple(Unzip3([]xtuple.Triple[int, string, uint32]{{1, "1", 2}, {2, "2", 3}, {3, "3", 4}})), xtuple.Triple[[]int, []string, []uint32]{Item1: []int{1, 2, 3}, Item2: []string{"1", "2", "3"}, Item3: []uint32{2, 3, 4}})
	xtestingEqual(t, xtuple.NewTriple(Unzip3([]xtuple.Triple[string, uint32, byte]{{")", 0, '0'}, {"(", 9, '9'}, {"*", 8, '8'}})), xtuple.Triple[[]string, []uint32, []byte]{Item1: []string{")", "(", "*"}, Item2: []uint32{0, 9, 8}, Item3: []byte{'0', '9', '8'}})
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
