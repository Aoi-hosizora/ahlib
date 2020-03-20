package xnumber

import (
	"fmt"
)

func RenderLatency(ns float64) string {
	if ns <= 0 {
		return "0.0000ns"
	}
	if ns < 1e3 {
		return fmt.Sprintf("%.4fns", ns)
	}
	us := ns / 1e3
	if us < 1e3 {
		return fmt.Sprintf("%.4fµs", us)
	}
	ms := us / 1e3
	if ms < 1e3 {
		return fmt.Sprintf("%.4fms", ms)
	}
	s := ms / 1e3
	if s < 60 {
		return fmt.Sprintf("%.4fs", s)
	}
	min := s / 60
	return fmt.Sprintf("%.4fmin", min)
}

func RenderByte(b float64) string {
	if b <= 0 {
		return "0B"
	}
	if b < 1024 {
		return fmt.Sprintf("%dB", int(b))
	}
	kb := b / 1024.0
	if kb < 1024 {
		return fmt.Sprintf("%.2fKB", kb)
	}
	mb := kb / 1024.0
	return fmt.Sprintf("%.2fMB", mb)
}
