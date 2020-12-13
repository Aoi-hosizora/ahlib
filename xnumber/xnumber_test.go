package xnumber

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"math"
	"testing"
)

func TestAccuracy(t *testing.T) {
	xtesting.True(t, NewAccuracy(1e-3).Equal(0.3333, 0.3334))
	xtesting.True(t, NewAccuracy(1e-3).NotEqual(0.333, 0.334))
	xtesting.True(t, NewAccuracy(1e-3).Greater(0.334, 0.333))
	xtesting.True(t, NewAccuracy(1e-3).GreaterOrEqual(0.3334, 0.3333))
	xtesting.True(t, NewAccuracy(1e-3).Smaller(0.333, 0.334))
	xtesting.True(t, NewAccuracy(1e-3).SmallerOrEqual(0.3333, 0.3334))

	xtesting.True(t, EqualInAccuracy(0.3333, 0.3334))
	xtesting.True(t, NotEqualInAccuracy(0.333, 0.334))
	xtesting.True(t, GreaterInAccuracy(0.334, 0.333))
	xtesting.True(t, GreaterOrEqualInAccuracy(0.3334, 0.3333))
	xtesting.True(t, SmallerInAccuracy(0.333, 0.334))
	xtesting.True(t, SmallerOrEqualInAccuracy(0.3333, 0.3334))
}

func TestRenderByte(t *testing.T) {
	xtesting.Equal(t, RenderByte(-1025), "-1.00KB")
	xtesting.Equal(t, RenderByte(-5), "-5B")
	xtesting.Equal(t, RenderByte(0), "0B")
	xtesting.Equal(t, RenderByte(1023), "1023B")
	xtesting.Equal(t, RenderByte(1024), "1.00KB")
	xtesting.Equal(t, RenderByte(1030), "1.01KB")
	xtesting.Equal(t, RenderByte(1536), "1.50KB")
	xtesting.Equal(t, RenderByte(2048), "2.00KB")
	xtesting.Equal(t, RenderByte(1024*1024), "1.00MB")
	xtesting.Equal(t, RenderByte(2.51*1024*1024), "2.51MB")
	xtesting.Equal(t, RenderByte(1024*1024*1024), "1.00GB")
	xtesting.Equal(t, RenderByte(2.51*1024*1024*1024), "2.51GB")
	xtesting.Equal(t, RenderByte(1024*1024*1024*1024), "1.00TB")
	xtesting.Equal(t, RenderByte(1.1*1024*1024*1024*1024), "1.10TB")
}

func TestBool(t *testing.T) {
	xtesting.Equal(t, Bool(true), 1)
	xtesting.Equal(t, Bool(true), 1)
	xtesting.Equal(t, Bool(false), 0)
	xtesting.Equal(t, Bool(false), 0)
}

func TestMinMax(t *testing.T) {
	xtesting.EqualValue(t, MinInt8, math.MinInt8)
	xtesting.EqualValue(t, MinInt16, math.MinInt16)
	xtesting.EqualValue(t, MinInt32, math.MinInt32)
	xtesting.EqualValue(t, MinInt64, math.MinInt64)
	xtesting.EqualValue(t, MinUint8, 0)
	xtesting.EqualValue(t, MinUint16, 0)
	xtesting.EqualValue(t, MinUint32, 0)
	xtesting.EqualValue(t, MinUint64, 0)

	xtesting.EqualValue(t, MaxInt8, math.MaxInt8)
	xtesting.EqualValue(t, MaxInt16, math.MaxInt16)
	xtesting.EqualValue(t, MaxInt32, math.MaxInt32)
	xtesting.EqualValue(t, MaxInt64, math.MaxInt64)
	xtesting.EqualValue(t, MaxUint8, 0xff)
	xtesting.EqualValue(t, MaxUint16, 0xffff)
	xtesting.EqualValue(t, MaxUint32, 0xffffffff)
	xtesting.EqualValue(t, MaxUint64, uint64(0xffffffffffffffff))

	xtesting.True(t, EqualInAccuracy(float64(MaxFloat32), math.MaxFloat32))
	xtesting.True(t, EqualInAccuracy(float64(SmallestNonzeroFloat32), math.SmallestNonzeroFloat32))
	xtesting.True(t, EqualInAccuracy(MaxFloat64, math.MaxFloat64))
	xtesting.True(t, EqualInAccuracy(SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))
}

func TestParse(t *testing.T) {
	i, _ := ParseInt("9223372036854775807", 10)
	xtesting.Equal(t, i, 9223372036854775807)
	u, _ := ParseUint("18446744073709551615", 10)
	xtesting.Equal(t, u, uint(18446744073709551615))
	i8, _ := ParseInt8("127", 10)
	xtesting.Equal(t, i8, int8(127))
	u8, _ := ParseUint8("255", 10)
	xtesting.Equal(t, u8, uint8(255))
	i16, _ := ParseInt16("32767", 10)
	xtesting.Equal(t, i16, int16(32767))
	u16, _ := ParseUint16("65535", 10)
	xtesting.Equal(t, u16, uint16(65535))
	i32, _ := ParseInt32("2147483647", 10)
	xtesting.Equal(t, i32, int32(2147483647))
	u32, _ := ParseUint32("4294967295", 10)
	xtesting.Equal(t, u32, uint32(4294967295))
	i64, _ := ParseInt64("9223372036854775807", 10)
	xtesting.Equal(t, i64, int64(9223372036854775807))
	u64, _ := ParseUint64("18446744073709551615", 10)
	xtesting.Equal(t, u64, uint64(18446744073709551615))
	f32, _ := ParseFloat32("0.5")
	xtesting.Equal(t, f32, float32(0.5))
	f64, _ := ParseFloat64("0.5")
	xtesting.Equal(t, f64, 0.5)

	_, err := ParseInt8("a", 10) // no number
	xtesting.NotNil(t, err)
	_, err = ParseInt8("a", 11) // success
	xtesting.Nil(t, err)
	_, err = ParseInt32("2147483648", 10) // overflow
	xtesting.NotNil(t, err)
	_, err = ParseInt64("10", 37) // base err
	xtesting.NotNil(t, err)
}

func TestParseOr(t *testing.T) {
	xtesting.Equal(t, ParseIntOr("9223372036854775807", 10, 0), 9223372036854775807)
	xtesting.Equal(t, ParseUintOr("18446744073709551615", 10, 0), uint(18446744073709551615))
	xtesting.Equal(t, ParseInt8Or("127", 10, 0), int8(127))
	xtesting.Equal(t, ParseUint8Or("255", 10, 0), uint8(255))
	xtesting.Equal(t, ParseInt16Or("32767", 10, 0), int16(32767))
	xtesting.Equal(t, ParseUint16Or("65535", 10, 0), uint16(65535))
	xtesting.Equal(t, ParseInt32Or("2147483647", 10, 0), int32(2147483647))
	xtesting.Equal(t, ParseUint32Or("4294967295", 10, 0), uint32(4294967295))
	xtesting.Equal(t, ParseInt64Or("9223372036854775807", 10, 0), int64(9223372036854775807))
	xtesting.Equal(t, ParseUint64Or("18446744073709551615", 10, 0), uint64(18446744073709551615))
	xtesting.Equal(t, ParseFloat32Or("0.5", 0), float32(0.5))
	xtesting.Equal(t, ParseFloat64Or("0.5", 0), 0.5)

	xtesting.Equal(t, ParseIntOr("", 10, 9223372036854775807), 9223372036854775807)
	xtesting.Equal(t, ParseUintOr("", 10, 18446744073709551615), uint(18446744073709551615))
	xtesting.Equal(t, ParseInt8Or("", 10, 127), int8(127))
	xtesting.Equal(t, ParseUint8Or("", 10, 255), uint8(255))
	xtesting.Equal(t, ParseInt16Or("", 10, 32767), int16(32767))
	xtesting.Equal(t, ParseUint16Or("", 10, 65535), uint16(65535))
	xtesting.Equal(t, ParseInt32Or("", 10, 2147483647), int32(2147483647))
	xtesting.Equal(t, ParseUint32Or("", 10, 4294967295), uint32(4294967295))
	xtesting.Equal(t, ParseInt64Or("", 10, 9223372036854775807), int64(9223372036854775807))
	xtesting.Equal(t, ParseUint64Or("", 10, 18446744073709551615), uint64(18446744073709551615))
	xtesting.Equal(t, ParseFloat32Or("", 0.5), float32(0.5))
	xtesting.Equal(t, ParseFloat64Or("", 0.5), 0.5)
}

func TestAtoX(t *testing.T) {
	i, _ := Atoi("9223372036854775807")
	xtesting.Equal(t, i, 9223372036854775807)
	u, _ := Atou("18446744073709551615")
	xtesting.Equal(t, u, uint(18446744073709551615))
	i8, _ := Atoi8("127")
	xtesting.Equal(t, i8, int8(127))
	u8, _ := Atou8("255")
	xtesting.Equal(t, u8, uint8(255))
	i16, _ := Atoi16("32767")
	xtesting.Equal(t, i16, int16(32767))
	u16, _ := Atou16("65535")
	xtesting.Equal(t, u16, uint16(65535))
	i32, _ := Atoi32("2147483647")
	xtesting.Equal(t, i32, int32(2147483647))
	u32, _ := Atou32("4294967295")
	xtesting.Equal(t, u32, uint32(4294967295))
	i64, _ := Atoi64("9223372036854775807")
	xtesting.Equal(t, i64, int64(9223372036854775807))
	u64, _ := Atou64("18446744073709551615")
	xtesting.Equal(t, u64, uint64(18446744073709551615))
	f32, _ := Atof32("0.5")
	xtesting.Equal(t, f32, float32(0.5))
	f64, _ := Atof64("0.5")
	xtesting.Equal(t, f64, 0.5)
}

func TestAtoXOr(t *testing.T) {
	xtesting.Equal(t, AtoiOr("9223372036854775807", 0), 9223372036854775807)
	xtesting.Equal(t, AtouOr("18446744073709551615", 0), uint(18446744073709551615))
	xtesting.Equal(t, Atoi8Or("127", 0), int8(127))
	xtesting.Equal(t, Atou8Or("255", 0), uint8(255))
	xtesting.Equal(t, Atoi16Or("32767", 0), int16(32767))
	xtesting.Equal(t, Atou16Or("65535", 0), uint16(65535))
	xtesting.Equal(t, Atoi32Or("2147483647", 0), int32(2147483647))
	xtesting.Equal(t, Atou32Or("4294967295", 0), uint32(4294967295))
	xtesting.Equal(t, Atoi64Or("9223372036854775807", 0), int64(9223372036854775807))
	xtesting.Equal(t, Atou64Or("18446744073709551615", 0), uint64(18446744073709551615))
	xtesting.Equal(t, Atof32Or("0.5", 0), float32(0.5))
	xtesting.Equal(t, Atof64Or("0.5", 0), 0.5)

	xtesting.Equal(t, AtoiOr("", 9223372036854775807), 9223372036854775807)
	xtesting.Equal(t, AtouOr("", 18446744073709551615), uint(18446744073709551615))
	xtesting.Equal(t, Atoi8Or("", 127), int8(127))
	xtesting.Equal(t, Atou8Or("", 255), uint8(255))
	xtesting.Equal(t, Atoi16Or("", 32767), int16(32767))
	xtesting.Equal(t, Atou16Or("", 65535), uint16(65535))
	xtesting.Equal(t, Atoi32Or("", 2147483647), int32(2147483647))
	xtesting.Equal(t, Atou32Or("", 4294967295), uint32(4294967295))
	xtesting.Equal(t, Atoi64Or("", 9223372036854775807), int64(9223372036854775807))
	xtesting.Equal(t, Atou64Or("", 18446744073709551615), uint64(18446744073709551615))
	xtesting.Equal(t, Atof32Or("", 0.5), float32(0.5))
	xtesting.Equal(t, Atof64Or("", 0.5), 0.5)
}

func TestFormat(t *testing.T) {
	xtesting.Equal(t, FormatInt(9223372036854775807, 10), "9223372036854775807")
	xtesting.Equal(t, FormatUint(18446744073709551615, 10), "18446744073709551615")
	xtesting.Equal(t, FormatInt8(127, 10), "127")
	xtesting.Equal(t, FormatUint8(255, 10), "255")
	xtesting.Equal(t, FormatInt16(32767, 10), "32767")
	xtesting.Equal(t, FormatUint16(65535, 10), "65535")
	xtesting.Equal(t, FormatInt32(2147483647, 10), "2147483647")
	xtesting.Equal(t, FormatUint32(4294967295, 10), "4294967295")
	xtesting.Equal(t, FormatInt64(9223372036854775807, 10), "9223372036854775807")
	xtesting.Equal(t, FormatUint64(18446744073709551615, 10), "18446744073709551615")
	xtesting.Equal(t, FormatFloat32(0.5, 'f', -1), "0.5")
	xtesting.Equal(t, FormatFloat64(0.5, 'f', -1), "0.5")
	xtesting.Equal(t, FormatFloat32(0.5555, 'e', 2), "5.55e-01")
	xtesting.Equal(t, FormatFloat64(0.5555, 'e', 2), "5.55e-01")
}

func TestXtoa(t *testing.T) {
	xtesting.Equal(t, Itoa(9223372036854775807), "9223372036854775807")
	xtesting.Equal(t, Utoa(18446744073709551615), "18446744073709551615")
	xtesting.Equal(t, I8toa(127), "127")
	xtesting.Equal(t, U8toa(255), "255")
	xtesting.Equal(t, I16toa(32767), "32767")
	xtesting.Equal(t, U16toa(65535), "65535")
	xtesting.Equal(t, I32toa(2147483647), "2147483647")
	xtesting.Equal(t, U32toa(4294967295), "4294967295")
	xtesting.Equal(t, I64toa(9223372036854775807), "9223372036854775807")
	xtesting.Equal(t, U64toa(18446744073709551615), "18446744073709551615")
	xtesting.Equal(t, F32toa(0.5), "0.5")
	xtesting.Equal(t, F64toa(0.5), "0.5")
}
