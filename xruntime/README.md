# xruntime

## Dependencies

+ None

## Documents

### Types

+ `type TraceFrame struct`
+ `type TraceStack []*TraceFrame`

### Variables

+ None

### Constants

+ None

### Functions

+ `func RuntimeTraceStack(skip int) TraceStack`
+ `func RuntimeTraceStackWithInfo(skip int) (stack TraceStack, filename string, funcname string, lineIndex int, lineText string)`

### Methods

+ `func (t *TraceFrame) String() string`
+ `func (t *TraceStack) String() string`
