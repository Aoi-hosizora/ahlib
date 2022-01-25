package xmodule

import (
	"reflect"
	"sync"
)

// ModuleName represents a global module name, and it could not be empty, '-' and '~'.
type ModuleName string

// String returns the string value of ModuleName.
func (m ModuleName) String() string {
	return string(m)
}

// ModuleContainer represents a module container.
type ModuleContainer struct {
	byName map[ModuleName]interface{}
	muName sync.RWMutex
	byType map[reflect.Type]interface{}
	muType sync.RWMutex
	logger Logger
}

// NewModuleContainer creates an empty ModuleContainer, using DefaultLogger with LogAll flag and default formats.
func NewModuleContainer() *ModuleContainer {
	return &ModuleContainer{
		byName: make(map[ModuleName]interface{}),
		byType: make(map[reflect.Type]interface{}),
		logger: DefaultLogger(LogAll, nil, nil),
	}
}

// =======
// methods
// =======

// SetLogger sets the Logger for ModuleContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // logs all messages (default)
// 	SetLogger(DefaultLogger(LogSilent)) // disable all logger
func (m *ModuleContainer) SetLogger(logger Logger) {
	m.logger = logger
}

const (
	panicInvalidModuleName   = "xmodule: invalid module name (empty, '-' and '~')"
	panicNilModule           = "xmodule: nil module"
	panicInvalidInterfacePtr = "xmodule: invalid interface pointer"
	panicNotImplemented      = "xmodule: module does not implement the interface"
	panicModuleNotFound      = "xmodule: module not found"

	panicInjectIntoNilValue     = "xmodule: inject into nil value"
	panicInjectIntoNonStructPtr = "xmodule: inject into non-struct pointer"
	panicNotInjectedAllFields   = "xmodule: not all fields with module tag are injected"
)

// assertModuleName checks given ModuleName, panics if using invalid name, that is empty, '-' and '~'.
func assertModuleName(name ModuleName) {
	if name == "" || name == "-" || name == "~" {
		panic(panicInvalidModuleName)
	}
}

// assertModuleType checks whether given module is nil, panics if not, and returns the reflect.Type.
func assertModuleType(module interface{}) reflect.Type {
	if module == nil {
		panic(panicNilModule)
	}
	return reflect.TypeOf(module)
}

// assertInterfacePtr checks whether given value is an interface pointer, panics if not, and returns the interface reflect.Type.
func assertInterfacePtr(interfacePtr interface{}) reflect.Type {
	if interfacePtr == nil {
		panic(panicInvalidInterfacePtr)
	}
	itfType := reflect.TypeOf(interfacePtr)
	if itfType.Kind() != reflect.Ptr {
		panic(panicInvalidInterfacePtr)
	}
	itfType = itfType.Elem() // interface type
	if itfType.Kind() != reflect.Interface {
		panic(panicInvalidInterfacePtr)
	}
	return itfType
}

// ProvideName provides a module using a ModuleName, panics when using invalid module name or nil module.
//
// Example:
// 	ProvideName(ModuleName("module"), &Module{})
// 	RemoveByName(ModuleName("module"))
// 	GetByName(ModuleName("module"))
// 	MustGetByName(ModuleName("module"))
func (m *ModuleContainer) ProvideName(name ModuleName, module interface{}) {
	assertModuleName(name)
	typ := assertModuleType(module)
	m.muName.Lock()
	m.byName[name] = module
	m.muName.Unlock()
	m.logger.PrvName(name.String(), typ.String())
}

// ProvideType provides a module using its type, panics when using nil module.
//
// Example:
// 	ProvideType(&Module{})
// 	RemoveByType(&Module{})
// 	GetByType(&Module{})
// 	MustGetByType(&Module{})
func (m *ModuleContainer) ProvideType(module interface{}) {
	typ := assertModuleType(module)
	m.muType.Lock()
	m.byType[typ] = module
	m.muType.Unlock()
	m.logger.PrvType(typ.String())
}

// ProvideImpl provides a module using the interface type, panics when using invalid interface pointer or nil module.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Module{})
// 	RemoveByImpl((*Interface)(nil))
// 	GetByImpl((*Interface)(nil))
// 	MustGetByImpl((*Interface)(nil))
func (m *ModuleContainer) ProvideImpl(interfacePtr interface{}, moduleImpl interface{}) {
	itfType := assertInterfacePtr(interfacePtr) // interface type
	innerType := assertModuleType(moduleImpl)   // inner type
	if !innerType.Implements(itfType) {
		panic(panicNotImplemented)
	}
	m.muType.Lock()
	m.byType[itfType] = moduleImpl
	m.muType.Unlock()
	m.logger.PrvImpl(itfType.String(), innerType.String())
}

// RemoveByName remove a module with a ModuleName from container, panics when using invalid module name.
func (m *ModuleContainer) RemoveByName(name ModuleName) {
	assertModuleName(name)
	m.muName.Lock()
	delete(m.byName, name)
	m.muName.Unlock()
}

// RemoveByType remove given module with its type from container, panics when using nil module.
func (m *ModuleContainer) RemoveByType(moduleType interface{}) {
	typ := assertModuleType(moduleType)
	m.muType.Lock()
	delete(m.byType, typ)
	m.muType.Unlock()
}

// RemoveByImpl remove a module with given interface pointer's type from container, panics when using invalid interface pointer.
func (m *ModuleContainer) RemoveByImpl(interfacePtr interface{}) {
	itfType := assertInterfacePtr(interfacePtr) // interface type
	m.muType.Lock()
	delete(m.byType, itfType)
	m.muType.Unlock()
}

// GetByName returns the module provided by name, panics when using invalid module name.
func (m *ModuleContainer) GetByName(name ModuleName) (module interface{}, exist bool) {
	assertModuleName(name)
	m.muName.RLock()
	module, exist = m.byName[name]
	m.muName.RUnlock()
	return
}

// MustGetByName returns a module provided by name, panics when using invalid module name or module not found.
func (m *ModuleContainer) MustGetByName(name ModuleName) interface{} {
	module, exist := m.GetByName(name)
	if !exist {
		panic(panicModuleNotFound)
	}
	return module
}

// GetByType returns a module provided by type, panics when using nil type.
func (m *ModuleContainer) GetByType(moduleType interface{}) (module interface{}, exist bool) {
	typ := assertModuleType(moduleType)
	m.muType.RLock()
	module, exist = m.byType[typ]
	m.muType.RUnlock()
	return
}

// MustGetByType returns a module provided by type, panics when using nil type or module not found.
func (m *ModuleContainer) MustGetByType(moduleType interface{}) interface{} {
	module, exist := m.GetByType(moduleType)
	if !exist {
		panic(panicModuleNotFound)
	}
	return module
}

// GetByImpl returns a module by interface pointer, panics when using invalid interface pointer.
func (m *ModuleContainer) GetByImpl(interfacePtr interface{}) (module interface{}, exist bool) {
	itfType := assertInterfacePtr(interfacePtr) // interface type
	m.muType.RLock()
	module, exist = m.byType[itfType]
	m.muType.RUnlock()
	return
}

// MustGetByImpl returns a module by moduleType, panics when using invalid interface pointer or module not found.
func (m *ModuleContainer) MustGetByImpl(interfacePtr interface{}) interface{} {
	module, exist := m.GetByImpl(interfacePtr)
	if !exist {
		panic(panicModuleNotFound)
	}
	return module
}

// Inject injects into injectee fields using module tag, returns true if all fields with `module` tag are injected, that means found and assignable.
//
// Example:
// 	type Struct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func (m *ModuleContainer) Inject(injectee interface{}) (allInjected bool) {
	return coreInject(m, injectee, false)
}

// MustInject injects into injectee fields using module tag, panics when some fields with `module` tag are not injected, that means not found or un-assignable.
//
// Example:
// 	type Struct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func (m *ModuleContainer) MustInject(v interface{}) {
	coreInject(m, v, true)
}

// coreInject is the core implementation for Inject and MustInject.
func coreInject(mc *ModuleContainer, injectee interface{}, force bool) (allInjected bool) {
	if injectee == nil {
		panic(panicInjectIntoNilValue)
	}
	val := reflect.ValueOf(injectee)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr {
		panic(panicInjectIntoNonStructPtr)
	}
	typName := typ.String()
	typ = typ.Elem()
	val = val.Elem()
	if typ.Kind() != reflect.Struct {
		panic(panicInjectIntoNonStructPtr)
	}

	allInjected = true // record is all injected
	injectedCount := 0 // injected fields count

	// for each field
	for idx := 0; idx < typ.NumField(); idx++ {
		// check tag
		field := typ.Field(idx)
		moduleTag := field.Tag.Get("module")
		if moduleTag == "" || moduleTag == "-" {
			continue
		}

		// check existence
		var module interface{}
		var exist bool
		if moduleTag != "~" {
			// inject by name
			mc.muName.RLock()
			module, exist = mc.byName[ModuleName(moduleTag)]
			mc.muName.RUnlock()
		} else {
			// inject by type or impl
			mc.muType.RLock()
			module, exist = mc.byType[field.Type]
			mc.muType.RUnlock()
		}
		if !exist {
			if force {
				// specific module not found -> panic
				panic(panicNotInjectedAllFields)
			}
			allInjected = false
			continue
		}

		// check injectable
		fieldVal := val.Field(idx)
		moduleVal := reflect.ValueOf(module)
		settable := fieldVal.CanSet() && moduleVal.Type().AssignableTo(field.Type)
		if !settable {
			if force {
				// cannot assign module to field -> panic
				panic(panicNotInjectedAllFields)
			}
			allInjected = false
			continue
		}

		// inject field
		fieldVal.Set(moduleVal)
		mc.logger.InjField(moduleTag, typName, field.Name, field.Type.String())
		injectedCount++
	}

	mc.logger.InjFinish(typName, injectedCount, allInjected)
	return allInjected
}

// ======
// global
// ======

// _mc is a global ModuleContainer.
var _mc = NewModuleContainer()

// SetLogger sets the Logger for ModuleContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // logs all messages (default)
// 	SetLogger(DefaultLogger(LogSilent)) // disable all logger
func SetLogger(logger Logger) {
	_mc.SetLogger(logger)
}

// ProvideName provides a module using a ModuleName, panics when using invalid module name or nil module.
//
// Example:
// 	ProvideName(ModuleName("module"), &Module{})
// 	RemoveByName(ModuleName("module"))
// 	GetByName(ModuleName("module"))
// 	MustGetByName(ModuleName("module"))
func ProvideName(name ModuleName, module interface{}) {
	_mc.ProvideName(name, module)
}

// ProvideType provides a module using its type, panics when using nil module.
//
// Example:
// 	ProvideType(&Module{})
// 	RemoveByType(&Module{})
// 	GetByType(&Module{})
// 	MustGetByType(&Module{})
func ProvideType(module interface{}) {
	_mc.ProvideType(module)
}

// ProvideImpl provides a module using the interface type, panics when using invalid interface pointer or nil module.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Module{})
// 	RemoveByImpl((*Interface)(nil))
// 	GetByImpl((*Interface)(nil))
// 	MustGetByImpl((*Interface)(nil))
func ProvideImpl(interfacePtr interface{}, moduleImpl interface{}) {
	_mc.ProvideImpl(interfacePtr, moduleImpl)
}

// RemoveByName remove a module with a ModuleName from container, panics when using invalid module name.
func RemoveByName(name ModuleName) {
	_mc.RemoveByName(name)
}

// RemoveByType remove given module with its type from container, panics when using nil module.
func RemoveByType(module interface{}) {
	_mc.RemoveByType(module)
}

// RemoveByImpl remove a module with given interface pointer's type from container, panics when using invalid interface pointer.
func RemoveByImpl(interfacePtr interface{}) {
	_mc.RemoveByImpl(interfacePtr)
}

// GetByName returns the module provided by name, panics when using invalid module name.
func GetByName(name ModuleName) (module interface{}, exist bool) {
	return _mc.GetByName(name)
}

// MustGetByName returns a module provided by name, panics when using invalid module name or module not found.
func MustGetByName(name ModuleName) interface{} {
	return _mc.MustGetByName(name)
}

// GetByType returns a module provided by type, panics when using nil type.
func GetByType(moduleType interface{}) (module interface{}, exist bool) {
	return _mc.GetByType(moduleType)
}

// MustGetByType returns a module provided by type, panics when using nil type or module not found.
func MustGetByType(moduleType interface{}) interface{} {
	return _mc.MustGetByType(moduleType)
}

// GetByImpl returns a module by interface pointer, panics when using invalid interface pointer.
func GetByImpl(interfacePtr interface{}) (module interface{}, exist bool) {
	return _mc.GetByImpl(interfacePtr)
}

// MustGetByImpl returns a module by moduleType, panics when using invalid interface pointer or module not found.
func MustGetByImpl(interfacePtr interface{}) interface{} {
	return _mc.MustGetByImpl(interfacePtr)
}

// Inject injects into injectee fields using module tag, returns true if all fields with `module` tag are injected, that means found and assignable.
//
// Example:
// 	type Struct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func Inject(injectee interface{}) (allInjected bool) {
	return _mc.Inject(injectee)
}

// MustInject injects into injectee fields using module tag, panics when some fields with `module` tag are not injected, that means not found or un-assignable.
//
// Example:
// 	type Struct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func MustInject(injectee interface{}) {
	_mc.MustInject(injectee)
}
