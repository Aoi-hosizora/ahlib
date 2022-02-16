package xstring

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Bool returns given t if given value is true, otherwise returns given f.
func Bool(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

// MaskToken masks a token string and returns the result, using given mask rune and indices for mask characters, this function also supports minus index.
func MaskToken(s string, mask rune, indices ...int) string {
	return coreMaskToken(s, string(mask), true, indices...)
}

// MaskTokenR masks a token string and returns the result, using given mask rune and indices for non-mask characters, this function also supports minus index,
func MaskTokenR(s string, mask rune, indices ...int) string {
	return coreMaskToken(s, string(mask), false, indices...)
}

// StringMaskToken masks a token string and returns the result, using given mask string and indices for mask characters, this function also supports minus index.
func StringMaskToken(s string, mask string, indices ...int) string {
	return coreMaskToken(s, mask, true, indices...)
}

// StringMaskTokenR masks a token string and returns the result, using given mask string and indices for non-mask characters, this function also supports minus index,
func StringMaskTokenR(s string, mask string, indices ...int) string {
	return coreMaskToken(s, mask, false, indices...)
}

// coreMaskToken is the core implementation of MaskToken and MaskTokenR.
func coreMaskToken(s string, mask string, toMask bool, indices ...int) string {
	switch {
	case len(s) == 0: // empty
		return ""
	case len(indices) == 0: // no change or full change
		if toMask {
			return s
		}
		return strings.Repeat(mask, len(s))
	}

	length := 0
	for range s {
		length++
	}
	newIndices := make(map[int]struct{}) // idx map
	idxs := make([]int, 0, len(indices))
	for _, i := range indices {
		if 0 <= i && i < length {
			idxs = append(idxs, i)
		} else if -length <= i && i < 0 {
			idxs = append(idxs, length+i)
		}
	}
	sort.Ints(idxs)
	for _, index := range idxs {
		newIndices[index] = struct{}{}
	}

	sb := strings.Builder{}
	sb.Grow(len(s))
	// use index to write mask or character
	for i, ch := range s {
		_, contains := newIndices[i]
		if (toMask && contains) || (!toMask && !contains) {
			sb.WriteString(mask)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

// EncodeUrlValues encodes the values (see url.Values) into url encoded form ("bar=baz&foo=quux") sorted by key with escape.
// The escapeFunc can be url.QueryEscape, url.PathEscape or the functions you defined, use nil for no escape.
func EncodeUrlValues(values map[string][]string, escapeFunc func(string) string) string {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	sb.Grow(len(values) * 4)
	for _, k := range keys {
		key := k
		if escapeFunc != nil {
			key = escapeFunc(key)
		}
		for _, v := range values[k] {
			val := v
			if escapeFunc != nil {
				val = escapeFunc(val)
			}
			if sb.Len() > 0 {
				sb.WriteString("&")
			}
			sb.WriteString(key) // escaped
			sb.WriteString("=")
			sb.WriteString(val) // escaped
		}
	}

	return sb.String()
}

const (
	panicIndexOutOfRange = "xstring: index out of range"
)

// SplitAndGet returns the string item from the split result slices, this also supports minus index.
func SplitAndGet(s string, sep string, index int) string {
	sp := strings.Split(s, sep)
	l := len(sp)

	if index >= 0 && index < l {
		return sp[index]
	}
	if newIndex := l + index; newIndex >= 0 && newIndex < l {
		return sp[newIndex]
	}

	panic(panicIndexOutOfRange)
}

// StringSliceToMap returns a string-string map from given string slice, notes that extra argument will be skipped.
func StringSliceToMap(args []string) map[string]string {
	l := len(args)
	out := make(map[string]string, l/2)
	for i := 0; i < l; i += 2 {
		if i+1 >= l {
			break // ignore the extra arg
		}
		key, value := args[i], args[i+1]
		out[key] = value
	}
	return out
}

// SliceToStringMap returns a string-interface{} map from given interface{} slice, notes that nil arguments and extra argument will be skipped.
func SliceToStringMap(args []interface{}) map[string]interface{} {
	l := len(args)
	out := make(map[string]interface{}, l/2)

	for i := 0; i < l; i += 2 {
		if i+1 >= l {
			break // ignore the extra arg
		}
		key := ""
		keyItf, value := args[i], args[i+1]
		if keyItf == nil {
			i--
			continue
		}
		if s, ok := keyItf.(string); ok {
			key = s
		} else if bs, ok := keyItf.([]byte); ok {
			key = FastBtos(bs)
		} else {
			key = fmt.Sprintf("%v", keyItf) // %v
		}
		out[key] = value
	}

	return out
}

var (
	errUnsupportedVersionString = errors.New("xstring: unsupported version string")
)

// SemanticVersion parses given semantic version string to (major, minor, patch) version number, notes that this function supports
// "1", "1.1", "1.1.1", "v1", "v1.1", "v1.1.1" format.
func SemanticVersion(semver string) (uint64, uint64, uint64, error) {
	if len(semver) == 0 || (len(semver) == 1 && !IsArabicNumber(rune(semver[0]))) {
		return 0, 0, 0, errUnsupportedVersionString
	}
	if semver[0] == 'v' {
		semver = semver[1:]
	}
	first := strings.IndexByte(semver, '.')
	last := strings.LastIndexByte(semver, '.')
	if first == -1 {
		// 1 / v1
		p1, err1 := strconv.ParseUint(semver, 10, 64)
		if err1 != nil {
			return 0, 0, 0, errUnsupportedVersionString
		}
		return p1, 0, 0, nil
	}
	if first == last {
		// 1.1 / v1.1
		p1, err1 := strconv.ParseUint(semver[:first], 10, 64)
		p2, err2 := strconv.ParseUint(semver[first+1:], 10, 64)
		if err1 != nil || err2 != nil {
			return 0, 0, 0, errUnsupportedVersionString
		}
		return p1, p2, 0, nil
	}
	// 1.1.1 / v1.1.1
	p1, err1 := strconv.ParseUint(semver[:first], 10, 64)
	p2, err2 := strconv.ParseUint(semver[first+1:last], 10, 64)
	p3, err3 := strconv.ParseUint(semver[last+1:], 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, errUnsupportedVersionString
	}
	return p1, p2, p3, nil
}
