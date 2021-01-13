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

func TestTimeUUID(t *testing.T) {
	log.Println(TimeUUID(time.Now(), 5))
	log.Println(TimeUUID(time.Now(), 24))
	log.Println(TimeUUID(time.Now(), 30))
}

func TestRandLetterNumberString(t *testing.T) {
	log.Println(RandCapitalLetterString(20))
	log.Println(RandCapitalLetterString(20))
	log.Println(RandLowercaseLetterString(20))
	log.Println(RandLowercaseLetterString(20))
	log.Println(RandLetterString(20))
	log.Println(RandLetterString(20))
	log.Println(RandNumberString(20))
	log.Println(RandNumberString(20))
	log.Println(RandCapitalLetterNumberString(20))
	log.Println(RandCapitalLetterNumberString(20))
	log.Println(RandLowercaseLetterNumberString(20))
	log.Println(RandLowercaseLetterNumberString(20))
}

// func TestToSnakeCase(t *testing.T) {
// 	xtesting.Equal(t, ToSnakeCase(""), "")
// 	xtesting.Equal(t, ToSnakeCase("AoiHosizora"), "aoi_hosizora")
// 	xtesting.Equal(t, ToSnakeCase("abc0d1EdF"), "abc0d1_ed_f")
// 	xtesting.Equal(t, ToSnakeCase("私達isわたしたち"), "私達isわたしたち")
// 	xtesting.Equal(t, ToSnakeCase("a bC"), "a_b_c")
// }

func TestRemoveSpaces(t *testing.T) {
	xtesting.Equal(t, RemoveBlanks(""), "")
	xtesting.Equal(t, RemoveBlanks("a b  c 　d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveBlanks("a b 	 c d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveBlanks("a b \n	 c d   e f"), "a b c d e f")
	xtesting.Equal(t, RemoveBlanks("\n　"), "")
	xtesting.Equal(t, RemoveBlanks("　\n	"), "")
}

func TestDefaultMaskToken(t *testing.T) {
	xtesting.Equal(t, DefaultMaskToken(""), "")
	xtesting.Equal(t, DefaultMaskToken(" "), "*")
	xtesting.Equal(t, DefaultMaskToken("a"), "*")
	xtesting.Equal(t, DefaultMaskToken("aa"), "**")
	xtesting.Equal(t, DefaultMaskToken("aaa"), "**a")
	xtesting.Equal(t, DefaultMaskToken("aaaa"), "a**a")
	xtesting.Equal(t, DefaultMaskToken("aaaaa"), "a***a")
	xtesting.Equal(t, DefaultMaskToken("aaaaaa"), "aa**aa")
	xtesting.Equal(t, DefaultMaskToken("あ"), "*")
	xtesting.Equal(t, DefaultMaskToken("あa"), "**")
	xtesting.Equal(t, DefaultMaskToken("あaa"), "**a")
	xtesting.Equal(t, DefaultMaskToken("あaaa"), "あ**a")
	xtesting.Equal(t, DefaultMaskToken("あaaaa"), "あ***a")
	xtesting.Equal(t, DefaultMaskToken("あaaaaa"), "あa**aa")
}

func TestStringToBytes(t *testing.T) {
	xtesting.Equal(t, FastStob(""), []byte{})
	xtesting.Equal(t, FastStob("abcdefg"), []byte("abcdefg"))

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
	bs02 := FastStob(str2)
	log.Println(time.Now().Sub(start).String())

	xtesting.Equal(t, bs01, bs1)
	xtesting.Equal(t, bs02, bs2)
}

func TestBytesToString(t *testing.T) {
	xtesting.Equal(t, FastBtos(nil), "")
	xtesting.Equal(t, FastBtos([]byte{}), "")
	xtesting.Equal(t, FastBtos([]byte("abcdefg")), "abcdefg")

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
	str02 := FastBtos(bs2)
	log.Println(time.Now().Sub(start).String())

	xtesting.Equal(t, str01, str1)
	xtesting.Equal(t, str02, str2)
}

func TestEncodeUrlValues(t *testing.T) {
	m1 := map[string][]string{}
	xtesting.Equal(t, EncodeUrlValues(m1, nil), "")

	m2 := map[string][]string{"a": {}}
	xtesting.Equal(t, EncodeUrlValues(m2, nil), "")

	m3 := map[string][]string{"a": {"b"}}
	xtesting.Equal(t, EncodeUrlValues(m3, nil), "a=b")

	m4 := map[string][]string{"a": {"b"}, "b": {"c"}}
	xtesting.Equal(t, EncodeUrlValues(m4, nil), "a=b&b=c")

	m5 := map[string][]string{"a": {"a1", "a2", "a3"}, "b": {"b1", "b2", "b3"}}
	xtesting.Equal(t, EncodeUrlValues(m5, nil), "a=a1&a=a2&a=a3&b=b1&b=b2&b=b3")
}

func TestIsMask(t *testing.T) {
	xtesting.False(t, IsMark([]rune(" ")[0]))
	xtesting.False(t, IsMark([]rune("a")[0]))
	xtesting.True(t, IsMark([]rune("\u0301")[0]))
}

func TestPadLeftRight(t *testing.T) {
	xtesting.Equal(t, PadLeft("test", '0', 4), "test")
	xtesting.Equal(t, PadLeft("test", '0', 5), "0test")
	xtesting.Equal(t, PadLeft("test", '0', 6), "00test")
	xtesting.Equal(t, PadLeft("测试テスト", '零', 6), "零测试テスト")

	xtesting.Equal(t, PadRight("test", '0', 4), "test")
	xtesting.Equal(t, PadRight("test", '0', 5), "test0")
	xtesting.Equal(t, PadRight("test", '0', 6), "test00")
	xtesting.Equal(t, PadRight("测试テスト", '零', 6), "测试テスト零")
}

func TestGetLeftRight(t *testing.T) {
	xtesting.Equal(t, GetLeft("", 0), "")
	xtesting.Equal(t, GetLeft("", 3), "")
	xtesting.Equal(t, GetLeft("123", 0), "")
	xtesting.Equal(t, GetLeft("123", 3), "123")
	xtesting.Equal(t, GetLeft("1234", 3), "123")
	xtesting.Equal(t, GetLeft("测试テスト", 3), "测试テ")

	xtesting.Equal(t, GetRight("", 0), "")
	xtesting.Equal(t, GetRight("", 3), "")
	xtesting.Equal(t, GetRight("123", 0), "")
	xtesting.Equal(t, GetRight("123", 3), "123")
	xtesting.Equal(t, GetRight("1234", 3), "234")
	xtesting.Equal(t, GetRight("测试テスト", 3), "テスト")
}


func TestSplitAndGet(t *testing.T) {
	xtesting.Equal(t, SplitAndGet("", "", 0), "")
	xtesting.Equal(t, SplitAndGet(" ", " ", 0), "")
	xtesting.Equal(t, SplitAndGet("a b", "", 0), "a")
	xtesting.Equal(t, SplitAndGet("a b", "", -1), "b")
	xtesting.Equal(t, SplitAndGet("a b", " ", 0), "a")
	xtesting.Equal(t, SplitAndGet("a b", " ", 1), "b")
	xtesting.Equal(t, SplitAndGet("a b", " ", -1), "b")
	xtesting.Equal(t, SplitAndGet("a b", " ", -2), "a")
	xtesting.Panic(t, func() { SplitAndGet("a b", " ", 2) })
}
