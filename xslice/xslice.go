package xslice

import (
	"math/rand"
	"reflect"
)

func Shuffle(slice []interface{}, source rand.Source) {
	random := rand.New(source)
	for i := len(slice) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Reverse(slice []interface{}) []interface{} {
	sliceCopy := slice
	if len(sliceCopy) == 0 {
		return sliceCopy
	}
	for i, j := 0, len(sliceCopy)-1; i < j; i, j = i+1, j-1 {
		sliceCopy[i], sliceCopy[j] = sliceCopy[j], sliceCopy[i]
	}
	return sliceCopy
}

func IndexOf(slice []interface{}, value interface{}) (index int) {
	for idx, val := range slice {
		if val == value {
			return idx
		}
	}
	return -1
}

func Contains(slice []interface{}, value interface{}) bool {
	return IndexOf(slice, value) != -1
}

func Map(slice []interface{}, mapFunc func(interface{}) interface{}) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = mapFunc(slice[idx])
	}
	return out
}

// Delete the value in slice, n is delete time, -1 for all
func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	if slice == nil {
		return nil
	}
	cnt := 0
	if n <= 0 {
		n = len(slice)
	}
	idx := IndexOf(slice, value)
	for idx != -1 && cnt < n {
		if len(slice) == idx+1 {
			slice = slice[:idx]
		} else {
			slice = append(slice[:idx], slice[idx+1:]...)
		}
		cnt++
		idx = IndexOf(slice, value)
	}
	return slice
}

func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	return Delete(slice, value, -1)
}

func SliceDiff(s1 []interface{}, s2 []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, item1 := range s1 {
		exist := false
		for _, item2 := range s2 {
			if reflect.DeepEqual(item1, item2) {
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

func Equal(s1 []interface{}, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, item := range s1 {
		if !Contains(s2, item) {
			return false
		}
	}
	for _, item := range s2 {
		if !Contains(s1, item) {
			return false
		}
	}
	return true
}
