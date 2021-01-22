package xstring

import (
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// Capitalize capitalizes the first letter of the whole string, this will ignore all trailing spaces.
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	first := string(unicode.ToUpper(r[0]))
	if len(s) == 1 {
		return first
	}
	return first + string(r[1:])
}

// Uncapitalize uncapitalizes the first letter of the whole string, this will ignore all trailing spaces.
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	first := string(unicode.ToLower(r[0]))
	if len(s) == 1 {
		return first
	}
	return first + string(r[1:])
}

// CapitalizeAll capitalizes all the first letter in words of the whole string, this will trim the trailing space.
func CapitalizeAll(s string) string {
	sp := strings.Split(s, " ")
	out := make([]string, 0, len(sp))
	for _, word := range sp {
		if len(word) != 0 {
			out = append(out, Capitalize(word))
		}
	}
	return strings.Join(out, " ")
}

// UncapitalizeAll uncapitalizes all the first letter in words of the whole string, this will trim the trailing space.
func UncapitalizeAll(s string) string {
	sp := strings.Split(s, " ")
	out := make([]string, 0, len(sp))
	for _, word := range sp {
		if len(word) != 0 {
			out = append(out, Uncapitalize(word))
		}
	}
	return strings.Join(out, " ")
}

// RemoveBlanks replaces all blanks (\s and a wide space) to a single space. About blank, see unicode.IsSpace.
func RemoveBlanks(s string) string {
	// [\t\n\v\f\r\x85\xA0 　]
	s = regexp.MustCompile(`[\s　]+`).ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// caseHelper split a string to a word array using default and given word separator. Default separators are "_", "-". ".", " ", "　".
func caseHelper(s string, seps ...string) []string {
	seps = append(seps, "_", "-", ".", "　")
	oldNews := make([]string, 0, len(seps)*2)
	for _, rule := range seps {
		if rule != "" {
			oldNews = append(oldNews, rule, " ")
		}
	}
	replacer := strings.NewReplacer(oldNews...)

	re := regexp.MustCompile(`([a-z])([A-Z])`)
	s = re.ReplaceAllString(s, `$1 $2`)                         // split lowercase and capital
	s = strings.Join(strings.Fields(strings.TrimSpace(s)), " ") // remove duplicate spaces
	s = replacer.Replace(s)                                     // split by rules

	words := strings.Fields(strings.ToLower(s))
	return words
}

// PascalCase rewrites string in pascal case using word separator. By default "_", "-". ".", " ", "　" are treated as word separator.
func PascalCase(s string, seps ...string) string {
	wordArray := caseHelper(s, seps...)
	for i, word := range wordArray {
		wordArray[i] = Capitalize(word)
	}
	return strings.Join(wordArray, "")
}

// CamelCase rewrites string in camel case using word separator. By default "_", "-". ".", " ", "　" are treated as word separator.
func CamelCase(s string, seps ...string) string {
	wordArray := caseHelper(s, seps...)
	for i, word := range wordArray {
		if i > 0 {
			wordArray[i] = Capitalize(word)
		}
	}
	return strings.Join(wordArray, "")
}

// SnakeCase rewrites string in snake case using word separator. By default "_", "-". ".", " ", "　" are treated as word separator.
func SnakeCase(s string, seps ...string) string {
	wordArray := caseHelper(s, seps...)
	return strings.Join(wordArray, "_")
}

// KebabCase rewrites string in kebab case using word separator. By default "_", "-". ".", " ", "　" are treated as word separator.
func KebabCase(s string, seps ...string) string {
	wordArray := caseHelper(s, seps...)
	return strings.Join(wordArray, "-")
}

// TimeUUID creates a uuid from given time. If the count is larger than 23, the remaining bits will be filled by rand numbers.
func TimeUUID(t time.Time, count int) string {
	layoutWithNanosecond := "20060102150405.000000000"
	uuid := t.Format(layoutWithNanosecond)
	uuid = uuid[:14] + uuid[15:]
	l := len(uuid) // 23

	if count <= l {
		return uuid[:count]
	}
	return uuid + RandNumberString(count-l)
}

// RandString generates a string by given rune slice in random order.
func RandString(count int, runes []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

var (
	capitalLetterRunes         = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowercaseLetterRunes       = []rune("abcdefghijklmnopqrstuvwxyz")
	allcaseLetterRunes         = append(capitalLetterRunes, lowercaseLetterRunes...)
	numberRunes                = []rune("0123456789")
	capitalLetterNumberRunes   = append(capitalLetterRunes, numberRunes...)
	lowercaseLetterNumberRunes = append(lowercaseLetterRunes, numberRunes...)
)

// RandCapitalLetterString generates a random string combined by capital letters, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ.
func RandCapitalLetterString(count int) string {
	return RandString(count, capitalLetterRunes)
}

// RandLowercaseLetterString generates a random string combined by lowercase letters, that is abcdefghijklmnopqrstuvwxyz.
func RandLowercaseLetterString(count int) string {
	return RandString(count, lowercaseLetterRunes)
}

// RandLetterString generates a random string combined by allcase letters, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ + abcdefghijklmnopqrstuvwxyz.
func RandLetterString(count int) string {
	return RandString(count, allcaseLetterRunes)
}

// RandNumberString generates a random string combined by numbers, that is 0123456789.
func RandNumberString(count int) string {
	return RandString(count, numberRunes)
}

// RandCapitalLetterNumberString generates a random string combined by capital letters and numbers, that is ABCDEFGHIJKLMNOPQRSTUVWXYZ + 0123456789.
func RandCapitalLetterNumberString(count int) string {
	return RandString(count, capitalLetterNumberRunes)
}

// RandLowercaseLetterNumberString generates a random string combined by lowercase letters and numbers, that is abcdefghijklmnopqrstuvwxyz + 0123456789.
func RandLowercaseLetterNumberString(count int) string {
	return RandString(count, lowercaseLetterNumberRunes)
}

// DefaultMaskToken masks and returns a token string. Here the masked result only shows the first and last two characters at most,
// and the length of characters shown is restrict less than masked.
func DefaultMaskToken(s string) string {
	switch {
	case len(s) == 0:
		return ""
	}

	r := []rune(s)
	switch l := len(r); l {
	case 1:
		return "*" // * -> 0:1
	case 2:
		return "**" // ** -> 0:2
	case 3:
		return "**" + string(r[2]) // **3 -> 1:2
	case 4:
		return "***" + string(r[3]) // ***4 -> 1:3
	case 5:
		return string(r[0]) + "***" + string(r[4]) // 1***5 -> 2:3
	case 6:
		return string(r[0]) + "****" + string(r[5]) // 1****6 -> 2:4
	case 7:
		return string(r[0]) + "****" + string(r[5:7]) // 1****67 -> 3:4
	case 8:
		return string(r[0]) + "*****" + string(r[6:8]) // 1*****78 -> 3:5
	default:
		return string(r[0:2]) + strings.Repeat("*", l-4) + string(r[l-2:l]) // 12*****89 -> 4:5
	}
}

// MaskToken masks and returns a token string with given mask run and indices array, which support minus index.
func MaskToken(s string, mask rune, indices ...int) string {
	switch {
	case len(s) == 0:
		return ""
	case len(indices) == 0:
		return s
	}

	r := []rune(s)
	l := len(r)
	idxs := make([]int, 0, len(indices))
	for _, index := range indices {
		if 0 <= index && index < l {
			idxs = append(idxs, index)
		} else if -l <= index && index < 0 {
			idxs = append(idxs, l+index)
		}
	}
	sort.Ints(idxs)

	idxKvs := make(map[int]bool)
	for i, index := range idxs {
		if i == 0 || idxs[i-1] != index {
			idxKvs[index] = true
		}
	}

	sb := strings.Builder{}
	for i := range r {
		if _, ok := idxKvs[i]; !ok {
			sb.WriteRune(r[i])
		} else {
			sb.WriteRune(mask)
		}
	}
	return sb.String()
}

// FastStob fast casts string to []byte in an unsafe ways.
func FastStob(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
}

// FastBtos fast casts []byte to string in an unsafe ways.
func FastBtos(bs []byte) string {
	if bs == nil || len(bs) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&bs))
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
			sb.WriteString(key)
			sb.WriteString("=")
			sb.WriteString(val)
		}
	}

	return sb.String()
}

// PadLeft returns the string with length of totalLength, which is padded by char in left.
func PadLeft(s string, char rune, totalLength int) string {
	l := len([]rune(s))
	sp := strings.Builder{}
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	sp.WriteString(s)
	return sp.String()
}

// PadRight returns the string with length of totalLength, which is padded by char in right.
func PadRight(s string, char rune, totalLength int) string {
	l := len([]rune(s))
	sp := strings.Builder{}
	sp.WriteString(s)
	for i := 0; i < totalLength-l; i++ {
		sp.WriteRune(char)
	}
	return sp.String()
}

// GetLeft gets the left part of the string with length.
func GetLeft(s string, length int) string {
	if length <= 0 {
		return ""
	}

	runes := []rune(s)
	l := len(runes)
	if l <= length {
		return s
	}
	return string(runes[:length])
}

// GetRight gets the right part of the string with length.
func GetRight(s string, length int) string {
	if length <= 0 {
		return ""
	}

	runes := []rune(s)
	l := len(runes)
	if l <= length {
		return s
	}
	return string(runes[l-length:])
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
