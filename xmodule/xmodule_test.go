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
	xtesting.Equal(t, ModuleName("-").String(), "-")
	xtesting.Equal(t, ModuleName("~").String(), "~")
	xtesting.Equal(t, ModuleName("test").String(), "test")
}

func TestProvideName(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
	for _, tc := range []struct {
		giveName   ModuleName
		giveModule interface{}
		wantPanicP bool
		wantPanicR bool
	}{
		{ModuleName(""), 0, true, true},
		{ModuleName("-"), 0, true, true},
		{ModuleName("~"), 0, true, true},
		{ModuleName("0"), nil, true, false},
		{ModuleName("int"), 12, false, false},                    // int
		{ModuleName("uint"), uint(12), false, false},             // uint
		{ModuleName("float"), 12.5, false, false},                // float
		{ModuleName("bool"), true, false, false},                 // bool
		{ModuleName("string"), "a", false, false},                // string
		{ModuleName("array"), [2]string{"1", "2"}, false, false}, // array
		{ModuleName("slice"), []string{"1", "2"}, false, false},  // slice
		{ModuleName("pointer"), &struct{}{}, false, false},       // pointer
		{ModuleName("struct"), struct{ int }{1}, false, false},   // struct
	} {
		if tc.wantPanicP {
			xtesting.Panic(t, func() { ProvideName(tc.giveName, tc.giveModule) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByName(tc.giveName) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideName(tc.giveName, tc.giveModule)
			xtesting.Equal(t, _mc.byName[tc.giveName], tc.giveModule)
			RemoveByName(tc.giveName)
			_, ok := _mc.byName[tc.giveName]
			xtesting.False(t, ok)
		}
	}
}

func TestProvideType(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
	for _, tc := range []struct {
		giveModule interface{}
		wantPanicP bool
		wantPanicR bool
	}{
		{nil, true, true},
		{12, false, false},                  // int
		{uint(12), false, false},            // uint
		{12.5, false, false},                // float
		{true, false, false},                // bool
		{"a", false, false},                 // string
		{[2]string{"1", "2"}, false, false}, // array
		{[]string{"1", "2"}, false, false},  // slice
		{&struct{}{}, false, false},         // pointer
		{struct{ int }{1}, false, false},    // struct
	} {
		if tc.wantPanicP {
			xtesting.Panic(t, func() { ProvideType(tc.giveModule) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByType(tc.giveModule) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideType(tc.giveModule)
			xtesting.Equal(t, _mc.byType[reflect.TypeOf(tc.giveModule)], tc.giveModule)
			RemoveByType(tc.giveModule)
			_, ok := _mc.byType[reflect.TypeOf(tc.giveModule)]
			xtesting.False(t, ok)
		}
	}
}

func TestProvideImpl(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
	for _, tc := range []struct {
		givePtr    interface{}
		giveImpl   interface{}
		wantPanicP bool
		wantPanicR bool
	}{
		{nil, 0, true, false},
		{0, nil, true, true},
		{0, errors.New(""), true, true},  // non ptr
		{t, "", true, true},              // non itf
		{(*error)(nil), "", true, false}, // non impl
		{(*error)(nil), errors.New("test"), false, false},
		{(*fmt.Stringer)(nil), &strings.Builder{}, false, false},
	} {
		if tc.wantPanicP {
			xtesting.Panic(t, func() { ProvideImpl(tc.givePtr, tc.giveImpl) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByImpl(tc.givePtr) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideImpl(tc.givePtr, tc.giveImpl)
			xtesting.Equal(t, _mc.byType[reflect.TypeOf(tc.givePtr).Elem()], tc.giveImpl)
			RemoveByImpl(tc.givePtr)
			_, ok := _mc.byType[reflect.TypeOf(tc.givePtr).Elem()]
			xtesting.False(t, ok)
		}
	}
}

func TestGetByName(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
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

	xtesting.Panic(t, func() { MustGetByName("") })
	xtesting.Panic(t, func() { MustGetByName("~") })
	xtesting.Panic(t, func() { MustGetByName("-") })
	xtesting.Panic(t, func() { MustGetByName("not exist") })
}

func TestGetByType(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
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
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
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
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))

	t.Run("abnormal", func(t *testing.T) {
		type allIgnored struct {
			unexportedField string
			ExportedField1  string
			ExportedField2  string `module:""`
			ExportedField3  string `module:"-"`
		}
		intValue := 0
		test1 := &allIgnored{}

		for _, tc := range []struct {
			giveCtrl  interface{}
			wantAll   bool
			wantPanic bool
		}{
			{nil, true, true},        // nil
			{struct{}{}, true, true}, // struct
			{&intValue, true, true},  // *int
			{test1, true, false},     // *struct
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
	})

	type testStruct struct {
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

	t.Run("normal", func(t *testing.T) {
		test2 := &testStruct{}
		xtesting.True(t, Inject(test2))
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
	})

	t.Run("not_all", func(t *testing.T) {
		// cannot assign
		ProvideName("uint", 12)
		test := &testStruct{}
		xtesting.False(t, Inject(test))
		xtesting.Panic(t, func() { MustInject(test) })

		// module not found
		ProvideName("uint", uint(12))
		_mc.byName = map[ModuleName]interface{}{}
		test = &testStruct{}
		xtesting.False(t, Inject(test))
		xtesting.Panic(t, func() { MustInject(test) })
	})
}

func TestLogger(t *testing.T) {
	xtesting.EqualValue(t, LogPrvName, 1)    // 00001
	xtesting.EqualValue(t, LogPrvType, 2)    // 00010
	xtesting.EqualValue(t, LogPrvImpl, 4)    // 00100
	xtesting.EqualValue(t, LogInjField, 8)   // 01000
	xtesting.EqualValue(t, LogInjFinish, 16) // 10000
	xtesting.EqualValue(t, LogSilent, 0)     // 00000
	xtesting.EqualValue(t, LogAll, 31)       // 11111

	type testStruct struct {
		unexported bool
		WithoutTag bool
		EmptyTag   bool `module:""`
		IgnoreTag  bool `module:"-"`

		Int    int         `module:"int"`
		Uint   uint        `module:"uint"`
		Float  float64     `module:"~"`
		String string      `module:"~"`
		Itf    interface{} `module:"~"`
		Error  error       `module:"~"`
	}

	for _, tc := range []struct {
		name         string
		giveLevel    LogLevel
		giveIgnore   bool
		giveMismatch bool
		giveCustom   bool
	}{
		{"LogSilent", LogSilent, false, false, false},
		{"LogPrvName", LogPrvName, false, false, false},
		{"LogPrvType", LogPrvType, false, false, false},
		{"LogPrvImpl", LogPrvImpl, false, false, false},
		{"LogPrvName | LogPrvType", LogPrvName | LogPrvType, false, false, false},
		{"LogPrvType | LogPrvImpl", LogPrvType | LogPrvImpl, false, false, false},
		{"LogPrvName | LogPrvImpl", LogPrvImpl | LogPrvName, false, false, false},
		{"LogInjField", LogInjField, false, false, false},
		{"LogInjFinish with module not found", LogInjFinish, true, false, false},
		{"LogInjFinish with cannot assign", LogInjFinish, false, true, false},
		{"LogInjField | LogInjFinish", LogInjField | LogInjFinish, false, false, false},
		{"LogAll", LogAll, false, false, false},
		{"LogAll with custom function", LogAll, false, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mc := NewModuleContainer()
			if !tc.giveCustom {
				mc.SetLogger(DefaultLogger(tc.giveLevel, nil, nil))
			} else {
				mc.SetLogger(DefaultLogger(tc.giveLevel, func(moduleName, moduleType string) {
					fmt.Printf("[Xmodule] Prv: %s <-- %s\n", moduleName, moduleType)
				}, func(moduleName string, structName string, additional string) {
					fmt.Printf("[Xmodule] Inj: %s --> %s %s\n", moduleName, structName, additional)
				}))
			}

			// prv
			if !tc.giveMismatch {
				mc.ProvideName("int", 1)
			} else {
				mc.ProvideName("int", "1")
			}
			if !tc.giveIgnore {
				mc.ProvideName("uint", uint(1))
				mc.ProvideType(1.0)
			}
			mc.ProvideType("test")
			mc.ProvideImpl((*interface{})(nil), struct{}{})
			mc.ProvideImpl((*error)(nil), errors.New("test"))

			// inj
			_ = mc.Inject(&testStruct{})
		})
	}
}
