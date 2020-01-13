package xstring

import "strings"

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
