package xmodule

import (
	"fmt"
	"reflect"
	"sync"
)

// ModuleName represents a global module name, and it could not be empty, - and ~.
type ModuleName string

// String returns the string value of ModuleName.
func (s ModuleName) String() string {
	return string(s)
}

// ModuleContainer represents a module container.
type ModuleContainer struct {
	// provByName saves the modules provided by name.
	provByName map[ModuleName]interface{}

	// muByName locks the provByName.
	muByName sync.Mutex

	// provByType saves the modules provided by type.
	provByType map[reflect.Type]interface{}

	// muByType locks the provByType.
	muByType sync.Mutex

	// logger represents the log for ModuleContainer.
	logger Logger
}

// NewModuleContainer creates an empty ModuleContainer with Logger with LogAll flag.
func NewModuleContainer() *ModuleContainer {
	return &ModuleContainer{
		provByName: make(map[ModuleName]interface{}),
		provByType: make(map[reflect.Type]interface{}),
		logger:     DefaultLogger(LogAll),
	}
}

// SetLogger sets the Logger for ModuleContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // set to default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func (m *ModuleContainer) SetLogger(logger Logger) {
	m.logger = logger
}

const (
	invalidModuleNamePanic = "xmodule: using invalid module name (empty, '-' and '~')"
	nilModulePanic         = "xmodule: using nil module"
	nilInterfacePtrPanic   = "xmodule: using nil interface pointer"
	NonInterfacePtrPanic   = "xmodule: using non-interface pointer"

	notImplementInterfacePanic = "xmodule: module (type of %s) do not implement the interface (type of %s)"
	moduleNotFoundPanic        = "xmodule: module not found"

	injectIntoNilPanic          = "xmodule: inject into nil struct"
	injectIntoNonStructPtrPanic = "xmodule: inject into non-struct pointer"
	notAllFieldsInjectedPanic   = "xmodule: not all fields with module tag are injected"
)

// ProvideName provides a module using a ModuleName, panics when using invalid module name or nil module.
func (m *ModuleContainer) ProvideName(name ModuleName, module interface{}) {
	if name == "" || name == "-" || name == "~" {
		panic(invalidModuleNamePanic)
	}
	if module == nil {
		panic(nilModulePanic)
	}

	m.muByName.Lock()
	m.provByName[name] = module
	m.muByName.Unlock()

	m.logger.LogName(name.String(), reflect.TypeOf(module).String())
}

// ProvideType provides a module using its type, panics when using nil module.
func (m *ModuleContainer) ProvideType(module interface{}) {
	if module == nil {
		panic(nilModulePanic)
	}
	typ := reflect.TypeOf(module)

	m.muByType.Lock()
	m.provByType[typ] = module
	m.muByType.Unlock()

	m.logger.LogType(typ.String())
}

// ProvideImpl provides a module using the interface type, panics when using invalid interface pointer or nil module.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Module{})
// 	GetByImpl((*Interface)(nil))
func (m *ModuleContainer) ProvideImpl(interfacePtr interface{}, moduleImpl interface{}) {
	if interfacePtr == nil {
		panic(nilInterfacePtrPanic)
	}
	if moduleImpl == nil {
		panic(nilModulePanic)
	}

	itfTyp := reflect.TypeOf(interfacePtr)
	if itfTyp.Kind() != reflect.Ptr {
		panic(NonInterfacePtrPanic)
	}
	itfTyp = itfTyp.Elem()
	if itfTyp.Kind() != reflect.Interface {
		panic(NonInterfacePtrPanic)
	}
	modTyp := reflect.TypeOf(moduleImpl)
	if !modTyp.Implements(itfTyp) {
		panic(fmt.Sprintf(notImplementInterfacePanic, modTyp.String(), itfTyp.String()))
	}

	m.muByType.Lock()
	m.provByType[itfTyp] = moduleImpl // interface type
	m.muByType.Unlock()

	m.logger.LogImpl(itfTyp.String(), modTyp.String())
}

// GetByName returns the module provided by name, panics when using invalid module name.
func (m *ModuleContainer) GetByName(name ModuleName) (module interface{}, exist bool) {
	if name == "" || name == "~" || name == "-" {
		panic(invalidModuleNamePanic)
	}

	m.muByName.Lock()
	module, exist = m.provByName[name]
	m.muByName.Unlock()
	return
}

// MustGetByName returns a module provided by name, panics when using invalid module name or module not found.
func (m *ModuleContainer) MustGetByName(name ModuleName) interface{} {
	module, exist := m.GetByName(name)
	if !exist {
		panic(moduleNotFoundPanic)
	}
	return module
}

// GetByType returns a module provided by type, panics when using nil type.
func (m *ModuleContainer) GetByType(moduleType interface{}) (module interface{}, exist bool) {
	if moduleType == nil {
		panic(nilModulePanic)
	}

	typ := reflect.TypeOf(moduleType)
	m.muByType.Lock()
	module, exist = m.provByType[typ]
	m.muByType.Unlock()
	return
}

// MustGetByType returns a module provided by type, panics when using nil type or module not found.
func (m *ModuleContainer) MustGetByType(moduleType interface{}) interface{} {
	module, exist := m.GetByType(moduleType)
	if !exist {
		panic(moduleNotFoundPanic)
	}
	return module
}

// GetByImpl returns a module by interface pointer, panics when using invalid interface pointer.
func (m *ModuleContainer) GetByImpl(interfacePtr interface{}) (module interface{}, exist bool) {
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

	m.muByType.Lock()
	module, exist = m.provByType[itfTyp] // interface type
	m.muByType.Unlock()
	return
}

// MustGetByImpl returns a module by moduleType, panics when using invalid interface pointer or module not found.
func (m *ModuleContainer) MustGetByImpl(interfacePtr interface{}) interface{} {
	module, exist := m.GetByImpl(interfacePtr)
	if !exist {
		panic(moduleNotFoundPanic)
	}
	return module
}

// ====
// core
// ====

// Inject injects into struct fields using its module tag, returns true if all fields with `module` tag has been injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func (m *ModuleContainer) Inject(ctrl interface{}) (allInjected bool) {
	return coreInject(m, ctrl, false)
}

// MustInject injects into struct fields using its module tag, panics when not all fields with `module` tag are injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func (m *ModuleContainer) MustInject(ctrl interface{}) {
	coreInject(m, ctrl, true)
}

// coreInject is the core implementation for Inject and MustInject.
func coreInject(d *ModuleContainer, ctrl interface{}, force bool) bool {
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
		moduleTag := field.Tag.Get("module")
		if moduleTag == "-" || moduleTag == "" {
			continue
		}

		// find
		var module interface{}
		var exist bool
		if moduleTag != "~" {
			// inject by name
			d.muByName.Lock()
			module, exist = d.provByName[ModuleName(moduleTag)]
			d.muByName.Unlock()
		} else {
			// inject by type or impl
			d.muByType.Lock()
			module, exist = d.provByType[field.Type]
			d.muByType.Unlock()
		}

		// exist
		if !exist {
			if force {
				// if force inject and module not found, panic
				panic(notAllFieldsInjectedPanic)
			}
			allInjected = false
			continue
		}

		// inject value
		fieldVal := ctrlVal.Field(idx)
		if fieldVal.IsValid() && fieldVal.CanSet() {
			fieldVal.Set(reflect.ValueOf(module))
			d.logger.LogInject(reflect.TypeOf(ctrl).String(), field.Type.String(), field.Name)
		}
	}

	return allInjected
}

// _module is a global ModuleContainer.
var _module = NewModuleContainer()

// SetLogger sets the Logger for ModuleContainer.
//
// Example:
// 	SetLogger(DefaultLogger(LogAll))    // set to default logger
// 	SetLogger(DefaultLogger(LogSilent)) // disable logger
func SetLogger(logger Logger) {
	_module.SetLogger(logger)
}

// ProvideName provides a module using a ModuleName, panics when using invalid module name or nil module.
func ProvideName(name ModuleName, module interface{}) {
	_module.ProvideName(name, module)
}

// ProvideType provides a module using its type, panics when using nil module.
func ProvideType(module interface{}) {
	_module.ProvideType(module)
}

// ProvideImpl provides a module using the interface type, panics when using invalid interface pointer or nil module.
//
// Example:
// 	ProvideImpl((*Interface)(nil), &Module{})
// 	GetByImpl((*Interface)(nil))
func ProvideImpl(interfacePtr interface{}, moduleImpl interface{}) {
	_module.ProvideImpl(interfacePtr, moduleImpl)
}

// GetByName returns the module provided by name, panics when using invalid module name.
func GetByName(name ModuleName) (module interface{}, exist bool) {
	return _module.GetByName(name)
}

// MustGetByName returns a module provided by name, panics when using invalid module name or module not found.
func MustGetByName(name ModuleName) interface{} {
	return _module.MustGetByName(name)
}

// GetByType returns a module provided by type, panics when using nil type.
func GetByType(moduleType interface{}) (module interface{}, exist bool) {
	return _module.GetByType(moduleType)
}

// MustGetByType returns a module provided by type, panics when using nil type or module not found.
func MustGetByType(moduleType interface{}) interface{} {
	return _module.MustGetByType(moduleType)
}

// GetByImpl returns a module by interface pointer, panics when using invalid interface pointer.
func GetByImpl(interfacePtr interface{}) (module interface{}, exist bool) {
	return _module.GetByImpl(interfacePtr)
}

// MustGetByImpl returns a module by moduleType, panics when using invalid interface pointer or module not found.
func MustGetByImpl(interfacePtr interface{}) interface{} {
	return _module.MustGetByImpl(interfacePtr)
}

// Inject injects into struct fields using its module tag, returns true if all fields with `module` tag has been injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func Inject(ctrl interface{}) (allInjected bool) {
	return _module.Inject(ctrl)
}

// MustInject injects into struct fields using its module tag, panics when not all fields with `module` tag are injected.
//
// Example:
// 	type AStruct struct {
// 		unexportedField string                 // -> ignore
// 		ExportedField1  string                 // -> ignore
// 		ExportedField2  string `module:""`     // -> ignore
// 		ExportedField3  string `module:"-"`    // -> ignore
// 		ExportedField4  string `module:"name"` // -> inject by name
// 		ExportedField5  string `module:"~"`    // -> inject by type or impl
// 	}
func MustInject(ctrl interface{}) {
	_module.MustInject(ctrl)
}
