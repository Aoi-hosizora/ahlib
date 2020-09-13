package xslice

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestSti(t *testing.T) {
	xtesting.Equal(t, Sti([]string{"123", "456"}), []interface{}{"123", "456"})
	// xtesting.Equal(t, Sti(""), []interface{}(nil))
	xtesting.Equal(t, Sti([]string{}), []interface{}{})
	// xtesting.Equal(t, Sti("") == nil, true)
	xtesting.Equal(t, Sti(nil) == nil, true)

	num := 2000000

	start := time.Now()
	arr := make([]int, 0)
	for i := 0; i < num; i++ {
		arr = append(arr, i)
	}
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	itf := make([]interface{}, num)
	for i := 0; i < num; i++ {
		itf[i] = arr[i]
	}
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	_ = Sti(arr)
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	v := reflect.ValueOf(arr)
	itf = make([]interface{}, num)
	for i := 0; i < num; i++ {
		itf[i] = v.Index(i).Interface()
	}
	log.Println(time.Now().Sub(start).String())
}

func TestIts(t *testing.T) {
	xtesting.Equal(t, Its([]interface{}{"123", "456"}, ""), []string{"123", "456"})
	xtesting.Equal(t, Its(nil, 0), nil)
	// xtesting.Equal(t, Its(nil, nil), nil)

	log.Println(ItsToString([]interface{}{1, 2, 3}))
	log.Println(ItsOfInt([]interface{}{}))
	log.Println(ItsOfInt([]interface{}{1, 2}))
	log.Println(ItsOfInt8([]interface{}{int8(1), int8(2)}))
	log.Println(ItsOfInt16([]interface{}{int16(1), int16(2)}))
	log.Println(ItsOfInt32([]interface{}{int32(1), int32(2)}))
	log.Println(ItsOfInt64([]interface{}{int64(1), int64(2)}))
	log.Println(ItsOfUint([]interface{}{uint(1), uint(2)}))
	log.Println(ItsOfUint8([]interface{}{uint8(1), uint8(2)}))
	log.Println(ItsOfUint16([]interface{}{uint16(1), uint16(2)}))
	log.Println(ItsOfUint32([]interface{}{uint32(1), uint32(2)}))
	log.Println(ItsOfUint64([]interface{}{uint64(1), uint64(2)}))
	log.Println(ItsOfFloat32([]interface{}{float32(0.1), float32(2.0)}))
	log.Println(ItsOfFloat64([]interface{}{0.1, 2.0}))
	log.Println(ItsOfString([]interface{}{"1", "2"}))
	log.Println(ItsOfByte([]interface{}{byte('1'), byte('2')}))
	log.Println(ItsOfRune([]interface{}{'1', '2'}))

	num := 2000000
	arr := make([]interface{}, num)
	for i := 0; i < num; i++ {
		arr[i] = i
	}

	start := time.Now()
	arr2 := make([]int, num)
	for i := 0; i < num; i++ {
		arr2[i] = arr[i].(int)
	}
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	_ = Its(arr, 0)
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	_ = ItsOfInt(arr)
	log.Println(time.Now().Sub(start).String())
}

func TestShuffle(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	Shuffle(a)
	log.Println(a)
	Shuffle(a)
	log.Println(a)

	Shuffle([]interface{}{})
}

func TestShuffleNew(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	log.Println(ShuffleNew(a))
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})

	ShuffleNew([]interface{}{})
}

func TestReverse(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	Reverse(a)
	xtesting.Equal(t, a, []interface{}{4, 3, 2, 1})
	Reverse(a)
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})

	Reverse([]interface{}{})
}

func TestReverseNew(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	xtesting.Equal(t, ReverseNew(a), []interface{}{4, 3, 2, 1})
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})
	xtesting.Equal(t, ReverseNew(ReverseNew(a)), []interface{}{1, 2, 3, 4})
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})

	ReverseNew([]interface{}{})
}

func TestForEach(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"1", "2", "3", "4", "5"}
	ForEach(Sti(s1), func(i interface{}) {
		log.Println(i)
	})
	ForEach(Sti(s2), func(i interface{}) {
		log.Println(i)
	})
}

func TestGoForEach(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"1", "2", "3", "4", "5"}
	GoForEach(Sti(s1), func(i interface{}) {
		log.Println(i)
	})
	GoForEach(Sti(s2), func(i interface{}) {
		log.Println(i)
	})
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"1", "2", "3", "4", "5"}
	s3 := []int{2, 3, 4, 5, 6}
	s4 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	xtesting.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return strconv.Itoa(i.(int)) }), Sti(s2))
	xtesting.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return i.(int) + 1 }), Sti(s3))
	xtesting.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return float32(i.(int)) }), Sti(s4))
}

func TestIndexOf(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	xtesting.Equal(t, IndexOf(Sti(s), 1), 0)
	xtesting.Equal(t, IndexOf(Sti(s), 6), -1)
	xtesting.Equal(t, IndexOf(Sti(s), nil), -1)
	xtesting.Equal(t, IndexOfWith([]interface{}{1, 2, 3, 4, 5}, 3, func(i, j interface{}) bool {
		return i.(int) == j.(int)-1
	}), 1)
}

func TestContains(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	xtesting.Equal(t, Contains(Sti(s), 1), true)
	xtesting.Equal(t, Contains(Sti(s), 6), false)
	xtesting.Equal(t, Contains(Sti(s), nil), false)
	xtesting.Equal(t, ContainsWith(Sti(s), 4, func(i, j interface{}) bool {
		return i.(int) == j.(int)-1
	}), true)
}

func TestCount(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	xtesting.Equal(t, Count(s, 1), 2)
	xtesting.Equal(t, Count(s, 2), 3)
	xtesting.Equal(t, Count(s, 3), 1)
	xtesting.Equal(t, Count(s, 4), 0)
	xtesting.Equal(t, Count(s, 5), 2)
	xtesting.Equal(t, Count(s, 6), 1)
	xtesting.Equal(t, Count(s, ""), 0)
}

func TestDelete(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3, 1}

	s = ItsOfInt(Delete(Sti(s), 1, 1))
	xtesting.Equal(t, s, []int{5, 2, 1, 2, 3, 1})
	s = ItsOfInt(Delete(Sti(s), 1, 2))
	xtesting.Equal(t, s, []int{5, 2, 2, 3})
	s = ItsOfInt(Delete(Sti(s), 6, 1))
	xtesting.Equal(t, s, []int{5, 2, 2, 3})
	s = ItsOfInt(Delete(Sti(s), 2, -1))
	xtesting.Equal(t, s, []int{5, 3})
	s = ItsOfInt(Delete(Sti(s), nil, -1))
	xtesting.Equal(t, s, []int{5, 3})

	ss := ItsOfInt(Delete(nil, 2, -1))
	xtesting.Equal(t, ss == nil, false)
}

func TestDeleteAll(t *testing.T) {
	xtesting.Equal(t, DeleteAll(Sti([]int{1, 5, 2, 1, 2, 3, 1}), 1), Sti([]int{5, 2, 2, 3}))
}

func TestDiff(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	xtesting.Equal(t, Diff(Sti(slice1), Sti(slice2)), Sti([]int{2, 3, 3}))
}

func TestUnion(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	xtesting.Equal(t, Union(Sti(slice1), Sti(slice2)), Sti([]int{1, 2, 1, 3, 4, 3, 5, 6}))
}

func TestIntersection(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	xtesting.Equal(t, Intersection(Sti(slice1), Sti(slice2)), Sti([]int{1, 1, 4}))
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 5, 6, 7}
	s2 := []int{2, 1, 7, 5, 6}
	s3 := []int{1, 5, 6, 7}
	s4 := []int{1, 7, 6, 5}
	s5 := []int{1, 2, 2, 5, 6, 7}
	xtesting.Equal(t, Equal(Sti(s1), Sti(s2)), true)
	xtesting.Equal(t, Equal(Sti(s1), Sti(s3)), false)
	xtesting.Equal(t, Equal(Sti(s1), Sti(s4)), false)
	xtesting.Equal(t, Equal(Sti(s1), Sti(s5)), false)
	xtesting.Equal(t, Equal(Sti(s3), Sti(s4)), true)
}

func TestToSet(t *testing.T) {
	s1 := []interface{}{1, 2, 3, 1, 2, 3, 4, 5, 6}
	s2 := make([]interface{}, 0)
	s3 := []interface{}{1}
	s4 := []interface{}{1, 1, 1, 1, 1}

	xtesting.Equal(t, ToSet(s1), []interface{}{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, ToSet(s2), []interface{}{})
	xtesting.Equal(t, ToSet(s3), []interface{}{1})
	xtesting.Equal(t, ToSet(s4), []interface{}{1})
}

func TestRange(t *testing.T) {
	xtesting.Equal(t, Range(0, 10, 1), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	xtesting.Equal(t, Range(1, 10, 2), []int{1, 3, 5, 7, 9})
	xtesting.Equal(t, Range(0, 10, 2), []int{0, 2, 4, 6, 8, 10})
	xtesting.Equal(t, Range(0, 1, 2), []int{0})
}

func TestReverseRange(t *testing.T) {
	xtesting.Equal(t, ReverseRange(0, 10, 1), []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	xtesting.Equal(t, ReverseRange(1, 10, 2), []int{10, 8, 6, 4, 2})
	xtesting.Equal(t, ReverseRange(1, 9, 2), []int{9, 7, 5, 3, 1})
	xtesting.Equal(t, ReverseRange(0, 1, 2), []int{1})
}

func TestGenerate(t *testing.T) {
	xtesting.Equal(t, GenerateByIndex([]int{}, func(i int) interface{} { return strconv.Itoa(i) }), []interface{}{})
	xtesting.Equal(t, GenerateByIndex([]int{1, 5, 9}, func(i int) interface{} { return strconv.Itoa(i) }), []interface{}{"1", "5", "9"})
	xtesting.Equal(t, GenerateByIndex(Range(0, 5, 1), func(i int) interface{} { return strconv.Itoa(i) }), []interface{}{"0", "1", "2", "3", "4", "5"})
	xtesting.Equal(t, GenerateByIndex(Range(0, 5, 2), func(i int) interface{} { return i + 1 }), []interface{}{1, 3, 5})
	xtesting.Equal(t, GenerateByIndex(Range(0, 1, 2), func(i int) interface{} { return i - 1 }), []interface{}{-1})
}

func TestSort(t *testing.T) {
	a := Range(1, 50, 1)
	r1 := Sti(a)
	r2 := Sti(ReverseRange(1, 50, 1))

	for range Range(0, 4, 1) {
		aa := ShuffleNew(Sti(a))

		Sort(aa, func(i, j int) bool {
			return aa[i].(int) < aa[j].(int)
		})
		xtesting.Equal(t, aa, r1)
	}

	for range Range(0, 4, 1) {
		aa := ShuffleNew(Sti(a))

		sortInterface := ReverseSortSlice(NewSortSlice(aa, func(i, j int) bool {
			return aa[i].(int) < aa[j].(int)
		}))
		sort.Sort(sortInterface)
		xtesting.Equal(t, aa, r2)
	}
}

func TestStable(t *testing.T) {
	type typ struct {
		a int
		b string
	}
	a := []*typ{{2, "c"}, {1, "b"}, {2, "d"}, {1, "c"}, {2, "a"}, {1, "a"}, {2, "b"}}
	r := Sti([]*typ{{1, "b"}, {1, "c"}, {1, "a"}, {2, "c"}, {2, "d"}, {2, "a"}, {2, "b"}})

	for range Range(0, 4, 1) {
		aa := Sti(a)
		StableSort(aa, func(i, j int) bool {
			return aa[i].(*typ).a < aa[j].(*typ).a
		})
		xtesting.Equal(t, aa, r)
	}
}
