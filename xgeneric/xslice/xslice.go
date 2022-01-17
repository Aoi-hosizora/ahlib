//go:build go1.18
// +build go1.18

package xslice

import (
	"constraints"
	"math/rand"
	"sort"
	"time"
)

// Equaller represents an equality function for two values, is used in XXXWith methods.
type Equaller[T any] func(i, j T) bool

// Lesser represents a less function for sort, see sort.Interface.
type Lesser[T any] func(i, j T) bool

// defaultLesser represents a default Equaller, it just checks order by `<` with constraints.Ordered.
func defaultLesser[T constraints.Ordered]() Lesser[T] {
	return func(i, j T) bool {
		return i < j
	}
}

// defaultEqualler represents a default Equaller, it just checks equality by `==` with comparable.
func defaultEqualler[T comparable]() Equaller[T] {
	return func(i, j T) bool {
		return i == j
	}
}

// ShuffleSelf shuffles the []T slice directly.
func ShuffleSelf[T any](slice []T) {
	rand.Seed(time.Now().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle shuffles the []T slice and returns the result.
func Shuffle[T any](slice []T) []T {
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
func Reverse[T any](slice []T) []T {
	out := make([]T, len(slice))
	copy(out, slice)
	ReverseSelf(out)
	return out
}

// SortSelf sorts the []T slice directly.
func SortSelf[T constraints.Ordered](slice []T) {
	SortSelfWith(slice, defaultLesser[T]())
}

// Sort sorts the []T slice directly and returns the result.
func Sort[T constraints.Ordered](slice []T) []T {
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
func SortWith[T any](slice []T, less Lesser[T]) []T {
	out := make([]T, len(slice))
	copy(out, slice)
	SortSelfWith(out, less)
	return out
}

// StableSortSelf sorts the []T slice in stable directly.
func StableSortSelf[T constraints.Ordered](slice []T) {
	StableSortSelfWith(slice, defaultLesser[T]())
}

// StableSort sorts the []T slice in stable directly and returns the result.
func StableSort[T constraints.Ordered](slice []T) []T {
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
func StableSortWith[T any](slice []T, less Lesser[T]) []T {
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

// Insert inserts value to the position of index in []T slice.
func Insert[T any](slice []T, value T, index int) []T {
	if len(slice) == 0 {
		return []T{value}
	}
	if index <= 0 {
		return append([]T{value}, slice...)
	}
	if index >= len(slice) {
		return append(slice, value)
	}
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

// Delete deletes value from []T slice in n times.
func Delete[T comparable](slice []T, value T, n int) []T {
	return DeleteWith(slice, value, n, defaultEqualler[T]())
}

// DeleteWith deletes value from []T slice in n times with Equaller.
func DeleteWith[T any](slice []T, value T, n int, equaller Equaller[T]) []T {
	slice = append(make([]T, 0, len(slice)), slice...) // <<<
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
		idx = IndexOfWith(slice, value, equaller)
	}
	return slice
}

// DeleteAll deletes value from []T slice in all.
func DeleteAll[T comparable](slice []T, value T) []T {
	return DeleteWith(slice, value, -1, defaultEqualler[T]())
}

// DeleteAllWith deletes value from []T slice in all with Equaller.
func DeleteAllWith[T any](slice []T, value T, equaller Equaller[T]) []T {
	return DeleteWith(slice, value, -1, equaller)
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
func Deduplicate[T comparable](slice []T) []T {
	return DeduplicateWith(slice, defaultEqualler[T]())
}

// DeduplicateWith removes the duplicate items from []T slice as a set with Equaller.
func DeduplicateWith[T any](slice []T, equaller Equaller[T]) []T {
	result := make([]T, 0, 0)
	for _, item := range slice {
		if !ContainsWith(result, item, equaller) {
			result = append(result, item)
		}
	}
	return result
}

// ElementMatch checks if two []T slice equal without order.
func ElementMatch[T comparable](slice1, slice2 []T) bool {
	return ElementMatchWith(slice1, slice2, defaultEqualler[T]())
}

// ElementMatchWith checks if two []T slice equal without order with Equaller.
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
