package xcollection

import "reflect"

// Convert specific slice to slice that element type in interface{}
// Example:
// 		Sti([]int{0, 1}) -> []interface{}{interface{}(0), interface{}(1)}
func Sti(slice interface{}) []interface{} {
	if slice == nil {
		return nil
	}
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	if v.IsValid() && t.Kind() != reflect.Slice {
		return nil
	}
	arr := make([]interface{}, v.Len())
	for idx := 0; idx < v.Len(); idx++ {
		arr[idx] = v.Index(idx).Interface()
	}
	return arr
}

// Convert slice that element type is interface{} to specific slice
// Example:
// 		Its([]interface{}{interface{}(0), interface{}(1)}, reflect.TypeOf(0)).([]int) -> []int{0, 1}
func Its(slice []interface{}, elType reflect.Type) interface{} {
	if slice == nil || elType == nil {
		return nil
	}
	si := reflect.MakeSlice(reflect.SliceOf(elType), len(slice), len(slice))
	for idx := range slice {
		v := reflect.ValueOf(slice[idx])
		if v.Type() != elType {
			return nil
		}
		si.Index(idx).Set(v)
	}
	return si.Interface()
}

func IndexOfSlice(slice []interface{}, value interface{}) (index int) {
	for idx, val := range slice {
		if val == value {
			return idx
		}
	}
	return -1
}

func DeleteInSlice(slice []interface{}, value interface{}, n int) []interface{} {
	cnt := 0
	if n <= 0 {
		n = len(slice)
	}
	idx := IndexOfSlice(slice, value)
	for idx != -1 && cnt < n {
		if len(slice) == idx+1 {
			slice = slice[:idx]
		} else {
			slice = append(slice[:idx], slice[idx+1:]...)
		}
		cnt++
		idx = IndexOfSlice(slice, value)
	}
	return slice
}
