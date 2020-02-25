package xslice

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

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

func TestShuffle(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())

	emptySlice := make([]interface{}, 0)
	Shuffle(emptySlice, source)
	assert.Equal(t, emptySlice, []interface{}{})

	oneElementSlice := []interface{}{11}
	Shuffle(oneElementSlice, source)
	assert.Equal(t, oneElementSlice, []interface{}{11})

	slice := []interface{}{"a", "b", "c"}
	Shuffle(slice, source)
	assert.Contains(t, slice, "a")
	assert.Contains(t, slice, "b")
	assert.Contains(t, slice, "c")
}

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(Sti([]int{1, 2, 3, 4, 5, 6})), Sti([]int{6, 5, 4, 3, 2, 1}))
	assert.Equal(t, Reverse(Sti([]int{})), Sti([]int{}))
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

	ss := Its(Delete(nil, 2, -1), 0)
	assert.Equal(t, ss == nil, true)
}

func TestDeleteAll(t *testing.T) {
	assert.Equal(t, DeleteAll(Sti([]int{1, 5, 2, 1, 2, 3, 1}), 1), Sti([]int{5, 2, 2, 3}))
}

func TestSliceDiff(t *testing.T) {
	slice1 := []int{1, 2, 1, 3, 4, 3}
	slice2 := []int{1, 5, 6, 4}
	assert.Equal(t, SliceDiff(Sti(slice1), Sti(slice2)), Sti([]int{2, 3, 3}))
}
