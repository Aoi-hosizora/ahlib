package xgopool

import (
	"context"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestSimpleGo(t *testing.T) {
	t.Run("cap", func(t *testing.T) {
		xtesting.Panic(t, func() { New(0) })
		xtesting.Panic(t, func() { New(-1) })
		xtesting.NotPanic(t, func() { New(1) })

		serializedPool := New(1)
		xtesting.Equal(t, serializedPool.WorkersCap(), int32(1))
		xtesting.Panic(t, func() { serializedPool.SetWorkersCap(0) })
		xtesting.Panic(t, func() { serializedPool.SetWorkersCap(-1) })
		xtesting.NotPanic(t, func() { serializedPool.SetWorkersCap(2) })
		xtesting.Equal(t, serializedPool.WorkersCap(), int32(2))
		xtesting.NotPanic(t, func() { serializedPool.SetWorkersCap(100) })
		xtesting.Equal(t, serializedPool.WorkersCap(), int32(100))

		xtesting.Equal(t, WorkersCap(), int32(10000))
		SetWorkersCap(100)
		xtesting.Equal(t, WorkersCap(), int32(100))
		SetWorkersCap(10000)
		xtesting.Equal(t, WorkersCap(), int32(10000))
		xtesting.Equal(t, NumWorkers(), int32(0))
		xtesting.Equal(t, NumTasks(), int32(0))
	})

	t.Run("incr", func(t *testing.T) {
		wg := sync.WaitGroup{}
		tmp := 0
		incr := func() { tmp++; wg.Done() }

		wg.Add(1)
		Go(incr)
		wg.Wait()
		xtesting.Equal(t, tmp, 1)
		//
		wg.Add(1)
		Go(incr)
		wg.Wait()
		xtesting.Equal(t, tmp, 2)
		//
		numTasks := _defaultPool.numTasks
		Go(nil)
		xtesting.Equal(t, _defaultPool.numTasks, numTasks)
		xtesting.Equal(t, tmp, 2)
		//
		wg.Add(1)
		Go(incr)
		wg.Wait()
		xtesting.Equal(t, tmp, 3)

		xtesting.Equal(t, NumWorkers(), int32(0))
		xtesting.Equal(t, NumTasks(), int32(0))
	})

	t.Run("decr", func(t *testing.T) {
		wg := sync.WaitGroup{}
		tmp := 3
		decr := func(_ context.Context) { tmp--; wg.Done() }

		wg.Add(1)
		CtxGo(context.Background(), decr)
		wg.Wait()
		xtesting.Equal(t, tmp, 2)
		//
		wg.Add(1)
		CtxGo(context.Background(), decr)
		wg.Wait()
		xtesting.Equal(t, tmp, 1)
		//
		numTasks := _defaultPool.numTasks
		CtxGo(context.Background(), nil)
		xtesting.Equal(t, _defaultPool.numTasks, numTasks)
		xtesting.Equal(t, tmp, 1)
		//
		wg.Add(1)
		CtxGo(context.Background(), decr)
		wg.Wait()
		xtesting.Equal(t, tmp, 0)

		xtesting.Equal(t, NumWorkers(), int32(0))
		xtesting.Equal(t, NumTasks(), int32(0))
	})

	t.Run("atomic add", func(t *testing.T) {
		wg := sync.WaitGroup{}
		p := New(100)
		n := int32(0)
		for i := 0; i < 2000; i++ {
			wg.Add(1)
			p.Go(func() {
				atomic.AddInt32(&n, 1)
				wg.Done()
			})
		}
		wg.Wait()
		xtesting.Equal(t, n, int32(2000))

		p.SetWorkersCap(2000)
		n = 0
		for i := 0; i < 2000; i++ {
			wg.Add(1)
			p.Go(func() {
				atomic.AddInt32(&n, 1)
				wg.Done()
			})
		}
		wg.Wait()
		xtesting.Equal(t, n, int32(2000))
	})
}

func TestPanicGo(t *testing.T) {
	terminated := make(chan struct{})
	done := func() { terminated <- struct{}{} }
	wait := func() { <-terminated }

	t.Run("default panic handler", func(t *testing.T) {
		originHandler := _defaultPool.panicHandler
		SetPanicHandler(func(ctx context.Context, i interface{}) {
			originHandler(ctx, i)
			done()
		})
		defer SetPanicHandler(originHandler)
		Go(func() { panic("panic") })
		// Warning: Panicked with `panic`
		wait()
	})

	t.Run("panic with no panic handler 1", func(t *testing.T) {
		g := New(100)
		g.SetPanicHandler(nil)
		g.Go(func() {
			defer done()
			defer func() { err := recover(); xtesting.Equal(t, err, "panic") }()
			panic("panic")
		})
		wait()
	})

	t.Run("panic with no panic handler 2", func(t *testing.T) {
		g := New(100)
		g.SetPanicHandler(nil)
		_testFlag = true
		g.Go(func() {
			panic("panic for testing ...")
		})
		// panic for testing ...
		for _testFlag {
		}
	})

	t.Run("panic with msg panic handler", func(t *testing.T) {
		g := New(100)
		msg := new(string)
		g.SetPanicHandler(func(_ context.Context, i interface{}) {
			*msg = fmt.Sprintf("%v_%v", i, i)
			done()
		})
		g.Go(func() { panic("panic") })
		wait()
		xtesting.Equal(t, *msg, "panic_panic")
	})

	t.Run("panic with context panic handler", func(t *testing.T) {
		g := New(100)
		msg := new(string)
		g.SetPanicHandler(func(ctx context.Context, i interface{}) {
			*msg = fmt.Sprintf("%v_%v", i, ctx.Value("id"))
			done()
		})
		g.CtxGo(context.WithValue(context.Background(), "id", 333), func(ctx context.Context) {
			xtesting.Equal(t, ctx.Value("id"), 333)
			panic("panic")
		})
		wait()
		xtesting.Equal(t, *msg, "panic_333")
	})
}

func TestWorkersCap(t *testing.T) {
	const N = 5000
	integers := new([]int)
	matchedIntegers := make([]int, 0, N)
	for i := 0; i < N; i++ {
		matchedIntegers = append(matchedIntegers, i)
	}

	t.Run("serialize", func(t *testing.T) {
		*integers = make([]int, 0, N)
		serializedPool := New(1)
		for i := 0; i < N; i++ {
			i := i
			serializedPool.Go(func() {
				xtesting.Equal(t, serializedPool.NumWorkers(), int32(1))
				*integers = append(*integers, i)
			})
		}
		terminated := make(chan struct{})
		serializedPool.Go(func() {
			xtesting.Equal(t, *integers, matchedIntegers)
			close(terminated)
		})
		<-terminated
	})

	t.Run("serialized parallel", func(t *testing.T) {
		*integers = make([]int, 0, N)
		parallelizedPool := New(2)
		terminated := make(chan struct{})
		for i := 0; i < N; i++ {
			i := i
			parallelizedPool.Go(func() {
				xtesting.True(t, parallelizedPool.NumWorkers() <= 2)
				xtesting.Equal(t, parallelizedPool.NumTasks(), int32(0))
				*integers = append(*integers, i)
				terminated <- struct{}{}
			})
			<-terminated
		}
		close(terminated)
		xtesting.Equal(t, *integers, matchedIntegers)
	})

	t.Run("parallel", func(t *testing.T) {
		*integers = make([]int, 0, N)
		parallelizedPool := New(5)
		mu := sync.Mutex{}
		wg := sync.WaitGroup{}
		for i := 0; i < N; i++ {
			i := i
			wg.Add(1)
			parallelizedPool.Go(func() {
				mu.Lock()
				*integers = append(*integers, i)
				mu.Unlock()
				wg.Done()
			})
		}
		wg.Wait()
		xtesting.ElementMatch(t, *integers, matchedIntegers)
	})
}

// benchmarks are referred from https://github.com/bytedance/gopkg/blob/develop/util/gopool/pool_test.go.

const benchmarkTimes = 20000

func doCopyStack(_, b int) int {
	if b < 100 {
		return doCopyStack(0, b+1)
	}
	return 0
}

func testFunc() {
	doCopyStack(0, 0)
}

func BenchmarkPool(b *testing.B) {
	b.Run("GoPool_xgopool", func(b *testing.B) {
		p := New(int32(runtime.GOMAXPROCS(0)))
		var wg sync.WaitGroup
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(benchmarkTimes)
			for j := 0; j < benchmarkTimes; j++ {
				p.CtxGo(context.Background(), func(_ context.Context) {
					testFunc()
					wg.Done()
				})
			}
			wg.Wait()
		}
	})
	// "github.com/bytedance/gopkg/util/gopool"
	// b.Run("GoPool_gopkg", func(b *testing.B) {
	// 	p := gopool.NewPool("", int32(runtime.GOMAXPROCS(0)), gopool.NewConfig())
	// 	var wg sync.WaitGroup
	// 	b.ReportAllocs()
	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		wg.Add(benchmarkTimes)
	// 		for j := 0; j < benchmarkTimes; j++ {
	// 			p.Go(func() {
	// 				testFunc()
	// 				wg.Done()
	// 			})
	// 		}
	// 		wg.Wait()
	// 	}
	// })
	b.Run("StdGo", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(benchmarkTimes)
			for j := 0; j < benchmarkTimes; j++ {
				go func() {
					testFunc()
					wg.Done()
				}()
			}
			wg.Wait()
		}
	})
}

/*
	goos: windows
	goarch: amd64
	pkg: github.com/Aoi-hosizora/ahlib/xgopool
	cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
	BenchmarkPool
	BenchmarkPool/GoPool_xgopool
	BenchmarkPool/GoPool_xgopool-8         	      52	  20352171 ns/op	  383157 B/op	   22036 allocs/op
	BenchmarkPool/GoPool_gopkg
	BenchmarkPool/GoPool_gopkg-8           	     128	  18193410 ns/op	  367304 B/op	   22530 allocs/op
	BenchmarkPool/StdGo
	BenchmarkPool/StdGo-8                  	      33	  70946273 ns/op	  320000 B/op	   20000 allocs/op
	PASS
*/
