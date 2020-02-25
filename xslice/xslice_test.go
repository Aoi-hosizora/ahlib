package xslice

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
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
	assert.Equal(t, Its(i, ""), interface{}(s))
	assert.Equal(t, Its(nil, 0), nil)
	assert.Equal(t, Its(nil, nil), nil)
}

func TestIndexOf(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	assert.Equal(t, IndexOf(Sti(s), 1), 0)
	assert.Equal(t, IndexOf(Sti(s), 6), -1)
	assert.Equal(t, IndexOf(Sti(s), nil), -1)
}

func TestContains(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3}
	assert.Equal(t, Contains(Sti(s), 1), true)
	assert.Equal(t, Contains(Sti(s), 6), false)
	assert.Equal(t, Contains(Sti(s), nil), false)
}

func TestDelete(t *testing.T) {
	s := []int{1, 5, 2, 1, 2, 3, 1}

	s = Its(Delete(Sti(s), 1, 1), 0).([]int)
	assert.Equal(t, s, []int{5, 2, 1, 2, 3, 1})
	s = Its(Delete(Sti(s), 1, 2), 0).([]int)
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = Its(Delete(Sti(s), 6, 1), 0).([]int)
	assert.Equal(t, s, []int{5, 2, 2, 3})
	s = Its(Delete(Sti(s), 2, -1), 0).([]int)
	assert.Equal(t, s, []int{5, 3})
	s = Its(Delete(Sti(s), nil, -1), 0).([]int)
	assert.Equal(t, s, []int{5, 3})

	ss := Its(Delete(nil, 2, -1), 0)
	assert.Equal(t, ss == nil, true)
}

func TestDeleteAll(t *testing.T) {
	assert.Equal(t, Its(DeleteAll(Sti([]int{1, 5, 2, 1, 2, 3, 1}), 1), 0).([]int), []int{5, 2, 2, 3})
}

func TestSliceDiff(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	assert.Equal(t, Its(SliceDiff(Sti(slice1), Sti(slice2)), 0).([]int), []int{2, 3, 3})
}
