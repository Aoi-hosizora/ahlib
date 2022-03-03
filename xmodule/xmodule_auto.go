package xmodule

import (
	"errors"
	"fmt"
	"reflect"
)

// _moduleTagName is the tag name which represents current struct field is a module.
const _moduleTagName = "module"

const (
	panicInjectIntoNilValue     = "xmodule: inject into nil value"
	panicInjectIntoNonStructPtr = "xmodule: inject into non-struct pointer"
	panicNotInjectedAllFields   = "xmodule: not all fields with module tag are injected" // TODO

	panicInvalidProvider      = "xmodule: using invalid module provider"
	errModulesCycleDependency = "xmodule: modules dependency in cycle"
	errRequiredModuleNotFound = "xmodule: module %s required by %s not found"
)

// ==============
// inject related
// ==============

// Inject injects into injectee fields using `module` tag, returns true if all module fields are injected, that means found and assignable,
// panics when injectee passed is nil or not a structure pointer.
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
// 	all := m.Inject(&Struct{})
func (m *ModuleContainer) Inject(injectee interface{}) (allInjected bool) {
	return coreInject(m, injectee, false)
}

// MustInject injects into injectee fields using `module` tag, panics when injectee passed is nil or not a structure pointer, or there are some
// module fields tag are not injected, that means not found or un-assignable.
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
// 	m.MustInject(&Struct{})
func (m *ModuleContainer) MustInject(injectee interface{}) {
	_ = coreInject(m, injectee, true)
}

// coreInject is the core implementation for Inject and MustInject.
func coreInject(mc *ModuleContainer, injectee interface{}, force bool) (allInjected bool) {
	// check parameters
	if injectee == nil {
		panic(panicInjectIntoNilValue)
	}
	val := reflect.ValueOf(injectee)
	typ := val.Type()
	if typ.Kind() != reflect.Ptr {
		panic(panicInjectIntoNonStructPtr)
	}
	injecteeName := typ.String()
	val = val.Elem()
	typ = typ.Elem()
	if typ.Kind() != reflect.Struct {
		panic(panicInjectIntoNonStructPtr)
	}

	// inject to struct fields
	return injectToStructFields(mc, typ, val, injecteeName, true, force)
}

// injectToStructFields injects modules to given struct's fields, panics only when some fields are not injected and force flag is enabled.
func injectToStructFields(mc *ModuleContainer, structType reflect.Type, structValue reflect.Value, injecteeTypeName string, lock, force bool) (allInjected bool) {
	allInjected = true // module fields are all injected
	injectedCount := 0 // count of injected module fields

	// for each struct field
	for idx := 0; idx < structType.NumField(); idx++ {
		sf, sv := structType.Field(idx), structValue.Field(idx)
		moduleTag := sf.Tag.Get(_moduleTagName)
		if moduleTag == "" || moduleTag == "-" {
			continue // not a module field
		}

		// inject to a module field
		injected := injectToSingleField(mc, moduleTag, sf.Type, sv, lock)
		if !injected {
			if force {
				// failed to inject (specific module not found / cannot inject module to field) -> panic
				panic(panicNotInjectedAllFields)
			}
			allInjected = false
			continue
		}

		mc.logger.InjField(moduleTag, injecteeTypeName, sf.Name, sf.Type.String())
		injectedCount++
	}

	mc.logger.InjFinish(injecteeTypeName, injectedCount, allInjected)
	return allInjected
}

// injectToSingleField checks whether the module for specific field exists and injectable, and injects it to given field.
func injectToSingleField(mc *ModuleContainer, moduleTag string, fieldType reflect.Type, fieldValue reflect.Value, lock bool) (injected bool) {
	// generate key in module map
	key := mkey{}
	if moduleTag != "~" {
		// by name
		key.name = ModuleName(moduleTag) // by name, field tag -> module tag
	} else {
		key.typ = fieldType // by type or intf, field type -> module type
	}

	// check module existence
	if lock {
		mc.mu.RLock()
	}
	module, exist := mc.modules[key]
	if lock {
		mc.mu.RUnlock()
	}
	if !exist {
		// specific module not found
		return false
	}

	// check field injectable
	moduleVal := reflect.ValueOf(module)
	moduleType := moduleVal.Type()
	injectable := fieldValue.CanSet() && moduleType.AssignableTo(fieldType)
	if !injectable {
		// cannot inject module to field
		return false
	}

	// inject module to field
	fieldValue.Set(moduleVal)
	return true
}

// ====================
// auto provide related
// ====================

type ModuleProvider struct {
	mod interface{}
	key mkey

	modType reflect.Type
	byName  bool
	byType  bool
	byIntf  bool
}

type dModuleProvider struct {
	provider     *ModuleProvider
	moduleFields []*reflect.StructField
	moduleKeys   []mkey
}

func NameProvider(name ModuleName, module interface{}) *ModuleProvider {
	ensureModuleName(name)
	typ := ensureModuleType(module)
	return &ModuleProvider{mod: module, key: nameKey(name), modType: typ, byName: true}
}

func TypeProvider(module interface{}) *ModuleProvider {
	typ := ensureModuleType(module)
	return &ModuleProvider{mod: module, key: typeKey(typ), modType: typ, byType: true}
}

func IntfProvider(interfacePtr interface{}, moduleImpl interface{}) *ModuleProvider {
	intfType := ensureInterfacePtr(interfacePtr)               // interface type
	typ := ensureModuleTypeWithInterface(moduleImpl, intfType) // module type
	return &ModuleProvider{mod: moduleImpl, key: typeKey(intfType), modType: typ, byIntf: true}
}

func (m *ModuleContainer) AutoProvide(providers ...*ModuleProvider) error {
	return coreAutoProvide(m, providers, false)
}

func (m *ModuleContainer) MustAutoProvide(providers ...*ModuleProvider) {
	_ = coreAutoProvide(m, providers, true)
}

// coreAutoProvide is the core implementation for AutoProvide and MustAutoProvide.
func coreAutoProvide(mc *ModuleContainer, providers []*ModuleProvider, force bool) error {
	independent, dependent := analyseProvidersDependency(providers)
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// independent providers -> provide directly
	if len(independent) > 0 {
		for _, p := range independent {
			autoProvideAfterInject(mc, p, false)
		}
	}

	// dependent providers -> provide in dependency graph order
	if len(dependent) == 0 {
		return nil
	}
	err := provideInDependencyGraphOrder(mc, dependent)
	if err != nil && force {
		panic(err)
	}
	return err
}

// analyseProvidersDependency checks given ModuleProvider, analyses and extracts them into two providers for independent module and dependent with other modules.
func analyseProvidersDependency(providers []*ModuleProvider) (independent []*ModuleProvider, dependent map[mkey]*dModuleProvider) {
	independent = make([]*ModuleProvider, 0)
	dependent = make(map[mkey]*dModuleProvider, 0)

	for _, p := range providers {
		// check current provider
		if p == nil {
			continue
		}
		if p.mod == nil || (p.key.name == "" && p.key.typ == nil) {
			panic(panicInvalidProvider)
		}

		// check struct pointer type
		typ := reflect.TypeOf(p.mod)
		var fields []*reflect.StructField
		var keys []mkey
		if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}

			// extract module fields
			fields = make([]*reflect.StructField, 0)
			keys = make([]mkey, 0)
			for idx := 0; idx < typ.NumField(); idx++ {
				sf := typ.Field(idx)
				tag := sf.Tag.Get(_moduleTagName)
				if tag == "" || tag == "-" {
					continue
				}
				key := mkey{}
				if tag != "~" {
					// by name
					key.name = ModuleName(tag) // by name, field tag -> module tag
				} else {
					key.typ = sf.Type // by type or intf, field type -> module type
				}
				fields = append(fields, &sf)
				keys = append(keys, key)
			}
		}

		// add current provider to result slices
		if len(fields) == 0 {
			independent = append(independent, p)
		} else {
			dmp := &dModuleProvider{provider: p, moduleFields: fields, moduleKeys: keys}
			dependent[p.key] = dmp
		}
	}

	return independent, dependent
}

// provideInDependencyGraphOrder generates dependency graph for given dependent module providers and provides these modules in graph order.
func provideInDependencyGraphOrder(mc *ModuleContainer, providers map[mkey]*dModuleProvider) error {
	// graph whose node is not provided and has some dependents yet
	graph := make(map[mkey][]mkey, len(providers))
	for key, dmp := range providers {
		graph[key] = dmp.moduleKeys // neighbors of node
	}

	// traverse if there are some nodes still not provided
	for len(graph) > 0 {
		// find nodes which can be provided
		zeroOutDegree := make([]mkey, 0) // zero out degree -> no dependents in the current graph
		for key, neighborKeys := range graph {
			contains := false
			for _, neighborKey := range neighborKeys {
				if _, ok := graph[neighborKey]; ok {
					contains = true
					break
				}
			}
			if !contains {
				zeroOutDegree = append(zeroOutDegree, key)
			}
		}
		if len(zeroOutDegree) == 0 {
			return errors.New(errModulesCycleDependency)
		}

		// provide nodes which has zero out degree
		for _, key := range zeroOutDegree {
			dmp := providers[key]
			for _, moduleKey := range dmp.moduleKeys { // these module nodes must not be in the graph
				if _, ok := mc.modules[moduleKey]; !ok {
					return fmt.Errorf(errRequiredModuleNotFound, moduleKey.String(), key.String())
				}
			}
			autoProvideAfterInject(mc, dmp.provider, true)
			delete(graph, key) // just delete the current node from graph
		}
	}

	return nil
}

// autoProvideAfterInject injects given ModuleProvider for struct pointer type, and provides it to ModuleContainer for all types.
func autoProvideAfterInject(mc *ModuleContainer, p *ModuleProvider, inject bool) {
	// check struct pointer type and inject
	if inject {
		val := reflect.ValueOf(p.mod)
		typ := val.Type()
		if typ.Kind() == reflect.Ptr {
			injecteeName := typ.String()
			val = val.Elem()
			typ = typ.Elem()
			if typ.Kind() == reflect.Struct {
				_ = injectToStructFields(mc, typ, val, injecteeName, false, true)
			}
		}
	}

	// provide module
	mc.modules[p.key] = p.mod
	switch {
	case p.byName:
		mc.logger.PrvName(p.key.name.String(), p.modType.String())
	case p.byType:
		mc.logger.PrvType(p.key.typ.String())
	case p.byIntf:
		mc.logger.PrvIntf(p.key.typ.String(), p.modType.String())
	}
}
