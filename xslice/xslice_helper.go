package xslice

import (
	"reflect"
)

// Sti means slice to interface slice.
// Example:
// 	Sti([]int{0, 1}) -> []interface{}{0, 1}
func Sti(slice interface{}) []interface{} {
	if slice == nil {
		return nil
	}

	v := reflect.ValueOf(slice)
	if v.Type().Kind() != reflect.Slice {
		panic("Sti: parameter must be a slice")
	}

	l := v.Len()
	arr := make([]interface{}, l)
	for idx := 0; idx < l; idx++ {
		arr[idx] = v.Index(idx).Interface()
	}
	return arr
}

// Its means interface slice to slice.
// Example:
// 	Its([]interface{}{0, 1}, 0).([]int) -> []int{0, 1}
func Its(slice []interface{}, model interface{}) interface{} {
	if model == nil {
		panic("Its: model must be non-nil")
	}
	if slice == nil {
		return nil
	}

	t := reflect.TypeOf(model)
	l := len(slice)

	out := reflect.MakeSlice(reflect.SliceOf(t), l, l)
	for idx := range slice {
		v := reflect.ValueOf(slice[idx])
		out.Index(idx).Set(v)
	}
	return out.Interface()
}

func StiOfInt(slice []int) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfInt8(slice []int8) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfInt16(slice []int16) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfInt32(slice []int32) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfInt64(slice []int64) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfUint(slice []uint) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfUint8(slice []uint8) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfUint16(slice []uint16) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfUint32(slice []uint32) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfUint64(slice []uint64) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfFloat32(slice []float32) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfFloat64(slice []float64) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfByte(slice []byte) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfRune(slice []rune) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfString(slice []string) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
	}
	return out
}

func StiOfBool(slice []bool) []interface{} {
	out := make([]interface{}, len(slice))
	for idx := range slice {
		out[idx] = slice[idx]
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

func ItsOfInt8(slice []interface{}) []int8 {
	out := make([]int8, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int8)
	}
	return out
}

func ItsOfInt16(slice []interface{}) []int16 {
	out := make([]int16, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(int16)
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

func ItsOfUint(slice []interface{}) []uint {
	out := make([]uint, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint)
	}
	return out
}

func ItsOfUint8(slice []interface{}) []uint8 {
	out := make([]uint8, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint8)
	}
	return out
}

func ItsOfUint16(slice []interface{}) []uint16 {
	out := make([]uint16, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(uint16)
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

func ItsOfString(slice []interface{}) []string {
	out := make([]string, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(string)
	}
	return out
}

func ItsOfBool(slice []interface{}) []bool {
	out := make([]bool, len(slice))
	for idx := range slice {
		out[idx] = slice[idx].(bool)
	}
	return out
}
