package xstring

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(ChatAt(str, 0)) + SubStringFrom(str, 1)
}

func Uncapitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(ChatAt(str, 0)) + SubStringFrom(str, 1)
}

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
	repeat := func(count int, char string) string {
		out := ""
		for idx := 0; idx < count; idx++ {
			out += char
		}
		return out
	}

	curr := 0
	out := ""
	for _, c := range jsonString {
		switch c {
		case '{', '[':
			curr++
			out += string(c) + "\n" + repeat(curr*intent, char)
		case '}', ']':
			curr--
			out += "\n" + repeat(curr*intent, char) + string(c)
		case ',':
			out += ",\n" + repeat(curr*intent, char)
		case ':':
			out += ": "
		case ' ', '\n', '\t':
			// pass
		default:
			out += string(c)
		}
	}
	return out
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
		return "*"
	case 2:
		return "*" + string(r[1:])
	case 3:
		return "**" + string(r[2:3])
	case 4:
		return string(r[0:1]) + "**" + string(r[3:4])
	case 5:
		return string(r[0:1]) + "***" + string(r[4:5])
	default:
		return string(r[0:2]) + strings.Repeat("*", l-4) + string(r[l-2:]) // <<< Default
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
