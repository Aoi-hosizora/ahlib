package xgopool

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
)

// GoPool represents a simple goroutine pool with workers capacity, panic handler, worker pool, task pool and related fields.
type GoPool struct {
	workersCap   int32
	panicHandler func(context.Context, interface{})

	workerPool  *sync.Pool
	numWorkers  int32
	workerMutex *sync.Mutex

	taskPool  *sync.Pool
	numTasks  int32
	taskMutex *sync.Mutex
	taskHead  *task
	taskTail  *task
}

const (
	panicNonPositiveCap = "xgopool: non-positive workers capacity"
)

// New creates an empty GoPool with given workers capacity.
func New(cap int32) *GoPool {
	if cap <= 0 {
		panic(panicNonPositiveCap)
	}
	return &GoPool{
		workersCap:   cap,
		panicHandler: func(ctx context.Context, i interface{}) { log.Printf("Warning: Panic with `%v`", i) },
		workerPool:   &sync.Pool{New: func() interface{} { return &worker{} }},
		workerMutex:  &sync.Mutex{},
		taskPool:     &sync.Pool{New: func() interface{} { return &task{} }},
		taskMutex:    &sync.Mutex{},
	}
}

// SetWorkersCap sets workers capacity dynamically.
func (g *GoPool) SetWorkersCap(cap int32) {
	if cap <= 0 {
		panic(panicNonPositiveCap)
	}
	atomic.StoreInt32(&g.workersCap, cap)
}

// SetPanicHandler sets panic handlers for goroutine function invoking.
func (g *GoPool) SetPanicHandler(handler func(context.Context, interface{})) {
	g.panicHandler = handler
}

// WorkersCap returns the current workers capacity.
func (g *GoPool) WorkersCap() int32 {
	return atomic.LoadInt32(&g.workersCap)
}

// NumWorkers returns the current workers count.
func (g *GoPool) NumWorkers() int32 {
	return atomic.LoadInt32(&g.numWorkers)
}

// NumTasks returns the count of current workers waiting.
func (g *GoPool) NumTasks() int32 {
	return atomic.LoadInt32(&g.numTasks)
}

// Go creates a task and waits for a worker to be scheduled and invoke the task function.
func (g *GoPool) Go(f func()) {
	if f == nil {
		return
	}
	g.CtxGo(context.Background(), func(ctx context.Context) { f() })
}

// CtxGo creates a task and waits for a worker to be scheduled and invoke the task function, this method can take context.Context as parameter.
func (g *GoPool) CtxGo(ctx context.Context, f func(context.Context)) {
	if f == nil {
		return
	}
	t := g.getTask(ctx, f)
	g.enqueueTask(t) // numTasks++

	g.workerMutex.Lock()
	if g.NumWorkers() < g.WorkersCap() {
		w := g.getWorker() // numWorkers++
		go w.start(g)
	}
	g.workerMutex.Unlock()
}

// task represents a goroutine task, with context.Context, given function and next pointer for linked list.
type task struct {
	ctx  context.Context
	f    func(context.Context)
	next *task
}

// getTask returns an empty task structure from task sync.Pool and sets the given parameters.
func (g *GoPool) getTask(ctx context.Context, f func(context.Context)) *task {
	t := g.taskPool.Get().(*task)
	t.ctx = ctx
	t.f = f
	t.next = nil
	return t
}

// recycleTask empties the given task structure and recycles to task sync.Pool.
func (g *GoPool) recycleTask(t *task) {
	t.ctx = nil
	t.f = nil
	t.next = nil
	g.taskPool.Put(t)
}

// enqueueTask enqueues the given task to GoPool's task linked list.
func (g *GoPool) enqueueTask(t *task) {
	g.taskMutex.Lock()
	defer g.taskMutex.Unlock()
	if g.taskHead == nil {
		g.taskHead = t
		g.taskTail = t
	} else {
		g.taskTail.next = t
		g.taskTail = t
	}
	atomic.AddInt32(&g.numTasks, 1)
}

// dequeueTask dequeues a task from the head of GoPool's task linked list, returns false if the task list is empty.
func (g *GoPool) dequeueTask() (*task, bool) {
	g.taskMutex.Lock()
	defer g.taskMutex.Unlock()
	if g.taskHead == nil {
		return nil, false
	}
	t := g.taskHead
	g.taskHead = g.taskHead.next
	atomic.AddInt32(&g.numTasks, -1)
	return t, true
}

// worker represents a goroutine worker, used to execute task.
type worker struct{}

// getWorker returns an empty worker structure from worker sync.Pool.
func (g *GoPool) getWorker() *worker {
	atomic.AddInt32(&g.numWorkers, 1)
	return g.workerPool.Get().(*worker)
}

// recycleWorker recycles to worker sync.Pool.
func (g *GoPool) recycleWorker(w *worker) {
	g.workerPool.Put(w)
	atomic.AddInt32(&g.numWorkers, -1)
}

// _testFlag is only used when testing the xgopool package, `true` value represents that now is testing.
var _testFlag atomic.Value

// start dequeues a task from the head of GoPool's task linked list, and invokes the given function with panic handler.
func (w *worker) start(g *GoPool) {
	for {
		t, ok := g.dequeueTask() // numTasks--
		if !ok {
			break
		}
		func() {
			defer func() {
				if err := recover(); err != nil {
					if g.panicHandler != nil {
						g.panicHandler(t.ctx, err)
					} else {
						if _testFlag.Load() == true {
							// enter only when testing the xgopool package, needn't worry about the performance
							defer func() {
								log.Printf("Panic when testing: `%v`", recover())
								_testFlag.Store(false)
							}()
						}
						panic(err)
					}
				}
			}()
			t.f(t.ctx)
		}()
		g.recycleTask(t)
	}
	g.workerMutex.Lock()
	g.recycleWorker(w) // numWorkers--
	g.workerMutex.Unlock()
}

// _defaultPool is a global GoPool.
var _defaultPool = New(10000)

// SetWorkersCap sets workers capacity dynamically.
func SetWorkersCap(cap int32) {
	_defaultPool.SetWorkersCap(cap)
}

// SetPanicHandler sets panic handlers for goroutine function invoking.
func SetPanicHandler(handler func(context.Context, interface{})) {
	_defaultPool.SetPanicHandler(handler)
}

// WorkersCap returns the current workers capacity.
func WorkersCap() int32 {
	return _defaultPool.WorkersCap()
}

// NumWorkers returns the current workers count.
func NumWorkers() int32 {
	return _defaultPool.NumWorkers()
}

// NumTasks returns the count of current workers waiting.
func NumTasks() int32 {
	return _defaultPool.NumTasks()
}

// Go creates a task and waits for a worker to be scheduled and invoke the task function.
func Go(f func()) {
	_defaultPool.Go(f)
}

// CtxGo creates a task and waits for a worker to be scheduled and invoke the task function, this method can take context.Context as parameter.
func CtxGo(ctx context.Context, f func(context.Context)) {
	_defaultPool.CtxGo(ctx, f)
}
