package xnumber

import (
	"strconv"
)

// parse

// ParseInt parses a string to int using given base.
func ParseInt(s string, base int) (int, error) {
	i, e := strconv.ParseInt(s, base, 0)
	return int(i), e
}

// ParseInt8 parses a string to int8 using given base.
func ParseInt8(s string, base int) (int8, error) {
	i, e := strconv.ParseInt(s, base, 8)
	return int8(i), e
}

// ParseInt16 parses a string to int16 using given base.
func ParseInt16(s string, base int) (int16, error) {
	i, e := strconv.ParseInt(s, base, 16)
	return int16(i), e
}

// ParseInt32 parses a string to int32 using given base.
func ParseInt32(s string, base int) (int32, error) {
	i, e := strconv.ParseInt(s, base, 32)
	return int32(i), e
}

// ParseInt64 parses a string to int64 using given base.
func ParseInt64(s string, base int) (int64, error) {
	i, e := strconv.ParseInt(s, base, 64)
	return i, e
}

// ParseUint parses a string to uint using given base.
func ParseUint(s string, base int) (uint, error) {
	u, e := strconv.ParseUint(s, base, 0)
	return uint(u), e
}

// ParseUint8 parses a string to uint8 using given base.
func ParseUint8(s string, base int) (uint8, error) {
	u, e := strconv.ParseUint(s, base, 8)
	return uint8(u), e
}

// ParseUint16 parses a string to uint16 using given base.
func ParseUint16(s string, base int) (uint16, error) {
	u, e := strconv.ParseUint(s, base, 16)
	return uint16(u), e
}

// ParseUint32 parses a string to uint32 using given base.
func ParseUint32(s string, base int) (uint32, error) {
	u, e := strconv.ParseUint(s, base, 32)
	return uint32(u), e
}

// ParseUint64 parses a string to uint64 using given base.
func ParseUint64(s string, base int) (uint64, error) {
	u, e := strconv.ParseUint(s, base, 64)
	return u, e
}

// ParseFloat32 parses a string to float32.
func ParseFloat32(s string) (float32, error) {
	f, e := strconv.ParseFloat(s, 32)
	return float32(f), e
}

// ParseFloat64 parses a string to float64.
func ParseFloat64(s string) (float64, error) {
	f, e := strconv.ParseFloat(s, 64)
	return f, e
}

// parseOr

// ParseIntOr parses a string to int using given base with a fallback value.
func ParseIntOr(s string, base int, o int) int {
	i, e := ParseInt(s, base)
	if e != nil {
		return o
	}
	return i
}

// ParseInt8Or parses a string to int8 using given base with a fallback value.
func ParseInt8Or(s string, base int, o int8) int8 {
	i, e := ParseInt8(s, base)
	if e != nil {
		return o
	}
	return i
}

// ParseInt16Or parses a string to int16 using given base with a fallback value.
func ParseInt16Or(s string, base int, o int16) int16 {
	i, e := ParseInt16(s, base)
	if e != nil {
		return o
	}
	return i
}

// ParseInt32Or parses a string to int32 using given base with a fallback value.
func ParseInt32Or(s string, base int, o int32) int32 {
	i, e := ParseInt32(s, base)
	if e != nil {
		return o
	}
	return i
}

// ParseInt64Or parses a string to int64 using given base with a fallback value.
func ParseInt64Or(s string, base int, o int64) int64 {
	i, e := ParseInt64(s, base)
	if e != nil {
		return o
	}
	return i
}

// ParseUintOr parses a string to uint using given base with a fallback value.
func ParseUintOr(s string, base int, o uint) uint {
	u, e := ParseUint(s, base)
	if e != nil {
		return o
	}
	return u
}

// ParseUint8Or parses a string to uint8 using given base with a fallback value.
func ParseUint8Or(s string, base int, o uint8) uint8 {
	u, e := ParseUint8(s, base)
	if e != nil {
		return o
	}
	return u
}

// ParseUint16Or parses a string to uint16 using given base with a fallback value.
func ParseUint16Or(s string, base int, o uint16) uint16 {
	u, e := ParseUint16(s, base)
	if e != nil {
		return o
	}
	return u
}

// ParseUint32Or parses a string to uint32 using given base with a fallback value.
func ParseUint32Or(s string, base int, o uint32) uint32 {
	u, e := ParseUint32(s, base)
	if e != nil {
		return o
	}
	return u
}

// ParseUint64Or parses a string to uint64 using given base with a fallback value.
func ParseUint64Or(s string, base int, o uint64) uint64 {
	u, e := ParseUint64(s, base)
	if e != nil {
		return o
	}
	return u
}

// ParseFloat32Or parses a string to float32 with a fallback value.
func ParseFloat32Or(s string, o float32) float32 {
	f, e := ParseFloat32(s)
	if e != nil {
		return o
	}
	return f
}

// ParseFloat64Or parses a string to float64 with a fallback value.
func ParseFloat64Or(s string, o float64) float64 {
	f, e := ParseFloat64(s)
	if e != nil {
		return o
	}
	return f
}

// atoX

// Atoi parses a string to int in base 10.
func Atoi(s string) (int, error) {
	return ParseInt(s, 10)
}

// Atoi8 parses a string to int8 in base 10.
func Atoi8(s string) (int8, error) {
	return ParseInt8(s, 10)
}

// Atoi16 parses a string to int8 in base 10.
func Atoi16(s string) (int16, error) {
	return ParseInt16(s, 10)
}

// Atoi32 parses a string to int32 in base 10.
func Atoi32(s string) (int32, error) {
	return ParseInt32(s, 10)
}

// Atoi64 parses a string to int64 in base 10.
func Atoi64(s string) (int64, error) {
	return ParseInt64(s, 10)
}

// Atou parses a string to uint in base 10.
func Atou(s string) (uint, error) {
	return ParseUint(s, 10)
}

// Atou8 parses a string to uint8 in base 10.
func Atou8(s string) (uint8, error) {
	return ParseUint8(s, 10)
}

// Atou16 parses a string to uint16 in base 10.
func Atou16(s string) (uint16, error) {
	return ParseUint16(s, 10)
}

// Atou32 parses a string to uint32 in base 10.
func Atou32(s string) (uint32, error) {
	return ParseUint32(s, 10)
}

// Atou64 parses a string to uint64 in base 10.
func Atou64(s string) (uint64, error) {
	return ParseUint64(s, 10)
}

// Atof32 parses a string to float32, is same as ParseFloat32.
func Atof32(s string) (float32, error) {
	return ParseFloat32(s)
}

// Atof64 parses a string to float32, is same as ParseFloat64.
func Atof64(s string) (float64, error) {
	return ParseFloat64(s)
}

// atoXOr

// AtoiOr parses a string to int in base 10 with a fallback value.
func AtoiOr(s string, o int) int {
	i, e := Atoi(s)
	if e != nil {
		return o
	}
	return i
}

// Atoi8Or parses a string to int8 in base 10 with a fallback value.
func Atoi8Or(s string, o int8) int8 {
	i, e := Atoi8(s)
	if e != nil {
		return o
	}
	return i
}

// Atoi16Or parses a string to int8 in base 10 with a fallback value.
func Atoi16Or(s string, o int16) int16 {
	i, e := Atoi16(s)
	if e != nil {
		return o
	}
	return i
}

// Atoi32Or parses a string to int32 in base 10 with a fallback value.
func Atoi32Or(s string, o int32) int32 {
	i, e := Atoi32(s)
	if e != nil {
		return o
	}
	return i
}

// Atoi64Or parses a string to int64 in base 10 with a fallback value.
func Atoi64Or(s string, o int64) int64 {
	i, e := Atoi64(s)
	if e != nil {
		return o
	}
	return i
}

// AtouOr parses a string to uint in base 10 with a fallback value.
func AtouOr(s string, o uint) uint {
	u, e := Atou(s)
	if e != nil {
		return o
	}
	return u
}

// Atou8Or parses a string to uint8 in base 10 with a fallback value.
func Atou8Or(s string, o uint8) uint8 {
	u, e := Atou8(s)
	if e != nil {
		return o
	}
	return u
}

// Atou16Or parses a string to uint16 in base 10 with a fallback value.
func Atou16Or(s string, o uint16) uint16 {
	u, e := Atou16(s)
	if e != nil {
		return o
	}
	return u
}

// Atou32Or parses a string to uint32 in base 10 with a fallback value.
func Atou32Or(s string, o uint32) uint32 {
	u, e := Atou32(s)
	if e != nil {
		return o
	}
	return u
}

// Atou64Or parses a string to uint64 in base 10 with a fallback value.
func Atou64Or(s string, o uint64) uint64 {
	u, e := Atou64(s)
	if e != nil {
		return o
	}
	return u
}

// Atof32Or parses a string to float32 with a fallback value.
func Atof32Or(s string, o float32) float32 {
	f, e := Atof32(s)
	if e != nil {
		return o
	}
	return f
}

// Atof64Or parses a string to float32 with a fallback value.
func Atof64Or(s string, o float64) float64 {
	f, e := Atof64(s)
	if e != nil {
		return o
	}
	return f
}

// format

// FormatInt formats a int to string using given base.
func FormatInt(i int, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt8 formats a int8 to string using given base.
func FormatInt8(i int8, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt16 formats a int16 to string using given base.
func FormatInt16(i int16, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt32 formats a int32 to string using given base.
func FormatInt32(i int32, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt64 formats a int64 to string using given base.
func FormatInt64(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// FormatUint formats a uint to string using given base.
func FormatUint(u uint, base int) string {
	return strconv.FormatUint(uint64(u), base)
}

// FormatUint8 formats a uint8 to string using given base.
func FormatUint8(u uint8, base int) string {
	return strconv.FormatUint(uint64(u), base)
}

// FormatUint16 formats a uint16 to string using given base.
func FormatUint16(u uint16, base int) string {
	return strconv.FormatUint(uint64(u), base)
}

// FormatUint32 formats a uint32 to string using given base.
func FormatUint32(u uint32, base int) string {
	return strconv.FormatUint(uint64(u), base)
}

// FormatUint64 formats a uint64 to string using given base.
func FormatUint64(u uint64, base int) string {
	return strconv.FormatUint(u, base)
}

// FormatFloat32 formats a float32 to string using given format and precision.
func FormatFloat32(f float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, 32)
}

// FormatFloat64 formats a float64 to string using given format and precision.
func FormatFloat64(f float64, fmt byte, prec int) string {
	return strconv.FormatFloat(f, fmt, prec, 64)
}

// Xtoa

// Itoa formats a int to string in base 10.
func Itoa(i int) string {
	return FormatInt(i, 10)
}

// I8toa formats a int8 to string in base 10.
func I8toa(i int8) string {
	return FormatInt8(i, 10)
}

// I16toa formats a int16 to string in base 10.
func I16toa(i int16) string {
	return FormatInt16(i, 10)
}

// I32toa formats a int32 to string in base 10.
func I32toa(i int32) string {
	return FormatInt32(i, 10)
}

// I64toa formats a int64 to string in base 10.
func I64toa(i int64) string {
	return FormatInt64(i, 10)
}

// Utoa formats a uint to string in base 10.
func Utoa(u uint) string {
	return FormatUint(u, 10)
}

// U8toa formats a uint8 to string in base 10.
func U8toa(u uint8) string {
	return FormatUint8(u, 10)
}

// U16toa formats a uint16 to string in base 10.
func U16toa(u uint16) string {
	return FormatUint16(u, 10)
}

// U32toa formats a uint32 to string in base 10.
func U32toa(u uint32) string {
	return FormatUint32(u, 10)
}

// U64toa formats a uint64 to string in base 10.
func U64toa(u uint64) string {
	return FormatUint64(u, 10)
}

// F32toa formats a float32 to string using default format.
func F32toa(f float32) string {
	return FormatFloat32(f, 'f', -1)
}

// F64toa formats a float64 to string using default format.
func F64toa(f float64) string {
	return FormatFloat64(f, 'f', -1)
}
