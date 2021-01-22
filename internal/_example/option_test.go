package _example

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestOption(t *testing.T) {
	for _, tc := range []struct {
		name string
		give *serviceOptions
		want *serviceOptions
	}{
		// normal
		{"default", New().options, &serviceOptions{2, "", true}},
		{"WithArg1", New(WithArg1(0)).options, &serviceOptions{0, "", true}},
		{"WithArg2", New(WithArg2("test")).options, &serviceOptions{2, "test", true}},
		{"WithArg3", New(WithArg3(false)).options, &serviceOptions{2, "", false}},
		{"WithArgs", New(WithArg1(0), WithArg2("test"), WithArg3(false)).options, &serviceOptions{0, "test", false}},

		// abnormal
		{"nil_option", New(nil).options, &serviceOptions{2, "", true}},
		{"multi_options", New(WithArg1(0), WithArg1(2), WithArg1(1)).options, &serviceOptions{1, "", true}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			xtesting.EqualValue(t, tc.give, tc.want)
		})
	}

	service := New(WithArg1(0), WithArg2("test"), WithArg3(false))
	xtesting.Equal(t, service.DoSomething(), "0-test-false")
}
