package xslice

import (
	"math/rand"
	"reflect"
)

// Example:
// 		Sti([]int{0, 1}) -> []interface{}{interface{}(0), interface{}(1)}
func Sti(slice interface{}) []interface{} {
	if slice == nil {
		return nil
	}
	v := reflect.ValueOf(slice)
	if v.IsValid() && v.Kind() != reflect.Slice {
		return nil
	}
	arr := make([]interface{}, v.Len())
	for idx := 0; idx < v.Len(); idx++ {
		arr[idx] = v.Index(idx).Interface()
	}
	return arr
}

// Example:
// 		Its([]interface{}{interface{}(0), interface{}(1)}, 0).([]int) -> []int{0, 1}
func Its(slice []interface{}, model interface{}) interface{} {
	if slice == nil || model == nil {
		return nil
	}
	t := reflect.TypeOf(model)
	si := reflect.MakeSlice(reflect.SliceOf(t), len(slice), len(slice))
	for idx := range slice {
		v := reflect.ValueOf(slice[idx])
		si.Index(idx).Set(v)
		// -> panic
	}
	return si.Interface()
}

func Shuffle(slice []interface{}, source rand.Source) {
	random := rand.New(source)
	for i := len(slice) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
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

// Delete the value in slice, n is delete time, -1 for all
// Example:
// 		Its(Delete(Sti([]int{1, 5, 2, 1, 2, 3, 1}), 1, 1), 0).([]int) == []int{5, 2, 1, 2, 3, 1}
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
