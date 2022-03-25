# xerror

## Dependencies

+ (xtesting)

## Documents

### Types

+ `type Wrapper interface`
+ `type Matcher interface`
+ `type Assigner interface`
+ `type MultiError interface`
+ `type ErrorGroup struct`

### Variables

+ `var DefaultExecutor func`

### Constants

+ None

### Functions

+ `func Combine(errs ...error) error`
+ `func Separate(err error) []error`
+ `func NewErrorGroup(ctx context.Context) *ErrorGroup`

### Methods

+ `func (eg *ErrorGroup) SetGoExecutor(executor func(f func()))`
+ `func (eg *ErrorGroup) Go(f func(ctx context.Context) error)`
+ `func (eg *ErrorGroup) Wait() error`
