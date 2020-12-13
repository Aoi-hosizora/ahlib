package xnumber

import (
	"strconv"
)

// parse

// ParseInt parses a string to int (32bit / 64bit) using given base, see strconv.ParseInt.
func ParseInt(s string, base int) (int, error) {
	i, e := strconv.ParseInt(s, base, 0)
	return int(i), e
}

// ParseInt8 parses a string to int8 using given base, see strconv.ParseInt.
func ParseInt8(s string, base int) (int8, error) {
	i, e := strconv.ParseInt(s, base, 8)
	return int8(i), e
}

// ParseInt16 parses a string to int16 using given base, see strconv.ParseInt.
func ParseInt16(s string, base int) (int16, error) {
	i, e := strconv.ParseInt(s, base, 16)
	return int16(i), e
}

// ParseInt32 parses a string to int32 using given base, see strconv.ParseInt.
func ParseInt32(s string, base int) (int32, error) {
	i, e := strconv.ParseInt(s, base, 32)
	return int32(i), e
}

// ParseInt64 parses a string to int64 using given base, see strconv.ParseInt.
func ParseInt64(s string, base int) (int64, error) {
	i, e := strconv.ParseInt(s, base, 64)
	return i, e
}

// ParseUint parses a string to uint (32bit / 64bit) using given base, see strconv.ParseUint.
func ParseUint(s string, base int) (uint, error) {
	i, e := strconv.ParseUint(s, base, 0)
	return uint(i), e
}

// ParseUint8 parses a string to uint8 using given base, see strconv.ParseUint.
func ParseUint8(s string, base int) (uint8, error) {
	i, e := strconv.ParseUint(s, base, 8)
	return uint8(i), e
}

// ParseUint16 parses a string to uint16 using given base, see strconv.ParseUint.
func ParseUint16(s string, base int) (uint16, error) {
	i, e := strconv.ParseUint(s, base, 16)
	return uint16(i), e
}

// ParseUint32 parses a string to uint32 using given base, see strconv.ParseUint.
func ParseUint32(s string, base int) (uint32, error) {
	i, e := strconv.ParseUint(s, base, 32)
	return uint32(i), e
}

// ParseUint64 parses a string to uint64 using given base, see strconv.ParseUint.
func ParseUint64(s string, base int) (uint64, error) {
	i, e := strconv.ParseUint(s, base, 64)
	return i, e
}

// ParseFloat32 parses a string to float32, see strconv.ParseFloat.
func ParseFloat32(s string) (float32, error) {
	f, e := strconv.ParseFloat(s, 32)
	return float32(f), e
}

// ParseFloat64 parses a string to float64, see strconv.ParseFloat.
func ParseFloat64(s string) (float64, error) {
	f, e := strconv.ParseFloat(s, 64)
	return f, e
}

// atoX

// Atoi parses a string to int (32bit / 64bit) in base 10.
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

// Atou parses a string to uint (32bit / 64bit) in base 10.
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

// format

// FormatInt formats a int to string using given base, see strconv.FormatInt.
func FormatInt(i int, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt8 formats a int8 to string using given base, see strconv.FormatInt.
func FormatInt8(i int8, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt16 formats a int16 to string using given base, see strconv.FormatInt.
func FormatInt16(i int16, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt32 formats a int32 to string using given base, see strconv.FormatInt.
func FormatInt32(i int32, base int) string {
	return strconv.FormatInt(int64(i), base)
}

// FormatInt64 formats a int64 to string using given base, see strconv.FormatInt.
func FormatInt64(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// FormatUint formats a uint to string using given base, see strconv.FormatUint.
func FormatUint(i uint, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint8 formats a uint8 to string using given base, see strconv.FormatUint.
func FormatUint8(i uint8, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint16 formats a uint16 to string using given base, see strconv.FormatUint.
func FormatUint16(i uint16, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint32 formats a uint32 to string using given base, see strconv.FormatUint.
func FormatUint32(i uint32, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint64 formats a uint64 to string using given base, see strconv.FormatUint.
func FormatUint64(i uint64, base int) string {
	return strconv.FormatUint(i, base)
}

// FormatFloat32 formats a float32 to string using given format and precision, see strconv.FormatFloat.
func FormatFloat32(f float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, 32)
}

// FormatFloat64 formats a float64 to string using given format and precision, see strconv.FormatFloat.
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
func Utoa(i uint) string {
	return FormatUint(i, 10)
}

// U8toa formats a uint8 to string in base 10.
func U8toa(i uint8) string {
	return FormatUint8(i, 10)
}

// U16toa formats a uint16 to string in base 10.
func U16toa(i uint16) string {
	return FormatUint16(i, 10)
}

// U32toa formats a uint32 to string in base 10.
func U32toa(i uint32) string {
	return FormatUint32(i, 10)
}

// U64toa formats a uint64 to string in base 10.
func U64toa(i uint64) string {
	return FormatUint64(i, 10)
}

// F32toa formats a float32 to string using given 'f' format and -1 precision, see strconv.FormatFloat.
func F32toa(f float32) string {
	return FormatFloat32(f, 'f', -1)
}

// F64toa formats a float64 to string using given 'f' format and -1 precision, see strconv.FormatFloat.
func F64toa(f float64) string {
	return FormatFloat64(f, 'f', -1)
}
