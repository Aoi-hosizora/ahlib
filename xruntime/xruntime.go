package xruntime

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type Stack struct {
	Index     int
	Filename  string
	Function  string
	Pc        uintptr
	LineIndex int
	Line      string
}

func (s *Stack) String() string {
	return fmt.Sprintf("%s:%d (0x%x)\n\t%s: %s", s.Filename, s.LineIndex, s.Pc, s.Function, s.Line)
}

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

func GetStackWithInfo(skip int) (stacks []*Stack, filename string, funcname string, lineIndex int, line string) {
	stacks = GetStack(skip + 1)
	if len(stacks) == 0 {
		return []*Stack{}, "", "", -1, ""
	}
	top := stacks[0]
	return stacks, top.Filename, top.Function, top.LineIndex, top.Line
}

func PrintStacks(stacks []*Stack) {
	for _, s := range stacks {
		fmt.Println(color.Red.Sprint(s.String()))
	}
}
