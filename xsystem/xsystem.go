package xsystem

func BitNumber() int {
	return 32 << (^uint(0) >> 63)
}

func IsX86() bool {
	return BitNumber() == 32
}

func IsX64() bool {
	return BitNumber() == 64
}
