package xdi

import (
	"reflect"
)

type DiContainer interface {
	Provide(service interface{})
	ProvideByName(name string, service interface{})
	ProvideImpl(interfacePtr interface{}, impl interface{})
	Inject(ctrl interface{})
}

type _DiContainer struct {
	_provByType map[reflect.Type]interface{}
	_provByName map[string]interface{}
}

func NewDiContainer() DiContainer {
	var dic = &_DiContainer{}
	dic._provByType = make(map[reflect.Type]interface{})
	dic._provByName = make(map[string]interface{})
	return dic
}

// service: can be normal type or struct
func (d *_DiContainer) Provide(service interface{}) {
	t := reflect.TypeOf(service)
	d._provByType[t] = service
}

// name: could not be ~, can be normal type or struct
func (d *_DiContainer) ProvideByName(name string, service interface{}) {
	if name == "~" {
		panic("Could not preserved key ~ as the name of service")
	}
	d._provByName[name] = service
}

// interfacePtr: (*Interface)(nil), impl: Struct or *Struct
func (d *_DiContainer) ProvideImpl(interfacePtr interface{}, impl interface{}) {
	it := reflect.TypeOf(interfacePtr).Elem()
	st := reflect.TypeOf(impl)
	// fmt.Println(it, st)
	if !st.Implements(it) {
		panic("Could not implement type of " + it.String() + " by " + st.String())
	}
	d._provByType[it] = impl
}

// diTag: "" || - -> ignore
// diTag: ~       -> auto inject
// diTag: name    -> inject by name
func (d *_DiContainer) Inject(ctrl interface{}) {
	var ctrlType = reflect.TypeOf(ctrl)
	if ctrlType.Kind() != reflect.Ptr || ctrlType.Elem().Kind() != reflect.Struct {
		panic("Object for injection should be pointer to struct, have " + ctrlType.String())
	}
	ctrlType = ctrlType.Elem()

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
			ctrlField := reflect.ValueOf(ctrl).Elem().Field(fieldIdx)
			srvValue := reflect.ValueOf(service)
			ctrlField.Set(srvValue)
		}
	}
}
