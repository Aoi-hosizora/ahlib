package xnumber

import (
	"fmt"
	"math"
	_ "runtime"
	_ "unsafe"
)

// Accuracy represents an accuracy with some compare methods in accuracy.
type Accuracy func() float64

// NewAccuracy creates an Accuracy, using eps as its accuracy.
func NewAccuracy(eps float64) Accuracy {
	return func() float64 {
		return eps
	}
}

// Equal checks eq between two float64.
func (eps Accuracy) Equal(a, b float64) bool {
	return math.Abs(a-b) < eps()
}

// NotEqual checks ne between two float64.
func (eps Accuracy) NotEqual(a, b float64) bool {
	return math.Abs(a-b) >= eps()
}

// Greater checks gt between two float64.
func (eps Accuracy) Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > eps()
}

// Less checks lt between two float64.
func (eps Accuracy) Less(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > eps()
}

// GreaterOrEqual checks gte between two float64.
func (eps Accuracy) GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < eps()
}

// LessOrEqual checks lte between two float64.
func (eps Accuracy) LessOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < eps()
}

// _acc represents a default Accuracy with 1e-3 as default accuracy.
var _acc = NewAccuracy(1e-3)

// EqualInAccuracy checks eq between two float64 in default Accuracy: 1e-3.
func EqualInAccuracy(a, b float64) bool {
	return _acc.Equal(a, b)
}

// NotEqualInAccuracy checks ne between two float64 in default Accuracy: 1e-3.
func NotEqualInAccuracy(a, b float64) bool {
	return _acc.NotEqual(a, b)
}

// GreaterInAccuracy checks gt between two float64 in default Accuracy: 1e-3.
func GreaterInAccuracy(a, b float64) bool {
	return _acc.Greater(a, b)
}

// LessInAccuracy checks lt between two float64 in default Accuracy: 1e-3.
func LessInAccuracy(a, b float64) bool {
	return _acc.Less(a, b)
}

// GreaterOrEqualInAccuracy checks gte between two float64 in default Accuracy: 1e-3.
func GreaterOrEqualInAccuracy(a, b float64) bool {
	return _acc.GreaterOrEqual(a, b)
}

// LessOrEqualInAccuracy checks lte between two float64 in default Accuracy: 1e-3.
func LessOrEqualInAccuracy(a, b float64) bool {
	return _acc.LessOrEqual(a, b)
}

// RenderByte renders a byte size to string (using %.2f), support `B` `KB` `MB` `GB` `TB`.
func RenderByte(bytes float64) string {
	divider := float64(1024)

	minus := false
	if bytes < 0 {
		bytes = -bytes
		minus = true
	} else if bytes == 0 {
		return "0B"
	}
	ret := func(s string) string {
		if minus {
			return fmt.Sprintf("-%s", s)
		}
		return s
	}

	// 1 - 1023B
	b := bytes
	if LessInAccuracy(b, divider) {
		return ret(fmt.Sprintf("%dB", int(b)))
	}

	// 1 - 1023K
	kb := bytes / divider
	if LessInAccuracy(kb, divider) {
		return ret(fmt.Sprintf("%.2fKB", kb))
	}

	// 1 - 1023M
	mb := kb / divider
	if LessInAccuracy(mb, divider) {
		return ret(fmt.Sprintf("%.2fMB", mb))
	}

	// 1 - 1023G
	gb := mb / divider
	if LessInAccuracy(gb, divider) {
		return ret(fmt.Sprintf("%.2fGB", gb))
	}

	// 1T -
	tb := gb / divider
	return ret(fmt.Sprintf("%.2fTB", tb))
}

// Bool returns 1 if value is true, otherwise returns 0.
func Bool(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntSize returns the int size (32 / 64).
func IntSize() int {
	const intSize = 32 << (^uint(0) >> 63)
	return intSize
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
