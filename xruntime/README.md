# xruntime

## Dependencies

+ (xtesting)

## Documents

### Types

+ `type TraceFrame struct`
+ `type TraceStack []*TraceFrame`
+ `type ProxyEnv struct`
+ `type NetAddrType string`
+ `type ConcreteNetAddr struct`

### Variables

+ None

### Constants

+ `const PprofGoroutineProfile string`
+ `const PprofThreadcreateProfile string`
+ `const PprofHeapProfile string`
+ `const PprofAllocsProfile string`
+ `const PprofBlockProfle string`
+ `const PprofMutexProfile string`
+ `const TCPAddrType NetAddrType`
+ `const UDPAddrType NetAddrType`
+ `const IPAddrType NetAddrType`
+ `const UnixAddrType NetAddrType`

### Functions

+ `func NameOfFunction(f interface{}) string`
+ `func RawStack(all bool) []byte`
+ `func RuntimeTraceStack(skip uint) TraceStack`
+ `func RuntimeTraceStackWithInfo(skip uint) (stack TraceStack, filename string, funcName string, lineIndex int, lineText string)`
+ `func HackHideString(given unsafe.Pointer, givenLength int, hidden string) (dataAddr uintptr)`
+ `func HackHideStringAfterString(given, hidden string) string`
+ `func HackExtractHiddenString(given unsafe.Pointer, givenLength int) (hidden string)`
+ `func HackExtractHiddenStringAfterString(given string) (hidden string)`
+ `func SignalName(sig syscall.Signal) string`
+ `func SignalReadableName(sig syscall.Signal) string`
+ `func GetProxyEnv() *ProxyEnv`
+ `func ParseNetAddr(addr net.Addr) (*ConcreteNetAddr, bool)`

### Methods

+ `func (t *TraceFrame) String() string`
+ `func (t *TraceStack) String() string`
+ `func (p *ProxyEnv) PrintLog(logFunc func(string), prefix string)`
