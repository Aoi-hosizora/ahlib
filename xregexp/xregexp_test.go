package xregexp

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestRegexp(t *testing.T) {
	xtesting.Equal(t, EmailRegex.MatchString("aaa@bbb.ccc"), true)
	xtesting.Equal(t, EmailRegex.MatchString("a@b.c"), true)
	xtesting.Equal(t, EmailRegex.MatchString("a"), false)
}
