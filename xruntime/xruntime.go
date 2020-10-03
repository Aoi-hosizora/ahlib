package xruntime

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Stack struct {
	Index     int     `json:"index"`
	Filename  string  `json:"filename"`
	Function  string  `json:"function"`
	Pc        uintptr `json:"pc"`
	LineIndex int     `json:"line_index"`
	Line      string  `json:"line"`
}

func (s *Stack) String() string {
	return fmt.Sprintf("%s:%d (0x%x)\n\t%s: %s", s.Filename, s.LineIndex, s.Pc, s.Function, s.Line)
}

// GetStack returns a slice of Stack from runtime stacks using given skip.
func GetStack(skip int) []*Stack {
	skip++
	out := make([]*Stack, 0)
	for i := skip; ; i++ {
		pc, filename, lineNumber, ok := runtime.Caller(i)
		if !ok {
			break
		}
		function := runtime.FuncForPC(pc).Name()
		_, function = filepath.Split(function)

		lineContent := "?"
		if filename != "" {
			if data, err := ioutil.ReadFile(filename); err == nil {
				lines := bytes.Split(data, []byte{'\n'})
				if lineNumber > 0 && lineNumber <= len(lines) {
					lineContent = string(bytes.TrimSpace(lines[lineNumber-1]))
				}
			}
		}
		out = append(out, &Stack{
			Index:     i,
			Filename:  filename,
			Function:  function,
			Pc:        pc,
			LineIndex: lineNumber,
			Line:      lineContent,
		})
	}

	return out
}

// GetStackWithInfo returns some information from the first runtime stack using given skip.
func GetStackWithInfo(skip int) (stacks []*Stack, filename string, funcname string, lineIndex int, line string) {
	skip++
	stacks = GetStack(skip)
	if len(stacks) == 0 {
		return []*Stack{}, "", "", -1, ""
	}
	top := stacks[0]
	return stacks, top.Filename, top.Function, top.LineIndex, top.Line
}

// PrintStacks prints a slice of stacks.
func PrintStacks(stacks []*Stack) {
	for _, s := range stacks {
		fmt.Println(s.String())
	}
}

// PrintStacksRed prints a slice of stacks using xcolor.Red.
func PrintStacksRed(stacks []*Stack) {
	l := log.New(os.Stderr, "", 0)
	xcolor.ForceColor()
	for _, s := range stacks {
		line := s.String()
		for _, s := range strings.Split(line, "\n") {
			l.Println(xcolor.Red.Sprint(s))
		}
	}
}
