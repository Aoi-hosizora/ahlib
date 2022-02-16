package xerror

import (
	"context"
	"errors"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"strconv"
	"testing"
	"time"
)

type stringError string

func (m stringError) Error() string {
	return string(m)
}

func TestMultiError(t *testing.T) {
	me := &multiError{errs: []error{}}
	xtesting.Equal(t, me.Errors(), []error{})
	xtesting.Equal(t, me.Error(), "")

	me = &multiError{errs: []error{nil}}
	xtesting.Equal(t, me.Errors(), []error{nil})
	xtesting.Panic(t, func() { _ = me.Error() })

	test := errors.New("test")
	me = &multiError{errs: []error{test}}
	xtesting.Equal(t, me.Errors(), []error{test})
	xtesting.Equal(t, me.Error(), "test")
	xtesting.Equal(t, errors.Is(me, test), true)
	xtesting.Equal(t, errors.Is(me, errors.New("test")), false)
	e1 := new(error)
	xtesting.Equal(t, errors.As(me, e1), true)
	xtesting.Equal(t, *e1, me) // <<<
	e2 := &strconv.NumError{}
	xtesting.Equal(t, errors.As(me, &e2), false)

	test1, test2 := errors.New("test1"), stringError("test2")
	me = &multiError{errs: []error{test1, test2}}
	xtesting.Equal(t, me.Errors(), []error{test1, test2})
	xtesting.Equal(t, me.Error(), "test1; test2")
	xtesting.Equal(t, errors.Is(me, test1), true)
	xtesting.Equal(t, errors.Is(me, test2), true)
	xtesting.Equal(t, errors.Is(me, errors.New("test2")), false)
	e3 := new(error)
	xtesting.Equal(t, errors.As(me, e3), true)
	xtesting.Equal(t, *e3, me) // <<<
	e4 := new(stringError)
	xtesting.Equal(t, errors.As(me, e4), true)
	xtesting.Equal(t, *e4, test2)
}

func TestCombine(t *testing.T) {
	for _, tc := range []struct {
		give []error
		want error
	}{
		{[]error{}, nil},
		{[]error{nil}, nil},
		{[]error{nil, nil}, nil},
		{[]error{nil, nil, errors.New("1")}, errors.New("1")},
		{[]error{nil, stringError("1"), nil}, stringError("1")},
		{[]error{nil, errors.New("1"), nil, errors.New("2")}, &multiError{errs: []error{errors.New("1"), errors.New("2")}}},
		{[]error{&multiError{errs: []error{errors.New("1"), stringError("2")}}, nil, errors.New("3")},
			&multiError{errs: []error{errors.New("1"), stringError("2"), errors.New("3")}}},
		{[]error{nil, &multiError{errs: []error{stringError("1")}}, nil, &multiError{errs: []error{errors.New("2"), errors.New("3")}}},
			&multiError{errs: []error{stringError("1"), errors.New("2"), errors.New("3")}}},
		{[]error{nil, &multiError{errs: []error{errors.New("1"), nil}}, nil, &multiError{errs: []error{errors.New("2"), nil, stringError("3")}}},
			&multiError{errs: []error{errors.New("1"), nil, errors.New("2"), nil, stringError("3")}}}, // <<< nil
		{[]error{errors.New("1"), nil, &multiError{errs: []error{nil}}, &multiError{errs: []error{nil, nil}}},
			&multiError{errs: []error{errors.New("1"), nil, nil, nil}}}, // <<< nil
		{[]error{&multiError{errs: []error{&multiError{errs: []error{stringError("?")}}}}},
			&multiError{errs: []error{&multiError{errs: []error{stringError("?")}}}}}, // <<<
	} {
		xtesting.Equal(t, Combine(tc.give...), tc.want)
	}

	t.Run("Append", func(t *testing.T) {
		err := errors.New("1")
		for i := 2; i <= 10; i++ {
			err = Combine(err, errors.New(strconv.Itoa(i))) // <- Append
		}
		me, ok := err.(MultiError)
		xtesting.True(t, ok)
		xtesting.Equal(t, len(me.Errors()), 10)
		for i := 0; i < 10; i++ {
			xtesting.Equal(t, me.Errors()[i].Error(), strconv.Itoa(i+1))
		}
	})
}

func TestSeparate(t *testing.T) {
	for _, tc := range []struct {
		give error
		want []error
	}{
		{nil, nil},
		{errors.New("test"), []error{errors.New("test")}},
		{stringError("test"), []error{stringError("test")}},
		{&multiError{}, []error{}},
		{&multiError{errs: []error{nil}}, []error{nil}},                               // <<< nil
		{&multiError{errs: []error{nil, nil}}, []error{nil, nil}},                     // <<< nil
		{&multiError{errs: []error{errors.New("test")}}, []error{errors.New("test")}}, // <<< nil
		{&multiError{errs: []error{errors.New("1"), stringError("2")}}, []error{errors.New("1"), stringError("2")}},
		{&multiError{errs: []error{nil, errors.New("test")}}, []error{nil, errors.New("test")}}, // <<< nil
		{&multiError{errs: []error{nil, stringError("1"), &multiError{errs: []error{stringError("2")}}}},
			[]error{nil, stringError("1"), &multiError{errs: []error{stringError("2")}}}}, // <<<
	} {
		xtesting.Equal(t, Separate(tc.give), tc.want)
	}
}

func TestErrorGroup(t *testing.T) {
	t.Run("Zero ErrorGroup", func(t *testing.T) {
		eg := &ErrorGroup{}
		eg.SetGoExecutor(DefaultExecutor)
		// 1. test context and panic
		eg.Go(func(ctx context.Context) error {
			xtesting.Equal(t, ctx, context.Background())
			return nil
		})
		eg.Go(func(ctx context.Context) error {
			panic("test") // will be recovered
		})
		xtesting.Nil(t, eg.Wait())
		// 2. test error and deadlock
		eg.Go(func(ctx context.Context) error {
			return errors.New("test")
		})
		eg.Go(func(ctx context.Context) error {
			// select {
			// case <-ctx.Done(): // cannot done context => deadlock
			// }
			return nil
		})
		xtesting.Equal(t, eg.Wait(), errors.New("test"))
		// 3. test executor
		eg = &ErrorGroup{}
		xtesting.Nil(t, eg.goExecutor)
		eg.Go(func(ctx context.Context) error { return nil }) // set eg.goExecutor
		xtesting.Nil(t, eg.Wait())
		xtesting.NotNil(t, eg.goExecutor)
	})

	t.Run("NewErrorGroup", func(t *testing.T) {
		eg := NewErrorGroup(context.WithValue(context.Background(), "key", "value"))
		eg.SetGoExecutor(DefaultExecutor)
		// 1. test context and panic
		xtesting.NotPanic(t, func() { eg.Go(nil) })
		eg.Go(func(ctx context.Context) error {
			xtesting.Equal(t, ctx.Value("key"), "value")
			return nil
		})
		eg.Go(func(ctx context.Context) error {
			panic("test") // will be recovered
		})
		xtesting.Nil(t, eg.Wait())
		// 2. test error and cancelation
		eg.Go(func(ctx context.Context) error {
			return errors.New("test")
		})
		var outErr error
		eg.Go(func(ctx context.Context) error {
			select {
			case <-ctx.Done(): // deadlock
				outErr = ctx.Err()
			}
			return nil
		})
		xtesting.Equal(t, eg.Wait(), errors.New("test"))
		xtesting.Equal(t, outErr, errors.New("context canceled"))
		// 3. test timeout
		f := func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Millisecond * 200):
			}
			return nil
		}
		ctx1, cancel1 := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel1()
		eg = NewErrorGroup(ctx1)
		eg.Go(f)
		xtesting.Equal(t, eg.Wait().Error(), "context deadline exceeded") // ctx.Done
		ctx2, cancel2 := context.WithTimeout(context.Background(), time.Millisecond*300)
		defer cancel2()
		eg = NewErrorGroup(ctx2)
		eg.Go(f)
		xtesting.Nil(t, eg.Wait()) // time.After
	})
}
