package xmap

import (
	"fmt"
)

// SliceToInterfaceMap returns a map[interface{}]interface{} from an interface slice.
func SliceToInterfaceMap(args []interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{})
	l := len(args)
	for i := 0; i < l; i += 2 {
		keyIdx := i
		valueIdx := i + 1
		if i+1 >= l {
			break
		}
		key := args[keyIdx]
		value := args[valueIdx]
		out[key] = value
	}

	return out
}

// SliceToInterfaceMap returns a map[string]interface{} from an interface slice.
func SliceToStringMap(args []interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	l := len(args)
	for i := 0; i < l; i += 2 {
		keyIdx := i
		valueIdx := i + 1
		if i+1 >= l {
			break
		}

		keyItf := args[keyIdx]
		value := args[valueIdx]
		key := ""
		if keyItf == nil {
			continue
		}
		if k, ok := keyItf.(string); ok {
			key = k
		} else {
			key = fmt.Sprintf("%v", keyItf)
		}
		out[key] = value
	}

	return out
}
