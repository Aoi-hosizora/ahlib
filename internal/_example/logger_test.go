package _example

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		name       string
		giveLogger Level
		giveLevel  Level
		giveString string
		wantString string
	}{
		{"debug for silence level", Debug, Silence, "debug", ""},
		{"info for silence level", Info, Silence, "debug", ""},
		{"warn for silence level", Warn, Silence, "debug", ""},
		{"error for silence level", Error, Silence, "debug", ""},
		{"panic for silence level", Panic, Silence, "debug", ""},

		{"debug for debug level", Debug, Debug, "debug", "[DEBUG] debug"},
		{"debug for info level", Debug, Info, "info", ""},
		{"info for info level", Info, Info, "info", "[INFO ] info"},
		{"info for warn level", Info, Warn, "warn", ""},
		{"warn for warn level", Warn, Warn, "warn", "[WARN ] warn"},
		{"warn for error level", Warn, Error, "error", ""},
		{"error for error level", Error, Error, "error", "[ERROR] error"},
		{"error for panic level", Error, Panic, "panic", ""},
		{"panic for panic level", Panic, Panic, "panic", "[PANIC] panic"},
		{"panic for fatal level", Panic, Fatal, "fatal", ""},
		{"fatal for fatal level", Fatal, Fatal, "fatal", "[FATAL] fatal"},
	}
	for _, tc := range tests {
		out := &strings.Builder{}
		logger := DefaultLogger(out, tc.giveLevel)
		var f func(string)
		switch tc.giveLogger {
		case Debug:
			f = logger.Debug
		case Info:
			f = logger.Info
		case Warn:
			f = logger.Warn
		case Error:
			f = logger.Error
		case Panic:
			f = logger.Panic
		case Fatal:
			f = logger.Fatal
		default:
			f = func(string) {}
		}
		t.Run(tc.name, func(t *testing.T) {
			if tc.giveLogger != Fatal {
				if tc.giveLogger == Panic {
					xtesting.Panic(t, func() { f(tc.giveString) })
				} else {
					xtesting.NotPanic(t, func() { f(tc.giveString) })
				}
				xtesting.Equal(t, out.String(), tc.wantString)
			} else {
				r := WithStub(func() { f(tc.giveString) })
				xtesting.Equal(t, r.Exited, true)
				xtesting.Equal(t, r.ExitCode, 1)
				xtesting.Equal(t, out.String(), tc.wantString)
			}
		})
	}
}
