package xstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	assert.Equal(t, Capitalize("abc"), "Abc")
	assert.Equal(t, Capitalize("Abc"), "Abc")
	assert.Equal(t, Capitalize(""), "")
}

func TestUncapitalize(t *testing.T) {
	assert.Equal(t, Uncapitalize("Abc"), "abc")
	assert.Equal(t, Uncapitalize("abc"), "abc")
	assert.Equal(t, Uncapitalize(""), "")
}

func TestMarshalJson(t *testing.T) {
	a := struct {
		F1 string `json:"f1"`
		F2 struct{ F3 int }
	}{
		F1: "a",
		F2: struct{ F3 int }{F3: 3},
	}
	assert.Equal(t, MarshalJson(a), "{\"f1\":\"a\",\"F2\":{\"F3\":3}}")
}

func TestPrettyJson(t *testing.T) {
	from := "{\"a\": \"b\", \"c\": {\"d\": \"e\", \"f\": 0}, \"g\": [{\"h\": 1}, {\"h\": 1}]}"
	to := "{\n" +
		"    \"a\": \"b\",\n" +
		"    \"c\": {\n" +
		"        \"d\": \"e\",\n" +
		"        \"f\": 0\n" +
		"    },\n" +
		"    \"g\": [\n" +
		"        {\n" +
		"            \"h\": 1\n" +
		"        },\n" +
		"        {\n" +
		"            \"h\": 1\n" +
		"        }\n" +
		"    ]\n" +
		"}"
	assert.Equal(t, PrettyJson(from, 4, " "), to)
}
