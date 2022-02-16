package xmodule

import (
	"reflect"
)

const (
	panicInjectIntoNilValue     = "xmodule: inject into nil value"
	panicInjectIntoNonStructPtr = "xmodule: inject into non-struct pointer"
	panicNotInjectedAllFields   = "xmodule: not all fields with module tag are injected"
)

// ==============
// inject related
// ==============

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

// TODO split coreInject

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
		// TODO <<<<<<
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

// TODO AutoProvide

// ====================
// auto provide related
// ====================

type ModuleProvider struct {
	module interface{}
	// ...
}

func ProviderForName(name ModuleName, module interface{}) *ModuleProvider {
	ensureModuleName(name)
	_ = ensureModuleType(module)
	return &ModuleProvider{module: module} // ...
}

func ProviderForType(module interface{}) *ModuleProvider {
	typ := ensureModuleType(module)
}

func ProviderForIntf(interfacePtr interface{}, moduleImpl interface{}) *ModuleProvider {
}

func (m *ModuleContainer) AutoProvide(providers ...*ModuleProvider) error {

}
