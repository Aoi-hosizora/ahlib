package _example

import (
	"io"
)

type Logger interface {
	Debug(m string)
	Info(m string)
	Warn(m string)
	Error(m string)
	Panic(m string)
	Fatal(m string)
}

type Level uint8

const (
	Debug = iota + 1
	Info
	Warn
	Error
	Panic
	Fatal
	Silence
)

type defaultLogger struct {
	writer io.Writer
	level  Level
}

var _ Logger = (*defaultLogger)(nil)

func (d *defaultLogger) Debug(m string) {
	if d.level <= Debug {
		_, _ = d.writer.Write([]byte("[DEBUG] " + m))
	}
}

func (d *defaultLogger) Info(m string) {
	if d.level <= Info {
		_, _ = d.writer.Write([]byte("[INFO ] " + m))
	}
}

func (d *defaultLogger) Warn(m string) {
	if d.level <= Warn {
		_, _ = d.writer.Write([]byte("[WARN ] " + m))
	}
}

func (d *defaultLogger) Error(m string) {
	if d.level <= Error {
		_, _ = d.writer.Write([]byte("[ERROR] " + m))
	}
}

func (d *defaultLogger) Panic(m string) {
	if d.level <= Panic {
		_, _ = d.writer.Write([]byte("[PANIC] " + m))
	}
	panic(m)
}

func (d *defaultLogger) Fatal(m string) {
	if d.level <= Fatal {
		_, _ = d.writer.Write([]byte("[FATAL] " + m))
	}
	Exit(1) // <<<
}

func DefaultLogger(writer io.Writer, level Level) *defaultLogger {
	return &defaultLogger{writer: writer, level: level}
}
