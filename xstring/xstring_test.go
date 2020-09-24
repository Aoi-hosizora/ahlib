package xstring

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
	"time"
)

func TestCapitalize(t *testing.T) {
	xtesting.Equal(t, Capitalize("abc"), "Abc")
	xtesting.Equal(t, Capitalize("Abc"), "Abc")
	xtesting.Equal(t, Capitalize(""), "")
}

func TestUncapitalize(t *testing.T) {
	xtesting.Equal(t, Uncapitalize("Abc"), "abc")
	xtesting.Equal(t, Uncapitalize("abc"), "abc")
	xtesting.Equal(t, Uncapitalize(""), "")
}

func TestChatAt(t *testing.T) {
	str := "テスト测试測試test"
	xtesting.Equal(t, str, str)
	xtesting.Equal(t, ChatAt(str, 0), "テ")
	xtesting.Equal(t, ChatAt(str, 3), "测")
	xtesting.Equal(t, ChatAt(str, 4), "试")
	xtesting.Equal(t, ChatAt(str, 5), "測")
	xtesting.Equal(t, ChatAt(str, 7), "t")
}

func TestSubString(t *testing.T) {
	str := "テスト测试測試test"
	xtesting.Equal(t, str, str)
	xtesting.Equal(t, SubStringTo(str, 3), "テスト")
	xtesting.Equal(t, SubString(str, 3, 5), "测试")
	xtesting.Equal(t, SubString(str, 5, 7), "測試")
	xtesting.Equal(t, SubStringFrom(str, 7), "test")
}

func TestToRuneToByte(t *testing.T) {
	log.Printf("%T", 'a')         // int32
	log.Printf("%T", ToRune("a")) // int32
	log.Printf("%T", ToByte("a")) // uint8
	log.Printf("%T", "a"[0])      // uint8

	xtesting.Equal(t, ToRune("a"), 'a')
	xtesting.Equal(t, ToRune("bcd"), 'b')
	xtesting.Equal(t, ToRune(""), rune(0))

	xtesting.Equal(t, ToByte("a"), byte('a'))
	xtesting.Equal(t, ToByte("bcd"), byte('b'))
	xtesting.Equal(t, ToByte(""), byte(0))
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
	log.Println(RandLetterNumberString(20))
	log.Println(RandLetterNumberString(20))

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
	xtesting.Equal(t, PrettifyJson(from, 4, " "), to)
}

func TestToSnakeCase(t *testing.T) {
	xtesting.Equal(t, ToSnakeCase(""), "")
	xtesting.Equal(t, ToSnakeCase("AoiHosizora"), "aoi_hosizora")
	xtesting.Equal(t, ToSnakeCase("abc0d1EdF"), "abc0d1_ed_f")
	xtesting.Equal(t, ToSnakeCase("私達isわたしたち"), "私達isわたしたち")
	xtesting.Equal(t, ToSnakeCase("a bC"), "a_b_c")
}

func TestIsLowercase(t *testing.T) {
	xtesting.Equal(t, IsLowercase(ToRune("A")), false)
	xtesting.Equal(t, IsLowercase(ToRune("Z")), false)
	xtesting.Equal(t, IsLowercase(ToRune("a")), true)
	xtesting.Equal(t, IsLowercase(ToRune("z")), true)
	xtesting.Equal(t, IsLowercase(ToRune("0")), false)
	xtesting.Equal(t, IsLowercase(ToRune("")), false)
	xtesting.Equal(t, IsLowercase(ToRune("我")), false)
}

func TestIsUppercase(t *testing.T) {
	xtesting.Equal(t, IsUppercase(ToRune("A")), true)
	xtesting.Equal(t, IsUppercase(ToRune("Z")), true)
	xtesting.Equal(t, IsUppercase(ToRune("a")), false)
	xtesting.Equal(t, IsUppercase(ToRune("z")), false)
	xtesting.Equal(t, IsUppercase(ToRune("0")), false)
	xtesting.Equal(t, IsUppercase(ToRune("")), false)
	xtesting.Equal(t, IsUppercase(ToRune("我")), false)
}

func TestRemoveSpaces(t *testing.T) {
	xtesting.Equal(t, RemoveSpaces(""), "")
	xtesting.Equal(t, RemoveSpaces("a b  c 　d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveSpaces("a b 	 c d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveSpaces("a b \n	 c d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveSpaces("\n　"), "")
	xtesting.Equal(t, RemoveSpaces("　\n	"), "")
}

func TestMaskToken(t *testing.T) {
	xtesting.Equal(t, MaskToken(""), "")
	xtesting.Equal(t, MaskToken(" "), "*")
	xtesting.Equal(t, MaskToken("a"), "*")
	xtesting.Equal(t, MaskToken("aa"), "*a")
	xtesting.Equal(t, MaskToken("aaa"), "**a")
	xtesting.Equal(t, MaskToken("aaaa"), "a**a")
	xtesting.Equal(t, MaskToken("aaaaa"), "a***a")
	xtesting.Equal(t, MaskToken("aaaaaa"), "aa**aa")
	xtesting.Equal(t, MaskToken("あ"), "*")
	xtesting.Equal(t, MaskToken("あa"), "*a")
	xtesting.Equal(t, MaskToken("あaa"), "**a")
	xtesting.Equal(t, MaskToken("あaaa"), "あ**a")
	xtesting.Equal(t, MaskToken("あaaaa"), "あ***a")
	xtesting.Equal(t, MaskToken("あaaaaa"), "あa**aa")
}

func TestStringToBytes(t *testing.T) {
	xtesting.Equal(t, StringToBytes(""), []byte{})
	xtesting.Equal(t, StringToBytes("abcdefg"), []byte("abcdefg"))

	cnt := 2000000

	bs1 := make([]byte, cnt, cnt)
	bs2 := make([]byte, cnt, cnt)
	for i := 0; i < cnt; i++ {
		bs1[i] = 'A'
	}
	for i := 0; i < cnt; i++ {
		bs2[i] = 'B'
	}
	str1 := string(bs1)
	str2 := string(bs2)

	start := time.Now()
	bs01 := []byte(str1)
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	bs02 := StringToBytes(str2)
	log.Println(time.Now().Sub(start).String())

	xtesting.Equal(t, bs01, bs1)
	xtesting.Equal(t, bs02, bs2)
}

func TestBytesToString(t *testing.T) {
	xtesting.Equal(t, BytesToString(nil), "")
	xtesting.Equal(t, BytesToString([]byte{}), "")
	xtesting.Equal(t, BytesToString([]byte("abcdefg")), "abcdefg")

	cnt := 2000000

	bs1 := make([]byte, cnt, cnt)
	bs2 := make([]byte, cnt, cnt)
	for i := 0; i < cnt; i++ {
		bs1[i] = 'A'
	}
	for i := 0; i < cnt; i++ {
		bs2[i] = 'B'
	}
	str1 := string(bs1)
	str2 := string(bs2)

	start := time.Now()
	str01 := string(bs1)
	log.Println(time.Now().Sub(start).String())

	start = time.Now()
	str02 := BytesToString(bs2)
	log.Println(time.Now().Sub(start).String())

	xtesting.Equal(t, str01, str1)
	xtesting.Equal(t, str02, str2)
}
