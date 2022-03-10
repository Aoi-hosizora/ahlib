# xruntime

## Dependencies

+ xtesting*

## Documents

### Types

+ `type TraceFrame struct`
+ `type TraceStack []*TraceFrame`

### Variables

+ None

### Constants

+ `const PprofGoroutineProfile string`
+ `const PprofThreadcreateProfile string`
+ `const PprofHeapProfile string`
+ `const PprofAllocsProfile string`
+ `const PprofBlockProfle string`
+ `const PprofMutexProfile string`

### Functions

+ `func RawStack(all bool) []byte`
+ `func RuntimeTraceStack(skip int) TraceStack`
+ `func RuntimeTraceStackWithInfo(skip int) (stack TraceStack, filename string, funcName string, lineIndex int, lineText string)`
+ `func SignalName(sig syscall.Signal) string`
+ `func SignalReadableName(sig syscall.Signal) string`
+ `func GetProxyEnv() (httpProxy string, httpsProxy string, socksProxy string)`

### Methods

+ `func (t *TraceFrame) String() string`
+ `func (t *TraceStack) String() string`
