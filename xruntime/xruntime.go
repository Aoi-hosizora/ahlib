package xruntime

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// ================
// function related
// ================

// NameOfFunction returns given function's name by searching runtime.Func by PC.
func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// ===================
// trace stack related
// ===================

// RawStack returns the raw debug trace stack of the calling goroutine from runtime.Stack, if all is true, it will also return other
// goroutines' trace stack. Also see debug.Stack and runtime.Stack for more information.
//
// Returned value is just like:
// 	goroutine 19 [running]:
// 	github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x7?)
// 		.../xruntime/xruntime.go:46 +0x6a
// 	github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0?)
// 		.../xruntime/xruntime_test.go:65 +0x30
// 	testing.tRunner(0xc000085380, 0xf682b0)
// 		.../src/testing/testing.go:1439 +0x102
// 	created by testing.(*T).Run
// 		.../src/testing/testing.go:1486 +0x35f
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
// 	File: .../xruntime/xruntime_test.go:145
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
// 	File: .../xruntime/xruntime_test.go:145
// 	Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack.func1
// 		stack := RuntimeTraceStack(0)
// 	File: .../xruntime/xruntime_test.go:147
// 	Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
// 		}()
// 	File: .../src/testing/testing.go:1439
// 	Func: testing.tRunner
// 		fn(t)
// 	File: .../src/runtime/asm_amd64.s:1571
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

// ========================
// hack hide string related
// ========================

var (
	// _magicBytes represents a slice of bytes, and are used in HackHideString and HackGetHiddenString.
	_magicBytes = []byte{0, 'h', 'i', 0, 'd', 'e'}

	_magicLength = len(_magicBytes)
)

// HackHideString hides given hidden string after given data address (in heap space) and returns the new data address. Note that this is an unsafe function.
//
// Example:
// 	handler := func() {}
// 	funcSize := int(reflect.TypeOf(func() {}).Size()) // which always equals to 8 in x86-64 machine
// 	handlerFn := (*struct{fn uintptr})(unsafe.Pointer(&handler)).fn
// 	hackedFn := HackHideString(handlerFn, funcSize, "handlerName")
// 	hackedHandler := *(*func())(unsafe.Pointer(&struct{fn uintptr}{fn: hackedFn}))
// 	realHandlerName := HackGetHiddenString(hackedFn, funcSize) // got "handlerName"
// 	hackedHandler() // hackedHandler can be invoked normally
//go:nosplit
func HackHideString(given uintptr, givenLength int, hidden string) (dataAddr uintptr) {
	hiddenHeader := (*reflect.StringHeader)(unsafe.Pointer(&hidden))
	hiddenLen := int32(hiddenHeader.Len) // only store int32 length
	hiddenBs := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: hiddenHeader.Data, Len: int(hiddenLen), Cap: int(hiddenLen)}))
	bs := bytes.Buffer{}
	bs.Grow(givenLength + _magicLength + 4 + int(hiddenLen))

	// given
	for i := 0; i < givenLength; i++ {
		bs.WriteByte(*(*byte)(unsafe.Pointer(given + uintptr(i))))
	}
	bs.Write(_magicBytes) // \x00, 'h', 'i', \x00, 'd', 'e', # = 6

	// hidden
	bs.Write([]byte{byte(hiddenLen), byte(hiddenLen >> 8), byte(hiddenLen >> 16), byte(hiddenLen >> 24)}) // lo -> hi, # = 4
	bs.Write(hiddenBs)

	fakeSlice := bs.Bytes() // its Data part will be used to interpret as other types
	return (*reflect.SliceHeader)(unsafe.Pointer(&fakeSlice)).Data
}

// HackHideStringAfterString hides given hidden string after given string (in heap space) and returns the new string. Note that this is an unsafe function.
//
// Example:
// 	httpMethod := "GET"
// 	hackedMethod := xruntime.HackHideStringAfterString(httpMethod, "handlerName")
// 	realHandlerName := xruntime.HackGetHiddenStringAfterString(hackedMethod) // got "handlerName"
// 	fmt.Println(hackedMethod) // hackedMethod can be printed normally
//go:nosplit
func HackHideStringAfterString(given, hidden string) string {
	givenHeader := (*reflect.StringHeader)(unsafe.Pointer(&given))
	fakeData := HackHideString(givenHeader.Data, givenHeader.Len, hidden)
	fakeHeader := &reflect.StringHeader{Data: fakeData, Len: givenHeader.Len}
	return *(*string)(unsafe.Pointer(fakeHeader))
}

// HackGetHiddenString get the hidden string from given unsafe.Pointer, see HackHideString for example. Note that this is an unsafe function.
//go:nosplit
func HackGetHiddenString(given uintptr, givenLength int) (hidden string) {
	// check magic
	for i, ch := range _magicBytes {
		if *(*byte)(unsafe.Pointer(given + uintptr(givenLength+i))) != ch {
			return ""
		}
	}

	// get hidden data
	digits := *(*[4]byte)(unsafe.Pointer(given + uintptr(givenLength+_magicLength))) // lo -> hi
	hiddenData := given + uintptr(givenLength+_magicLength+4)
	hiddenLen := int32(digits[0]) + int32(digits[1])<<8 + int32(digits[2])<<16 + int32(digits[3])<<24

	// construct string
	fakeHeader := &reflect.StringHeader{Data: hiddenData, Len: int(hiddenLen)}
	return *(*string)(unsafe.Pointer(fakeHeader))
}

// HackGetHiddenStringAfterString get the hidden string from given string, see HackHideStringAfterString for example. Note that this is an unsafe function.
//go:nosplit
func HackGetHiddenStringAfterString(given string) (hidden string) {
	givenHeader := (*reflect.StringHeader)(unsafe.Pointer(&given))
	return HackGetHiddenString(givenHeader.Data, givenHeader.Len)
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

// ProxyEnv represents some proxy environment variables, returned by GetProxyEnv.
type ProxyEnv struct {
	NoProxy    string
	HttpProxy  string
	HttpsProxy string
	SocksProxy string
}

// GetProxyEnv lookups and returns four proxy environments, including no_proxy, http_proxy, https_proxy and socks_proxy.
func GetProxyEnv() *ProxyEnv {
	np := strings.TrimSpace(os.Getenv("no_proxy"))
	hp := strings.TrimSpace(os.Getenv("http_proxy"))
	hsp := strings.TrimSpace(os.Getenv("https_proxy"))
	ssp := strings.TrimSpace(os.Getenv("socks_proxy"))
	return &ProxyEnv{NoProxy: np, HttpProxy: hp, HttpsProxy: hsp, SocksProxy: ssp}
}

// PrintLog prints each proxy variables if it is not empty, using given logFunc (default to log.Println) and prefix (default to empty).
//
// Example:
// 	proxyEnv := xruntime.GetProxyEnv()
// 	proxyEnv.PrintLog(nil, "[Gin] ")
func (p *ProxyEnv) PrintLog(logFunc func(string), prefix string) {
	if logFunc == nil {
		logFunc = func(s string) {
			log.Println(s)
		}
	}

	if p.NoProxy != "" {
		logFunc(fmt.Sprintf("%sUsing no_proxy: %s", prefix, p.NoProxy))
	}
	if p.HttpProxy != "" {
		logFunc(fmt.Sprintf("%sUsing http_proxy: %s", prefix, p.HttpProxy))
	}
	if p.HttpsProxy != "" {
		logFunc(fmt.Sprintf("%sUsing https_proxy: %s", prefix, p.HttpsProxy))
	}
	if p.SocksProxy != "" {
		logFunc(fmt.Sprintf("%sUsing socks_proxy: %s", prefix, p.SocksProxy))
	}
}

// NetAddrType describes a concrete address type, including TCPAddrType, UDPAddrType, IPAddrType and UnixAddrType.
type NetAddrType string

const (
	// TCPAddrType represents a net.TCPAddr concrete address type.
	TCPAddrType NetAddrType = "TCPAddr"

	// UDPAddrType represents a net.UDPAddr concrete address type.
	UDPAddrType NetAddrType = "UDPAddr"

	// IPAddrType represents a net.IPAddr concrete address type.
	IPAddrType NetAddrType = "IPAddr"

	// UnixAddrType represents a net.UnixAddr concrete address type.
	UnixAddrType NetAddrType = "UnixAddr"
)

// String returns the string value of NetAddrType.
func (n NetAddrType) String() string {
	return string(n)
}

// ConcreteNetAddr represents a concrete net.Addr type, which includes implementation of net.TCPAddr, net.UDPAddr, net.IPAddr and net.UnixAddr.
type ConcreteNetAddr struct {
	Type NetAddrType

	TCPAddr  *net.TCPAddr
	UDPAddr  *net.UDPAddr
	IPAddr   *net.IPAddr
	UnixAddr *net.UnixAddr

	// for TCP, UDP, IP
	IP   net.IP
	Zone string

	// for TCP, UDP
	Port int

	// for Unix
	Name string
	Net  string
}

// ParseNetAddr parses given net.Addr value to ConcreteNetAddr, returns error if given address is not TCP, UDP, IP, Unix address.
func ParseNetAddr(addr net.Addr) (*ConcreteNetAddr, bool) {
	if addr == nil {
		return nil, false
	}

	var out *ConcreteNetAddr
	switch addr := addr.(type) {
	case *net.TCPAddr:
		out = &ConcreteNetAddr{Type: TCPAddrType, TCPAddr: addr, IP: addr.IP, Zone: addr.Zone, Port: addr.Port}
	case *net.UDPAddr:
		out = &ConcreteNetAddr{Type: UDPAddrType, UDPAddr: addr, IP: addr.IP, Zone: addr.Zone, Port: addr.Port}
	case *net.IPAddr:
		out = &ConcreteNetAddr{Type: IPAddrType, IPAddr: addr, IP: addr.IP, Zone: addr.Zone}
	case *net.UnixAddr:
		out = &ConcreteNetAddr{Type: UnixAddrType, UnixAddr: addr, Name: addr.Name, Net: addr.Net}
	default:
		return nil, false
	}
	return out, true
}
