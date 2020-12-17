// +build linux

// Copyright (c) 2017, Jeremy Jay
// All rights reserved.
// https://github.com/pbnjay/memory

package xsystem

import "syscall"

func totalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	_ = syscall.Sysinfo(in)

	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}
