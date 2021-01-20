package xstring

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestCapitalize(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", " "},
		{"abc", "Abc"},
		{"Abc", "Abc"},
		{"abc def", "Abc def"},
		{"测试", "测试"},
		{"テス", "テス"},
		{"тест", "Тест"},
	} {
		xtesting.Equal(t, Capitalize(tc.give), tc.want)
	}
}

func TestUncapitalize(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", " "},
		{"Abc", "abc"},
		{"abc", "abc"},
		{"Abc Def", "abc Def"},
		{"测试", "测试"},
		{"テス", "テス"},
		{"Тест", "тест"},
	} {
		xtesting.Equal(t, Uncapitalize(tc.give), tc.want)
	}
}

func TestCapitalizeAll(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", ""},
		{"abc", "Abc"},
		{"abc def", "Abc Def"},
		{"abc Def ", "Abc Def"},
		{"测试 测试", "测试 测试"},
		{"テス テス", "テス テス"},
		{"тест тест", "Тест Тест"},
	} {
		xtesting.Equal(t, CapitalizeAll(tc.give), tc.want)
	}
}

func TestUncapitalizeAll(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", ""},
		{"Abc", "abc"},
		{"Abc Def", "abc def"},
		{"abc Def ", "abc def"},
		{"测试 测试", "测试 测试"},
		{"テス テス", "テス テス"},
		{"Тест Тест", "тест тест"},
	} {
		xtesting.Equal(t, UncapitalizeAll(tc.give), tc.want)
	}
}

func TestRemoveBlanks(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", ""},
		{" \n　", ""},
		{"\t\v ", ""},
		{"a b", "a b"},
		{"a  b", "a b"},
		{" a　b\tc\r\nd", "a b c d"},
		{"ab cd　 ef\n\tgh\n", "ab cd ef gh"},
		{"\t\t\r\r\n\n", ""},
	} {
		xtesting.Equal(t, RemoveBlanks(tc.give), tc.want)
	}
}

func TestCaseHelper(t *testing.T) {
	for _, tc := range []struct {
		giveString string
		giveSeps   []string
		want       []string
	}{
		{"", []string{}, []string{}},
		{" ", []string{}, []string{}},
		{" a ", []string{}, []string{"a"}},
		{"A", []string{}, []string{"a"}},
		{"a", []string{"a"}, []string{}},
		{"AaA", []string{"a"}, []string{"a", "a"}},
		{"AaA", []string{"A"}, []string{"a"}},
		{"ABcdEFghIJ", []string{}, []string{"abcd", "efgh", "ij"}},
		{"abCDefGHij", []string{}, []string{"ab", "cdef", "ghij"}},
		{"!", []string{"!"}, []string{}},
		{"!!a!!", []string{"!"}, []string{"a"}},
		{"aB_c-d.e+f g　h?i", []string{"+", "?", "!"}, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}},
		{"тестТест", []string{}, []string{"тесттест"}},
		{"测试andテスandтест", []string{}, []string{"测试andテスandтест"}},
		{"测试 and テス and тест", []string{"and"}, []string{"测试", "テス", "тест"}},
		{"mix-Mix_Mix.Mix?Mix", []string{"?"}, []string{"mix", "mix", "mix", "mix", "mix"}},
	} {
		words := caseHelper(tc.giveString, tc.giveSeps...)
		xtesting.Equal(t, words, tc.want)
	}
}

func TestXXXCase(t *testing.T) {
	for _, tc := range []struct {
		give  string
		wantP string
		wantC string
		wantS string
		wantK string
	}{
		{"", "", "", "", ""},
		{"a", "A", "a", "a", "a"},
		{"A", "A", "a", "a", "a"},
		{"abc", "Abc", "abc", "abc", "abc"},
		{"a b", "AB", "aB", "a_b", "a-b"},
		{"ab cd", "AbCd", "abCd", "ab_cd", "ab-cd"},
		{"TestPascalCase", "TestPascalCase", "testPascalCase", "test_pascal_case", "test-pascal-case"},
		{"testCamelCase", "TestCamelCase", "testCamelCase", "test_camel_case", "test-camel-case"},
		{"test_snake_case", "TestSnakeCase", "testSnakeCase", "test_snake_case", "test-snake-case"},
		{"test-kebab-case", "TestKebabCase", "testKebabCase", "test_kebab_case", "test-kebab-case"},
		{"testMixed_Case-Test.Test", "TestMixedCaseTestTest", "testMixedCaseTestTest", "test_mixed_case_test_test", "test-mixed-case-test-test"},
	} {
		xtesting.Equal(t, PascalCase(tc.give), tc.wantP)
		xtesting.Equal(t, CamelCase(tc.give), tc.wantC)
		xtesting.Equal(t, SnakeCase(tc.give), tc.wantS)
		xtesting.Equal(t, KebabCase(tc.give), tc.wantK)
	}
}

func TestTimeUUID(t *testing.T) {
	zero := time.Time{}
	now := time.Date(2021, time.Month(1), 18, 15, 07, 25, 123456789, time.UTC)
	for _, tc := range []struct {
		giveTime  time.Time
		giveCount int
		want      string
	}{
		{now, 0, ""},
		{now, 1, "2"},
		{zero, 4, "0001"},
		{now, 5, "20210"},
		{now, 8, "20210118"},
		{zero, 8, "00010101"},
		{now, 14, "20210118150725"},
		{now, 23, "20210118150725123456789"},
		{zero, 23, "00010101000000000000000"},
		{now, 30, "20210118150725123456789"}, // 25
		{now, 38, "20210118150725123456789"},
	} {
		uuid := TimeUUID(tc.giveTime, tc.giveCount)
		if tc.giveCount <= 23 {
			xtesting.Equal(t, uuid, tc.want)
		} else {
			xtesting.Equal(t, len(uuid), tc.giveCount)
			xtesting.Equal(t, uuid[:23], tc.want)

			for i := 0; i < 4; i++ {
				uuid1 := TimeUUID(tc.giveTime, tc.giveCount)
				time.Sleep(2 * time.Nanosecond)
				uuid2 := TimeUUID(tc.giveTime, tc.giveCount)
				xtesting.NotEqual(t, uuid1, uuid2)
				log.Println(uuid1[23:], uuid2[23:])
			}
		}
	}
}

func TestRandXXXString(t *testing.T) {
	for _, tc := range []struct {
		giveFn    func(int) string
		giveCount int
		show      bool
	}{
		{RandCapitalLetterString, 0, false},
		{RandCapitalLetterString, 5, false},
		{RandCapitalLetterString, 20, true},

		{RandLowercaseLetterString, 0, false},
		{RandLowercaseLetterString, 5, false},
		{RandLowercaseLetterString, 20, true},

		{RandLetterString, 0, false},
		{RandLetterString, 5, false},
		{RandLetterString, 20, true},

		{RandNumberString, 0, false},
		{RandNumberString, 5, false},
		{RandNumberString, 20, true},

		{RandCapitalLetterNumberString, 0, false},
		{RandCapitalLetterNumberString, 5, false},
		{RandCapitalLetterNumberString, 20, true},

		{RandLowercaseLetterNumberString, 0, false},
		{RandLowercaseLetterNumberString, 5, false},
		{RandLowercaseLetterNumberString, 20, true},
	} {
		r := tc.giveFn(tc.giveCount)
		if tc.giveCount == 0 {
			xtesting.Equal(t, r, "")
		} else {
			xtesting.Equal(t, len(r), tc.giveCount)

			for i := 0; i < 8; i++ {
				r1 := tc.giveFn(tc.giveCount)
				time.Sleep(2 * time.Nanosecond)
				r2 := tc.giveFn(tc.giveCount)
				xtesting.NotEqual(t, r1, r2)

				if tc.show {
					log.Println(r1, r2)
				}
			}
			if tc.show {
				fmt.Println()
			}
		}
	}
}

func TestDefaultMaskToken(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", "*"},

		{"a", "*"},                     // 1
		{"aa", "**"},                   // 2
		{"aaa", "**a"},                 // 3
		{"aaaa", "***a"},               // 4
		{"aaaaa", "a***a"},             // 5
		{"aaaaaa", "a****a"},           // 6
		{"aaaaaaa", "a****aa"},         // 7
		{"aaaaaaaa", "a*****aa"},       // 8
		{"aaaaaaaaa", "aa*****aa"},     // 9 <<<
		{"aaaaaaaaaa", "aa******aa"},   // 10
		{"aaaaaaaaaaa", "aa*******aa"}, // 11

		{"テ", "*"},
		{"テス", "**"},
		{"テスa", "**a"},
		{"テスaa", "***a"},
		{"テスaaa", "テ***a"},
		{"テスaaaa", "テ****a"},
		{"テスaaaaa", "テ****aa"},
		{"テスaaaaaa", "テ*****aa"},
		{"テスaaaaaaa", "テス*****aa"},
		{"テスaaaaaaaa", "テス******aa"},
		{"テスaaaaaaaaa", "テス*******aa"},
	} {
		xtesting.Equal(t, DefaultMaskToken(tc.give), tc.want)
	}
}

func TestMaskToken(t *testing.T) {
	for _, tc := range []struct {
		giveString  string
		giveIndices []int
		want        string
	}{
		{"", []int{1}, ""},

		{"a", []int{}, "a"},
		{"a", []int{0}, "*"}, // <<<
		{"a", []int{1}, "a"},
		{"a", []int{-1}, "*"}, // <<<
		{"a", []int{-2}, "a"},
		{"a", []int{0, 1}, "*"},
		{"a", []int{0, -1}, "*"},

		{"aa", []int{}, "aa"},
		{"aa", []int{0}, "*a"}, // <<<
		{"aa", []int{1}, "a*"}, // <<<
		{"aa", []int{2}, "aa"},
		{"aa", []int{-1}, "a*"}, // <<<
		{"aa", []int{-2}, "*a"}, // <<<
		{"aa", []int{-3}, "aa"},
		{"aa", []int{0, 1}, "**"},
		{"aa", []int{-1, -2}, "**"},

		{"aaa", []int{}, "aaa"},
		{"aaa", []int{0}, "*aa"},  // <<<
		{"aaa", []int{1}, "a*a"},  // <<<
		{"aaa", []int{2}, "aa*"},  // <<<
		{"aaa", []int{-1}, "aa*"}, // <<<
		{"aaa", []int{-2}, "a*a"}, // <<<
		{"aaa", []int{-3}, "*aa"}, // <<<
		{"aaa", []int{0, 1}, "**a"},
		{"aaa", []int{-1, -3}, "*a*"},
		{"aaa", []int{1, -1}, "a**"},
	} {
		xtesting.Equal(t, MaskToken(tc.giveString, '*', tc.giveIndices...), tc.want)
	}
}

func TestFastStob(t *testing.T) {
	for _, tc := range []struct {
		give string
		want []byte
	}{
		{"", []byte{}},
		{"a", []byte{'a'}},
		{"hello", []byte{'h', 'e', 'l', 'l', 'o'}},
		{"a b c", []byte{'a', ' ', 'b', ' ', 'c'}},
		{"测试", []byte("测试")},
		{"テス", []byte("テス")},
		{"тест", []byte("тест")},
	} {
		xtesting.Equal(t, FastStob(tc.give), tc.want)
	}
}

func TestFastBtos(t *testing.T) {
	for _, tc := range []struct {
		give []byte
		want string
	}{
		{[]byte{}, ""},
		{[]byte{'a'}, "a"},
		{[]byte{'h', 'e', 'l', 'l', 'o'}, "hello"},
		{[]byte{'a', ' ', 'b', ' ', 'c'}, "a b c"},
		{[]byte("测试"), "测试"},
		{[]byte("テス"), "テス"},
		{[]byte("тест"), "тест"},
	} {
		xtesting.Equal(t, FastBtos(tc.give), tc.want)
	}
}

func BenchmarkFastStob(b *testing.B) {
	s := "hello world"

	b.Run("FastStob", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FastStob(s)
		}
	})
	b.Run("ConvertToBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = []byte(s)
		}
	})
}

func BenchmarkFastBtos(b *testing.B) {
	bs := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}

	b.Run("FastBtos", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FastBtos(bs)
		}
	})
	b.Run("ConvertToString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(bs)
		}
	})
}

func TestEncodeUrlValues(t *testing.T) {
	for _, tc := range []struct {
		give           map[string][]string
		giveEscapeFunc func(string) string
		want           string
	}{
		{map[string][]string{}, nil, ""},
		{map[string][]string{"a": {}}, nil, ""},
		{map[string][]string{"a": {""}}, nil, "a="},
		{map[string][]string{"": {"a"}}, nil, "=a"},
		{map[string][]string{"a": {"", "", ""}}, nil, "a=&a=&a="},

		{map[string][]string{"a": {"b"}, "c": {}}, nil, "a=b"},
		{map[string][]string{"a": {"b"}, "c": {"d"}}, nil, "a=b&c=d"},
		{map[string][]string{"c": {"e", "d"}, "a": {"b"}}, nil, "a=b&c=e&c=d"},
		{map[string][]string{"a": {"a1", "a2", "a3"}, "b": {"b1", "b2", "b3"}}, nil, "a=a1&a=a2&a=a3&b=b1&b=b2&b=b3"},

		{map[string][]string{"q": {"a+b"}, "order": {"-true"}, "%": {"?"}}, url.QueryEscape, "%25=%3F&order=-true&q=a%2Bb"},
		{map[string][]string{"test": {"测试", "テス", "тест"}}, url.QueryEscape, "test=%E6%B5%8B%E8%AF%95&test=%E3%83%86%E3%82%B9&test=%D1%82%D0%B5%D1%81%D1%82"},
		{map[string][]string{"q": {"a+b"}, "order": {"-true"}, "%": {"?"}}, url.PathEscape, "%25=%3F&order=-true&q=a+b"},
		{map[string][]string{"test": {"测试", "テス", "тест"}}, url.PathEscape, "test=%E6%B5%8B%E8%AF%95&test=%E3%83%86%E3%82%B9&test=%D1%82%D0%B5%D1%81%D1%82"},
		{map[string][]string{"a": {"b"}, "c": {"d"}}, func(s string) string { return "?" }, "?=?&?=?"},
	} {
		xtesting.Equal(t, EncodeUrlValues(tc.give, tc.giveEscapeFunc), tc.want)
	}
}

func TestPadLeftRight(t *testing.T) {
	for _, tc := range []struct {
		give       string
		giveChar   rune
		giveLength int
		wantLeft   string
		wantRight  string
	}{
		{"", '0', -1, "", ""},
		{"", '0', 0, "", ""},
		{"", '0', 1, "0", "0"},
		{"", '0', 2, "00", "00"},
		{"1", '0', -1, "1", "1"},
		{"1", '0', 0, "1", "1"},
		{"1", '0', 1, "1", "1"},
		{"1", '0', 2, "01", "10"},
		{"1", '0', 3, "001", "100"},

		{"test", '0', -1, "test", "test"},
		{"test", '0', 3, "test", "test"},
		{"test", '0', 4, "test", "test"},
		{"test", '0', 5, "0test", "test0"},
		{"test", '0', 6, "00test", "test00"},

		{"测试テスтест", '零', 7, "测试テスтест", "测试テスтест"},
		{"测试テスтест", '零', 8, "测试テスтест", "测试テスтест"},
		{"测试テスтест", '零', 9, "零测试テスтест", "测试テスтест零"},
		{"测试テスтест", '零', 10, "零零测试テスтест", "测试テスтест零零"},
	} {
		xtesting.Equal(t, PadLeft(tc.give, tc.giveChar, tc.giveLength), tc.wantLeft)
		xtesting.Equal(t, PadRight(tc.give, tc.giveChar, tc.giveLength), tc.wantRight)
	}
}

func TestGetLeftRight(t *testing.T) {
	for _, tc := range []struct {
		give       string
		giveLength int
		wantLeft   string
		wantRight  string
	}{
		{"", -1, "", ""},
		{"", 0, "", ""},
		{"", 3, "", ""},

		{"123", -1, "", ""},
		{"123", 0, "", ""},
		{"123", 1, "1", "3"},
		{"123", 2, "12", "23"},
		{"123", 3, "123", "123"},
		{"123", 4, "123", "123"},
		{"1234", -1, "", ""},
		{"1234", 3, "123", "234"},
		{"1234", 4, "1234", "1234"},
		{"1234", 5, "1234", "1234"},

		{"测试テス", 3, "测试テ", "试テス"},
		{"测试テス", 6, "测试テス", "测试テス"},
		{"测试テスтест", 5, "测试テスт", "スтест"},
		{"测试テスтест", 10, "测试テスтест", "测试テスтест"},
	} {
		xtesting.Equal(t, GetLeft(tc.give, tc.giveLength), tc.wantLeft)
		xtesting.Equal(t, GetRight(tc.give, tc.giveLength), tc.wantRight)
	}
}

func TestSplitAndGet(t *testing.T) {
	for _, tc := range []struct {
		give      string
		giveSep   string
		giveIndex int
		want      string
		wantPanic bool
	}{
		{"", "", 0, "", true},
		{"", "", 1, "", true},
		{"", "", -1, "", true},
		{" ", "", 0, " ", false},
		{" ", "", -1, " ", false},

		{"a b", "", 0, "a", false},
		{"a b", "", 1, " ", false},
		{"a b", "", 2, "b", false},
		{"a b", "", 3, "", true},
		{"a b", "", -1, "b", false},
		{"a b", "", -2, " ", false},
		{"a b", "", -3, "a", false},
		{"a b", "", -4, "", true},

		{"a b", " ", 0, "a", false},
		{"a b", " ", 1, "b", false},
		{"a b", " ", 2, "", true},
		{"a b", " ", -1, "b", false},
		{"a b", " ", -2, "a", false},
		{"a b", " ", -3, "", true},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { SplitAndGet(tc.give, tc.giveSep, tc.giveIndex) })
		} else {
			xtesting.Equal(t, SplitAndGet(tc.give, tc.giveSep, tc.giveIndex), tc.want)
		}
	}
}
