# xruntime

### Functions

+ `type Stack struct {}`
+ `(s *Stack) String() string`
+ `GetStack(skip int) []*Stack`
+ `GetStackWithInfo(skip int) (stacks []*Stack, filename string, funcname string, lineIndex int, line string)`
+ `PrintStacks(stacks []*Stack)`
