package xmodule

import (
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xerror"
	"reflect"
)

// _moduleTagName is the "module" tag name which is used to regard struct field as a module.
const _moduleTagName = "module"

const (
	panicInjectIntoNilValue     = "xmodule: inject into nil value"
	panicInjectIntoNonStructPtr = "xmodule: inject into non-struct pointer"
	panicInvalidProvider        = "xmodule: using nil or invalid module provider"

	errRequiredModuleNotFound = "xmodule: module '%s' required by injectee '%s' is not found"
	errMismatchedModuleType   = "xmodule: module type '%s' mismatches with field type '%s'"
	errModulesCycleDependency = "xmodule: given provided modules have cycle dependency"
)

// =================
// injection related
// =================

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
// 	all := m.Inject(&Struct{})
func (m *ModuleContainer) Inject(injectee interface{}) error {
	return coreInject(m, injectee, false)
}

// MustInject injects into given injectee's module fields, panics when injectee passed is nil or not a structure pointer, or there are some fields
// can not be injected for several reasons. Note that remaining fields will stop injecting once error happened. See Inject for more details.
func (m *ModuleContainer) MustInject(injectee interface{}) {
	_ = coreInject(m, injectee, true)
}

// coreInject is the core implementation for Inject and MustInject.
func coreInject(mc *ModuleContainer, injectee interface{}, force bool) error {
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
	errs := injectToStructFields(mc, typ, val, injecteeName, true, force)
	return xerror.Combine(errs...)
}

// injectToStructFields tries to inject modules to given structure, returns errors when some fields can not be injected.
func injectToStructFields(mc *ModuleContainer, structType reflect.Type, structValue reflect.Value, injecteeTypeName string, lock, force bool) []error {
	var errs []error
	totalCount := 0
	injectedCount := 0

	// for each struct field
	for idx := 0; idx < structType.NumField(); idx++ {
		sf, sv := structType.Field(idx), structValue.Field(idx)
		moduleTag := sf.Tag.Get(_moduleTagName)
		if moduleTag == "" || moduleTag == "-" {
			continue // not a module field
		}

		// inject to module field
		totalCount++
		err := injectToSingleField(mc, moduleTag, sf.Type, sv, injecteeTypeName, lock)
		if err != nil {
			if force {
				panic(err.Error())
			}
			errs = append(errs, err)
			continue // no force -> just continue
		}

		mc.logger.InjField(moduleTag, injecteeTypeName, sf.Name, sf.Type.String())
		injectedCount++
	}

	if totalCount > 0 {
		mc.logger.InjFinish(injecteeTypeName, injectedCount, totalCount)
	}
	return errs
}

// injectToSingleField checks whether the module for given field exists and type matches, and injects it to given field if check passed.
func injectToSingleField(mc *ModuleContainer, fieldTag string, fieldType reflect.Type, fieldValue reflect.Value, injecteeTypeName string, lock bool) error {
	// generate key by field tag and type
	key := mkeyFromField(fieldTag, fieldType) // by name (...), type or intf (~)

	// check module existence
	if lock {
		mc.mu.RLock()
	}
	module, exist := mc.modules[key]
	if lock {
		mc.mu.RUnlock()
	}
	if !exist {
		// specific module is not found
		return fmt.Errorf(errRequiredModuleNotFound, fieldType.String(), injecteeTypeName)
	}

	// check module injectable
	moduleVal := reflect.ValueOf(module)
	moduleType := moduleVal.Type()
	if !moduleType.AssignableTo(fieldType) {
		// module type mismatches with field
		return fmt.Errorf(errMismatchedModuleType, moduleType.String(), fieldType.String())
	}

	// check field assignable, inject module to field
	if fieldValue.CanSet() {
		fieldValue.Set(moduleVal)
	}
	return nil
}

// ====================
// auto provide related
// ====================

// ModuleProvider represents a type for module provider used by AutoProvide. Note that you must create this value by NameProvider, TypeProvider
// or IntfProvider, otherwise it may panic when invoking AutoProvide.
type ModuleProvider struct {
	mod interface{}
	key mkey

	modType reflect.Type
	byName  bool
	byType  bool
	byIntf  bool

	depKeys []mkey // late, update by analyseDependency
}

// NameProvider creates a ModuleProvider, it can be used to provide a module using given ModuleName.
func NameProvider(name ModuleName, module interface{}) *ModuleProvider {
	ensureModuleName(name)
	typ := ensureModuleType(module)
	return &ModuleProvider{mod: module, key: nameKey(name), modType: typ, byName: true}
}

// TypeProvider creates a ModuleProvider, it can be used to provide a module using its type.
func TypeProvider(module interface{}) *ModuleProvider {
	typ := ensureModuleType(module)
	return &ModuleProvider{mod: module, key: typeKey(typ), modType: typ, byType: true}
}

// IntfProvider creates a ModuleProvider, it can be used to provide a module using given interface pointer type.
func IntfProvider(interfacePtr interface{}, moduleImpl interface{}) *ModuleProvider {
	intfType := ensureInterfacePtr(interfacePtr)                   // interface type
	modType := ensureModuleTypeWithInterface(moduleImpl, intfType) // module type
	return &ModuleProvider{mod: moduleImpl, key: typeKey(intfType), modType: modType, byIntf: true}
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
func (m *ModuleContainer) AutoProvide(providers ...*ModuleProvider) error {
	return coreAutoProvide(m, providers)
}

// MustAutoProvide processes with given ModuleProvider-s, injects them if necessary and provides them in dependency order, panics when error happens.
// See AutoProvide for more details.
func (m *ModuleContainer) MustAutoProvide(providers ...*ModuleProvider) {
	err := coreAutoProvide(m, providers)
	if err != nil {
		panic(err.Error())
	}
}

// coreAutoProvide is the core implementation for AutoProvide and MustAutoProvide.
func coreAutoProvide(mc *ModuleContainer, providers []*ModuleProvider) error {
	indeps, depGraph := analyseDependency(providers)
	mc.mu.Lock() // modifying container is safe
	defer mc.mu.Unlock()

	// 1. independent providers
	if len(indeps) > 0 {
		for _, p := range indeps {
			_ = injectAndProvide(mc, p, false) // never error
		}
	}

	// 2. dependent providers
	for len(depGraph) > 0 {
		// providable module: no dependent in graph || all dependents have been provided
		providable, err := findProvidableModules(mc, depGraph)
		if err != nil {
			return err
		}
		for _, p := range providable {
			err := injectAndProvide(mc, p, true)
			if err != nil {
				return err
			}
			// just delete the current module provider from graph
			delete(depGraph, p.key)
		}
	}

	return nil
}

// analyseDependency checks given ModuleProvider-s, analyses and splits them into independent providers and module dependency graph.
func analyseDependency(providers []*ModuleProvider) (indeps []*ModuleProvider, depGraph map[mkey]*ModuleProvider) {
	indeps = make([]*ModuleProvider, 0)
	depGraph = make(map[mkey]*ModuleProvider, 0)

	for _, p := range providers {
		// check current provider
		if p == nil {
			continue
		}
		if p.mod == nil || (p.key.name == "" && p.key.typ == nil) {
			panic(panicInvalidProvider)
		}

		// extract module's dependents
		if p.modType.Kind() == reflect.Ptr && p.modType.Elem().Kind() == reflect.Struct {
			typ := p.modType.Elem()
			for idx := 0; idx < typ.NumField(); idx++ {
				sf := typ.Field(idx)
				moduleTag := sf.Tag.Get(_moduleTagName)
				if moduleTag == "" || moduleTag == "-" {
					continue
				}
				key := mkeyFromField(moduleTag, sf.Type) // by name (...), type or intf (~)
				p.depKeys = append(p.depKeys, key)
			}
		}

		// add current provider to result
		if len(p.depKeys) == 0 {
			indeps = append(indeps, p) // module that have no dependent
		} else {
			depGraph[p.key] = p // add module to dependency graph
		}
	}
	return indeps, depGraph
}

// findProvidableModules finds providable modules in current dependency graph, returns error if cycle dependency happens, or some dependents not found.
func findProvidableModules(mc *ModuleContainer, graph map[mkey]*ModuleProvider) ([]*ModuleProvider, error) {
	// find nodes which may be able to be provided
	providable := make([]*ModuleProvider, 0)
	for _, p := range graph {
		out := false
		for _, depKey := range p.depKeys {
			if _, ok := graph[depKey]; ok {
				out = true
				break
			}
		}
		if !out {
			// node that has zero out degree -> no dependent in graph currently
			providable = append(providable, p)
		}
	}
	if len(graph) > 0 && len(providable) == 0 {
		return nil, errors.New(errModulesCycleDependency)
	}

	// check existence of these nodes all dependents
	for _, p := range providable {
		for _, depKey := range p.depKeys { // these dependents must not be in graph currently
			if _, ok := mc.modules[depKey]; !ok {
				return nil, fmt.Errorf(errRequiredModuleNotFound, depKey.String(), p.modType.String())
			}
		}
	}
	return providable, nil
}

// injectAndProvide injects to given ModuleProvider only for struct pointer type, and provides it to ModuleContainer for all types.
func injectAndProvide(mc *ModuleContainer, p *ModuleProvider, needInject bool) error {
	// check module injectable and do inject
	if needInject {
		val := reflect.ValueOf(p.mod)
		typ := val.Type()
		if typ.Kind() == reflect.Ptr {
			injecteeName := typ.String()
			typ, val = typ.Elem(), val.Elem()
			if typ.Kind() == reflect.Struct {
				// inject something to module first
				errs := injectToStructFields(mc, typ, val, injecteeName, false, false)
				if len(errs) > 0 {
					return xerror.Combine(errs...)
				}
			}
		}
	}

	// provide module after injecting
	mc.modules[p.key] = p.mod
	switch {
	case p.byName:
		mc.logger.PrvName(p.key.name.String(), p.modType.String())
	case p.byType:
		mc.logger.PrvType(p.key.typ.String())
	case p.byIntf:
		mc.logger.PrvIntf(p.key.typ.String(), p.modType.String())
	}
	return nil
}
