package xset

import (
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()

	set.Add(1, 2, 1, 2, 3, 4, 1)
	assert.Equal(t, xslice.Its(set.Slice(), 0).([]int), []int{1, 2, 3, 4})

	set.Remove(1, 4)
	assert.Equal(t, xslice.Its(set.Slice(), 0).([]int), []int{2, 3})

	set.Clear()
	assert.Equal(t, set.Size(), 0)

	set = FromSlice(xslice.Sti([]int{1, 5, 2, 1, 3, 5, 4}))
	assert.Equal(t, xslice.Its(set.Slice(), 0).([]int), []int{1, 5, 2, 3, 4})

	slice1 := FromSlice(xslice.Sti([]int{1, 2, 3, 4}))
	slice2 := FromSlice(xslice.Sti([]int{1, 5, 6, 4}))
	assert.Equal(t, slice1.Union(slice2), FromSlice(xslice.Sti([]int{1, 2, 3, 4, 5, 6})))
	assert.Equal(t, slice1.Intersect(slice2), FromSlice(xslice.Sti([]int{1, 4})))
	assert.Equal(t, slice1.Diff(slice2), FromSlice(xslice.Sti([]int{2, 3})))
}
