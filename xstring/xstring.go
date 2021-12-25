package xstring

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

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

// CapitalizeAll capitalizes all the first letter in words of the whole string, words are split by blank character, see xstring.IsBlank.
func CapitalizeAll(s string) string {
	newWord := true
	sp := strings.Builder{}
	for _, v := range s {
		if newWord {
			newWord = false
			sp.WriteRune(unicode.ToUpper(v))
		} else {
			sp.WriteRune(v)
		}
		newWord = IsBlank(v)
	}
	return sp.String()
}

// UncapitalizeAll uncapitalizes all the first letter in words of the whole string, words are split by blank character, see xstring.IsBlank.
func UncapitalizeAll(s string) string {
	newWord := true
	sp := strings.Builder{}
	for _, v := range s {
		if newWord {
			newWord = false
			sp.WriteRune(unicode.ToLower(v))
		} else {
			sp.WriteRune(v)
		}
		newWord = IsBlank(v)
	}
	return sp.String()
}

// blankRe represents the blank regexp, that is [ \t\n\v\f\r\x85\xA0\u3000] (including unicode.IsSpace and ideographic space \u3000).
// Note that this regexp does not equal to /[\s　]/.
var blankRe = regexp.MustCompile("[ \\t\\n\\v\\f\\r\\x85\\xA0\u3000]+")

// IsBlank checks if given rune is a space or a blank, that is [ \t\n\v\f\r\x85\xA0\u3000], also see unicode.IsSpace.
func IsBlank(r rune) bool {
	return unicode.IsSpace(r) || r == '\u3000' // "　"
}

// RemoveBlanks replaces all blanks to a single space " ", also see xstring.IsBlank and unicode.IsSpace.
func RemoveBlanks(s string) string {
	s = blankRe.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// CaseSplitter is the special word splitter used in SplitToWords, and it means also to split the string by different cases, such as "helloWorld" to ["hello", "world"].
const CaseSplitter = ""

var defaultSplitters = []string{CaseSplitter, "_", "-", "."}

// SplitToWords splits a single string to a word array using default and given word separator. Default separators are [ \t\n\v\f\r\x85\xA0\u3000] (blank) and
// [_-.], you can set the seps parameter to use different word separators (but blank characters are just treated as word separator).
func SplitToWords(s string, seps ...string) []string { // caseHelper
	// separators
	if len(seps) == 0 {
		seps = defaultSplitters
	}
	separators := make([]string, 0, len(seps))
	splitCase := false
	for _, sep := range seps {
		if sep == CaseSplitter {
			splitCase = true
		} else if len(sep) > 0 {
			separators = append(separators, sep)
		}
	}

	// replacer
	oldNews := make([]string, 0, len(separators)*2+2)
	oldNews = append(oldNews, "　", " ") // "\u3000" => " "
	for _, rule := range separators {
		oldNews = append(oldNews, rule, " ")
	}
	replacer := strings.NewReplacer(oldNews...)

	// split
	if splitCase {
		sb := strings.Builder{}
		lastLower := false
		for i, r := range s {
			currLower := !unicode.IsUpper(r)
			if i > 0 && !currLower && lastLower {
				sb.WriteRune(' ') // split by case
			}
			sb.WriteRune(r)
			lastLower = currLower
		}
		s = sb.String()
	}
	s = replacer.Replace(s) // split by rules
	// words := strings.Fields(strings.ToLower(s))
	words := strings.Fields(s)
	return words
}

// PascalCase rewrites string in pascal case using word separator. By default, [ \t\n\v\f\r\x85\xA0\u3000] and [_-.] are treated as word separator.
func PascalCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = Capitalize(word)
	}
	return strings.Join(wordArray, "")
}

// CamelCase rewrites string in camel case using word separator. By default, [ \t\n\v\f\r\x85\xA0\u3000] and [_-.] are treated as word separator.
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

// SnakeCase rewrites string in snake case using word separator. By default, [ \t\n\v\f\r\x85\xA0\u3000] and [_-.] are treated as word separator.
func SnakeCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = strings.ToLower(word)
	}
	return strings.Join(wordArray, "_")
}

// KebabCase rewrites string in kebab case using word separator. By default, [ \t\n\v\f\r\x85\xA0\u3000] and [_-.] are treated as word separator.
func KebabCase(s string, extraSeps ...string) string {
	wordArray := SplitToWords(s, append(defaultSplitters, extraSeps...)...)
	for i, word := range wordArray {
		wordArray[i] = strings.ToLower(word)
	}
	return strings.Join(wordArray, "-")
}

// TimeUUID creates a uuid from given time. If the count is larger than 23, the remaining bits will be filled by rand numbers.
func TimeUUID(t time.Time, count int) string {
	layoutWithNanosecond := "20060102150405.000000000"
	uuid := t.Format(layoutWithNanosecond)
	uuid = uuid[:14] + uuid[15:]
	l := len(uuid) // 23

	if count <= l {
		return uuid[:count]
	}
	return uuid + RandNumberString(count-l)
}

// RandString generates a string by given rune slice in random order.
func RandString(count int, runes []rune) string {
	rand.Seed(time.Now().UnixNano())
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

// MaskToken masks a token string and returns the result, using given mask rune and indices for mask characters, this function also supports minus index.
func MaskToken(s string, mask rune, indices ...int) string {
	return coreMaskToken(s, mask, true, indices...)
}

// MaskTokenR masks a token string and returns the result, using given mask rune and indices for non-mask characters, this function also supports minus index,
func MaskTokenR(s string, mask rune, indices ...int) string {
	return coreMaskToken(s, mask, false, indices...)
}

// coreMaskToken is the core implementation of MaskToken and MaskTokenR.
func coreMaskToken(s string, mask rune, usedToMask bool, indices ...int) string {
	switch {
	case len(s) == 0: // empty
		return ""
	case len(indices) == 0: // no change or full change
		if usedToMask {
			return s
		}
		return strings.Repeat(string(mask), len(s))
	}

	runes := []rune(s)
	length := len(runes)
	newIndices := make(map[int]bool) // idx-true map

	idxs := make([]int, 0, len(indices)) // temp sorted indices
	for _, index := range indices {
		if 0 <= index && index < length {
			idxs = append(idxs, index)
		} else if -length <= index && index < 0 {
			idxs = append(idxs, length+index)
		}
	}
	sort.Ints(idxs)
	for i, index := range idxs {
		if i == 0 || idxs[i-1] != index {
			newIndices[index] = true // <<<
		}
	}

	sb := strings.Builder{}
	if usedToMask {
		// use index to write mask
		for i, ch := range runes {
			if _, ok := newIndices[i]; !ok {
				sb.WriteRune(ch)
			} else {
				sb.WriteRune(mask) // has index, write mask
			}
		}
	} else {
		// use index to write character
		for i, ch := range runes {
			if _, ok := newIndices[i]; !ok {
				sb.WriteRune(mask)
			} else {
				sb.WriteRune(ch) // has index, write character
			}
		}
	}
	return sb.String()
}

// FastStob fast casts string to []byte in an unsafe ways.
func FastStob(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

// FastBtos fast casts []byte to string in an unsafe ways.
func FastBtos(bs []byte) string {
	if bs == nil || len(bs) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&bs))
}

const (
	// utf8BomString is UTF8 BOM character in string, U+FEFF, 0xEF 0xBB 0xBF.
	utf8BomString = "\xef\xbb\xbf"

	// utf8ReplacementString is UTF8 replacement character in string, U+FFFD, 0xEF 0xBF 0xBD.
	utf8ReplacementString = "\xef\xbf\xbd"
)

var (
	// utf8BomBytes is UTF8 BOM character in bytes, U+FEFF, 0xEF 0xBB 0xBF.
	utf8BomBytes = []byte{0xEF, 0xBB, 0xBF}

	// utf8ReplacementString is UTF8 replacement character in bytes, U+FFFD, 0xEF 0xBF 0xBD.
	utf8ReplacementBytes = []byte{0xEF, 0xBF, 0xBD}
)

// TrimUTF8Bom trims BOM (byte order mark, U+FEFF, that is 0xEF 0xBB 0xBF in UTF-8) from a string.
// See https://en.wikipedia.org/wiki/Byte_order_mark#Byte_order_marks_by_encoding and https://www.compart.com/en/unicode/U+FEFF for details.
func TrimUTF8Bom(s string) string {
	return strings.TrimPrefix(s, utf8BomString)
}

// TrimUTF8BomBytes trims BOM (byte order mark, U+FEFF, that is 0xEF 0xBB 0xBF in UTF-8) from a bytes.
// See https://en.wikipedia.org/wiki/Byte_order_mark#Byte_order_marks_by_encoding and https://www.compart.com/en/unicode/U+FEFF for details.
func TrimUTF8BomBytes(bs []byte) []byte {
	return bytes.TrimPrefix(bs, utf8BomBytes)
}

// TrimUTF8Replacement trims replacement character (�, U+FFFD, that is 0xEF 0xBF 0xBD in UTF-8) from a string.
// See https://en.wikipedia.org/wiki/Specials_(Unicode_block)#Replacement_character and https://www.compart.com/en/unicode/U+FFFD for details.
func TrimUTF8Replacement(s string) string {
	return strings.TrimPrefix(s, utf8ReplacementString)
}

// TrimUTF8ReplacementBytes trims replacement character (�, U+FFFD, that is 0xEF 0xBF 0xBD in UTF-8) from a bytes.
// See https://en.wikipedia.org/wiki/Specials_(Unicode_block)#Replacement_character and https://www.compart.com/en/unicode/U+FFFD for details.
func TrimUTF8ReplacementBytes(bs []byte) []byte {
	return bytes.TrimPrefix(bs, utf8ReplacementBytes)
}

// EncodeUrlValues encodes the values (see url.Values) into url encoded form ("bar=baz&foo=quux") sorted by key with escape.
// The escapeFunc can be url.QueryEscape, url.PathEscape or the functions you defined, use nil for no escape.
func EncodeUrlValues(values map[string][]string, escapeFunc func(string) string) string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	for _, k := range keys {
		key := k
		if escapeFunc != nil {
			key = escapeFunc(key)
		}
		for _, v := range values[k] {
			val := v
			if escapeFunc != nil {
				val = escapeFunc(val)
			}
			if sb.Len() > 0 {
				sb.WriteString("&")
			}
			sb.WriteString(key) // escaped
			sb.WriteString("=")
			sb.WriteString(val) // escaped
		}
	}

	return sb.String()
}

// PadLeft returns the string with length of totalLength, which is padded by char in left.
func PadLeft(s string, char rune, totalLength int) string {
	l := 0 // length
	for range s {
		l++
	}
	sp := strings.Builder{}
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	sp.WriteString(s)
	return sp.String()
}

// PadRight returns the string with length of totalLength, which is padded by char in right.
func PadRight(s string, char rune, totalLength int) string {
	l := 0 // length
	for range s {
		l++
	}
	sp := strings.Builder{}
	sp.WriteString(s)
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	return sp.String()
}

// GetLeft gets the left part of the string with length.
func GetLeft(s string, length int) string {
	sp := strings.Builder{}
	idx := 0
	for _, v := range s {
		if idx < length {
			idx++
			sp.WriteRune(v)
		} else {
			break
		}
	}
	return sp.String()
}

// GetRight gets the right part of the string with length.
func GetRight(s string, length int) string {
	l := 0 // length
	for range s {
		l++
	}
	sp := strings.Builder{}
	idx := 0
	for _, v := range s {
		if idx >= l-length {
			sp.WriteRune(v)
		}
		idx++
	}
	return sp.String()
}

const (
	panicIndexOutOfRange = "xstring: index out of range"
)

// SplitAndGet returns the string item from the split result slices, this also supports minus index.
func SplitAndGet(s string, sep string, index int) string {
	sp := strings.Split(s, sep)
	l := len(sp)

	if index >= 0 && index < l {
		return sp[index]
	}
	if newIndex := l + index; newIndex >= 0 && newIndex < l {
		return sp[newIndex]
	}

	panic(panicIndexOutOfRange)
}

// SliceToStringMap returns a string-interface{} map from given interface{} slice.
func SliceToStringMap(args ...interface{}) map[string]interface{} {
	l := len(args)
	out := make(map[string]interface{}, l/2)

	for i := 0; i < l; i += 2 {
		if i+1 >= l {
			break // ignore the final arg
		}
		key := ""
		keyItf, value := args[i], args[i+1]
		if keyItf == nil {
			i--
			continue
		}
		if k, ok := keyItf.(string); ok {
			key = k
		} else {
			key = fmt.Sprintf("%v", keyItf) // %v
		}
		out[key] = value
	}

	return out
}
