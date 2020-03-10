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

func IsUppercase(char rune) bool {
	return char >= []rune("A")[0] && char <= []rune("Z")[0]
}

func IsLowercase(char rune) bool {
	return char >= []rune("a")[0] && char <= []rune("z")[0]
}

func ToSnakeCase(str string) string {
	out := ""
	newStr := Uncapitalize(str)
	for _, ch := range []rune(newStr) {
		if IsUppercase(ch) {
			out += "_" + strings.ToLower(string(ch))
		} else if ch == []rune(" ")[0] {
			out += "_"
		} else {
			out += string(ch)
		}
	}
	return out
}
