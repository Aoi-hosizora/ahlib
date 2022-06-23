package xerror

import (
	"context"
	"errors"
	"strings"
	"sync"
)

// ==========
// interfaces
// ==========

// Wrapper is an interface used to identify errors which has Unwrap method, can be used in errors.Unwrap function.
type Wrapper interface {
	Unwrap() error
}

// Matcher is an interface used to identify errors which has Is method, can be used in errors.Is function.
type Matcher interface {
	Is(error) bool
}

// Assigner is an interface used to identify errors which has As method, can be used in errors.As function.
type Assigner interface {
	As(interface{}) bool
}

// ===========
// multi error
// ===========

// MultiError is an interface representing error groups, types implement this interface can be returned by xerror.Combine or
// several methods in github.com/uber-go/multierr package.
type MultiError interface {
	Errors() []error
}

// multiError is an unexported error type implements MultiError interface, can be returned by xerror.Combine.
type multiError struct {
	errs []error
}

var (
	_ MultiError = (*multiError)(nil)
)

// Errors implements MultiError interface.
func (m *multiError) Errors() []error {
	return m.errs // items are all non-nillable, if used in a safe manner
}

// Is implements Matcher interface.
func (m *multiError) Is(target error) bool {
	for _, err := range m.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// As implements Assigner interface.
func (m *multiError) As(target interface{}) bool {
	for _, err := range m.errs {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

// Error implements error interface.
func (m *multiError) Error() string {
	switch len(m.errs) {
	case 0:
		return ""
	case 1:
		return m.errs[0].Error() // non-nillable
	}
	sb := strings.Builder{}
	for _, err := range m.errs {
		if sb.Len() > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(err.Error()) // non-nillable
	}
	return sb.String()
}

// Combine combines given errors to a single error, there are some situations:
// 1. If pass empty errors, or all errors passed are nil, it will return a nil error.
// 2. If pass a single non-nil error, it will return this single error directly.
// 3. If more than one error passed are non-nil, it returns a MultiError containing all these non-nil errors.
// 4. If some errors are MultiError, the internal errors contained will be flatted.
func Combine(errs ...error) error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0] // maybe nil
	}
	notnil := make([]error, 0)
	for _, err := range errs {
		if err == nil {
			continue
		}
		if me, ok := err.(MultiError); ok {
			notnil = append(notnil, me.Errors()...)
		} else {
			notnil = append(notnil, err)
		}
	}
	switch len(notnil) {
	case 0:
		return nil
	case 1:
		return notnil[0] // single error (non-nil)
	default:
		return &multiError{errs: notnil} // multiple errors (all non-nil)
	}
}

// Separate separates given error to multiple errors that given error is composed of (that is MultiError). If given error is
// nil, a nil slice is returned.
func Separate(err error) []error {
	if err == nil {
		return nil
	}
	me, ok := err.(MultiError)
	if !ok {
		return []error{err}
	}
	errs := me.Errors()
	out := make([]error, len(errs))
	copy(out, errs)
	return out
}

// ===========
// error group
// ===========

// ErrorGroup is a sync.WaitGroup wrapper that can used to synchronization, error propagation, and context cancellation for
// groups of goroutines, refers to https://pkg.go.dev/golang.org/x/sync/errgroup for more details.
//
// A zero ErrorGroup is valid and does not cancel on error.
type ErrorGroup struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg      sync.WaitGroup
	err     error
	errOnce sync.Once

	mu         sync.RWMutex
	goExecutor func(f func())
}

// NewErrorGroup returns a new ErrorGroup with cancelable context derived from given context, and the default goroutine executor.
func NewErrorGroup(ctx context.Context) *ErrorGroup {
	ctx, cancel := context.WithCancel(ctx)
	return &ErrorGroup{ctx: ctx, cancel: cancel, goExecutor: defaultExecutor}
}

// defaultExecutor is the default goroutine executor for ErrorGroup, including create goroutine by `go` keyword and panic
// recovery with no logging.
var defaultExecutor = func(f func()) {
	go func() {
		defer func() {
			_ = recover()
		}()
		f()
	}()
}

// SetGoExecutor sets goroutine executor, can be used to change the behavior of `go` keyword, you can use this executor to
// add recover behavior for goroutine.
//
// Example:
// 	// custom recover behavior
// 	eg := NewErrorGroup(context.Background())
// 	eg.SetGoExecutor(func(f func()) {
// 		go func() {
// 			defer func() {
// 				if v := recover(); v != nil {
// 					log.Printf("Warning: Panic with %v", v)
// 				}
// 			}()
// 			f()
// 		}()
// 	})
//
// 	// use xgopool goroutine pool
// 	eg := NewErrorGroup(context.Background())
// 	gp := xgopool.New(runtime.NumCPU() * 10)
// 	gp.SetPanicHandler(func(_ context.Context, v interface{}) {
// 		log.Printf("Warning: Panic with %v", v)
// 	})
// 	eg.SetGoExecutor(func(f func()) {
// 		gp.Go(f)
// 	})
func (eg *ErrorGroup) SetGoExecutor(executor func(f func())) {
	if executor != nil {
		eg.mu.Lock()
		eg.goExecutor = executor
		eg.mu.Unlock()
	}
}

// Go calls given function in a new goroutine using specific executor. The first call to return a non-nil error cancels the
// group, its error will be returned by Wait.
//
// If using a zero ErrorGroup, ctx will be Background, otherwise it will be the context derived from given context passed
// to NewErrorGroup.
//
// Example:
// 	eg := NewErrorGroup(context.Background())
//
// 	// in select statement
// 	eg.Go(func(ctx context.Context) error {
// 		select {
// 		case ...:
// 		case <-ctx.Done():
// 		}
// 		return nil
// 	})
//
// 	// in cancelable http requesting
// 	eg.Go(func(ctx context.Context) error {
// 		req, _ := http.NewRequestWithContext(ctx, "GET", "...", nil)
// 		// ...
// 		return nil
// 	})
func (eg *ErrorGroup) Go(f func(ctx context.Context) error) {
	if f == nil {
		return
	}

	// get executor and context
	eg.mu.RLock()
	executor := eg.goExecutor
	ctx := eg.ctx
	eg.mu.RUnlock()
	if executor == nil {
		eg.mu.Lock()
		if eg.goExecutor == nil {
			eg.goExecutor = defaultExecutor
		}
		executor = eg.goExecutor
		eg.mu.Unlock()
	}
	if ctx == nil {
		ctx = context.Background()
	}

	// execute with goroutine
	eg.wg.Add(1)
	executor(func() {
		defer eg.wg.Done()

		err := f(ctx) // call given function
		if err != nil {
			eg.errOnce.Do(func() {
				eg.err = err // record the first error
				if eg.cancel != nil {
					eg.cancel() // cancel the context also
				}
			})
		}
	})
}

// Wait blocks until all function calls from the Go method have returned, then returns the first non-nil error (if any)
// from them.
func (eg *ErrorGroup) Wait() error {
	eg.wg.Wait()
	if eg.cancel != nil {
		eg.cancel()
	}
	return eg.err
}
