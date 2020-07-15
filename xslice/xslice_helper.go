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
		panic("interface{} is not a slice")
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

// noinspection GoUnusedExportedFunction
func ItsToString(slice []interface{}) []string {
	out := make([]string, len(slice))
	for idx := range slice {
		out[idx] = fmt.Sprintf("%v", slice[idx])
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfInt(slice []interface{}) []int {
	out := make([]int, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfUint(slice []interface{}) []uint {
	out := make([]uint, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfUint32(slice []interface{}) []uint32 {
	out := make([]uint32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint32)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfUint64(slice []interface{}) []uint64 {
	out := make([]uint64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint64)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfInt32(slice []interface{}) []int32 {
	out := make([]int32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int32)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfInt64(slice []interface{}) []int64 {
	out := make([]int64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int64)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfFloat32(slice []interface{}) []float32 {
	out := make([]float32, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(float32)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfFloat64(slice []interface{}) []float64 {
	out := make([]float64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(float64)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfComplex64(slice []interface{}) []complex64 {
	out := make([]complex64, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(complex64)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfComplex128(slice []interface{}) []complex128 {
	out := make([]complex128, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(complex128)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfString(slice []interface{}) []string {
	out := make([]string, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(string)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfByte(slice []interface{}) []byte {
	out := make([]byte, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(byte)
	}
	return out
}

// noinspection GoUnusedExportedFunction
func ItsOfRune(slice []interface{}) []rune {
	out := make([]rune, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(rune)
	}
	return out
}
