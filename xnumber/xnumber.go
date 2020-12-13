package xnumber

import (
	"fmt"
	"math"
)

// Accuracy represents a type of function, which includes some compare functions using accuracy.
type Accuracy func() float64

// NewAccuracy creates an Accuracy, using eps as accuracy.
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
	return !eps.Equal(a, b)
}

// Greater checks gt between two float64.
func (eps Accuracy) Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > eps()
}

// Smaller checks lt between two float64.
func (eps Accuracy) Smaller(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > eps()
}

// GreaterOrEqual checks gte between two float64.
func (eps Accuracy) GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < eps()
}

// SmallerOrEqual checks lte between two float64.
func (eps Accuracy) SmallerOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < eps()
}

// _acc represents a default Accuracy with 1e-3 as default accuracy.
var _acc = NewAccuracy(1e-3)

// EqualInAccuracy checks eq between two float64 in default Accuracy.
func EqualInAccuracy(a, b float64) bool {
	return _acc.Equal(a, b)
}

// NotEqualInAccuracy checks ne between two float64 in default Accuracy.
func NotEqualInAccuracy(a, b float64) bool {
	return _acc.NotEqual(a, b)
}

// GreaterInAccuracy checks gt between two float64 in default Accuracy.
func GreaterInAccuracy(a, b float64) bool {
	return _acc.Greater(a, b)
}

// SmallerInAccuracy checks lt between two float64 in default Accuracy.
func SmallerInAccuracy(a, b float64) bool {
	return _acc.Smaller(a, b)
}

// GreaterOrEqualInAccuracy checks gte between two float64 in default Accuracy.
func GreaterOrEqualInAccuracy(a, b float64) bool {
	return _acc.GreaterOrEqual(a, b)
}

// SmallerOrEqualInAccuracy checks lte between two float64 in default Accuracy.
func SmallerOrEqualInAccuracy(a, b float64) bool {
	return _acc.SmallerOrEqual(a, b)
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
	if SmallerInAccuracy(b, divider) {
		return ret(fmt.Sprintf("%dB", int(b)))
	}

	// 1 - 1023K
	kb := bytes / divider
	if SmallerInAccuracy(kb, divider) {
		return ret(fmt.Sprintf("%.2fKB", kb))
	}

	// 1 - 1023M
	mb := kb / divider
	if SmallerInAccuracy(mb, divider) {
		return ret(fmt.Sprintf("%.2fMB", mb))
	}

	// 1 - 1023G
	gb := mb / divider
	if SmallerInAccuracy(gb, divider) {
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
