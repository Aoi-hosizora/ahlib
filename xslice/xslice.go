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

func init() {
	// for coreShuffle
	rand.Seed(time.Now().UnixNano())
}

// coreShuffle is the implementation for ShuffleSelf, Shuffle, ShuffleSelfG and ShuffleG.
func coreShuffle(slice innerSlice) {
	for i := slice.length() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		itemI, itemJ := slice.get(i), slice.get(j)
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

// coreReverse is the implementation for ReverseSelf Reverse, ReverseSelfG and ReverseG.
func coreReverse(slice innerSlice) {
	for i, j := 0, slice.length()-1; i < j; i, j = i+1, j-1 {
		itemI, itemJ := slice.get(i), slice.get(j)
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

// coreSort is the implementation for SortSelf, Sort, StableSortSelf, StableSort, SortSelfG, SortG, StableSortSelfG and StableSortG, using sort.Slice and sort.SliceStable.
func coreSort(slice innerSlice, less Lesser, stable bool) {
	ss := &sortableSlice{slice: slice, less: less}
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

// coreIndexOf is the implementation for IndexOf, IndexOfWith, IndexOfG and IndexOfWithG.
func coreIndexOf(slice innerSlice, value interface{}, equaller Equaller) int {
	length := slice.length()
	for idx := 0; idx < length; idx++ {
		val := slice.get(idx)
		if equaller(val, value) {
			return idx
		}
	}
	return -1
}

// LastIndexOf returns the last index of value in the []interface{} slice.
func LastIndexOf(slice []interface{}, value interface{}) int {
	return coreLastIndexOf(checkInterfaceSliceParam(slice), value, defaultEqualler)
}

// LastIndexOfWith returns the last index of value in the []interface{} slice with Equaller.
func LastIndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreLastIndexOf(checkInterfaceSliceParam(slice), value, equaller)
}

// LastIndexOfG returns the last index of value in the []T slice, is the generic function of IndexOf.
func LastIndexOfG(slice interface{}, value interface{}) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreLastIndexOf(s, v, defaultEqualler)
}

// LastIndexOfWithG returns the last index of value in the []T slice with Equaller, is the generic function of IndexOfWith.
func LastIndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int {
	s, v := checkSliceInterfaceAndElemParam(slice, value)
	return coreLastIndexOf(s, v, equaller)
}

// coreLastIndexOf is the implementation for LastIndexOf, LastIndexOfWith, LastIndexOfG and LastIndexOfWithG.
func coreLastIndexOf(slice innerSlice, value interface{}, equaller Equaller) int {
	for idx := slice.length() - 1; idx >= 0; idx-- {
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

// coreContains is the implementation for Contains, ContainsWith, ContainsG and ContainsWithG.
func coreContains(slice innerSlice, value interface{}, equaller Equaller) bool {
	length := slice.length()
	for idx := 0; idx < length; idx++ {
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

// coreCount is the implementation for Count, CountWith, CountG and CountWithG.
func coreCount(slice innerSlice, value interface{}, equaller Equaller) int {
	cnt := 0
	length := slice.length()
	for idx := 0; idx < length; idx++ {
		val := slice.get(idx)
		if equaller(val, value) {
			cnt++
		}
	}
	return cnt
}

// InsertSelf inserts values into []interface{} slice at index position using the space of given slice.
func InsertSelf(slice []interface{}, index int, values ...interface{}) []interface{} {
	return coreInsert(checkInterfaceSliceParam(slice), checkInterfaceSliceParam(values), index, true).actual().([]interface{})
}

// Insert inserts values into []interface{} slice at index position using a new slice space to store.
func Insert(slice []interface{}, index int, values ...interface{}) []interface{} {
	return coreInsert(checkInterfaceSliceParam(slice), checkInterfaceSliceParam(values), index, false).actual().([]interface{})
}

// InsertSelfG inserts values into []T slice at index position using the space of given slice, is the generic function of InsertSelf.
func InsertSelfG(slice interface{}, index int, values interface{}) interface{} {
	s, v := checkTwoSliceInterfaceParam(slice, values)
	return coreInsert(s, v, index, true).actual()
}

// InsertG inserts values into []T slice at index position using a new slice space to store, is the generic function of Insert.
func InsertG(slice interface{}, index int, values interface{}) interface{} {
	s, v := checkTwoSliceInterfaceParam(slice, values)
	return coreInsert(s, v, index, false).actual()
}

// coreInsert is the implementation for InsertSelf, Insert, InsertSelfG and InsertG.
func coreInsert(slice, values innerSlice, index int, self bool) innerSlice {
	if values.length() == 0 {
		if self {
			return slice
		}
		return cloneInnerSliceItems(slice, 0)
	}
	if self {
		slice.insert(index, values)
		return slice
	}
	newSlice := cloneInnerSliceItems(slice, values.length())
	newSlice.insert(index, values)
	return newSlice
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

// coreDelete is the implementation for Delete, DeleteWith, DeleteAll, DeleteAllWith, DeleteG, DeleteWithG, DeleteAllG and DeleteAllWithG.
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

// ContainsAll returns true if values in []interface{} subset are all in the []interface{} list.
func ContainsAll(list, subset []interface{}) bool {
	return coreContainsAll(checkInterfaceSliceParam(list), checkInterfaceSliceParam(subset), defaultEqualler)
}

// ContainsAllWith returns true if values in []interface{} subset are all in the []interface{} list.
func ContainsAllWith(list, subset []interface{}, equaller Equaller) bool {
	return coreContainsAll(checkInterfaceSliceParam(list), checkInterfaceSliceParam(subset), equaller)
}

// ContainsAllG returns true if values in []T subset are all in the []T list, is the generic function of ContainsAll.
func ContainsAllG(list, subset interface{}) bool {
	s1, s2 := checkTwoSliceInterfaceParam(list, subset)
	return coreContainsAll(s1, s2, defaultEqualler)
}

// ContainsAllWithG returns true if values in []T subset are all in the []T list, is the generic function of ContainsAllWith.
func ContainsAllWithG(list, subset interface{}, equaller Equaller) bool {
	s1, s2 := checkTwoSliceInterfaceParam(list, subset)
	return coreContainsAll(s1, s2, equaller)
}

// coreContainsAll is the implementation for ContainsAll, ContainsAllWith, ContainsAllG and ContainsAllWithG.
func coreContainsAll(list, subset innerSlice, equaller Equaller) bool {
	for i := 0; i < subset.length(); i++ {
		val := subset.get(i)
		if !coreContains(list, val, equaller) {
			return false
		}
	}
	return true
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

// coreDiff is the implementation for Diff, DiffWith, DiffG and DiffWithG.
func coreDiff(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeSameTypeInnerSlice(slice1, 0, 0)
	length1 := slice1.length()
	for i1 := 0; i1 < length1; i1++ {
		item1 := slice1.get(i1)
		if !coreContains(slice2, item1, equaller) {
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

// coreUnion is the implementation for Union, UnionWith, UnionG and UnionWithG.
func coreUnion(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeSameTypeInnerSlice(slice1, 0, slice1.length())
	length1 := slice1.length()
	for i1 := 0; i1 < length1; i1++ {
		item1 := slice1.get(i1)
		result.append(item1)
	}
	length2 := slice2.length()
	for i2 := 0; i2 < length2; i2++ {
		item2 := slice2.get(i2)
		if !coreContains(slice1, item2, equaller) {
			result.append(item2)
		}
	}
	return result
}

// Intersect returns the intersection of two []interface{} slices.
func Intersect(slice1, slice2 []interface{}) []interface{} {
	return coreIntersect(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

// IntersectWith returns the intersection of two []interface{} slices with Equaller.
func IntersectWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreIntersect(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller).actual().([]interface{})
}

// IntersectG returns the intersection of two []T slices, is the generic function of Intersect.
func IntersectG(slice1, slice2 interface{}) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreIntersect(s1, s2, defaultEqualler).actual()
}

// IntersectWithG returns the intersection of two []T slices with Equaller, is the generic function of IntersectWith.
func IntersectWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreIntersect(s1, s2, equaller).actual()
}

// coreIntersect is the implementation for Intersect, IntersectWith, IntersectG and IntersectWithG.
func coreIntersect(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeSameTypeInnerSlice(slice1, 0, 0)
	length1 := slice1.length()
	for i1 := 0; i1 < length1; i1++ {
		item1 := slice1.get(i1)
		if coreContains(slice2, item1, equaller) {
			result.append(item1)
		}
	}
	return result
}

// Deduplicate removes the duplicate items from []interface{} slice as a set.
func Deduplicate(slice []interface{}) []interface{} {
	return coreDeduplicate(checkInterfaceSliceParam(slice), defaultEqualler).actual().([]interface{})
}

// DeduplicateWith removes the duplicate items from []interface{} slice as a set with Equaller.
func DeduplicateWith(slice []interface{}, equaller Equaller) []interface{} {
	return coreDeduplicate(checkInterfaceSliceParam(slice), equaller).actual().([]interface{})
}

// DeduplicateG removes the duplicate items from []T slice as a set, is the generic function of Deduplicate.
func DeduplicateG(slice interface{}) interface{} {
	return coreDeduplicate(checkSliceInterfaceParam(slice), defaultEqualler).actual()
}

// DeduplicateWithG removes the duplicate items from []T slice as a set with Equaller, is the generic function of DeduplicateWith.
func DeduplicateWithG(slice interface{}, equaller Equaller) interface{} {
	return coreDeduplicate(checkSliceInterfaceParam(slice), equaller).actual()
}

// coreDeduplicate is the implementation for Deduplicate, DeduplicateWith, DeduplicateG and DeduplicateWithG.
func coreDeduplicate(slice innerSlice, equaller Equaller) innerSlice {
	result := makeSameTypeInnerSlice(slice, 0, 0)
	length := slice.length()
	for idx := 0; idx < length; idx++ {
		item := slice.get(idx)
		if !coreContains(result, item, equaller) { // O(n^2)
			result.append(item)
		}
	}
	return result
}

// ElementMatch checks whether two []interface{} slice equal without order.
func ElementMatch(slice1, slice2 []interface{}) bool {
	return coreElementMatch(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), defaultEqualler)
}

// ElementMatchWith checks whether two []interface{} slice equal without order with Equaller.
func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool {
	return coreElementMatch(checkInterfaceSliceParam(slice1), checkInterfaceSliceParam(slice2), equaller)
}

// ElementMatchG checks whether two []T slice equal without order, is the generic function of ElementMatch.
func ElementMatchG(slice1, slice2 interface{}) bool {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreElementMatch(s1, s2, defaultEqualler)
}

// ElementMatchWithG checks whether two []T slice equal without order with Equaller, is the generic function of ElementMatchWith.
func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool {
	s1, s2 := checkTwoSliceInterfaceParam(slice1, slice2)
	return coreElementMatch(s1, s2, equaller)
}

// coreElementMatch is the implementation for ElementMatch, ElementMatchWith, ElementMatchG and ElementMatchWithG.
func coreElementMatch(slice1, slice2 innerSlice, equaller Equaller) bool {
	extra1 := makeSameTypeInnerSlice(slice1, 0, 0)
	extra2 := makeSameTypeInnerSlice(slice2, 0, 0)

	length1 := slice1.length()
	length2 := slice2.length()
	visited := make([]bool, length2)
	for idx1 := 0; idx1 < length1; idx1++ {
		item1 := slice1.get(idx1)
		exist := false
		for idx2 := 0; idx2 < length2; idx2++ {
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

	for idx2 := 0; idx2 < length2; idx2++ {
		item2 := slice2.get(idx2)
		if !visited[idx2] {
			extra2.append(item2)
		}
	}

	return extra1.length() == 0 && extra2.length() == 0
}

// Repeat generates a []interface{} with given value repeated given count.
func Repeat(value interface{}, count uint) []interface{} {
	return coreRepeat(value, count, false).actual().([]interface{})
}

// RepeatG generates a []T with given value repeated given count, is the generic function of Repeat.
func RepeatG(value interface{}, count uint) interface{} {
	return coreRepeat(value, count, true).actual().(interface{})
}

// coreRepeat is the implementation for Repeat and RepeatG.
func coreRepeat(value interface{}, count uint, g bool) innerSlice {
	if value == nil {
		g = false
	}
	slice := makeItemTypeInnerSlice(value, int(count), int(count), g)
	for i := 0; i < int(count); i++ {
		slice.set(i, value)
	}
	return slice
}
