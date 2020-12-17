package xsystem

import (
	"runtime"
)

const intSize = 32 << (^uint(0) >> 63)

// BitNumber returns the os bit number (32 or 64) of the OS.
func BitNumber() int {
	return intSize
}

// Is32Bit returns true if the os bit number is 32.
func Is32Bit() bool {
	return intSize == 32
}

// Is64Bit returns true if the os bit number is 64.
func Is64Bit() bool {
	return intSize == 64
}

// GetOsName returns the running program's operating system target, see runtime.GOOS.
// 	$ go tool dist list
// 	aix/ppc64
// 	android/386
// 	android/amd64
// 	android/arm
// 	android/arm64
// 	darwin/amd64
// 	darwin/arm64
// 	dragonfly/amd64
// 	freebsd/386
// 	freebsd/amd64
// 	freebsd/arm
// 	freebsd/arm64
// 	illumos/amd64
// 	js/wasm
// 	linux/386
// 	linux/amd64
// 	linux/arm
// 	linux/arm64
// 	linux/mips
// 	linux/mips64
// 	linux/mips64le
// 	linux/mipsle
// 	linux/ppc64
// 	linux/ppc64le
// 	linux/riscv64
// 	linux/s390x
// 	netbsd/386
// 	netbsd/amd64
// 	netbsd/arm
// 	netbsd/arm64
// 	openbsd/386
// 	openbsd/amd64
// 	openbsd/arm
// 	openbsd/arm64
// 	plan9/386
// 	plan9/amd64
// 	plan9/arm
// 	solaris/amd64
// 	windows/386
// 	windows/amd64
// 	windows/arm
func GetOsName() string {
	return runtime.GOOS
}

// GetOsArch returns the running program's architecture target, see runtime.GOARCH.
func GetOsArch() string {
	return runtime.GOARCH
}

// GetTotalMemory returns the total memory size, see https://github.com/pbnjay/memory.
func GetTotalMemory() uint64 {
	return totalMemory()
}
