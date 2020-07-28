package xstring

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.Replace(str, string(str[0]), strings.ToUpper(string(str[0])), 1)
}

func Uncapitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.Replace(str, string(str[0]), strings.ToLower(string(str[0])), 1)
}

func ToRune(char string) rune {
	if char == "" {
		return 0
	}
	return []rune(char)[0]
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

func RemoveSpaces(str string) string {
	r, _ := regexp.Compile(`[\s　]+`) // BS \n \t
	str = r.ReplaceAllString(str, " ")
	return strings.TrimSpace(str)
}

// return string(byte[]), return "" if err
func MarshalJson(object interface{}) string {
	j, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	return string(j)
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

// Capital + Lowercase
func RandLetterString(count int) string {
	return RandString(count, LetterRunes)
}

// Only number
func RandNumberString(count int) string {
	return RandString(count, NumberRunes)
}

func MaskToken(token string) string {
	switch len(token) {
	case 0:
		return ""
	case 1:
		return "*"
	case 2:
		return "*" + token[1:]
	case 3:
		return "**" + token[2:3]
	case 4:
		return token[0:1] + "**" + token[3:4]
	case 5:
		return token[0:1] + "***" + token[4:5]
	default:
		return token[0:2] + strings.Repeat("*", len(token)-4) + token[len(token)-2:] // <<< Default
	}
}
