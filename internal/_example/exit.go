package _example

import (
	"os"
)

var _exit = func(code int) { os.Exit(code) }

func Exit(code int) {
	_exit(code)
}

type StubbedExit struct {
	Exited   bool
	ExitCode int
	prev     func(int)
}

func (s *StubbedExit) UnStub() {
	_exit = s.prev
}

func (s *StubbedExit) exit(code int) {
	s.Exited = true
	s.ExitCode = code
}

func WithStub(f func()) *StubbedExit {
	s := &StubbedExit{prev: _exit}
	defer s.UnStub()
	_exit = s.exit
	if f != nil {
		f()
	}
	return s
}
