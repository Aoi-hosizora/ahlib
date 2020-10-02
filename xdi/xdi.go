package xdi

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"reflect"
	"sync"
)

// ServiceName represents a global service name.
type ServiceName string

// String returns a string type of ServiceName.
func (s *ServiceName) String() string {
	return string(*s)
}

// DiContainer represents a container for DI.
type DiContainer struct {
	// provByName saves the services provided by name.
	provByName map[ServiceName]interface{}

	// muByName locks the provByName.
	muByName sync.Mutex

	// provByType saves the services provided by type.
	provByType map[reflect.Type]interface{}

	// muByType locks the provByType.
	muByType sync.Mutex

	// logger for DiContainer.
	logger Logger
}

// NewDiContainer creates a default DiContainer.
func NewDiContainer() *DiContainer {
	return &DiContainer{
		provByType: make(map[reflect.Type]interface{}),
		provByName: make(map[ServiceName]interface{}),
		logger:     DefaultLogger(LogAll),
	}
}

// SetLogger sets another Logger for DiContainer.
// 	SetLogger(xdi.DefaultLogger()) // set default logger
// 	SetLogger(xdi.SilentLogger())  // disable logger
func (d *DiContainer) SetLogger(logger Logger) {
	_di.logger = logger
}

// ProvideName provides a service using a ServiceName, panic when using `~` or nil service.
func (d *DiContainer) ProvideName(name ServiceName, service interface{}) {
	if name == "~" {
		panic("could not provide service using '~' name")
	}
	if service == nil {
		panic("could not provide nil service")
	}
	typ := reflect.TypeOf(service)

	d.muByName.Lock()
	d.provByName[name] = service
	d.muByName.Unlock()

	d.logger.LogName(string(name), typ.String())
}

// ProvideType provides a service using its type, panic when service is nil.
func (d *DiContainer) ProvideType(service interface{}) {
	if service == nil {
		panic("could not provide nil service")
	}
	typ := reflect.TypeOf(service)

	d.muByType.Lock()
	d.provByType[typ] = service
	d.muByType.Unlock()

	d.logger.LogType(typ.String())
}

// ProvideImpl provides a service using the interface type, panic when wrong interfacePtr or nil serviceImpl.
// Example:
// 	ProvideImpl((*Interface)(nil), &Struct{})
// 	GetByType(Interface(nil))
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	if interfacePtr == nil {
		panic("interfacePtr could not be nil")
	}
	it := reflect.TypeOf(interfacePtr)
	if it.Kind() != reflect.Ptr {
		panic("interfacePtr must be an pointer of interface")
	}
	it = it.Elem()
	if it.Kind() != reflect.Interface {
		panic("interfacePtr must be an pointer of interface")
	}

	if serviceImpl == nil {
		panic("could not provide nil service")
	}
	st := reflect.TypeOf(serviceImpl)
	if !st.Implements(it) {
		panic(fmt.Sprintf("could not implement type %s by %s", it.String(), st.String()))
	}

	d.muByType.Lock()
	d.provByType[it] = serviceImpl
	d.muByType.Unlock()

	d.logger.LogImpl(it.String(), st.String())
}

// GetByName returns a service by ServiceName.
func (d *DiContainer) GetByName(name ServiceName) (service interface{}, exist bool) {
	if name == "~" {
		panic("could not get the service of ~")
	}

	service, exist = d.provByName[name]
	return
}

// GetByNameForce returns a service by ServiceName, panic when service not found.
func (d *DiContainer) GetByNameForce(name ServiceName) interface{} {
	service, exist := d.GetByName(name)
	if !exist {
		panic(fmt.Sprintf("service with name %s is not found", name))
	}
	return service
}

// GetByType returns a service by serviceType.
func (d *DiContainer) GetByType(serviceType interface{}) (service interface{}, exist bool) {
	if serviceType == nil {
		panic("could not get nil type service")
	}

	typ := reflect.TypeOf(serviceType)
	service, exist = d.provByType[typ]
	return
}

// GetByTypeForce returns a service by interface pointer, panic when service not found.
func (d *DiContainer) GetByTypeForce(serviceType interface{}) interface{} {
	service, exist := d.GetByType(serviceType)
	if !exist {
		panic(fmt.Sprintf("service with type %s is not found", reflect.TypeOf(serviceType).String()))
	}
	return service
}

// GetByImpl returns a service by interface pointer, panic when wrong interfacePtr.
func (d *DiContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	if interfacePtr == nil {
		panic("interfacePtr could not be nil")
	}
	it := reflect.TypeOf(interfacePtr)
	if it.Kind() != reflect.Ptr {
		panic("interfacePtr must be an pointer of interface")
	}
	it = it.Elem()
	if it.Kind() != reflect.Interface {
		panic("interfacePtr must be an pointer of interface")
	}

	service, exist = d.provByType[it]
	return
}

// GetByImplForce returns a service by serviceType, panic when wrong interfacePtr or service not found.
func (d *DiContainer) GetByImplForce(interfacePtr interface{}) interface{} {
	service, exist := d.GetByImpl(interfacePtr)
	if !exist {
		panic(fmt.Sprintf("service with type %s is not found", reflect.TypeOf(interfacePtr).Elem().String()))
	}
	return service
}

// Inject injects fields into struct by di tag, and returns if all fields with di tag is injected.
// Example:
// 	`di:""`       // -> ignore
// 	`di:"-"`      // -> ignore
// 	`di:"~"`      // -> auto inject
// 	`di:"name"`   // -> inject by name
func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	return d.inject(ctrl, false)
}

// MustInject injects fields into struct, same with Inject, but panic when not all field with di tag is injected.
func (d *DiContainer) MustInject(ctrl interface{}) {
	d.inject(ctrl, true)
}

// inject injects fields into struct, is the inner implement for Inject and MustInject.
func (d *DiContainer) inject(ctrl interface{}, force bool) bool {
	if ctrl == nil {
		panic("Inject: object for injection must noe bt nil")
	}
	ctrlType := xreflect.ElemType(ctrl)
	ctrlValue := xreflect.ElemValue(ctrl)
	if ctrlType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Inject: object for injection must be struct, have `%s` with `%s`", ctrlType.Kind().String(), ctrlType.String()))
	}

	// record is all injected
	allInjected := true

	// for each field
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
			// inject by type
			service, exist = d.provByType[field.Type]
		} else {
			// inject by name
			service, exist = d.provByName[ServiceName(diTag)]
		}

		// exist
		if !exist {
			if force {
				// if force inject and service not found, panic
				panic("Inject: there are some fields could not be injected")
			}
			allInjected = false
			continue
		}

		// inject
		fieldType := ctrlType.Field(fieldIdx)
		fieldValue := ctrlValue.Field(fieldIdx)
		if fieldValue.IsValid() && fieldValue.CanSet() {
			// set value
			fieldValue.Set(reflect.ValueOf(service))
			d.logger.LogInject(reflect.TypeOf(ctrl).String(), fieldType.Type.String(), fieldType.Name)
		}
	}

	return allInjected
}

// _di is a global DiContainer.
var _di = NewDiContainer()

// SetLogger sets another Logger for DiContainer.
// 	xdi.SetLogger(xdi.DefaultLogger()) // set default logger
// 	xdi.SetLogger(xdi.SilentLogger())  // disable logger
func SetLogger(logger Logger) {
	_di.SetLogger(logger)
}

// ProvideName provides a service using a ServiceName, panic when name is `~` (preserve name).
func ProvideName(name ServiceName, service interface{}) {
	_di.ProvideName(name, service)
}

// ProvideType provides a service using its type, panic when service is nil.
func ProvideType(service interface{}) {
	_di.ProvideType(service)
}

// ProvideImpl provides a service using the interface type, panic when wrong interfacePtr or nil serviceImpl.
func ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	_di.ProvideImpl(interfacePtr, serviceImpl)
}

// GetByName returns a service by ServiceName.
func GetByName(name ServiceName) (service interface{}, exist bool) {
	return _di.GetByName(name)
}

// GetByNameForce returns a service by ServiceName, panic when service not found.
func GetByNameForce(name ServiceName) interface{} {
	return _di.GetByNameForce(name)
}

// GetByType returns a service by serviceType.
func GetByType(serviceType interface{}) (service interface{}, exist bool) {
	return _di.GetByType(serviceType)
}

// GetByTypeForce returns a service by serviceType, panic when service not found.
func GetByTypeForce(serviceType interface{}) interface{} {
	return _di.GetByTypeForce(serviceType)
}

// GetByImpl returns a service by interface pointer, panic when wrong interfacePtr.
func GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	return _di.GetByImpl(interfacePtr)
}

// GetByImplForce returns a service by serviceType, panic when wrong interfacePtr or service not found.
func GetByImplForce(interfacePtr interface{}) interface{} {
	return _di.GetByImplForce(interfacePtr)
}

// Inject injects fields into struct by di tag, and returns if all fields with di tag is injected.
// Example:
// 	`di:""`       // -> ignore
// 	`di:"-"`      // -> ignore
// 	`di:"~"`      // -> auto inject
// 	`di:"name"`   // -> inject by name
func Inject(ctrl interface{}) (allInjected bool) {
	return _di.Inject(ctrl)
}

// MustInject injects fields into struct, same with Inject, but panic when not all field with di tag is injected.
func MustInject(ctrl interface{}) {
	_di.MustInject(ctrl)
}
