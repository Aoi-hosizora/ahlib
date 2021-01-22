package xmodule

import (
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"strings"
	"testing"
)

func TestModuleName(t *testing.T) {
	xtesting.Equal(t, ModuleName("").String(), "")
	xtesting.Equal(t, ModuleName("test").String(), "test")
}

func TestProvideName(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	for _, tc := range []struct {
		giveName   ModuleName
		giveModule interface{}
		wantPanic  bool
	}{
		{ModuleName(""), 0, true},
		{ModuleName("-"), 0, true},
		{ModuleName("~"), 0, true},
		{ModuleName("0"), nil, true},
		{ModuleName("int"), 12, false},                    // int
		{ModuleName("uint"), uint(12), false},             // uint
		{ModuleName("float"), 12.5, false},                // float
		{ModuleName("bool"), true, false},                 // bool
		{ModuleName("string"), "a", false},                // string
		{ModuleName("array"), [2]string{"1", "2"}, false}, // array
		{ModuleName("slice"), []string{"1", "2"}, false},  // slice
		{ModuleName("pointer"), &struct{}{}, false},       // pointer
		{ModuleName("struct"), struct{ int }{1}, false},   // struct
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { ProvideName(tc.giveName, tc.giveModule) })
		} else {
			ProvideName(tc.giveName, tc.giveModule)
			xtesting.Equal(t, _mc.provByName[tc.giveName], tc.giveModule)
		}
	}
}

func TestProvideType(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	for _, tc := range []struct {
		giveModule interface{}
		wantPanic  bool
	}{
		{nil, true},
		{12, false},                  // int
		{uint(12), false},            // uint
		{12.5, false},                // float
		{true, false},                // bool
		{"a", false},                 // string
		{[2]string{"1", "2"}, false}, // array
		{[]string{"1", "2"}, false},  // slice
		{&struct{}{}, false},         // pointer
		{struct{ int }{1}, false},    // struct
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { ProvideType(tc.giveModule) })
		} else {
			ProvideType(tc.giveModule)
			xtesting.Equal(t, _mc.provByType[reflect.TypeOf(tc.giveModule)], tc.giveModule)
		}
	}
}

func TestProvideImpl(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	for _, tc := range []struct {
		givePtr   interface{}
		giveImpl  interface{}
		wantPanic bool
	}{
		{nil, 0, true},
		{0, nil, true},
		{0, errors.New(""), true}, // non ptr
		{t, "", true},             // non itf
		{(*error)(nil), "", true}, // non impl
		{(*error)(nil), errors.New("test"), false},
		{(*fmt.Stringer)(nil), &strings.Builder{}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { ProvideImpl(tc.givePtr, tc.giveImpl) })
		} else {
			ProvideImpl(tc.givePtr, tc.giveImpl)
			xtesting.Equal(t, _mc.provByType[reflect.TypeOf(tc.givePtr).Elem()], tc.giveImpl)
		}
	}
}

func TestGetByName(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	ProvideName("int", 12)
	ProvideName("uint", uint(12))
	ProvideName("float", 12.5)
	ProvideName("bool", true)
	ProvideName("string", "a")
	ProvideName("array", [2]string{"1", "2"})
	ProvideName("slice", []string{"1", "2"})
	ProvideName("pointer", &struct{}{})
	ProvideName("struct", struct{ int }{1})

	for _, tc := range []struct {
		giveName   ModuleName
		wantModule interface{}
		wantPanic  bool
	}{
		{ModuleName(""), nil, true},
		{ModuleName("~"), nil, true},
		{ModuleName("-"), nil, true},
		{ModuleName("int"), 12, false},                    // int
		{ModuleName("uint"), uint(12), false},             // uint
		{ModuleName("float"), 12.5, false},                // float
		{ModuleName("bool"), true, false},                 // bool
		{ModuleName("string"), "a", false},                // string
		{ModuleName("array"), [2]string{"1", "2"}, false}, // array
		{ModuleName("slice"), []string{"1", "2"}, false},  // slice
		{ModuleName("pointer"), &struct{}{}, false},       // pointer
		{ModuleName("struct"), struct{ int }{1}, false},   // struct
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { GetByName(tc.giveName) })
		} else {
			module, ok := GetByName(tc.giveName)
			xtesting.Equal(t, module, tc.wantModule)
			xtesting.True(t, ok)
			xtesting.Equal(t, MustGetByName(tc.giveName), tc.wantModule)
		}
	}

	xtesting.Panic(t, func() { MustGetByName("not exist") })
}

func TestGetByType(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	ProvideType(12)
	ProvideType(uint(12))
	ProvideType(12.5)
	ProvideType(true)
	ProvideType("a")
	ProvideType([2]string{"1", "2"})
	ProvideType([]string{"1", "2"})
	ProvideType(&struct{}{})
	ProvideType(struct{ int }{1})

	for _, tc := range []struct {
		wantModule interface{}
		wantPanic  bool
	}{
		{nil, true},
		{nil, true},
		{nil, true},
		{12, false},                  // int
		{uint(12), false},            // uint
		{12.5, false},                // float
		{true, false},                // bool
		{"a", false},                 // string
		{[2]string{"1", "2"}, false}, // array
		{[]string{"1", "2"}, false},  // slice
		{&struct{}{}, false},         // pointer
		{struct{ int }{1}, false},    // struct
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { GetByType(tc.wantModule) })
		} else {
			module, ok := GetByType(tc.wantModule)
			xtesting.Equal(t, module, tc.wantModule)
			xtesting.True(t, ok)
			xtesting.Equal(t, MustGetByType(tc.wantModule), tc.wantModule)
		}
	}

	xtesting.Panic(t, func() { MustGetByType(struct{}{}) })
}

func TestGetByImpl(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))
	ProvideImpl((*error)(nil), errors.New("test"))
	ProvideImpl((*fmt.Stringer)(nil), &strings.Builder{})

	for _, tc := range []struct {
		givePtr    interface{}
		wantModule interface{}
		wantPanic  bool
	}{
		{nil, 0, true},
		{0, errors.New(""), true}, // non ptr
		{t, "", true},             // non itf
		{(*error)(nil), errors.New("test"), false},
		{(*fmt.Stringer)(nil), &strings.Builder{}, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { GetByImpl(tc.givePtr) })
		} else {
			module, ok := GetByImpl(tc.givePtr)
			xtesting.Equal(t, module, tc.wantModule)
			xtesting.True(t, ok)
			xtesting.Equal(t, MustGetByImpl(tc.givePtr), tc.wantModule)
		}
	}

	xtesting.Panic(t, func() { MustGetByImpl((*fmt.GoStringer)(nil)) })
}

func TestInject(t *testing.T) {
	SetLogger(DefaultLogger(LogSilent))

	// abnormal
	type testStruct1 struct {
		unexportedField string
		ExportedField1  string
		ExportedField2  string `module:""`
		ExportedField3  string `module:"-"`
	}
	test1 := &testStruct1{}
	dummy := 0
	for _, tc := range []struct {
		giveCtrl  interface{}
		wantAll   bool
		wantPanic bool
	}{
		{nil, true, true},
		{testStruct1{}, true, true},
		{&dummy, true, true},
		{test1, true, false},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { Inject(tc.giveCtrl) })
		} else {
			xtesting.Equal(t, Inject(tc.giveCtrl), tc.wantAll)
		}
	}
	for _, tc := range []struct {
		give interface{}
		want interface{}
	}{
		{test1.unexportedField, ""},
		{test1.ExportedField1, ""},
		{test1.ExportedField2, ""},
		{test1.ExportedField3, ""},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}

	// normal
	ProvideName("int", 12)
	ProvideName("uint", uint(12))
	ProvideName("float", 12.5)
	ProvideName("bool", true)
	ProvideName("string", "a")
	ProvideName("array", [2]string{"1", "2"})
	ProvideName("slice", []string{"1", "2"})
	ProvideName("pointer", &struct{}{})
	ProvideName("struct", struct{ int }{1})
	ProvideType(12)
	ProvideType(uint(12))
	ProvideType(12.5)
	ProvideType(true)
	ProvideType("a")
	ProvideType([2]string{"1", "2"})
	ProvideType([]string{"1", "2"})
	ProvideType(&struct{}{})
	ProvideType(struct{ int }{1})
	ProvideImpl((*error)(nil), errors.New("test"))
	ProvideImpl((*fmt.Stringer)(nil), &strings.Builder{})

	type testStruct2 struct {
		Int1     int           `module:"int"`
		Uint1    uint          `module:"uint"`
		Float1   float64       `module:"float"`
		Bool1    bool          `module:"bool"`
		String1  string        `module:"string"`
		Array1   [2]string     `module:"array"`
		Slice1   []string      `module:"slice"`
		Pointer1 *struct{}     `module:"pointer"`
		Struct1  struct{ int } `module:"struct"`
		Int2     int           `module:"~"`
		Uint2    uint          `module:"~"`
		Float2   float64       `module:"~"`
		Bool2    bool          `module:"~"`
		String2  string        `module:"~"`
		Array2   [2]string     `module:"~"`
		Slice2   []string      `module:"~"`
		Pointer2 *struct{}     `module:"~"`
		Struct2  struct{ int } `module:"~"`
		Err      error         `module:"~"`
		Sb       fmt.Stringer  `module:"~"`
	}

	test2 := &testStruct2{}
	all := Inject(test2)
	xtesting.True(t, all)
	xtesting.NotPanic(t, func() { MustInject(test2) })

	for _, tc := range []struct {
		give interface{}
		want interface{}
	}{
		{test2.Int1, 12},
		{test2.Uint1, uint(12)},
		{test2.Float1, 12.5},
		{test2.Bool1, true},
		{test2.String1, "a"},
		{test2.Array1, [2]string{"1", "2"}},
		{test2.Slice1, []string{"1", "2"}},
		{test2.Pointer1, &struct{}{}},
		{test2.Struct1, struct{ int }{1}},
		{test2.Int2, 12},
		{test2.Uint2, uint(12)},
		{test2.Float2, 12.5},
		{test2.Bool2, true},
		{test2.String2, "a"},
		{test2.Array2, [2]string{"1", "2"}},
		{test2.Slice2, []string{"1", "2"}},
		{test2.Pointer2, &struct{}{}},
		{test2.Struct2, struct{ int }{1}},
		{test2.Err, errors.New("test")},
		{test2.Sb, &strings.Builder{}},
	} {
		xtesting.Equal(t, tc.give, tc.want)
	}

	type testStruct3 struct {
		Struct2 struct{} `module:"struct2"`
		Struct3 struct{} `module:"~"`
	}
	test3 := &testStruct3{}
	xtesting.False(t, Inject(test3))
	xtesting.Panic(t, func() { MustInject(test3) })
}

func TestLogger(t *testing.T) {
	xtesting.EqualValue(t, LogSilent, 0) // 0000
	xtesting.EqualValue(t, LogName, 1)   // 0001
	xtesting.EqualValue(t, LogType, 2)   // 0010
	xtesting.EqualValue(t, LogImpl, 4)   // 0100
	xtesting.EqualValue(t, LogInject, 8) // 1000
	xtesting.EqualValue(t, LogAll, 15)   // 1111

	type testStruct struct {
		Int    int    `module:"int"`
		String string `module:"~"`
		Err    error  `module:"~"`
	}

	for _, tc := range []struct {
		str       string
		giveLevel LogLevel
		change    bool
	}{
		{"LogSilent", LogSilent, false},
		{"LogName", LogName, false},
		{"LogType", LogType, false},
		{"LogImpl", LogImpl, false},
		{"LogName | LogType", LogName | LogType, false},
		{"LogType | LogImpl", LogType | LogImpl, false},
		{"LogName | LogImpl", LogName | LogImpl, false},
		{"LogInject", LogInject, false},
		{"LogAll", LogAll, false},
		{"LogAll 2", LogAll, true},
	} {
		fmt.Println(tc.str)
		SetLogger(DefaultLogger(tc.giveLevel))
		if tc.change {
			LogLeftArrow = func(arg1, arg2, arg3 string) {
				fmt.Printf("[XMODULE] %-4s %-15s <-- %s\n", arg1, arg2, arg3)
			}
			LogRightArrow = func(arg1, arg2, arg3 string) {
				fmt.Printf("[XMODULE] %-4s %-15s --> %s\n", arg1, arg2, arg3)
			}
		}

		ProvideName("int", 0)
		ProvideType("test")
		ProvideImpl((*error)(nil), errors.New("test"))
		GetByName("int")
		GetByType("")
		GetByImpl((*error)(nil))
		Inject(&testStruct{})
	}
}
