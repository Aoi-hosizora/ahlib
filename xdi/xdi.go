package xdi

import (
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

// service: can be normal type or struct
func (d *DiContainer) Provide(service interface{}) {
	if service == nil {
		panic("could not provide nil service")
	}
	t := reflect.TypeOf(service)
	d._provByType[t] = service
}

// name: could not be ~, can be normal type or struct
func (d *DiContainer) ProvideByName(name string, service interface{}) {
	if name == "~" {
		panic("could not provide service using ~ name")
	}
	d._provByName[name] = service
}

// interfacePtr: (*Interface)(nil), impl: Struct or *Struct
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, impl interface{}) {
	it := reflect.TypeOf(interfacePtr)
	if reflect.TypeOf(it).Kind() != reflect.Ptr {
		panic("parameter of impl could be only ptr")
	}
	it = it.Elem()
	st := reflect.TypeOf(impl)
	// fmt.Println(it, st)
	if !st.Implements(it) {
		panic("could not implement type of " + it.String() + " by " + st.String())
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
func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	var ctrlType = xcommon.ElemType(ctrl)
	var ctrlValue = xcommon.ElemValue(ctrl)
	if ctrlType.Kind() != reflect.Struct {
		panic("Object for injection should be struct, have " + ctrlType.String())
	}
	allInjected = true

	for fieldIdx := 0; fieldIdx < ctrlType.NumField(); fieldIdx++ {
		field := ctrlType.Field(fieldIdx)
		diTag := field.Tag.Get("di")
		if diTag == "-" || diTag == "" {
			continue
		}

		var service interface{}
		var exist bool
		if diTag == "~" {
			service, exist = d._provByType[field.Type]
		} else {
			service, exist = d._provByName[diTag]
		}

		if !exist {
			allInjected = false
		} else {
			ctrlField := ctrlValue.Field(fieldIdx)

			if ctrlField.IsValid() && ctrlField.CanSet() {
				srvValue := reflect.ValueOf(service)
				ctrlField.Set(srvValue)
			}
		}
	}
	return allInjected
}

// check if all field needed inject is not nil
func AllInjected(ctrl interface{}) bool {
	var ctrlType = xcommon.ElemType(ctrl)
	var ctrlValue = xcommon.ElemValue(ctrl)
	if ctrlType.Kind() != reflect.Struct {
		return true
	}

	for idx := 0; idx < ctrlType.NumField(); idx++ {
		field := ctrlType.Field(idx)
		diTag := field.Tag.Get("di")
		if diTag == "" || diTag == "-" {
			continue
		}

		ctrlField := ctrlValue.Field(idx)
		if ctrlField.IsValid() && ctrlField.CanSet() && ctrlField.IsZero() {
			return false
		}
	}
	return true
}
