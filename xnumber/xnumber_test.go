package xnumber

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestAccuracy(t *testing.T) {
	acc := NewAccuracy(1e-2)

	for _, tc := range []struct {
		giveFn func(a, b float64) bool
		giveA  float64
		giveB  float64
		want   bool
	}{
		{EqualInAccuracy, 0.333, 0.332, false}, // 0.333 - 0.332 = 0.001 != 0.001
		{EqualInAccuracy, 0.333, 0.333, true},
		{EqualInAccuracy, 0.333, 0.334, false},
		{EqualInAccuracy, 0.3333, 0.3332, true}, // 0.3333 - 0.3332 = 0.0001 != 0.0001
		{EqualInAccuracy, 0.3333, 0.3333, true},
		{EqualInAccuracy, 0.3333, 0.3334, true},
		{acc.Equal, 0.33, 0.32, false},
		{acc.Equal, 0.33, 0.33, true},
		{acc.Equal, 0.33, 0.34, false},
		{acc.Equal, 0.333, 0.332, true},
		{acc.Equal, 0.333, 0.333, true},
		{acc.Equal, 0.333, 0.334, true},

		{NotEqualInAccuracy, 0.333, 0.332, true},
		{NotEqualInAccuracy, 0.333, 0.333, false},
		{NotEqualInAccuracy, 0.333, 0.334, true},
		{NotEqualInAccuracy, 0.3333, 0.3332, false},
		{NotEqualInAccuracy, 0.3333, 0.3333, false},
		{NotEqualInAccuracy, 0.3333, 0.3334, false},
		{acc.NotEqual, 0.33, 0.32, true},
		{acc.NotEqual, 0.33, 0.33, false},
		{acc.NotEqual, 0.33, 0.34, true},
		{acc.NotEqual, 0.333, 0.332, false},
		{acc.NotEqual, 0.333, 0.333, false},
		{acc.NotEqual, 0.333, 0.334, false},

		{GreaterInAccuracy, 0.333, 0.332, true},
		{GreaterInAccuracy, 0.333, 0.333, false},
		{GreaterInAccuracy, 0.333, 0.334, false},
		{GreaterInAccuracy, 0.3333, 0.3332, false},
		{GreaterInAccuracy, 0.3333, 0.3333, false},
		{GreaterInAccuracy, 0.3333, 0.3334, false},
		{acc.Greater, 0.33, 0.32, true},
		{acc.Greater, 0.33, 0.33, false},
		{acc.Greater, 0.33, 0.34, false},
		{acc.Greater, 0.333, 0.332, false},
		{acc.Greater, 0.333, 0.333, false},
		{acc.Greater, 0.333, 0.334, false},

		{GreaterOrEqualInAccuracy, 0.333, 0.332, true},
		{GreaterOrEqualInAccuracy, 0.333, 0.333, true},
		{GreaterOrEqualInAccuracy, 0.333, 0.334, false},
		{GreaterOrEqualInAccuracy, 0.3333, 0.3332, true},
		{GreaterOrEqualInAccuracy, 0.3333, 0.3333, true},
		{GreaterOrEqualInAccuracy, 0.3333, 0.3334, true},
		{acc.GreaterOrEqual, 0.33, 0.32, true},
		{acc.GreaterOrEqual, 0.33, 0.33, true},
		{acc.GreaterOrEqual, 0.33, 0.34, false},
		{acc.GreaterOrEqual, 0.333, 0.332, true},
		{acc.GreaterOrEqual, 0.333, 0.333, true},
		{acc.GreaterOrEqual, 0.333, 0.334, true},

		{LessInAccuracy, 0.333, 0.332, false},
		{LessInAccuracy, 0.333, 0.333, false},
		{LessInAccuracy, 0.333, 0.334, true},
		{LessInAccuracy, 0.3333, 0.3332, false},
		{LessInAccuracy, 0.3333, 0.3333, false},
		{LessInAccuracy, 0.3333, 0.3334, false},
		{acc.Less, 0.33, 0.32, false},
		{acc.Less, 0.33, 0.33, false},
		{acc.Less, 0.33, 0.34, true},
		{acc.Less, 0.333, 0.332, false},
		{acc.Less, 0.333, 0.333, false},
		{acc.Less, 0.333, 0.334, false},

		{LessOrEqualInAccuracy, 0.333, 0.332, false},
		{LessOrEqualInAccuracy, 0.333, 0.333, true},
		{LessOrEqualInAccuracy, 0.333, 0.334, true},
		{LessOrEqualInAccuracy, 0.3333, 0.3332, true},
		{LessOrEqualInAccuracy, 0.3333, 0.3333, true},
		{LessOrEqualInAccuracy, 0.3333, 0.3334, true},
		{acc.LessOrEqual, 0.33, 0.32, false},
		{acc.LessOrEqual, 0.33, 0.33, true},
		{acc.LessOrEqual, 0.33, 0.34, true},
		{acc.LessOrEqual, 0.333, 0.332, true},
		{acc.LessOrEqual, 0.333, 0.333, true},
		{acc.LessOrEqual, 0.333, 0.334, true},
	} {
		xtesting.Equal(t, tc.giveFn(tc.giveA, tc.giveB), tc.want)
	}
}

func TestFormatByteSize(t *testing.T) {
	for _, tc := range []struct {
		give float64
		want string
	}{
		{-1025, "-1.00KB"},
		{-5, "-5B"},
		{0, "0B"},
		{1023, "1023B"},
		{1024, "1.00KB"},
		{1030, "1.01KB"},
		{1536, "1.50KB"},
		{2048, "2.00KB"},
		{1024 * 1024, "1.00MB"},
		{2.51 * 1024 * 1024, "2.51MB"},
		{1024 * 1024 * 1024, "1.00GB"},
		{2.51 * 1024 * 1024 * 1024, "2.51GB"},
		{1024 * 1024 * 1024 * 1024, "1.00TB"},
		{1.1 * 1024 * 1024 * 1024 * 1024, "1.10TB"},
	} {
		xtesting.Equal(t, FormatByteSize(tc.give), tc.want)
	}
}

func TestBool(t *testing.T) {
	xtesting.Equal(t, Bool(true), 1)
	xtesting.Equal(t, Bool(false), 0)
}

func TestIntBitLength(t *testing.T) {
	xtesting.Equal(t, IntBitLength(), strconv.IntSize)
}

func TestFastrand(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Println(FastrandUint32())
	}
	for i := 0; i < 5; i++ {
		log.Println(FastrandUint64())
	}
	for i := 0; i < 5; i++ {
		log.Println(FastrandInt32())
	}
	for i := 0; i < 5; i++ {
		log.Println(FastrandInt64())
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, tc := range []struct {
		give int
		want bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{1023, false},
		{1024, true},
		{2047, false},
		{2048, true},
		{65535, false},
		{65536, true},
		{1073741823, false},
		{1073741824, true},
		{2147483647, false},
	} {
		t.Run(strconv.Itoa(tc.give), func(t *testing.T) {
			xtesting.Equal(t, IsPowerOfTwo(tc.give), tc.want)
		})
	}
}

func TestMinMax(t *testing.T) {
	for _, tc := range []struct {
		give interface{}
		want interface{}
	}{
		{MinInt8, math.MinInt8},
		{MinInt16, math.MinInt16},
		{MinInt32, math.MinInt32},
		{MinInt64, math.MinInt64},
		{MinUint8, 0},
		{MinUint16, 0},
		{MinUint32, 0},
		{MinUint64, 0},

		{MaxInt8, math.MaxInt8},
		{MaxInt16, math.MaxInt16},
		{MaxInt32, math.MaxInt32},
		{MaxInt64, math.MaxInt64},
		{MaxUint8, 0xff},
		{MaxUint16, 0xffff},
		{MaxUint32, 0xffffffff},
		{MaxUint64, uint64(0xffffffffffffffff)},
	} {
		xtesting.EqualValue(t, tc.give, tc.want)
	}

	xtesting.True(t, EqualInAccuracy(float64(MaxFloat32), math.MaxFloat32))
	xtesting.True(t, EqualInAccuracy(float64(SmallestNonzeroFloat32), math.SmallestNonzeroFloat32))
	xtesting.True(t, EqualInAccuracy(MaxFloat64, math.MaxFloat64))
	xtesting.True(t, EqualInAccuracy(SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))
}

func TestOverflowWhen(t *testing.T) {
	// signed add (positive+positive)
	xtesting.True(t, OverflowWhenAddInt8(MaxInt8, 1))
	xtesting.True(t, OverflowWhenAddInt16(MaxInt16, 1))
	xtesting.True(t, OverflowWhenAddInt32(MaxInt32, 1))
	xtesting.True(t, OverflowWhenAddInt64(MaxInt64, 1))
	xtesting.False(t, OverflowWhenAddInt8(MaxInt8-1, 1))
	xtesting.False(t, OverflowWhenAddInt16(MaxInt16-1, 1))
	xtesting.False(t, OverflowWhenAddInt32(MaxInt32-1, 1))
	xtesting.False(t, OverflowWhenAddInt64(MaxInt64-1, 1))

	// signed add (negative+negative)
	xtesting.True(t, OverflowWhenAddInt8(MinInt8, -1))
	xtesting.True(t, OverflowWhenAddInt16(MinInt16, -1))
	xtesting.True(t, OverflowWhenAddInt32(MinInt32, -1))
	xtesting.True(t, OverflowWhenAddInt64(MinInt64, -1))
	xtesting.False(t, OverflowWhenAddInt8(MinInt8+1, -1))
	xtesting.False(t, OverflowWhenAddInt16(MinInt16+1, -1))
	xtesting.False(t, OverflowWhenAddInt32(MinInt32+1, -1))
	xtesting.False(t, OverflowWhenAddInt64(MinInt64+1, -1))

	// signed subtract (negative-positive)
	xtesting.True(t, OverflowWhenSubtractInt8(MinInt8, 1))
	xtesting.True(t, OverflowWhenSubtractInt16(MinInt16, 1))
	xtesting.True(t, OverflowWhenSubtractInt32(MinInt32, 1))
	xtesting.True(t, OverflowWhenSubtractInt64(MinInt64, 1))
	xtesting.False(t, OverflowWhenSubtractInt8(MinInt8+1, 1))
	xtesting.False(t, OverflowWhenSubtractInt16(MinInt16+1, 1))
	xtesting.False(t, OverflowWhenSubtractInt32(MinInt32+1, 1))
	xtesting.False(t, OverflowWhenSubtractInt64(MinInt64+1, 1))

	// signed subtract (positive-negative)
	xtesting.True(t, OverflowWhenSubtractInt8(MaxInt8, -1))
	xtesting.True(t, OverflowWhenSubtractInt16(MaxInt16, -1))
	xtesting.True(t, OverflowWhenSubtractInt32(MaxInt32, -1))
	xtesting.True(t, OverflowWhenSubtractInt64(MaxInt64, -1))
	xtesting.False(t, OverflowWhenSubtractInt8(MaxInt8-1, -1))
	xtesting.False(t, OverflowWhenSubtractInt16(MaxInt16-1, -1))
	xtesting.False(t, OverflowWhenSubtractInt32(MaxInt32-1, -1))
	xtesting.False(t, OverflowWhenSubtractInt64(MaxInt64-1, -1))

	// unsigned add (positive+positive)
	xtesting.True(t, OverflowWhenAddUint8(MaxUint8, 1))
	xtesting.True(t, OverflowWhenAddUint16(MaxUint16, 1))
	xtesting.True(t, OverflowWhenAddUint32(MaxUint32, 1))
	xtesting.True(t, OverflowWhenAddUint64(MaxUint64, 1))
	xtesting.False(t, OverflowWhenAddUint8(MaxUint8-1, 1))
	xtesting.False(t, OverflowWhenAddUint16(MaxUint16-1, 1))
	xtesting.False(t, OverflowWhenAddUint32(MaxUint32-1, 1))
	xtesting.False(t, OverflowWhenAddUint64(MaxUint64-1, 1))

	// unsigned subtract (positive-positive)
	xtesting.True(t, OverflowWhenSubtractUint8(1, 2))
	xtesting.True(t, OverflowWhenSubtractUint16(1, 2))
	xtesting.True(t, OverflowWhenSubtractUint32(1, 2))
	xtesting.True(t, OverflowWhenSubtractUint64(1, 2))
	xtesting.False(t, OverflowWhenSubtractUint8(2, 1))
	xtesting.False(t, OverflowWhenSubtractUint16(2, 1))
	xtesting.False(t, OverflowWhenSubtractUint32(2, 1))
	xtesting.False(t, OverflowWhenSubtractUint64(2, 1))

	originBitLength := intBitLength
	defer func() { intBitLength = originBitLength }()
	intBitLength = 32
	xtesting.True(t, OverflowWhenAddInt(int(MaxInt32), 1))
	xtesting.False(t, OverflowWhenAddInt(int(MaxInt32-1), 1))
	xtesting.True(t, OverflowWhenAddInt(int(MinInt32), -1))
	xtesting.False(t, OverflowWhenAddInt(int(MinInt32+1), -1))
	xtesting.True(t, OverflowWhenSubtractInt(int(MinInt32), 1))
	xtesting.False(t, OverflowWhenSubtractInt(int(MinInt32+1), 1))
	xtesting.True(t, OverflowWhenSubtractInt(int(MaxInt32), -1))
	xtesting.False(t, OverflowWhenSubtractInt(int(MaxInt32-1), -1))
	xtesting.True(t, OverflowWhenAddUint(uint(MaxUint32), 1))
	xtesting.False(t, OverflowWhenAddUint(uint(MaxUint32-1), 1))
	xtesting.True(t, OverflowWhenSubtractUint(uint(1), 2))
	xtesting.False(t, OverflowWhenSubtractUint(uint(2), 1))
	intBitLength = 64
	xtesting.True(t, OverflowWhenAddInt(int(MaxInt64), 1))
	xtesting.False(t, OverflowWhenAddInt(int(MaxInt64-1), 1))
	xtesting.True(t, OverflowWhenAddInt(int(MinInt64), -1))
	xtesting.False(t, OverflowWhenAddInt(int(MinInt64+1), -1))
	xtesting.True(t, OverflowWhenSubtractInt(int(MinInt64), 1))
	xtesting.False(t, OverflowWhenSubtractInt(int(MinInt64+1), 1))
	xtesting.True(t, OverflowWhenSubtractInt(int(MaxInt64), -1))
	xtesting.False(t, OverflowWhenSubtractInt(int(MaxInt64-1), -1))
	xtesting.True(t, OverflowWhenAddUint(uint(MaxUint64), 1))
	xtesting.False(t, OverflowWhenAddUint(uint(MaxUint64-1), 1))
	xtesting.True(t, OverflowWhenSubtractUint(uint(1), 2))
	xtesting.False(t, OverflowWhenSubtractUint(uint(2), 1))
}

func TestParse(t *testing.T) {
	i, _ := ParseInt("9223372036854775807", 10)
	xtesting.Equal(t, i, 9223372036854775807)
	i8, _ := ParseInt8("127", 10)
	xtesting.Equal(t, i8, int8(127))
	i16, _ := ParseInt16("32767", 10)
	xtesting.Equal(t, i16, int16(32767))
	i32, _ := ParseInt32("2147483647", 10)
	xtesting.Equal(t, i32, int32(2147483647))
	i64, _ := ParseInt64("9223372036854775807", 10)
	xtesting.Equal(t, i64, int64(9223372036854775807))

	u, _ := ParseUint("18446744073709551615", 10)
	xtesting.Equal(t, u, uint(18446744073709551615))
	u8, _ := ParseUint8("255", 10)
	xtesting.Equal(t, u8, uint8(255))
	u16, _ := ParseUint16("65535", 10)
	xtesting.Equal(t, u16, uint16(65535))
	u32, _ := ParseUint32("4294967295", 10)
	xtesting.Equal(t, u32, uint32(4294967295))
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
	xtesting.Equal(t, ParseInt8Or("127", 10, 0), int8(127))
	xtesting.Equal(t, ParseInt16Or("32767", 10, 0), int16(32767))
	xtesting.Equal(t, ParseInt32Or("2147483647", 10, 0), int32(2147483647))
	xtesting.Equal(t, ParseInt64Or("9223372036854775807", 10, 0), int64(9223372036854775807))
	xtesting.Equal(t, ParseUintOr("18446744073709551615", 10, 0), uint(18446744073709551615))
	xtesting.Equal(t, ParseUint8Or("255", 10, 0), uint8(255))
	xtesting.Equal(t, ParseUint16Or("65535", 10, 0), uint16(65535))
	xtesting.Equal(t, ParseUint32Or("4294967295", 10, 0), uint32(4294967295))
	xtesting.Equal(t, ParseUint64Or("18446744073709551615", 10, 0), uint64(18446744073709551615))
	xtesting.Equal(t, ParseFloat32Or("0.5", 0), float32(0.5))
	xtesting.Equal(t, ParseFloat64Or("0.5", 0), 0.5)

	xtesting.Equal(t, ParseIntOr("", 10, 9223372036854775807), 9223372036854775807)
	xtesting.Equal(t, ParseInt8Or("", 10, 127), int8(127))
	xtesting.Equal(t, ParseInt16Or("", 10, 32767), int16(32767))
	xtesting.Equal(t, ParseInt32Or("", 10, 2147483647), int32(2147483647))
	xtesting.Equal(t, ParseInt64Or("", 10, 9223372036854775807), int64(9223372036854775807))
	xtesting.Equal(t, ParseUintOr("", 10, 18446744073709551615), uint(18446744073709551615))
	xtesting.Equal(t, ParseUint8Or("", 10, 255), uint8(255))
	xtesting.Equal(t, ParseUint16Or("", 10, 65535), uint16(65535))
	xtesting.Equal(t, ParseUint32Or("", 10, 4294967295), uint32(4294967295))
	xtesting.Equal(t, ParseUint64Or("", 10, 18446744073709551615), uint64(18446744073709551615))
	xtesting.Equal(t, ParseFloat32Or("", 0.5), float32(0.5))
	xtesting.Equal(t, ParseFloat64Or("", 0.5), 0.5)
}

func TestAtoX(t *testing.T) {
	i, _ := Atoi("9223372036854775807")
	xtesting.Equal(t, i, 9223372036854775807)
	i8, _ := Atoi8("127")
	xtesting.Equal(t, i8, int8(127))
	i16, _ := Atoi16("32767")
	xtesting.Equal(t, i16, int16(32767))
	u32, _ := Atou32("4294967295")
	xtesting.Equal(t, u32, uint32(4294967295))
	i64, _ := Atoi64("9223372036854775807")
	xtesting.Equal(t, i64, int64(9223372036854775807))

	u, _ := Atou("18446744073709551615")
	xtesting.Equal(t, u, uint(18446744073709551615))
	u8, _ := Atou8("255")
	xtesting.Equal(t, u8, uint8(255))
	u16, _ := Atou16("65535")
	xtesting.Equal(t, u16, uint16(65535))
	i32, _ := Atoi32("2147483647")
	xtesting.Equal(t, i32, int32(2147483647))
	u64, _ := Atou64("18446744073709551615")
	xtesting.Equal(t, u64, uint64(18446744073709551615))

	f32, _ := Atof32("0.5")
	xtesting.Equal(t, f32, float32(0.5))
	f64, _ := Atof64("0.5")
	xtesting.Equal(t, f64, 0.5)
}

func TestAtoXOr(t *testing.T) {
	xtesting.Equal(t, AtoiOr("9223372036854775807", 0), 9223372036854775807)
	xtesting.Equal(t, Atoi8Or("127", 0), int8(127))
	xtesting.Equal(t, Atoi16Or("32767", 0), int16(32767))
	xtesting.Equal(t, Atoi32Or("2147483647", 0), int32(2147483647))
	xtesting.Equal(t, Atoi64Or("9223372036854775807", 0), int64(9223372036854775807))
	xtesting.Equal(t, AtouOr("18446744073709551615", 0), uint(18446744073709551615))
	xtesting.Equal(t, Atou8Or("255", 0), uint8(255))
	xtesting.Equal(t, Atou16Or("65535", 0), uint16(65535))
	xtesting.Equal(t, Atou32Or("4294967295", 0), uint32(4294967295))
	xtesting.Equal(t, Atou64Or("18446744073709551615", 0), uint64(18446744073709551615))
	xtesting.Equal(t, Atof32Or("0.5", 0), float32(0.5))
	xtesting.Equal(t, Atof64Or("0.5", 0), 0.5)

	xtesting.Equal(t, AtoiOr("", 9223372036854775807), 9223372036854775807)
	xtesting.Equal(t, Atoi8Or("", 127), int8(127))
	xtesting.Equal(t, Atoi16Or("", 32767), int16(32767))
	xtesting.Equal(t, Atoi32Or("", 2147483647), int32(2147483647))
	xtesting.Equal(t, Atoi64Or("", 9223372036854775807), int64(9223372036854775807))
	xtesting.Equal(t, AtouOr("", 18446744073709551615), uint(18446744073709551615))
	xtesting.Equal(t, Atou8Or("", 255), uint8(255))
	xtesting.Equal(t, Atou16Or("", 65535), uint16(65535))
	xtesting.Equal(t, Atou32Or("", 4294967295), uint32(4294967295))
	xtesting.Equal(t, Atou64Or("", 18446744073709551615), uint64(18446744073709551615))
	xtesting.Equal(t, Atof32Or("", 0.5), float32(0.5))
	xtesting.Equal(t, Atof64Or("", 0.5), 0.5)
}

func TestFormat(t *testing.T) {
	xtesting.Equal(t, FormatInt(9223372036854775807, 10), "9223372036854775807")
	xtesting.Equal(t, FormatInt8(127, 10), "127")
	xtesting.Equal(t, FormatInt16(32767, 10), "32767")
	xtesting.Equal(t, FormatInt32(2147483647, 10), "2147483647")
	xtesting.Equal(t, FormatInt64(9223372036854775807, 10), "9223372036854775807")
	xtesting.Equal(t, FormatUint(18446744073709551615, 10), "18446744073709551615")
	xtesting.Equal(t, FormatUint8(255, 10), "255")
	xtesting.Equal(t, FormatUint16(65535, 10), "65535")
	xtesting.Equal(t, FormatUint32(4294967295, 10), "4294967295")
	xtesting.Equal(t, FormatUint64(18446744073709551615, 10), "18446744073709551615")
	xtesting.Equal(t, FormatFloat32(0.5, 'f', -1), "0.5")
	xtesting.Equal(t, FormatFloat64(0.5, 'f', -1), "0.5")
	xtesting.Equal(t, FormatFloat32(0.5555, 'e', 2), "5.55e-01")
	xtesting.Equal(t, FormatFloat64(0.5555, 'e', 2), "5.55e-01")
}

func TestXtoa(t *testing.T) {
	xtesting.Equal(t, Itoa(9223372036854775807), "9223372036854775807")
	xtesting.Equal(t, I8toa(127), "127")
	xtesting.Equal(t, I16toa(32767), "32767")
	xtesting.Equal(t, I32toa(2147483647), "2147483647")
	xtesting.Equal(t, I64toa(9223372036854775807), "9223372036854775807")
	xtesting.Equal(t, Utoa(18446744073709551615), "18446744073709551615")
	xtesting.Equal(t, U8toa(255), "255")
	xtesting.Equal(t, U16toa(65535), "65535")
	xtesting.Equal(t, U32toa(4294967295), "4294967295")
	xtesting.Equal(t, U64toa(18446744073709551615), "18446744073709551615")
	xtesting.Equal(t, F32toa(0.5), "0.5")
	xtesting.Equal(t, F64toa(0.5), "0.5")
}

func TestRange(t *testing.T) {
	for _, tc := range []struct {
		giveFrom int8
		giveTo   int8
		giveStep int8
		want     []int8
	}{
		{1, 1, 1, nil},
		{1, 1, -1, nil},
		{1, 10, 0, nil},
		{10, 1, 0, nil},
		{1, 10, -1, nil},
		{10, 1, 1, nil},
		{1, 2, 1, []int8{1}},
		{1, 0, -1, []int8{1}},
		{1, 2, 100, []int8{1}},
		{1, 0, -100, []int8{1}},

		{0, 10, 1, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{1, 10, 1, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{0, 10, 2, []int8{0, 2, 4, 6, 8}},
		{1, 10, 2, []int8{1, 3, 5, 7, 9}},
		{0, 10, 3, []int8{0, 3, 6, 9}},
		{1, 10, 3, []int8{1, 4, 7}},
		{0, 10, 4, []int8{0, 4, 8}},
		{1, 10, 4, []int8{1, 5, 9}},
		{0, 10, 5, []int8{0, 5}},
		{1, 10, 5, []int8{1, 6}},
		{0, 10, 6, []int8{0, 6}},
		{1, 10, 6, []int8{1, 7}},
		{0, 10, 9, []int8{0, 9}},
		{1, 10, 9, []int8{1}},
		{0, 10, 10, []int8{0}},
		{1, 10, 10, []int8{1}},
		{0, 10, 100, []int8{0}},
		{1, 10, 100, []int8{1}},

		{11, 0, -1, []int8{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{10, 0, -1, []int8{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{11, 0, -2, []int8{11, 9, 7, 5, 3, 1}},
		{10, 0, -2, []int8{10, 8, 6, 4, 2}},
		{11, 0, -3, []int8{11, 8, 5, 2}},
		{10, 0, -3, []int8{10, 7, 4, 1}},
		{11, 0, -4, []int8{11, 7, 3}},
		{10, 0, -4, []int8{10, 6, 2}},
		{11, 0, -5, []int8{11, 6, 1}},
		{10, 0, -5, []int8{10, 5}},
		{11, 0, -6, []int8{11, 5}},
		{10, 0, -6, []int8{10, 4}},
		{11, 0, -9, []int8{11, 2}},
		{10, 0, -9, []int8{10, 1}},
		{11, 0, -10, []int8{11, 1}},
		{10, 0, -10, []int8{10}},
		{11, 0, -100, []int8{11}},
		{10, 0, -100, []int8{10}},
	} {
		t.Run(fmt.Sprintf("%d_%d_%d", tc.giveFrom, tc.giveTo, tc.giveStep), func(t *testing.T) {
			intSlice := IntRange(int(tc.giveFrom), int(tc.giveTo), int(tc.giveStep))
			int8Slice := Int8Range(tc.giveFrom, tc.giveTo, tc.giveStep)
			int16Slice := Int16Range(int16(tc.giveFrom), int16(tc.giveTo), int16(tc.giveStep))
			int32Slice := Int32Range(int32(tc.giveFrom), int32(tc.giveTo), int32(tc.giveStep))
			int64Slice := Int64Range(int64(tc.giveFrom), int64(tc.giveTo), int64(tc.giveStep))
			for _, slice := range []interface{}{intSlice, int8Slice, int16Slice, int32Slice, int64Slice} {
				val := reflect.ValueOf(slice)
				xtesting.Equal(t, val.Len(), len(tc.want))
				for i := 0; i < val.Len(); i++ {
					xtesting.Equal(t, int8(val.Index(i).Int()), tc.want[i])
				}
			}

			if tc.giveFrom >= 0 && tc.giveTo >= 0 {
				flag := int8(+1)
				if tc.giveStep < 0 {
					flag = int8(-1)
				}
				uintSlice := UintRange(uint(tc.giveFrom), uint(tc.giveTo), uint(flag*tc.giveStep), flag == -1)
				uint8Slice := Uint8Range(uint8(tc.giveFrom), uint8(tc.giveTo), uint8(flag*tc.giveStep), flag == -1)
				uint16Slice := Uint16Range(uint16(tc.giveFrom), uint16(tc.giveTo), uint16(flag*tc.giveStep), flag == -1)
				uint32Slice := Uint32Range(uint32(tc.giveFrom), uint32(tc.giveTo), uint32(flag*tc.giveStep), flag == -1)
				uint64Slice := Uint64Range(uint64(tc.giveFrom), uint64(tc.giveTo), uint64(flag*tc.giveStep), flag == -1)
				for _, slice := range []interface{}{uintSlice, uint8Slice, uint16Slice, uint32Slice, uint64Slice} {
					val := reflect.ValueOf(slice)
					xtesting.Equal(t, val.Len(), len(tc.want))
					for i := 0; i < val.Len(); i++ {
						xtesting.Equal(t, int8(val.Index(i).Uint()), tc.want[i])
					}
				}
			}
		})
	}

	t.Run("overflow", func(t *testing.T) {
		xtesting.Equal(t, IntRange(int(MaxInt64-3), int(MaxInt64), 2), []int{int(MaxInt64 - 3), int(MaxInt64 - 1)})
		xtesting.Equal(t, Int8Range(MaxInt8-3, MaxInt8, 2), []int8{MaxInt8 - 3, MaxInt8 - 1})
		xtesting.Equal(t, Int16Range(MaxInt16-3, MaxInt16, 2), []int16{MaxInt16 - 3, MaxInt16 - 1})
		xtesting.Equal(t, Int32Range(MaxInt32-3, MaxInt32, 2), []int32{MaxInt32 - 3, MaxInt32 - 1})
		xtesting.Equal(t, Int64Range(MaxInt64-3, MaxInt64, 2), []int64{MaxInt64 - 3, MaxInt64 - 1})

		xtesting.Equal(t, IntRange(int(MinInt64+3), int(MinInt64), -2), []int{int(MinInt64 + 3), int(MinInt64 + 1)})
		xtesting.Equal(t, Int8Range(MinInt8+3, MinInt8, -2), []int8{MinInt8 + 3, MinInt8 + 1})
		xtesting.Equal(t, Int16Range(MinInt16+3, MinInt16, -2), []int16{MinInt16 + 3, MinInt16 + 1})
		xtesting.Equal(t, Int32Range(MinInt32+3, MinInt32, -2), []int32{MinInt32 + 3, MinInt32 + 1})
		xtesting.Equal(t, Int64Range(MinInt64+3, MinInt64, -2), []int64{MinInt64 + 3, MinInt64 + 1})

		xtesting.Equal(t, UintRange(uint(MaxUint64-3), uint(MaxUint64), 2), []uint{uint(MaxUint64 - 3), uint(MaxUint64 - 1)})
		xtesting.Equal(t, Uint8Range(MaxUint8-3, MaxUint8, 2), []uint8{MaxUint8 - 3, MaxUint8 - 1})
		xtesting.Equal(t, Uint16Range(MaxUint16-3, MaxUint16, 2), []uint16{MaxUint16 - 3, MaxUint16 - 1})
		xtesting.Equal(t, Uint32Range(MaxUint32-3, MaxUint32, 2), []uint32{MaxUint32 - 3, MaxUint32 - 1})
		xtesting.Equal(t, Uint64Range(MaxUint64-3, MaxUint64, 2), []uint64{MaxUint64 - 3, MaxUint64 - 1})

		xtesting.Equal(t, UintRange(3, 0, 2, true), []uint{3, 1})
		xtesting.Equal(t, Uint8Range(3, 0, 2, true), []uint8{3, 1})
		xtesting.Equal(t, Uint16Range(3, 0, 2, true), []uint16{3, 1})
		xtesting.Equal(t, Uint32Range(3, 0, 2, true), []uint32{3, 1})
		xtesting.Equal(t, Uint64Range(3, 0, 2, true), []uint64{3, 1})
	})
}

func TestReverseSlice(t *testing.T) {
	for _, tc := range []struct {
		give []int8
		want []int8
	}{
		{[]int8{}, []int8{}},
		{[]int8{0}, []int8{0}},
		{[]int8{1, 2}, []int8{2, 1}},
		{[]int8{1, 2, 3}, []int8{3, 2, 1}},
		{[]int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}},
		{[]int8{1, 1, 3, 2, 2}, []int8{2, 2, 3, 1, 1}},
		{[]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int8{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{[]int8{0, 1, 4, 7, 2, 5, 8, 3, 6, 9}, []int8{9, 6, 3, 8, 5, 2, 7, 4, 1, 0}},
	} {
		t.Run(fmt.Sprintf("%v", tc.give), func(t *testing.T) {
			l := len(tc.give)
			intSlice, int8Slice, int16Slice, int32Slice, int64Slice := make([]int, l), make([]int8, l), make([]int16, l), make([]int32, l), make([]int64, l)
			uintSlice, uint8Slice, uint16Slice, uint32Slice, uint64Slice := make([]uint, l), make([]uint8, l), make([]uint16, l), make([]uint32, l), make([]uint64, l)
			for _, obj := range []struct {
				slice    interface{}
				f        interface{}
				unsigned bool
			}{
				{intSlice, ReverseIntSlice, false},
				{int8Slice, ReverseInt8Slice, false},
				{int16Slice, ReverseInt16Slice, false},
				{int32Slice, ReverseInt32Slice, false},
				{int64Slice, ReverseInt64Slice, false},
				{uintSlice, ReverseUintSlice, true},
				{uint8Slice, ReverseUint8Slice, true},
				{uint16Slice, ReverseUint16Slice, true},
				{uint32Slice, ReverseUint32Slice, true},
				{uint64Slice, ReverseUint64Slice, true},
			} {
				val := reflect.ValueOf(obj.slice)
				for idx := 0; idx < val.Len(); idx++ {
					if !obj.unsigned {
						val.Index(idx).SetInt(int64(tc.give[idx]))
					} else {
						val.Index(idx).SetUint(uint64(tc.give[idx]))
					}
				}
				reflect.ValueOf(obj.f).Call([]reflect.Value{val})
				xtesting.Equal(t, val.Len(), len(tc.want))
				for i := 0; i < val.Len(); i++ {
					if !obj.unsigned {
						xtesting.Equal(t, int8(val.Index(i).Int()), tc.want[i])
					} else {
						xtesting.Equal(t, int8(val.Index(i).Uint()), tc.want[i])
					}
				}
			}
		})
	}
}
