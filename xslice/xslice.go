package xslice

import (
	"math/rand"
	"sync"
	"time"
)

// Equaller represents how two data equal, used in XXXWith methods.
type Equaller func(i, j interface{}) bool

// Shuffle shuffles the slice directly.
func Shuffle(slice []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ShuffleNew shuffles the old slice and return a new one.
func ShuffleNew(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for idx, s := range slice {
		newSlice[idx] = s
	}
	Shuffle(newSlice)
	return newSlice
}

// Reverse reverses the slice directly.
func Reverse(slice []interface{}) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ReverseNew reverse the old slice and return a new one.
func ReverseNew(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for idx, s := range slice {
		newSlice[idx] = s
	}
	Reverse(newSlice)
	return newSlice
}

// ForEach invokes function for each item in slice.
func ForEach(slice []interface{}, fn func(interface{})) {
	for idx := range slice {
		fn(slice[idx])
	}
}

// GoForEach invokes a goroutine with a sync.WaitGroup, which invokes function for each item in slice.
func GoForEach(slice []interface{}, each func(interface{})) {
	if len(slice) == 0 {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for idx := range slice {
		go func(i interface{}) {
			each(i)
			wg.Done()
		}(slice[idx])
	}
	wg.Wait()
}

// Map maps a slice and return a new slice.
func Map(slice []interface{}, mapper func(interface{}) interface{}) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = mapper(slice[idx])
	}
	return out
}

// IndexOfWith returns the first index of value in the slice with Equaller.
func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	for idx, val := range slice {
		if equaller(value, val) {
			return idx
		}
	}
	return -1
}

// IndexOf returns the first index of value in the slice without Equaller.
func IndexOf(slice []interface{}, value interface{}) int {
	return IndexOfWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// ContainsWith checks whether the value is in the slice with Equaller.
func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return IndexOfWith(slice, value, equaller) != -1
}

// Contains checks whether the value is in the slice without Equaller.
func Contains(slice []interface{}, value interface{}) bool {
	return ContainsWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// CountWith returns the count of value in the slice with Equaller.
func CountWith(slice []interface{}, value interface{}, equaller Equaller) int {
	cnt := 0
	for _, val := range slice {
		if equaller(value, val) {
			cnt++
		}
	}
	return cnt
}

// Count returns the count of value in the slice without Equaller.
func Count(slice []interface{}, value interface{}) int {
	return CountWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// DeleteWith deletes a value from slice for n times with Equaller.
func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{} {
	if slice == nil {
		return nil
	}
	cnt := 0
	if n <= 0 {
		n = len(slice)
	}
	idx := IndexOfWith(slice, value, equaller)
	for idx != -1 && cnt < n {
		if len(slice) == idx+1 {
			slice = slice[:idx]
		} else {
			slice = append(slice[:idx], slice[idx+1:]...)
		}
		cnt++
		idx = IndexOfWith(slice, value, equaller)
	}
	return slice
}

// Delete deletes a value from slice for n times without Equaller.
func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	return DeleteWith(slice, value, n, func(i, j interface{}) bool {
		return i == j
	})
}

// DeleteAllWith deletes a value from slice for all with Equaller.
func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	return DeleteWith(slice, value, -1, equaller)
}

// DeleteAll deletes a value from slice for all without Equaller.
func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	return DeleteAllWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// DiffWith returns the difference of two slices with Equaller.
func DiffWith(s1 []interface{}, s2 []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item1 := range s1 {
		exist := false
		for _, item2 := range s2 {
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, item1)
		}
	}
	return result
}

// Diff returns the difference of two slices without Equaller.
func Diff(s1 []interface{}, s2 []interface{}) []interface{} {
	return DiffWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// UnionWith returns the union of two slices with Equaller.
func UnionWith(s1 []interface{}, s2 []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item1 := range s1 {
		result = append(result, item1)
	}
	for _, item2 := range s2 {
		exist := false
		for _, item1 := range s1 {
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, item2)
		}
	}
	return result
}

// Union returns the union of two slices without Equaller.
func Union(s1 []interface{}, s2 []interface{}) []interface{} {
	return UnionWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// IntersectionWith returns the intersection of two slices with Equaller.
func IntersectionWith(s1 []interface{}, s2 []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item1 := range s1 {
		for _, item2 := range s2 {
			if equaller(item1, item2) {
				result = append(result, item1)
				break
			}
		}
	}
	return result
}

// Intersection returns the intersection of two slices without Equaller.
func Intersection(s1 []interface{}, s2 []interface{}) []interface{} {
	return IntersectionWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// ToSetWith removes the duplicate items in a slice as a set with Equaller.
func ToSetWith(slice []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item := range slice {
		if CountWith(result, item, equaller) == 0 {
			result = append(result, item)
		}
	}
	return result
}

// ToSet removes the duplicate items in a slice as a set with Equaller.
func ToSet(s []interface{}) []interface{} {
	return ToSetWith(s, func(i, j interface{}) bool {
		return i == j
	})
}

// EqualWith checks two slice is equal in elements with Equaller.
func EqualWith(s1 []interface{}, s2 []interface{}, equaller Equaller) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, item := range s1 {
		if !ContainsWith(s2, item, equaller) {
			return false
		}
	}
	for _, item := range s2 {
		if !ContainsWith(s1, item, equaller) {
			return false
		}
	}
	return true
}

// Equal checks two slice is equal in elements without Equaller.
func Equal(s1 []interface{}, s2 []interface{}) bool {
	return EqualWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// Range generates an integer slice from small to large with step.
func Range(min, max, step int) []int {
	if min >= max {
		panic("min should less then max")
	} else if step <= 0 {
		panic("step should larger than 0")
	}

	out := make([]int, 0)
	for idx := min; idx <= max; idx += step {
		out = append(out, idx)
	}
	return out
}

// ReverseRange generates an reverse integer slice from small to large with step.
func ReverseRange(min, max, step int) []int {
	if min >= max {
		panic("min should less then max")
	} else if step <= 0 {
		panic("step should larger than 0")
	}

	out := make([]int, 0)
	for idx := max; idx >= min; idx -= step {
		out = append(out, idx)
	}
	return out
}

// GenerateByIndex generates a slice by indies and a generate function.
func GenerateByIndex(indies []int, fn func(i int) interface{}) []interface{} {
	out := make([]interface{}, len(indies))
	for idx, num := range indies {
		out[idx] = fn(num)
	}
	return out
}
