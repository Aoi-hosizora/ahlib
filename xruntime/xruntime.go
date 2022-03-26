package xruntime

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// ===================
// trace stack related
// ===================

// RawStack returns the raw debug trace stack of the calling goroutine from runtime.Stack, if all is true, it will also return other
// goroutines' trace stack. Also see debug.Stack and runtime.Stack for more information.
//
// Returned value is just like:
// 	goroutine 19 [running]:
// 	github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x47)
// 		.../xruntime/xruntime.go:30 +0x9f
// 	github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0)
// 		.../xruntime/xruntime_test.go:20 +0x30
// 	testing.tRunner(0xc0000851e0, 0x224440)
// 		.../src/testing/testing.go:1259 +0x102
// 	created by testing.(*T).Run
// 		.../src/testing/testing.go:1306 +0x35a
func RawStack(all bool) []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, all)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

// TraceFrame represents a frame of the runtime trace stack, also see RuntimeTraceStack.
type TraceFrame struct {
	// Index represents the index of the frame in stack, 0 identifying the caller of this xruntime package.
	Index int

	// FuncPC represents the function's program count.
	FuncPC uintptr

	// FuncFullName represents the function's fill name, including the package full name and the function name.
	FuncFullName string

	// FuncName represents the function's name, including the package short name and the function name.
	FuncName string

	// Filename represents the file's full name.
	Filename string

	// LineIndex represents the line number in the file, starts from 1.
	LineIndex int

	// LineText represents the line text in the file，"?" if the text cannot be got.
	LineText string
}

// String returns the formatted TraceFrame.
//
// Returned value is just like:
// 	File: .../xruntime/xruntime_test.go:100
// 	Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack.func1
// 		stack := RuntimeTraceStack(0)
func (t *TraceFrame) String() string {
	return fmt.Sprintf("File: %s:%d\nFunc: %s\n\t%s", t.Filename, t.LineIndex, t.FuncFullName, t.LineText)
}

// TraceStack represents the runtime trace stack, consists of some TraceFrame, also see RuntimeTraceStack.
type TraceStack []*TraceFrame

// String returns the formatted TraceStack.
//
// Returned value is just like:
// 	File: .../xruntime/xruntime_test.go:100
// 	Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack.func1
// 		stack := RuntimeTraceStack(0)
// 	File: .../xruntime/xruntime_test.go:102
// 	Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
// 		}()
// 	File: D:/Development/Go/src/testing/testing.go:1439
// 	Func: testing.tRunner
// 		fn(t)
// 	File: D:/Development/Go/src/runtime/asm_amd64.s:1571
// 	Func: runtime.goexit
// 		BYTE	$0x90	// NOP
func (t TraceStack) String() string {
	l := len(t)
	sb := strings.Builder{}
	for i, frame := range t {
		sb.WriteString(frame.String())
		if i != l-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// RuntimeTraceStack returns TraceStack (a slice of TraceFrame) from runtime.Caller using given skip (0 identifying the caller of RuntimeTraceStack).
func RuntimeTraceStack(skip uint) TraceStack {
	frames := make([]*TraceFrame, 0)
	for i := skip + 1; ; i++ {
		funcPC, filename, lineIndex, ok := runtime.Caller(int(i))
		if !ok {
			break
		}

		// func
		funcObj := runtime.FuncForPC(funcPC)
		funcFullName := funcObj.Name()
		_, funcName := filepath.Split(funcFullName)

		// file
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
		frame := &TraceFrame{Index: int(i), FuncPC: funcPC, FuncFullName: funcFullName, FuncName: funcName, Filename: filename, LineIndex: lineIndex, LineText: lineText}
		frames = append(frames, frame)
	}

	return frames
}

// RuntimeTraceStackWithInfo returns TraceStack (a slice of TraceFrame) from runtime.Caller using given skip, with some information of the first TraceFrame's.
func RuntimeTraceStackWithInfo(skip uint) (stack TraceStack, filename string, funcName string, lineIndex int, lineText string) {
	stack = RuntimeTraceStack(skip + 1)
	if len(stack) == 0 {
		return []*TraceFrame{}, "", "", 0, ""
	}
	top := stack[0]
	return stack, top.Filename, top.FuncName, top.LineIndex, top.LineText
}

// ============================
// mass constants and functions
// ============================

// Pprof profile names, see pprof.Lookup (runtime/pprof) and pprof.Handler (net/http/pprof) for more details.
const (
	PprofGoroutineProfile    = "goroutine"
	PprofThreadcreateProfile = "threadcreate"
	PprofHeapProfile         = "heap"
	PprofAllocsProfile       = "allocs"
	PprofBlockProfle         = "block"
	PprofMutexProfile        = "mutex"
)

// signalNames is used by SignalName.
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

// signalReadableNames is used by SignalReadableName.
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

// GetProxyEnv lookups and returns three proxy environments, including http_proxy, https_proxy and socks_proxy.
func GetProxyEnv() (httpProxy string, httpsProxy string, socksProxy string) {
	hp := strings.TrimSpace(os.Getenv("http_proxy"))
	hsp := strings.TrimSpace(os.Getenv("https_proxy"))
	ssp := strings.TrimSpace(os.Getenv("socks_proxy"))
	return hp, hsp, ssp
}
