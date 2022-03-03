package xmodule

import (
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSimpleTypes(t *testing.T) {
	xtesting.Equal(t, ModuleName("").String(), "")
	xtesting.Equal(t, ModuleName("-").String(), "-")
	xtesting.Equal(t, ModuleName("~").String(), "~")
	xtesting.Equal(t, ModuleName("test").String(), "test")

	xtesting.Equal(t, nameKey("").String(), "<invalid>")
	xtesting.Equal(t, nameKey("-").String(), "name:-")
	xtesting.Equal(t, nameKey("~").String(), "name:~")
	xtesting.Equal(t, nameKey("name").String(), "name:name")

	xtesting.Equal(t, typeKey(nil).String(), "<invalid>")
	xtesting.Equal(t, typeKey(reflect.TypeOf("s")).String(), "type:string")
	xtesting.Equal(t, typeKey(reflect.TypeOf(&strconv.NumError{})).String(), "type:*strconv.NumError")
	xtesting.Equal(t, typeKey(reflect.TypeOf(new(fmt.Stringer)).Elem()).String(), "type:fmt.Stringer")
}

func TestProvideByName(t *testing.T) {
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
			xtesting.Panic(t, func() { ProvideByName(tc.giveName, tc.giveModule) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByName(tc.giveName) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideByName(tc.giveName, tc.giveModule)
			xtesting.Equal(t, _mc.modules[nameKey(tc.giveName)], tc.giveModule)
			xtesting.True(t, RemoveByName(tc.giveName))
			_, ok := _mc.modules[nameKey(tc.giveName)]
			xtesting.False(t, ok)
			xtesting.False(t, RemoveByName(tc.giveName))
		}
	}
}

func TestProvideByType(t *testing.T) {
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
			xtesting.Panic(t, func() { ProvideByType(tc.giveModule) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByType(tc.giveModule) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideByType(tc.giveModule)
			xtesting.Equal(t, _mc.modules[typeKey(reflect.TypeOf(tc.giveModule))], tc.giveModule)
			xtesting.True(t, RemoveByType(tc.giveModule))
			_, ok := _mc.modules[typeKey(reflect.TypeOf(tc.giveModule))]
			xtesting.False(t, ok)
			xtesting.False(t, RemoveByType(tc.giveModule))
		}
	}
}

func TestProvideByIntf(t *testing.T) {
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
			xtesting.Panic(t, func() { ProvideByIntf(tc.givePtr, tc.giveImpl) })
		}
		if tc.wantPanicR {
			xtesting.Panic(t, func() { RemoveByIntf(tc.givePtr) })
		}
		if !tc.wantPanicP && !tc.wantPanicR {
			ProvideByIntf(tc.givePtr, tc.giveImpl)
			xtesting.Equal(t, _mc.modules[typeKey(reflect.TypeOf(tc.givePtr).Elem())], tc.giveImpl)
			xtesting.True(t, RemoveByIntf(tc.givePtr))
			_, ok := _mc.modules[typeKey(reflect.TypeOf(tc.givePtr).Elem())]
			xtesting.False(t, ok)
			xtesting.False(t, RemoveByIntf(tc.givePtr))
		}
	}
}

func TestGetByName(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
	ProvideByName("int", 12)
	ProvideByName("uint", uint(12))
	ProvideByName("float", 12.5)
	ProvideByName("bool", true)
	ProvideByName("string", "a")
	ProvideByName("array", [2]string{"1", "2"})
	ProvideByName("slice", []string{"1", "2"})
	ProvideByName("pointer", &struct{}{})
	ProvideByName("struct", struct{ int }{1})

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
	ProvideByType(12)
	ProvideByType(uint(12))
	ProvideByType(12.5)
	ProvideByType(true)
	ProvideByType("a")
	ProvideByType([2]string{"1", "2"})
	ProvideByType([]string{"1", "2"})
	ProvideByType(&struct{}{})
	ProvideByType(struct{ int }{1})

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

func TestGetByIntf(t *testing.T) {
	defer func() { _mc = NewModuleContainer() }()
	SetLogger(DefaultLogger(LogSilent, nil, nil))
	ProvideByIntf((*error)(nil), errors.New("test"))
	ProvideByIntf((*fmt.Stringer)(nil), &strings.Builder{})

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
			xtesting.Panic(t, func() { GetByIntf(tc.givePtr) })
		} else {
			module, ok := GetByIntf(tc.givePtr)
			xtesting.Equal(t, module, tc.wantModule)
			xtesting.True(t, ok)
			xtesting.Equal(t, MustGetByIntf(tc.givePtr), tc.wantModule)
		}
	}

	xtesting.Panic(t, func() { MustGetByIntf((*fmt.GoStringer)(nil)) })
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
		test1 := &allIgnored{}

		for _, tc := range []struct {
			giveCtrl  interface{}
			wantErr   bool
			wantPanic bool
		}{
			{nil, false, true},        // nil
			{struct{}{}, false, true}, // struct
			{new(int), true, true},    // *int
			{test1, false, false},     // *struct
		} {
			if tc.wantPanic {
				xtesting.Panic(t, func() { _ = Inject(tc.giveCtrl) })
			} else {
				xtesting.Equal(t, Inject(tc.giveCtrl) != nil, tc.wantErr)
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
	ProvideByName("int", 12)
	ProvideByName("uint", uint(12))
	ProvideByName("float", 12.5)
	ProvideByName("bool", true)
	ProvideByName("string", "a")
	ProvideByName("array", [2]string{"1", "2"})
	ProvideByName("slice", []string{"1", "2"})
	ProvideByName("pointer", &struct{}{})
	ProvideByName("struct", struct{ int }{1})
	ProvideByType(12)
	ProvideByType(uint(12))
	ProvideByType(12.5)
	ProvideByType(true)
	ProvideByType("a")
	ProvideByType([2]string{"1", "2"})
	ProvideByType([]string{"1", "2"})
	ProvideByType(&struct{}{})
	ProvideByType(struct{ int }{1})
	ProvideByIntf((*error)(nil), errors.New("test"))
	ProvideByIntf((*fmt.Stringer)(nil), &strings.Builder{})

	t.Run("normal", func(t *testing.T) {
		test2 := &testStruct{}
		xtesting.Nil(t, Inject(test2))
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

	t.Run("some errors", func(t *testing.T) {
		// type mismatch
		ProvideByName("uint", 12)
		test := &testStruct{}
		xtesting.NotNil(t, Inject(test))
		xtesting.Panic(t, func() { MustInject(test) })

		// module not found
		ProvideByName("uint", uint(12))
		_mc.modules = map[mkey]interface{}{}
		test = &testStruct{}
		xtesting.NotNil(t, Inject(test))
		xtesting.Panic(t, func() { MustInject(test) })
	})
}

func TestAutoProvide(t *testing.T) {
	t.Run("providers", func(t *testing.T) {
		xtesting.Panic(t, func() { NameProvider("", 0) })
		xtesting.Panic(t, func() { NameProvider("~", struct{}{}) })
		xtesting.Panic(t, func() { NameProvider("xxx", nil) })
		xtesting.NotPanic(t, func() { NameProvider("xxx", "module") })
		xtesting.Panic(t, func() { TypeProvider(nil) })
		xtesting.NotPanic(t, func() { TypeProvider("module") })
		xtesting.Panic(t, func() { IntfProvider(0, struct{}{}) })
		xtesting.Panic(t, func() { IntfProvider(fmt.Stringer(nil), ModuleName("x")) })
		xtesting.NotPanic(t, func() { IntfProvider((*fmt.Stringer)(nil), ModuleName("x")) })

		_mc = NewModuleContainer()
		_mc.SetLogger(DefaultLogger(LogSilent, nil, nil))
		xtesting.Panic(t, func() { _ = _mc.AutoProvide(&ModuleProvider{}) })
		xtesting.NotPanic(t, func() { _ = _mc.AutoProvide(nil, NameProvider("xxx", 0)) })
		xtesting.NotPanic(t, func() { _mc.MustAutoProvide(nil, nil, nil) })
	})

	t.Run("simple", func(t *testing.T) {
		defer func() { _mc = NewModuleContainer() }()
		SetLogger(DefaultLogger(LogSilent, nil, nil))

		repository := []int{1, 2, 3, 4}
		type DBer interface {
			GetData() interface{}
		}
		type DB struct {
			DBer // embed directly
			Flag int8
		}
		type Service struct {
			Repository []int `module:"repo"`
			DB         DBer  `module:"~"`
			Variable   string
		}
		type Controller struct {
			Service *Service `module:"~"`
		}

		providers := []*ModuleProvider{
			TypeProvider(&Service{Variable: "hello world"}),
			NameProvider("repo", repository),
			IntfProvider((*DBer)(nil), &DB{Flag: 3}),
		}
		xtesting.Nil(t, AutoProvide(providers...))
		ctrl := &Controller{Service: MustGetByType(&Service{}).(*Service)}
		xtesting.Equal(t, ctrl.Service.Variable, "hello world")
		xtesting.Equal(t, ctrl.Service.Repository, []int{1, 2, 3, 4})
		xtesting.Equal(t, ctrl.Service.DB.(*DB).Flag, int8(3))

		_mc.modules = map[mkey]interface{}{}
		xtesting.NotPanic(t, func() { _mc.MustAutoProvide(providers...) })
		ctrl = &Controller{Service: MustGetByType(&Service{}).(*Service)}
		xtesting.Equal(t, ctrl.Service.DB.(*DB).Flag, int8(3))
	})

	t.Run("complex", func(t *testing.T) {
		defer func() { _mc = NewModuleContainer() }()
		SetLogger(DefaultLogger(LogSilent, nil, nil))

		type (
			IE interface{ EEE() }
			ID interface{ DDD() }
			F  struct{}
			E  struct {
				IE
				F *F     `module:"fff"`
				X string `module:"str1"`
				Y string `module:"str2"`
				Z string `module:"~"`
			}
			C struct {
				E IE     `module:"~"`
				F *F     `module:"fff"`
				X string `module:"str2"`
				Y string `module:"~"`
				Z uint64 `module:"~"`
			}
			D struct {
				ID
				C *C    `module:"~"`
				X int32 `module:"int1"`
				Y int64 `module:"int2"`
				Z string
			}
			B struct {
				C *C    `module:"~"`
				D ID    `module:"~"`
				X int64 `module:"int2"`
				Y int8
			}
			A struct {
				B *B     `module:"bbb"`
				C *C     `module:"~"`
				D *D     `module:"ddd"`
				E IE     `module:"~"`
				X string `module:"str1"`
				Y string `module:"str2"`
				Z int32  `module:"int1"`
				W int64  `module:"int2"`
			}
			O struct {
				A *A `module:"~"`
				F *F `module:"~"`
				X string
				Y int64 `module:"int2"`
			}
		)
		providers := []*ModuleProvider{
			TypeProvider(&O{X: "xxx"}),
			TypeProvider(&A{}),
			NameProvider("bbb", &B{Y: 127}),
			TypeProvider(&C{}),
			IntfProvider((*ID)(nil), &D{Z: "zzz"}), NameProvider("ddd", &D{Z: "zzz2"}),
			IntfProvider((*IE)(nil), &E{}),
			NameProvider("fff", &F{}), TypeProvider(&F{}),
			TypeProvider("abc"), TypeProvider(uint64(789)),
			NameProvider("int1", int32(111)), NameProvider("int2", int64(222)),
			NameProvider("str1", "sss"), NameProvider("str2", "ttt"),
		}
		_mc = NewModuleContainer()
		_mc.SetLogger(DefaultLogger(LogPrvName|LogPrvType|LogPrvIntf|LogInjFinish, nil, nil))
		xtesting.Nil(t, AutoProvide(providers...))
		fmt.Println("==============")
		_mc = NewModuleContainer()
		_mc.SetLogger(DefaultLogger(LogInjField|LogInjFinish, nil, nil))
		xtesting.NotPanic(t, func() { MustAutoProvide(providers...) })

		o := _mc.MustGetByType(&O{}).(*O)
		xtesting.Equal(t, o.X, "xxx")
		xtesting.Equal(t, o.Y, int64(222))
		xtesting.Equal(t, *o.F, F{})
		xtesting.Equal(t, o.A.X, "sss")
		xtesting.Equal(t, o.A.Y, "ttt")
		xtesting.Equal(t, o.A.Z, int32(111))
		xtesting.Equal(t, o.A.W, int64(222))
		xtesting.Equal(t, o.A.B.X, int64(222))
		xtesting.Equal(t, o.A.B.Y, int8(127))
		xtesting.Equal(t, o.A.B.C, o.A.C)
		xtesting.Equal(t, o.A.B.D.(*D).Y, int64(222))
		xtesting.Equal(t, o.A.B.D.(*D).Z, "zzz")
		xtesting.Equal(t, o.A.C.X, "ttt")
		xtesting.Equal(t, o.A.C.Y, "abc")
		xtesting.Equal(t, o.A.C.Z, uint64(789))
		xtesting.Equal(t, o.A.C.E, o.A.E)
		xtesting.Equal(t, *o.A.C.F, F{})
		xtesting.Equal(t, o.A.D.X, int32(111))
		xtesting.Equal(t, o.A.D.Y, int64(222))
		xtesting.Equal(t, o.A.D.Z, "zzz2")
		xtesting.Equal(t, o.A.D.C, o.A.C)
		xtesting.Equal(t, o.A.E.(*E).X, "sss")
		xtesting.Equal(t, o.A.E.(*E).Y, "ttt")
		xtesting.Equal(t, o.A.E.(*E).Z, "abc")
		xtesting.Equal(t, *o.A.E.(*E).F, F{})
	})

	t.Run("errors", func(t *testing.T) {
		defer func() { _mc = NewModuleContainer() }()
		SetLogger(DefaultLogger(LogSilent, nil, nil))
		restore := func() { _mc.modules = map[mkey]interface{}{} }

		type Dep1 struct {
			X string `module:"x"`
			Y string `module:"y"`
		}
		type Dep2 struct {
			D *Dep2 `module:"~"`
		}
		type Dep3 struct {
			D interface{} `module:"dep4"`
		}
		type Dep4 struct {
			D *Dep3 `module:"~"`
		}

		restore()
		xtesting.NotNil(t, AutoProvide(TypeProvider(&Dep1{}))) // module not found
		restore()
		xtesting.NotNil(t, AutoProvide(NameProvider("x", &Dep2{}))) // module not found
		restore()
		xtesting.NotNil(t, AutoProvide(TypeProvider(&Dep1{}), NameProvider("x", 0))) // type mismatch
		restore()
		xtesting.NotNil(t, AutoProvide(TypeProvider(&Dep1{}), NameProvider("x", 0), NameProvider("y", false))) // type mismatch, multi error
		restore()
		xtesting.NotNil(t, AutoProvide(TypeProvider(&Dep2{}))) // dependency cycle
		restore()
		xtesting.NotNil(t, AutoProvide(TypeProvider(&Dep3{}), NameProvider("dep4", &Dep4{}))) // dependency cycle

		restore()
		xtesting.Panic(t, func() { MustAutoProvide(TypeProvider(&Dep1{})) })
		restore()
		xtesting.Panic(t, func() { MustAutoProvide(NameProvider("x", &Dep2{})) })
		restore()
		xtesting.Panic(t, func() { MustAutoProvide(TypeProvider(&Dep1{}), NameProvider("x", 0)) })
		restore()
		xtesting.Panic(t, func() { MustAutoProvide(TypeProvider(&Dep1{}), NameProvider("x", 0), NameProvider("y", false)) })
		restore()
		xtesting.Panic(t, func() { MustAutoProvide(TypeProvider(&Dep2{})) })
		restore()
		xtesting.Panic(t, func() { MustAutoProvide(TypeProvider(&Dep3{}), NameProvider("dep4", &Dep4{})) })
	})
}

func TestLogger(t *testing.T) {
	xtesting.EqualValue(t, LogPrvName, 1)    // 00001
	xtesting.EqualValue(t, LogPrvType, 2)    // 00010
	xtesting.EqualValue(t, LogPrvIntf, 4)    // 00100
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
		{"LogPrvIntf", LogPrvIntf, false, false, false},
		{"LogPrvName | LogPrvType", LogPrvName | LogPrvType, false, false, false},
		{"LogPrvType | LogPrvIntf", LogPrvType | LogPrvIntf, false, false, false},
		{"LogPrvName | LogPrvIntf", LogPrvIntf | LogPrvName, false, false, false},
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
				mc.ProvideByName("int", 1)
			} else {
				mc.ProvideByName("int", "1")
			}
			if !tc.giveIgnore {
				mc.ProvideByName("uint", uint(1))
				mc.ProvideByType(1.0)
			}
			mc.ProvideByType("test")
			mc.ProvideByIntf((*interface{})(nil), struct{}{})
			mc.ProvideByIntf((*error)(nil), errors.New("test"))

			// inj
			_ = mc.Inject(&testStruct{})
		})
	}
}
