package xslice

import (
	"fmt"
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

func ItsToString(slice []interface{}) []string {
	out := make([]string, len(slice))
	for idx := range slice {
		out[idx] = fmt.Sprintf("%v", slice[idx])
	}
	return out
}

func ItsOfInt(slice []interface{}) []int {
	out := make([]int, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int)
	}
	return out
}

func ItsOfUint(slice []interface{}) []uint {
	out := make([]uint, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint)
	}
	return out
}

func ItsOfUint32(slice []interface{}) []uint32 {
	out := make([]uint32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint32)
	}
	return out
}

func ItsOfUint64(slice []interface{}) []uint64 {
	out := make([]uint64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint64)
	}
	return out
}

func ItsOfInt32(slice []interface{}) []int32 {
	out := make([]int32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int32)
	}
	return out
}

func ItsOfInt64(slice []interface{}) []int64 {
	out := make([]int64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int64)
	}
	return out
}

func ItsOfFloat32(slice []interface{}) []float32 {
	out := make([]float32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(float32)
	}
	return out
}

func ItsOfFloat64(slice []interface{}) []float64 {
	out := make([]float64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(float64)
	}
	return out
}

func ItsOfString(slice []interface{}) []string {
	out := make([]string, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(string)
	}
	return out
}

func ItsOfByte(slice []interface{}) []byte {
	out := make([]byte, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(byte)
	}
	return out
}

func ItsOfRune(slice []interface{}) []rune {
	out := make([]rune, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(rune)
	}
	return out
}
