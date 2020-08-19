package xslice

import (
	"math/rand"
	"sync"
	"time"
)

// An Equaller represents how two data equal, used in `XXXWith` methods.
type Equaller func(i, j interface{}) bool

// Shuffle the slice directly.
func Shuffle(slice []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle the old slice and return a new one.
func ShuffleNew(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for idx, s := range slice {
		newSlice[idx] = s
	}
	Shuffle(newSlice)
	return newSlice
}

// Reverse the slice directly.
func Reverse(slice []interface{}) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Reverse the old slice and return a new one.
func ReverseNew(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for idx, s := range slice {
		newSlice[idx] = s
	}
	Reverse(newSlice)
	return newSlice
}

// ForEach each item in slice.
func ForEach(slice []interface{}, each func(interface{})) {
	for idx := range slice {
		each(slice[idx])
	}
}

// Use `go` and `WaitGroup` to ForEach the slice concurrently.
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

// `Map` a slice and return a new slice.
func Map(slice []interface{}, mapper func(interface{}) interface{}) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = mapper(slice[idx])
	}
	return out
}

// `IndexOf` with equaller.
func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	for idx, val := range slice {
		if equaller(val, value) {
			return idx
		}
	}
	return -1
}

// `IndexOf` with normal equal.
func IndexOf(slice []interface{}, value interface{}) int {
	return IndexOfWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// `Contains` with equaller.
func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return IndexOfWith(slice, value, equaller) != -1
}

// `Contains` with normal equal.
func Contains(slice []interface{}, value interface{}) bool {
	return ContainsWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// `Count` with equaller.
func CountWith(slice []interface{}, value interface{}, equaller Equaller) int {
	cnt := 0
	for _, item := range slice {
		if equaller(item, value) {
			cnt++
		}
	}
	return cnt
}

// `Count` with equaller.
func Count(slice []interface{}, value interface{}) int {
	return CountWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// `Delete` with equaller.
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

// `Delete` with normal equal.
func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	return DeleteWith(slice, value, n, func(i, j interface{}) bool {
		return i == j
	})
}

// `DeleteAll` with equaller.
func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	return DeleteWith(slice, value, -1, equaller)
}

// `DeleteAll` with normal equal.
func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	return DeleteAllWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

// `Diff` with equaller.
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

// `Diff` with normal equal.
func Diff(s1 []interface{}, s2 []interface{}) []interface{} {
	return DiffWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// `Union` with equaller.
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

// `Union` with normal equal.
func Union(s1 []interface{}, s2 []interface{}) []interface{} {
	return UnionWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// `Intersection` with equaller.
func IntersectionWith(s1 []interface{}, s2 []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item1 := range s1 {
		for _, item2 := range s2 {
			if equaller(item1, item2) {
				result = append(result, item1)
			}
		}
	}
	return result
}

// `Intersection` with normal equal.
func Intersection(s1 []interface{}, s2 []interface{}) []interface{} {
	return IntersectionWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// `Equal` with equaller.
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
		if !ContainsWith(s2, item, equaller) {
			return false
		}
	}
	return true
}

// `Equal` with normal equal.
func Equal(s1 []interface{}, s2 []interface{}) bool {
	return EqualWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}

// `ToSet` with equaller.
func ToSetWith(slice []interface{}, equaller Equaller) []interface{} {
	result := make([]interface{}, 0)
	for _, item := range slice {
		if CountWith(result, item, equaller) == 0 {
			result = append(result, item)
		}
	}
	return result
}

// `ToSet` with normal equal.
func ToSet(s []interface{}) []interface{} {
	return ToSetWith(s, func(i, j interface{}) bool {
		return i == j
	})
}
