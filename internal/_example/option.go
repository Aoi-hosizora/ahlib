package _example

import (
	"fmt"
)

type Service struct {
	options *serviceOptions
}

type serviceOptions struct {
	Arg1 int
	Arg2 string
	Arg3 bool
}

func (s *Service) DoSomething() string {
	return fmt.Sprintf("%d-%s-%v", s.options.Arg1, s.options.Arg2, s.options.Arg3)
}

type ServiceOption func(*serviceOptions)

func New(options ...ServiceOption) *Service {
	so := &serviceOptions{
		Arg1: 2, // default option value
		Arg3: true,
	}
	for _, o := range options {
		if o != nil {
			o(so) // apply option
		}
	}
	return &Service{options: so}
}

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
