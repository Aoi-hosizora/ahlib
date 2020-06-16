package xdi

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"github.com/gookit/color"
	"reflect"
)

// DiContainer log function, yellow for type, red for name
type LogFunc func(kind string, parentName string, fieldName string, fieldType string)

type DiContainer struct {
	_provByType     map[reflect.Type]interface{}
	_provByName     map[string]interface{}
	_provideLogMode bool
	_injectLogMode  bool
	_logFunc        LogFunc
}

func NewDiContainer() *DiContainer {
	var dic = &DiContainer{
		_provByType:     make(map[reflect.Type]interface{}),
		_provByName:     make(map[string]interface{}),
		_provideLogMode: true,
		_injectLogMode:  true,
		_logFunc: func(method string, parent string, name string, t string) {
			method += ":"
			if parent != "" {
				parent = fmt.Sprintf("(%s).", color.Yellow.Sprint(parent))
			}
			name = color.Yellow.Sprint(name)
			t = color.Yellow.Sprint(t)
			fmt.Printf("[XDI] %-12s %s%s (%s)\n", method, parent, name, t)
		},
	}
	return dic
}

var (
	provideNilPanic        = "could not provide nil service"
	providePreservePanic   = "could not provide service using ~ name"
	provideNonPtrImplPanic = "could not provide a non pointer implementation"
	notImplPanic           = "could not implement type %s by %s"
	injectNonStructPanic   = "object for injection should be struct, have %s with %s"
	injectFailedPanic      = "there are some fields could not be injected"
)

// setup a DiContainer log mode, log when Provide and Inject
func (d *DiContainer) SetLogMode(provideLogMode bool, injectLogMode bool) {
	_di._provideLogMode = provideLogMode
	_di._injectLogMode = injectLogMode
}

// setup a DiContainer how to log
// kind: Provide or Inject
// parentName: field's parent when inject
func (d *DiContainer) SetLogFunc(logFunc LogFunc) {
	_di._logFunc = logFunc
}

// service: can be normal type or struct
func (d *DiContainer) Provide(service interface{}) {
	if service == nil {
		panic(provideNilPanic)
	}
	t := reflect.TypeOf(service)
	d._provByType[t] = service
	if d._provideLogMode {
		d._logFunc("Provide", "", "_", t.String())
	}
}

// name: could not be ~, can be normal type or struct
func (d *DiContainer) ProvideByName(name string, service interface{}) {
	if name == "~" {
		panic(providePreservePanic)
	}
	d._provByName[name] = service
	if d._provideLogMode {
		d._logFunc("ProvideName", "", name, reflect.TypeOf(service).String())
	}
}

// interfacePtr: (*Interface)(nil), impl: Struct or *Struct
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, impl interface{}) {
	it := reflect.TypeOf(interfacePtr)
	if reflect.TypeOf(it).Kind() != reflect.Ptr {
		panic(provideNonPtrImplPanic)
	}
	it = it.Elem()
	st := reflect.TypeOf(impl)
	if !st.Implements(it) {
		panic(fmt.Sprintf(notImplPanic, it.String(), st.String()))
	}
	d._provByType[it] = impl
	if d._provideLogMode {
		d._logFunc("ProvideImpl", "", "_", it.String())
	}
}

// get data by type
func (d *DiContainer) GetProvide(srvType interface{}) (service interface{}, exist bool) {
	service, exist = d._provByType[reflect.TypeOf(srvType)]
	return
}

// get data by name
func (d *DiContainer) GetProvideByName(name string) (service interface{}, exist bool) {
	service, exist = d._provByName[name]
	return
}

// diTag: "" || - -> ignore
// diTag: ~       -> auto inject
// diTag: name    -> inject by name
func (d *DiContainer) inject(ctrl interface{}, force bool) bool {
	var ctrlType = xreflect.ElemType(ctrl)
	var ctrlValue = xreflect.ElemValue(ctrl)
	if ctrlType.Kind() != reflect.Struct {
		panic(fmt.Sprintf(injectNonStructPanic, ctrlType.Kind().String(), ctrlType.String()))
	}
	allInjected := true

	for fieldIdx := 0; fieldIdx < ctrlType.NumField(); fieldIdx++ {
		// check
		field := ctrlType.Field(fieldIdx)
		diTag := field.Tag.Get("di")
		if diTag == "-" || diTag == "" {
			continue
		}

		// find
		var service interface{}
		var exist bool
		if diTag == "~" {
			service, exist = d._provByType[field.Type]
		} else {
			service, exist = d._provByName[diTag]
		}

		// exist
		if !exist {
			allInjected = false
			if force {
				panic(injectFailedPanic)
			}
			continue
		}

		// inject
		fieldType := ctrlType.Field(fieldIdx)
		fieldValue := ctrlValue.Field(fieldIdx)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(service))
			if d._injectLogMode {
				d._logFunc("Inject", reflect.TypeOf(ctrl).String(), fieldType.Name, fieldType.Type.String())
			}
		}
	}
	return allInjected
}

func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	return d.inject(ctrl, false)
}

func (d *DiContainer) MustInject(ctrl interface{}) {
	d.inject(ctrl, true)
}

// A DiContainer that used for global
var _di = NewDiContainer()

// A global Provide function for DiContainer
func Provide(service interface{}) {
	_di.Provide(service)
}

// A global ProvideByName function for DiContainer
func ProvideByName(name string, service interface{}) {
	_di.ProvideByName(name, service)
}

// A global ProvideImpl function for DiContainer
func ProvideImpl(interfacePtr interface{}, impl interface{}) {
	_di.ProvideImpl(interfacePtr, impl)
}

// A global GetProvide function for DiContainer
func GetProvide(srvType interface{}) (service interface{}, exist bool) {
	return _di.GetProvide(srvType)
}

// A global Provide function for DiContainer
func GetProvideByName(name string) (service interface{}, exist bool) {
	return _di.GetProvideByName(name)
}

// A global Inject function for DiContainer
func Inject(ctrl interface{}) (allInjected bool) {
	return _di.Inject(ctrl)
}

// A global MustInject function for DiContainer
func MustInject(ctrl interface{}) {
	_di.MustInject(ctrl)
}

// A global SetLogMode function for DiContainer
// only for global DiContainer, not work for NewDiContainer
func SetLogMode(provideLogMode bool, injectLogMode bool) {
	_di._provideLogMode = provideLogMode
	_di._injectLogMode = injectLogMode
}

// A global SetLogFunc function for DiContainer
// only for global DiContainer, not work for NewDiContainer
func SetLogFunc(logFunc LogFunc) {
	_di._logFunc = logFunc
}
