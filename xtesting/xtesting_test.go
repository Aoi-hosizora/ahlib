package xtesting

import (
	"regexp"
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 0, 0)
	Equal(t, interface{}(0), 0)
	Equal(t, nil, (*int)(nil))
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 0, int32(0))
	NotEqual(t, interface{}(0), nil)
	NotEqual(t, nil, 0)
}

func TestMatchRegex(t *testing.T) {
	re := regexp.MustCompile(`^[abc]*[0-9A-Z]$`)
	MatchRegex(t, "aA", re)
	MatchRegex(t, "a0", re)
	MatchRegex(t, "X", re)
	MatchRegex(t, "bbX", re)
}

func TestNotMatchRegex(t *testing.T) {
	re := regexp.MustCompile(`^[abc]*[0-9A-Z]$`)
	NotMatchRegex(t, "", re)
	NotMatchRegex(t, "00", re)
	NotMatchRegex(t, "Aa", re)
	NotMatchRegex(t, "ca", re)
	NotMatchRegex(t, "bbZZ", re)
}
