//go:build go1.18
// +build go1.18

package xgslice

import (
	"github.com/Aoi-hosizora/ahlib/xgeneric/xsugar"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xtuple"
	"math/rand"
	"sort"
	"time"
)

// ====================
// xslice compatibility
// ====================

// Equaller represents an equality function for two values, is used in XXXWith methods.
type Equaller[T any] func(i, j T) bool

// Equaller2 represents an equality function for two values in different types, is used in XXXWith methods.
type Equaller2[T1, T2 any] func(i T1, j T2) bool

// Lesser represents a less function for sort, see sort.Interface.
type Lesser[T any] func(i, j T) bool

// defaultLesser represents a default Equaller, it just checks order by `<` between xsugar.Ordered types.
func defaultLesser[T xsugar.Ordered]() Lesser[T] {
	return func(i, j T) bool {
		return i < j
	}
}

// defaultEqualler represents a default Equaller, it just checks equality by `==` between two comparable values.
func defaultEqualler[T comparable]() Equaller[T] {
	return func(i, j T) bool {
		return i == j
	}
}

// defaultEqualler2 represents a default Equaller2 with the same type, it just checks equality by `==` between two comparable values.
func defaultEqualler2[T comparable]() Equaller2[T, T] {
	return func(i, j T) bool {
		return i == j
	}
}

func init() {
	// for ShuffleSelf
	rand.Seed(time.Now().UnixNano())
}

// ShuffleSelf shuffles the []T slice, by modifying given slice directly.
func ShuffleSelf[T any](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle shuffles the []T slice and returns the result.
func Shuffle[S ~[]T, T any](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	ShuffleSelf(out)
	return out
}

// ReverseSelf reverses the []T slice, by modifying given slice directly.
func ReverseSelf[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Reverse reverses the []T slice and returns the result.
func Reverse[S ~[]T, T any](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	ReverseSelf(out)
	return out
}

// SortSelf sorts the []T slice, by modifying given slice directly.
func SortSelf[T xsugar.Ordered](slice []T) {
	SortSelfWith(slice, defaultLesser[T]())
}

// Sort sorts the []T slice and returns the result.
func Sort[S ~[]T, T xsugar.Ordered](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	SortSelf(out)
	return out
}

// SortSelfWith sorts the []T slice with less function, by modifying given slice directly.
func SortSelfWith[T any](slice []T, less Lesser[T]) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// SortWith sorts the []T slice with less function and returns the result.
func SortWith[S ~[]T, T any](slice S, less Lesser[T]) S {
	out := make([]T, len(slice))
	copy(out, slice)
	SortSelfWith(out, less)
	return out
}

// StableSortSelf sorts the []T slice in stable, by modifying given slice directly.
func StableSortSelf[T xsugar.Ordered](slice []T) {
	StableSortSelfWith(slice, defaultLesser[T]())
}

// StableSort sorts the []T slice in stable and returns the result.
func StableSort[S ~[]T, T xsugar.Ordered](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	StableSortSelf(out)
	return out
}

// StableSortSelfWith sorts the []T slice in stable with less function, by modifying given slice directly.
func StableSortSelfWith[T any](slice []T, less Lesser[T]) {
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// StableSortWith sorts the []T slice in stable with less function and returns the result.
func StableSortWith[S ~[]T, T any](slice S, less Lesser[T]) S {
	out := make([]T, len(slice))
	copy(out, slice)
	StableSortSelfWith(out, less)
	return out
}

// IndexOf returns the first index of value in the []T slice.
func IndexOf[T comparable](slice []T, value T) int {
	return IndexOfWith(slice, value, defaultEqualler[T]())
}

// IndexOfWith returns the first index of value in the []T slice with Equaller.
func IndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int {
	for idx, item := range slice {
		if equaller(item, value) {
			return idx
		}
	}
	return -1
}

// LastIndexOf returns the last index of value in the []T slice.
func LastIndexOf[T comparable](slice []T, value T) int {
	return LastIndexOfWith(slice, value, defaultEqualler[T]())
}

// LastIndexOfWith returns the last index of value in the []T slice with Equaller.
func LastIndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int {
	for idx := len(slice) - 1; idx >= 0; idx-- {
		if equaller(slice[idx], value) {
			return idx
		}
	}
	return -1
}

// Contains returns true if value is in the []T slice.
func Contains[T comparable](slice []T, value T) bool {
	return ContainsWith(slice, value, defaultEqualler[T]())
}

// ContainsWith returns true if value is in the []T slice with Equaller.
func ContainsWith[T any](slice []T, value T, equaller Equaller[T]) bool {
	for _, item := range slice {
		if equaller(item, value) {
			return true
		}
	}
	return false
}

// Count returns the count of value in the []T slice.
func Count[T comparable](slice []T, value T) int {
	return CountWith(slice, value, defaultEqualler[T]())
}

// CountWith returns the count of value in the []T slice with Equaller.
func CountWith[T any](slice []T, value T, equaller Equaller[T]) int {
	cnt := 0
	for _, val := range slice {
		if equaller(val, value) {
			cnt++
		}
	}
	return cnt
}

// Insert inserts values into []T slice at index position using a new slice space to store.
func Insert[S ~[]T, T any](slice S, index int, values ...T) S {
	out := make([]T, len(slice), len(slice)+len(values))
	copy(out, slice)
	return InsertSelf(S(out), index, values...)
}

// InsertSelf inserts values into []T slice at index position using the space of given slice.
func InsertSelf[S ~[]T, T any](slice S, index int, values ...T) S {
	switch {
	case len(values) == 0:
		return slice
	case len(slice) == 0 || index >= len(slice):
		return append(slice, values...)
	default:
		if index <= 0 {
			index = 0
		}
		expanded := append(slice, values...)
		shifted := append(expanded[:index+len(values)], slice[index:]...)
		copy(shifted[index:], values)
		return shifted
	}
}

// getCap returns the capacity from given capArg and minimum capacity, used by functions which have `capArg ...int` argument.
func getCap(capArg []int, minimum int) int {
	capacity := 0
	if capArg != nil && len(capArg) > 0 {
		capacity = capArg[0]
	}
	if minimum < 0 {
		minimum = 0
	}
	if capacity < minimum {
		capacity = minimum
	}
	return capacity
}

// Delete deletes value from []T slice in n times.
func Delete[S ~[]T, T comparable](slice S, value T, n int, capArg ...int) S {
	return DeleteWith(slice, value, n, defaultEqualler[T](), capArg...)
}

// DeleteWith deletes value from []T slice in n times with Equaller.
func DeleteWith[S ~[]T, T any](slice S, value T, n int, equaller Equaller[T], capArg ...int) S {
	if n <= 0 {
		n = len(slice)
	}
	out := make([]T, 0, getCap(capArg, 0))
	cnt := 0
	for idx, item := range slice { // O(n)
		if cnt >= n {
			out = append(out, slice[idx:]...)
			break
		}
		if equaller(item, value) {
			cnt++
		} else {
			out = append(out, item)
		}
	}
	return out
}

// DeleteAll deletes value from []T slice in all.
func DeleteAll[S ~[]T, T comparable](slice S, value T, capArg ...int) S {
	return DeleteWith(slice, value, -1, defaultEqualler[T](), capArg...)
}

// DeleteAllWith deletes value from []T slice in all with Equaller.
func DeleteAllWith[S ~[]T, T any](slice S, value T, equaller Equaller[T], capArg ...int) S {
	return DeleteWith(slice, value, -1, equaller, capArg...)
}

// DeleteSelf deletes value from []T slice in n times, by modifying given slice directly.
func DeleteSelf[S ~[]T, T comparable](slice S, value T, n int) S {
	return DeleteSelfWith(slice, value, n, defaultEqualler[T]())
}

// DeleteSelfWith deletes value from []T slice in n times with Equaller, by modifying given slice directly.
func DeleteSelfWith[S ~[]T, T any](slice S, value T, n int, equaller Equaller[T]) S {
	if n <= 0 {
		n = len(slice)
	}
	cnt := 0
	idx := IndexOfWith(slice, value, equaller)
	for idx != -1 && cnt < n {
		if idx == len(slice)-1 {
			slice = slice[:idx]
		} else {
			slice = append(slice[:idx], slice[idx+1:]...)
		}
		cnt++
		idx = IndexOfWith(slice, value, equaller) // O(n^2)
	}
	return slice
}

// DeleteAllSelf deletes value from []T slice in all, by modifying given slice directly.
func DeleteAllSelf[S ~[]T, T comparable](slice S, value T) S {
	return DeleteSelfWith(slice, value, -1, defaultEqualler[T]())
}

// DeleteAllSelfWith deletes value from []T slice in all with Equaller, by modifying given slice directly.
func DeleteAllSelfWith[S ~[]T, T any](slice S, value T, equaller Equaller[T]) S {
	return DeleteSelfWith(slice, value, -1, equaller)
}

// ContainsAll returns true if values in []T subset are all in the []T list.
func ContainsAll[T comparable](list, subset []T) bool {
	return ContainsAllWith(list, subset, defaultEqualler[T]())
}

// ContainsAllWith returns true if values in []T subset are all in the []T list with Equaller.
func ContainsAllWith[T any](list, subset []T, equaller Equaller[T]) bool {
	for _, item := range subset {
		if !ContainsWith(list, item, equaller) {
			return false
		}
	}
	return true
}

// Diff returns the difference of two []T slices.
func Diff[S ~[]T, T comparable](slice1, slice2 S, capArg ...int) S {
	return DiffWith(slice1, slice2, defaultEqualler[T](), capArg...)
}

// DiffWith returns the difference of two []T slices with Equaller.
func DiffWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T], capArg ...int) S {
	result := make([]T, 0, getCap(capArg, 0))
	for _, item1 := range slice1 {
		if !ContainsWith(slice2, item1, equaller) {
			result = append(result, item1)
		}
	}
	return result
}

// Union returns the union of two []T slices.
func Union[S ~[]T, T comparable](slice1, slice2 S, capArg ...int) S {
	return UnionWith(slice1, slice2, defaultEqualler[T](), capArg...)
}

// UnionWith returns the union of two []T slices with Equaller.
func UnionWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T], capArg ...int) S {
	result := make([]T, len(slice1), getCap(capArg, len(slice1)))
	copy(result, slice1)
	for _, item2 := range slice2 {
		if !ContainsWith(slice1, item2, equaller) {
			result = append(result, item2)
		}
	}
	return result
}

// Intersect returns the intersection of two []T slices.
func Intersect[S ~[]T, T comparable](slice1, slice2 S, capArg ...int) S {
	return IntersectWith(slice1, slice2, defaultEqualler[T](), capArg...)
}

// IntersectWith returns the intersection of two []T slices with Equaller.
func IntersectWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T], capArg ...int) S {
	result := make([]T, 0, getCap(capArg, 0))
	for _, item1 := range slice1 {
		if ContainsWith(slice2, item1, equaller) {
			result = append(result, item1)
		}
	}
	return result
}

// Deduplicate removes the duplicate items from []T slice as a set.
func Deduplicate[T comparable, S ~[]T](slice S, capArg ...int) S {
	return DeduplicateWith(slice, defaultEqualler[T](), capArg...)
}

// DeduplicateWith removes the duplicate items from []T slice as a set with Equaller.
func DeduplicateWith[S ~[]T, T any](slice S, equaller Equaller[T], capArg ...int) S {
	result := make([]T, 0, getCap(capArg, 0))
	for _, item := range slice {
		if !ContainsWith(result, item, equaller) { // O(n^2)
			result = append(result, item)
		}
	}
	return result
}

// DeduplicateSelf removes the duplicate items from []T slice as a set, by modifying given slice directly.
func DeduplicateSelf[S ~[]T, T comparable](slice S) S {
	return DeduplicateSelfWith(slice, defaultEqualler[T]())
}

// DeduplicateSelfWith removes the duplicate items from []T slice as a set with Equaller, by modifying given slice directly.
func DeduplicateSelfWith[S ~[]T, T any](slice S, equaller Equaller[T]) S {
	if len(slice) <= 1 {
		return slice
	}
	i := 1
	for _, item := range slice[1:] {
		if !ContainsWith(slice[:i], item, equaller) {
			slice[i] = item
			i++
		}
	}
	return slice[:i]
}

// Compact removes the duplicate items in neighbor from []T slice.
func Compact[S ~[]T, T comparable](slice S, capArg ...int) S {
	return CompactWith(slice, defaultEqualler[T](), capArg...)
}

// CompactWith removes the duplicate items in neighbor from []T slice with Equaller.
func CompactWith[S ~[]T, T any](slice S, equaller Equaller[T], capArg ...int) S {
	if len(slice) <= 1 {
		return slice
	}
	result := make([]T, 1, getCap(capArg, 1))
	last := slice[0]
	result[0] = last
	for _, item := range slice[1:] { // O(n)
		if !equaller(item, last) {
			result = append(result, item)
			last = item
		}
	}
	return result
}

// CompactSelf removes the duplicate items in neighbor from []T slice, by modifying given slice directly.
func CompactSelf[S ~[]T, T comparable](slice S) S {
	return CompactSelfWith(slice, defaultEqualler[T]())
}

// CompactSelfWith removes the duplicate items in neighbor from []T slice with Equaller, by modifying given slice directly.
func CompactSelfWith[S ~[]T, T any](slice S, equaller Equaller[T]) S {
	if len(slice) <= 1 {
		return slice
	}
	i := 1
	last := slice[0]
	for _, item := range slice[1:] {
		if !equaller(item, last) {
			slice[i] = item
			i++
			last = item
		}
	}
	return slice[:i]
}

// Equal checks whether two []T slices equal (the same length and all elements equal).
func Equal[T comparable](slice1, slice2 []T) bool {
	return EqualWith(slice1, slice2, defaultEqualler2[T]())
}

// EqualWith checks whether two slices equal (the same length and all elements equal) with Equaller.
func EqualWith[T1, T2 any](slice1 []T1, slice2 []T2, equaller Equaller2[T1, T2]) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for idx := range slice1 {
		if !equaller(slice1[idx], slice2[idx]) {
			return false
		}
	}
	return true
}

// ElementMatch checks whether two []T slices equal (ignore the order of the elements, but the number of duplicate elements should match).
func ElementMatch[T comparable](slice1, slice2 []T) bool {
	return ElementMatchWith(slice1, slice2, defaultEqualler2[T]())
}

// ElementMatchWith checks whether two slices equal (ignore the order of the elements, but the number of duplicate elements should match) with Equaller.
func ElementMatchWith[T1, T2 any](slice1 []T1, slice2 []T2, equaller Equaller2[T1, T2]) bool {
	visited := make([]bool, len(slice2))
	for _, item1 := range slice1 {
		exist := false
		for i2, item2 := range slice2 {
			if visited[i2] {
				continue
			}
			if equaller(item1, item2) {
				visited[i2] = true
				exist = true
				break
			}
		}
		if !exist {
			return false
		}
	}
	for i2 := range slice2 {
		if !visited[i2] {
			return false
		}
	}
	return true
}

// Repeat generates a []T slice with given value and repeat count.
func Repeat[T any](value T, count uint) []T {
	out := make([]T, 0, count)
	for i := 0; i < int(count); i++ {
		out = append(out, value)
	}
	return out
}

// ==================
// fp-style functions
// ==================

const (
	panicNilEachFunc      = "xgslice: nil each function"
	panicNilMapFunc       = "xgslice: nil map function"
	panicNilExpandFunc    = "xgslice: nil expand function"
	panicNilReduceFunc    = "xgslice: nil reduce function"
	panicNilPredicateFunc = "xgslice: nil predicate function"
)

// Foreach invokes given function for each item of given slice.
func Foreach[T any](slice []T, f func(T)) {
	if f == nil {
		panic(panicNilEachFunc)
	}
	for _, item := range slice {
		f(item)
	}
}

// Map maps given slice to another slice using mapper function.
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	if f == nil {
		panic(panicNilMapFunc)
	}
	out := make([]T2, 0, len(slice))
	for _, item := range slice {
		out = append(out, f(item))
	}
	return out
}

// Expand maps and expands given slice to another slice using expand function.
func Expand[T1, T2 any](slice []T1, f func(T1) []T2, capArg ...int) []T2 {
	if f == nil {
		panic(panicNilExpandFunc)
	}
	out := make([]T2, 0, getCap(capArg, 0))
	for _, item := range slice {
		out = append(out, f(item)...)
	}
	return out
}

// Reduce reduces given slice to a single value using initial value and left reducer function.
func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	if f == nil {
		panic(panicNilReduceFunc)
	}
	for _, item := range slice {
		initial = f(initial, item)
	}
	return initial
}

// Filter filters given slice and returns a new slice using given predicate function.
func Filter[S ~[]T, T any](slice S, f func(T) bool, capArg ...int) S {
	if f == nil {
		panic(panicNilPredicateFunc)
	}
	out := make([]T, 0, getCap(capArg, 0))
	for _, item := range slice {
		if f(item) {
			out = append(out, item)
		}
	}
	return out
}

// Any checks whether given slice contains an item that satisfied given predicate function.
func Any[T any](slice []T, f func(T) bool) bool {
	if f == nil {
		panic(panicNilPredicateFunc)
	}
	if len(slice) == 0 {
		return true
	}
	for _, item := range slice {
		if f(item) {
			return true
		}
	}
	return false
}

// All checks whether items from given slice that all satisfied given predicate function.
func All[T any](slice []T, f func(T) bool) bool {
	if f == nil {
		panic(panicNilPredicateFunc)
	}
	for _, item := range slice {
		if !f(item) {
			return false
		}
	}
	return true
}

// Zip zips given two slices to a tuple slice, its length is the less one of two slices.
func Zip[T1, T2 any](slice1 []T1, slice2 []T2) []xtuple.Tuple[T1, T2] {
	l := len(slice1)
	if l2 := len(slice2); l2 < l {
		l = l2
	}
	out := make([]xtuple.Tuple[T1, T2], 0, l)
	for i := 0; i < l; i++ {
		out = append(out, xtuple.NewTuple(slice1[i], slice2[i]))
	}
	return out
}

// Zip3 zips given three slices to a triple slice, its length is the less one of three slices.
func Zip3[T1, T2, T3 any](slice1 []T1, slice2 []T2, slice3 []T3) []xtuple.Triple[T1, T2, T3] {
	l := len(slice1)
	if l2 := len(slice2); l2 < l {
		l = l2
	}
	if l3 := len(slice3); l3 < l {
		l = l3
	}
	out := make([]xtuple.Triple[T1, T2, T3], 0, l)
	for i := 0; i < l; i++ {
		out = append(out, xtuple.NewTriple(slice1[i], slice2[i], slice3[i]))
	}
	return out
}

// Unzip unzips given tuple slice to two slices.
func Unzip[T1, T2 any](slice []xtuple.Tuple[T1, T2]) ([]T1, []T2) {
	slice1 := make([]T1, 0, len(slice))
	slice2 := make([]T2, 0, len(slice))
	for _, item := range slice {
		slice1 = append(slice1, item.Item1)
		slice2 = append(slice2, item.Item2)
	}
	return slice1, slice2
}

// Unzip3 unzips given triple slice to three slices.
func Unzip3[T1, T2, T3 any](slice []xtuple.Triple[T1, T2, T3]) ([]T1, []T2, []T3) {
	slice1 := make([]T1, 0, len(slice))
	slice2 := make([]T2, 0, len(slice))
	slice3 := make([]T3, 0, len(slice))
	for _, item := range slice {
		slice1 = append(slice1, item.Item1)
		slice2 = append(slice2, item.Item2)
		slice3 = append(slice3, item.Item3)
	}
	return slice1, slice2, slice3
}

// =====================================
// funtions from golang.org/x/exp/slices
// =====================================

// Clone returns a copy of given slice, and the elements are copied using assignment.
func Clone[S ~[]T, T any](slice S) S {
	newSlice := make(S, len(slice), cap(slice))
	for i := 0; i < len(slice); i++ {
		newSlice[i] = slice[i]
	}
	return newSlice
}

// Clip removes unused capacity from given slice, returning slice[:len(s):len(s)].
func Clip[S ~[]T, T any](slice S) S {
	return slice[:len(slice):len(slice)]
}

// Grow increases given slice's capacity, if necessary, to guarantee space for another n elements.
func Grow[S ~[]T, T any](slice S, n int) S {
	if n < 0 {
		n = 0
	}
	return append(slice, make(S, n)...)[:len(slice)]
}
