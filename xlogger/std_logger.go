package xlogger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type StdLogger struct {
	mu  sync.Mutex
	out io.Writer
	buf []byte
}

func NewStdLogger(out io.Writer) *StdLogger {
	return &StdLogger{out: out}
}

func (l *StdLogger) Writer() io.Writer {
	return l.out
}

func (l *StdLogger) Output(s string) {
	now := time.Now()
	t := fmt.Sprintf("[%s] ", now.Format(time.RFC3339))
	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf = l.buf[:0]
	l.buf = append(l.buf, []byte(t)...)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, _ = l.out.Write(l.buf)
}

func (l *StdLogger) Outputf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(s)
}

func (l *StdLogger) Outputln(format string) {
	s := fmt.Sprintln(format)
	l.Output(s)
}

var _stdLogger = NewStdLogger(os.Stderr)

// noinspection GoUnusedExportedFunction
func Writer() io.Writer {
	return _stdLogger.Writer()
}

func Output(s string) {
	_stdLogger.Output(s)
}

func Outputf(format string, v ...interface{}) {
	_stdLogger.Outputf(format, v...)
}

func Outputln(format string) {
	_stdLogger.Outputln(format)
}
