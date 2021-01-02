package _example

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestExit(t *testing.T) {
	for _, tc := range []struct {
		name         string
		give         *StubbedExit
		wantExited   bool
		wantExitCode int
	}{
		{"no_exit", WithStub(func() {}), false, 0},
		{"exited", WithStub(func() { Exit(1) }), true, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			xtesting.Equal(t, tc.give.Exited, tc.wantExited)
			xtesting.Equal(t, tc.give.ExitCode, tc.wantExitCode)
		})
	}
}
