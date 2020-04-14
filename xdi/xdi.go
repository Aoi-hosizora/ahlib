package xdi

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcommon"
	"reflect"
)

type DiContainer struct {
	_provByType map[reflect.Type]interface{}
	_provByName map[string]interface{}
}

func NewDiContainer() *DiContainer {
	var dic = &DiContainer{
		_provByType: make(map[reflect.Type]interface{}),
		_provByName: make(map[string]interface{}),
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

// service: can be normal type or struct
func (d *DiContainer) Provide(service interface{}) {
	if service == nil {
		panic(provideNilPanic)
	}
	t := reflect.TypeOf(service)
	d._provByType[t] = service
}

// name: could not be ~, can be normal type or struct
func (d *DiContainer) ProvideByName(name string, service interface{}) {
	if name == "~" {
		panic(providePreservePanic)
	}
	d._provByName[name] = service
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
	var ctrlType = xcommon.ElemType(ctrl)
	var ctrlValue = xcommon.ElemValue(ctrl)
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
		ctrlField := ctrlValue.Field(fieldIdx)
		if ctrlField.IsValid() && ctrlField.CanSet() {
			ctrlField.Set(reflect.ValueOf(service))
		}
	}
	return allInjected
}

func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	return d.inject(ctrl, false)
}

func (d *DiContainer) InjectForce(ctrl interface{}) {
	d.inject(ctrl, true)
}
