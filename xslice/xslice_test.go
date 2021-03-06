package xslice

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestCheckParam(t *testing.T) {
	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{nil, []interface{}{}},
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}},
		{[]interface{}{1, nil, "2", false, 3.3}, []interface{}{1, nil, "2", false, 3.3}},
	} {
		xtesting.Equal(t, checkInterfaceSliceParam(tc.give).origin, tc.want)
	}

	for _, tc := range []struct {
		give      interface{}
		want      interface{}
		wantPanic bool
	}{
		{nil, nil, true},
		{0, nil, true},
		{[]interface{}(nil), []interface{}{}, false},
		{[]int(nil), []int{}, false},
		{[]int{}, []int{}, false},
		{[]int{1}, []int{1}, false},
		{[]int{1, 1, 1}, []int{1, 1, 1}, false},
		{[]int{1, 3, 0, 2}, []int{1, 3, 0, 2}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { checkSliceInterfaceParam(tc.give) })
		} else {
			xtesting.Equal(t, checkSliceInterfaceParam(tc.give).origin, tc.want)
		}
	}

	for _, tc := range []struct {
		give1     interface{}
		give2     interface{}
		want1     interface{}
		want2     interface{}
		wantPanic bool
	}{
		{nil, []interface{}{}, nil, nil, true},
		{[]interface{}{}, nil, nil, nil, true},
		{nil, 0, nil, nil, true},
		{0, nil, nil, nil, true},
		{0, []int{}, nil, nil, true},
		{[]int{}, 0, nil, nil, true},
		{[]int{}, []string{}, nil, nil, true},
		{[]int(nil), []int(nil), []int{}, []int{}, false},
		{[]int{}, []int{}, []int{}, []int{}, false},
		{[]int{0}, []int{1}, []int{0}, []int{1}, false},
		{[]int{0, 0, 0}, []int{1, 1, 1}, []int{0, 0, 0}, []int{1, 1, 1}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { checkTwoSliceInterfaceParam(tc.give1, tc.give2) })
		} else {
			s1, s2 := checkTwoSliceInterfaceParam(tc.give1, tc.give2)
			xtesting.Equal(t, s1.origin, tc.want1)
			xtesting.Equal(t, s2.origin, tc.want2)
		}
	}

	for _, tc := range []struct {
		give1     interface{}
		give2     interface{}
		want1     interface{}
		want2     interface{}
		wantPanic bool
	}{
		{nil, 0, nil, nil, true},
		{0, 0, nil, nil, true},
		{[]int{}, "0", nil, nil, true},
		{[]int{}, nil, []int{}, 0, false},
		{[]int{}, 0, []int{}, 0, false},
		{[]int(nil), 0, []int{}, 0, false},
		{[]int{0}, 0, []int{0}, 0, false},
		{[]int{1, 1, 1}, 1, []int{1, 1, 1}, 1, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { checkSliceInterfaceAndElemParam(tc.give1, tc.give2) })
		} else {
			s, v := checkSliceInterfaceAndElemParam(tc.give1, tc.give2)
			xtesting.Equal(t, s.origin, tc.want1)
			xtesting.Equal(t, v, tc.want2)
		}
	}
}

func TestInnerOfInterfaceSlice(t *testing.T) {
	slice := checkInterfaceSliceParam([]interface{}{1, 2, 3, 4, 5, 6})
	// actual
	xtesting.Equal(t, slice.actual(), []interface{}{1, 2, 3, 4, 5, 6})
	// length
	xtesting.Equal(t, slice.length(), 6)
	// get
	xtesting.Equal(t, slice.get(0), 1)
	xtesting.Equal(t, slice.get(5), 6)
	xtesting.Panic(t, func() { slice.get(-1) })
	xtesting.Panic(t, func() { slice.get(6) })
	// set
	slice.set(0, 11)
	xtesting.Equal(t, slice.get(0), 11)
	slice.set(5, 66)
	xtesting.Equal(t, slice.get(5), 66)
	xtesting.Panic(t, func() { slice.set(-1, 0) })
	xtesting.Panic(t, func() { slice.set(6, 0) })
	// slice
	xtesting.Equal(t, slice.slice(0, 0), []interface{}{})
	xtesting.Equal(t, slice.slice(0, 1), []interface{}{11})
	xtesting.Equal(t, slice.slice(5, 6), []interface{}{66})
	xtesting.Equal(t, slice.slice(1, 5), []interface{}{2, 3, 4, 5})
	xtesting.Panic(t, func() { slice.slice(-1, 0) })
	xtesting.Panic(t, func() { slice.slice(6, 7) })
	xtesting.Panic(t, func() { slice.slice(4, 3) })
	// remove
	slice.remove(5)
	xtesting.Equal(t, slice.origin, []interface{}{11, 2, 3, 4, 5})
	slice.remove(0)
	xtesting.Equal(t, slice.origin, []interface{}{2, 3, 4, 5})
	slice.remove(2)
	xtesting.Equal(t, slice.origin, []interface{}{2, 3, 5})
	xtesting.Panic(t, func() { slice.remove(-1) })
	xtesting.Panic(t, func() { slice.remove(3) })
	// replace
	slice.replace([]interface{}{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, slice.origin, []interface{}{1, 2, 3, 4, 5, 6})
	xtesting.Panic(t, func() { slice.replace(nil) })
	xtesting.Panic(t, func() { slice.replace(0) })
	xtesting.Panic(t, func() { slice.replace([]string{""}) })
	// append
	slice.append(7)
	xtesting.Equal(t, slice.origin, []interface{}{1, 2, 3, 4, 5, 6, 7})
	slice.append(nil)
	xtesting.Equal(t, slice.origin, []interface{}{1, 2, 3, 4, 5, 6, 7, nil})
	slice.append("0")
	xtesting.Equal(t, slice.origin, []interface{}{1, 2, 3, 4, 5, 6, 7, nil, "0"})
}

func TestInnerInterfaceWrappedSlice(t *testing.T) {
	slice := checkSliceInterfaceParam([]int{1, 2, 3, 4, 5, 6})
	// actual
	xtesting.Equal(t, slice.actual(), []int{1, 2, 3, 4, 5, 6})
	// length
	xtesting.Equal(t, slice.length(), 6)
	// get
	xtesting.Equal(t, slice.get(0), 1)
	xtesting.Equal(t, slice.get(5), 6)
	xtesting.Panic(t, func() { slice.get(-1) })
	xtesting.Panic(t, func() { slice.get(6) })
	// set
	slice.set(0, nil)
	xtesting.Equal(t, slice.get(0), 0)
	slice.set(0, 11)
	xtesting.Equal(t, slice.get(0), 11)
	slice.set(5, 66)
	xtesting.Equal(t, slice.get(5), 66)
	xtesting.Panic(t, func() { slice.set(-1, 0) })
	xtesting.Panic(t, func() { slice.set(6, 0) })
	xtesting.Panic(t, func() { slice.set(0, "") })
	// slice
	xtesting.Equal(t, slice.slice(0, 0), []interface{}{})
	xtesting.Equal(t, slice.slice(0, 1), []interface{}{11})
	xtesting.Equal(t, slice.slice(5, 6), []interface{}{66})
	xtesting.Equal(t, slice.slice(1, 5), []interface{}{2, 3, 4, 5})
	xtesting.Panic(t, func() { slice.slice(-1, 0) })
	xtesting.Panic(t, func() { slice.slice(6, 7) })
	xtesting.Panic(t, func() { slice.slice(4, 3) })
	// remove
	slice.remove(5)
	xtesting.Equal(t, slice.origin, []int{11, 2, 3, 4, 5})
	slice.remove(0)
	xtesting.Equal(t, slice.origin, []int{2, 3, 4, 5})
	slice.remove(2)
	xtesting.Equal(t, slice.origin, []int{2, 3, 5})
	xtesting.Panic(t, func() { slice.remove(-1) })
	xtesting.Panic(t, func() { slice.remove(3) })
	// replace
	slice.replace([]int{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, slice.origin, []int{1, 2, 3, 4, 5, 6})
	xtesting.Panic(t, func() { slice.replace(nil) })
	xtesting.Panic(t, func() { slice.replace(0) })
	xtesting.Panic(t, func() { slice.replace([]string{""}) })
	// append
	slice.append(7)
	xtesting.Equal(t, slice.origin, []int{1, 2, 3, 4, 5, 6, 7})
	slice.append(nil)
	xtesting.Equal(t, slice.origin, []int{1, 2, 3, 4, 5, 6, 7, 0})
	xtesting.Panic(t, func() { slice.append("0") })
}

func TestCloneAndMakeSlice(t *testing.T) {
	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{nil, []interface{}{}},
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}},
		{[]interface{}{1, nil, "2", false, 3.3}, []interface{}{1, nil, "2", false, 3.3}},
	} {
		xtesting.Equal(t, cloneInterfaceSlice(tc.give), tc.want)
	}

	for _, tc := range []struct {
		give      interface{}
		want      interface{}
		wantPanic bool
	}{
		{nil, nil, true},
		{0, nil, true},
		{[]interface{}(nil), []interface{}{}, false},
		{[]int(nil), []int{}, false},
		{[]int{}, []int{}, false},
		{[]int{1}, []int{1}, false},
		{[]int{1, 1, 1}, []int{1, 1, 1}, false},
		{[]int{1, 3, 0, 2}, []int{1, 3, 0, 2}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { cloneSliceInterface(tc.give) })
		} else {
			xtesting.Equal(t, cloneSliceInterface(tc.give), tc.want)
		}
	}

	for _, tc := range []struct {
		giveType  innerSlice
		giveLen   int
		giveCap   int
		want      interface{}
		wantPanic bool
	}{
		{nil, 0, 0, nil, true},
		{&innerOfInterfaceSlice{}, -1, 0, nil, true},
		{&innerOfInterfaceSlice{}, 0, 0, []interface{}{}, false},
		{&innerOfInterfaceSlice{}, 1, 1, []interface{}{nil}, false},
		{&innerOfInterfaceSlice{}, 3, 0, []interface{}{nil, nil, nil}, false},
		{&innerInterfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, -1, 0, []int{}, true},
		{&innerInterfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 0, 0, []int{}, false},
		{&innerInterfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 1, 1, []int{0}, false},
		{&innerInterfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 3, 1, []int{0, 0, 0}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { makeInnerSlice(tc.giveType, tc.giveLen, tc.giveCap) })
		} else {
			xtesting.Equal(t, makeInnerSlice(tc.giveType, tc.giveLen, tc.giveCap).actual(), tc.want)
		}
	}
}

func TestShuffle(t *testing.T) {
	// Shuffle & ShuffleSelf
	for _, tc := range []struct {
		give []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]interface{}{"1", 2, 3.0, uint(4)}},
	} {
		me := make([]interface{}, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			result := Shuffle(tc.give)
			xtesting.Equal(t, tc.give, me)
			xtesting.ElementMatch(t, result, me)
			fmt.Println(result)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			ShuffleSelf(tc.give)
			xtesting.ElementMatch(t, tc.give, me)
			fmt.Println(tc.give)
		}
	}

	// ShuffleG & ShuffleSelfG
	for _, tc := range []struct {
		give []int
	}{
		{[]int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 2, 3, 4}},
	} {
		me := make([]int, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			result := ShuffleG(tc.give)
			xtesting.Equal(t, tc.give, me)
			xtesting.ElementMatch(t, result, me)
			fmt.Println(result)
		}
		for i := 0; i < 2; i++ {
			time.Sleep(2 * time.Nanosecond)
			ShuffleSelfG(tc.give)
			xtesting.ElementMatch(t, tc.give, me)
			fmt.Println(tc.give)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{0}, []interface{}{0}},
		{[]interface{}{1, 2, 3}, []interface{}{3, 2, 1}},
	} {
		me := make([]interface{}, 0, len(tc.give))
		for _, item := range tc.give {
			me = append(me, item)
		}
		result := Reverse(tc.give)
		xtesting.Equal(t, tc.give, me)
		xtesting.Equal(t, result, tc.want)
		ReverseSelf(tc.give)
		xtesting.Equal(t, tc.give, tc.want)
	}

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
		result := ReverseG(tc.give)
		xtesting.Equal(t, tc.give, me)
		xtesting.Equal(t, result, tc.want)
		ReverseSelfG(tc.give)
		xtesting.Equal(t, tc.give, tc.want)
	}
}

func TestSort(t *testing.T) {
	le := func(i, j interface{}) bool { return i.(int) < j.(int) }

	for _, tc := range []struct {
		give      []interface{}
		giveLess  Lesser
		want      []interface{}
		wantPanic bool
	}{
		{[]interface{}{}, nil, nil, true},
		{[]interface{}{}, le, []interface{}{}, false},
		{[]interface{}{0}, le, []interface{}{0}, false},
		{[]interface{}{1, 1, 1}, le, []interface{}{1, 1, 1}, false},
		{[]interface{}{4, 3, 2, 1}, le, []interface{}{1, 2, 3, 4}, false},
		{[]interface{}{8, 1, 6, 8, 1, 2}, le, []interface{}{1, 1, 2, 6, 8, 8}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { Sort(tc.give, tc.giveLess) })
		} else {
			me := make([]interface{}, 0, len(tc.give))
			for _, item := range tc.give {
				me = append(me, item)
			}
			result := Sort(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, me)
			xtesting.Equal(t, result, tc.want)
			SortSelf(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, tc.want)
		}
	}

	for _, tc := range []struct {
		give      []int
		giveLess  Lesser
		want      []int
		wantPanic bool
	}{
		{[]int{}, nil, nil, true},
		{[]int{}, le, []int{}, false},
		{[]int{0}, le, []int{0}, false},
		{[]int{1, 1, 1}, le, []int{1, 1, 1}, false},
		{[]int{4, 3, 2, 1}, le, []int{1, 2, 3, 4}, false},
		{[]int{8, 1, 6, 8, 1, 2}, le, []int{1, 1, 2, 6, 8, 8}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { SortG(tc.give, tc.giveLess) })
		} else {
			me := make([]int, 0, len(tc.give))
			for _, item := range tc.give {
				me = append(me, item)
			}
			result := SortG(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, me)
			xtesting.Equal(t, result, tc.want)
			SortSelfG(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, tc.want)
		}
	}
}

func TestStableSort(t *testing.T) {
	type testStruct struct {
		value int
		other int
	}
	new1 := func(v int) testStruct { return testStruct{value: v} }
	new2 := func(v, o int) testStruct { return testStruct{value: v, other: o} }
	le := func(i, j interface{}) bool { return i.(testStruct).value < j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveLess  Lesser
		want      []interface{}
		wantPanic bool
	}{
		{[]interface{}{}, nil, nil, true},
		{[]interface{}{}, le, []interface{}{}, false},
		{[]interface{}{new1(0)}, le, []interface{}{new1(0)}, false},
		{[]interface{}{new2(1, 3), new2(1, 2), new2(1, 1)}, le,
			[]interface{}{new2(1, 3), new2(1, 2), new2(1, 1)}, false},
		{[]interface{}{new1(4), new1(3), new1(2), new1(1)}, le,
			[]interface{}{new1(1), new1(2), new1(3), new1(4)}, false},
		{[]interface{}{new2(8, 2), new2(1, 2), new1(6), new2(8, 1), new2(1, 1), new1(2)}, le,
			[]interface{}{new2(1, 2), new2(1, 1), new1(2), new1(6), new2(8, 2), new2(8, 1)}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { Sort(tc.give, tc.giveLess) })
		} else {
			me := make([]interface{}, 0, len(tc.give))
			for _, item := range tc.give {
				me = append(me, item)
			}
			result := StableSort(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, me)
			xtesting.Equal(t, result, tc.want)
			StableSortSelf(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, tc.want)
		}
	}

	for _, tc := range []struct {
		give      []testStruct
		giveLess  Lesser
		want      []testStruct
		wantPanic bool
	}{
		{[]testStruct{}, nil, nil, true},
		{[]testStruct{}, le, []testStruct{}, false},
		{[]testStruct{new1(0)}, le, []testStruct{new1(0)}, false},
		{[]testStruct{new2(1, 3), new2(1, 2), new2(1, 1)}, le,
			[]testStruct{new2(1, 3), new2(1, 2), new2(1, 1)}, false},
		{[]testStruct{new1(4), new1(3), new1(2), new1(1)}, le,
			[]testStruct{new1(1), new1(2), new1(3), new1(4)}, false},
		{[]testStruct{new2(8, 2), new2(1, 2), new1(6), new2(8, 1), new2(1, 1), new1(2)}, le,
			[]testStruct{new2(1, 2), new2(1, 1), new1(2), new1(6), new2(8, 2), new2(8, 1)}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { SortG(tc.give, tc.giveLess) })
		} else {
			me := make([]testStruct, 0, len(tc.give))
			for _, item := range tc.give {
				me = append(me, item)
			}
			result := StableSortG(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, me)
			xtesting.Equal(t, result, tc.want)
			StableSortSelfG(tc.give, tc.giveLess)
			xtesting.Equal(t, tc.give, tc.want)
		}
	}

}

type testStruct struct {
	value int
	now   time.Time
}

func (t testStruct) String() string {
	return strconv.Itoa(t.value)
}

func newTestStruct(value int) testStruct {
	time.Sleep(2 * time.Nanosecond)
	return testStruct{value: value, now: time.Now()}
}

func newTestStructSlice1(s []interface{}) []interface{} {
	newSlice := make([]interface{}, len(s))
	for idx, item := range s {
		newSlice[idx] = newTestStruct(item.(int))
	}
	return newSlice
}

func newTestStructSlice2(s []int) []testStruct {
	newSlice := make([]testStruct, len(s))
	for idx, item := range s {
		newSlice[idx] = newTestStruct(item)
	}
	return newSlice
}

func toInterfaceSlice(s []interface{}) []interface{} {
	out := make([]interface{}, len(s))
	for idx, item := range s {
		out[idx] = item.(testStruct).value
	}
	return out
}

func toIntSlice(s interface{}) []int {
	out := make([]int, len(s.([]testStruct)))
	for idx, item := range s.([]testStruct) {
		out[idx] = item.value
	}
	return out
}

func TestIndexOf(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 2, 3}
	s2 := []int{1, 5, 2, 1, 2, 3}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		want      int
	}{
		{[]interface{}{}, 0, -1},
		{s1, -1, -1},
		{s1, 0, -1},
		{s1, 1, 0},
		{s1, 2, 2},
		{s1, 3, 5},
		{s1, 4, -1},
		{s1, 5, 1},
		{s1, 6, -1},
	} {
		xtesting.Equal(t, IndexOf(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, IndexOfWith(give, giveValue, eq), tc.want)
	}

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
		xtesting.Equal(t, IndexOfG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, IndexOfWithG(give, giveValue, eq), tc.want)
	}
}

func TestContains(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 2, 3}
	s2 := []int{1, 5, 2, 1, 2, 3}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		want      bool
	}{
		{[]interface{}{}, 0, false},
		{s1, -1, false},
		{s1, 0, false},
		{s1, 1, true},
		{s1, 2, true},
		{s1, 3, true},
		{s1, 4, false},
		{s1, 5, true},
		{s1, 6, false},
	} {
		xtesting.Equal(t, Contains(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, ContainsWith(give, giveValue, eq), tc.want)
	}

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
		xtesting.Equal(t, ContainsG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, ContainsWithG(give, giveValue, eq), tc.want)
	}
}

func TestCount(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		want      int
	}{
		{[]interface{}{}, 0, 0},
		{s1, -1, 0},
		{s1, 0, 0},
		{s1, 1, 2},
		{s1, 2, 3},
		{s1, 3, 1},
		{s1, 4, 0},
		{s1, 5, 2},
		{s1, 6, 1},
		{s1, 7, 0},
	} {
		xtesting.Equal(t, Count(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, CountWith(give, giveValue, eq), tc.want)
	}

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
		xtesting.Equal(t, CountG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, CountWithG(give, giveValue, eq), tc.want)
	}
}

func TestDelete(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		giveN     int
		want      []interface{}
	}{
		{[]interface{}{}, 0, 1, []interface{}{}},
		{s1, -1, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 0, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 1, 1, []interface{}{5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 1, 2, []interface{}{5, 2, 5, 2, 6, 3, 2}},
		{s1, 2, 1, []interface{}{1, 5, 1, 5, 2, 6, 3, 2}},
		{s1, 2, 2, []interface{}{1, 5, 1, 5, 6, 3, 2}},
		{s1, 2, -1, []interface{}{1, 5, 1, 5, 6, 3}},
		{s1, 3, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 2}},
		{s1, 4, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 5, 1, []interface{}{1, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 6, 1, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{s1, 7, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, Delete(tc.give, tc.giveValue, tc.giveN), tc.want)
		give, giveValue := newTestStructSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, toInterfaceSlice(DeleteWith(give, giveValue, tc.giveN, eq)), tc.want)
	}

	for _, tc := range []struct {
		give      []int
		giveValue int
		giveN     int
		want      []int
	}{
		{[]int{}, 0, 1, []int{}},
		{s2, -1, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 0, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 1, 1, []int{5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 1, 2, []int{5, 2, 5, 2, 6, 3, 2}},
		{s2, 2, 1, []int{1, 5, 1, 5, 2, 6, 3, 2}},
		{s2, 2, 2, []int{1, 5, 1, 5, 6, 3, 2}},
		{s2, 2, -1, []int{1, 5, 1, 5, 6, 3}},
		{s2, 3, 1, []int{1, 5, 2, 1, 5, 2, 6, 2}},
		{s2, 4, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 5, 1, []int{1, 2, 1, 5, 2, 6, 3, 2}},
		{s2, 6, 1, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{s2, 7, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, DeleteG(tc.give, tc.giveValue, tc.giveN), tc.want)
		give, giveValue := newTestStructSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, toIntSlice(DeleteWithG(give, giveValue, tc.giveN, eq)), tc.want)
	}
}

func TestDeleteAll(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		want      []interface{}
	}{
		{[]interface{}{}, 0, []interface{}{}},
		{s1, -1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 0, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 1, []interface{}{5, 2, 5, 2, 6, 3, 2}},
		{s1, 2, []interface{}{1, 5, 1, 5, 6, 3}},
		{s1, 3, []interface{}{1, 5, 2, 1, 5, 2, 6, 2}},
		{s1, 4, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, 5, []interface{}{1, 2, 1, 2, 6, 3, 2}},
		{s1, 6, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{s1, 7, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, DeleteAll(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, toInterfaceSlice(DeleteAllWith(give, giveValue, eq)), tc.want)
	}

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
		xtesting.Equal(t, DeleteAllG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestStructSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, toIntSlice(DeleteAllWithG(give, giveValue, eq)), tc.want)
	}
}

func TestDiff(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  []interface{}
	}{
		{[]interface{}{}, []interface{}{}, []interface{}{}},
		{s1, []interface{}{}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, []interface{}{1}, []interface{}{5, 2, 5, 2, 6, 3, 2}},
		{s1, []interface{}{1, 2}, []interface{}{5, 5, 6, 3}},
		{s1, []interface{}{1, 2, 3}, []interface{}{5, 5, 6}},
		{s1, []interface{}{1, 2, 3, 4}, []interface{}{5, 5, 6}},
		{s1, []interface{}{1, 2, 3, 4, 5}, []interface{}{6}},
		{s1, []interface{}{1, 2, 3, 4, 5, 6}, []interface{}{}},
		{s1, []interface{}{6}, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{s1, []interface{}{6, 5}, []interface{}{1, 2, 1, 2, 3, 2}},
		{s1, []interface{}{6, 5, 4}, []interface{}{1, 2, 1, 2, 3, 2}},
		{s1, []interface{}{6, 5, 4, 3}, []interface{}{1, 2, 1, 2, 2}},
		{s1, []interface{}{6, 5, 4, 3, 2}, []interface{}{1, 1}},
		{s1, []interface{}{6, 5, 4, 3, 2, 1}, []interface{}{}},
	} {
		xtesting.Equal(t, Diff(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice1(tc.give1), newTestStructSlice1(tc.give2)
		xtesting.Equal(t, toInterfaceSlice(DiffWith(give1, give2, eq)), tc.want)
	}

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
		xtesting.Equal(t, DiffG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice2(tc.give1), newTestStructSlice2(tc.give2)
		xtesting.Equal(t, toIntSlice(DiffWithG(give1, give2, eq)), tc.want)
	}
}

func TestUnion(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  []interface{}
	}{
		{[]interface{}{}, []interface{}{}, []interface{}{}},
		{s1, []interface{}{}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{s1, []interface{}{11}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11}},
		{s1, []interface{}{11, 2}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11}},
		{s1, []interface{}{11, 2, 13}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13}},
		{s1, []interface{}{11, 2, 13, 14}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14}},
		{s1, []interface{}{11, 2, 13, 14, 5}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14}},
		{s1, []interface{}{11, 2, 13, 14, 5, 16}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 11, 13, 14, 16}},
	} {
		xtesting.Equal(t, Union(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice1(tc.give1), newTestStructSlice1(tc.give2)
		xtesting.Equal(t, toInterfaceSlice(UnionWith(give1, give2, eq)), tc.want)
	}

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
		xtesting.Equal(t, UnionG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice2(tc.give1), newTestStructSlice2(tc.give2)
		xtesting.Equal(t, toIntSlice(UnionWithG(give1, give2, eq)), tc.want)
	}
}

func TestIntersection(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  []interface{}
	}{
		{[]interface{}{}, []interface{}{}, []interface{}{}},
		{s1, []interface{}{}, []interface{}{}},
		{s1, []interface{}{1}, []interface{}{1, 1}},
		{s1, []interface{}{1, 2}, []interface{}{1, 2, 1, 2, 2}},
		{s1, []interface{}{1, 2, 3}, []interface{}{1, 2, 1, 2, 3, 2}},
		{s1, []interface{}{1, 2, 3, 4}, []interface{}{1, 2, 1, 2, 3, 2}},
		{s1, []interface{}{1, 2, 3, 4, 5}, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{s1, []interface{}{1, 2, 3, 4, 5, 6}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, Intersection(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice1(tc.give1), newTestStructSlice1(tc.give2)
		xtesting.Equal(t, toInterfaceSlice(IntersectionWith(give1, give2, eq)), tc.want)
	}

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
		xtesting.Equal(t, IntersectionG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice2(tc.give1), newTestStructSlice2(tc.give2)
		xtesting.Equal(t, toIntSlice(IntersectionWithG(give1, give2, eq)), tc.want)
	}
}

func TestToSet(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 1, 1}, []interface{}{1}},
		{[]interface{}{2, 1, 1}, []interface{}{2, 1}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, []interface{}{1, 5, 2, 6, 3}},
	} {
		xtesting.Equal(t, ToSet(tc.give), tc.want)
		give := newTestStructSlice1(tc.give)
		xtesting.Equal(t, toInterfaceSlice(ToSetWith(give, eq)), tc.want)
	}

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
		xtesting.Equal(t, ToSetG(tc.give), tc.want)
		give := newTestStructSlice2(tc.give)
		xtesting.Equal(t, toIntSlice(ToSetWithG(give, eq)), tc.want)
	}
}

func TestElementMatch(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  bool
	}{
		{[]interface{}{}, []interface{}{}, true},
		{[]interface{}{0}, []interface{}{}, false},
		{[]interface{}{0}, []interface{}{0}, true},
		{[]interface{}{1, 1, 1}, []interface{}{1}, false},
		{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}, true},
		{[]interface{}{1, 2, 1}, []interface{}{1, 2, 1}, true},
		{[]interface{}{1, 2, 1}, []interface{}{1, 2, 2}, false},
	} {
		xtesting.Equal(t, ElementMatch(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice1(tc.give1), newTestStructSlice1(tc.give2)
		xtesting.Equal(t, ElementMatchWith(give1, give2, eq), tc.want)
	}

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
		xtesting.Equal(t, ElementMatchG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestStructSlice2(tc.give1), newTestStructSlice2(tc.give2)
		xtesting.Equal(t, ElementMatchWithG(give1, give2, eq), tc.want)
	}
}

func TestRange(t *testing.T) {
	for _, tc := range []struct {
		giveMin   int
		giveMax   int
		giveStep  int
		want      []int
		wantPanic bool
	}{
		{0, 0, 1, []int{0}, false},
		{0, 1, 1, []int{0, 1}, false},
		{0, 1, 2, []int{0}, false},
		{1, 1, 1, []int{1}, false},
		{1, 10, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{1, 10, 2, []int{1, 3, 5, 7, 9}, false},
		{0, 9, 2, []int{0, 2, 4, 6, 8}, false},

		{2, 1, 1, nil, true},
		{1, 2, 0, nil, true},
		{1, 2, -1, nil, true},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { Range(tc.giveMin, tc.giveMax, tc.giveStep) })
		} else {
			xtesting.Equal(t, Range(tc.giveMin, tc.giveMax, tc.giveStep), tc.want)
		}
	}
}

func TestReverseRange(t *testing.T) {
	for _, tc := range []struct {
		giveMin   int
		giveMax   int
		giveStep  int
		want      []int
		wantPanic bool
	}{
		{0, 0, 1, []int{0}, false},
		{0, 1, 1, []int{1, 0}, false},
		{0, 1, 2, []int{1}, false},
		{1, 1, 1, []int{1}, false},
		{1, 10, 1, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
		{1, 10, 2, []int{10, 8, 6, 4, 2}, false},
		{0, 9, 2, []int{9, 7, 5, 3, 1}, false},

		{2, 1, 1, nil, true},
		{1, 2, 0, nil, true},
		{1, 2, -1, nil, true},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { ReverseRange(tc.giveMin, tc.giveMax, tc.giveStep) })
		} else {
			xtesting.Equal(t, ReverseRange(tc.giveMin, tc.giveMax, tc.giveStep), tc.want)
		}
	}
}
