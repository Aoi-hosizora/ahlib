package xdi

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
)

func TestByName(t *testing.T) {
	sn := ServiceName("")
	xtesting.Equal(t, sn.String(), "")
	sn = "test"
	xtesting.Equal(t, sn.String(), "test")

	_, ok := GetByName("a")
	xtesting.False(t, ok)
	xtesting.Panic(t, func() { GetByNameForce("a") })

	ProvideName("a", "a")
	a, ok := GetByName("a")
	xtesting.Equal(t, a, "a")
	xtesting.True(t, ok)
	a = GetByNameForce("a")
	xtesting.Equal(t, a, "a")

	ProvideName("a", 5)
	a, ok = GetByName("a")
	xtesting.Equal(t, a, 5)
	xtesting.True(t, ok)
	a = GetByNameForce("a")
	xtesting.Equal(t, a, 5)

	xtesting.Panic(t, func() { ProvideName("a", nil) })
	xtesting.Panic(t, func() { ProvideName("", 0) })
	xtesting.Panic(t, func() { ProvideName("-", 0) })
	xtesting.Panic(t, func() { ProvideName("~", 0) })
}

func TestByType(t *testing.T) {
	_, ok := GetByType("")
	xtesting.False(t, ok)
	xtesting.Panic(t, func() { GetByTypeForce("") })

	ProvideType("a")
	a, ok := GetByType("")
	xtesting.Equal(t, a, "a")
	xtesting.True(t, ok)
	a = GetByTypeForce("")
	xtesting.Equal(t, a, "a")

	ProvideType("aa")
	a, ok = GetByType("")
	xtesting.Equal(t, a, "aa")
	xtesting.True(t, ok)
	a = GetByTypeForce("")
	xtesting.Equal(t, a, "aa")

	xtesting.Panic(t, func() { ProvideType(nil) })
	xtesting.Panic(t, func() { GetByType(nil) })
}

type testByImplInterface1 interface{ A() }

type testByImplStruct1 struct{ I int }

func (t *testByImplStruct1) A() {}

type testByImplStruct2 struct{}

func TestByImpl(t *testing.T) {
	i := (*testByImplInterface1)(nil)
	s := &testByImplStruct1{I: 0}

	_, ok := GetByImpl(i)
	xtesting.False(t, ok)
	xtesting.Panic(t, func() { GetByImplForce(i) })

	ProvideImpl(i, s)
	a, ok := GetByImpl(i)
	xtesting.Equal(t, a, testByImplInterface1(s))
	xtesting.True(t, ok)
	aa, ok := a.(*testByImplStruct1)
	xtesting.Equal(t, aa, s)
	xtesting.True(t, ok)
	a = GetByImplForce(i)
	xtesting.Equal(t, a, s)

	s = &testByImplStruct1{I: 5}

	ProvideImpl(i, s)
	a, ok = GetByImpl(i)
	xtesting.Equal(t, a, testByImplInterface1(s))
	xtesting.True(t, ok)
	aa, ok = a.(*testByImplStruct1)
	xtesting.Equal(t, aa, s)
	xtesting.True(t, ok)
	a = GetByImplForce(i)
	xtesting.Equal(t, a, s)

	n := 0
	ptr := &n
	xtesting.Panic(t, func() { ProvideImpl(nil, 0) })
	xtesting.Panic(t, func() { ProvideImpl("0", 0) })
	xtesting.Panic(t, func() { ProvideImpl(0, 0) })
	xtesting.Panic(t, func() { ProvideImpl(ptr, 0) })
	xtesting.Panic(t, func() { ProvideImpl(i, nil) })
	xtesting.Panic(t, func() { ProvideImpl(i, &testByImplStruct2{}) })
	xtesting.Panic(t, func() { GetByImpl(nil) })
	xtesting.Panic(t, func() { GetByImpl("0") })
	xtesting.Panic(t, func() { GetByImpl(0) })
	xtesting.Panic(t, func() { GetByImpl(ptr) })
}

func TestInject(t *testing.T) {
	xtesting.Panic(t, func() { Inject(nil) })
	xtesting.Panic(t, func() { Inject("") })

	type A struct {
		// type
		I   int     `di:"~"`
		I8  int8    `di:"~"`
		I16 int16   `di:"~"`
		I32 int32   `di:"~"`
		I64 int64   `di:"~"`
		U   uint    `di:"~"`
		U8  uint8   `di:"~"`
		U16 uint16  `di:"~"`
		U32 uint32  `di:"~"`
		U64 uint64  `di:"~"`
		F32 float32 `di:"~"`
		F64 float64 `di:"~"`

		// name
		B  bool              `di:"b"`
		S  string            `di:"s"`
		BS []byte            `di:"bs"`
		IS []int             `di:"is"`
		SS []string          `di:"ss"`
		FA [3]float64        `di:"fa"`
		BA [2]bool           `di:"ba"`
		M  map[string]string `di:"m"`

		Useless1 int         `di:""`
		Useless2 chan func() `di:"-"`
	}
	a := &A{}

	ProvideType(1)
	ProvideType(int8(1))
	ProvideType(int16(1))
	ProvideType(int32(1))
	ProvideType(int64(1))
	ProvideType(uint(1))
	ProvideType(uint8(1))
	ProvideType(uint16(1))
	ProvideType(uint32(1))
	ProvideType(uint64(1))
	ProvideType(float32(0.5))
	ProvideType(0.5)

	all := Inject(a)
	xtesting.False(t, all)
	xtesting.Panic(t, func() { MustInject(a) })

	ProvideName("b", true)
	ProvideName("s", "sss")
	ProvideName("bs", []byte("sss"))
	ProvideName("is", []int{1, 2, 3})
	ProvideName("ss", []string{"1", "2", "3"})
	ProvideName("fa", [3]float64{0, 1.5, 0.5})
	ProvideName("ba", [2]bool{true, false})
	ProvideName("m", map[string]string{"a": "aa", "b": "bb"})

	xtesting.True(t, Inject(a))
	xtesting.NotPanic(t, func() { MustInject(a) })

	xtesting.Equal(t, a.I, 1)
	xtesting.Equal(t, a.I8, int8(1))
	xtesting.Equal(t, a.I16, int16(1))
	xtesting.Equal(t, a.I32, int32(1))
	xtesting.Equal(t, a.I64, int64(1))
	xtesting.Equal(t, a.U, uint(1))
	xtesting.Equal(t, a.U8, uint8(1))
	xtesting.Equal(t, a.U16, uint16(1))
	xtesting.Equal(t, a.U32, uint32(1))
	xtesting.Equal(t, a.U64, uint64(1))
	xtesting.Equal(t, a.F32, float32(0.5))
	xtesting.Equal(t, a.F64, 0.5)
	xtesting.Equal(t, a.B, true)
	xtesting.Equal(t, a.S, "sss")
	xtesting.Equal(t, a.BS, []byte("sss"))
	xtesting.Equal(t, a.IS, []int{1, 2, 3})
	xtesting.Equal(t, a.SS, []string{"1", "2", "3"})
	xtesting.Equal(t, a.FA, [...]float64{0, 1.5, 0.5})
	xtesting.Equal(t, a.BA, [2]bool{true, false})
	xtesting.Equal(t, a.M, map[string]string{"a": "aa", "b": "bb"})
}

func TestLogger(t *testing.T) {
	xtesting.EqualValue(t, LogSilent, 0) // 0000
	xtesting.EqualValue(t, LogName, 1)   // 0001
	xtesting.EqualValue(t, LogType, 2)   // 0010
	xtesting.EqualValue(t, LogImpl, 4)   // 0100
	xtesting.EqualValue(t, LogInject, 8) // 1000
	xtesting.EqualValue(t, LogAll, 15)   // 1111

	type a struct {
		A string `di:"a"`
	}

	log.Println("LogAll")
	SetLogger(DefaultLogger(LogAll))
	ProvideName("a", "a")
	ProvideType("a")
	ProvideImpl((*testByImplInterface1)(nil), &testByImplStruct1{})
	GetByName("a")
	GetByType("")
	GetByImpl((*testByImplInterface1)(nil))
	Inject(&a{})

	log.Println("LogSilent")
	SetLogger(DefaultLogger(LogSilent))
	ProvideName("a", "a")
	ProvideType("a")
	ProvideImpl((*testByImplInterface1)(nil), &testByImplStruct1{})
	GetByName("a")
	GetByType("")
	GetByImpl((*testByImplInterface1)(nil))
	Inject(&a{})

	log.Println("LogName | LogImpl")
	SetLogger(DefaultLogger(LogName | LogImpl))
	ProvideName("a", "a")
	ProvideType("a")
	ProvideImpl((*testByImplInterface1)(nil), &testByImplStruct1{})
	GetByName("a")
	GetByType("")
	GetByImpl((*testByImplInterface1)(nil))
	Inject(&a{})
}
