package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"testing"
)

func TestElemType(t *testing.T) {
	var a ****int
	t1 := ElemType(a).String()
	xtesting.Equal(t, t1, "int")

	var b int
	t2 := ElemType(b).String()
	xtesting.Equal(t, t2, "int")
}

func TestElemValue(t *testing.T) {
	var a *****int
	v1 := ElemValue(a)
	xtesting.Equal(t, v1.IsValid(), false)

	var b int
	v2 := ElemValue(b).Interface()
	xtesting.Equal(t, v2, 0)

	var c = &b
	v3 := ElemValue(c).Interface()
	xtesting.Equal(t, v3, 0)
}

func TestUnexportedField(t *testing.T) {
	type s struct {
		a string
		b int64
		c uint64
		d float64
	}
	ss := &s{}
	el := reflect.ValueOf(ss).Elem()

	xtesting.Equal(t, GetUnexportedField(el.FieldByName("a")), "")
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("b")), int64(0))
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("c")), uint64(0))
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("d")), 0.0)

	SetUnexportedField(el.FieldByName("a"), "string")
	SetUnexportedField(el.FieldByName("b"), int64(9223372036854775807))
	SetUnexportedField(el.FieldByName("c"), uint64(18446744073709551615))
	SetUnexportedField(el.FieldByName("d"), 0.333)

	xtesting.Equal(t, GetUnexportedField(el.FieldByName("a")), "string")
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("b")), int64(9223372036854775807))
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("c")), uint64(18446744073709551615))
	xtesting.Equal(t, GetUnexportedField(el.FieldByName("d")), 0.333)
}

func TestGetStructFields(t *testing.T) {
	a := struct {
		A int
		B string
		C float64
	}{}
	fields := GetStructFields(a)

	xtesting.Equal(t, len(fields), 3)
	xtesting.True(t, fields[0].Name == "A" && fields[0].Type == reflect.TypeOf(0))
	xtesting.True(t, fields[1].Name == "B" && fields[1].Type == reflect.TypeOf(""))
	xtesting.True(t, fields[2].Name == "C" && fields[2].Type == reflect.TypeOf(0.))
}

func TestGet(t *testing.T) {
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

	ii, _ := GetInt(i)
	xtesting.Equal(t, ii, int64(i))
	ii, _ = GetInt(i8)
	xtesting.Equal(t, ii, int64(i8))
	ii, _ = GetInt(i16)
	xtesting.Equal(t, ii, int64(i16))
	ii, _ = GetInt(i32)
	xtesting.Equal(t, ii, int64(i32))
	ii, _ = GetInt(i64)
	xtesting.Equal(t, ii, i64)
	_, err := GetInt("")
	xtesting.NotNil(t, err)

	uu, _ := GetUint(u)
	xtesting.Equal(t, uu, uint64(u))
	uu, _ = GetUint(u8)
	xtesting.Equal(t, uu, uint64(u8))
	uu, _ = GetUint(u16)
	xtesting.Equal(t, uu, uint64(u16))
	uu, _ = GetUint(u32)
	xtesting.Equal(t, uu, uint64(u32))
	uu, _ = GetUint(u64)
	xtesting.Equal(t, uu, u64)
	uu, _ = GetUint(up)
	xtesting.Equal(t, uu, uint64(up))
	_, err = GetUint("")
	xtesting.NotNil(t, err)

	ff, _ := GetFloat(f32)
	xtesting.InDelta(t, ff, 0.1, 1e-3)
	ff, _ = GetFloat(f64)
	xtesting.InDelta(t, ff, 0.1, 1e-3)
	_, err = GetFloat("")
	xtesting.NotNil(t, err)

	cc, _ := GetComplex(c64)
	xtesting.InDelta(t, real(cc), 0.1, 1e-3)
	xtesting.InDelta(t, imag(cc), 0.1, 1e-3)
	cc, _ = GetComplex(c128)
	xtesting.InDelta(t, real(cc), 0.1, 1e-3)
	xtesting.InDelta(t, imag(cc), 0.1, 1e-3)
	_, err = GetComplex("")
	xtesting.NotNil(t, err)

	ss, _ := GetString(str1)
	xtesting.Equal(t, ss, str1)
	ss, _ = GetString(str2)
	xtesting.Equal(t, ss, str2)
	ss, _ = GetString(str3)
	xtesting.Equal(t, ss, str3)
	_, err = GetString(0)
	xtesting.NotNil(t, err)

	bb, _ := GetBool(true)
	xtesting.Equal(t, bb, true)
	_, err = GetBool(0)
	xtesting.NotNil(t, err)
	bb, _ = GetBool(false)
	xtesting.Equal(t, bb, false)
}

func TestFlag(t *testing.T) {
	v, _ := SmpvalOf(0)
	xtesting.Equal(t, v.Flag(), Int)
	v, _ = SmpvalOf(uint(0))
	xtesting.Equal(t, v.Flag(), Uint)
	v, _ = SmpvalOf(0.)
	xtesting.Equal(t, v.Flag(), Float)
	v, _ = SmpvalOf(0 + 0i)
	xtesting.Equal(t, v.Flag(), Complex)
	v, _ = SmpvalOf(true)
	xtesting.Equal(t, v.Flag(), Bool)
	v, _ = SmpvalOf("")
	xtesting.Equal(t, v.Flag(), String)

	s, _ := SmplenOf(0)
	xtesting.Equal(t, s.Flag(), Int)
	s, _ = SmplenOf(uint(0))
	xtesting.Equal(t, s.Flag(), Uint)
	s, _ = SmplenOf(0.)
	xtesting.Equal(t, s.Flag(), Float)
	s, _ = SmplenOf(0 + 0i)
	xtesting.Equal(t, s.Flag(), Complex)
	s, _ = SmplenOf(true)
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

	v, _ := SmpvalOf(i)
	xtesting.Equal(t, v.Int(), int64(i))
	v, _ = SmpvalOf(i8)
	xtesting.Equal(t, v.Int(), int64(i8))
	v, _ = SmpvalOf(i16)
	xtesting.Equal(t, v.Int(), int64(i16))
	v, _ = SmpvalOf(i32)
	xtesting.Equal(t, v.Int(), int64(i32))
	v, _ = SmpvalOf(i64)
	xtesting.Equal(t, v.Int(), i64)

	v, _ = SmpvalOf(u)
	xtesting.Equal(t, v.Uint(), uint64(u))
	v, _ = SmpvalOf(u8)
	xtesting.Equal(t, v.Uint(), uint64(u8))
	v, _ = SmpvalOf(u16)
	xtesting.Equal(t, v.Uint(), uint64(u16))
	v, _ = SmpvalOf(u32)
	xtesting.Equal(t, v.Uint(), uint64(u32))
	v, _ = SmpvalOf(u64)
	xtesting.Equal(t, v.Uint(), u64)
	v, _ = SmpvalOf(up)
	xtesting.Equal(t, v.Uint(), uint64(up))

	v, _ = SmpvalOf(f32)
	xtesting.InDelta(t, v.Float(), float64(f32), 1e-3)
	v, _ = SmpvalOf(f64)
	xtesting.InDelta(t, v.Float(), f64, 1e-3)

	c, _ := SmpvalOf(c64)
	xtesting.InDelta(t, real(c.Complex()), float64(real(c64)), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), float64(imag(c64)), 1e-3)
	c, _ = SmpvalOf(c128)
	xtesting.InDelta(t, real(c.Complex()), real(c128), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), imag(c128), 1e-3)

	v, _ = SmpvalOf(str1)
	xtesting.Equal(t, v.String(), str1)
	v, _ = SmpvalOf(str2)
	xtesting.Equal(t, v.String(), str2)
	v, _ = SmpvalOf(str3)
	xtesting.Equal(t, v.String(), str3)

	// noinspection GoBoolExpressions
	v, _ = SmpvalOf(b1)
	xtesting.Equal(t, v.Bool(), true)
	// noinspection GoBoolExpressions
	v, _ = SmpvalOf(b2)
	xtesting.Equal(t, v.Bool(),false)

	_, err := SmpvalOf(m1)
	xtesting.NotNil(t, err)
	_, err = SmpvalOf(m2)
	xtesting.NotNil(t, err)
	_, err = SmpvalOf(m3)
	xtesting.NotNil(t, err)
	_, err = SmpvalOf(s)
	xtesting.NotNil(t, err)
	_, err = SmpvalOf(p)
	xtesting.NotNil(t, err)
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

	sze, _ := SmplenOf(i)
	xtesting.Equal(t, sze.Int(), int64(i))
	sze, _ = SmplenOf(i8)
	xtesting.Equal(t, sze.Int(), int64(i8))
	sze, _ = SmplenOf(i16)
	xtesting.Equal(t, sze.Int(), int64(i16))
	sze, _ = SmplenOf(i32)
	xtesting.Equal(t, sze.Int(), int64(i32))
	sze, _ = SmplenOf(i64)
	xtesting.Equal(t, sze.Int(), i64)

	sze, _ = SmplenOf(u)
	xtesting.Equal(t, sze.Uint(), uint64(u))
	sze, _ = SmplenOf(u8)
	xtesting.Equal(t, sze.Uint(), uint64(u8))
	sze, _ = SmplenOf(u16)
	xtesting.Equal(t, sze.Uint(), uint64(u16))
	sze, _ = SmplenOf(u32)
	xtesting.Equal(t, sze.Uint(), uint64(u32))
	sze, _ = SmplenOf(u64)
	xtesting.Equal(t, sze.Uint(), u64)
	sze, _ = SmplenOf(up)
	xtesting.Equal(t, sze.Uint(), uint64(up))

	sze, _ = SmplenOf(f32)
	xtesting.InDelta(t, sze.Float(), float64(f32), 1e-3)
	sze, _ = SmplenOf(f32)
	xtesting.InDelta(t, sze.Float(), f64, 1e-3)

	c, _ := SmplenOf(c64)
	xtesting.InDelta(t, real(c.Complex()), float64(real(c64)), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), float64(imag(c64)), 1e-3)
	c, _ = SmplenOf(c128)
	xtesting.InDelta(t, real(c.Complex()), real(c128), 1e-3)
	xtesting.InDelta(t, imag(c.Complex()), imag(c128), 1e-3)

	sze, _ = SmplenOf(str1)
	xtesting.Equal(t, sze.Int(), int64(4))
	sze, _ = SmplenOf(str2)
	xtesting.Equal(t, sze.Int(), int64(5))
	sze, _ = SmplenOf(str3)
	xtesting.Equal(t, sze.Int(), int64(0))

	// noinspection GoBoolExpressions
	sze, _ = SmplenOf(b1)
	xtesting.Equal(t, sze.Bool(), true)
	// noinspection GoBoolExpressions
	sze, _ = SmplenOf(b2)
	xtesting.Equal(t, sze.Bool(), false)

	sze, _ = SmplenOf(m1)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze, _ = SmplenOf(m2)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze, _ = SmplenOf(m3)
	xtesting.Equal(t, sze.Int(), int64(3))

	_, err := SmplenOf(s)
	xtesting.NotNil(t, err)
	_, err = SmplenOf(p)
	xtesting.NotNil(t, err)
}
