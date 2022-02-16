package xpointer

// ===
// ptr
// ===

// BoolPtr returns a pointer pointed to given bool value.
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns a pointer pointed to given string value.
func StringPtr(v string) *string {
	return &v
}

// IntPtr returns a pointer pointed to given int value.
func IntPtr(v int) *int {
	return &v
}

// Int8Ptr returns a pointer pointed to given int8 value.
func Int8Ptr(v int8) *int8 {
	return &v
}

// Int16Ptr returns a pointer pointed to given int16 value.
func Int16Ptr(v int16) *int16 {
	return &v
}

// Int32Ptr returns a pointer pointed to given int32 value.
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int64Ptr returns a pointer pointed to given int64 value.
func Int64Ptr(v int64) *int64 {
	return &v
}

// UintPtr returns a pointer pointed to given uint value.
func UintPtr(v uint) *uint {
	return &v
}

// Uint8Ptr returns a pointer pointed to given uint8 value.
func Uint8Ptr(v uint8) *uint8 {
	return &v
}

// Uint16Ptr returns a pointer pointed to given uint16 value.
func Uint16Ptr(v uint16) *uint16 {
	return &v
}

// Uint32Ptr returns a pointer pointed to given uint32 value.
func Uint32Ptr(v uint32) *uint32 {
	return &v
}

// Uint64Ptr returns a pointer pointed to given uint64 value.
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Float32Ptr returns a pointer pointed to given float32 value.
func Float32Ptr(v float32) *float32 {
	return &v
}

// Float64Ptr returns a pointer pointed to given float64 value.
func Float64Ptr(v float64) *float64 {
	return &v
}

// Complex64Ptr returns a pointer pointed to given complex64 value.
func Complex64Ptr(v complex64) *complex64 {
	return &v
}

// Complex128Ptr returns a pointer pointed to given complex128 value.
func Complex128Ptr(v complex128) *complex128 {
	return &v
}

// BytePtr returns a pointer pointed to given byte value.
func BytePtr(v byte) *byte {
	return &v
}

// RunePtr returns a pointer pointed to given rune value.
func RunePtr(v rune) *rune {
	return &v
}

// ===
// val
// ===

// BoolVal returns a bool value from given pointer, returns the fallback value when nil.
func BoolVal(p *bool, o bool) bool {
	if p == nil {
		return o
	}
	return *p
}

// StringVal returns a string value from given pointer, returns the fallback value when nil.
func StringVal(p *string, o string) string {
	if p == nil {
		return o
	}
	return *p
}

// IntVal returns an int value from given pointer, returns the fallback value when nil.
func IntVal(p *int, o int) int {
	if p == nil {
		return o
	}
	return *p
}

// Int8Val returns an int8 value from given pointer, returns the fallback value when nil.
func Int8Val(p *int8, o int8) int8 {
	if p == nil {
		return o
	}
	return *p
}

// Int16Val returns an int16 value from given pointer, returns the fallback value when nil.
func Int16Val(p *int16, o int16) int16 {
	if p == nil {
		return o
	}
	return *p
}

// Int32Val returns an int32 value from given pointer, returns the fallback value when nil.
func Int32Val(p *int32, o int32) int32 {
	if p == nil {
		return o
	}
	return *p
}

// Int64Val returns an int64 value from given pointer, returns the fallback value when nil.
func Int64Val(p *int64, o int64) int64 {
	if p == nil {
		return o
	}
	return *p
}

// UintVal returns an uint value from given pointer, returns the fallback value when nil.
func UintVal(p *uint, o uint) uint {
	if p == nil {
		return o
	}
	return *p
}

// Uint8Val returns an uint8 value from given pointer, returns the fallback value when nil.
func Uint8Val(p *uint8, o uint8) uint8 {
	if p == nil {
		return o
	}
	return *p
}

// Uint16Val returns an uint16 value from given pointer, returns the fallback value when nil.
func Uint16Val(p *uint16, o uint16) uint16 {
	if p == nil {
		return o
	}
	return *p
}

// Uint32Val returns an uint32 value from given pointer, returns the fallback value when nil.
func Uint32Val(p *uint32, o uint32) uint32 {
	if p == nil {
		return o
	}
	return *p
}

// Uint64Val returns an uint64 value from given pointer, returns the fallback value when nil.
func Uint64Val(p *uint64, o uint64) uint64 {
	if p == nil {
		return o
	}
	return *p
}

// Float32Val returns a float32 value from given pointer, returns the fallback value when nil.
func Float32Val(p *float32, o float32) float32 {
	if p == nil {
		return o
	}
	return *p
}

// Float64Val returns a float64 value from given pointer, returns the fallback value when nil.
func Float64Val(p *float64, o float64) float64 {
	if p == nil {
		return o
	}
	return *p
}

// Complex64Val returns a complex64 value from given pointer, returns the fallback value when nil.
func Complex64Val(p *complex64, o complex64) complex64 {
	if p == nil {
		return o
	}
	return *p
}

// Complex128Val returns a complex128 value from given pointer, returns the fallback value when nil.
func Complex128Val(p *complex128, o complex128) complex128 {
	if p == nil {
		return o
	}
	return *p
}

// ByteVal returns a byte value from given pointer, returns the fallback value when nil.
func ByteVal(p *byte, o byte) byte {
	if p == nil {
		return o
	}
	return *p
}

// RuneVal returns a rune value from given pointer, returns the fallback value when nil.
func RuneVal(p *rune, o rune) rune {
	if p == nil {
		return o
	}
	return *p
}
