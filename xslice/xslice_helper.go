package xslice

import (
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
