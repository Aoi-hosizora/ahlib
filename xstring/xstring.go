package xstring

import (
	"strings"
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
	replace := func(src string) string {
		return strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(src,
					"\t", " "),
				"\n", " "),
			"  ", " ")
	}

	length := len(str)
	newStr := replace(str)
	for length != len(newStr) {
		length = len(newStr)
		newStr = replace(newStr)
	}
	return strings.TrimSpace(newStr)
}
