// +build freebsd openbsd dragonfly netbsd

// Copyright (c) 2017, Jeremy Jay
// All rights reserved.
// https://github.com/pbnjay/memory

package xsystem

func sysTotalMemory() uint64 {
	s, err := sysctlUint64("hw.physmem")
	if err != nil {
		return 0
	}
	return s
}
