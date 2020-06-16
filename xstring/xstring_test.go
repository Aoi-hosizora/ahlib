package xstring

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCapitalize(t *testing.T) {
	assert.Equal(t, Capitalize("abc"), "Abc")
	assert.Equal(t, Capitalize("Abc"), "Abc")
	assert.Equal(t, Capitalize(""), "")
}

func TestUncapitalize(t *testing.T) {
	assert.Equal(t, Uncapitalize("Abc"), "abc")
	assert.Equal(t, Uncapitalize("abc"), "abc")
	assert.Equal(t, Uncapitalize(""), "")
}

func TestMarshalJson(t *testing.T) {
	a := struct {
		F1 string `json:"f1"`
		F2 struct{ F3 int }
	}{
		F1: "a",
		F2: struct{ F3 int }{F3: 3},
	}
	assert.Equal(t, MarshalJson(a), "{\"f1\":\"a\",\"F2\":{\"F3\":3}}")
}

func TestCurrentTimeUuid(t *testing.T) {
	log.Println(CurrentTimeUuid(5))
	log.Println(CurrentTimeUuid(24))
	log.Println(CurrentTimeUuid(30))
}

func TestRandLetterNumberString(t *testing.T) {
	log.Println(RandLetterString(20))
	log.Println(RandLetterString(20))
	log.Println(RandNumberString(20))
	log.Println(RandNumberString(20))

	log.Println(RandString(32, CapitalLetterRunes))
	log.Println(RandString(32, LowercaseLetterRunes))
	log.Println(RandString(32, NumberRunes))

	log.Println(RandString(32, LetterRunes))
	log.Println(RandString(32, LetterNumberRunes))
	log.Println(RandString(32, CapitalLetterNumberRunes))
	log.Println(RandString(32, LowercaseLetterNumberRunes))
}

func TestPrettifyJson(t *testing.T) {
	from := "{\"a\": \"b\", \"c\": {\"d\": \"e\", \"f\": 0}, \"g\": [{\"h\": 1}, {\"h\": 1}]}"
	to := "{\n" +
		"    \"a\": \"b\",\n" +
		"    \"c\": {\n" +
		"        \"d\": \"e\",\n" +
		"        \"f\": 0\n" +
		"    },\n" +
		"    \"g\": [\n" +
		"        {\n" +
		"            \"h\": 1\n" +
		"        },\n" +
		"        {\n" +
		"            \"h\": 1\n" +
		"        }\n" +
		"    ]\n" +
		"}"
	assert.Equal(t, PrettifyJson(from, 4, " "), to)
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, ToSnakeCase(""), "")
	assert.Equal(t, ToSnakeCase("AoiHosizora"), "aoi_hosizora")
	assert.Equal(t, ToSnakeCase("abc0d1EdF"), "abc0d1_ed_f")
	assert.Equal(t, ToSnakeCase("私達isわたしたち"), "私達isわたしたち")
	assert.Equal(t, ToSnakeCase("a bC"), "a_b_c")
}

func TestIsLowercase(t *testing.T) {
	assert.Equal(t, IsLowercase(ToRune("A")), false)
	assert.Equal(t, IsLowercase(ToRune("Z")), false)
	assert.Equal(t, IsLowercase(ToRune("a")), true)
	assert.Equal(t, IsLowercase(ToRune("z")), true)
	assert.Equal(t, IsLowercase(ToRune("0")), false)
	assert.Equal(t, IsLowercase(ToRune("")), false)
	assert.Equal(t, IsLowercase(ToRune("我")), false)
}

func TestIsUppercase(t *testing.T) {
	assert.Equal(t, IsUppercase(ToRune("A")), true)
	assert.Equal(t, IsUppercase(ToRune("Z")), true)
	assert.Equal(t, IsUppercase(ToRune("a")), false)
	assert.Equal(t, IsUppercase(ToRune("z")), false)
	assert.Equal(t, IsUppercase(ToRune("0")), false)
	assert.Equal(t, IsUppercase(ToRune("")), false)
	assert.Equal(t, IsUppercase(ToRune("我")), false)
}

func TestRemoveSpaces(t *testing.T) {
	assert.Equal(t, RemoveSpaces(""), "")
	assert.Equal(t, RemoveSpaces("a b  c 　d   e f"), "a b c d e f")
	assert.Equal(t, RemoveSpaces("a b 	 c d   e f"), "a b c d e f")
	assert.Equal(t, RemoveSpaces("a b \n	 c d   e f"), "a b c d e f")
	assert.Equal(t, RemoveSpaces("\n　"), "")
	assert.Equal(t, RemoveSpaces("　\n	"), "")
}
