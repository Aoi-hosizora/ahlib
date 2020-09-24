package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"math"
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

func TestBoolVal(t *testing.T) {
	xtesting.Equal(t, BoolVal(true), 1)
	xtesting.Equal(t, BoolVal(false), 0)
}

func TestIsEqual(t *testing.T) {
	a := interface{}(1)
	b := 1
	c := int32(1)
	d := &b

	if !IsEqual(a, b) {
		t.Fatal("a and b is equal, but got not equal")
	}
	if IsEqual(a, c) {
		t.Fatal("a and c is not equal, but got equal")
	}
	if IsEqual(b, c) {
		t.Fatal("b and c is not equal, but got equal")
	}
	if !IsEqual(d, b) {
		t.Fatal("d and b is equal, but got not equal")
	}

	p1 := interface{}(nil)
	p2 := interface{}(nil)
	var p3 *int = nil
	p4 := &b
	var p5 interface{} = &b
	if !IsEqual(p1, p2) {
		t.Fatal("p1 and p2 is equal, but got not equal")
	}
	if !IsEqual(p2, p3) {
		t.Fatal("p2 and p3 is equal, nut got not equal")
	}
	if !IsEqual(p4, p5) {
		t.Fatal("p4 and p5 is equal, nut got not equal")
	}

	var s0 []string
	var s00 []string
	if !IsEqual(s0, s00) {
		t.Fatal("s0 and s00 is equal, but got not equal")
	}

	s1 := []int{1, 2, 3}
	s2 := []int{3, 2, 1}
	s3 := []interface{}{1, 2, 3}
	if IsEqual(s1, s2) {
		t.Fatal("s1 and s2 is not equal, but got equal")
	}
	if IsEqual(s1, s3) {
		t.Fatal("s1 and s3 is not equal, but got equal")
	}

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{3, 2, 1}
	a3 := [4]int{3, 2, 1}
	if IsEqual(a1, a2) {
		t.Fatal("a1 and a2 is not equal, but got equal")
	}
	if IsEqual(a2, a3) {
		t.Fatal("a2 and a3 is not equal, but got equal")
	}

	m0 := map[int]int{}
	m00 := map[int]int{}
	if !IsEqual(m0, m00) {
		t.Fatal("m0 and m00 is equal, but got not equal")
	}

	m1 := map[int]int{1: 1, 2: 2}
	m2 := map[int]int{2: 2, 1: 1}
	m3 := map[int]interface{}{2: 2, 1: 1}
	if !IsEqual(m1, m2) {
		t.Fatal("m1 and m2 is equal, but got equal")
	}
	if IsEqual(m1, m3) {
		t.Fatal("m1 and m3 is not equal, but got equal")
	}

	f0 := func() {}
	f00 := func() {}
	if IsEqual(f0, f00) {
		t.Fatal("f0 and f00 is not equal, but got not equal")
	}

	f1 := func() {}
	f2 := func(int) {}
	if IsEqual(f1, f1) {
		// Func values are deeply equal if both are nil; otherwise they are not deeply equal.
		t.Fatal("f1 and f1 is not equal, but got equal")
	}
	if IsEqual(f1, f2) {
		t.Fatal("f1 and f2 is not equal, but got equal")
	}
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
	str1 := "test"
	str2 := "测试テスト"
	str3 := ""
	b1 := true
	b2 := false

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

	ff, _ := GetFloat(f32)
	xtesting.True(t, math.Abs(ff-0.1) < 1e-3)
	ff, _ = GetFloat(f64)
	xtesting.True(t, math.Abs(ff-0.1) < 1e-3)

	ss, _ := GetString(str1)
	xtesting.Equal(t, ss, str1)
	ss, _ = GetString(str2)
	xtesting.Equal(t, ss, str2)
	ss, _ = GetString(str3)
	xtesting.Equal(t, ss, str3)

	// noinspection GoBoolExpressions
	bb, _ := GetBool(b1)
	// noinspection GoBoolExpressions
	xtesting.Equal(t, bb, b1)
	// noinspection GoBoolExpressions
	bb, _ = GetBool(b2)
	// noinspection GoBoolExpressions
	xtesting.Equal(t, bb, b2)
}

func TestFlag(t *testing.T) {
	i := 0
	u := uint(0)
	s := ""
	b := false

	v, _ := IufsOf(i)
	xtesting.Equal(t, v.Flag(), Int)
	v, _ = IufsOf(u)
	xtesting.Equal(t, v.Flag(), Uint)
	v, _ = IufsOf(s)
	xtesting.Equal(t, v.Flag(), String)
	// noinspection GoBoolExpressions
	v, _ = IufsOf(b)
	xtesting.Equal(t, v.Flag(), Int)
}

func TestIufs(t *testing.T) {
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

	v, _ := IufsOf(i)
	xtesting.Equal(t, v.Int(), int64(i))
	v, _ = IufsOf(i8)
	xtesting.Equal(t, v.Int(), int64(i8))
	v, _ = IufsOf(i16)
	xtesting.Equal(t, v.Int(), int64(i16))
	v, _ = IufsOf(i32)
	xtesting.Equal(t, v.Int(), int64(i32))
	v, _ = IufsOf(i64)
	xtesting.Equal(t, v.Int(), i64)

	v, _ = IufsOf(u)
	xtesting.Equal(t, v.Uint(), uint64(u))
	v, _ = IufsOf(u8)
	xtesting.Equal(t, v.Uint(), uint64(u8))
	v, _ = IufsOf(u16)
	xtesting.Equal(t, v.Uint(), uint64(u16))
	v, _ = IufsOf(u32)
	xtesting.Equal(t, v.Uint(), uint64(u32))
	v, _ = IufsOf(u64)
	xtesting.Equal(t, v.Uint(), u64)
	v, _ = IufsOf(up)
	xtesting.Equal(t, v.Uint(), uint64(up))

	v, _ = IufsOf(f32)
	xtesting.True(t, math.Abs(v.Float()-float64(f32)) < 1e-3)
	v, _ = IufsOf(f64)
	xtesting.True(t, math.Abs(v.Float()-f64) < 1e-3)

	v, _ = IufsOf(str1)
	xtesting.Equal(t, v.String(), str1)
	v, _ = IufsOf(str2)
	xtesting.Equal(t, v.String(), str2)
	v, _ = IufsOf(str3)
	xtesting.Equal(t, v.String(), str3)

	// noinspection GoBoolExpressions
	v, _ = IufsOf(b1)
	xtesting.Equal(t, v.Int(), int64(1))
	// noinspection GoBoolExpressions
	v, _ = IufsOf(b2)
	xtesting.Equal(t, v.Int(), int64(0))

	_, err := IufsOf(m1)
	xtesting.NotNil(t, err)
	_, err = IufsOf(m2)
	xtesting.NotNil(t, err)
	_, err = IufsOf(m3)
	xtesting.NotNil(t, err)
	_, err = IufsOf(s)
	xtesting.NotNil(t, err)
	_, err = IufsOf(p)
	xtesting.NotNil(t, err)
}

func TestIufSize(t *testing.T) {
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

	sze, _ := IufSizeOf(i)
	xtesting.Equal(t, sze.Int(), int64(i))
	sze, _ = IufSizeOf(i8)
	xtesting.Equal(t, sze.Int(), int64(i8))
	sze, _ = IufSizeOf(i16)
	xtesting.Equal(t, sze.Int(), int64(i16))
	sze, _ = IufSizeOf(i32)
	xtesting.Equal(t, sze.Int(), int64(i32))
	sze, _ = IufSizeOf(i64)
	xtesting.Equal(t, sze.Int(), i64)

	sze, _ = IufSizeOf(u)
	xtesting.Equal(t, sze.Uint(), uint64(u))
	sze, _ = IufSizeOf(u8)
	xtesting.Equal(t, sze.Uint(), uint64(u8))
	sze, _ = IufSizeOf(u16)
	xtesting.Equal(t, sze.Uint(), uint64(u16))
	sze, _ = IufSizeOf(u32)
	xtesting.Equal(t, sze.Uint(), uint64(u32))
	sze, _ = IufSizeOf(u64)
	xtesting.Equal(t, sze.Uint(), u64)
	sze, _ = IufSizeOf(up)
	xtesting.Equal(t, sze.Uint(), uint64(up))

	sze, _ = IufSizeOf(f32)
	xtesting.True(t, math.Abs(sze.Float()-float64(f32)) < 1e-3)
	sze, _ = IufSizeOf(f32)
	xtesting.True(t, math.Abs(sze.Float()-f64) < 1e-3)

	sze, _ = IufSizeOf(str1)
	xtesting.Equal(t, sze.Int(), int64(4))
	sze, _ = IufSizeOf(str2)
	xtesting.Equal(t, sze.Int(), int64(5))
	sze, _ = IufSizeOf(str3)
	xtesting.Equal(t, sze.Int(), int64(0))

	// noinspection GoBoolExpressions
	sze, _ = IufSizeOf(b1)
	xtesting.Equal(t, sze.Int(), int64(1))
	// noinspection GoBoolExpressions
	sze, _ = IufSizeOf(b2)
	xtesting.Equal(t, sze.Int(), int64(0))

	sze, _ = IufSizeOf(m1)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze, _ = IufSizeOf(m2)
	xtesting.Equal(t, sze.Int(), int64(3))
	sze, _ = IufSizeOf(m3)
	xtesting.Equal(t, sze.Int(), int64(3))

	_, err := IufSizeOf(s)
	xtesting.NotNil(t, err)
	_, err = IufSizeOf(p)
	xtesting.NotNil(t, err)
}
