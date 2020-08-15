package xpointer

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestPtr(t *testing.T) {
	stringVal := "0"
	xtesting.Equal(t, StringPtr("0"), &stringVal)
	boolVal := true
	xtesting.Equal(t, BoolPtr(true), &boolVal)
	byteVal := byte(0)
	xtesting.Equal(t, BytePtr(0), &byteVal)
	runeVal := rune(0)
	xtesting.Equal(t, RunePtr(0), &runeVal)
	intVal := 0
	xtesting.Equal(t, IntPtr(0), &intVal)
	int8Val := int8(0)
	xtesting.Equal(t, Int8Ptr(0), &int8Val)
	int16Val := int16(0)
	xtesting.Equal(t, Int16Ptr(0), &int16Val)
	int32Val := int32(0)
	xtesting.Equal(t, Int32Ptr(0), &int32Val)
	int64Val := int64(0)
	xtesting.Equal(t, Int64Ptr(0), &int64Val)
	uintVal := uint(0)
	xtesting.Equal(t, UintPtr(0), &uintVal)
	uint8Val := uint8(0)
	xtesting.Equal(t, Uint8Ptr(0), &uint8Val)
	uint16Val := uint16(0)
	xtesting.Equal(t, Uint16Ptr(0), &uint16Val)
	uint32Val := uint32(0)
	xtesting.Equal(t, Uint32Ptr(0), &uint32Val)
	uint64Val := uint64(0)
	xtesting.Equal(t, Uint64Ptr(0), &uint64Val)
	float32Val := float32(0.1)
	xtesting.Equal(t, Float32Ptr(0.1), &float32Val)
	float64Val := 0.1
	xtesting.Equal(t, Float64Ptr(0.1), &float64Val)
	complex64Val := complex64(0 + 1i)
	xtesting.Equal(t, Complex64Ptr(0+1i), &complex64Val)
	complex128Val := 0 + 1i
	xtesting.Equal(t, Complex128Ptr(0+1i), &complex128Val)
	interfaceVal := interface{}(nil)
	xtesting.Equal(t, InterfacePtr(nil), &interfaceVal)
}

func TestVal(t *testing.T) {
	xtesting.Equal(t, StringVal(nil, "0"), "0")
	xtesting.Equal(t, BoolVal(nil, false), false)
	xtesting.Equal(t, ByteVal(nil, 0), byte(0))
	xtesting.Equal(t, RuneVal(nil, 0), rune(0))
	xtesting.Equal(t, IntVal(nil, 0), 0)
	xtesting.Equal(t, Int8Val(nil, 0), int8(0))
	xtesting.Equal(t, Int16Val(nil, 0), int16(0))
	xtesting.Equal(t, Int32Val(nil, 0), int32(0))
	xtesting.Equal(t, Int64Val(nil, 0), int64(0))
	xtesting.Equal(t, UintVal(nil, 0), uint(0))
	xtesting.Equal(t, Uint8Val(nil, 0), uint8(0))
	xtesting.Equal(t, Uint16Val(nil, 0), uint16(0))
	xtesting.Equal(t, Uint32Val(nil, 0), uint32(0))
	xtesting.Equal(t, Uint64Val(nil, 0), uint64(0))
	xtesting.Equal(t, Float32Val(nil, 0.1), float32(0.1))
	xtesting.Equal(t, Float64Val(nil, 0.1), 0.1)
	xtesting.Equal(t, Complex64Val(nil, 0+1i), complex64(0+1i))
	xtesting.Equal(t, Complex128Val(nil, 0+1i), 0+1i)
	xtesting.Equal(t, InterfaceVal(nil, nil), nil)

	stringVal := "0"
	xtesting.Equal(t, StringVal(&stringVal, ""), stringVal)
	boolVal := true
	xtesting.Equal(t, BoolVal(&boolVal, false), boolVal)
	byteVal := byte(0)
	xtesting.Equal(t, ByteVal(&byteVal, 0), byteVal)
	runeVal := rune(0)
	xtesting.Equal(t, RuneVal(&runeVal, 0), runeVal)
	intVal := 0
	xtesting.Equal(t, IntVal(&intVal, 0), intVal)
	int8Val := int8(0)
	xtesting.Equal(t, Int8Val(&int8Val, 0), int8Val)
	int16Val := int16(0)
	xtesting.Equal(t, Int16Val(&int16Val, 0), int16Val)
	int32Val := int32(0)
	xtesting.Equal(t, Int32Val(&int32Val, 0), int32Val)
	int64Val := int64(0)
	xtesting.Equal(t, Int64Val(&int64Val, 0), int64Val)
	uintVal := uint(0)
	xtesting.Equal(t, UintVal(&uintVal, 0), uintVal)
	uint8Val := uint8(0)
	xtesting.Equal(t, Uint8Val(&uint8Val, 0), uint8Val)
	uint16Val := uint16(0)
	xtesting.Equal(t, Uint16Val(&uint16Val, 0), uint16Val)
	uint32Val := uint32(0)
	xtesting.Equal(t, Uint32Val(&uint32Val, 0), uint32Val)
	uint64Val := uint64(0)
	xtesting.Equal(t, Uint64Val(&uint64Val, 0), uint64Val)
	float32Val := float32(0.1)
	xtesting.Equal(t, Float32Val(&float32Val, 0), float32Val)
	float64Val := 0.1
	xtesting.Equal(t, Float64Val(&float64Val, 0), float64Val)
	complex64Val := complex64(0 + 1i)
	xtesting.Equal(t, Complex64Val(&complex64Val, 0), complex64Val)
	complex128Val := 0 + 1i
	xtesting.Equal(t, Complex128Val(&complex128Val, 0), complex128Val)
	interfaceVal := interface{}(0)
	xtesting.Equal(t, InterfaceVal(&interfaceVal, 0), interfaceVal)
}
