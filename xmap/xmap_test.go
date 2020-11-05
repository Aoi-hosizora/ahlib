package xmap

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	xtesting.Equal(t, SliceToInterfaceMap(nil), map[interface{}]interface{}{})
	xtesting.Equal(t, SliceToInterfaceMap([]interface{}{false, "b", "c"}), map[interface{}]interface{}{false: "b"})
	xtesting.Equal(t, SliceToInterfaceMap([]interface{}{"a", "b", "c", "d"}), map[interface{}]interface{}{"a": "b", "c": "d"})
	xtesting.Equal(t, SliceToInterfaceMap([]interface{}{1, 2}), map[interface{}]interface{}{1: 2})
	xtesting.Equal(t, SliceToInterfaceMap([]interface{}{nil, "  ", nil, " "}), map[interface{}]interface{}{nil: " "})

	xtesting.Equal(t, SliceToStringMap(nil), map[string]interface{}{})
	xtesting.Equal(t, SliceToStringMap([]interface{}{"a", "b", "c"}), map[string]interface{}{"a": "b"})
	xtesting.Equal(t, SliceToStringMap([]interface{}{"a", "b", "c", "d"}), map[string]interface{}{"a": "b", "c": "d"})
	xtesting.Equal(t, SliceToStringMap([]interface{}{1, 2}), map[string]interface{}{"1": 2})
	xtesting.Equal(t, SliceToStringMap([]interface{}{nil, "  ", " ", " "}), map[string]interface{}{" ": " "})
}
