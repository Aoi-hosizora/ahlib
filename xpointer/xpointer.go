package xpointer

func StringPtr(v string) *string {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

func BytePtr(v byte) *byte {
	return &v
}

func RunePtr(v rune) *rune {
	return &v
}

func IntPtr(v int) *int {
	return &v
}

func Int8Ptr(v int8) *int8 {
	return &v
}

func Int16Ptr(v int16) *int16 {
	return &v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func UintPtr(v uint) *uint {
	return &v
}

func Uint8Ptr(v uint8) *uint8 {
	return &v
}

func Uint16Ptr(v uint16) *uint16 {
	return &v
}

func Uint32Ptr(v uint32) *uint32 {
	return &v
}

func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Float32Ptr(v float32) *float32 {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func Complex64Ptr(v complex64) *complex64 {
	return &v
}

func Complex128Ptr(v complex128) *complex128 {
	return &v
}

func InterfacePtr(v interface{}) *interface{} {
	return &v
}

func StringVal(p *string, defaultValue string) string {
	if p == nil {
		return defaultValue
	}
	return *p
}

func BoolVal(p *bool, defaultValue bool) bool {
	if p == nil {
		return defaultValue
	}
	return *p
}

func ByteVal(p *byte, defaultValue byte) byte {
	if p == nil {
		return defaultValue
	}
	return *p
}

func RuneVal(p *rune, defaultValue rune) rune {
	if p == nil {
		return defaultValue
	}
	return *p
}

func IntVal(p *int, defaultValue int) int {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Int8Val(p *int8, defaultValue int8) int8 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Int16Val(p *int16, defaultValue int16) int16 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Int32Val(p *int32, defaultValue int32) int32 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Int64Val(p *int64, defaultValue int64) int64 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func UintVal(p *uint, defaultValue uint) uint {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Uint8Val(p *uint8, defaultValue uint8) uint8 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Uint16Val(p *uint16, defaultValue uint16) uint16 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Uint32Val(p *uint32, defaultValue uint32) uint32 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Uint64Val(p *uint64, defaultValue uint64) uint64 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Float32Val(p *float32, defaultValue float32) float32 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Float64Val(p *float64, defaultValue float64) float64 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Complex64Val(p *complex64, defaultValue complex64) complex64 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func Complex128Val(p *complex128, defaultValue complex128) complex128 {
	if p == nil {
		return defaultValue
	}
	return *p
}

func InterfaceVal(p *interface{}, defaultValue interface{}) interface{} {
	if p == nil {
		return defaultValue
	}
	return *p
}
