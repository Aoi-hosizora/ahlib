package xstring

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// Capitalize capitalizes string.
func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(ChatAt(str, 0)) + SubStringFrom(str, 1)
}

// Uncapitalize uncapitalizes string.
func Uncapitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(ChatAt(str, 0)) + SubStringFrom(str, 1)
}

// RemoveSpaces removes ` ` `\n` `\t` from string.
func RemoveSpaces(str string) string {
	r, _ := regexp.Compile(`[\s　]+`) // BS \n \t
	str = r.ReplaceAllString(str, " ")
	return strings.TrimSpace(str)
}

func ChatAt(str string, idx int) string {
	return string([]rune(str)[idx])
}

func SubString(str string, f int, t int) string {
	return string([]rune(str)[f:t])
}

func SubStringFrom(str string, f int) string {
	return string([]rune(str)[f:])
}

func SubStringTo(str string, t int) string {
	return string([]rune(str)[:t])
}

func ToRune(char string) rune {
	if char == "" {
		return 0
	}
	return []rune(char)[0]
}

func ToByte(char string) byte {
	if char == "" {
		return 0
	}
	return char[0]
}

func IsUppercase(char rune) bool {
	return char >= ToRune("A") && char <= ToRune("Z")
}

func IsLowercase(char rune) bool {
	return char >= ToRune("a") && char <= ToRune("z")
}

func ToSnakeCase(str string) string {
	out := ""
	newStr := Uncapitalize(str)
	for _, ch := range []rune(newStr) {
		if IsUppercase(ch) {
			out += "_" + strings.ToLower(string(ch))
		} else if ch == ToRune(" ") {
			out += "_"
		} else {
			out += string(ch)
		}
	}
	return out
}

func PrettifyJson(jsonString string, intent int, char string) string {
	curr := 0
	sb := strings.Builder{}
	jsonString = RemoveSpaces(jsonString)
	l := len(jsonString)
	for idx, c := range jsonString {
		switch c {
		case '{', '[':
			curr++
			sb.WriteRune(c)
			if idx+1 < l && jsonString[idx+1] != '}' && jsonString[idx+1] != ']' {
				sb.WriteString("\n")
				sb.WriteString(strings.Repeat(char, curr*intent))
			}
		case '}', ']':
			curr--
			if idx-1 > -1 && jsonString[idx-1] != '{' && jsonString[idx-1] != '[' {
				sb.WriteString("\n")
				sb.WriteString(strings.Repeat(char, curr*intent))
			}
			sb.WriteRune(c)
		case ',':
			sb.WriteString(",\n")
			sb.WriteString(strings.Repeat(char, curr*intent))
		case ':':
			sb.WriteString(": ")
		// case ' ', '\n', '\t':
		// 	// pass
		default:
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func CurrentTimeUuid(count int) string {
	return TimeUuid(time.Now(), count)
}

// count: [0, 21+]
func TimeUuid(t time.Time, count int) string {
	nanosecondLayout := "20060102150405.0000000"
	uuid := t.Format(nanosecondLayout)
	uuid = uuid[:14] + uuid[15:]

	if count <= len(uuid) {
		return uuid[:count]
	} else {
		return uuid + RandNumberString(count-len(uuid))
	}
}

func RandString(count int, runes []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

var (
	CapitalLetterRunes   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LowercaseLetterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	NumberRunes          = []rune("0123456789")

	LetterRunes                = append(CapitalLetterRunes, LowercaseLetterRunes...)
	LetterNumberRunes          = append(LetterRunes, NumberRunes...)
	CapitalLetterNumberRunes   = append(CapitalLetterRunes, NumberRunes...)
	LowercaseLetterNumberRunes = append(LowercaseLetterRunes, NumberRunes...)
)

// Capital + Lower: ABCDEFGHIJKLMNOPQRSTUVWXYZ + abcdefghijklmnopqrstuvwxyz
func RandLetterString(count int) string {
	return RandString(count, LetterRunes)
}

// Only number: 0123456789
func RandNumberString(count int) string {
	return RandString(count, NumberRunes)
}

// Letter + Number: abcdefghijklmnopqrstuvwxyz + 0123456789
func RandLetterNumberString(count int) string {
	return RandString(count, LowercaseLetterNumberRunes)
}

func MaskToken(token string) string {
	r := []rune(token)
	switch l := len(r); l {
	case 0:
		return ""
	case 1:
		return "*" // *
	case 2:
		return "*" + string(r[1:]) // *1
	case 3:
		return "**" + string(r[2:3]) // **2
	case 4:
		return string(r[0:1]) + "**" + string(r[3:4]) // 0**3
	case 5:
		return string(r[0:1]) + "***" + string(r[4:5]) // 0***4
	default:
		return string(r[0:2]) + strings.Repeat("*", l-4) + string(r[l-2:]) // 01***56
	}
}

// Unsafe case to []byte.
func StringToBytes(str string) []byte {
	if str == "" {
		return []byte{}
	}
	return *(*[]byte)(unsafe.Pointer(&str))
}

// Unsafe case to string.
func BytesToString(bs []byte) string {
	if bs == nil || len(bs) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&bs))
}

// MapSliceToMap returns map[string]string to map[string][]string.
func MapSliceToMap(m map[string][]string) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		if l := len(v); l > 0 {
			out[k] = v[l-1]
		}
	}
	return out
}

// MapToMapSlice returns map[string][]string to map[string]string.
func MapToMapSlice(m map[string]string) map[string][]string {
	out := make(map[string][]string)
	for k, v := range m {
		out[k] = []string{v}
	}
	return out
}

func QueryString(values map[string][]string) string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	for _, k := range keys {
		for _, v := range values[k] {
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

// StringInterface returns a string value from any interface.
func StringInterface(i interface{}) string {
	s, ok := i.(string)
	if ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// ErrorInterface returns an error value from any interface.
func ErrorInterface(i interface{}) error {
	s, ok := i.(error)
	if ok {
		return s
	}
	return fmt.Errorf("%v", i)
}

// DefaultFormatString returns the default format string of interface, equals to '%v'.
func DefaultFormatString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// GoSyntaxString returns the go-syntax representation of interface, equals to '%#v'.
func GoSyntaxString(i interface{}) string {
	return fmt.Sprintf("%#v", i)
}

// IsMark determines whether the rune is a marker
func IsMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}

// IsEmpty returns true if the string is empty.
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty.
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace).
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank.
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

// PadLeft returns the string with length, which is padded by paddingChar in left.
func PadLeft(str string, paddingChar rune, totalLength int) string {
	l := len([]rune(str))
	sp := strings.Builder{}
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(paddingChar)
	}
	sp.WriteString(str)
	return sp.String()
}

// PadRight returns the string with length, which is padded by paddingChar in right.
func PadRight(str string, paddingChar rune, totalLength int) string {
	l := len([]rune(str))
	sp := strings.Builder{}
	sp.WriteString(str)
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(paddingChar)
	}
	return sp.String()
}

// GetLeft gets the left part of the string with length.
func GetLeft(str string, length int) string {
	runes := []rune(str)
	l := len(runes)
	if l <= length {
		return str
	}
	return string(runes[:length])
}

// GetRight gets the right part of the string with length.
func GetRight(str string, length int) string {
	runes := []rune(str)
	l := len(runes)
	if l <= length {
		return str
	}
	return string(runes[l-length:])
}
