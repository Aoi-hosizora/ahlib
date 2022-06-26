package xstring

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"net/url"
	"strings"
	"testing"
	"time"
	"unicode"
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
		{" ", " "},
		{"abc", "Abc"},
		{"abc def", "Abc Def"},
		{"abc\tDef ", "Abc\tDef "},
		{" abc\ndef\r\n　ghi\v", " Abc\nDef\r\n　Ghi\v"},
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
		{" ", " "},
		{"Abc", "abc"},
		{"Abc Def", "abc def"},
		{"abc\tDef ", "abc\tdef "},
		{" abc\nDef\r\n　Ghi\v", " abc\ndef\r\n　ghi\v"},
		{"测试 测试", "测试 测试"},
		{"テス テス", "テス テス"},
		{"Тест Тест", "тест тест"},
	} {
		xtesting.Equal(t, UncapitalizeAll(tc.give), tc.want)
	}
}

func TestIsArabicNumber(t *testing.T) {
	xtesting.True(t, unicode.IsNumber('０'), true)
	xtesting.True(t, unicode.IsNumber('５'), true)
	for _, tc := range []struct {
		give rune
		want bool
	}{
		{' ', false},
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
		{'０', false}, // <- true for unicode.IsNumber
		{'１', false},
		{'２', false},
		{'３', false},
		{'４', false},
		{'５', false},
		{'６', false},
		{'７', false},
		{'８', false},
		{'９', false},
	} {
		t.Run(string(tc.give), func(t *testing.T) {
			xtesting.Equal(t, IsArabicNumber(tc.give), tc.want)
		})
	}
}

func TestIsBlank(t *testing.T) {
	for _, tc := range []struct {
		give rune
		want bool
	}{
		{' ', true},
		{'\t', true},
		{'\n', true},
		{'\r', true},
		{'\v', true},
		{'\f', true},
		{'　', true},
		{'\x85', true},
		{'\xA0', true},
		{0x2000, true},
		{0x2001, true},
		{0x2002, true},
		{0x2003, true},
		{0x2004, true},
		{0x2005, true},
		{0x2006, true},
		{0x2007, true},
		{0x2008, true},
		{0x2009, true},
		{0x200A, true},
		{0x200B, false},
		{0x200F, false},
		{0x2028, true},
		{0x2029, true},
		{0x202F, true},
		{0x205F, true},
		{'a', false},
		{'0', false},
		{'测', false},
	} {
		xtesting.Equal(t, IsBlank(tc.give), tc.want)
	}
}

func TestTrimBlanks(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", ""},
		{"\t", ""},
		{"\n|\r", "|"},
		{"\t\t\r\r\n\n", ""},
		{" a　b\tc\r\nd\v\f", "a　b\tc\r\nd"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, TrimBlanks(tc.give), tc.want)
		})
	}
}

func TestReplaceAndRemoveExtraBlanks(t *testing.T) {
	for _, tc := range []struct {
		give  string
		want1 string
		want2 string
	}{
		{"", "", ""},
		{" ", "", ""},
		{" \n　", "", ""},
		{"\t\v ", "", ""},
		{"a b", "a|b", "a b"},
		{"a  b", "a|b", "a b"},
		{" a　b\tc\r\nd", "a|b|c|d", "a b c d"},
		{"   ab cd　 ef\n\tgh\n", "ab|cd|ef|gh", "ab cd ef gh"},
		{"\t\t\r\r\n\n", "", ""},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, ReplaceExtraBlanks(tc.give, "|"), tc.want1)
			xtesting.Equal(t, RemoveExtraBlanks(tc.give), tc.want2)
		})
	}
}

func TestIsBreakLine(t *testing.T) {
	for _, tc := range []struct {
		give rune
		want bool
	}{
		{' ', false},
		{'\t', false},
		{'\n', true},
		{'\r', true},
		{'\v', false},
		{'\f', false},
		{'a', false},
		{'0', false},
		{'测', false},
	} {
		xtesting.Equal(t, IsBreakLine(tc.give), tc.want)
	}
}

func TestTrimBreakLines(t *testing.T) {
	for _, tc := range []struct {
		give string
		want string
	}{
		{"", ""},
		{" ", " "},
		{"\n", ""},
		{"\n|\r", "|"},
		{"\n\r\t\r\t\n", "\t\r\t"},
		{" a　b\tc\r\nd\r\n", " a　b\tc\r\nd"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, TrimBreakLines(tc.give), tc.want)
		})
	}
}

func TestReplaceAndRemoveExtraBreakLines(t *testing.T) {
	for _, tc := range []struct {
		give  string
		want1 string
		want2 string
	}{
		{"", "", ""},
		{" ", " ", " "},
		{" \n\n　", " |　", " \n　"},
		{"\t\r\n", "\t", "\t"},
		{"a\nb", "a|b", "a\nb"},
		{"a\n\rb", "a|b", "a\nb"},
		{" a\rb\tc\n\nd", " a|b\tc|d", " a\nb\tc\nd"},
		{"  \n ab\t cd　 ef\n\rgh\n", "  | ab\t cd　 ef|gh", "  \n ab\t cd　 ef\ngh"},
		{"\n\r\t\r\t\n", "\t|\t", "\t\n\t"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, ReplaceExtraBreakLines(tc.give, "|"), tc.want1)
			xtesting.Equal(t, RemoveExtraBreakLines(tc.give), tc.want2)
		})
	}
}

func TestSplitToWords(t *testing.T) {
	for _, tc := range []struct {
		giveString string
		giveSeps   []string
		want       []string
	}{
		{"", []string{}, []string{}},
		{"a", []string{}, []string{"a"}},
		{"ABC", []string{}, []string{"ABC"}},

		{"a", []string{"a"}, []string{}},
		{"abc", []string{"a"}, []string{"bc"}},
		{"abc", []string{"b"}, []string{"a", "c"}},
		{"abc", []string{"c"}, []string{"ab"}},
		{"abc", []string{"a", "b"}, []string{"c"}},
		{"abc", []string{"a", "b", "c"}, []string{}},
		{"abc#d", []string{"c", "#"}, []string{"ab", "d"}},

		{" ", []string{}, []string{}},
		{"a b", []string{}, []string{"a", "b"}},
		{" a  b   c    ", []string{}, []string{"a", "b", "c"}},
		{" a  b   c    ", []string{"="}, []string{"a", "b", "c"}},

		{"Abc", []string{}, []string{"Abc"}},
		{"aBc", []string{}, []string{"a", "Bc"}},
		{"abC", []string{}, []string{"ab", "C"}},
		{"aBC", []string{}, []string{"a", "BC"}},
		{"Abc", []string{"="}, []string{"Abc"}},
		{"aBc", []string{"="}, []string{"aBc"}},
		{"abC", []string{"="}, []string{"abC"}},
		{"aBC", []string{"="}, []string{"aBC"}},
		{"Abc", []string{"=", CaseSplitter}, []string{"Abc"}},
		{"aBc", []string{"=", CaseSplitter}, []string{"a", "Bc"}},
		{"abC", []string{"=", CaseSplitter}, []string{"ab", "C"}},
		{"aBC", []string{"=", CaseSplitter}, []string{"a", "BC"}},

		{"ab cd_ef.ghIj-kl", []string{}, []string{"ab", "cd", "ef", "gh", "Ij", "kl"}},
		{"ab cd_ef.ghIj-kl", []string{"="}, []string{"ab", "cd_ef.ghIj-kl"}},
		{"ab cd_ef.ghIj-kl", []string{CaseSplitter}, []string{"ab", "cd_ef.gh", "Ij-kl"}},
		{"ab cd_ef.ghIj-kl", []string{"_", "-"}, []string{"ab", "cd", "ef.ghIj", "kl"}},
		{"ab cd_ef.ghIj-kl", []string{" ", ".", ".", CaseSplitter}, []string{"ab", "cd_ef", "gh", "Ij-kl"}},

		{"!", []string{"!"}, []string{}},
		{"!!a!!", []string{"!"}, []string{"a"}},
		{"aB_c-d.e+f g　h?i", []string{CaseSplitter, "+", "?", "!"}, []string{"a", "B_c-d.e", "f", "g", "h", "i"}},
		{"aB_c-d.e+f g　h?i", []string{CaseSplitter, "_", "-", ".", "+"}, []string{"a", "B", "c", "d", "e", "f", "g", "h?i"}},
		{"тестТест-a", []string{}, []string{"тест", "Тест", "a"}},
		{"тестТест-a", []string{"-"}, []string{"тестТест", "a"}},
		{"测试andテスТестOr", []string{}, []string{"测试andテス", "Тест", "Or"}},
		{"测试 and テス тест or", []string{"and"}, []string{"测试", "テス", "тест", "or"}},
		{"mix-Mix_Mix.Mix?mix", []string{"-", "_", ".", "?"}, []string{"mix", "Mix", "Mix", "Mix", "mix"}},
	} {
		t.Run(tc.giveString, func(t *testing.T) {
			words := SplitToWords(tc.giveString, tc.giveSeps...)
			xtesting.Equal(t, words, tc.want)
		})
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
		{"abCdEF", "AbCdEF", "abCdEF", "ab_cd_ef", "ab-cd-ef"},
		{"AAaaAA_bbBBbb", "AAaaAABbBBbb", "aaaaAABbBBbb", "aaaa_aa_bb_bbbb", "aaaa-aa-bb-bbbb"},
		{"a1a b2B C3c D4D", "A1aB2BC3cD4D", "a1aB2BC3cD4D", "a1a_b2_b_c3c_d4_d", "a1a-b2-b-c3c-d4-d"},
		{"IPv4Address-and-Port", "IPv4AddressAndPort", "ipv4AddressAndPort", "ipv4_address_and_port", "ipv4-address-and-port"},
		{"TestPascalCase", "TestPascalCase", "testPascalCase", "test_pascal_case", "test-pascal-case"},
		{"testCamelCase", "TestCamelCase", "testCamelCase", "test_camel_case", "test-camel-case"},
		{"test_snake_case", "TestSnakeCase", "testSnakeCase", "test_snake_case", "test-snake-case"},
		{"test-kebab-case", "TestKebabCase", "testKebabCase", "test_kebab_case", "test-kebab-case"},
		{"testMixed_Case-Test.Test", "TestMixedCaseTestTest", "testMixedCaseTestTest", "test_mixed_case_test_test", "test-mixed-case-test-test"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, PascalCase(tc.give), tc.wantP)
			xtesting.Equal(t, CamelCase(tc.give), tc.wantC)
			xtesting.Equal(t, SnakeCase(tc.give), tc.wantS)
			xtesting.Equal(t, KebabCase(tc.give), tc.wantK)
		})
	}
}

// showFn represents if tests need to show shuffle and random result.
var showFn = func() bool { return false }

func TestTimeID(t *testing.T) {
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
		uuid := TimeID(tc.giveTime, tc.giveCount)
		if tc.giveCount <= 23 {
			xtesting.Equal(t, uuid, tc.want)
		} else {
			xtesting.Equal(t, len(uuid), tc.giveCount)
			xtesting.Equal(t, uuid[:23], tc.want)

			for i := 0; i < 4; i++ {
				time.Sleep(2 * time.Nanosecond)
				uuid1 := TimeID(tc.giveTime, tc.giveCount)
				time.Sleep(2 * time.Nanosecond)
				uuid2 := TimeID(tc.giveTime, tc.giveCount)
				xtesting.NotEqual(t, uuid1, uuid2)
				if showFn() {
					fmt.Println(uuid1, uuid2)
				}
			}
		}
	}
}

func TestRandString(t *testing.T) {
	for _, tc := range []struct {
		giveFn    func(int) string
		giveCount int
	}{
		{RandCapitalLetterString, 0},
		{RandCapitalLetterString, 5},
		{RandCapitalLetterString, 20},

		{RandLowercaseLetterString, 0},
		{RandLowercaseLetterString, 5},
		{RandLowercaseLetterString, 20},

		{RandLetterString, 0},
		{RandLetterString, 5},
		{RandLetterString, 20},

		{RandNumberString, 0},
		{RandNumberString, 5},
		{RandNumberString, 20},

		{RandCapitalLetterNumberString, 0},
		{RandCapitalLetterNumberString, 5},
		{RandCapitalLetterNumberString, 20},

		{RandLowercaseLetterNumberString, 0},
		{RandLowercaseLetterNumberString, 5},
		{RandLowercaseLetterNumberString, 20},
	} {
		r := tc.giveFn(tc.giveCount)
		if tc.giveCount == 0 {
			xtesting.Equal(t, r, "")
		} else {
			xtesting.Equal(t, len(r), tc.giveCount)

			for i := 0; i < 4; i++ {
				time.Sleep(2 * time.Nanosecond)
				r1 := tc.giveFn(tc.giveCount)
				time.Sleep(2 * time.Nanosecond)
				r2 := tc.giveFn(tc.giveCount)
				xtesting.NotEqual(t, r1, r2)
				if showFn() {
					fmt.Println(r1, r2)
				}
			}
		}
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
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FastStob(s)
		}
	})
	b.Run("ConvertToBytes", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = []byte(s)
		}
	})
}

func BenchmarkFastBtos(b *testing.B) {
	bs := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}

	b.Run("FastBtos", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FastBtos(bs)
		}
	})
	b.Run("ConvertToString", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = string(bs)
		}
	})
}

func TestTrimUTF8XXX(t *testing.T) {
	// Bom
	for _, tc := range []struct {
		giveStr string
		giveBs  []byte
		wantStr string
		wantBs  []byte
	}{
		{"", []byte{}, "", []byte{}},
		{"test", []byte{'t', 'e', 's', 't'}, "test", []byte{'t', 'e', 's', 't'}},
		{"\xef\xbb\xbf", []byte{0xEF, 0xBB, 0xBF}, "", []byte{}},
		{"\xef\xbb\xbftest", []byte{0xEF, 0xBB, 0xBF, 't', 'e', 's', 't'}, "test", []byte{'t', 'e', 's', 't'}},
	} {
		xtesting.Equal(t, TrimUTF8Bom(tc.giveStr), tc.wantStr)
		xtesting.Equal(t, TrimUTF8BomBytes(tc.giveBs), tc.wantBs)
	}

	// Rc
	for _, tc := range []struct {
		giveStr string
		giveBs  []byte
		wantStr string
		wantBs  []byte
	}{
		{"", []byte{}, "", []byte{}},
		{"test", []byte{'t', 'e', 's', 't'}, "test", []byte{'t', 'e', 's', 't'}},
		{"\xef\xbf\xbd", []byte{0xEF, 0xBF, 0xBD}, "", []byte{}},
		{"\xef\xbf\xbdtest", []byte{0xEF, 0xBF, 0xBD, 't', 'e', 's', 't'}, "test", []byte{'t', 'e', 's', 't'}},
	} {
		xtesting.Equal(t, TrimUTF8Replacement(tc.giveStr), tc.wantStr)
		xtesting.Equal(t, TrimUTF8ReplacementBytes(tc.giveBs), tc.wantBs)
	}
}

func TestPadGetLeftRight(t *testing.T) {
	for _, tc := range []struct {
		give       string
		givePad    rune
		giveLength int
		wantPadL   string
		wantPadR   string
		wantGetL   string
		wantGetR   string
		wantL      string
		wantR      string
	}{
		{"", '0', -2, "", "", "", "", "", ""},
		{"", '0', -1, "", "", "", "", "", ""},
		{"", '0', 0, "", "", "", "", "", ""},
		{"", '0', 1, "0", "0", "", "", "0", "0"},
		{"", '0', 2, "00", "00", "", "", "00", "00"},

		{"1", '0', -2, "1", "1", "", "", "", ""},
		{"1", '0', -1, "1", "1", "", "", "", ""},
		{"1", '0', 0, "1", "1", "", "", "", ""},
		{"1", '0', 1, "1", "1", "1", "1", "1", "1"},
		{"1", '0', 2, "01", "10", "1", "1", "01", "10"},
		{"1", '0', 3, "001", "100", "1", "1", "001", "100"},

		{"123", '0', -1, "123", "123", "", "", "", ""},
		{"123", '0', 0, "123", "123", "", "", "", ""},
		{"123", '0', 1, "123", "123", "1", "3", "1", "3"},
		{"123", '0', 2, "123", "123", "12", "23", "12", "23"},
		{"123", '0', 3, "123", "123", "123", "123", "123", "123"},
		{"123", '0', 4, "0123", "1230", "123", "123", "0123", "1230"},
		{"123", '0', 5, "00123", "12300", "123", "123", "00123", "12300"},

		{"test", ' ', -1, "test", "test", "", "", "", ""},
		{"test", ' ', 3, "test", "test", "tes", "est", "tes", "est"},
		{"test", ' ', 4, "test", "test", "test", "test", "test", "test"},
		{"test", ' ', 5, " test", "test ", "test", "test", " test", "test "},
		{"test", ' ', 8, "    test", "test    ", "test", "test", "    test", "test    "},

		{"测试テスтест", '零', 1, "测试テスтест", "测试テスтест", "测", "т", "测", "т"},
		{"测试テスтест", '零', 3, "测试テスтест", "测试テスтест", "测试テ", "ест", "测试テ", "ест"},
		{"测试テスтест", '零', 6, "测试テスтест", "测试テスтест", "测试テスте", "テスтест", "测试テスте", "テスтест"},
		{"测试テスтест", '零', 7, "测试テスтест", "测试テスтест", "测试テスтес", "试テスтест", "测试テスтес", "试テスтест"},
		{"测试テスтест", '零', 8, "测试テスтест", "测试テスтест", "测试テスтест", "测试テスтест", "测试テスтест", "测试テスтест"},
		{"测试テスтест", '零', 9, "零测试テスтест", "测试テスтест零", "测试テスтест", "测试テスтест", "零测试テスтест", "测试テスтест零"},
		{"测试テスтест", '零', 11, "零零零测试テスтест", "测试テスтест零零零", "测试テスтест", "测试テスтест", "零零零测试テスтест", "测试テスтест零零零"},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, PadLeft(tc.give, tc.givePad, tc.giveLength), tc.wantPadL)
			xtesting.Equal(t, PadRight(tc.give, tc.givePad, tc.giveLength), tc.wantPadR)
			xtesting.Equal(t, GetLeft(tc.give, tc.giveLength), tc.wantGetL)
			xtesting.Equal(t, GetRight(tc.give, tc.giveLength), tc.wantGetR)
			xtesting.Equal(t, GetOrPadLeft(tc.give, tc.giveLength, tc.givePad), tc.wantL)
			xtesting.Equal(t, GetOrPadRight(tc.give, tc.giveLength, tc.givePad), tc.wantR)
		})
	}
}

func TestBool(t *testing.T) {
	for _, tc := range []struct {
		giveBool bool
		giveT    string
		giveF    string
		want     string
	}{
		{true, "T", "F", "T"},
		{true, "true", "false", "true"},
		{true, "1", "0", "1"},
		{false, "T", "F", "F"},
		{false, "true", "false", "false"},
		{false, "1", "0", "0"},
	} {
		xtesting.Equal(t, Bool(tc.giveBool, tc.giveT, tc.giveF), tc.want)
	}
}

func TestExtraSpace(t *testing.T) {
	for _, tc := range []struct {
		give       string
		wantLeftS  string
		wantRightS string
		wantLeftB  string
		wantRightB string
	}{
		{"", "", "", "", ""},
		{" ", "  ", "  ", " ", " "},
		{"\t", " \t", "\t ", "\t", "\t"},
		{"\r\n", " \r\n", "\r\n ", "\r\n", "\r\n"},
		{"a", " a", "a ", " a", "a "},
		{"  a|\n\n\n ", "   a|\n\n\n ", "  a|\n\n\n  ", " a|", "a| "},
		{"\r\ta|　", " \r\ta|　", "\r\ta|　 ", " a|", "a| "},
	} {
		t.Run(tc.give, func(t *testing.T) {
			xtesting.Equal(t, ExtraSpaceOnLeftIfNotEmpty(tc.give), tc.wantLeftS)
			xtesting.Equal(t, ExtraSpaceOnRightIfNotEmpty(tc.give), tc.wantRightS)
			xtesting.Equal(t, ExtraSpaceOnLeftIfNotBlank(tc.give), tc.wantLeftB)
			xtesting.Equal(t, ExtraSpaceOnRightIfNotBlank(tc.give), tc.wantRightB)
		})
	}
}

func TestMaskToken(t *testing.T) {
	for _, tc := range []struct {
		giveString  string
		giveIndices []int
		want        string
		wantR       string
	}{
		{"", []int{}, "", ""},
		{"", []int{-2}, "", ""},
		{"", []int{-1}, "", ""},
		{"", []int{0}, "", ""},
		{"", []int{1}, "", ""},
		{"", []int{2}, "", ""},

		{"a", []int{}, "a", "*"},
		{"a", []int{-3}, "a", "*"},
		{"a", []int{-2}, "a", "*"},
		{"a", []int{-1}, "*", "a"}, // <<<
		{"a", []int{0}, "*", "a"},  // <<<
		{"a", []int{1}, "a", "*"},
		{"a", []int{2}, "a", "*"},
		{"a", []int{3}, "a", "*"},
		{"a", []int{0, 1}, "*", "a"},
		{"a", []int{-1, -2}, "*", "a"},
		{"a", []int{0, -1}, "*", "a"}, // <<<

		{"aa", []int{}, "aa", "**"},
		{"aa", []int{-3}, "aa", "**"},
		{"aa", []int{-2}, "*a", "a*"}, // <<<
		{"aa", []int{-1}, "a*", "*a"}, // <<<
		{"aa", []int{0}, "*a", "a*"},  // <<<
		{"aa", []int{1}, "a*", "*a"},  // <<<
		{"aa", []int{2}, "aa", "**"},
		{"aa", []int{3}, "aa", "**"},
		{"aa", []int{-3, 0}, "*a", "a*"},
		{"aa", []int{0, 1}, "**", "aa"}, // <<<
		{"aa", []int{1, 2}, "a*", "*a"},
		{"aa", []int{-1, -2}, "**", "aa"}, // <<<
		{"aa", []int{-2, -3}, "*a", "a*"},
		{"aa", []int{0, -1}, "**", "aa"}, // <<<

		{"aaa", []int{}, "aaa", "***"},
		{"aaa", []int{0}, "*aa", "a**"},
		{"aaa", []int{1}, "a*a", "*a*"},
		{"aaa", []int{2}, "aa*", "**a"},
		{"aaa", []int{-1}, "aa*", "**a"},
		{"aaa", []int{-2}, "a*a", "*a*"},
		{"aaa", []int{-3}, "*aa", "a**"},
		{"aaa", []int{0, 1}, "**a", "aa*"},
		{"aaa", []int{-1, -3}, "*a*", "a*a"},
		{"aaa", []int{1, -1}, "a**", "*aa"},

		{"pwd123abcpwd", []int{3, 4, 5, 6, 7, 8, 20}, "pwd******pwd", "***123abc***"},
		{"pwd123abcpwd", []int{0, 1, 2, 9, 10, 11, 20}, "***123abc***", "pwd******pwd"},
		{"pwd123abcpwd", []int{-3, -4, -5, -6, -7, -8, -20}, "pwd1******wd", "****23abcp**"},
		{"pwd123abcpwd", []int{0, -1, -2, -9, -10, -11, -20}, "****23abcp**", "pwd1******wd"},
	} {
		xtesting.Equal(t, MaskToken(tc.giveString, '*', tc.giveIndices...), tc.want)
		xtesting.Equal(t, MaskTokenR(tc.giveString, '*', tc.giveIndices...), tc.wantR)
		xtesting.Equal(t, StringMaskToken(tc.giveString, "#$%", tc.giveIndices...), strings.ReplaceAll(tc.want, "*", "#$%"))
		xtesting.Equal(t, StringMaskTokenR(tc.giveString, "#$%", tc.giveIndices...), strings.ReplaceAll(tc.wantR, "*", "#$%"))
	}
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

func TestStringSliceToStringMap(t *testing.T) {
	type Map = map[string]string
	for _, tc := range []struct {
		give []string
		want Map
	}{
		{nil, Map{}},
		{[]string{}, Map{}},
		{[]string{""}, Map{}},
		{[]string{"xxx"}, Map{}},

		{[]string{"a", "b"}, Map{"a": "b"}},
		{[]string{"a", "b", "c"}, Map{"a": "b"}},
		{[]string{"a", "", "b", "c", "d"}, Map{"a": "", "b": "c"}},
		{[]string{"a", "", "a", "b", "c"}, Map{"a": "b"}},
		{[]string{"", "", "", " ", " ", "", "x"}, Map{"": " ", " ": ""}},
	} {
		xtesting.Equal(t, StringSliceToMap(tc.give), tc.want)
	}
}

func TestSliceToStringMap(t *testing.T) {
	type Map = map[string]interface{}
	for _, tc := range []struct {
		give []interface{}
		want Map
	}{
		{nil, Map{}},
		{[]interface{}{}, Map{}},
		{[]interface{}{1}, Map{}},
		{[]interface{}{nil}, Map{}},

		{[]interface{}{"1", 2}, Map{"1": 2}},
		{[]interface{}{1, uint32(2)}, Map{"1": uint32(2)}},
		{[]interface{}{uint(1), 2.3}, Map{"1": 2.3}},
		{[]interface{}{true, 2i}, Map{"true": 2i}},
		{[]interface{}{1 + 2i, false}, Map{"(1+2i)": false}},
		{[]interface{}{[]byte("key"), []byte("value")}, Map{"key": []byte("value")}},

		{[]interface{}{nil, 2}, Map{}},
		{[]interface{}{1, 2, 3}, Map{"1": 2}},
		{[]interface{}{nil, 2, 3}, Map{"2": 3}},
		{[]interface{}{1, 2, nil, 3, 4}, Map{"1": 2, "3": 4}},
		{[]interface{}{1, 2, 1, 3, 1, 4, 2, 3}, Map{"1": 4, "2": 3}},
		{[]interface{}{1, "2", []byte("3"), 4.4}, Map{"1": "2", "3": 4.4}},
		{[]interface{}{true, 2, 3.3, true}, Map{"true": 2, "3.3": true}},
	} {
		xtesting.Equal(t, SliceToStringMap(tc.give), tc.want)
	}
}

func TestSemanticVersion(t *testing.T) {
	for _, tc := range []struct {
		give      string
		wantP1    uint64
		wantP2    uint64
		wantP3    uint64
		wantError bool
	}{
		{"", 0, 0, 0, true},
		{".1", 0, 0, 0, true},
		{"v", 0, 0, 0, true},
		{"vw", 0, 0, 0, true},
		{"w1", 0, 0, 0, true},
		{"0v", 0, 0, 0, true},
		{"0.v", 0, 0, 0, true},
		{"0.0.v", 0, 0, 0, true},
		{"0v", 0, 0, 0, true},
		{"v0.v", 0, 0, 0, true},
		{"v0.v.0", 0, 0, 0, true},

		{"1", 1, 0, 0, false},
		{"1.2", 1, 2, 0, false},
		{"1.2.3", 1, 2, 3, false},
		{"v1", 1, 0, 0, false},
		{"v1.2", 1, 2, 0, false},
		{"v1.2.3", 1, 3, 3, false},
	} {
		t.Run(tc.give, func(t *testing.T) {
			p1, p2, p3, err := SemanticVersion(tc.give)
			xtesting.Equal(t, err != nil, tc.wantError)
			if err != nil {
				xtesting.Equal(t, p1, tc.wantP1)
				xtesting.Equal(t, p2, tc.wantP2)
				xtesting.Equal(t, p3, tc.wantP3)
			}
		})
	}
}
