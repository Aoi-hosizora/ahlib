package xstring

import (
	"encoding/json"
)

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
