# xgopool

## Dependencies

+ xtesting*

## Documents

### Types

+ `type GoPool struct`

### Variables

+ None

### Constants

+ None

### Functions

+ `func New(cap int32) *GoPool`
+ `func SetWorkersCap(cap int32)`
+ `func SetPanicHandler(handler func(context.Context, interface{}))`
+ `func WorkersCap() int32`
+ `func NumWorkers() int32`
+ `func NumTasks() int32`
+ `func Go(f func())`
+ `func CtxGo(ctx context.Context, f func(context.Context))`

### Methods

+ `func (g *GoPool) SetWorkersCap(cap int32)`
+ `func (g *GoPool) SetPanicHandler(handler func(context.Context, interface{}))`
+ `func (g *GoPool) WorkersCap() int32`
+ `func (g *GoPool) NumWorkers() int32`
+ `func (g *GoPool) NumTasks() int32`
+ `func (g *GoPool) Go(f func())`
+ `func (g *GoPool) CtxGo(ctx context.Context, f func(context.Context))`
