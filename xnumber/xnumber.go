package xnumber

import (
	"fmt"
	"math"
)

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
