package xsystem

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"runtime"
	"testing"
)

func TestBitNumber(t *testing.T) {
	log.Println(BitNumber())
	log.Println(IsX64())
	log.Println(IsX86())

	bit := 32 << (^uint(0) >> 63)
	xtesting.Equal(t, BitNumber(), bit)
	xtesting.Equal(t, IsX64(), bit == 64)
	xtesting.Equal(t, IsX86(), bit == 32)
}

func TestOs(t *testing.T) {
	log.Println(GetOsName())
	log.Println(GetOsArch())
	log.Println(GetTotalMemory())
	log.Println(GetTotalMemory() / 1024.0 / 1024 / 1024)

	xtesting.Equal(t, GetOsName(), runtime.GOOS)
	xtesting.Equal(t, GetOsArch(), runtime.GOARCH)
	xtesting.Equal(t, GetTotalMemory(), sysTotalMemory())
}
