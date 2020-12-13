package xdi

import (
	"fmt"
	"reflect"
	"sync"
)

// ServiceName represents a global service name, and it could not be [ |-|~].
type ServiceName string

// String returns the string value of ServiceName.
func (s ServiceName) String() string {
	return string(s)
}

// DiContainer represents a dependency injection container, or a module container.
type DiContainer struct {
	// provByName saves the services provided by name.
	provByName map[ServiceName]interface{}

	// muByName locks the provByName.
	muByName sync.Mutex

	// provByType saves the services provided by type.
	provByType map[reflect.Type]interface{}

	// muByType locks the provByType.
	muByType sync.Mutex

	// logger represents the log for DiContainer.
	logger Logger
}

// NewDiContainer creates an empty DiContainer with Logger with LogAll flag.
func NewDiContainer() *DiContainer {
	return &DiContainer{
		provByName: make(map[ServiceName]interface{}),
		provByType: make(map[reflect.Type]interface{}),
		logger:     DefaultLogger(LogAll),
	}
}

// SetLogger sets the Logger for DiContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // set default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func (d *DiContainer) SetLogger(logger Logger) {
	_di.logger = logger
}

var (
	invalidServiceNamePanic = "xdi: using invalid service name (empty, '-' and '~')"
	nilServicePanic         = "xdi: using nil service"
	nilInterfacePtrPanic    = "xdi: using nil interface pointer"
	NonInterfacePtrPanic    = "xdi: using non-interface pointer"

	notImplementInterfacePanic = "xdi: service (type of %s) do not implement the interface (type of %s)"
	serviceNotFoundPanic       = "xdi: service not found"

	injectIntoNilPanic          = "xdi: inject into nil struct"
	injectIntoNonStructPtrPanic = "xdi: inject into non-struct pointer"
	notAllFieldsInjectedPanic   = "xdi: not all fields with di tag are injected"
)

// ProvideName provides a service using a ServiceName, panics when using invalid service name or nil service.
func (d *DiContainer) ProvideName(name ServiceName, service interface{}) {
	if name == "" || name == "-" || name == "~" {
		panic(invalidServiceNamePanic)
	}
	if service == nil {
		panic(nilServicePanic)
	}

	d.muByName.Lock()
	d.provByName[name] = service
	d.muByName.Unlock()

	d.logger.LogName(name.String(), reflect.TypeOf(service).String())
}

// ProvideType provides a service using its type, panics when using nil service.
func (d *DiContainer) ProvideType(service interface{}) {
	if service == nil {
		panic(nilServicePanic)
	}
	typ := reflect.TypeOf(service)

	d.muByType.Lock()
	d.provByType[typ] = service
	d.muByType.Unlock()

	d.logger.LogType(typ.String())
}

// ProvideImpl provides a service using the interface type, panics when using invalid interface pointer or nil service.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Service{})
// 	GetByImpl((*Interface)(nil))
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	if interfacePtr == nil {
		panic(nilInterfacePtrPanic)
	}
	if serviceImpl == nil {
		panic(nilServicePanic)
	}

	itfTyp := reflect.TypeOf(interfacePtr)
	if itfTyp.Kind() != reflect.Ptr {
		panic(NonInterfacePtrPanic)
	}
	itfTyp = itfTyp.Elem()
	if itfTyp.Kind() != reflect.Interface {
		panic(NonInterfacePtrPanic)
	}
	srvTyp := reflect.TypeOf(serviceImpl)
	if !srvTyp.Implements(itfTyp) {
		panic(fmt.Sprintf(notImplementInterfacePanic, srvTyp.String(), itfTyp.String()))
	}

	d.muByType.Lock()
	d.provByType[itfTyp] = serviceImpl // interface type
	d.muByType.Unlock()

	d.logger.LogImpl(itfTyp.String(), srvTyp.String())
}

// GetByName returns the service provided by name, panics when using invalid service name.
func (d *DiContainer) GetByName(name ServiceName) (service interface{}, exist bool) {
	if name == "" || name == "~" || name == "-" {
		panic(invalidServiceNamePanic)
	}

	d.muByName.Lock()
	service, exist = d.provByName[name]
	d.muByName.Unlock()
	return
}

// GetByNameForce returns a service provided by name, panics when using invalid service name or service not found.
func (d *DiContainer) GetByNameForce(name ServiceName) interface{} {
	service, exist := d.GetByName(name)
	if !exist {
		panic(serviceNotFoundPanic)
	}
	return service
}

// GetByType returns a service provided by type, panics when using nil type.
func (d *DiContainer) GetByType(serviceType interface{}) (service interface{}, exist bool) {
	if serviceType == nil {
		panic(nilServicePanic)
	}

	typ := reflect.TypeOf(serviceType)
	d.muByType.Lock()
	service, exist = d.provByType[typ]
	d.muByType.Unlock()
	return
}

// GetByTypeForce returns a service provided by type, panics when using nil type or service not found.
func (d *DiContainer) GetByTypeForce(serviceType interface{}) interface{} {
	service, exist := d.GetByType(serviceType)
	if !exist {
		panic(serviceNotFoundPanic)
	}
	return service
}

// GetByImpl returns a service by interface pointer, panics when using invalid interface pointer.
func (d *DiContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	if interfacePtr == nil {
		panic(nilInterfacePtrPanic)
	}
	itfTyp := reflect.TypeOf(interfacePtr)
	if itfTyp.Kind() != reflect.Ptr {
		panic(NonInterfacePtrPanic)
	}
	itfTyp = itfTyp.Elem()
	if itfTyp.Kind() != reflect.Interface {
		panic(NonInterfacePtrPanic)
	}

	d.muByType.Lock()
	service, exist = d.provByType[itfTyp] // interface type
	d.muByType.Unlock()
	return
}

// GetByImplForce returns a service by serviceType, panics when using invalid interface pointer or service not found.
func (d *DiContainer) GetByImplForce(interfacePtr interface{}) interface{} {
	service, exist := d.GetByImpl(interfacePtr)
	if !exist {
		panic(serviceNotFoundPanic)
	}
	return service
}

// ====
// core
// ====

// Inject injects into struct fields using its di tag.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string             // -> ignore
// 		ExportedField1  string             // -> ignore
// 		ExportedField2  string `di:""`     // -> ignore
// 		ExportedField3  string `di:"-"`    // -> ignore
// 		ExportedField4  string `di:"name"` // -> inject by name
// 		ExportedField5  string `di:"~"`    // -> inject by type or impl
// 	}
func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	return coreInject(d, ctrl, false)
}

// InjectForce injects into struct fields using its di tag, panics when not all fields are injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string             // -> ignore
// 		ExportedField1  string             // -> ignore
// 		ExportedField2  string `di:""`     // -> ignore
// 		ExportedField3  string `di:"-"`    // -> ignore
// 		ExportedField4  string `di:"name"` // -> inject by name
// 		ExportedField5  string `di:"~"`    // -> inject by type or impl
// 	}
func (d *DiContainer) InjectForce(ctrl interface{}) {
	coreInject(d, ctrl, true)
}

// coreInject is the core implementation for Inject and InjectForce.
func coreInject(d *DiContainer, ctrl interface{}, force bool) bool {
	if ctrl == nil {
		panic(injectIntoNilPanic)
	}
	ctrlTyp := reflect.TypeOf(ctrl)
	ctrlVal := reflect.ValueOf(ctrl)
	if ctrlTyp.Kind() != reflect.Ptr {
		panic(injectIntoNonStructPtrPanic)
	}
	ctrlTyp = ctrlTyp.Elem()
	ctrlVal = ctrlVal.Elem()
	if ctrlTyp.Kind() != reflect.Struct {
		panic(injectIntoNonStructPtrPanic)
	}

	// record is all injected
	allInjected := true

	// for each field
	for idx := 0; idx < ctrlTyp.NumField(); idx++ {
		// check
		field := ctrlTyp.Field(idx)
		diTag := field.Tag.Get("di")
		if diTag == "-" || diTag == "" {
			continue
		}

		// find
		var service interface{}
		var exist bool
		if diTag != "~" {
			// inject by name
			d.muByName.Lock()
			service, exist = d.provByName[ServiceName(diTag)]
			d.muByName.Unlock()
		} else {
			// inject by type or impl
			d.muByType.Lock()
			service, exist = d.provByType[field.Type]
			d.muByType.Unlock()
		}

		// exist
		if !exist {
			if force {
				// if force inject and service not found, panic
				panic(notAllFieldsInjectedPanic)
			}
			allInjected = false
			continue
		}

		// inject value
		fieldVal := ctrlVal.Field(idx)
		if fieldVal.IsValid() && fieldVal.CanSet() {
			fieldVal.Set(reflect.ValueOf(service))
			d.logger.LogInject(reflect.TypeOf(ctrl).String(), field.Type.String(), field.Name)
		}
	}

	return allInjected
}

// _di is a global DiContainer.
var _di = NewDiContainer()

// SetLogger sets the Logger for DiContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // set default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func SetLogger(logger Logger) {
	_di.SetLogger(logger)
}

// ProvideName provides a service using a ServiceName, panics when using invalid service name or nil service.
func ProvideName(name ServiceName, service interface{}) {
	_di.ProvideName(name, service)
}

// ProvideType provides a service using its type, panics when using nil service.
func ProvideType(service interface{}) {
	_di.ProvideType(service)
}

// ProvideImpl provides a service using the interface type, panics when using invalid interface pointer or nil service.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Service{})
// 	GetByImpl((*Interface)(nil))
func ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	_di.ProvideImpl(interfacePtr, serviceImpl)
}

// GetByName returns the service provided by name, panics when using invalid service name.
func GetByName(name ServiceName) (service interface{}, exist bool) {
	return _di.GetByName(name)
}

// GetByNameForce returns a service provided by name, panics when using invalid service name or service not found.
func GetByNameForce(name ServiceName) interface{} {
	return _di.GetByNameForce(name)
}

// GetByType returns a service provided by type, panics when using nil type.
func GetByType(serviceType interface{}) (service interface{}, exist bool) {
	return _di.GetByType(serviceType)
}

// GetByTypeForce returns a service provided by type, panics when using nil type or service not found.
func GetByTypeForce(serviceType interface{}) interface{} {
	return _di.GetByTypeForce(serviceType)
}

// GetByImpl returns a service by interface pointer, panics when using invalid interface pointer.
func GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	return _di.GetByImpl(interfacePtr)
}

// GetByImplForce returns a service by serviceType, panics when using invalid interface pointer or service not found.
func GetByImplForce(interfacePtr interface{}) interface{} {
	return _di.GetByImplForce(interfacePtr)
}

// Inject injects into struct fields using its di tag.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string             // -> ignore
// 		ExportedField1  string             // -> ignore
// 		ExportedField2  string `di:""`     // -> ignore
// 		ExportedField3  string `di:"-"`    // -> ignore
// 		ExportedField4  string `di:"name"` // -> inject by name
// 		ExportedField5  string `di:"~"`    // -> inject by type or impl
// 	}
func Inject(ctrl interface{}) (allInjected bool) {
	return _di.Inject(ctrl)
}

// InjectForce injects into struct fields using its di tag, panics when not all fields are injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string             // -> ignore
// 		ExportedField1  string             // -> ignore
// 		ExportedField2  string `di:""`     // -> ignore
// 		ExportedField3  string `di:"-"`    // -> ignore
// 		ExportedField4  string `di:"name"` // -> inject by name
// 		ExportedField5  string `di:"~"`    // -> inject by type or impl
// 	}
func InjectForce(ctrl interface{}) {
	_di.InjectForce(ctrl)
}
