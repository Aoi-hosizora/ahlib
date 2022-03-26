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

// TODO
// https://pkg.go.dev/golang.org/x/exp/slices

// ====================
// xslice compatibility
// ====================

// Equaller represents an equality function for two values, is used in XXXWith methods.
type Equaller[T any] func(i, j T) bool

// Lesser represents a less function for sort, see sort.Interface.
type Lesser[T any] func(i, j T) bool

// defaultLesser represents a default Equaller, it just checks order by `<` between xsugar.Ordered types.
func defaultLesser[T xsugar.Ordered]() Lesser[T] {
	return func(i, j T) bool {
		return i < j
	}
}

// defaultEqualler represents a default Equaller, it just checks equality by `==` between comparable types.
func defaultEqualler[T comparable]() Equaller[T] {
	return func(i, j T) bool {
		return i == j
	}
}

func init() {
	// for ShuffleSelf
	rand.Seed(time.Now().UnixNano())
}

// ShuffleSelf shuffles the []T slice directly.
func ShuffleSelf[T any](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle shuffles the []T slice and returns the result.
func Shuffle[T any, S ~[]T](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	ShuffleSelf(out)
	return out
}

// ReverseSelf reverses the []T slice directly.
func ReverseSelf[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Reverse reverses the []T slice and returns the result.
func Reverse[T any, S ~[]T](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	ReverseSelf(out)
	return out
}

// SortSelf sorts the []T slice directly.
func SortSelf[T xsugar.Ordered](slice []T) {
	SortSelfWith(slice, defaultLesser[T]())
}

// Sort sorts the []T slice directly and returns the result.
func Sort[T xsugar.Ordered, S ~[]T](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	SortSelf(out)
	return out
}

// SortSelfWith sorts the []T slice with less function directly.
func SortSelfWith[T any](slice []T, less Lesser[T]) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// SortWith sorts the []T slice with less function and returns the result.
func SortWith[T any, S ~[]T](slice S, less Lesser[T]) S {
	out := make([]T, len(slice))
	copy(out, slice)
	SortSelfWith(out, less)
	return out
}

// StableSortSelf sorts the []T slice in stable directly.
func StableSortSelf[T xsugar.Ordered](slice []T) {
	StableSortSelfWith(slice, defaultLesser[T]())
}

// StableSort sorts the []T slice in stable directly and returns the result.
func StableSort[T xsugar.Ordered, S ~[]T](slice S) S {
	out := make([]T, len(slice))
	copy(out, slice)
	StableSortSelf(out)
	return out
}

// StableSortSelfWith sorts the []T slice in stable with less function directly.
func StableSortSelfWith[T any](slice []T, less Lesser[T]) {
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// StableSortWith sorts the []T slice in stable with less function and returns the result.
func StableSortWith[T any, S ~[]T](slice S, less Lesser[T]) S {
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
	for idx, val := range slice {
		if equaller(val, value) {
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
	for _, val := range slice {
		if equaller(val, value) {
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

// InsertSelf inserts values into []T slice at index position using the space of given slice.
func InsertSelf[T any, S ~[]T](slice S, index int, values ...T) S {
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

// Insert inserts values into []T slice at index position using a new slice space to store.
func Insert[T any, S ~[]T](slice S, index int, values ...T) S {
	out := make([]T, 0, len(slice)+len(values))
	for _, v := range slice {
		out = append(out, v)
	}
	return InsertSelf(S(out), index, values...)
}

// Delete deletes value from []T slice in n times.
func Delete[T comparable, S ~[]T](slice S, value T, n int) S {
	return DeleteWith(slice, value, n, defaultEqualler[T]())
}

// DeleteWith deletes value from []T slice in n times with Equaller.
func DeleteWith[T any, S ~[]T](slice S, value T, n int, equaller Equaller[T]) S {
	out := append(make([]T, 0, len(slice)), slice...)
	if n <= 0 {
		n = len(out)
	}
	cnt := 0
	idx := IndexOfWith(out, value, equaller)
	for idx != -1 && cnt < n {
		if idx == len(out)-1 {
			out = out[:idx]
		} else {
			out = append(out[:idx], out[idx+1:]...)
		}
		cnt++
		idx = IndexOfWith(out, value, equaller)
	}
	return out
}

// DeleteAll deletes value from []T slice in all.
func DeleteAll[T comparable, S ~[]T](slice S, value T) S {
	return DeleteWith(slice, value, -1, defaultEqualler[T]())
}

// DeleteAllWith deletes value from []T slice in all with Equaller.
func DeleteAllWith[T any, S ~[]T](slice S, value T, equaller Equaller[T]) S {
	return DeleteWith(slice, value, -1, equaller)
}

// ContainsAll returns true if values in []T subset are all in the []T list.
func ContainsAll[T comparable](list, subset []T) bool {
	return ContainsAllWith(list, subset, defaultEqualler[T]())
}

// ContainsAllWith returns true if values in []T subset are all in the []T list with Equaller.
func ContainsAllWith[T any](list, subset []T, equaller Equaller[T]) bool {
	for _, val := range subset {
		if !ContainsWith(list, val, equaller) {
			return false
		}
	}
	return true
}

// Diff returns the difference of two []T slices.
func Diff[T comparable](slice1, slice2 []T) []T {
	return DiffWith(slice1, slice2, defaultEqualler[T]())
}

// DiffWith returns the difference of two []T slices with Equaller.
func DiffWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T {
	result := make([]T, 0, 0)
	for _, item1 := range slice1 {
		if !ContainsWith(slice2, item1, equaller) {
			result = append(result, item1)
		}
	}
	return result
}

// Union returns the union of two []T slices.
func Union[T comparable](slice1, slice2 []T) []T {
	return UnionWith(slice1, slice2, defaultEqualler[T]())
}

// UnionWith returns the union of two []T slices with Equaller.
func UnionWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T {
	result := make([]T, len(slice1))
	copy(result, slice1)
	for _, item2 := range slice2 {
		if !ContainsWith(slice1, item2, equaller) {
			result = append(result, item2)
		}
	}
	return result
}

// Intersect returns the intersection of two []T slices.
func Intersect[T comparable](slice1, slice2 []T) []T {
	return IntersectWith(slice1, slice2, defaultEqualler[T]())
}

// IntersectWith returns the intersection of two []T slices with Equaller.
func IntersectWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T {
	result := make([]T, 0, 0)
	for _, item1 := range slice1 {
		if ContainsWith(slice2, item1, equaller) {
			result = append(result, item1)
		}
	}
	return result
}

// Deduplicate removes the duplicate items from []T slice as a set.
func Deduplicate[T comparable, S ~[]T](slice S) S {
	return DeduplicateWith(slice, defaultEqualler[T]())
}

// DeduplicateWith removes the duplicate items from []T slice as a set with Equaller.
func DeduplicateWith[T any, S ~[]T](slice S, equaller Equaller[T]) S {
	result := make([]T, 0, 0)
	for _, item := range slice {
		if !ContainsWith(result, item, equaller) {
			result = append(result, item) // O(n^2)
		}
	}
	return result
}

// ElementMatch checks whether two []T slice equal without order.
func ElementMatch[T comparable](slice1, slice2 []T) bool {
	return ElementMatchWith(slice1, slice2, defaultEqualler[T]())
}

// ElementMatchWith checks whether two []T slice equal without order with Equaller.
func ElementMatchWith[T any](slice1, slice2 []T, equaller Equaller[T]) bool {
	extra1, extra2 := make([]T, 0, 0), make([]T, 0, 0)
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
			extra1 = append(extra1, item1)
		}
	}
	for i2, item2 := range slice2 {
		if !visited[i2] {
			extra2 = append(extra2, item2)
		}
	}
	return len(extra1) == 0 && len(extra2) == 0
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
func Expand[T1, T2 any](slice []T1, f func(T1) []T2) []T2 {
	if f == nil {
		panic(panicNilExpandFunc)
	}
	out := make([]T2, 0, len(slice))
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
func Filter[T any, S ~[]T](slice S, f func(T) bool) S {
	if f == nil {
		panic(panicNilPredicateFunc)
	}
	out := make([]T, 0)
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
