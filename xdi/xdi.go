package xdi

import (
	"reflect"
)

type DiContainer struct {
	_provByType map[reflect.Type]interface{}
	_provByName map[string]interface{}
}

func NewDiContainer() *DiContainer {
	var dic = &DiContainer{}
	dic._provByType = make(map[reflect.Type]interface{})
	dic._provByName = make(map[string]interface{})
	return dic
}

// service: can be normal type or struct
func (d *DiContainer) Provide(service interface{}) {
	t := reflect.TypeOf(service)
	d._provByType[t] = service
}

// name: could not be ~, can be normal type or struct
func (d *DiContainer) ProvideByName(name string, service interface{}) {
	if name == "~" {
		panic("Could not preserved key ~ as the name of service")
	}
	d._provByName[name] = service
}

// interfacePtr: (*Interface)(nil), impl: Struct or *Struct
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, impl interface{}) {
	it := reflect.TypeOf(interfacePtr).Elem()
	st := reflect.TypeOf(impl)
	// fmt.Println(it, st)
	if !st.Implements(it) {
		panic("Could not implement type of " + it.String() + " by " + st.String())
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
func (d *DiContainer) Inject(ctrl interface{}) {
	var ctrlType = reflect.TypeOf(ctrl)
	for ctrlType.Kind() == reflect.Ptr {
		ctrlType = ctrlType.Elem()
	}
	if ctrlType.Kind() != reflect.Struct {
		panic("Object for injection should be struct, have " + ctrlType.String())
	}

	for fieldIdx := 0; fieldIdx < ctrlType.NumField(); fieldIdx++ {
		field := ctrlType.Field(fieldIdx)
		diTag := field.Tag.Get("di")
		if diTag == "-" || diTag == "" {
			continue
		}

		var service interface{}
		var ok bool
		if diTag == "~" {
			service, ok = d._provByType[field.Type]
		} else {
			service, ok = d._provByName[diTag]
		}
		if ok {
			ctrlValue := reflect.ValueOf(ctrl)
			for ctrlValue.Kind() == reflect.Ptr {
				ctrlValue = ctrlValue.Elem()
			}
			ctrlField := ctrlValue.Field(fieldIdx)

			if ctrlField.IsValid() && ctrlField.CanSet() {
				srvValue := reflect.ValueOf(service)
				ctrlField.Set(srvValue)
			}
		}
	}
}

// check if di field is nil
func HasNilDi(ctrl interface{}) bool {
	var ctrlType = reflect.TypeOf(ctrl)
	for ctrlType.Kind() == reflect.Ptr {
		ctrlType = ctrlType.Elem()
	}
	if ctrlType.Kind() != reflect.Struct {
		panic("Object for injection should be struct, have " + ctrlType.String())
	}

	for idx := 0; idx < ctrlType.NumField(); idx++ {
		field := ctrlType.Field(idx)
		diTag := field.Tag.Get("di")
		if diTag == "" || diTag == "-" {
			continue
		}

		ctrlValue := reflect.ValueOf(ctrl)
		for ctrlValue.Kind() == reflect.Ptr {
			ctrlValue = ctrlValue.Elem()
		}
		ctrlField := ctrlValue.Field(idx)

		if ctrlField.IsValid() && ctrlField.CanSet() {
			if ctrlField.IsZero() {
				return true
			}
		}
	}
	return false
}
