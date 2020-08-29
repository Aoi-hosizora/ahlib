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

func TestNil(t *testing.T) {
	Nil(t, nil)
	a := (*int)(nil)
	Nil(t, a)
	b := (map[string]string)(nil)
	Nil(t, b)
}

func TestNotNil(t *testing.T) {
	NotNil(t, 0)
	NotNil(t, "")
	NotNil(t, false)
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}

func TestEqualSlice(t *testing.T) {
	a := make([]interface{}, 0)
	b := make([]interface{}, 0)
	EqualSlice(t, a, b)

	c := []interface{}{1, 2, 3, 3}
	_ = []interface{}{1, 2, 3}
	e := []interface{}{3, 2, 3, 1}
	// EqualSlice(t, c, d)
	EqualSlice(t, c, e)
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
