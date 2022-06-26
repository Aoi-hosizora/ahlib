package xslice

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

// =============================
// testing on internal functions
// =============================

func TestCheckParam(t *testing.T) {
	// checkInterfaceSliceParam
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
		xtesting.Equal(t, checkInterfaceSliceParam(tc.give).(*interfaceItemSlice).origin, tc.want)
	}

	// checkSliceInterfaceParam
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
			xtesting.Equal(t, checkSliceInterfaceParam(tc.give).(*interfaceWrappedSlice).val.Interface(), tc.want)
		}
	}

	// checkTwoSliceInterfaceParam
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
			xtesting.Equal(t, s1.(*interfaceWrappedSlice).val.Interface(), tc.want1)
			xtesting.Equal(t, s2.(*interfaceWrappedSlice).val.Interface(), tc.want2)
		}
	}

	// checkSliceInterfaceAndElemParam
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
			xtesting.Equal(t, s.(*interfaceWrappedSlice).val.Interface(), tc.want1)
			xtesting.Equal(t, v, tc.want2)
		}
	}
}

func TestInterfaceItemSlice(t *testing.T) {
	slice := checkInterfaceSliceParam(append(make([]interface{}, 0, 10), 1, 2, 3, 4, 5, 6))
	// actual
	xtesting.Equal(t, slice.actual(), []interface{}{1, 2, 3, 4, 5, 6})
	// length
	xtesting.Equal(t, slice.length(), 6)
	// capacity
	xtesting.Equal(t, slice.capacity(), 10)
	// get
	xtesting.Equal(t, slice.get(0), 1)
	xtesting.Equal(t, slice.get(5), 6)
	xtesting.Panic(t, func() { slice.get(-1) })
	xtesting.Panic(t, func() { slice.get(6) })
	// slice
	xtesting.Equal(t, slice.slice(0, 0).actual(), []interface{}{})
	xtesting.Equal(t, slice.slice(0, slice.length()).actual(), []interface{}{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, slice.slice(3, slice.capacity()).actual(), []interface{}{4, 5, 6, nil, nil, nil, nil})
	xtesting.Panic(t, func() { slice.slice(2, 1) })
	xtesting.Panic(t, func() { slice.slice(3, slice.capacity()+1) })
	// set
	slice.set(0, 11)
	xtesting.Equal(t, slice.get(0), 11)
	slice.set(5, 66)
	xtesting.Equal(t, slice.get(5), 66)
	xtesting.Panic(t, func() { slice.set(-1, 0) })
	xtesting.Panic(t, func() { slice.set(6, 0) })
	// insert
	slice.insert(-1, &interfaceItemSlice{[]interface{}{-3, -2, -1}})
	xtesting.Equal(t, slice.actual(), []interface{}{-3, -2, -1, 11, 2, 3, 4, 5, 66})
	xtesting.Equal(t, slice.capacity(), 10)
	slice.insert(3, &interfaceItemSlice{[]interface{}{}})
	xtesting.Equal(t, slice.actual(), []interface{}{-3, -2, -1, 11, 2, 3, 4, 5, 66})
	slice.insert(3, &interfaceItemSlice{[]interface{}{0, 0}})
	xtesting.Equal(t, slice.actual(), []interface{}{-3, -2, -1, 0, 0, 11, 2, 3, 4, 5, 66})
	xtesting.NotEqual(t, slice.capacity(), 10)
	slice.insert(1000, &interfaceItemSlice{[]interface{}{7, 8}})
	xtesting.Equal(t, slice.actual(), []interface{}{-3, -2, -1, 0, 0, 11, 2, 3, 4, 5, 66, 7, 8})
	slice.insert(6, &interfaceItemSlice{[]interface{}{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}})
	xtesting.Equal(t, slice.actual(), []interface{}{-3, -2, -1, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 3, 4, 5, 66, 7, 8})
	xtesting.Panic(t, func() { slice.insert(0, nil) })
	xtesting.Panic(t, func() { slice.insert(0, &interfaceWrappedSlice{}) })

	slice.(*interfaceItemSlice).origin = []interface{}{11, 2, 3, 4, 5, 66}
	// remove
	slice.remove(5)
	xtesting.Equal(t, slice.actual(), []interface{}{11, 2, 3, 4, 5})
	slice.remove(0)
	xtesting.Equal(t, slice.actual(), []interface{}{2, 3, 4, 5})
	slice.remove(2)
	xtesting.Equal(t, slice.actual(), []interface{}{2, 3, 5})
	xtesting.Panic(t, func() { slice.remove(-1) })
	xtesting.Panic(t, func() { slice.remove(3) })
	// append
	slice.append(7)
	xtesting.Equal(t, slice.actual(), []interface{}{2, 3, 5, 7})
	slice.append(nil)
	xtesting.Equal(t, slice.actual(), []interface{}{2, 3, 5, 7, nil})
	slice.append("0")
	xtesting.Equal(t, slice.actual(), []interface{}{2, 3, 5, 7, nil, "0"})
}

func TestInterfaceWrappedSlice(t *testing.T) {
	slice := checkSliceInterfaceParam(append(make([]int, 0, 10), 1, 2, 3, 4, 5, 6))
	// actual
	xtesting.Equal(t, slice.actual(), []int{1, 2, 3, 4, 5, 6})
	// length
	xtesting.Equal(t, slice.length(), 6)
	// capacity
	xtesting.Equal(t, slice.capacity(), 10)
	// get
	xtesting.Equal(t, slice.get(0), 1)
	xtesting.Equal(t, slice.get(5), 6)
	xtesting.Panic(t, func() { slice.get(-1) })
	xtesting.Panic(t, func() { slice.get(6) })
	// slice
	xtesting.Equal(t, slice.slice(0, 0).actual(), []int{})
	xtesting.Equal(t, slice.slice(0, slice.length()).actual(), []int{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, slice.slice(3, slice.capacity()).actual(), []int{4, 5, 6, 0, 0, 0, 0})
	xtesting.Panic(t, func() { slice.slice(2, 1) })
	xtesting.Panic(t, func() { slice.slice(3, slice.capacity()+1) })
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
	// insert
	typ := reflect.TypeOf([]int{})
	slice.insert(-1, &interfaceWrappedSlice{reflect.ValueOf([]int{-3, -2, -1}), typ})
	xtesting.Equal(t, slice.actual(), []int{-3, -2, -1, 11, 2, 3, 4, 5, 66})
	xtesting.Equal(t, slice.capacity(), 10)
	slice.insert(3, &interfaceWrappedSlice{reflect.ValueOf([]int{}), typ})
	xtesting.Equal(t, slice.actual(), []int{-3, -2, -1, 11, 2, 3, 4, 5, 66})
	slice.insert(3, &interfaceWrappedSlice{reflect.ValueOf([]int{0, 0}), typ})
	xtesting.Equal(t, slice.actual(), []int{-3, -2, -1, 0, 0, 11, 2, 3, 4, 5, 66})
	xtesting.NotEqual(t, slice.capacity(), 10)
	slice.insert(1000, &interfaceWrappedSlice{reflect.ValueOf([]int{7, 8}), typ})
	xtesting.Equal(t, slice.actual(), []int{-3, -2, -1, 0, 0, 11, 2, 3, 4, 5, 66, 7, 8})
	slice.insert(6, &interfaceWrappedSlice{reflect.ValueOf([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), typ})
	xtesting.Equal(t, slice.actual(), []int{-3, -2, -1, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 3, 4, 5, 66, 7, 8})
	xtesting.Panic(t, func() { slice.insert(0, nil) })
	xtesting.Panic(t, func() { slice.insert(0, &interfaceItemSlice{}) })

	slice.(*interfaceWrappedSlice).val = reflect.ValueOf([]int{11, 2, 3, 4, 5, 66})
	// remove
	slice.remove(5)
	xtesting.Equal(t, slice.actual(), []int{11, 2, 3, 4, 5})
	slice.remove(0)
	xtesting.Equal(t, slice.actual(), []int{2, 3, 4, 5})
	slice.remove(2)
	xtesting.Equal(t, slice.actual(), []int{2, 3, 5})
	xtesting.Panic(t, func() { slice.remove(-1) })
	xtesting.Panic(t, func() { slice.remove(3) })
	// append
	slice.append(7)
	xtesting.Equal(t, slice.actual(), []int{2, 3, 5, 7})
	slice.append(nil)
	xtesting.Equal(t, slice.actual(), []int{2, 3, 5, 7, 0})
	xtesting.Panic(t, func() { slice.append("0") })
}

func TestCloneAndMakeSlice(t *testing.T) {
	// cloneInterfaceSlice
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

	// cloneSliceInterface
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
		{[]interface{}{1, 3, nil, 2}, []interface{}{1, 3, nil, 2}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { cloneSliceInterface(tc.give) })
		} else {
			xtesting.Equal(t, cloneSliceInterface(tc.give), tc.want)
		}
	}

	// cloneSliceInterfaceFromInterfaceSlice
	for _, tc := range []struct {
		give1     []interface{}
		give2     interface{}
		want      interface{}
		wantPanic bool
	}{
		{nil, nil, nil, true},
		{nil, 0, nil, true},
		{nil, []int{}, []int{}, false},
		{[]interface{}{nil}, []int{}, []int{0}, false},
		{[]interface{}{1, 2}, []int{}, []int{1, 2}, false},
		{[]interface{}{uint(1), "2"}, []uint{}, nil, true},
		{[]interface{}{uint(1), uint(2)}, []uint32{}, nil, true},
		{[]interface{}{uint(1), uint(2)}, []uint{}, []uint{1, 2}, false},
		{[]interface{}{nil, nil, "x"}, []string{}, []string{"", "", "x"}, false},
		{[]interface{}{nil, 1, "x", false}, []interface{}{}, []interface{}{nil, 1, "x", false}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { cloneSliceInterfaceFromInterfaceSlice(tc.give1, tc.give2) })
		} else {
			xtesting.Equal(t, cloneSliceInterfaceFromInterfaceSlice(tc.give1, tc.give2), tc.want)
		}
	}

	// makeSameTypeInnerSlice
	for _, tc := range []struct {
		giveType  innerSlice
		giveLen   int
		giveCap   int
		want      interface{}
		wantPanic bool
	}{
		{nil, 0, 0, nil, true},
		{&interfaceItemSlice{}, -1, 0, []interface{}{}, false}, // no panic
		{&interfaceItemSlice{}, 0, 0, []interface{}{}, false},
		{&interfaceItemSlice{}, 1, 1, []interface{}{nil}, false},
		{&interfaceItemSlice{}, 3, 0, []interface{}{nil, nil, nil}, false},
		{&interfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, -1, 0, []int{}, false}, // no panic
		{&interfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 0, 0, []int{}, false},
		{&interfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 1, 1, []int{0}, false},
		{&interfaceWrappedSlice{typ: reflect.TypeOf([]int{})}, 3, 1, []int{0, 0, 0}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { makeSameTypeInnerSlice(tc.giveType, tc.giveLen, tc.giveCap) })
		} else {
			xtesting.Equal(t, makeSameTypeInnerSlice(tc.giveType, tc.giveLen, tc.giveCap).actual(), tc.want)
		}
	}

	// makeItemTypeInnerSlice
	for _, tc := range []struct {
		giveValue interface{}
		giveG     bool
		giveLen   int
		giveCap   int
		want      interface{}
		wantPanic bool
	}{
		{nil, true, 0, 0, nil, true},
		{0, false, -1, 0, []interface{}{}, false}, // no panic
		{0, false, 0, 0, []interface{}{}, false},
		{0, false, 1, 1, []interface{}{nil}, false},
		{0, false, 3, 0, []interface{}{nil, nil, nil}, false},
		{struct{}{}, false, -1, 0, []interface{}{}, false}, // no panic
		{struct{}{}, false, 0, 0, []interface{}{}, false},
		{struct{}{}, false, 1, 1, []interface{}{nil}, false},
		{struct{}{}, false, 3, 0, []interface{}{nil, nil, nil}, false},
		{0, true, -1, 0, []int{}, false}, // no panic
		{0, true, 0, 0, []int{}, false},
		{0, true, 1, 1, []int{0}, false},
		{0, true, 3, 0, []int{0, 0, 0}, false},
		{struct{}{}, true, -1, 0, []struct{}{}, false}, // no panic
		{struct{}{}, true, 0, 0, []struct{}{}, false},
		{struct{}{}, true, 1, 1, []struct{}{{}}, false},
		{struct{}{}, true, 3, 1, []struct{}{{}, {}, {}}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { makeItemTypeInnerSlice(tc.giveValue, tc.giveLen, tc.giveCap, tc.giveG) })
		} else {
			xtesting.Equal(t, makeItemTypeInnerSlice(tc.giveValue, tc.giveLen, tc.giveCap, tc.giveG).actual(), tc.want)
		}
	}

	// cloneInnerSliceItems
	typ := reflect.TypeOf([]int{})
	for _, tc := range []struct {
		give      innerSlice
		giveExtra int
		want      interface{}
		wantCap   int
		wantPanic bool
	}{
		{nil, 0, nil, 0, true},

		{&interfaceItemSlice{[]interface{}{}}, 0, []interface{}{}, 0, false},
		{&interfaceItemSlice{[]interface{}{1}}, 1, []interface{}{1}, 2, false},
		{&interfaceItemSlice{[]interface{}{1, 1, 1}}, -1, []interface{}{1, 1, 1}, 3, false},
		{&interfaceItemSlice{[]interface{}{1, nil, "2", false, 3.3}}, 20, []interface{}{1, nil, "2", false, 3.3}, 25, false},

		{&interfaceWrappedSlice{reflect.ValueOf([]int{}), typ}, 0, []int{}, 0, false},
		{&interfaceWrappedSlice{reflect.ValueOf([]int{1}), typ}, 1, []int{1}, 2, false},
		{&interfaceWrappedSlice{reflect.ValueOf([]int{1, 1, 1}), typ}, -1, []int{1, 1, 1}, 3, false},
		{&interfaceWrappedSlice{reflect.ValueOf([]int{1, 3, 0, 2}), typ}, 5, []int{1, 3, 0, 2}, 9, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { cloneInnerSliceItems(tc.give, tc.giveExtra) })
		} else {
			ii := cloneInnerSliceItems(tc.give, tc.giveExtra)
			xtesting.Equal(t, ii.actual(), tc.want)
			xtesting.Equal(t, ii.capacity(), tc.wantCap)
		}
	}
}

// =============================
// testing on exported functions
// =============================

func TestShuffle(t *testing.T) {
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
		give     []interface{}
		giveLess Lesser
		want     []interface{}
	}{
		{[]interface{}{}, le, []interface{}{}},
		{[]interface{}{0}, le, []interface{}{0}},
		{[]interface{}{1, 1, 1}, le, []interface{}{1, 1, 1}},
		{[]interface{}{4, 3, 2, 1}, le, []interface{}{1, 2, 3, 4}},
		{[]interface{}{8, 1, 6, 8, 1, 2}, le, []interface{}{1, 1, 2, 6, 8, 8}},
	} {
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

	for _, tc := range []struct {
		give     []int
		giveLess Lesser
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
		result := SortG(tc.give, tc.giveLess)
		xtesting.Equal(t, tc.give, me)
		xtesting.Equal(t, result, tc.want)
		SortSelfG(tc.give, tc.giveLess)
		xtesting.Equal(t, tc.give, tc.want)
	}
}

func TestStableSort(t *testing.T) {
	type tuple = [2]int
	new1 := func(v int) tuple { return tuple{v, 0} }
	new2 := func(v, o int) tuple { return tuple{v, o} }
	le := func(i, j interface{}) bool { return i.(tuple)[0] < j.(tuple)[0] }

	for _, tc := range []struct {
		give     []interface{}
		giveLess Lesser
		want     []interface{}
	}{
		{[]interface{}{}, le, []interface{}{}},
		{[]interface{}{new1(0)}, le, []interface{}{new1(0)}},
		{[]interface{}{new2(1, 3), new2(1, 2), new2(1, 1)}, le,
			[]interface{}{new2(1, 3), new2(1, 2), new2(1, 1)}},
		{[]interface{}{new1(4), new1(3), new1(2), new1(1)}, le,
			[]interface{}{new1(1), new1(2), new1(3), new1(4)}},
		{[]interface{}{new2(8, 2), new2(1, 2), new1(6), new2(8, 1), new2(1, 1), new1(2)}, le,
			[]interface{}{new2(1, 2), new2(1, 1), new1(2), new1(6), new2(8, 2), new2(8, 1)}},
	} {
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

	for _, tc := range []struct {
		give     []tuple
		giveLess Lesser
		want     []tuple
	}{
		{[]tuple{}, le, []tuple{}},
		{[]tuple{new1(0)}, le, []tuple{new1(0)}},
		{[]tuple{new2(1, 3), new2(1, 2), new2(1, 1)}, le,
			[]tuple{new2(1, 3), new2(1, 2), new2(1, 1)}},
		{[]tuple{new1(4), new1(3), new1(2), new1(1)}, le,
			[]tuple{new1(1), new1(2), new1(3), new1(4)}},
		{[]tuple{new2(8, 2), new2(1, 2), new1(6), new2(8, 1), new2(1, 1), new1(2)}, le,
			[]tuple{new2(1, 2), new2(1, 1), new1(2), new1(6), new2(8, 2), new2(8, 1)}},
	} {
		me := make([]tuple, 0, len(tc.give))
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

type testStruct struct {
	value int
}

func (t testStruct) String() string {
	return strconv.Itoa(t.value)
}

func newTestStruct(value int) testStruct {
	return testStruct{value: value}
}

func newTestSlice1(s []interface{}) []interface{} {
	newSlice := make([]interface{}, len(s))
	for idx, item := range s {
		newSlice[idx] = newTestStruct(item.(int))
	}
	return newSlice
}

func newTestSlice2(s []int) []testStruct {
	newSlice := make([]testStruct, len(s))
	for idx, item := range s {
		newSlice[idx] = newTestStruct(item)
	}
	return newSlice
}

func testToItfSlice(s []interface{}) []interface{} {
	out := make([]interface{}, len(s))
	for idx, item := range s {
		out[idx] = item.(testStruct).value
	}
	return out
}

func testToIntSlice(s interface{}) []int {
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
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
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
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, IndexOfWithG(give, giveValue, eq), tc.want)
	}
}

func TestLastIndexOf(t *testing.T) {
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
		{s1, 1, 3},
		{s1, 2, 4},
		{s1, 3, 5},
		{s1, 4, -1},
		{s1, 5, 1},
		{s1, 6, -1},
	} {
		xtesting.Equal(t, LastIndexOf(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, LastIndexOfWith(give, giveValue, eq), tc.want)
	}

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
		xtesting.Equal(t, LastIndexOfG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, LastIndexOfWithG(give, giveValue, eq), tc.want)
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
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
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
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
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
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
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
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
		xtesting.Equal(t, CountWithG(give, giveValue, eq), tc.want)
	}
}

func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		give       []interface{}
		giveValues []interface{}
		giveIndex  int
		want       []interface{}
	}{
		{[]interface{}{}, []interface{}{}, -2, []interface{}{}},
		{[]interface{}{}, []interface{}{1, 2}, -1, []interface{}{1, 2}},
		{[]interface{}{}, []interface{}{0, 0, 0}, 0, []interface{}{0, 0, 0}},
		{[]interface{}{}, []interface{}{3}, 1, []interface{}{3}},
		{[]interface{}{1}, []interface{}{9}, -1, []interface{}{9, 1}},
		{[]interface{}{1}, []interface{}{9, 9, 9}, 0, []interface{}{9, 9, 9, 1}},
		{[]interface{}{1}, []interface{}{}, 1, []interface{}{1}},
		{[]interface{}{1}, []interface{}{0, 9}, 2, []interface{}{1, 0, 9}},
		{[]interface{}{1, 2}, []interface{}{-1}, -1, []interface{}{-1, 1, 2}},
		{[]interface{}{1, 2}, []interface{}{9, 9}, 0, []interface{}{9, 9, 1, 2}},
		{[]interface{}{1, 2}, []interface{}{3, 2, 1}, 1, []interface{}{1, 3, 2, 1, 2}},
		{[]interface{}{1, 2}, []interface{}{9, 9, 9}, 2, []interface{}{1, 2, 9, 9, 9}},
		{[]interface{}{1, 2, 3}, []interface{}{nil}, -1, []interface{}{nil, 1, 2, 3}},
		{[]interface{}{1, 2, 3}, []interface{}{"9", "8", "7"}, 0, []interface{}{"9", "8", "7", 1, 2, 3}},
		{[]interface{}{1, 2, 3}, []interface{}{}, 1, []interface{}{1, 2, 3}},
		{[]interface{}{1, 2, 3}, []interface{}{-2, -1}, 2, []interface{}{1, 2, -2, -1, 3}},
		{[]interface{}{1, 2, 3}, []interface{}{0, 9999, 999, 99, 9}, 4, []interface{}{1, 2, 3, 0, 9999, 999, 99, 9}},
	} {
		xtesting.Equal(t, Insert(tc.give, tc.giveIndex, tc.giveValues...), tc.want)
		xtesting.Equal(t, InsertSelf(tc.give, tc.giveIndex, tc.giveValues...), tc.want)
	}

	for _, tc := range []struct {
		give       []int
		giveValues []interface{} // <<<
		giveIndex  int
		want       []int
	}{
		{[]int{}, []interface{}{}, -2, []int{}},
		{[]int{}, []interface{}{1, 2}, -1, []int{1, 2}},
		{[]int{}, []interface{}{0, 0, 0}, 0, []int{0, 0, 0}},
		{[]int{}, []interface{}{3}, 1, []int{3}},
		{[]int{1}, []interface{}{9}, -1, []int{9, 1}},
		{[]int{1}, []interface{}{9, 9, 9}, 0, []int{9, 9, 9, 1}},
		{[]int{1}, []interface{}{}, 1, []int{1}},
		{[]int{1}, []interface{}{0, 9}, 2, []int{1, 0, 9}},
		{[]int{1, 2}, []interface{}{-1}, -1, []int{-1, 1, 2}},
		{[]int{1, 2}, []interface{}{9, 9}, 0, []int{9, 9, 1, 2}},
		{[]int{1, 2}, []interface{}{3, 2, 1}, 1, []int{1, 3, 2, 1, 2}},
		{[]int{1, 2}, []interface{}{9, 9, 9}, 2, []int{1, 2, 9, 9, 9}},
		{[]int{1, 2, 3}, []interface{}{-9}, -1, []int{-9, 1, 2, 3}},
		{[]int{1, 2, 3}, []interface{}{9, 8, 7}, 0, []int{9, 8, 7, 1, 2, 3}},
		{[]int{1, 2, 3}, []interface{}{}, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []interface{}{-2, -1}, 2, []int{1, 2, -2, -1, 3}},
		{[]int{1, 2, 3}, []interface{}{0, 9999, 999, 99, 9}, 4, []int{1, 2, 3, 0, 9999, 999, 99, 9}},
	} {
		xtesting.Equal(t, InsertG(tc.give, tc.giveIndex, tc.giveValues...), tc.want)
		xtesting.Equal(t, InsertSelfG(tc.give, tc.giveIndex, tc.giveValues...), tc.want)
	}

	xtesting.NotPanic(t, func() { InsertG([]int{}, 0, nil) })
	xtesting.Panic(t, func() { InsertG([]int{}, 0, []uint{}) })
	xtesting.Panic(t, func() { InsertG([]int{}, 0, []interface{}{1, 2, 3}) })

	give1 := append(make([]interface{}, 0, 6), 1, 2, 3)
	addr1 := (*reflect.SliceHeader)(unsafe.Pointer(&give1)).Data
	give1_ := Insert(give1, 0)
	xtesting.NotEqual(t, addr1, (*reflect.SliceHeader)(unsafe.Pointer(&give1_)).Data)
	give1 = InsertSelf(give1, 1, 4, 5)
	xtesting.Equal(t, cap(give1), 6)
	xtesting.Equal(t, addr1, (*reflect.SliceHeader)(unsafe.Pointer(&give1)).Data)
	give1 = InsertSelf(give1, 0, 6, 7, 8)
	xtesting.NotEqual(t, cap(give1), 6)
	xtesting.NotEqual(t, addr1, (*reflect.SliceHeader)(unsafe.Pointer(&give1)).Data)

	give2 := append(make([]int, 0, 6), 1, 2, 3)
	addr2 := (*reflect.SliceHeader)(unsafe.Pointer(&give2)).Data
	addr2_ := InsertG(give2, 0)
	xtesting.NotEqual(t, addr2, (*reflect.SliceHeader)(unsafe.Pointer(&addr2_)).Data)
	give2 = InsertSelfG(give2, 1, 4, 5).([]int)
	xtesting.Equal(t, cap(give2), 6)
	xtesting.Equal(t, addr2, (*reflect.SliceHeader)(unsafe.Pointer(&give2)).Data)
	give2 = InsertSelfG(give2, 0, 4, 5).([]int)
	xtesting.NotEqual(t, cap(give2), 6)
	xtesting.NotEqual(t, addr2, (*reflect.SliceHeader)(unsafe.Pointer(&give2)).Data)
}

func TestDelete(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		giveN     int
		want      []interface{}
	}{
		{[]interface{}{}, 0, 1, []interface{}{}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, -1, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 0, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, 1, []interface{}{5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, 2, []interface{}{5, 2, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, 1, []interface{}{1, 5, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, 2, []interface{}{1, 5, 1, 5, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, -1, []interface{}{1, 5, 1, 5, 6, 3}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 3, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 4, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 5, 1, []interface{}{1, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 6, 1, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 7, 1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, Delete(tc.give, tc.giveValue, tc.giveN), tc.want)
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
		d1 := DeleteSelf(tc.give, tc.giveValue, tc.giveN)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToItfSlice(DeleteWith(give, giveValue, tc.giveN, eq)), tc.want)
		d2 := DeleteSelfWith(give, giveValue, tc.giveN, eq)
		xtesting.Equal(t, testToItfSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}

	for _, tc := range []struct {
		give      []int
		giveValue int
		giveN     int
		want      []int
	}{
		{[]int{}, 0, 1, []int{}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, -1, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 0, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, 1, []int{5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, 2, []int{5, 2, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, 1, []int{1, 5, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, 2, []int{1, 5, 1, 5, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, -1, []int{1, 5, 1, 5, 6, 3}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 3, 1, []int{1, 5, 2, 1, 5, 2, 6, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 4, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 5, 1, []int{1, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 6, 1, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 7, 1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, DeleteG(tc.give, tc.giveValue, tc.giveN), tc.want)
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
		d1 := DeleteSelfG(tc.give, tc.giveValue, tc.giveN).([]int)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToIntSlice(DeleteWithG(give, giveValue, tc.giveN, eq)), tc.want)
		d2 := DeleteSelfWithG(give, giveValue, tc.giveN, eq).([]testStruct)
		xtesting.Equal(t, testToIntSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}
}

func TestDeleteAll(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give      []interface{}
		giveValue int
		want      []interface{}
	}{
		{[]interface{}{}, 0, []interface{}{}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, -1, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 0, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, []interface{}{5, 2, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, []interface{}{1, 5, 1, 5, 6, 3}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 3, []interface{}{1, 5, 2, 1, 5, 2, 6, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 4, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 5, []interface{}{1, 2, 1, 2, 6, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 6, []interface{}{1, 5, 2, 1, 5, 2, 3, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, 7, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, DeleteAll(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice1(tc.give), newTestStruct(tc.giveValue)
		d1 := DeleteAllSelf(tc.give, tc.giveValue)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToItfSlice(DeleteAllWith(give, giveValue, eq)), tc.want)
		d2 := DeleteAllSelfWith(give, giveValue, eq)
		xtesting.Equal(t, testToItfSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}

	for _, tc := range []struct {
		give      []int
		giveValue int
		want      []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, -1, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 0, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 1, []int{5, 2, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 2, []int{1, 5, 1, 5, 6, 3}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 3, []int{1, 5, 2, 1, 5, 2, 6, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 4, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 5, []int{1, 2, 1, 2, 6, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 6, []int{1, 5, 2, 1, 5, 2, 3, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, 7, []int{1, 5, 2, 1, 5, 2, 6, 3, 2}},
	} {
		xtesting.Equal(t, DeleteAllG(tc.give, tc.giveValue), tc.want)
		give, giveValue := newTestSlice2(tc.give), newTestStruct(tc.giveValue)
		d1 := DeleteAllSelfG(tc.give, tc.giveValue).([]int)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToIntSlice(DeleteAllWithG(give, giveValue, eq)), tc.want)
		d2 := DeleteAllSelfWithG(give, giveValue, eq).([]testStruct)
		xtesting.Equal(t, testToIntSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}
}

func TestContainsAll(t *testing.T) {
	s1 := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	s2 := []int{1, 5, 2, 1, 5, 2, 6, 3, 2}
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  bool
	}{
		{[]interface{}{}, []interface{}{}, true},
		{[]interface{}{}, []interface{}{1, 1, 1}, false},
		{s1, []interface{}{}, true},
		{s1, []interface{}{1}, true},
		{s1, []interface{}{1, 0}, false},
		{s1, []interface{}{5, 2, 1}, true},
		{s1, []interface{}{5, 5, 5, 5}, true},
		{s1, []interface{}{2, 2, 2, 1, 5, 2, 1, 5, 2, 6, 3, 2}, true},
		{s1, []interface{}{1, 2, 3, 4, 5, 6}, false},
	} {
		xtesting.Equal(t, ContainsAll(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, ContainsAllWith(give1, give2, eq), tc.want)
	}

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  bool
	}{
		{[]int{}, []int{}, true},
		{[]int{}, []int{1, 1, 1}, false},
		{s2, []int{}, true},
		{s2, []int{1}, true},
		{s2, []int{1, 0}, false},
		{s2, []int{5, 2, 1}, true},
		{s2, []int{5, 5, 5, 5}, true},
		{s2, []int{2, 2, 2, 1, 5, 2, 1, 5, 2, 6, 3, 2}, true},
		{s2, []int{1, 2, 3, 4, 5, 6}, false},
	} {
		xtesting.Equal(t, ContainsAllG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, ContainsAllWithG(give1, give2, eq), tc.want)
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
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, testToItfSlice(DiffWith(give1, give2, eq)), tc.want)
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
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, testToIntSlice(DiffWithG(give1, give2, eq)), tc.want)
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
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, testToItfSlice(UnionWith(give1, give2, eq)), tc.want)
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
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, testToIntSlice(UnionWithG(give1, give2, eq)), tc.want)
	}
}

func TestIntersect(t *testing.T) {
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
		xtesting.Equal(t, Intersect(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, testToItfSlice(IntersectWith(give1, give2, eq)), tc.want)
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
		xtesting.Equal(t, IntersectG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, testToIntSlice(IntersectWithG(give1, give2, eq)), tc.want)
	}
}

func TestDeduplicate(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 1, 1}, []interface{}{1}},
		{[]interface{}{1, 1, 2, 1}, []interface{}{1, 2}},
		{[]interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}, []interface{}{1, 5, 2, 6, 3}},
	} {
		xtesting.Equal(t, Deduplicate(tc.give), tc.want)
		give := newTestSlice1(tc.give)
		d1 := DeduplicateSelf(tc.give)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToItfSlice(DeduplicateWith(give, eq)), tc.want)
		d2 := DeduplicateSelfWith(give, eq)
		xtesting.Equal(t, testToItfSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}

	for _, tc := range []struct {
		give []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{1, 1, 2, 1}, []int{1, 2}},
		{[]int{1, 5, 2, 1, 5, 2, 6, 3, 2}, []int{1, 5, 2, 6, 3}},
	} {
		xtesting.Equal(t, DeduplicateG(tc.give), tc.want)
		give := newTestSlice2(tc.give)
		d1 := DeduplicateSelfG(tc.give).([]int)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToIntSlice(DeduplicateWithG(give, eq)), tc.want)
		d2 := DeduplicateSelfWithG(give, eq).([]testStruct)
		xtesting.Equal(t, testToIntSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}
}

func TestCompact(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 1, 1}, []interface{}{1}},
		{[]interface{}{2, 2, 1, 1, 1, 2, 1, 3}, []interface{}{2, 1, 2, 1, 3}},
		{[]interface{}{1, 5, 5, 2, 1, 5, 2, 2, 6, 6, 6, 3, 2, 1, 1}, []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2, 1}},
	} {
		xtesting.Equal(t, Compact(tc.give), tc.want)
		give := newTestSlice1(tc.give)
		d1 := CompactSelf(tc.give)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToItfSlice(CompactWith(give, eq)), tc.want)
		d2 := CompactSelfWith(give, eq)
		xtesting.Equal(t, testToItfSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}

	for _, tc := range []struct {
		give []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{2, 2, 1, 1, 1, 2, 1, 3}, []int{2, 1, 2, 1, 3}},
		{[]int{1, 5, 5, 2, 1, 5, 2, 2, 6, 6, 6, 3, 2, 1, 1}, []int{1, 5, 2, 1, 5, 2, 6, 3, 2, 1}},
	} {
		xtesting.Equal(t, CompactG(tc.give), tc.want)
		give := newTestSlice2(tc.give)
		d1 := CompactSelfG(tc.give).([]int)
		xtesting.Equal(t, d1, tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d1)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&tc.give)).Data)
		xtesting.Equal(t, testToIntSlice(CompactWithG(give, eq)), tc.want)
		d2 := CompactSelfWithG(give, eq).([]testStruct)
		xtesting.Equal(t, testToIntSlice(d2), tc.want)
		xtesting.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&d2)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&give)).Data)
	}
}

func TestEqual(t *testing.T) {
	eq := func(i, j interface{}) bool { return i.(testStruct).value == j.(testStruct).value }

	for _, tc := range []struct {
		give1 []interface{}
		give2 []interface{}
		want  bool
	}{
		{[]interface{}{}, []interface{}{}, true},
		{[]interface{}{0}, []interface{}{}, false},
		{[]interface{}{}, []interface{}{0}, false},
		{[]interface{}{0}, []interface{}{0}, true},
		{[]interface{}{1, 1, 1}, []interface{}{1}, false},
		{[]interface{}{1}, []interface{}{1, 1, 1}, false},
		{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}, true},
		{[]interface{}{1, 2, 1}, []interface{}{1, 1, 2}, false},
		{[]interface{}{1, 1, 2, 3}, []interface{}{1, 2, 3, 1}, false},
		{[]interface{}{1, 1, 2, 2}, []interface{}{1, 1, 2}, false},
		{[]interface{}{1, 1, 2, 3}, []interface{}{1, 1, 2, 3}, true},
	} {
		xtesting.Equal(t, Equal(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, EqualWith(give1, give2, eq), tc.want)
	}

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  bool
	}{
		{[]int{}, []int{}, true},
		{[]int{0}, []int{}, false},
		{[]int{}, []int{0}, false},
		{[]int{0}, []int{0}, true},
		{[]int{1, 1, 1}, []int{1}, false},
		{[]int{1}, []int{1, 1, 1}, false},
		{[]int{1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{1, 2, 1}, []int{1, 1, 2}, false},
		{[]int{1, 1, 2, 3}, []int{1, 2, 3, 1}, false},
		{[]int{1, 1, 2, 2}, []int{1, 1, 2}, false},
		{[]int{1, 1, 2, 3}, []int{1, 1, 2, 3}, true},
	} {
		xtesting.Equal(t, EqualG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, EqualWithG(give1, give2, eq), tc.want)
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
		{[]interface{}{}, []interface{}{0}, false},
		{[]interface{}{0}, []interface{}{0}, true},
		{[]interface{}{1, 1, 1}, []interface{}{1}, false},
		{[]interface{}{1}, []interface{}{1, 1, 1}, false},
		{[]interface{}{1, 1, 1}, []interface{}{1, 1, 1}, true},
		{[]interface{}{1, 2, 1}, []interface{}{1, 1, 2}, true},
		{[]interface{}{1, 2, 3}, []interface{}{1, 2, 2}, false},
		{[]interface{}{1, 2, 2}, []interface{}{1, 2, 3}, false},
	} {
		xtesting.Equal(t, ElementMatch(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice1(tc.give1), newTestSlice1(tc.give2)
		xtesting.Equal(t, ElementMatchWith(give1, give2, eq), tc.want)
	}

	for _, tc := range []struct {
		give1 []int
		give2 []int
		want  bool
	}{
		{[]int{}, []int{}, true},
		{[]int{0}, []int{}, false},
		{[]int{}, []int{0}, false},
		{[]int{0}, []int{0}, true},
		{[]int{1, 1, 1}, []int{1}, false},
		{[]int{1}, []int{1, 1, 1}, false},
		{[]int{1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{1, 2, 1}, []int{1, 1, 2}, true},
		{[]int{1, 2, 3}, []int{1, 2, 2}, false},
		{[]int{1, 2, 2}, []int{1, 2, 3}, false},
	} {
		xtesting.Equal(t, ElementMatchG(tc.give1, tc.give2), tc.want)
		give1, give2 := newTestSlice2(tc.give1), newTestSlice2(tc.give2)
		xtesting.Equal(t, ElementMatchWithG(give1, give2, eq), tc.want)
	}
}

type testError struct{ m string }

func (t *testError) Error() string { return t.m }

func TestRepeat(t *testing.T) {
	for _, tc := range []struct {
		giveValue interface{}
		giveCount uint
		want      []interface{}
	}{
		{nil, 0, []interface{}{}},
		{nil, 2, []interface{}{nil, nil}},
		{true, 0, []interface{}{}},
		{true, 1, []interface{}{true}},
		{5, 5, []interface{}{5, 5, 5, 5, 5}},
		{"", 5, []interface{}{"", "", "", "", ""}},
		{uint(0), 2, []interface{}{uint(0), uint(0)}},
		{[]float64{1.1, 2.2}, 3, []interface{}{[]float64{1.1, 2.2}, []float64{1.1, 2.2}, []float64{1.1, 2.2}}},
		{error(nil), 2, []interface{}{nil, nil}},                                           // <<<
		{error((*testError)(nil)), 2, []interface{}{(*testError)(nil), (*testError)(nil)}}, // <<<
		{(*testError)(nil), 2, []interface{}{(*testError)(nil), (*testError)(nil)}},
		{&testError{"test"}, 2, []interface{}{&testError{"test"}, &testError{"test"}}},
	} {
		xtesting.Equal(t, Repeat(tc.giveValue, tc.giveCount), tc.want)
	}

	for _, tc := range []struct {
		giveValue interface{}
		giveCount uint
		want      interface{}
	}{
		{nil, 0, []interface{}{}},
		{nil, 2, []interface{}{nil, nil}},
		{true, 0, []bool{}},
		{true, 1, []bool{true}},
		{5, 5, []int{5, 5, 5, 5, 5}},
		{"", 5, []string{"", "", "", "", ""}},
		{uint(0), 2, []uint{uint(0), uint(0)}},
		{[]float64{1.1, 2.2}, 3, [][]float64{{1.1, 2.2}, {1.1, 2.2}, {1.1, 2.2}}},
		{error(nil), 2, []interface{}{nil, nil}},                                          // <<<
		{error((*testError)(nil)), 2, []*testError{(*testError)(nil), (*testError)(nil)}}, // <<<
		{(*testError)(nil), 2, []*testError{(*testError)(nil), (*testError)(nil)}},
		{&testError{"test"}, 2, []*testError{{"test"}, {"test"}}},
	} {
		xtesting.Equal(t, RepeatG(tc.giveValue, tc.giveCount), tc.want)
	}
}
