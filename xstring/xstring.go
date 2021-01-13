package xstring

import (
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// https://github.com/gobeam/Stringy

// Capitalize capitalizes the first letter of the whole string.
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	first := string(unicode.ToUpper(r[0]))
	if len(s) == 1 {
		return first
	}
	return first + string(r[1:])
}

// Uncapitalize uncapitalizes the first letter of the whole string.
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	first := string(unicode.ToLower(r[0]))
	if len(s) == 1 {
		return first
	}
	return first + string(r[1:])
}

// CapitalizeAll capitalizes all the first letter in words of the whole string.
func CapitalizeAll(s string) string {
	sp := strings.Split(s, " ")
	out := make([]string, 0, len(sp))
	for _, word := range sp {
		if len(word) != 0 {
			out = append(out, Capitalize(word))
		}
	}
	return strings.Join(out, " ")
}

// UncapitalizeAll uncapitalizes all the first letter in words of the whole string.
func UncapitalizeAll(s string) string {
	sp := strings.Split(s, " ")
	out := make([]string, 0, len(sp))
	for _, word := range sp {
		if len(word) != 0 {
			out = append(out, Uncapitalize(word))
		}
	}
	return strings.Join(out, " ")
}

// RemoveBlanks replaces all blanks (\s and a wide space) to a single space. About blank, see unicode.IsSpace.
func RemoveBlanks(s string) string {
	// [\t\n\v\f\r\x85\xA0 　]
	s = regexp.MustCompile(`[\s　]+`).ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// func ToSnakeCase(s string) string {
// 	// TODO
// 	out := ""
// 	newStr := Uncapitalize(s)
// 	for _, ch := range []rune(newStr) {
// 		if unicode.IsUpper(ch) {
// 			out += "_" + strings.ToLower(string(ch))
// 		} else if ch == ' ' {
// 			out += "_"
// 		} else {
// 			out += string(ch)
// 		}
// 	}
// 	return out
// }

// TimeUUID creates a uuid from given time. If the count is larger than 21, the remaining bits will be filled by rand numbers.
func TimeUUID(t time.Time, count int) string {
	layoutWithNanosecond := "20060102150405.0000000"
	uuid := t.Format(layoutWithNanosecond)
	uuid = uuid[:14] + uuid[15:]

	if count <= len(uuid) {
		return uuid[:count]
	} else {
		return uuid + RandNumberString(count-len(uuid))
	}
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

// DefaultMaskToken masks a token string, the masked result only shows the first two and last two characters.
func DefaultMaskToken(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	l := len(r)
	sb := strings.Builder{}
	switch l {
	case 1:
		return "*" // *
	case 2:
		return "**" // **
	case 3:
		return "**" + string(r[2]) // **2
	case 4:
		sb.WriteRune(r[0])
		sb.WriteString("**")
		sb.WriteRune(r[3]) // 0**3
	case 5:
		sb.WriteRune(r[0])
		sb.WriteString("***")
		sb.WriteRune(r[4]) // 0***4
	default:
		sb.WriteRune(r[0])
		sb.WriteRune(r[1])
		for i :=0; i < l - 4; i++ {
			sb.WriteRune('*')
		}
		sb.WriteRune(r[l-2])
		sb.WriteRune(r[l-1]) // 01***56
	}

	return sb.String()
}

// MaskToken masks a token string with given mask and indices, which support minus index.
func MaskToken(s string, mask rune, indices ...int) string {
	switch {
	case len(s) == 0:
		return ""
	case len(indices) == 0:
		return s
	}

	r := []rune(s)
	l := len(r)
	idxs := make([]int, 0, len(indices))
	for _, index := range indices {
		if 0 <= index && index < l {
			idxs = append(idxs, index)
		} else if -l <= index && index < 0 {
			idxs = append(idxs, l+index)
		}
	}
	sort.Ints(idxs)

	idxKvs := make(map[int]bool)
	for i, index := range idxs {
		if i == 0 || idxs[i-1] != index {
			idxKvs[index] = true
		}
	}

	sb := strings.Builder{}
	for i := range r {
		if _, ok := idxKvs[i]; !ok {
			sb.WriteRune(r[i])
		} else {
			sb.WriteRune(mask)
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
		if escapeFunc != nil {
			k = escapeFunc(k)
		}
		for _, v := range values[k] {
			if escapeFunc != nil {
				v = escapeFunc(v)
			}

			if sb.Len() > 0 {
				sb.WriteString("&")
			}
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(v)
		}
	}

	return sb.String()
}

// IsMark determines whether the rune is a marker.
func IsMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}

// PadLeft returns the string with length of totalLength, which is padded by char in left.
func PadLeft(s string, char rune, totalLength int) string {
	l := len([]rune(s))
	sp := strings.Builder{}
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	sp.WriteString(s)
	return sp.String()
}

// PadRight returns the string with length of totalLength, which is padded by char in right.
func PadRight(s string, char rune, totalLength int) string {
	l := len([]rune(s))
	sp := strings.Builder{}
	sp.WriteString(s)
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	return sp.String()
}

// GetLeft gets the left part of the string with length.
func GetLeft(s string, length int) string {
	runes := []rune(s)
	l := len(runes)
	if l <= length {
		return s
	}
	return string(runes[:length])
}

// GetRight gets the right part of the string with length.
func GetRight(s string, length int) string {
	runes := []rune(s)
	l := len(runes)
	if l <= length {
		return s
	}
	return string(runes[l-length:])
}

// SplitAndGet returns the string item from the split result slices, this also supports minus index.
func SplitAndGet(s string, sep string, index int) string {
	sp := strings.Split(s, sep)
	if len(sp) == 0 {
		return ""
	}

	if index >= 0 {
		return sp[index]
	} else {
		return sp[len(sp)+index]
	}
}
