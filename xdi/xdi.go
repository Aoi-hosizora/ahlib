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
// 	SetLogger(DefaultLogger(LogAll))    // set default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func (d *DiContainer) SetLogger(logger Logger) {
	_di.logger = logger
}

var (
	provideInvalidNamePanic    = "xdi: service name could not be empty, '-' or '~'"
	provideNilServicePanic     = "xdi: provide nil service"
	nilInterfacePtrPanic       = "xdi: nil interfacePtr"
	invalidInterfacePtrPanic   = "xdi: non-interface-pointer interfacePtr"
	notImplementInterfacePanic = "xdi: service (type of %s) do not implement the interface (type of %s)"

	serviceByNameNotFoundPanic = "xdi: service provided by name not found"
	serviceByTypeNotFoundPanic = "xdi: service provided by type not found"
	serviceByImplNotFoundPanic = "xdi: service provided by interface type not found"
	serviceByInvalidNamePanic  = "xdi: get service using empty, '-' or '~' name"
	serviceByNilPanic          = "xdi: get service using nil"

	injectIntoNilPanic          = "xdi: inject into nil"
	injectIntoNonStructPtrPanic = "xdi: inject into non-struct-pointer value"
	notAllFieldsInjectedPanic   = "xdi: not all fields with di tag are injected"
)

// ProvideName provides a service using a ServiceName, panic when using empty or `-` or `~` name or nil service.
func (d *DiContainer) ProvideName(name ServiceName, service interface{}) {
	if name == "" || name == "-" || name == "~" {
		panic(provideInvalidNamePanic)
	}
	if service == nil {
		panic(provideNilServicePanic)
	}

	d.muByName.Lock()
	d.provByName[name] = service
	d.muByName.Unlock()

	d.logger.LogName(name.String(), reflect.TypeOf(service).String())
}

// ProvideType provides a service using its type, panic when using nil service.
func (d *DiContainer) ProvideType(service interface{}) {
	if service == nil {
		panic(provideNilServicePanic)
	}
	typ := reflect.TypeOf(service)

	d.muByType.Lock()
	d.provByType[typ] = service
	d.muByType.Unlock()

	d.logger.LogType(typ.String())
}

// ProvideImpl provides a service using the interface type, panic when using invalid interfacePtr or nil serviceImpl.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Struct{})
// 	GetByType(Interface(nil))
func (d *DiContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	if interfacePtr == nil {
		panic(nilInterfacePtrPanic)
	}
	if serviceImpl == nil {
		panic(provideNilServicePanic)
	}

	itfTyp := reflect.TypeOf(interfacePtr)
	if itfTyp.Kind() != reflect.Ptr {
		panic(invalidInterfacePtrPanic)
	}
	itfTyp = itfTyp.Elem()
	if itfTyp.Kind() != reflect.Interface {
		panic(invalidInterfacePtrPanic)
	}
	srvTyp := reflect.TypeOf(serviceImpl)
	if !srvTyp.Implements(itfTyp) {
		panic(fmt.Sprintf(notImplementInterfacePanic, srvTyp.String(), itfTyp.String()))
	}

	d.muByType.Lock()
	d.provByType[itfTyp] = serviceImpl
	d.muByType.Unlock()

	d.logger.LogImpl(itfTyp.String(), srvTyp.String())
}

// GetByName returns the service provided by name, panics when using invalid service name, returns false if not exists.
func (d *DiContainer) GetByName(name ServiceName) (service interface{}, exist bool) {
	if name == "" || name == "~" || name == "-" {
		panic(serviceByInvalidNamePanic)
	}
	service, exist = d.provByName[name]
	return
}

// GetByNameForce returns a service provided by name, panics when using invalid service name or the service not found.
func (d *DiContainer) GetByNameForce(name ServiceName) interface{} {
	service, exist := d.GetByName(name)
	if !exist {
		panic(serviceByNameNotFoundPanic)
	}
	return service
}

// GetByType returns a service provided by type, panics when using nil type, returns false if not exists.
func (d *DiContainer) GetByType(serviceType interface{}) (service interface{}, exist bool) {
	if serviceType == nil {
		panic(serviceByNilPanic)
	}

	typ := reflect.TypeOf(serviceType)
	service, exist = d.provByType[typ]
	return
}

// GetByTypeForce returns a service provided by type, panics when using nil type or the service not found.
func (d *DiContainer) GetByTypeForce(serviceType interface{}) interface{} {
	service, exist := d.GetByType(serviceType)
	if !exist {
		panic(serviceByTypeNotFoundPanic)
	}
	return service
}

// GetByImpl returns a service by interface pointer, panics when using invalid interfacePtr, returns false if not exists.
func (d *DiContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	if interfacePtr == nil {
		panic(nilInterfacePtrPanic)
	}
	itfTyp := reflect.TypeOf(interfacePtr)
	if itfTyp.Kind() != reflect.Ptr {
		panic(invalidInterfacePtrPanic)
	}
	itfTyp = itfTyp.Elem()
	if itfTyp.Kind() != reflect.Interface {
		panic(invalidInterfacePtrPanic)
	}

	service, exist = d.provByType[itfTyp]
	return
}

// GetByImplForce returns a service by serviceType, panics when using invalid interfacePtr or the service not found.
func (d *DiContainer) GetByImplForce(interfacePtr interface{}) interface{} {
	service, exist := d.GetByImpl(interfacePtr)
	if !exist {
		panic(serviceByImplNotFoundPanic)
	}
	return service
}

// Inject injects into the struct fields by di tag, and returns true if all fields with di tag are both injected.
// Example:
// 	type AStruct struct {
// 		unexportedField1 string           // -> ignore
// 		ExportedField1 string             // -> ignore
// 		ExportedField2 string `di:""`     // -> ignore
// 		ExportedField3 string `di:"-"`    // -> ignore
// 		ExportedField4 string `di:"name"` // -> inject by name
// 		ExportedField5 string `di:"~"`    // -> inject by type
// 	}
func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool) {
	return d.inject(ctrl, false)
}

// InjectForce injects into the struct fields by di tag, same as Inject, but panic when not all fields with di tag are injected.
func (d *DiContainer) InjectForce(ctrl interface{}) {
	d.inject(ctrl, true)
}

// MustInject injects into the struct fields by di tag, same as Inject, but panic when not all fields with di tag are injected.
func (d *DiContainer) MustInject(ctrl interface{}) {
	d.InjectForce(ctrl)
}

// inject is the core implementation for Inject, InjectForce and MustInject.
func (d *DiContainer) inject(ctrl interface{}, force bool) bool {
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
	for fieldIdx := 0; fieldIdx < ctrlTyp.NumField(); fieldIdx++ {
		// check
		field := ctrlTyp.Field(fieldIdx)
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
				panic(notAllFieldsInjectedPanic)
			}
			allInjected = false
			continue
		}

		// inject
		fieldType := ctrlTyp.Field(fieldIdx)
		fieldValue := ctrlVal.Field(fieldIdx)
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

// SetLogger sets the Logger for DiContainer.
// 	SetLogger(DefaultLogger(LogAll))    // set default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func SetLogger(logger Logger) {
	_di.SetLogger(logger)
}

// ProvideName provides a service using a ServiceName, panic when using empty or `-` or `~` name or nil service.
func ProvideName(name ServiceName, service interface{}) {
	_di.ProvideName(name, service)
}

// ProvideType provides a service using its type, panic when using nil service.
func ProvideType(service interface{}) {
	_di.ProvideType(service)
}

// ProvideImpl provides a service using the interface type, panic when using invalid interfacePtr or nil serviceImpl.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Struct{})
// 	GetByType(Interface(nil))
func ProvideImpl(interfacePtr interface{}, serviceImpl interface{}) {
	_di.ProvideImpl(interfacePtr, serviceImpl)
}

// GetByName returns the service provided by name, panics when using invalid service name, returns false if not exists.
func GetByName(name ServiceName) (service interface{}, exist bool) {
	return _di.GetByName(name)
}

// GetByNameForce returns a service provided by name, panics when using invalid service name or the service not found.
func GetByNameForce(name ServiceName) interface{} {
	return _di.GetByNameForce(name)
}

// GetByType returns a service provided by type, panics when using nil type, returns false if not exists.
func GetByType(serviceType interface{}) (service interface{}, exist bool) {
	return _di.GetByType(serviceType)
}

// GetByTypeForce returns a service provided by type, panics when using nil type or the service not found.
func GetByTypeForce(serviceType interface{}) interface{} {
	return _di.GetByTypeForce(serviceType)
}

// GetByImpl returns a service by interface pointer, panics when using invalid interfacePtr, returns false if not exists.
func GetByImpl(interfacePtr interface{}) (service interface{}, exist bool) {
	return _di.GetByImpl(interfacePtr)
}

// GetByImplForce returns a service by serviceType, panics when using invalid interfacePtr or the service not found.
func GetByImplForce(interfacePtr interface{}) interface{} {
	return _di.GetByImplForce(interfacePtr)
}

// Inject injects into the struct fields by di tag, and returns true if all fields with di tag are both injected.
// Example:
// 	type AStruct struct {
// 		unexportedField1 string           // -> ignore
// 		ExportedField1 string             // -> ignore
// 		ExportedField2 string `di:""`     // -> ignore
// 		ExportedField3 string `di:"-"`    // -> ignore
// 		ExportedField4 string `di:"name"` // -> inject by name
// 		ExportedField5 string `di:"~"`    // -> inject by type
// 	}
func Inject(ctrl interface{}) (allInjected bool) {
	return _di.Inject(ctrl)
}

// InjectForce injects into the struct fields by di tag, same as Inject, but panic when not all fields with di tag are injected.
func InjectForce(ctrl interface{}) {
	_di.MustInject(ctrl)
}

// MustInject injects into the struct fields by di tag, same as Inject, but panic when not all fields with di tag are injected.
func MustInject(ctrl interface{}) {
	_di.MustInject(ctrl)
}
