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
