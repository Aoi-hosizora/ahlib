// +build !linux,!darwin,!windows

// Copyright (c) 2017, Jeremy Jay
// All rights reserved.
// https://github.com/pbnjay/memory

package memory

func totalMemory() uint64 {
	return 0
}
