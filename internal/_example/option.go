package _example

import (
	"fmt"
)

// 1. Define xxxOptions or xxxConfig struct, unexported.

type serviceOptions struct {
	Arg1 int
	Arg2 string
	Arg3 bool
}

// 2. Define XXXOption func, with xxxOptions as function parameter.

type ServiceOption func(*serviceOptions)

// 3. Define some WithXXX functions with XXXOption as returned value.

func WithArg1(i int) ServiceOption {
	return func(o *serviceOptions) {
		o.Arg1 = i
	}
}

func WithArg2(s string) ServiceOption {
	return func(o *serviceOptions) {
		o.Arg2 = s
	}
}

func WithArg3(b bool) ServiceOption {
	return func(o *serviceOptions) {
		o.Arg3 = b
	}
}

// 4. Use xxxOptions to destination struct, use a New function to create this struct.

type Service struct {
	options *serviceOptions
}

func New(options ...ServiceOption) *Service {
	// 4.1 setup default option value
	opts := &serviceOptions{
		Arg1: 2,
		Arg3: true,
	}
	// 4.2 apply options
	for _, option := range options {
		if option != nil {
			option(opts)
		}
	}
	return &Service{options: opts}
}

func (s *Service) DoSomething() string {
	return fmt.Sprintf("%d-%s-%v", s.options.Arg1, s.options.Arg2, s.options.Arg3)
}
