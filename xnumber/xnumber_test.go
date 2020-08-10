package xnumber

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRenderLatency(t *testing.T) {
	assert.Equal(t, RenderLatency(-0.1), "0.0000ns")
	assert.Equal(t, RenderLatency(0), "0.0000ns")
	assert.Equal(t, RenderLatency(999), "999.0000ns")
	assert.Equal(t, RenderLatency(10000), "10.0000µs")
	assert.Equal(t, RenderLatency(1000000), "1.0000ms")
	assert.Equal(t, RenderLatency(10000000000), "10.0000s")
	assert.Equal(t, RenderLatency(59000000000), "59.0000s")
	assert.Equal(t, RenderLatency(60000000000), "1.0000min")
}

func TestRenderByte(t *testing.T) {
	assert.Equal(t, RenderByte(-5), "0B")
	assert.Equal(t, RenderByte(0), "0B")
	assert.Equal(t, RenderByte(1023), "1023B")
	assert.Equal(t, RenderByte(1024), "1.00KB")
	assert.Equal(t, RenderByte(1536), "1.50KB")
	assert.Equal(t, RenderByte(2048), "2.00KB")
	assert.Equal(t, RenderByte(1024*1024), "1.00MB")
	assert.Equal(t, RenderByte(2.51*1024*1024), "2.51MB")
}

func TestParse(t *testing.T) {
	log.Println(ParseInt("9223372036854775807", 10))
	log.Println(ParseUint("18446744073709551615", 10))
	log.Println(ParseInt8("127", 10))
	log.Println(ParseUint8("255", 10))
	log.Println(ParseInt16("32767", 10))
	log.Println(ParseUint16("65535", 10))
	log.Println(ParseInt32("2147483647", 10))
	log.Println(ParseUint32("4294967295", 10))
	log.Println(ParseInt64("9223372036854775807", 10))
	log.Println(ParseUint64("18446744073709551615", 10))
	log.Println(ParseFloat32("0.7"))
	log.Println(ParseFloat64("0.7"))
}

func TestFormat(t *testing.T) {
	log.Println("\"" + FormatInt(9223372036854775807, 10) + "\"")
	log.Println("\"" + FormatUint(18446744073709551615, 10) + "\"")
	log.Println("\"" + FormatInt8(127, 10) + "\"")
	log.Println("\"" + FormatUint8(255, 10) + "\"")
	log.Println("\"" + FormatInt16(32767, 10) + "\"")
	log.Println("\"" + FormatUint16(65535, 10) + "\"")
	log.Println("\"" + FormatInt32(2147483647, 10) + "\"")
	log.Println("\"" + FormatUint32(4294967295, 10) + "\"")
	log.Println("\"" + FormatInt64(9223372036854775807, 10) + "\"")
	log.Println("\"" + FormatUint64(18446744073709551615, 10) + "\"")
	log.Println("\"" + FormatFloat32(0.7, 'f', -1) + "\"")
	log.Println("\"" + FormatFloat32(0.7, 'f', -1) + "\"")
}

func TestAtoi(t *testing.T) {
	log.Println(Atoi("9223372036854775807"))
	log.Println(Atou("18446744073709551615"))
	log.Println(Atoi8("127"))
	log.Println(Atou8("255"))
	log.Println(Atoi16("32767"))
	log.Println(Atou16("65535"))
	log.Println(Atoi32("2147483647"))
	log.Println(Atou32("4294967295"))
	log.Println(Atoi64("9223372036854775807"))
	log.Println(Atou64("18446744073709551615"))
	log.Println(Atof32("0.7"))
	log.Println(Atof64("0.7"))
}

func TestItoa(t *testing.T) {
	log.Println("\"" + Itoa(9223372036854775807) + "\"")
	log.Println("\"" + Utoa(18446744073709551615) + "\"")
	log.Println("\"" + I8toa(127) + "\"")
	log.Println("\"" + U8toa(255) + "\"")
	log.Println("\"" + I16toa(32767) + "\"")
	log.Println("\"" + U16toa(65535) + "\"")
	log.Println("\"" + I32toa(2147483647) + "\"")
	log.Println("\"" + U32toa(4294967295) + "\"")
	log.Println("\"" + I64toa(9223372036854775807) + "\"")
	log.Println("\"" + U64toa(18446744073709551615) + "\"")
	log.Println("\"" + F32toa(0.7) + "\"")
	log.Println("\"" + F64toa(0.7) + "\"")
}
