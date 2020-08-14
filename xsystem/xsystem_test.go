package xsystem

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
)

func TestBitNumber(t *testing.T) {
	log.Println(BitNumber())
	log.Println(IsX64())
	log.Println(IsX86())

	bit := 32 << (^uint(0) >> 63)
	xtesting.Equal(t, IsX64(), bit == 64)
	xtesting.Equal(t, IsX86(), bit == 32)
}
