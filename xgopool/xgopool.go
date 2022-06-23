package xgopool

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
)

// GoPool represents a simple goroutine pool with workers capacity, panic handler, worker pool, task pool and task queue. Please
// visit https://github.com/bytedance/gopkg/blob/develop/util/gopool/gopool.go for more details.
type GoPool struct {
	workersCap   int32 // atomic
	panicHandler func(context.Context, interface{})

	workerPool  *sync.Pool
	numWorkers  int32 // atomic
	workerMutex *sync.Mutex

	taskPool  *sync.Pool
	numTasks  int32 // atomic
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
		workersCap: cap,
		panicHandler: func(ctx context.Context, i interface{}) {
			log.Printf("xgopool warning: Goroutine panicked with `%v`", i)
		},
		workerPool:  &sync.Pool{New: func() interface{} { return &worker{} }},
		workerMutex: &sync.Mutex{},
		taskPool:    &sync.Pool{New: func() interface{} { return &task{} }},
		taskMutex:   &sync.Mutex{},
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

// Go creates a task and waits for a worker to be scheduled, and invokes the task function.
func (g *GoPool) Go(f func()) {
	if f != nil {
		g.CtxGo(context.Background(), func(context.Context) {
			f()
		})
	}
}

// CtxGo creates a task and waits for a worker to be scheduled and invoke the task function. Note that function in this method
// takes context.Context as parameter.
func (g *GoPool) CtxGo(ctx context.Context, f func(context.Context)) {
	if f != nil {
		t := g.getTask(ctx, f)
		g.enqueueTask(t) // numTasks++
		if g.NumWorkers() < g.WorkersCap() {
			w := g.getWorker() // numWorkers++
			go w.start()
		}
	}
}

// task represents a goroutine task, with context.Context, given function and next pointer for task linked list.
type task struct {
	ctx  context.Context
	f    func(context.Context)
	next *task
}

// getTask returns an empty task structure from task sync.Pool and initializes fields.
func (g *GoPool) getTask(ctx context.Context, f func(context.Context)) *task {
	t := g.taskPool.Get().(*task)
	t.ctx = ctx
	t.f = f
	t.next = nil
	return t
}

// recycleTask empties given task structure and recycles to task sync.Pool.
func (g *GoPool) recycleTask(t *task) {
	t.ctx = nil
	t.f = nil
	t.next = nil
	g.taskPool.Put(t)
}

// enqueueTask enqueues given task to GoPool's task linked list and updates numTasks.
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

// dequeueTask dequeues a task from the head of GoPool's task linked list and updates numTasks, returns false if the task list is empty.
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

// worker represents a goroutine worker, and is used to execute task.
type worker struct {
	g *GoPool
}

// getWorker returns an empty worker structure from worker sync.Pool and updates numWorkers.
func (g *GoPool) getWorker() *worker {
	g.workerMutex.Lock()
	defer g.workerMutex.Unlock()
	w := g.workerPool.Get().(*worker)
	w.g = g
	atomic.AddInt32(&g.numWorkers, 1)
	return w
}

// recycleWorker recycles to worker sync.Pool and updates numWorkers.
func (g *GoPool) recycleWorker(w *worker) {
	g.workerMutex.Lock()
	defer g.workerMutex.Unlock()
	w.g = nil
	g.workerPool.Put(w)
	atomic.AddInt32(&g.numWorkers, -1)
}

// _testFlag is only used when testing the xgopool package, it represents that now is testing if it equals to true.
var _testFlag atomic.Value

// start dequeues a task from the head of GoPool's task linked list, and invokes given function with panic handler.
func (w *worker) start() {
	defer w.g.recycleWorker(w) // numWorkers--
	for {
		t, ok := w.g.dequeueTask() // numTasks--
		if !ok {
			break
		}
		func() {
			defer func() {
				if hdl := w.g.panicHandler; hdl != nil {
					if i := recover(); i != nil {
						hdl(t.ctx, i)
					}
				} else if _testFlag.Load() == true {
					// enter only when testing xgopool package
					if i := recover(); i != nil {
						defer func() {
							log.Printf("Panic when testing: `%v`", i)
							_testFlag.Store(false)
						}()
					}
				}
			}()
			t.f(t.ctx)
		}()
		w.g.recycleTask(t)
	}
}

// _defaultPool is a global GoPool with capacity 10000.
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

// Go creates a task and waits for a worker to be scheduled, and invokes the task function.
func Go(f func()) {
	_defaultPool.Go(f)
}

// CtxGo creates a task and waits for a worker to be scheduled and invokes the task function. Note that function in this method
// takes context.Context as parameter.
func CtxGo(ctx context.Context, f func(context.Context)) {
	_defaultPool.CtxGo(ctx, f)
}
