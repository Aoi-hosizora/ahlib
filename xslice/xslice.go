package xslice

import (
	"math/rand"
	"time"
)

// Equaller represents how two data equal, used for XXXWith methods.
type Equaller func(i, j interface{}) bool

// defaultEqualler represents a default Equaller, it just checks equality directly.
var defaultEqualler Equaller = func(i, j interface{}) bool {
	return i == j
}

// ShuffleSelf shuffles the []interface{} slice directly.
func ShuffleSelf(slice []interface{}) {
	coreShuffle(checkSliceParam(slice))
}

// Shuffle shuffles the []interface{} slice and returns the result.
func Shuffle(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreShuffle(checkSliceParam(newSlice))
	return newSlice
}

// ShuffleSelfG shuffles the []T slice directly, is the generic function of ShuffleSelf.
func ShuffleSelfG(slice interface{}) {
	coreShuffle(checkInterfaceParam(slice))
}

// ShuffleG shuffles the []T slice and returns the result, is the generic function of Shuffle.
func ShuffleG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreShuffle(checkInterfaceParam(newSlice))
	return newSlice
}

// coreShuffle is the implementation for ShuffleSelf and Shuffle.
func coreShuffle(slice innerSlice) {
	rand.Seed(time.Now().UnixNano())
	for i := slice.length() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		itemJ := slice.get(j)
		itemI := slice.get(i)
		slice.set(i, itemJ)
		slice.set(j, itemI)
	}
}

// ReverseSelf reverses the []interface{} slice directly.
func ReverseSelf(slice []interface{}) {
	coreReverse(checkSliceParam(slice))
}

// Reverse reverses the []interface{} slice and returns the result.
func Reverse(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreReverse(checkSliceParam(newSlice))
	return newSlice
}

// ReverseSelfG reverses the []T slice directly, is the generic function of ReverseSelf.
func ReverseSelfG(slice interface{}) {
	coreReverse(checkInterfaceParam(slice))
}

// ReverseG reverses the []T slice and returns the result, is the generic function of Reverse.
func ReverseG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreReverse(checkInterfaceParam(newSlice))
	return newSlice
}

// coreReverse is the implementation for ReverseSelf and Reverse.
func coreReverse(slice innerSlice) {
	for i, j := 0, slice.length()-1; i < j; i, j = i+1, j-1 {
		itemJ := slice.get(j)
		itemI := slice.get(i)
		slice.set(i, itemJ)
		slice.set(j, itemI)
	}
}

// IndexOf returns the first index of value in the []interface{} slice.
func IndexOf(slice []interface{}, value interface{}) int {
	return coreIndexOf(checkSliceParam(slice), value, defaultEqualler)
}

// IndexOfWith returns the first index of value in the []interface{} slice with Equaller.
func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreIndexOf(checkSliceParam(slice), value, equaller)
}

// IndexOfG returns the first index of value in the []T slice, is the generic function of IndexOf.
func IndexOfG(slice interface{}, value interface{}) int {
	return coreIndexOf(checkInterfaceParam(slice), value, defaultEqualler)
}

// IndexOfWithG returns the first index of value in the []T slice with Equaller, is the generic function of IndexOfWith.
func IndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int {
	return coreIndexOf(checkInterfaceParam(slice), value, equaller)
}

// coreIndexOf is the implementation for IndexOf and IndexOfWith.
func coreIndexOf(slice innerSlice, value interface{}, equaller Equaller) int {
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(val, value) {
			return idx
		}
	}
	return -1
}

// Contains returns true if the value is in the []interface{} slice.
func Contains(slice []interface{}, value interface{}) bool {
	return coreContains(checkSliceParam(slice), value, defaultEqualler)
}

// ContainsWith returns true if the value is in the []interface{} slice with Equaller.
func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return coreContains(checkSliceParam(slice), value, equaller)
}

// Contains returns true if the value is in the []T slice, is the generic function of Contains.
func ContainsG(slice interface{}, value interface{}) bool {
	return coreContains(checkInterfaceParam(slice), value, defaultEqualler)
}

// ContainsWith returns true if the value is in the []T slice with Equaller, is the generic function of ContainsWith.
func ContainsWithG(slice interface{}, value interface{}, equaller Equaller) bool {
	return coreContains(checkInterfaceParam(slice), value, equaller)
}

// coreContains is the implementation for Contains and ContainsWith.
func coreContains(slice innerSlice, value interface{}, equaller Equaller) bool {
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(val, value) {
			return true
		}
	}
	return false
}

// Count returns the count of value in the []interface{} slice.
func Count(slice []interface{}, value interface{}) int {
	return coreCount(checkSliceParam(slice), value, defaultEqualler)
}

// CountWith returns the count of value in the []interface{} slice with Equaller.
func CountWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreCount(checkSliceParam(slice), value, equaller)
}

// CountG returns the count of value in the []T slice, is the generic function of Count.
func CountG(slice interface{}, value interface{}) int {
	return coreCount(checkInterfaceParam(slice), value, defaultEqualler)
}

// CountWithG returns the count of value in the []T slice with Equaller, is the generic function of CountWith.
func CountWithG(slice interface{}, value interface{}, equaller Equaller) int {
	return coreCount(checkInterfaceParam(slice), value, equaller)
}

// coreCount is the implementation for Count and CountWith.
func coreCount(slice innerSlice, value interface{}, equaller Equaller) int {
	cnt := 0
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(val, value) {
			cnt++
		}
	}
	return cnt
}

// Delete deletes a value from []interface{} slice in n times.
func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkSliceParam(newSlice), value, n, defaultEqualler).actual().([]interface{})
}

// DeleteWith deletes a value from []interface{} slice in n times with Equaller.
func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkSliceParam(newSlice), value, n, equaller).actual().([]interface{})
}

// DeleteG deletes a value from []T slice in n times, is the generic function of Delete.
func DeleteG(slice interface{}, value interface{}, n int) interface{} {
	newSlice := cloneSliceInterface(slice)
	return coreDelete(checkInterfaceParam(newSlice), value, n, defaultEqualler).actual()
}

// DeleteWithG deletes a value from []T slice in n times with Equaller, is the generic function of DeleteWith.
func DeleteWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{} {
	newSlice := cloneSliceInterface(slice)
	return coreDelete(checkInterfaceParam(newSlice), value, n, equaller).actual()
}

// DeleteAll deletes a value from []interface{} slice in all.
func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkSliceParam(newSlice), value, 0, defaultEqualler).actual().([]interface{})
}

// DeleteAllWith deletes a value from []interface{} slice in all with Equaller.
func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkSliceParam(newSlice), value, 0, equaller).actual().([]interface{})
}

// DeleteAllG deletes a value from []T slice in all, is the generic function of DeleteAll.
func DeleteAllG(slice interface{}, value interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	return coreDelete(checkInterfaceParam(newSlice), value, 0, defaultEqualler).actual()
}

// DeleteAllWithG deletes a value from []T slice in all with Equaller, is the generic function of DeleteAllWith.
func DeleteAllWithG(slice interface{}, value interface{}, equaller Equaller) interface{} {
	newSlice := cloneSliceInterface(slice)
	return coreDelete(checkInterfaceParam(newSlice), value, 0, equaller).actual()
}

// coreDelete is the implementation for Delete, DeleteWith, DeleteAll and DeleteAllWith.
func coreDelete(slice innerSlice, value interface{}, n int, equaller Equaller) innerSlice {
	if n <= 0 {
		n = slice.length()
	}
	cnt := 0
	idx := coreIndexOf(slice, value, equaller)
	for idx != -1 && cnt < n {
		slice.remove(idx)
		cnt++
		idx = coreIndexOf(slice, value, equaller)
	}
	return slice
}

// Diff returns the difference of two []interface{} slices.
func Diff(slice1, slice2 []interface{}) []interface{} {
	return coreDiff(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// DiffWith returns the difference of two []interface{} slices with Equaller.
func DiffWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreDiff(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

// DiffG returns the difference of two []T slices, is the generic function of Diff.
func DiffG(slice1, slice2 interface{}) interface{} {
	return coreDiff(checkInterfaceParam(slice1), checkInterfaceParam(slice2), defaultEqualler).actual()
}

// DiffWithG returns the difference of two []T slices with Equaller, is the generic function of DiffWith.
func DiffWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	return coreDiff(checkInterfaceParam(slice1), checkInterfaceParam(slice2), equaller).actual()
}

// coreDiff is the implementation for Diff and DiffWith.
func coreDiff(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, 0)
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		exist := false
		for i2 := 0; i2 < slice2.length(); i2++ {
			item2 := slice2.get(i2)
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result.append(item1)
		}
	}
	return result
}

// Union returns the union of two []interface{} slices.
func Union(slice1, slice2 []interface{}) []interface{} {
	return coreUnion(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// UnionWith returns the union of two []interface{} slices with Equaller.
func UnionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreUnion(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

// UnionG returns the union of two []interface{} slices, is the generic function of Union.
func UnionG(slice1, slice2 interface{}) interface{} {
	return coreUnion(checkInterfaceParam(slice1), checkInterfaceParam(slice2), defaultEqualler).actual()
}

// UnionWithG returns the union of two []interface{} slices with Equaller, is the generic function of UnionWith.
func UnionWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	return coreUnion(checkInterfaceParam(slice1), checkInterfaceParam(slice2), equaller).actual()
}

// coreUnion is the implementation for Union and UnionWith.
func coreUnion(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, slice1.length())
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		result.append(item1)
	}
	for i2 := 0; i2 < slice2.length(); i2++ {
		item2 := slice2.get(i2)
		exist := false
		for i1 := 0; i1 < slice1.length(); i1++ {
			item1 := slice1.get(i1)
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result.append(item2)
		}
	}
	return result
}

// Intersection returns the intersection of two []interface{} slices.
func Intersection(slice1, slice2 []interface{}) []interface{} {
	return coreIntersection(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// IntersectionWith returns the intersection of two []interface{} slices with Equaller.
func IntersectionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreIntersection(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

// IntersectionG returns the intersection of two []T slices, is the generic function of Intersection.
func IntersectionG(slice1, slice2 interface{}) interface{} {
	return coreIntersection(checkInterfaceParam(slice1), checkInterfaceParam(slice2), defaultEqualler).actual()
}

// IntersectionWithG returns the intersection of two []T slices with Equaller, is the generic function of IntersectionWith.
func IntersectionWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	return coreIntersection(checkInterfaceParam(slice1), checkInterfaceParam(slice2), equaller).actual()
}

// coreIntersection is the implementation for Intersection and IntersectionWith.
func coreIntersection(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, 0)
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		for i2 := 0; i2 < slice2.length(); i2++ {
			item2 := slice2.get(i2)
			if equaller(item1, item2) {
				result.append(item1)
				break
			}
		}
	}
	return result
}

// ToSet removes the duplicate items in []interface{} slice as a set.
func ToSet(slice []interface{}) []interface{} {
	return coreToSet(checkSliceParam(slice), defaultEqualler).actual().([]interface{})
}

// ToSetWith removes the duplicate items in []interface{} slice as a set with Equaller.
func ToSetWith(slice []interface{}, equaller Equaller) []interface{} {
	return coreToSet(checkSliceParam(slice), equaller).actual().([]interface{})
}

// ToSet removes the duplicate items in []interface{} slice as a set, is the generic function of ToSet.
func ToSetG(slice interface{}) interface{} {
	return coreToSet(checkInterfaceParam(slice), defaultEqualler).actual()
}

// ToSetWith removes the duplicate items in []interface{} slice as a set with Equaller, is the generic function of ToSetWith.
func ToSetWithG(slice interface{}, equaller Equaller) interface{} {
	return coreToSet(checkInterfaceParam(slice), equaller).actual()
}

// coreToSet is the implementation for ToSet and ToSetWith.
func coreToSet(slice innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice, 0, 0)
	for idx := 0; idx < slice.length(); idx++ {
		item := slice.get(idx)
		if coreCount(result, item, equaller) == 0 {
			result.append(item)
		}
	}
	return result
}

// ElementMatch checks two []interface{} slice equals without order.
func ElementMatch(slice1, slice2 []interface{}) bool {
	return coreElementMatch(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler)
}

// ElementMatchWith checks two []interface{} slice equals without order with Equaller.
func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool {
	return coreElementMatch(checkSliceParam(slice1), checkSliceParam(slice2), equaller)
}

// ElementMatch checks two []interface{} slice equals without order, is the generic function of ElementMatch.
func ElementMatchG(slice1, slice2 interface{}) bool {
	return coreElementMatch(checkInterfaceParam(slice1), checkInterfaceParam(slice2), defaultEqualler)
}

// ElementMatchWith checks two []interface{} slice equals without order with Equaller, is the generic function of ElementMatchWith.
func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool {
	return coreElementMatch(checkInterfaceParam(slice1), checkInterfaceParam(slice2), equaller)
}

// coreElementMatch is the implementation for ElementMatch and ElementMatchWith.
func coreElementMatch(slice1, slice2 innerSlice, equaller Equaller) bool {
	extra1 := makeInnerSlice(slice1, 0, 0)
	extra2 := makeInnerSlice(slice2, 0, 0)

	visited := make([]bool, slice2.length())
	for idx1 := 0; idx1 < slice1.length(); idx1++ {
		item1 := slice1.get(idx1)
		exist := false
		for idx2 := 0; idx2 < slice2.length(); idx2++ {
			if visited[idx2] {
				continue
			}
			item2 := slice2.get(idx2)
			if equaller(item1, item2) {
				visited[idx2] = true
				exist = true
				break
			}
		}
		if !exist {
			extra1.append(item1)
		}
	}

	for item2 := 0; item2 < slice2.length(); item2++ {
		if !visited[item2] {
			extra2.append(item2)
		}
	}

	return extra1.length() == 0 && extra2.length() == 0
}

const (
	minLargerThenMaxPanic = "xslice: min is larger than max"
	stepLessThenZeroPanic = "xslice: step is less then or equals to 0"
)

// Range generates a []int slice from min index to max index with step.
func Range(min, max, step int) []int {
	if min > max {
		panic(minLargerThenMaxPanic)
	} else if step <= 0 {
		panic(stepLessThenZeroPanic)
	}

	out := make([]int, 0)
	for idx := min; idx <= max; idx += step {
		out = append(out, idx)
	}
	return out
}

// ReverseRange generates a reverse []int slice from max index to min index with step.
func ReverseRange(min, max, step int) []int {
	if min > max {
		panic(minLargerThenMaxPanic)
	} else if step <= 0 {
		panic(stepLessThenZeroPanic)
	}

	out := make([]int, 0)
	for idx := max; idx >= min; idx -= step {
		out = append(out, idx)
	}
	return out
}
