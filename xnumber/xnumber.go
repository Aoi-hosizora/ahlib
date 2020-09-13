package xnumber

import (
	"fmt"
	"math"
	"strconv"
)

// accuracy

type Accuracy func() float64

// Default accuracy, use 1e-3.
var DefaultAccuracy = NewAccuracy(1e-3)

func NewAccuracy(eps float64) Accuracy {
	return func() float64 {
		return eps
	}
}

func (eps Accuracy) Equal(a, b float64) bool {
	return math.Abs(a-b) < eps()
}

func (eps Accuracy) NotEqual(a, b float64) bool {
	return !eps.Equal(a, b)
}

func (eps Accuracy) Greater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > eps()
}

func (eps Accuracy) Smaller(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > eps()
}

func (eps Accuracy) GreaterOrEqual(a, b float64) bool {
	return math.Max(a, b) == a || math.Abs(a-b) < eps()
}

func (eps Accuracy) SmallerOrEqual(a, b float64) bool {
	return math.Max(a, b) == b || math.Abs(a-b) < eps()
}

// render

func RenderByte(b float64) string {
	if DefaultAccuracy.SmallerOrEqual(b, 0) {
		return "0B"
	}
	if DefaultAccuracy.Smaller(b, 1024) {
		return fmt.Sprintf("%dB", int(b))
	}
	kb := b / 1024.0
	if DefaultAccuracy.Smaller(kb, 1024) {
		return fmt.Sprintf("%.2fKB", kb)
	}
	mb := kb / 1024.0
	if DefaultAccuracy.Smaller(mb, 1024) {
		return fmt.Sprintf("%.2fMB", mb)
	}
	gb := mb / 1024.0
	if DefaultAccuracy.Smaller(gb, 1024) {
		return fmt.Sprintf("%.2fGB", gb)
	}
	tb := gb / 1024.0
	return fmt.Sprintf("%.2fTB", tb)
}

func Bool(b bool) int {
	if b {
		return 1
	}
	return 0
}

// parse

func ParseInt(s string, base int) (int, error) {
	i, e := strconv.ParseInt(s, base, 0)
	return int(i), e
}

func ParseInt8(s string, base int) (int8, error) {
	i, e := strconv.ParseInt(s, base, 8)
	return int8(i), e
}

func ParseInt16(s string, base int) (int16, error) {
	i, e := strconv.ParseInt(s, base, 16)
	return int16(i), e
}

func ParseInt32(s string, base int) (int32, error) {
	i, e := strconv.ParseInt(s, base, 32)
	return int32(i), e
}

func ParseInt64(s string, base int) (int64, error) {
	i, e := strconv.ParseInt(s, base, 64)
	return i, e
}

func ParseUint(s string, base int) (uint, error) {
	i, e := strconv.ParseUint(s, base, 0)
	return uint(i), e
}

func ParseUint8(s string, base int) (uint8, error) {
	i, e := strconv.ParseUint(s, base, 8)
	return uint8(i), e
}

func ParseUint16(s string, base int) (uint16, error) {
	i, e := strconv.ParseUint(s, base, 16)
	return uint16(i), e
}

func ParseUint32(s string, base int) (uint32, error) {
	i, e := strconv.ParseUint(s, base, 32)
	return uint32(i), e
}

func ParseUint64(s string, base int) (uint64, error) {
	i, e := strconv.ParseUint(s, base, 64)
	return i, e
}

func ParseFloat32(s string) (float32, error) {
	f, e := strconv.ParseFloat(s, 32)
	return float32(f), e
}

func ParseFloat64(s string) (float64, error) {
	f, e := strconv.ParseFloat(s, 64)
	return f, e
}

// atoX

func Atoi(s string) (int, error) {
	return ParseInt(s, 10)
}

func Atoi8(s string) (int8, error) {
	return ParseInt8(s, 10)
}

func Atoi16(s string) (int16, error) {
	return ParseInt16(s, 10)
}

func Atoi32(s string) (int32, error) {
	return ParseInt32(s, 10)
}

func Atoi64(s string) (int64, error) {
	return ParseInt64(s, 10)
}

func Atou(s string) (uint, error) {
	return ParseUint(s, 10)
}

func Atou8(s string) (uint8, error) {
	return ParseUint8(s, 10)
}

func Atou16(s string) (uint16, error) {
	return ParseUint16(s, 10)
}

func Atou32(s string) (uint32, error) {
	return ParseUint32(s, 10)
}

func Atou64(s string) (uint64, error) {
	return ParseUint64(s, 10)
}

func Atof32(s string) (float32, error) {
	return ParseFloat32(s)
}

func Atof64(s string) (float64, error) {
	return ParseFloat64(s)
}

// format

func FormatInt(i int, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatInt8(i int8, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatInt16(i int16, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatInt32(i int32, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatInt64(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

func FormatUint(i uint, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatUint8(i uint8, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatUint16(i uint16, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatUint32(i uint32, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatUint64(i uint64, base int) string {
	return strconv.FormatUint(i, base)
}

func FormatFloat32(f float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, 32)
}

func FormatFloat64(f float64, fmt byte, prec int) string {
	return strconv.FormatFloat(f, fmt, prec, 64)
}

// Xtoa

func Itoa(i int) string {
	return FormatInt(i, 10)
}

func I8toa(i int8) string {
	return FormatInt8(i, 10)
}

func I16toa(i int16) string {
	return FormatInt16(i, 10)
}

func I32toa(i int32) string {
	return FormatInt32(i, 10)
}

func I64toa(i int64) string {
	return FormatInt64(i, 10)
}

func Utoa(i uint) string {
	return FormatUint(i, 10)
}

func U8toa(i uint8) string {
	return FormatUint8(i, 10)
}

func U16toa(i uint16) string {
	return FormatUint16(i, 10)
}

func U32toa(i uint32) string {
	return FormatUint32(i, 10)
}

func U64toa(i uint64) string {
	return FormatUint64(i, 10)
}

func F32toa(f float32) string {
	return FormatFloat32(f, 'f', -1)
}

func F64toa(f float64) string {
	return FormatFloat64(f, 'f', -1)
}

// min max

const (
	MinInt8   = int8(-128)                  // -1 << 7
	MinInt16  = int16(-32768)               // -1 << 15
	MinInt32  = int32(-2147483648)          // -1 << 31
	MinInt64  = int64(-9223372036854775808) // -1 << 63
	MinUint8  = uint8(0)                    // 0
	MinUint16 = uint16(0)                   // 0
	MinUint32 = uint32(0)                   // 0
	MinUint64 = uint64(0)                   // 0

	MaxInt8   = int8(127)                    // 1<<7 - 1
	MaxInt16  = int16(32767)                 //  1<<15 - 1
	MaxInt32  = int32(2147483647)            //  1<<31 - 1
	MaxInt64  = int64(9223372036854775807)   // 1<<63 - 1
	MaxUint8  = uint8(255)                   //  1<<8 - 1
	MaxUint16 = uint16(65535)                // 1<<16 - 1
	MaxUint32 = uint32(4294967295)           // 1<<32 - 1
	MaxUint64 = uint64(18446744073709551615) // 1<<64 - 1

	MaxFloat32             = float32(math.MaxFloat32)             // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = float32(math.SmallestNonzeroFloat32) // 1 / 2**(127 - 1 + 23)
	MaxFloat64             = float64(math.MaxFloat64)             // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = float64(math.SmallestNonzeroFloat64) // 1 / 2**(1023 - 1 + 52)
)
