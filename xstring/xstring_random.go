package xstring

import (
	"math/rand"
	"time"
)

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
