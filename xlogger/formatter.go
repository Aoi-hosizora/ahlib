package xlogger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"sync"
	"time"
)

type CustomerFormatter struct {
	RuntimeCaller func(*runtime.Frame) (function string, file string)
	ForceColor    bool

	isTerminal       bool
	terminalInitOnce sync.Once
}

func (f *CustomerFormatter) init(entry *logrus.Entry) {
	if entry.Logger != nil {
		f.isTerminal = checkIfTerminal(entry.Logger.Out)
	}
}

func (f *CustomerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	f.terminalInitOnce.Do(func() {
		f.init(entry)
	})

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestampFormat := time.RFC3339
	caller := ""
	if entry.HasCaller() {
		funcVal := fmt.Sprintf("%s()", entry.Caller.Function)
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if f.RuntimeCaller != nil {
			funcVal, fileVal = f.RuntimeCaller(entry.Caller)
		}
		if fileVal == "" {
			caller = funcVal
		} else if funcVal == "" {
			caller = fileVal
		} else {
			caller = fileVal + " " + funcVal
		}
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]
	message := strings.TrimSuffix(entry.Message, "\n")
	now := entry.Time.Format(timestampFormat)

	if f.ForceColor || (f.isTerminal && runtime.GOOS != "windows") {
		var levelColor int
		switch entry.Level {
		case logrus.DebugLevel, logrus.TraceLevel:
			levelColor = 37 // gray
		case logrus.WarnLevel:
			levelColor = 33 // yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = 31 // red
		default:
			levelColor = 36 // blue
		}
		_, _ = fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s]%s %-44s ", levelColor, levelText, now, caller, message)
	} else {
		_, _ = fmt.Fprintf(b, "%s[%s]%s %-44s ", levelText, now, caller, message)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
