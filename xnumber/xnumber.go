package xnumber

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xsystem"
	"math"
	"strconv"
)

// accuracy

type Accuracy func() float64

func NewAccuracy(eps float64) Accuracy {
	return func() float64 {
		return eps
	}
}

func (eps Accuracy) Equal(a, b float64) bool {
	return math.Abs(a-b) < eps()
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

// Deprecated: Use `time.Duration.String()` is better
func RenderLatency(ns float64) string {
	acc := NewAccuracy(1e-3)
	if acc.SmallerOrEqual(ns, 0) {
		return "0.0000ns"
	}
	if acc.Smaller(ns, 1e3) {
		return fmt.Sprintf("%.4fns", ns)
	}
	us := ns / 1e3
	if acc.Smaller(us, 1e3) {
		return fmt.Sprintf("%.4fµs", us)
	}
	ms := us / 1e3
	if acc.Smaller(ms, 1e3) {
		return fmt.Sprintf("%.4fms", ms)
	}
	s := ms / 1e3
	if acc.Smaller(s, 60) {
		return fmt.Sprintf("%.4fs", s)
	}
	m := s / 60
	return fmt.Sprintf("%.4fmin", m)
}

func RenderByte(b float64) string {
	acc := NewAccuracy(1e-3)
	if acc.SmallerOrEqual(b, 0) {
		return "0B"
	}
	if acc.Smaller(b, 1024) {
		return fmt.Sprintf("%dB", int(b))
	}
	kb := b / 1024.0
	if acc.Smaller(kb, 1024) {
		return fmt.Sprintf("%.2fKB", kb)
	}
	mb := kb / 1024.0
	return fmt.Sprintf("%.2fMB", mb)
}

// parse

func ParseInt(s string, base int) (int, error) {
	i, e := strconv.ParseInt(s, base, xsystem.BitNumber())
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
	i, e := strconv.ParseUint(s, base, xsystem.BitNumber())
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
