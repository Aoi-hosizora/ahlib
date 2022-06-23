package xmodule

import (
	"fmt"
	"reflect"
	"sync"
)

// ModuleName represents a global module name, and it could not be empty, '-' and '~'.
type ModuleName string

// String returns the string value of ModuleName.
func (m ModuleName) String() string {
	return string(m)
}

// mkey represents the key type of module map used in ModuleContainer, currently these fields are exclusive.
type mkey struct {
	name ModuleName   // by name
	typ  reflect.Type // by type or intf
}

// nameKey returns a mkey with given ModuleName.
func nameKey(name ModuleName) mkey {
	return mkey{name: name}
}

// typeKey returns a mkey with given reflect.Type.
func typeKey(typ reflect.Type) mkey {
	return mkey{typ: typ}
}

// mkeyFromField returns a mkey with given field information.
func mkeyFromField(fieldTag string, fieldType reflect.Type) mkey {
	if fieldTag != "~" {
		return mkey{name: ModuleName(fieldTag)} // by name -> use field tag
	}
	return mkey{typ: fieldType} // by type or intf -> use field type
}

// String returns the string value of mkey.
func (m mkey) String() string {
	if m.name != "" {
		return fmt.Sprintf("name:%s", m.name.String())
	}
	if m.typ != nil {
		return fmt.Sprintf("type:%s", m.typ.String())
	}
	return "<invalid>"
}

// ModuleContainer represents a module container, modules will be stored in a single map using its ModuleName (ProvideByName) or its
// reflect.Type (ProvideByType or ProvideByIntf).
type ModuleContainer struct {
	modules map[mkey]interface{}
	mu      sync.RWMutex
	logger  Logger
}

// NewModuleContainer creates an empty ModuleContainer, using DefaultLogger with LogAll flag and default formats.
func NewModuleContainer() *ModuleContainer {
	return &ModuleContainer{
		modules: make(map[mkey]interface{}),
		logger:  DefaultLogger(LogAll, nil, nil),
	}
}

// SetLogger sets given Logger for ModuleContainer, default logger can be got from DefaultLogger.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // logs all messages (default)
// 	SetLogger(DefaultLogger(LogSilent)) // disable all logger
func (m *ModuleContainer) SetLogger(logger Logger) {
	m.logger = logger
}

// ================
// ensure functions
// ================

const (
	panicInvalidModuleName   = "xmodule: invalid module name (empty, '-' and '~')"
	panicNilModule           = "xmodule: nil module"
	panicInvalidInterfacePtr = "xmodule: invalid interface pointer"
	panicNotImplemented      = "xmodule: module does not implement the interface"
	panicModuleNotFound      = "xmodule: module not found"
)

// ensureModuleName checks given ModuleName, panics if using invalid name, that is empty, '-' and '~'.
func ensureModuleName(name ModuleName) {
	if name == "" || name == "-" || name == "~" {
		panic(panicInvalidModuleName)
	}
}

// ensureModuleType checks whether given module is nil, panics if not, otherwise returns the reflect.Type.
func ensureModuleType(module interface{}) reflect.Type {
	if module == nil {
		panic(panicNilModule)
	}
	return reflect.TypeOf(module)
}

// ensureInterfacePtr checks whether given value is an interface pointer, panics if not, otherwise returns the interface reflect.Type.
func ensureInterfacePtr(interfacePtr interface{}) reflect.Type {
	if interfacePtr == nil {
		panic(panicInvalidInterfacePtr)
	}
	typ := reflect.TypeOf(interfacePtr)
	if typ.Kind() != reflect.Ptr {
		panic(panicInvalidInterfacePtr)
	}
	typ = typ.Elem() // interface type
	if typ.Kind() != reflect.Interface {
		panic(panicInvalidInterfacePtr)
	}
	return typ
}

// ensureModuleTypeWithInterface checks whether given value is nil, and whether it implements given interface type, panics if not, otherwise
// returns its reflect.Type.
func ensureModuleTypeWithInterface(moduleImpl interface{}, interfaceType reflect.Type) reflect.Type {
	typ := ensureModuleType(moduleImpl) // module type
	if !typ.Implements(interfaceType) {
		panic(panicNotImplemented)
	}
	return typ
}

// =============================
// methods: Provide, Remove, Get
// =============================

// ProvideByName provides a module using given ModuleName, panics when using invalid module name or nil module.
//
// Example:
// 	m := NewModuleContainer()
// 	m.ProvideByName(ModuleName("module"), &Module{})
// 	module, ok := m.GetByName(ModuleName("module"))
// 	module := m.MustGetByName(ModuleName("module"))
// 	removed := m.RemoveByName(ModuleName("module"))
func (m *ModuleContainer) ProvideByName(name ModuleName, module interface{}) {
	ensureModuleName(name)
	typ := ensureModuleType(module)
	m.mu.Lock()
	m.modules[nameKey(name)] = module
	m.mu.Unlock()
	m.logger.PrvName(name.String(), typ.String())
}

// ProvideByType provides a module using its type, panics when using nil module.
//
// Example:
// 	m := NewModuleContainer()
// 	m.ProvideByType(&Module{})
// 	module, ok := m.GetByType(&Module{})
// 	module := m.MustGetByType(&Module{})
// 	removed := m.RemoveByType(&Module{})
func (m *ModuleContainer) ProvideByType(module interface{}) {
	typ := ensureModuleType(module)
	m.mu.Lock()
	m.modules[typeKey(typ)] = module
	m.mu.Unlock()
	m.logger.PrvType(typ.String())
}

// ProvideByIntf provides a module using given interface pointer type, such as `(*Interface)(nil)` or `new(Interface)`, panics when using
// invalid interface pointer or nil module.
//
// Example:
// 	m := NewModuleContainer()
// 	m.ProvideByIntf((*Interface)(nil), &Module{})
// 	module, ok := m.GetByIntf((*Interface)(nil))
// 	module := m.MustGetByIntf((*Interface)(nil))
// 	removed := m.RemoveByIntf((*Interface)(nil))
func (m *ModuleContainer) ProvideByIntf(interfacePtr interface{}, moduleImpl interface{}) {
	intfType := ensureInterfacePtr(interfacePtr)                   // interface type
	modType := ensureModuleTypeWithInterface(moduleImpl, intfType) // module type
	m.mu.Lock()
	m.modules[typeKey(intfType)] = moduleImpl
	m.mu.Unlock()
	m.logger.PrvIntf(intfType.String(), modType.String())
}

// RemoveByName removes a module with a ModuleName from container, return true if module existed before removing, panics when using invalid
// module name.
func (m *ModuleContainer) RemoveByName(name ModuleName) (removed bool) {
	ensureModuleName(name)
	m.mu.Lock()
	l := len(m.modules)
	delete(m.modules, nameKey(name))
	removed = len(m.modules) != l
	m.mu.Unlock()
	return removed
}

// RemoveByType removes given module with its type from container, return true if module existed before removing, panics when using nil module.
func (m *ModuleContainer) RemoveByType(moduleType interface{}) (removed bool) {
	typ := ensureModuleType(moduleType)
	m.mu.Lock()
	l := len(m.modules)
	delete(m.modules, typeKey(typ))
	removed = len(m.modules) != l
	m.mu.Unlock()
	return removed
}

// RemoveByIntf removes a module with given interface pointer's type from container, return true if module existed before removing, panics when
// using invalid interface pointer.
func (m *ModuleContainer) RemoveByIntf(interfacePtr interface{}) (removed bool) {
	intfType := ensureInterfacePtr(interfacePtr) // interface type
	m.mu.Lock()
	l := len(m.modules)
	delete(m.modules, typeKey(intfType))
	removed = len(m.modules) != l
	m.mu.Unlock()
	return removed
}

// GetByName returns the module provided by name, panics when using invalid module name.
func (m *ModuleContainer) GetByName(name ModuleName) (module interface{}, exist bool) {
	ensureModuleName(name)
	m.mu.RLock()
	module, exist = m.modules[nameKey(name)]
	m.mu.RUnlock()
	return module, exist
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
	typ := ensureModuleType(moduleType)
	m.mu.RLock()
	module, exist = m.modules[typeKey(typ)]
	m.mu.RUnlock()
	return module, exist
}

// MustGetByType returns a module provided by type, panics when using nil type or module not found.
func (m *ModuleContainer) MustGetByType(moduleType interface{}) interface{} {
	module, exist := m.GetByType(moduleType)
	if !exist {
		panic(panicModuleNotFound)
	}
	return module
}

// GetByIntf returns a module by interface pointer, panics when using invalid interface pointer.
func (m *ModuleContainer) GetByIntf(interfacePtr interface{}) (module interface{}, exist bool) {
	intfType := ensureInterfacePtr(interfacePtr) // interface type
	m.mu.RLock()
	module, exist = m.modules[typeKey(intfType)]
	m.mu.RUnlock()
	return module, exist
}

// MustGetByIntf returns a module by interface pointer, panics when using invalid interface pointer or module not found.
func (m *ModuleContainer) MustGetByIntf(interfacePtr interface{}) interface{} {
	module, exist := m.GetByIntf(interfacePtr)
	if !exist {
		panic(panicModuleNotFound)
	}
	return module
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

// ProvideByName provides a module using given ModuleName, panics when using invalid module name or nil module.
//
// Example:
// 	xmodule.ProvideByName(ModuleName("module"), &Module{})
// 	module, ok := xmodule.GetByName(ModuleName("module"))
// 	module := xmodule.MustGetByName(ModuleName("module"))
// 	removed := xmodule.RemoveByName(ModuleName("module"))
func ProvideByName(name ModuleName, module interface{}) {
	_mc.ProvideByName(name, module)
}

// ProvideByType provides a module using its type, panics when using nil module.
//
// Example:
// 	xmodule.ProvideByType(&Module{})
// 	module, ok := xmodule.GetByType(&Module{})
// 	module := xmodule.MustGetByType(&Module{})
// 	removed := xmodule.RemoveByType(&Module{})
func ProvideByType(module interface{}) {
	_mc.ProvideByType(module)
}

// ProvideByIntf provides a module using given interface pointer type, such as `(*Interface)(nil)` or `new(Interface)`, panics when using
// invalid interface pointer or nil module.
//
// Example:
// 	xmodule.ProvideByIntf((*Interface)(nil), &Module{})
// 	module, ok := xmodule.GetByIntf((*Interface)(nil))
// 	module := xmodule.MustGetByIntf((*Interface)(nil))
// 	removed := xmodule.RemoveByIntf((*Interface)(nil))
func ProvideByIntf(interfacePtr interface{}, moduleImpl interface{}) {
	_mc.ProvideByIntf(interfacePtr, moduleImpl)
}

// RemoveByName removes a module with a ModuleName from container, return true if module existed before removing, panics when using invalid
// module name.
func RemoveByName(name ModuleName) (removed bool) {
	return _mc.RemoveByName(name)
}

// RemoveByType removes given module with its type from container, return true if module existed before removing, panics when using nil module.
func RemoveByType(module interface{}) (removed bool) {
	return _mc.RemoveByType(module)
}

// RemoveByIntf removes a module with given interface pointer's type from container, return true if module existed before removing, panics when
// using invalid interface pointer.
func RemoveByIntf(interfacePtr interface{}) (removed bool) {
	return _mc.RemoveByIntf(interfacePtr)
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

// GetByIntf returns a module by interface pointer, panics when using invalid interface pointer.
func GetByIntf(interfacePtr interface{}) (module interface{}, exist bool) {
	return _mc.GetByIntf(interfacePtr)
}

// MustGetByIntf returns a module by interface pointer, panics when using invalid interface pointer or module not found.
func MustGetByIntf(interfacePtr interface{}) interface{} {
	return _mc.MustGetByIntf(interfacePtr)
}

// Inject injects into given injectee's module fields, returns error if there are some fields can not be injected (possible reasons: specific
// module is not found, module type mismatches with field), panics when injectee passed is nil or not a structure pointer. Note that if error
// is returned, remaining fields will still be injected as usual.
//
// Example:
// 	type Struct struct {
// 		unexportedField string                 // -> ignore (unexported)
// 		ExportedField1  string                 // -> ignore (no module tag)
// 		ExportedField2  string `module:""`     // -> ignore (module tag is empty)
// 		ExportedField3  string `module:"-"`    // -> ignore (module tag is "-")
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or intf
// 	}
// 	m := NewModuleContainer()
// 	all := Inject(&Struct{})
func Inject(injectee interface{}) error {
	return _mc.Inject(injectee)
}

// MustInject injects into given injectee's module fields, panics when injectee passed is nil or not a structure pointer, or there are some fields
// can not be injected for several reasons. Note that remaining fields will stop injecting once error happened. See Inject for more details.
func MustInject(injectee interface{}) {
	_mc.MustInject(injectee)
}

// AutoProvide processes with given ModuleProvider-s, injects them if necessary (must be a pointer of struct), and provides them in dependency
// order, returns error if some fields from providers can not be injected (see Inject for more details), or some dependent modules is not found,
// or cycle dependency happens, panics when using invalid provider.
//
// Example:
// 	wellKnownList := []int{...}
// 	type Service struct {
// 		WellKnownList  []int     `module:"list"`
// 		AnotherService *ServiceB `module:"~"`
// 		Implement      Interface `module:"~"`
// 		LocalVariable  string // a local variable for Service
// 	}
// 	m := NewModuleContainer()
// 	_ = m.AutoProvide(
// 		TypeProvider(&Service{LocalVariable: "..."}),
// 		TypeProvider(&ServiceB{...}),
// 		NameProvider("list", wellKnownList),
// 		IntfProvider((*Interface)(nil), &Implement{}),
// 	)
// 	_ = m.MustGetByType(&Service{}).(*Service)
func AutoProvide(providers ...*ModuleProvider) error {
	return _mc.AutoProvide(providers...)
}

// MustAutoProvide processes with given ModuleProvider-s, injects them if necessary and provides them in dependency order, panics when error happens.
// See AutoProvide for more details.
func MustAutoProvide(providers ...*ModuleProvider) {
	_mc.MustAutoProvide(providers...)
}
