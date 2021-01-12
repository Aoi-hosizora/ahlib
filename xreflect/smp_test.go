package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestSmpFlag(t *testing.T) {
	v := SmpvalOf(0)
	xtesting.Equal(t, v.Flag(), Int)
	v = SmpvalOf(uint(0))
	xtesting.Equal(t, v.Flag(), Uint)
	v = SmpvalOf(0.)
	xtesting.Equal(t, v.Flag(), Float)
	v = SmpvalOf(0 + 0i)
	xtesting.Equal(t, v.Flag(), Complex)
	v = SmpvalOf(true)
	xtesting.Equal(t, v.Flag(), Bool)
	v = SmpvalOf("")
	xtesting.Equal(t, v.Flag(), String)

	s := SmplenOf(0)
	xtesting.Equal(t, s.Flag(), Int)
	s = SmplenOf(uint(0))
	xtesting.Equal(t, s.Flag(), Uint)
	s = SmplenOf(0.)
	xtesting.Equal(t, s.Flag(), Float)
	s = SmplenOf(0 + 0i)
	xtesting.Equal(t, s.Flag(), Complex)
	s = SmplenOf(true)
	xtesting.Equal(t, s.Flag(), Bool)
}

func TestSmpval(t *testing.T) {
	i := 9223372036854775807
	i8 := int8(127)
	i16 := int16(32767)
	i32 := int32(2147483647)
	i64 := int64(9223372036854775807)
	u := uint(18446744073709551615)
	u8 := uint8(255)
	u16 := uint16(65535)
	u32 := uint32(4294967295)
	u64 := uint64(18446744073709551615)
	up := uintptr(18446744073709551615)
	f32 := float32(0.1)
	f64 := 0.1
	c64 := complex64(0.1 + 0.1i)
	c128 := 0.1 + 0.1i
	str1 := "test"
	str2 := "测试テスト"
	str3 := ""
	b1 := true
	b2 := false
	m1 := []int{0, 1, 2}
	m2 := [...]int{0, 1, 2}
	m3 := map[int]int{0: 0, 1: 1, 2: 2}
	s := struct{}{}
	p := &struct{}{}

	v := SmpvalOf(i)
	xtesting.Equal(t, v.Int(), int64(i))
	v = SmpvalOf(i8)
	xtesting.Equal(t, v.Int(), int64(i8))
	v = SmpvalOf(i16)
	xtesting.Equal(t, v.Int(), int64(i16))
	v = SmpvalOf(i32)
	xtesting.Equal(t, v.Int(), int64(i32))
	v = SmpvalOf(i64)
	xtesting.Equal(t, v.Int(), i64)

	v = SmpvalOf(u)
	xtesting.Equal(t, v.Uint(), uint64(u))
	v = SmpvalOf(u8)
	xtesting.Equal(t, v.Uint(), uint64(u8))
	v = SmpvalOf(u16)
	xtesting.Equal(t, v.Uint(), uint64(u16))
	v = SmpvalOf(u32)
	xtesting.Equal(t, v.Uint(), uint64(u32))
	v = SmpvalOf(u64)
	xtesting.Equal(t, v.Uint(), u64)
	v = SmpvalOf(up)
	xtesting.Equal(t, v.Uint(), uint64(up))

	v = SmpvalOf(f32)
	xtesting.InDelta(t, v.Float(), float64(f32), 1e-3)
	v = SmpvalOf(f64)
	xtesting.InDelta(t, v.Float(), f64, 1e-3)

	c := SmpvalOf(c64)
	xtesting.InDelta(t, real(c.Complex()), float64(real(c64)), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), float64(imag(c64)), 1e-3)
	c = SmpvalOf(c128)
	xtesting.InDelta(t, real(c.Complex()), real(c128), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), imag(c128), 1e-3)

	v = SmpvalOf(str1)
	xtesting.Equal(t, v.String(), str1)
	v = SmpvalOf(str2)
	xtesting.Equal(t, v.String(), str2)
	v = SmpvalOf(str3)
	xtesting.Equal(t, v.String(), str3)

	// noinspection GoBoolExpressions
	v = SmpvalOf(b1)
	xtesting.Equal(t, v.Bool(), true)
	// noinspection GoBoolExpressions
	v = SmpvalOf(b2)
	xtesting.Equal(t, v.Bool(), false)

	xtesting.Panic(t, func() { SmpvalOf(m1) })
	xtesting.Panic(t, func() { SmpvalOf(m2) })
	xtesting.Panic(t, func() { SmpvalOf(m3) })
	xtesting.Panic(t, func() { SmpvalOf(s) })
	xtesting.Panic(t, func() { SmpvalOf(p) })
}

func TestSmplen(t *testing.T) {
	i := 9223372036854775807
	i8 := int8(127)
	i16 := int16(32767)
	i32 := int32(2147483647)
	i64 := int64(9223372036854775807)
	u := uint(18446744073709551615)
	u8 := uint8(255)
	u16 := uint16(65535)
	u32 := uint32(4294967295)
	u64 := uint64(18446744073709551615)
	up := uintptr(18446744073709551615)
	f32 := float32(0.1)
	f64 := 0.1
	c64 := complex64(0.1 + 0.1i)
	c128 := 0.1 + 0.1i
	str1 := "test"
	str2 := "测试テスト"
	str3 := ""
	b1 := true
	b2 := false
	m1 := []int{0, 1, 2}
	m2 := [...]int{0, 1, 2}
	m3 := map[int]int{0: 0, 1: 1, 2: 2}
	s := struct{}{}
	p := &struct{}{}

	sze := SmplenOf(i)
	xtesting.Equal(t, sze.Int(), int64(i))
	sze = SmplenOf(i8)
	xtesting.Equal(t, sze.Int(), int64(i8))
	sze = SmplenOf(i16)
	xtesting.Equal(t, sze.Int(), int64(i16))
	sze = SmplenOf(i32)
	xtesting.Equal(t, sze.Int(), int64(i32))
	sze = SmplenOf(i64)
	xtesting.Equal(t, sze.Int(), i64)

	sze = SmplenOf(u)
	xtesting.Equal(t, sze.Uint(), uint64(u))
	sze = SmplenOf(u8)
	xtesting.Equal(t, sze.Uint(), uint64(u8))
	sze = SmplenOf(u16)
	xtesting.Equal(t, sze.Uint(), uint64(u16))
	sze = SmplenOf(u32)
	xtesting.Equal(t, sze.Uint(), uint64(u32))
	sze = SmplenOf(u64)
	xtesting.Equal(t, sze.Uint(), u64)
	sze = SmplenOf(up)
	xtesting.Equal(t, sze.Uint(), uint64(up))

	sze = SmplenOf(f32)
	xtesting.InDelta(t, sze.Float(), float64(f32), 1e-3)
	sze = SmplenOf(f32)
	xtesting.InDelta(t, sze.Float(), f64, 1e-3)

	c := SmplenOf(c64)
	xtesting.InDelta(t, real(c.Complex()), float64(real(c64)), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), float64(imag(c64)), 1e-3)
	c = SmplenOf(c128)
	xtesting.InDelta(t, real(c.Complex()), real(c128), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), imag(c128), 1e-3)

	sze = SmplenOf(str1)
	xtesting.Equal(t, sze.Int(), int64(4))
	sze = SmplenOf(str2)
	xtesting.Equal(t, sze.Int(), int64(5))
	sze = SmplenOf(str3)
	xtesting.Equal(t, sze.Int(), int64(0))

	// noinspection GoBoolExpressions
	sze = SmplenOf(b1)
	xtesting.Equal(t, sze.Bool(), true)
	// noinspection GoBoolExpressions
	sze = SmplenOf(b2)
	xtesting.Equal(t, sze.Bool(), false)

	sze = SmplenOf(m1)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze = SmplenOf(m2)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze = SmplenOf(m3)
	xtesting.Equal(t, sze.Int(), int64(3))

	xtesting.Panic(t, func() { SmplenOf(s) })
	xtesting.Panic(t, func() { SmplenOf(p) })
}
