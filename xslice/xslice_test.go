package xslice

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSti(t *testing.T) {
	assert.Equal(t, Sti([]string{"123", "456"}), []interface{}{"123", "456"})
	// assert.Equal(t, Sti(""), []interface{}(nil))
	assert.Equal(t, Sti([]string{}), []interface{}{})
	// assert.Equal(t, Sti("") == nil, true)
	assert.Equal(t, Sti(nil) == nil, true)

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
	assert.Equal(t, Its([]interface{}{"123", "456"}, ""), []string{"123", "456"})
	assert.Equal(t, Its(nil, 0), nil)
	// assert.Equal(t, Its(nil, nil), nil)

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
	source := rand.NewSource(time.Now().UnixNano())
	a := []interface{}{1, 2, 3, 4}
	Shuffle(a, source)
	log.Println(a)
	Shuffle(a, source)
	log.Println(a)
}

func TestReverse(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	Reverse(a)
	assert.Equal(t, a, []interface{}{4, 3, 2, 1})
	Reverse(a)
	assert.Equal(t, a, []interface{}{1, 2, 3, 4})
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"1", "2", "3", "4", "5"}
	s3 := []int{2, 3, 4, 5, 6}
	s4 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	assert.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return strconv.Itoa(i.(int)) }), Sti(s2))
	assert.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return i.(int) + 1 }), Sti(s3))
	assert.Equal(t, Map(Sti(s1), func(i interface{}) interface{} { return float32(i.(int)) }), Sti(s4))
}

func TestIndexOf(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	assert.Equal(t, IndexOf(Sti(s), 1), 0)
	assert.Equal(t, IndexOf(Sti(s), 6), -1)
	assert.Equal(t, IndexOf(Sti(s), nil), -1)
	assert.Equal(t, IndexOfWith([]interface{}{1, 2, 3, 4, 5}, 3, func(i, j interface{}) bool {
		return i.(int) == j.(int)-1
	}), 1)
}

func TestContains(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	assert.Equal(t, Contains(Sti(s), 1), true)
	assert.Equal(t, Contains(Sti(s), 6), false)
	assert.Equal(t, Contains(Sti(s), nil), false)
	assert.Equal(t, ContainsWith(Sti(s), 4, func(i, j interface{}) bool {
		return i.(int) == j.(int)-1
	}), true)
}

func TestDelete(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3, 1}

	s = ItsOfInt(Delete(Sti(s), 1, 1))
	assert.Equal(t, s, []int{5, 2, 1, 2, 3, 1})
	s = ItsOfInt(Delete(Sti(s), 1, 2))
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = ItsOfInt(Delete(Sti(s), 6, 1))
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = ItsOfInt(Delete(Sti(s), 2, -1))
	assert.Equal(t, s, []int{5, 3})
	s = ItsOfInt(Delete(Sti(s), nil, -1))
	assert.Equal(t, s, []int{5, 3})

	ss := ItsOfInt(Delete(nil, 2, -1))
	assert.Equal(t, ss == nil, true)
}

func TestDeleteAll(t *testing.T) {
	assert.Equal(t, DeleteAll(Sti([]int{1, 5, 2, 1, 2, 3, 1}), 1), Sti([]int{5, 2, 2, 3}))
}

func TestSliceDiff(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	assert.Equal(t, Diff(Sti(slice1), Sti(slice2)), Sti([]int{2, 3, 3}))
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 5, 6, 7}
	s2 := []int{2, 1, 7, 5, 6}
	s3 := []int{1, 5, 6, 7}
	s4 := []int{1, 5, 8, 7}
	s5 := []int{1, 2, 2, 5, 6, 7}
	assert.Equal(t, Equal(Sti(s1), Sti(s2)), true)
	assert.Equal(t, Equal(Sti(s1), Sti(s3)), false)
	assert.Equal(t, Equal(Sti(s1), Sti(s4)), false)
	assert.Equal(t, Equal(Sti(s1), Sti(s5)), false)
}
