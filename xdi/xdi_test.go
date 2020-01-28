package xdi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type IServiceA interface {
	A() int
}

type IServiceB interface {
	B(param string) string
	C(param int) int
}

type IServiceC interface{}

type ServiceA struct{}

type ServiceB struct {
	SA IServiceA `di:"~"` // auto inject service by type of interface
}

type ServiceC struct{}

func (ServiceA) A() int {
	return 2
}

func (ServiceB) B(param string) string {
	return param + "123"
}

func (b ServiceB) C(param int) int {
	return param * b.SA.A()
}

type Controller struct {
	SA  *ServiceA `di:"a"` // inject service by name
	SB  IServiceB `di:"~"` // auto inject service by type of interface
	SSB *ServiceB `di:"~"` // auto inject service by type of struct
	SC  IServiceC `di:"-"` // not inject
	PD  int       `di:"d"` // inject data by name
}

func NewServiceA(dic *DiContainer) *ServiceA {
	a := &ServiceA{}
	dic.Inject(a)
	return &ServiceA{}
}

func NewServiceB(dic *DiContainer) *ServiceB {
	b := &ServiceB{}
	dic.Inject(b)
	return b
}

func NewServiceC(dic *DiContainer) *ServiceC {
	c := &ServiceC{}
	dic.Inject(c)
	return c
}

func Test_DiContainer_Inject(t *testing.T) {
	dic := NewDiContainer()

	dic.ProvideByName("a", NewServiceA(dic))
	dic.ProvideImpl((*IServiceA)(nil), *NewServiceA(dic))
	dic.Provide(NewServiceB(dic))
	dic.ProvideImpl((*IServiceB)(nil), NewServiceB(dic))
	dic.ProvideImpl((*IServiceC)(nil), NewServiceC(dic))
	dic.ProvideByName("d", 123)

	ctrl := &Controller{}
	ok := dic.Inject(ctrl)

	assert.Equal(t, ctrl.SA.A(), 2)
	assert.Equal(t, ctrl.SB.B("a"), "a123")
	assert.Equal(t, ctrl.SB.C(2), 4)
	assert.Equal(t, ctrl.SSB.B("a"), "a123")
	assert.Equal(t, ctrl.SSB.C(2), 4)
	assert.Equal(t, ctrl.SC == nil, true)
	assert.Equal(t, ctrl.PD, 123)

	ctrl2 := &Controller{}
	ctrl3 := &struct{ Other int `di:"o"` }{}

	assert.Equal(t, ok, true)
	assert.Equal(t, AllInjected(ctrl), true)
	assert.Equal(t, AllInjected(ctrl2), false)
	assert.Equal(t, dic.Inject(ctrl2), true)
	assert.Equal(t, AllInjected(ctrl3), false)
	assert.Equal(t, dic.Inject(ctrl3), false)
	// assert.Equal(t, dic.Inject(nil), true) -> panic
}
