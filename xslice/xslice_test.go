package xslice

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"math"
	"strconv"
	"testing"
)

func TestShuffle(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}

	ShuffleSelf(a)
	xtesting.ElementMatch(t, a, []interface{}{1, 2, 3, 4})
	log.Println("ShuffleSelf:", a)

	ShuffleSelf(a)
	xtesting.ElementMatch(t, a, []interface{}{1, 2, 3, 4})
	log.Println("ShuffleSelf:", a)

	b := make([]interface{}, 0)
	ShuffleSelf(b)
	xtesting.Equal(t, b, b)
}

func TestShuffleNew(t *testing.T) {
	aa := []interface{}{1, 2, 3, 4}

	a := Shuffle(aa)
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})
	xtesting.ElementMatch(t, aa, a)
	log.Println("Shuffle:", a)

	a = Shuffle(aa)
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})
	xtesting.ElementMatch(t, aa, a)
	log.Println("Shuffle:", a)

	b := make([]interface{}, 0)
	xtesting.Equal(t, b, []interface{}{})
	xtesting.Equal(t, Shuffle(b), b)
}

func TestReverse(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}

	ReverseSelf(a)
	xtesting.Equal(t, a, []interface{}{4, 3, 2, 1})

	ReverseSelf(a)
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})

	b := make([]interface{}, 0)
	ReverseSelf(b)
	xtesting.Equal(t, b, b)
}

func TestReverseNew(t *testing.T) {
	aa := []interface{}{1, 2, 3, 4}

	a := Reverse(aa)
	xtesting.Equal(t, a, []interface{}{4, 3, 2, 1})
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})

	a = Reverse(a)
	xtesting.Equal(t, a, []interface{}{1, 2, 3, 4})
	xtesting.Equal(t, aa, []interface{}{1, 2, 3, 4})

	b := make([]interface{}, 0)
	xtesting.Equal(t, b, []interface{}{})
	xtesting.Equal(t, Reverse(b), b)
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

	s2 := []int{1, 5, 2, 1, 2, 3}
	xtesting.Equal(t, IndexOfG(s2, 0), -1)
	xtesting.Equal(t, IndexOfG(s2, 1), 0)
	xtesting.Equal(t, IndexOfG(s2, 2), 2)
	xtesting.Equal(t, IndexOfG(s2, 3), 5)
	xtesting.Equal(t, IndexOfG(s2, 4), -1)
	xtesting.Equal(t, IndexOfG(s2, 5), 1)
	xtesting.Equal(t, IndexOfG(s2, 6), -1)
	xtesting.Equal(t, IndexOfG(s2, nil), -1)

	xtesting.Equal(t, IndexOfWithG(s2, -1, eq), -1)
	xtesting.Equal(t, IndexOfWithG(s2, 0, eq), 0)
	xtesting.Equal(t, IndexOfWithG(s2, 1, eq), 2)
	xtesting.Equal(t, IndexOfWithG(s2, 2, eq), 5)
	xtesting.Equal(t, IndexOfWithG(s2, 3, eq), -1)
	xtesting.Equal(t, IndexOfWithG(s2, 4, eq), 1)
	xtesting.Equal(t, IndexOfWithG(s2, 5, eq), -1)
	xtesting.Equal(t, IndexOfWithG(s2, 6, eq), -1)
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

	s2 := []int{1, 5, 2, 1, 2, 3, 1}
	s2 = DeleteG(s2, 0, 1).([]int)
	xtesting.Equal(t, s2, []int{1, 5, 2, 1, 2, 3, 1})
	s2 = DeleteG(s2, 1, 1).([]int)
	xtesting.Equal(t, s2, []int{5, 2, 1, 2, 3, 1})
	s2 = DeleteG(s2, 1, -1).([]int)
	xtesting.Equal(t, s2, []int{5, 2, 2, 3})
	s2 = DeleteG(s2, 2, 1).([]int)
	xtesting.Equal(t, s2, []int{5, 2, 3})
	s2 = DeleteG(s2, 2, 1).([]int)
	xtesting.Equal(t, s2, []int{5, 3})
	s2 = DeleteG(s2, 3, 1).([]int)
	xtesting.Equal(t, s2, []int{5})
	s2 = DeleteG(s2, 4, 1).([]int)
	xtesting.Equal(t, s2, []int{5})
	s2 = DeleteG(s2, 5, 1).([]int)
	xtesting.Equal(t, s2, []int{})
	s2 = DeleteG(s2, nil, 1).([]int)
	xtesting.Equal(t, s2, []int{})
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

	xtesting.Equal(t, DiffG([]int{4, 5, 2, 5, 1, 4, 2}, []int{1}), []int{4, 5, 2, 5, 4, 2})
	xtesting.Equal(t, DiffG([]int{4, 5, 2, 5, 1, 4, 2}, []int{4, 5, 2}), []int{1})
	xtesting.Equal(t, DiffG([]int{4, 5, 2, 5, 1, 4, 2}, []int{4, 5, 2, 5, 1, 4, 2}), []int{})
	xtesting.Equal(t, DiffG([]int{4, 5, 2, 5, 1, 4, 2}, []int{0}), []int{4, 5, 2, 5, 1, 4, 2})

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
	xtesting.Equal(t, ElementMatch(s1, s2), true)
	xtesting.Equal(t, ElementMatch(s1, s3), false)
	xtesting.Equal(t, ElementMatch(s1, s4), false)
	xtesting.Equal(t, ElementMatch(s1, s5), false)
	xtesting.Equal(t, ElementMatch(s3, s4), false)
	xtesting.Equal(t, ElementMatch(s5, s6), false)

	eq := func(i, j interface{}) bool { return math.Abs(float64(i.(int))-float64(j.(int))) <= 1 }
	xtesting.Equal(t, ElementMatchWith(s1, s2, eq), true)
	xtesting.Equal(t, ElementMatchWith(s1, s3, eq), false)
	xtesting.Equal(t, ElementMatchWith(s1, s4, eq), false)
	xtesting.Equal(t, ElementMatchWith(s1, s5, eq), false)
	xtesting.Equal(t, ElementMatchWith(s3, s4, eq), true)
	xtesting.Equal(t, ElementMatchWith(s5, s6, eq), true)
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
