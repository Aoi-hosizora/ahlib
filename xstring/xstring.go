package xstring

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// ==================
// capitalize related
// ==================

// Capitalize capitalizes the first letter of the whole string.
func Capitalize(s string) string {
	for i, v := range s {
		f := string(unicode.ToUpper(v))
		return f + s[i+len(f):]
	}
	return ""
}

// Uncapitalize uncapitalizes the first letter of the whole string.
func Uncapitalize(s string) string {
	for i, v := range s {
		f := string(unicode.ToLower(v))
		return f + s[i+len(f):]
	}
	return ""
}

// CapitalizeAll capitalizes all the first letter in words of the whole string, words are split by blank character.
func CapitalizeAll(s string) string {
	wordStart := true
	sb := strings.Builder{}
	sb.Grow(len(s))
	for _, v := range s {
		if wordStart {
			wordStart = false
			sb.WriteRune(unicode.ToUpper(v))
		} else {
			sb.WriteRune(v)
		}
		wordStart = IsBlank(v)
	}
	return sb.String()
}

// UncapitalizeAll uncapitalizes all the first letter in words of the whole string, words are split by blank character.
func UncapitalizeAll(s string) string {
	wordStart := true
	sb := strings.Builder{}
	sb.Grow(len(s))
	for _, v := range s {
		if wordStart {
			wordStart = false
			sb.WriteRune(unicode.ToLower(v))
		} else {
			sb.WriteRune(v)
		}
		wordStart = IsBlank(v)
	}
	return sb.String()
}

// ========================
// number and blank related
// ========================

// IsArabicNumber returns true if given rune is an arabic number, that is only "0", "1", "2", "3", "4", "5", "6", "7", "8", "9". Note that it is different with
// unicode.IsNumber.
func IsArabicNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsBlank checks whether given rune is a blank (including space, CR, LF, tabulation, ideographic space, etc.). In the Latin-1 space, a blank is [\t\n\v\f\r\x20\x85\xA0].
// Note that this equals to unicode.IsSpace, see IsBlank's source code for more details.
func IsBlank(r rune) bool {
	// https://compart.com/en/unicode/U+2000
	// 0x09 => '\t' / Tabulation
	// 0x0A => '\n' / Line Feed / LF
	// 0x0B => '\v' / Line Tabulation
	// 0x0C => '\f' / Form Feed
	// 0x0D => '\r' / Carriage Return / CR
	// 0x20 => ' ' / Space
	// 0x85 => Next Line / NEL
	// 0xA0 => No-Break Space / NBSP
	// 0x2000 => En Quad
	// 0x2001 => Em Quad
	// 0x2002 => En Space / ENSP
	// 0x2003 => Em Space / EMSP
	// 0x2004 => Three-Per-Em Space
	// 0x2005 => Four-Per-Em Space
	// 0x2006 => Six-Per-Em Space
	// 0x2007 => Figure Space
	// 0x2008 => Punctuation Space
	// 0x2009 => Thin Space
	// 0x200A => Hair Space
	// 0x2028 => Line Separator
	// 0x2029 => Paragraph Separator
	// 0x202F => Narrow No-Break Space / NNBSP
	// 0x205F => Medium Mathematical Space / MMSP
	// 0x3000 => '　' / Ideographic Space
	return unicode.IsSpace(r)
}

// TrimBlanks trims blanks left and right of the string. Note that this equals to strings.TrimSpace.
func TrimBlanks(s string) string {
	return strings.TrimSpace(s)
}

// RemoveBlanks replaces all blanks (including space, CR, LF, tabulation, ideographic space, etc.) to a single space " ", please see xstring.IsBlank and unicode.IsSpace
// for more details.
func RemoveBlanks(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s) / 2)
	space := false
	for _, r := range s {
		switch {
		case IsBlank(r):
			space = true
		default:
			if space == true && sb.Len() > 0 {
				space = false
				sb.WriteRune(' ')
			}
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// =================
// word case related
// =================

// CaseSplitter is the special word splitter used in SplitToWords, it means also to split the string by different cases, such as "helloWorld" to ["hello", "world"].
const CaseSplitter = ""

// defaultSplitters is the default splitter used in SplitToWords.
var defaultSplitters = []string{CaseSplitter, "_", "-", "."}

// SplitToWords splits given string to a word array using default or given word splitters. Here default splitters are blank characters (including space, CR, LF,
// tabulation, ideographic space, etc.), "_", "-" and ".", you can set the `seps` argument to replace the default word separators (such as new characters or CaseSplitter,
// but for blank characters).
func SplitToWords(s string, seps ...string) []string {
	// splitters
	if len(seps) == 0 {
		seps = defaultSplitters
	}
	splitters := make([]string, 0, len(seps))
	splitCase := false
	for _, sep := range seps {
		if sep == CaseSplitter {
			splitCase = true
		} else if len(sep) > 0 {
			splitters = append(splitters, sep)
		}
	}

	// replacer
	oldNews := make([]string, 0, len(splitters)*2)
	for _, rule := range splitters {
		oldNews = append(oldNews, rule, " ")
	}
	replacer := strings.NewReplacer(oldNews...)

	// split
	if splitCase {
		sb := strings.Builder{}
		sb.Grow(len(s))
		lastLower := false
		for i, r := range s {
			lower := !unicode.IsUpper(r)
			if i > 0 && !lower && lastLower {
				sb.WriteRune(' ') // split by case
			}
			sb.WriteRune(r)
			lastLower = lower
		}
		s = sb.String()
	}
	s = replacer.Replace(s) // split by splitters
	words := strings.Fields(s)
	return words
}

// PascalCase rewrites string in pascal case, by default, word splitters are blank characters, "_", "-" and ".", see SplitToWords fore more details.
func PascalCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = Capitalize(word)
	}
	return strings.Join(wordArray, "")
}

// CamelCase rewrites string in camel case, by default, word splitters are blank characters, "_", "-" and ".", see SplitToWords fore more details.
func CamelCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		if i == 0 {
			wordArray[i] = strings.ToLower(word)
		} else {
			wordArray[i] = Capitalize(word)
		}
	}
	return strings.Join(wordArray, "")
}

// SnakeCase rewrites string in snake case, by default, word splitters are blank characters, "_", "-" and ".", see SplitToWords fore more details.
func SnakeCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = strings.ToLower(word)
	}
	return strings.Join(wordArray, "_")
}

// KebabCase rewrites string in kebab case, by default, word splitters are blank characters, "_", "-" and ".", see SplitToWords fore more details.
func KebabCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = strings.ToLower(word)
	}
	return strings.Join(wordArray, "-")
}

// =====================
// uuid and rand related
// =====================

// TimeID creates an id from given time. If the count is larger than 23, the remaining bits will be filled by rand numbers.
func TimeID(t time.Time, count int) string {
	layoutWithNano := "20060102150405.000000000"
	uuid := t.Format(layoutWithNano)
	uuid = uuid[:14] + uuid[15:]
	l := len(uuid) // 23

	if count <= l {
		return uuid[:count]
	}
	return uuid + RandNumberString(count-l)
}

func init() {
	// for RandString
	rand.Seed(time.Now().UnixNano())
}

// RandString generates a string by given rune slice in random order.
func RandString(count int, runes []rune) string {
	b := make([]rune, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

var (
	capitalLetterRunes         = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowercaseLetterRunes       = []rune("abcdefghijklmnopqrstuvwxyz")
	allcaseLetterRunes         = append(capitalLetterRunes, lowercaseLetterRunes...)
	numberRunes                = []rune("0123456789")
	capitalLetterNumberRunes   = append(capitalLetterRunes, numberRunes...)
	lowercaseLetterNumberRunes = append(lowercaseLetterRunes, numberRunes...)
)

// RandCapitalLetterString generates a random string combined by capital letters, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ.
func RandCapitalLetterString(count int) string {
	return RandString(count, capitalLetterRunes)
}

// RandLowercaseLetterString generates a random string combined by lowercase letters, that is abcdefghijklmnopqrstuvwxyz.
func RandLowercaseLetterString(count int) string {
	return RandString(count, lowercaseLetterRunes)
}

// RandLetterString generates a random string combined by allcase letters, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ + abcdefghijklmnopqrstuvwxyz.
func RandLetterString(count int) string {
	return RandString(count, allcaseLetterRunes)
}

// RandNumberString generates a random string combined by numbers, that is 0123456789.
func RandNumberString(count int) string {
	return RandString(count, numberRunes)
}

// RandCapitalLetterNumberString generates a random string combined by capital letters and numbers, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ + 0123456789.
func RandCapitalLetterNumberString(count int) string {
	return RandString(count, capitalLetterNumberRunes)
}

// RandLowercaseLetterNumberString generates a random string combined by lowercase letters and numbers, that is abcdefghijklmnopqrstuvwxyz + 0123456789.
func RandLowercaseLetterNumberString(count int) string {
	return RandString(count, lowercaseLetterNumberRunes)
}

// ============
// fast related
// ============

// FastStob fast casts string to []byte. Note that this is an unsafe function.
func FastStob(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

// FastBtos fast casts []byte to string. Note that this is an unsafe function.
func FastBtos(bs []byte) string {
	if bs == nil || len(bs) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&bs))
}

// ===========================
// bom and replacement related
// ===========================

const (
	// utf8BomString is UTF8 BOM character in string, that is U+FEFF, or 0xEF 0xBB 0xBF.
	utf8BomString = "\xef\xbb\xbf"

	// utf8ReplacementString is UTF8 replacement character in string, that is U+FFFD, or 0xEF 0xBF 0xBD.
	utf8ReplacementString = "\xef\xbf\xbd"
)

var (
	// utf8BomString is UTF8 BOM character in string, that is U+FEFF, or 0xEF 0xBB 0xBF.
	utf8BomBytes = []byte{0xEF, 0xBB, 0xBF}

	// utf8ReplacementString is UTF8 replacement character in string, that is U+FFFD, or 0xEF 0xBF 0xBD.
	utf8ReplacementBytes = []byte{0xEF, 0xBF, 0xBD}
)

// TrimUTF8Bom trims BOM (byte order mark, U+FEFF, that is 0xEF 0xBB 0xBF in UTF-8) from a string. Please visit
// https://en.wikipedia.org/wiki/Byte_order_mark#Byte_order_marks_by_encoding and https://www.compart.com/en/unicode/U+FEFF for details.
func TrimUTF8Bom(s string) string {
	return strings.TrimPrefix(s, utf8BomString)
}

// TrimUTF8BomBytes trims BOM (byte order mark, U+FEFF, that is 0xEF 0xBB 0xBF in UTF-8) from a bytes. Please visit
// https://en.wikipedia.org/wiki/Byte_order_mark#Byte_order_marks_by_encoding and https://www.compart.com/en/unicode/U+FEFF for details.
func TrimUTF8BomBytes(bs []byte) []byte {
	return bytes.TrimPrefix(bs, utf8BomBytes)
}

// TrimUTF8Replacement trims replacement character (�, U+FFFD, that is 0xEF 0xBF 0xBD in UTF-8) from a string. Please visit
// https://en.wikipedia.org/wiki/Specials_(Unicode_block)#Replacement_character and https://www.compart.com/en/unicode/U+FFFD for details.
func TrimUTF8Replacement(s string) string {
	return strings.TrimPrefix(s, utf8ReplacementString)
}

// TrimUTF8ReplacementBytes trims replacement character (�, U+FFFD, that is 0xEF 0xBF 0xBD in UTF-8) from a bytes. Please visit
// https://en.wikipedia.org/wiki/Specials_(Unicode_block)#Replacement_character and https://www.compart.com/en/unicode/U+FFFD for details.
func TrimUTF8ReplacementBytes(bs []byte) []byte {
	return bytes.TrimPrefix(bs, utf8ReplacementBytes)
}

// ===================
// pad and get related
// ===================

// PadLeft returns the string with length of totalLength, which is padded by given pad rune to left, if given `totalLength` doesn't exceed the length of
// given string, this function does nothing.
func PadLeft(s string, pad rune, totalLength int) string {
	strLength := 0
	for range s {
		strLength++
	}
	if strLength >= totalLength {
		return s
	}
	sb := strings.Builder{}
	sb.Grow(totalLength)
	for i := 0; i < totalLength-strLength; i++ {
		sb.WriteRune(pad)
	}
	sb.WriteString(s)
	return sb.String()
}

// PadRight returns the string with length of totalLength, which is padded by given pad rune to right, if given `totalLength` doesn't exceed the length of
// given string, this function does nothing.
func PadRight(s string, pad rune, totalLength int) string {
	currLength := 0
	for range s {
		currLength++
	}
	if currLength >= totalLength {
		return s
	}
	sb := strings.Builder{}
	sb.Grow(totalLength)
	sb.WriteString(s)
	for i := 0; i < totalLength-currLength; i++ {
		sb.WriteRune(pad)
	}
	return sb.String()
}

// GetLeft gets the left part of given string with length, if given `length` exceeds the length of given string, this function does nothing.
func GetLeft(s string, length int) string {
	if length <= 0 {
		return ""
	}
	sb := strings.Builder{}
	sb.Grow(length)
	idx := 0
	for _, v := range s {
		if idx < length {
			idx++
			sb.WriteRune(v)
		} else {
			break
		}
	}
	return sb.String()
}

// GetRight gets the right part of given string with length, if given `length` exceeds the length of given string, this function does nothing.
func GetRight(s string, length int) string {
	if length <= 0 {
		return ""
	}
	strLength := 0
	for range s {
		strLength++
	}
	sb := strings.Builder{}
	sb.Grow(length)
	idx := 0
	for _, v := range s {
		if idx >= strLength-length {
			sb.WriteRune(v)
		}
		idx++
	}
	return sb.String()
}

// GetOrPadLeft gets the left part of given string, or pad given string by given rune to left, if `length` exceeds the length of given string, this
// function does pad, otherwise does get.
func GetOrPadLeft(s string, length int, pad rune) string {
	strLength := 0
	for range s {
		strLength++
	}
	if length == strLength {
		return s
	}
	if length < strLength {
		return GetLeft(s, length)
	}
	return PadLeft(s, pad, length)
}

// GetOrPadRight gets the right part of given string, or pad given string by given rune to right, if `length` exceeds the length of given string, this
// function does pad, otherwise does get.
func GetOrPadRight(s string, length int, pad rune) string {
	strLength := 0
	for range s {
		strLength++
	}
	if length == strLength {
		return s
	}
	if length < strLength {
		return GetRight(s, length)
	}
	return PadRight(s, pad, length)
}

// ExtraSpaceOnLeftIfNotEmpty returns a string from give string with an extra space on left if given string is not empty.
func ExtraSpaceOnLeftIfNotEmpty(s string) string {
	if len(s) == 0 {
		return ""
	}
	return " " + s
}

// ExtraSpaceOnRightIfNotEmpty returns a string from give string with an extra space on right if given string is not empty.
func ExtraSpaceOnRightIfNotEmpty(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s + " "
}

// ExtraSpaceOnLeftIfNotBlank returns a string from give string with an extra space on left if given string is not blank.
func ExtraSpaceOnLeftIfNotBlank(s string) string {
	s2 := TrimBlanks(s)
	if len(s2) == 0 {
		return s
	}
	return " " + s2
}

// ExtraSpaceOnRightIfNotBlank returns a string from give string with an extra space on right if given string is not blank.
func ExtraSpaceOnRightIfNotBlank(s string) string {
	s2 := TrimBlanks(s)
	if len(s2) == 0 {
		return s
	}
	return s2 + " "
}
