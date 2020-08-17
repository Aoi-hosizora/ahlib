package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"log"
	"testing"
)

func TestElemType(t *testing.T) {
	var a ****int
	t1 := ElemType(a).String()
	log.Println(t1)
	if !IsEqual(t1, "int") {
		t.Fatal("t1 is equal to int, nut got not equal")
	}

	var b int
	t2 := ElemType(b).String()
	log.Println(t2)
	if !IsEqual(t2, "int") {
		t.Fatal("t2 is equal to int, nut got not equal")
	}
}

func TestElemValue(t *testing.T) {
	var a *****int
	v1 := ElemValue(a)
	log.Println(v1.IsValid())
	if !IsEqual(v1.IsValid(), false) {
		t.Fatal("`v1.IsValid()` is equal to false, nut got not equal")
	}

	var b int
	v2 := ElemValue(b).Interface()
	log.Println(v2)
	if !IsEqual(v2, 0) {
		t.Fatal("v2 is equal to 0, nut got not equal")
	}

	var c = &b
	v3 := ElemValue(c).Interface()
	log.Println(v3)
	if !IsEqual(v3, 0) {
		t.Fatal("v3 is equal to 0, nut got not equal")
	}
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

	if xcondition.First(GetInt(i)) != int64(i) { t.Fail() }
	if xcondition.First(GetInt(i8)) != int64(i8) { t.Fatal() }
	if xcondition.First(GetInt(i16)) != int64(i16) { t.Fatal() }
	if xcondition.First(GetInt(i32)) != int64(i32) { t.Fatal() }
	if xcondition.First(GetInt(i64)) != i64 { t.Fatal() }
	if xcondition.First(GetUint(u)) != uint64(u) { t.Fatal() }
	if xcondition.First(GetUint(u8)) != uint64(u8) { t.Fatal() }
	if xcondition.First(GetUint(u16)) != uint64(u16) { t.Fatal() }
	if xcondition.First(GetUint(u32)) != uint64(u32) { t.Fatal() }
	if xcondition.First(GetUint(u64)) != u64 { t.Fatal() }
	if xcondition.First(GetUint(up)) != uint64(up) { t.Fatal() }
	if !xnumber.NewAccuracy(1e-3).Equal(xcondition.First(GetFloat(f32)).(float64), 0.1) { t.Fatal() }
	if !xnumber.NewAccuracy(1e-3).Equal(xcondition.First(GetFloat(f64)).(float64), 0.1) { t.Fatal() }
	if xcondition.First(GetString(str1)) != str1 { t.Fatal() }
	if xcondition.First(GetString(str2)) != str2 { t.Fatal() }
	if xcondition.First(GetString(str3)) != str3 { t.Fatal() }
	// noinspection GoBoolExpressions
	if xcondition.First(GetBool(b1)) != b1 { t.Fatal() }
	// noinspection GoBoolExpressions
	if xcondition.First(GetBool(b2)) != b2 { t.Fatal() }
}

func TestGetValueSize(t *testing.T) {
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
	s := &struct{}{}

	sze, _ := GetValueSize(i)
	if sze.Int() != int64(i) { t.Fatal() }
	sze, _ = GetValueSize(i8)
	if sze.Int() != int64(i8) { t.Fatal() }
	sze, _ = GetValueSize(i16)
	if sze.Int() != int64(i16) { t.Fatal() }
	sze, _ = GetValueSize(i32)
	if sze.Int() != int64(i32) { t.Fatal() }
	sze, _ = GetValueSize(i64)
	if sze.Int() != i64 { t.Fatal() }

	sze, _ = GetValueSize(u)
	if sze.Uint() != uint64(u) { t.Fatal() }
	sze, _ = GetValueSize(u8)
	if sze.Uint() != uint64(u8) { t.Fatal() }
	sze, _ = GetValueSize(u16)
	if sze.Uint() != uint64(u16) { t.Fatal() }
	sze, _ = GetValueSize(u32)
	if sze.Uint() != uint64(u32) { t.Fatal() }
	sze, _ = GetValueSize(u64)
	if sze.Uint() != u64 { t.Fatal() }
	sze, _ = GetValueSize(up)
	if sze.Uint() != uint64(up) { t.Fatal() }

	sze, _ = GetValueSize(f32)
	if !xnumber.NewAccuracy(1e-3).Equal(sze.Float(), float64(f32)) { t.Fatal() }
	sze, _ = GetValueSize(f32)
	if !xnumber.NewAccuracy(1e-3).Equal(sze.Float(), f64) { t.Fatal() }

	sze, _ = GetValueSize(str1)
	if sze.Int() != 4 { t.Fatal() }
	sze, _ = GetValueSize(str2)
	if sze.Int() != 5 { t.Fatal() }
	sze, _ = GetValueSize(str3)
	if sze.Int() != 0 { t.Fatal() }

	// noinspection GoBoolExpressions
	sze, _ = GetValueSize(b1)
	if sze.Int() == 0 { t.Fatal() }
	// noinspection GoBoolExpressions
	sze, _ = GetValueSize(b2)
	if sze.Int() != 0 { t.Fatal() }

	_, err := GetValueSize(s)
	if err == nil { t.Fatal() }
}