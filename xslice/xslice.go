package xslice

import (
	"math/rand"
	"sort"
	"time"
)

// Equaller represents an equality function for two interface{}, is used in XXXWith methods.
type Equaller func(i, j interface{}) bool

// Lesser represents a less function for sort, see sort.Interface.
type Lesser func(i, j interface{}) bool

// defaultEqualler represents a default Equaller, it just checks equality by `==`.
var defaultEqualler Equaller = func(i, j interface{}) bool {
	return i == j
}

// ShuffleSelf shuffles the []interface{} slice directly.
func ShuffleSelf(slice []interface{}) {
	coreShuffle(checkInterfaceSliceParam(slice))
}

// Shuffle shuffles the []interface{} slice and returns the result.
func Shuffle(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreShuffle(checkInterfaceSliceParam(newSlice))
	return newSlice
}

// ShuffleSelfG shuffles the []T slice directly, is the generic function of ShuffleSelf.
func ShuffleSelfG(slice interface{}) {
	coreShuffle(checkSliceInterfaceParam(slice))
}

// ShuffleG shuffles the []T slice and returns the result, is the generic function of Shuffle.
func ShuffleG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreShuffle(checkSliceInterfaceParam(newSlice))
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
	coreReverse(checkInterfaceSliceParam(slice))
}

// Reverse reverses the []interface{} slice and returns the result.
func Reverse(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreReverse(checkInterfaceSliceParam(newSlice))
	return newSlice
}

// ReverseSelfG reverses the []T slice directly, is the generic function of ReverseSelf.
func ReverseSelfG(slice interface{}) {
	coreReverse(checkSliceInterfaceParam(slice))
}

// ReverseG reverses the []T slice and returns the result, is the generic function of Reverse.
func ReverseG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreReverse(checkSliceInterfaceParam(newSlice))
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

// SortSelf sorts the []interface{} slice with less function directly.
func SortSelf(slice []interface{}, less Lesser) {
	coreSort(checkInterfaceSliceParam(slice), less, false)
}

// Sort sorts the []interface{} slice with less function and returns the result.
func Sort(slice []interface{}, less Lesser) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreSort(checkInterfaceSliceParam(newSlice), less, false)
	return newSlice
}

// SortSelfG sorts the []T slice with less function directly, is the generic function of SortSelf.
func SortSelfG(slice interface{}, less Lesser) {
	coreSort(checkSliceInterfaceParam(slice), less, false)
}

// SortG sorts the []T slice with less function and returns the result, is the generic function of Sort.
func SortG(slice interface{}, less Lesser) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreSort(checkSliceInterfaceParam(newSlice), less, false)
	return newSlice
}

// StableSortSelf sorts the []interface{} slice in stable with less function directly.
func StableSortSelf(slice []interface{}, less Lesser) {
	coreSort(checkInterfaceSliceParam(slice), less, true)
}

// StableSort sorts the []interface{} slice in stable with less function and returns the result.
func StableSort(slice []interface{}, less Lesser) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreSort(checkInterfaceSliceParam(newSlice), less, true)
	return newSlice
}

// StableSortSelfG sorts the []T slice in stable with less function directly, is the generic function of StableSortSelf.
func StableSortSelfG(slice interface{}, less Lesser) {
	coreSort(checkSliceInterfaceParam(slice), less, true)
}

// StableSortG sorts the []T slice with in stable less function and returns the result, is the generic function of StableSort.
func StableSortG(slice interface{}, less Lesser) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreSort(checkSliceInterfaceParam(newSlice), less, true)
	return newSlice
}

// coreSort is the implementation for SortSelf, Sort, StableSortSelf and StableSort, using sort.Slice and sort.SliceStable.
func coreSort(slice innerSlice, less Lesser, stable bool) {
	if less == nil {
		panic(panicNilLesser)
	}
	ss := &sortSlice{slice: slice, less: less}
	if stable {
		sort.Stable(ss)
	} else {
		sort.Sort(ss)
	}
}

// IndexOf returns the first index of value in the []interface{} slice.
func IndexOf(slice []interface{}, value interface{}) int {
	return coreIndexOf(checkInterfaceSliceParam(slice), value, defaultEqualler)
}

// IndexOfWith returns the first index of value in the []interface{} slice with Equaller.
func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreIndexOf(checkInterfaceSliceParam(slice), value, equaller)
}

// IndexOfG returns the first index of value in the []T slice, is the generic function of IndexOf.
func IndexOfG(slice interface{}, value interface{}) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreIndexOf(s, v, defaultEqualler)
}

// IndexOfWithG returns the first index of value in the []T slice with Equaller, is the generic function of IndexOfWith.
func IndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreIndexOf(s, v, equaller)
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

// Contains returns true if value is in the []interface{} slice.
func Contains(slice []interface{}, value interface{}) bool {
	return coreContains(checkInterfaceSliceParam(slice), value, defaultEqualler)
}

// ContainsWith returns true if value is in the []interface{} slice with Equaller.
func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return coreContains(checkInterfaceSliceParam(slice), value, equaller)
}

// ContainsG returns true if value is in the []T slice, is the generic function of Contains.
func ContainsG(slice interface{}, value interface{}) bool {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreContains(s, v, defaultEqualler)
}

// ContainsWithG returns true if value is in the []T slice with Equaller, is the generic function of ContainsWith.
func ContainsWithG(slice interface{}, value interface{}, equaller Equaller) bool {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreContains(s, v, equaller)
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
	return coreCount(checkInterfaceSliceParam(slice), value, defaultEqualler)
}

// CountWith returns the count of value in the []interface{} slice with Equaller.
func CountWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreCount(checkInterfaceSliceParam(slice), value, equaller)
}

// CountG returns the count of value in the []T slice, is the generic function of Count.
func CountG(slice interface{}, value interface{}) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreCount(s, v, defaultEqualler)
}

// CountWithG returns the count of value in the []T slice with Equaller, is the generic function of CountWith.
func CountWithG(slice interface{}, value interface{}, equaller Equaller) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreCount(s, v, equaller)
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

// Delete deletes value from []interface{} slice in n times.
func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkInterfaceSliceParam(newSlice), value, n, defaultEqualler).actual().([]interface{})
}

// DeleteWith deletes value from []interface{} slice in n times with Equaller.
func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkInterfaceSliceParam(newSlice), value, n, equaller).actual().([]interface{})
}

// DeleteG deletes value from []T slice in n times, is the generic function of Delete.
func DeleteG(slice interface{}, value interface{}, n int) interface{} {
	newSlice := cloneSliceInterface(slice)
	s, v := checkSliceInterfaceAndElemParam(newSlice, value)
	return coreDelete(s, v, n, defaultEqualler).actual()
}

// DeleteWithG deletes value from []T slice in n times with Equaller, is the generic function of DeleteWith.
func DeleteWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{} {
	newSlice := cloneSliceInterface(slice)
	s, v := checkSliceInterfaceAndElemParam(newSlice, value)
	return coreDelete(s, v, n, equaller).actual()
}

// DeleteAll deletes value from []interface{} slice in all.
func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkInterfaceSliceParam(newSlice), value, 0, defaultEqualler).actual().([]interface{})
}

// DeleteAllWith deletes value from []interface{} slice in all with Equaller.
func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	return coreDelete(checkInterfaceSliceParam(newSlice), value, 0, equaller).actual().([]interface{})
}

// DeleteAllG deletes value from []T slice in all, is the generic function of DeleteAll.
func DeleteAllG(slice interface{}, value interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	s, v := checkSliceInterfaceAndElemParam(newSlice, value)
	return coreDelete(s, v, 0, defaultEqualler).actual()
}

// DeleteAllWithG deletes value from []T slice in all with Equaller, is the generic function of DeleteAllWith.
func DeleteAllWithG(slice interface{}, value interface{}, equaller Equaller) interface{} {
	newSlice := cloneSliceInterface(slice)
	s, v := checkSliceInterfaceAndElemParam(newSlice, value)
	return coreDelete(s, v, 0, equaller).actual()
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
	return coreDiff(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// DiffWith returns the difference of two []interface{} slices with Equaller.
func DiffWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreDiff(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller).actual().([]interface{})
}

// DiffG returns the difference of two []T slices, is the generic function of Diff.
func DiffG(slice1, slice2 interface{}) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreDiff(s1, s2, defaultEqualler).actual()
}

// DiffWithG returns the difference of two []T slices with Equaller, is the generic function of DiffWith.
func DiffWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreDiff(s1, s2, equaller).actual()
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
	return coreUnion(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// UnionWith returns the union of two []interface{} slices with Equaller.
func UnionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreUnion(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller).actual().([]interface{})
}

// UnionG returns the union of two []T slices, is the generic function of Union.
func UnionG(slice1, slice2 interface{}) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreUnion(s1, s2, defaultEqualler).actual()
}

// UnionWithG returns the union of two []T slices with Equaller, is the generic function of UnionWith.
func UnionWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreUnion(s1, s2, equaller).actual()
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
	return coreIntersection(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// IntersectionWith returns the intersection of two []interface{} slices with Equaller.
func IntersectionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreIntersection(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller).actual().([]interface{})
}

// IntersectionG returns the intersection of two []T slices, is the generic function of Intersection.
func IntersectionG(slice1, slice2 interface{}) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreIntersection(s1, s2, defaultEqualler).actual()
}

// IntersectionWithG returns the intersection of two []T slices with Equaller, is the generic function of IntersectionWith.
func IntersectionWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreIntersection(s1, s2, equaller).actual()
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

// ToSet removes the duplicate items from []interface{} slice as a set.
func ToSet(slice []interface{}) []interface{} {
	return coreToSet(checkInterfaceSliceParam(slice), defaultEqualler).actual().([]interface{})
}

// ToSetWith removes the duplicate items from []interface{} slice as a set with Equaller.
func ToSetWith(slice []interface{}, equaller Equaller) []interface{} {
	return coreToSet(checkInterfaceSliceParam(slice), equaller).actual().([]interface{})
}

// ToSetG removes the duplicate items from []T slice as a set, is the generic function of ToSet.
func ToSetG(slice interface{}) interface{} {
	return coreToSet(checkSliceInterfaceParam(slice), defaultEqualler).actual()
}

// ToSetWithG removes the duplicate items from []T slice as a set with Equaller, is the generic function of ToSetWith.
func ToSetWithG(slice interface{}, equaller Equaller) interface{} {
	return coreToSet(checkSliceInterfaceParam(slice), equaller).actual()
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

// ElementMatch checks if two []interface{} slice equal without order.
func ElementMatch(slice1, slice2 []interface{}) bool {
	return coreElementMatch(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler)
}

// ElementMatchWith checks if two []interface{} slice equal without order with Equaller.
func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool {
	return coreElementMatch(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller)
}

// ElementMatchG checks if two []T slice equal without order, is the generic function of ElementMatch.
func ElementMatchG(slice1, slice2 interface{}) bool {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreElementMatch(s1, s2, defaultEqualler)
}

// ElementMatchWithG checks if two []T slice equal without order with Equaller, is the generic function of ElementMatchWith.
func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreElementMatch(s1, s2, equaller)
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

	for idx2 := 0; idx2 < slice2.length(); idx2++ {
		item2 := slice2.get(idx2)
		if !visited[idx2] {
			extra2.append(item2)
		}
	}

	return extra1.length() == 0 && extra2.length() == 0
}

const (
	panicMinLargerThenMax = "xslice: min is larger than max"
	panicStepLessThenZero = "xslice: step is less then or equals to 0"
)

// Range generates a []int slice from min index to max index with step.
// Deprecated: use xnumber.IntRange, xnumber.Int32Range, etc. instead.
func Range(min, max, step int) []int {
	if min > max {
		panic(panicMinLargerThenMax)
	} else if step <= 0 {
		panic(panicStepLessThenZero)
	}

	out := make([]int, 0)
	for idx := min; idx <= max; idx += step {
		out = append(out, idx)
	}
	return out
}

// ReverseRange generates a reverse []int slice from max index to min index with step.
// Deprecated: use xnumber.IntRange, xnumber.Int32Range, etc. instead.
func ReverseRange(min, max, step int) []int {
	if min > max {
		panic(panicMinLargerThenMax)
	} else if step <= 0 {
		panic(panicStepLessThenZero)
	}

	out := make([]int, 0)
	for idx := max; idx >= min; idx -= step {
		out = append(out, idx)
	}
	return out
}
