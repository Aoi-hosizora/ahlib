package xruntime

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// TraceFrame represents a line of the runtime trace stack.
type TraceFrame struct {
	// Index represents the index of frame in stack.
	Index int

	// PC represents the frame's program count.
	PC uintptr

	// Filename represents the file full name.
	Filename string

	// FuncFullName represents the function fill name.
	FuncFullName string

	// FuncName represents the function name.
	FuncName string

	// LineIndex represents the line index in the file.
	LineIndex int

	// LineText represents the line text in the file.
	LineText string
}

// String returns the formatted TraceFrame.
//
// Format like:
// 	.../xruntime/xruntime_test.go:10 xruntime.TestTraceStack
// 		stack := RuntimeTraceStack(0)
func (t *TraceFrame) String() string {
	return fmt.Sprintf("%s:%d %s\n\t%s", t.Filename, t.LineIndex, t.FuncName, t.LineText)
}

// TraceStack represents the runtime trace stack, that is a slice of TraceFrame.
type TraceStack []*TraceFrame

// String returns the formatted TraceStack.
//
// Format like:
// 	.../xruntime/xruntime_test.go:10 xruntime.TestTraceStack
// 		stack := RuntimeTraceStack(0)
// 	.../src/testing/testing.go:1127 testing.tRunner
// 		fn(t)
func (t TraceStack) String() string {
	l := len(t)
	sb := strings.Builder{}
	for i, frame := range t {
		sb.WriteString(fmt.Sprintf("%s:%d %s", frame.Filename, frame.LineIndex, frame.FuncName))
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("\t%s", frame.LineText))
		if i != l-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// RuntimeTraceStack returns a slice of TraceFrame from runtime trace stacks using given skip (start from 1).
func RuntimeTraceStack(skip int) TraceStack {
	frames := make([]*TraceFrame, 0)
	for i := skip; ; i++ {
		pc, filename, lineIndex, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// func
		funcObj := runtime.FuncForPC(pc)
		funcFullName := funcObj.Name()
		_, funcName := filepath.Split(funcFullName)

		// line
		lineText := "?"
		if filename != "" {
			if data, err := ioutil.ReadFile(filename); err == nil {
				lines := bytes.Split(data, []byte{'\n'})
				if lineIndex > 0 && lineIndex <= len(lines) {
					lineText = string(bytes.TrimSpace(lines[lineIndex-1]))
				}
			}
		}

		// out
		frames = append(frames, &TraceFrame{
			Index:        i,
			PC:           pc,
			Filename:     filename,
			FuncFullName: funcFullName,
			FuncName:     funcName,
			LineIndex:    lineIndex,
			LineText:     lineText,
		})
	}

	return frames
}

// RuntimeTraceStackWithInfo get a slice of TraceFrame, with some information from the first trace stack line using given skip.
func RuntimeTraceStackWithInfo(skip int) (stack TraceStack, filename string, funcname string, lineIndex int, lineText string) {
	skip++
	stack = RuntimeTraceStack(skip)
	if len(stack) == 0 {
		return []*TraceFrame{}, "", "", 0, ""
	}
	top := stack[0]
	return stack, top.Filename, top.FuncName, top.LineIndex, top.LineText
}

var signalNames = [...]string{
	1:  "SIGHUP",
	2:  "SIGINT",
	3:  "SIGQUIT",
	4:  "SIGILL",
	5:  "SIGTRAP",
	6:  "SIGABRT",
	7:  "SIGBUS",
	8:  "SIGFPE",
	9:  "SIGKILL",
	10: "SIGUSR1",
	11: "SIGSEGV",
	12: "SIGUSR2",
	13: "SIGPIPE",
	14: "SIGALRM",
	15: "SIGTERM",
}

var signalReadableNames = [...]string{
	1:  "hangup",
	2:  "interrupt",
	3:  "quit",
	4:  "illegal instruction",
	5:  "trace/breakpoint trap",
	6:  "aborted",
	7:  "bus error",
	8:  "floating point exception",
	9:  "killed",
	10: "user defined signal 1",
	11: "segmentation fault",
	12: "user defined signal 2",
	13: "broken pipe",
	14: "alarm clock",
	15: "terminated",
}

// SignalName returns the SIGXXX string from given syscall.Signal. Note that syscall.Signal.String() and xruntime.SignalReadableName
// will return the human-readable string value.
func SignalName(sig syscall.Signal) string {
	if sig >= 1 && int(sig) < len(signalNames) {
		return signalNames[sig]
	}
	return "signal " + strconv.Itoa(int(sig))
}

// SignalReadableName returns the human-readable name from given syscall.Signal, this function has the same result with syscall.Signal.String().
func SignalReadableName(sig syscall.Signal) string {
	if sig >= 1 && int(sig) < len(signalReadableNames) {
		return signalReadableNames[sig]
	}
	return "signal " + strconv.Itoa(int(sig))
}
