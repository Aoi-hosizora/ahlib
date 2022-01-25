package xnumber

import (
	"fmt"
	"math"
	_ "runtime"
	_ "unsafe"
)

// ================
// accuracy related
// ================

// Accuracy represents an accuracy with some compare methods in epsilon accuracy.
type Accuracy func() float64

// NewAccuracy creates an Accuracy, using eps as its accuracy.
func NewAccuracy(eps float64) Accuracy {
	return func() float64 {
		return eps
	}
}

// Equal checks equality (eq) between two float64.
func (eps Accuracy) Equal(a, b float64) bool {
	return math.Abs(a-b) < eps()
}

// NotEqual checks not-equality (ne) between two float64.
func (eps Accuracy) NotEqual(a, b float64) bool {
	return math.Abs(a-b) >= eps()
}

// Greater checks if float `a` is greater than (gt) `b`.
func (eps Accuracy) Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > eps()
}

// Less checks if float `a` is less than (lt) `b`.
func (eps Accuracy) Less(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > eps()
}

// GreaterOrEqual checks if float `a` is greater than or equals to (gte) `b`.
func (eps Accuracy) GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < eps()
}

// LessOrEqual checks if float `a` is less than or equals to (lte) `b`.
func (eps Accuracy) LessOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < eps()
}

// _acc is a default Accuracy with 1e-3 eps.
var _acc = NewAccuracy(1e-3)

// EqualInAccuracy checks equality (eq) between two float64 with 1e-3 accuracy.
func EqualInAccuracy(a, b float64) bool {
	return _acc.Equal(a, b)
}

// NotEqualInAccuracy checks not-equality (ne) between two float64 with 1e-3 accuracy.
func NotEqualInAccuracy(a, b float64) bool {
	return _acc.NotEqual(a, b)
}

// GreaterInAccuracy checks if float `a` is greater than (gt) `b` with 1e-3 accuracy.
func GreaterInAccuracy(a, b float64) bool {
	return _acc.Greater(a, b)
}

// LessInAccuracy checks if float `a` is less than (lt) `b` with 1e-3 accuracy.
func LessInAccuracy(a, b float64) bool {
	return _acc.Less(a, b)
}

// GreaterOrEqualInAccuracy checks if float `a` is greater than or equals to (gte) `b` with 1e-3 accuracy.
func GreaterOrEqualInAccuracy(a, b float64) bool {
	return _acc.GreaterOrEqual(a, b)
}

// LessOrEqualInAccuracy checks if float `a` is less than or equals to (lte) `b` with 1e-3 accuracy.
func LessOrEqualInAccuracy(a, b float64) bool {
	return _acc.LessOrEqual(a, b)
}

// ==============
// mass functions
// ==============

// FormatByteSize formats a byte size to string (using %.2f), supports `B` `KB` `MB` `GB` `TB` units.
func FormatByteSize(bytes float64) string {
	divisor := float64(1024)
	minus := false
	switch {
	case bytes == 0:
		return "0B"
	case bytes < 0:
		bytes = -bytes
		minus = true
	}
	ret := func(s string) string {
		if minus {
			return "-" + s
		}
		return s
	}

	// 1 - 1023B
	b := bytes
	if LessInAccuracy(b, divisor) {
		return ret(fmt.Sprintf("%dB", int(b)))
	}
	// 1 - 1023K
	kb := b / divisor
	if LessInAccuracy(kb, divisor) {
		return ret(fmt.Sprintf("%.2fKB", kb))
	}
	// 1 - 1023M
	mb := kb / divisor
	if LessInAccuracy(mb, divisor) {
		return ret(fmt.Sprintf("%.2fMB", mb))
	}
	// 1 - 1023G
	gb := mb / divisor
	if LessInAccuracy(gb, divisor) {
		return ret(fmt.Sprintf("%.2fGB", gb))
	}
	// 1T -
	tb := gb / divisor
	return ret(fmt.Sprintf("%.2fTB", tb))
}

// Bool returns 1 if the given value is true, otherwise returns 0.
func Bool(b bool) int {
	if b {
		return 1
	}
	return 0
}

// intBitLength represents the int or uint bit-length, usually it equals to 32 or 64.
var intBitLength = 32 << (^uint(0) >> 63) // <<< it should be `const` but use `var` for testing coverage

// IntBitLength returns the int or uint bit-length, usually it equals to 32 or 64.
func IntBitLength() int {
	return intBitLength
}

//go:linkname FastrandUint32 runtime.fastrand

// FastrandUint32 returns a random uint32 value using runtime.fastrand.
func FastrandUint32() uint32

// FastrandUint64 returns a random uint64 value using runtime.fastrand.
func FastrandUint64() uint64 {
	return (uint64(FastrandUint32()) << 32) | uint64(FastrandUint32())
}

// FastrandInt32 returns a random int32 value using runtime.fastrand.
func FastrandInt32() int32 {
	return int32(FastrandUint32() & (1<<31 - 1))
}

// FastrandInt64 returns a random int64 value using runtime.fastrand.
func FastrandInt64() int64 {
	return int64(FastrandUint64() & (1<<63 - 1))
}

// IsPowerOfTwo checks whether the given integer is power of two.
func IsPowerOfTwo(x int) bool {
	return (x & (-x)) == x
}

// =====================
// numeric range related
// =====================

const (
	MinInt8   = int8(-128)                  // -1 << 7,  see math.MinInt8.
	MinInt16  = int16(-32768)               // -1 << 15, see math.MinInt16.
	MinInt32  = int32(-2147483648)          // -1 << 31, see math.MinInt32.
	MinInt64  = int64(-9223372036854775808) // -1 << 63, see math.MinInt64.
	MinUint8  = uint8(0)                    // 0.
	MinUint16 = uint16(0)                   // 0.
	MinUint32 = uint32(0)                   // 0.
	MinUint64 = uint64(0)                   // 0.

	MaxInt8   = int8(127)                    // 1 << 7  - 1, see math.MaxInt8.
	MaxInt16  = int16(32767)                 // 1 << 15 - 1, see math.MaxInt16.
	MaxInt32  = int32(2147483647)            // 1 << 31 - 1, see math.MaxInt32.
	MaxInt64  = int64(9223372036854775807)   // 1 << 63 - 1, see math.MaxInt64.
	MaxUint8  = uint8(255)                   // 1 << 8  - 1, see math.MaxUint8.
	MaxUint16 = uint16(65535)                // 1 << 16 - 1, see math.MaxUint16.
	MaxUint32 = uint32(4294967295)           // 1 << 32 - 1, see math.MaxUint32.
	MaxUint64 = uint64(18446744073709551615) // 1 << 64 - 1, see math.MaxUint64.

	MaxFloat32             = float32(math.MaxFloat32)             // 2**127 * (2**24 - 1) / 2**23, see math.MaxFloat32.
	SmallestNonzeroFloat32 = float32(math.SmallestNonzeroFloat32) // 1 / 2**(127 - 1 + 23), see math.SmallestNonzeroFloat32.
	MaxFloat64             = float64(math.MaxFloat64)             // 2**1023 * (2**53 - 1) / 2**52, see math.MaxFloat64.
	SmallestNonzeroFloat64 = float64(math.SmallestNonzeroFloat64) // 1 / 2**(1023 - 1 + 52), see math.SmallestNonzeroFloat64.
)

// OverflowWhenAddInt8 checks whether overflow will happen when add int8 addend to int8 augend.
func OverflowWhenAddInt8(augend, addend int8) bool {
	return (augend > 0 && addend > 0 && augend > MaxInt8-addend) || (augend < 0 && addend < 0 && augend < MinInt8-addend)
}

// OverflowWhenAddInt16 checks whether overflow will happen when add int16 addend to int16 augend.
func OverflowWhenAddInt16(augend, addend int16) bool {
	return (augend > 0 && addend > 0 && augend > MaxInt16-addend) || (augend < 0 && addend < 0 && augend < MinInt16-addend)
}

// OverflowWhenAddInt32 checks whether overflow will happen when add int32 addend to int32 augend.
func OverflowWhenAddInt32(augend, addend int32) bool {
	return (augend > 0 && addend > 0 && augend > MaxInt32-addend) || (augend < 0 && addend < 0 && augend < MinInt32-addend)
}

// OverflowWhenAddInt64 checks whether overflow will happen when add int64 addend to int64 augend.
func OverflowWhenAddInt64(augend, addend int64) bool {
	return (augend > 0 && addend > 0 && augend > MaxInt64-addend) || (augend < 0 && addend < 0 && augend < MinInt64-addend)
}

// OverflowWhenAddInt checks whether overflow will happen when add int addend to int augend.
func OverflowWhenAddInt(augend, addend int) bool {
	if intBitLength <= 32 {
		return OverflowWhenAddInt32(int32(augend), int32(addend))
	}
	return OverflowWhenAddInt64(int64(augend), int64(addend))
}

// OverflowWhenSubtractInt8 checks whether overflow will happen when subtract int8 subtrahend from int8 minuend.
func OverflowWhenSubtractInt8(minuend, subtrahend int8) bool {
	return OverflowWhenAddInt8(minuend, -subtrahend)
}

// OverflowWhenSubtractInt16 checks whether overflow will happen when subtract int16 subtrahend from int16 minuend.
func OverflowWhenSubtractInt16(minuend, subtrahend int16) bool {
	return OverflowWhenAddInt16(minuend, -subtrahend)
}

// OverflowWhenSubtractInt32 checks whether overflow will happen when subtract int32 subtrahend from int32 minuend.
func OverflowWhenSubtractInt32(minuend, subtrahend int32) bool {
	return OverflowWhenAddInt32(minuend, -subtrahend)
}

// OverflowWhenSubtractInt64 checks whether overflow will happen when subtract int64 subtrahend from int64 minuend.
func OverflowWhenSubtractInt64(minuend, subtrahend int64) bool {
	return OverflowWhenAddInt64(minuend, -subtrahend)
}

// OverflowWhenSubtractInt checks whether overflow will happen when subtract int subtrahend from int minuend.
func OverflowWhenSubtractInt(minuend, subtrahend int) bool {
	if intBitLength <= 32 {
		return OverflowWhenSubtractInt32(int32(minuend), int32(subtrahend))
	}
	return OverflowWhenSubtractInt64(int64(minuend), int64(subtrahend))
}

// OverflowWhenAddUint8 checks whether overflow will happen when add uint8 addend to uint8 augend.
func OverflowWhenAddUint8(augend, addend uint8) bool {
	return augend > 0 && addend > 0 && augend > MaxUint8-addend
}

// OverflowWhenAddUint16 checks whether overflow will happen when add uint16 addend to uint16 augend.
func OverflowWhenAddUint16(augend, addend uint16) bool {
	return augend > 0 && addend > 0 && augend > MaxUint16-addend
}

// OverflowWhenAddUint32 checks whether overflow will happen when add uint32 addend to uint32 augend.
func OverflowWhenAddUint32(augend, addend uint32) bool {
	return augend > 0 && addend > 0 && augend > MaxUint32-addend
}

// OverflowWhenAddUint64 checks whether overflow will happen when add uint64 addend to uint64 augend.
func OverflowWhenAddUint64(augend, addend uint64) bool {
	return augend > 0 && addend > 0 && augend > MaxUint64-addend
}

// OverflowWhenAddUint checks whether overflow will happen when add uint addend to uint augend.
func OverflowWhenAddUint(augend, addend uint) bool {
	if intBitLength <= 32 {
		return OverflowWhenAddUint32(uint32(augend), uint32(addend))
	}
	return OverflowWhenAddUint64(uint64(augend), uint64(addend))
}

// OverflowWhenSubtractUint8 checks whether overflow will happen when subtract uint8 subtrahend from uint8 minuend.
func OverflowWhenSubtractUint8(minuend, subtrahend uint8) bool {
	return minuend > 0 && subtrahend > 0 && minuend < subtrahend
}

// OverflowWhenSubtractUint16 checks whether overflow will happen when subtract uint16 subtrahend from uint16 minuend.
func OverflowWhenSubtractUint16(minuend, subtrahend uint16) bool {
	return minuend > 0 && subtrahend > 0 && minuend < subtrahend
}

// OverflowWhenSubtractUint32 checks whether overflow will happen when subtract uint32 subtrahend from uint32 minuend.
func OverflowWhenSubtractUint32(minuend, subtrahend uint32) bool {
	return minuend > 0 && subtrahend > 0 && minuend < subtrahend
}

// OverflowWhenSubtractUint64 checks whether overflow will happen when subtract uint64 subtrahend from uint64 minuend.
func OverflowWhenSubtractUint64(minuend, subtrahend uint64) bool {
	return minuend > 0 && subtrahend > 0 && minuend < subtrahend
}

// OverflowWhenSubtractUint checks whether overflow will happen when subtract uint subtrahend from uint minuend.
func OverflowWhenSubtractUint(minuend, subtrahend uint) bool {
	if intBitLength <= 32 {
		return OverflowWhenSubtractUint32(uint32(minuend), uint32(subtrahend))
	}
	return OverflowWhenSubtractUint64(uint64(minuend), uint64(subtrahend))
}
