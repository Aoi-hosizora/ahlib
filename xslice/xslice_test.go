package xslice

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())

	emptyArray := make([]interface{}, 0)
	Shuffle(emptyArray, source)
	assert.Equal(t, emptyArray, []interface{}{})

	oneElementArray := []interface{}{11}
	Shuffle(oneElementArray, source)
	assert.Equal(t, oneElementArray, []interface{}{11})

	array := []interface{}{"a", "b", "c"}
	Shuffle(array, source)
	assert.Contains(t, array, "a")
	assert.Contains(t, array, "b")
	assert.Contains(t, array, "c")
}

func TestSti(t *testing.T) {
	s := []string{"123", "456"}
	i := []interface{}{interface{}("123"), interface{}("456")}
	assert.Equal(t, Sti(s), i)
	assert.Equal(t, Sti(""), []interface{}(nil))
	assert.Equal(t, Sti([]string{}), []interface{}{})
	assert.Equal(t, Sti("") == nil, true)
	assert.Equal(t, Sti(nil) == nil, true)
}

func TestIts(t *testing.T) {
	i := []interface{}{interface{}("123"), interface{}("456")}
	s := []string{"123", "456"}
	assert.Equal(t, Its(i, reflect.TypeOf("")), interface{}(s))
	assert.Equal(t, Its(i, reflect.TypeOf(0)), nil)
	assert.Equal(t, Its(nil, reflect.TypeOf(0)), nil)
	assert.Equal(t, Its(nil, nil), nil)
	assert.Equal(t, Its([]interface{}{0, "1"}, reflect.TypeOf(0)), nil)
}

func TestIndexOfSlice(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	assert.Equal(t, IndexOfSlice(Sti(s), 1), 0)
	assert.Equal(t, IndexOfSlice(Sti(s), 6), -1)
	assert.Equal(t, IndexOfSlice(Sti(s), nil), -1)
}

func TestDeleteInSlice(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3, 1}

	s = Its(DeleteInSlice(Sti(s), 1, 1), reflect.TypeOf(0)).([]int)
	assert.Equal(t, s, []int{5, 2, 1, 2, 3, 1})
	s = Its(DeleteInSlice(Sti(s), 1, 2), reflect.TypeOf(0)).([]int)
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = Its(DeleteInSlice(Sti(s), 6, 1), reflect.TypeOf(0)).([]int)
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = Its(DeleteInSlice(Sti(s), 2, -1), reflect.TypeOf(0)).([]int)
	assert.Equal(t, s, []int{5, 3})
	s = Its(DeleteInSlice(Sti(s), nil, -1), reflect.TypeOf(0)).([]int)
	assert.Equal(t, s, []int{5, 3})

	ss := Its(DeleteInSlice(nil, 2, -1), reflect.TypeOf(0))
	assert.Equal(t, ss == nil, true)
}
