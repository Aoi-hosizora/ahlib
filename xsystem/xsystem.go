package xsystem

import (
	"runtime"
)

func BitNumber() int {
	return 32 << (^uint(0) >> 63)
}

func IsX86() bool {
	return BitNumber() == 32
}

func IsX64() bool {
	return BitNumber() == 64
}

func GetOsName() string {
	return runtime.GOOS
}

func GetOsArch() string {
	return runtime.GOARCH
}

func GetTotalMemory() uint64 {
	return sysTotalMemory()
}
