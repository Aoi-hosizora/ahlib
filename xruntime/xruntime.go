package xruntime

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
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
func (t *TraceFrame) String() string {
	return fmt.Sprintf("%s:%d (0x%x)\n\t%s: %s", t.Filename, t.LineIndex, t.PC, t.FuncName, t.LineText)
}

// TrackStack represents the runtime trace stack, that is a slice of TraceFrame.
type TraceStack []*TraceFrame

// String returns the formatted TraceStack.
func (t *TraceStack) String() string {
	l := len(*t)
	sb := &strings.Builder{}
	for i, frame := range *t {
		sb.WriteString(fmt.Sprintf("%s:%d (0x%x)", frame.Filename, frame.LineIndex, frame.PC))
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("\t%s", frame.LineText))
		if i != l-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// RuntimeTraceStack returns a slice of TraceFrame from runtime trace stacks using given skip.
func RuntimeTraceStack(skip int) TraceStack {
	skip++
	frames := make([]*TraceFrame, 0)
	for i := skip; ; i++ {
		pc, filename, lineIndex, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// func
		funcFullName := runtime.FuncForPC(pc).Name()
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
		return []*TraceFrame{}, "", "", -1, ""
	}
	top := stack[0]
	return stack, top.Filename, top.FuncName, top.LineIndex, top.LineText
}
