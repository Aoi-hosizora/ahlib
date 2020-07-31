package xslice

import (
	"math/rand"
	"reflect"
)

type Equaller func(i, j interface{}) bool

func Shuffle(slice []interface{}, source rand.Source) {
	random := rand.New(source)
	for i := len(slice) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Reverse(slice []interface{}) {
	if len(slice) == 0 {
		return
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Map(slice []interface{}, mapper func(interface{}) interface{}) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = mapper(slice[idx])
	}
	return out
}

func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	for idx, val := range slice {
		if equaller(val, value) {
			return idx
		}
	}
	return -1
}

func IndexOf(slice []interface{}, value interface{}) int {
	return IndexOfWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return IndexOfWith(slice, value, equaller) != -1
}

func Contains(slice []interface{}, value interface{}) bool {
	return ContainsWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

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

func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	return DeleteWith(slice, value, n, func(i, j interface{}) bool {
		return i == j
	})
}

func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	return DeleteWith(slice, value, -1, equaller)
}

func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	return DeleteAllWith(slice, value, func(i, j interface{}) bool {
		return i == j
	})
}

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

func Diff(s1 []interface{}, s2 []interface{}) []interface{} {
	return DiffWith(s1, s2, func(i, j interface{}) bool {
		return reflect.DeepEqual(i, j)
	})
}

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

func Equal(s1 []interface{}, s2 []interface{}) bool {
	return EqualWith(s1, s2, func(i, j interface{}) bool {
		return i == j
	})
}
