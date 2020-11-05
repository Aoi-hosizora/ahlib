package xslice

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSti(t *testing.T) {
	xtesting.Equal(t, Sti([]string{"123", "456"}), []interface{}{"123", "456"})
	xtesting.Equal(t, Sti([]string{}), []interface{}{})
	xtesting.Equal(t, Sti([]int{0, 0., 0.0}), []interface{}{0, 0, 0})
	xtesting.Equal(t, Sti(nil) == nil, true)
	xtesting.PanicWithValue(t, "Sti: parameter must be a slice", func() { Sti(0) })

	num := 2000000

	start := time.Now()
	arr := make([]int, 0)
	for i := 0; i < num; i++ {
		arr = append(arr, i)
	}
	log.Println("[]int:", time.Now().Sub(start).String()) // 81.7816ms

	start = time.Now()
	_ = Sti(arr)
	log.Println("Sti:", time.Now().Sub(start).String()) // 271.2762ms

	start = time.Now()
	_ = StiOfInt(arr)
	log.Println("StiOfInt:", time.Now().Sub(start).String()) // 141.623ms
}

func TestIts(t *testing.T) {
	xtesting.Equal(t, Its([]interface{}{"123", "456"}, ""), []string{"123", "456"})
	xtesting.Equal(t, Its(nil, 0), nil)
	xtesting.PanicWithValue(t, "Its: model must be non-nil", func() { xtesting.Equal(t, Its([]interface{}{}, nil), nil) })

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
	log.Println("[]int:", time.Now().Sub(start).String()) // 23.9355ms

	start = time.Now()
	_ = Its(arr, 0)
	log.Println("Its:", time.Now().Sub(start).String()) // 230.385ms

	start = time.Now()
	_ = ItsOfInt(arr)
	log.Println("ItsOfInt:", time.Now().Sub(start).String()) // 10.9698ms
}

func TestStiOf(t *testing.T) {
	xtesting.Equal(t, StiOfInt([]int{1, 2}), []interface{}{1, 2})
	xtesting.Equal(t, StiOfInt8([]int8{1, 2}), []interface{}{int8(1), int8(2)})
	xtesting.Equal(t, StiOfInt16([]int16{1, 2}), []interface{}{int16(1), int16(2)})
	xtesting.Equal(t, StiOfInt32([]int32{1, 2}), []interface{}{int32(1), int32(2)})
	xtesting.Equal(t, StiOfInt64([]int64{1, 2}), []interface{}{int64(1), int64(2)})
	xtesting.Equal(t, StiOfUint([]uint{1, 2}), []interface{}{uint(1), uint(2)})
	xtesting.Equal(t, StiOfUint8([]uint8{1, 2}), []interface{}{uint8(1), uint8(2)})
	xtesting.Equal(t, StiOfUint16([]uint16{1, 2}), []interface{}{uint16(1), uint16(2)})
	xtesting.Equal(t, StiOfUint32([]uint32{1, 2}), []interface{}{uint32(1), uint32(2)})
	xtesting.Equal(t, StiOfUint64([]uint64{1, 2}), []interface{}{uint64(1), uint64(2)})
	xtesting.Equal(t, StiOfFloat32([]float32{0.1, 2.0}), []interface{}{float32(0.1), float32(2.0)})
	xtesting.Equal(t, StiOfFloat64([]float64{0.1, 2.0}), []interface{}{0.1, 2.0})
	xtesting.Equal(t, StiOfByte([]byte{'1', '2'}), []interface{}{byte('1'), byte('2')})
	xtesting.Equal(t, StiOfRune([]rune{'1', '2'}), []interface{}{'1', '2'})
	xtesting.Equal(t, StiOfString([]string{"1", "2"}), []interface{}{"1", "2"})
	xtesting.Equal(t, StiOfBool([]bool{false, true}), []interface{}{false, true})
}

func TestItsOf(t *testing.T) {
	xtesting.Equal(t, ItsOfInt([]interface{}{1, 2}), []int{1, 2})
	xtesting.Equal(t, ItsOfInt8([]interface{}{int8(1), int8(2)}), []int8{1, 2})
	xtesting.Equal(t, ItsOfInt16([]interface{}{int16(1), int16(2)}), []int16{1, 2})
	xtesting.Equal(t, ItsOfInt32([]interface{}{int32(1), int32(2)}), []int32{1, 2})
	xtesting.Equal(t, ItsOfInt64([]interface{}{int64(1), int64(2)}), []int64{1, 2})
	xtesting.Equal(t, ItsOfUint([]interface{}{uint(1), uint(2)}), []uint{1, 2})
	xtesting.Equal(t, ItsOfUint8([]interface{}{uint8(1), uint8(2)}), []uint8{1, 2})
	xtesting.Equal(t, ItsOfUint16([]interface{}{uint16(1), uint16(2)}), []uint16{1, 2})
	xtesting.Equal(t, ItsOfUint32([]interface{}{uint32(1), uint32(2)}), []uint32{1, 2})
	xtesting.Equal(t, ItsOfUint64([]interface{}{uint64(1), uint64(2)}), []uint64{1, 2})
	xtesting.Equal(t, ItsOfFloat32([]interface{}{float32(0.1), float32(2.0)}), []float32{0.1, 2.0})
	xtesting.Equal(t, ItsOfFloat64([]interface{}{0.1, 2.0}), []float64{0.1, 2.0})
	xtesting.Equal(t, ItsOfByte([]interface{}{byte('1'), byte('2')}), []byte{'1', '2'})
	xtesting.Equal(t, ItsOfRune([]interface{}{'1', '2'}), []rune{'1', '2'})
	xtesting.Equal(t, ItsOfString([]interface{}{"1", "2"}), []string{"1", "2"})
	xtesting.Equal(t, ItsOfBool([]interface{}{false, true}), []bool{false, true})
}

func TestShuffle(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}

	Shuffle(a)
	xtesting.ElementMatch(t, a, []interface{}{1, 2, 3, 4})
	log.Println("Shuffle:", a)

	Shuffle(a)
	xtesting.ElementMatch(t, a, []interface{}{1, 2, 3, 4})
	log.Println("Shuffle:", a)

	b := make([]interface{}, 0)
	Shuffle(b)
	xtesting.Equal(t, b, b)
}

func TestShuffleNew(t *testing.T) {
	aa := []interface{}{1, 2, 3, 4}

	a := ShuffleNew(aa)
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})
	xtesting.ElementMatch(t, aa, a)
	log.Println("ShuffleNew:", a)

	a = ShuffleNew(aa)
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})
	xtesting.ElementMatch(t, aa, a)
	log.Println("ShuffleNew:", a)

	b := make([]interface{}, 0)
	xtesting.Equal(t, b, []interface{}{})
	xtesting.Equal(t, ShuffleNew(b), b)
}

func TestReverse(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}

	Reverse(a)
	xtesting.Equal(t, a, []interface{}{4, 3, 2, 1})

	Reverse(a)
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})

	b := make([]interface{}, 0)
	Reverse(b)
	xtesting.Equal(t, b, b)
}

func TestReverseNew(t *testing.T) {
	aa := []interface{}{1, 2, 3, 4}

	a := ReverseNew(aa)
	xtesting.Equal(t, a, []interface{}{4, 3, 2, 1})
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})

	a = ReverseNew(a)
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})

	b := make([]interface{}, 0)
	xtesting.Equal(t, b, []interface{}{})
	xtesting.Equal(t, ReverseNew(b), b)
}

func TestForEach(t *testing.T) {
	ip := 0
	ForEach([]interface{}{1, 2, 3, 4, 5}, func(i interface{}) {
		ip += i.(int)
	})
	xtesting.Equal(t, ip, 15)

	sp := strings.Builder{}
	ForEach([]interface{}{"1", "2", "3", "4", "5"}, func(i interface{}) {
		sp.WriteString(i.(string))
	})
	xtesting.Equal(t, sp.String(), "12345")

	cnt := 0
	ForEach([]interface{}{}, func(interface{}) {
		cnt++
	})
	xtesting.Equal(t, cnt, 0)
}

func TestGoForEach(t *testing.T) {
	is := make([]int, 0)
	mu := sync.Mutex{}
	GoForEach([]interface{}{1, 2, 3, 4, 5}, func(i interface{}) {
		mu.Lock()
		is = append(is, i.(int))
		mu.Unlock()
	})
	xtesting.ElementMatch(t, Sti(is), []interface{}{1, 2, 3, 4, 5})

	ss := make([]string, 0)
	GoForEach([]interface{}{"1", "2", "3", "4", "5"}, func(i interface{}) {
		mu.Lock()
		ss = append(ss, i.(string))
		mu.Unlock()
	})
	xtesting.ElementMatch(t, Sti(ss), []interface{}{"1", "2", "3", "4", "5"})

	cnt := 0
	GoForEach([]interface{}{}, func(interface{}) {
		cnt++
	})
	xtesting.Equal(t, cnt, 0)
}

func TestMap(t *testing.T) {
	s1 := []interface{}{1, 2, 3, 4, 5}
	s2 := []interface{}{"1", "2", "3", "4", "5"}
	s3 := []interface{}{2, 3, 4, 5, 6}
	s4 := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	xtesting.Equal(t, Map(s1, func(i interface{}) interface{} { return strconv.Itoa(i.(int)) }), s2)
	xtesting.Equal(t, Map(s1, func(i interface{}) interface{} { return i.(int) + 1 }), s3)
	xtesting.Equal(t, Map(s1, func(i interface{}) interface{} { return float64(i.(int)) }), s4)

	xtesting.Equal(t, Map([]interface{}{}, func(i interface{}) interface{} { return nil }), []interface{}{})
}

func TestIndexOf(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 2, 3}
	xtesting.Equal(t, IndexOf(s, 0), -1)
	xtesting.Equal(t, IndexOf(s, 1), 0)
	xtesting.Equal(t, IndexOf(s, 2), 2)
	xtesting.Equal(t, IndexOf(s, 3), 5)
	xtesting.Equal(t, IndexOf(s, 4), -1)
	xtesting.Equal(t, IndexOf(s, 5), 1)
	xtesting.Equal(t, IndexOf(s, 6), -1)
	xtesting.Equal(t, IndexOf(s, nil), -1)

	eq := func(i, j interface{}) bool { return i.(int) == j.(int)-1 }
	xtesting.Equal(t, IndexOfWith(s, -1, eq), -1)
	xtesting.Equal(t, IndexOfWith(s, 0, eq), 0)
	xtesting.Equal(t, IndexOfWith(s, 1, eq), 2)
	xtesting.Equal(t, IndexOfWith(s, 2, eq), 5)
	xtesting.Equal(t, IndexOfWith(s, 3, eq), -1)
	xtesting.Equal(t, IndexOfWith(s, 4, eq), 1)
	xtesting.Equal(t, IndexOfWith(s, 5, eq), -1)
	xtesting.Equal(t, IndexOfWith(s, 6, eq), -1)
}

func TestContains(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 2, 3}
	xtesting.Equal(t, Contains(s, 0), false)
	xtesting.Equal(t, Contains(s, 1), true)
	xtesting.Equal(t, Contains(s, 2), true)
	xtesting.Equal(t, Contains(s, 3), true)
	xtesting.Equal(t, Contains(s, 4), false)
	xtesting.Equal(t, Contains(s, 5), true)
	xtesting.Equal(t, Contains(s, 6), false)
	xtesting.Equal(t, Contains(s, nil), false)

	eq := func(i, j interface{}) bool { return i.(int) == j.(int)-1 }
	xtesting.Equal(t, ContainsWith(s, -1, eq), false)
	xtesting.Equal(t, ContainsWith(s, 0, eq), true)
	xtesting.Equal(t, ContainsWith(s, 1, eq), true)
	xtesting.Equal(t, ContainsWith(s, 2, eq), true)
	xtesting.Equal(t, ContainsWith(s, 3, eq), false)
	xtesting.Equal(t, ContainsWith(s, 4, eq), true)
	xtesting.Equal(t, ContainsWith(s, 5, eq), false)
	xtesting.Equal(t, ContainsWith(s, 6, eq), false)
}

func TestCount(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 5, 2, 6, 3, 2}
	xtesting.Equal(t, Count(s, 0), 0)
	xtesting.Equal(t, Count(s, 1), 2)
	xtesting.Equal(t, Count(s, 2), 3)
	xtesting.Equal(t, Count(s, 3), 1)
	xtesting.Equal(t, Count(s, 4), 0)
	xtesting.Equal(t, Count(s, 5), 2)
	xtesting.Equal(t, Count(s, 6), 1)
	xtesting.Equal(t, Count(s, 7), 0)
	xtesting.Equal(t, Count(s, nil), 0)

	eq := func(i, j interface{}) bool { return i.(int) == j.(int)-1 }
	xtesting.Equal(t, CountWith(s, -1, eq), 0)
	xtesting.Equal(t, CountWith(s, 0, eq), 2)
	xtesting.Equal(t, CountWith(s, 1, eq), 3)
	xtesting.Equal(t, CountWith(s, 2, eq), 1)
	xtesting.Equal(t, CountWith(s, 3, eq), 0)
	xtesting.Equal(t, CountWith(s, 4, eq), 2)
	xtesting.Equal(t, CountWith(s, 5, eq), 1)
	xtesting.Equal(t, CountWith(s, 6, eq), 0)
}

func TestDelete(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 2, 3, 1}

	s = Delete(s, 0, 1)
	xtesting.Equal(t, s, []interface{}{1, 5, 2, 1, 2, 3, 1})
	s = Delete(s, 1, 1)
	xtesting.Equal(t, s, []interface{}{5, 2, 1, 2, 3, 1})
	s = Delete(s, 1, -1)
	xtesting.Equal(t, s, []interface{}{5, 2, 2, 3})
	s = Delete(s, 2, 1)
	xtesting.Equal(t, s, []interface{}{5, 2, 3})
	s = Delete(s, 2, 1)
	xtesting.Equal(t, s, []interface{}{5, 3})
	s = Delete(s, 3, 1)
	xtesting.Equal(t, s, []interface{}{5})
	s = Delete(s, 4, 1)
	xtesting.Equal(t, s, []interface{}{5})
	s = Delete(s, 5, 1)
	xtesting.Equal(t, s, []interface{}{})
	s = Delete(s, nil, 1)
	xtesting.Equal(t, s, []interface{}{})

	s = []interface{}{1, 5, 2, 1, 2, 3, 1}

	eq := func(i, j interface{}) bool { return i.(string) == strconv.Itoa(j.(int)) }
	s = DeleteWith(s, "5", -1, eq)
	xtesting.Equal(t, s, []interface{}{1, 2, 1, 2, 3, 1})
	s = DeleteWith(s, "1", 2, eq)
	xtesting.Equal(t, s, []interface{}{2, 2, 3, 1})
	s = DeleteWith(s, "1", 1, eq)
	xtesting.Equal(t, s, []interface{}{2, 2, 3})
	s = DeleteWith(s, "2", -1, eq)
	xtesting.Equal(t, s, []interface{}{3})
	s = DeleteWith(s, "3", -1, eq)
	xtesting.Equal(t, s, []interface{}{})
	s = DeleteWith(s, "4", -1, eq)
	xtesting.Equal(t, s, []interface{}{})

	xtesting.Equal(t, Delete(nil, 2, -1), []interface{}(nil))
}

func TestDeleteAll(t *testing.T) {
	s := []interface{}{1, 5, 2, 1, 2, 3, 1}

	s = DeleteAll(s, 0)
	xtesting.Equal(t, s, []interface{}{1, 5, 2, 1, 2, 3, 1})
	s = DeleteAll(s, 1)
	xtesting.Equal(t, s, []interface{}{5, 2, 2, 3})
	s = DeleteAll(s, 2)
	xtesting.Equal(t, s, []interface{}{5, 3})
	s = DeleteAll(s, 3)
	xtesting.Equal(t, s, []interface{}{5})
	s = DeleteAll(s, 4)
	xtesting.Equal(t, s, []interface{}{5})
	s = DeleteAll(s, 5)
	xtesting.Equal(t, s, []interface{}{})
	s = DeleteAll(s, nil)
	xtesting.Equal(t, s, []interface{}{})

	s = []interface{}{1, 5, 2, 1, 2, 3, 1}

	eq := func(i, j interface{}) bool { return i.(string) == strconv.Itoa(j.(int)) }
	s = DeleteAllWith(s, "5", eq)
	xtesting.Equal(t, s, []interface{}{1, 2, 1, 2, 3, 1})
	s = DeleteAllWith(s, "1", eq)
	xtesting.Equal(t, s, []interface{}{2, 2, 3})
	s = DeleteAllWith(s, "2", eq)
	xtesting.Equal(t, s, []interface{}{3})
	s = DeleteAllWith(s, "3", eq)
	xtesting.Equal(t, s, []interface{}{})
	s = DeleteAllWith(s, "4", eq)
	xtesting.Equal(t, s, []interface{}{})

	xtesting.Equal(t, DeleteAll(nil, 2), []interface{}(nil))
}

func TestDiff(t *testing.T) {
	xtesting.Equal(t, Diff([]interface{}{}, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Diff(nil, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Diff([]interface{}{}, nil), []interface{}{})

	xtesting.Equal(t, Diff([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{1}), []interface{}{4, 5, 2, 5, 4, 2})
	xtesting.Equal(t, Diff([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{4, 5, 2}), []interface{}{1})
	xtesting.Equal(t, Diff([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{4, 5, 2, 5, 1, 4, 2}), []interface{}{})
	xtesting.Equal(t, Diff([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{0}), []interface{}{4, 5, 2, 5, 1, 4, 2})

	eq := func(i, j interface{}) bool { return i.(string) == strconv.Itoa(j.(int)) }
	xtesting.Equal(t, DiffWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{1}, eq), []interface{}{"4", "5", "2", "5", "4", "2"})
	xtesting.Equal(t, DiffWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{4, 5, 2}, eq), []interface{}{"1"})
	xtesting.Equal(t, DiffWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{4, 5, 2, 5, 1, 4, 2}, eq), []interface{}{})
	xtesting.Equal(t, DiffWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{0}, eq), []interface{}{"4", "5", "2", "5", "1", "4", "2"})
}

func TestUnion(t *testing.T) {
	xtesting.Equal(t, Union([]interface{}{}, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Union(nil, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Union([]interface{}{}, nil), []interface{}{})

	xtesting.Equal(t, Union([]interface{}{}, []interface{}{1}), []interface{}{1})
	xtesting.Equal(t, Union([]interface{}{1, 1, 1, 1}, []interface{}{1, 1, 2}), []interface{}{1, 1, 1, 1, 2})
	xtesting.Equal(t, Union([]interface{}{1, 2, 3}, []interface{}{3, 2, 1}), []interface{}{1, 2, 3})
	xtesting.Equal(t, Union([]interface{}{1, 4, 3, 3, 5}, []interface{}{4, 6, 0}), []interface{}{1, 4, 3, 3, 5, 6, 0})

	eq := func(i, j interface{}) bool { return i.(string) == strconv.Itoa(j.(int)) }
	xtesting.Equal(t, UnionWith([]interface{}{}, []interface{}{1}, eq), []interface{}{1})
	xtesting.Equal(t, UnionWith([]interface{}{"1", "1", "1", "1"}, []interface{}{1, 1, 2}, eq), []interface{}{"1", "1", "1", "1", 2})
	xtesting.Equal(t, UnionWith([]interface{}{"1", "2", "3"}, []interface{}{3, 2, 1}, eq), []interface{}{"1", "2", "3"})
	xtesting.Equal(t, UnionWith([]interface{}{"1", "4", "3", "3", "5"}, []interface{}{4, 6, 0}, eq), []interface{}{"1", "4", "3", "3", "5", 6, 0})
}

func TestIntersection(t *testing.T) {
	xtesting.Equal(t, Intersection([]interface{}{}, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Intersection(nil, []interface{}{}), []interface{}{})
	xtesting.Equal(t, Intersection([]interface{}{}, nil), []interface{}{})

	xtesting.Equal(t, Intersection([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{1}), []interface{}{1})
	xtesting.Equal(t, Intersection([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{4, 5, 2}), []interface{}{4, 5, 2, 5, 4, 2})
	xtesting.Equal(t, Intersection([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{4, 5, 2, 5, 1, 4, 2}), []interface{}{4, 5, 2, 5, 1, 4, 2})
	xtesting.Equal(t, Intersection([]interface{}{4, 5, 2, 5, 1, 4, 2}, []interface{}{0}), []interface{}{})

	eq := func(i, j interface{}) bool { return i.(string) == strconv.Itoa(j.(int)) }
	xtesting.Equal(t, IntersectionWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{1}, eq), []interface{}{"1"})
	xtesting.Equal(t, IntersectionWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{4, 5, 2}, eq), []interface{}{"4", "5", "2", "5", "4", "2"})
	xtesting.Equal(t, IntersectionWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{4, 5, 2, 5, 1, 4, 2}, eq), []interface{}{"4", "5", "2", "5", "1", "4", "2"})
	xtesting.Equal(t, IntersectionWith([]interface{}{"4", "5", "2", "5", "1", "4", "2"}, []interface{}{0}, eq), []interface{}{})
}

func TestToSet(t *testing.T) {
	xtesting.Equal(t, ToSet([]interface{}{1, 2, 3, 1, 2, 3, 4, 5, 6}), []interface{}{1, 2, 3, 4, 5, 6})
	xtesting.Equal(t, ToSet(nil), []interface{}{})
	xtesting.Equal(t, ToSet([]interface{}{}), []interface{}{})
	xtesting.Equal(t, ToSet([]interface{}{1}), []interface{}{1})
	xtesting.Equal(t, ToSet([]interface{}{1, 0, 1, 0, 1, 0, 1}), []interface{}{1, 0})
	xtesting.Equal(t, ToSet([]interface{}{1, 1, 1, 1, 1}), []interface{}{1})

	eq := func(i, j interface{}) bool { return math.Abs(float64(i.(int))-float64(j.(int))) <= 1 }
	xtesting.Equal(t, ToSetWith([]interface{}{1, 2, 3, 1, 2, 3, 4, 5, 6}, eq), []interface{}{1, 3, 5})
	xtesting.Equal(t, ToSetWith(nil, eq), []interface{}{})
	xtesting.Equal(t, ToSetWith([]interface{}{}, eq), []interface{}{})
	xtesting.Equal(t, ToSetWith([]interface{}{1}, eq), []interface{}{1})
	xtesting.Equal(t, ToSetWith([]interface{}{1, 0, 1, 0, 1, 0, 1}, eq), []interface{}{1})
	xtesting.Equal(t, ToSetWith([]interface{}{1, 1, 1, 1, 1}, eq), []interface{}{1})
}

func TestEqual(t *testing.T) {
	s1 := []interface{}{1, 2, 5, 6, 7}
	s2 := []interface{}{2, 1, 7, 5, 6}
	s3 := []interface{}{1, 5, 6, 7}
	s4 := []interface{}{2, 7, 6, 8}
	s5 := []interface{}{1, 2, 2, 5, 6, 7}
	s6 := []interface{}{1, 2, 3, 5, 6, 7}
	xtesting.Equal(t, Equal(s1, s2), true)
	xtesting.Equal(t, Equal(s1, s3), false)
	xtesting.Equal(t, Equal(s1, s4), false)
	xtesting.Equal(t, Equal(s1, s5), false)
	xtesting.Equal(t, Equal(s3, s4), false)
	xtesting.Equal(t, Equal(s5, s6), false)

	eq := func(i, j interface{}) bool { return math.Abs(float64(i.(int))-float64(j.(int))) <= 1 }
	xtesting.Equal(t, EqualWith(s1, s2, eq), true)
	xtesting.Equal(t, EqualWith(s1, s3, eq), false)
	xtesting.Equal(t, EqualWith(s1, s4, eq), false)
	xtesting.Equal(t, EqualWith(s1, s5, eq), false)
	xtesting.Equal(t, EqualWith(s3, s4, eq), true)
	xtesting.Equal(t, EqualWith(s5, s6, eq), true)
}

func TestRange(t *testing.T) {
	xtesting.Panic(t, func() { Range(3, 1, 1) })
	xtesting.Panic(t, func() { Range(1, 3, 0) })
	xtesting.Equal(t, Range(0, 10, 1), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	xtesting.Equal(t, Range(1, 10, 2), []int{1, 3, 5, 7, 9})
	xtesting.Equal(t, Range(0, 10, 2), []int{0, 2, 4, 6, 8, 10})
	xtesting.Equal(t, Range(0, 1, 2), []int{0})
}

func TestReverseRange(t *testing.T) {
	xtesting.Panic(t, func() { ReverseRange(3, 1, 1) })
	xtesting.Panic(t, func() { ReverseRange(1, 3, 0) })
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
