# xmodule

## Dependencies

+ xtesting*
+ xcolor

## Documents

### Types

+ `type ModuleName string`
+ `type ModuleContainer struct`
+ `type LogLevel uint8`
+ `type Logger interface`

### Variables

+ None

### Constants

+ `const LogPrvName LogLevel`
+ `const LogPrvType LogLevel`
+ `const LogPrvIntf LogLevel`
+ `const LogInjField LogLevel`
+ `const LogInjFinish LogLevel`
+ `const LogAll LogLevel`
+ `const LogSilent LogLevel`

### Functions

+ `func NewModuleContainer() *ModuleContainer`
+ `func SetLogger(logger Logger)`
+ `func ProvideByName(name ModuleName, module interface{})`
+ `func ProvideByType(module interface{})`
+ `func ProvideByIntf(interfacePtr interface{}, moduleImpl interface{})`
+ `func RemoveByName(name ModuleName) (removed bool)`
+ `func RemoveByType(module interface{}) (removed bool)`
+ `func RemoveByIntf(interfacePtr interface{}) (removed bool)`
+ `func GetByName(name ModuleName) (module interface{}, exist bool)`
+ `func MustGetByName(name ModuleName) interface{}`
+ `func GetByType(moduleType interface{}) (module interface{}, exist bool)`
+ `func MustGetByType(moduleType interface{}) interface{}`
+ `func GetByIntf(interfacePtr interface{}) (module interface{}, exist bool)`
+ `func MustGetByIntf(interfacePtr interface{}) interface{}`
+ `func Inject(injectee interface{}) (allInjected bool)`
+ `func MustInject(injectee interface{})`
+ `func DefaultLogger(level LogLevel, logPrvFunc func(moduleName, moduleType string), logInjFunc func(moduleName, injecteeType, addition string)) Logger`

### Methods

+ `func (m ModuleName) String() string`
+ `func (m *ModuleContainer) SetLogger(logger Logger)`
+ `func (m *ModuleContainer) ProvideByName(name ModuleName, module interface{})`
+ `func (m *ModuleContainer) ProvideByType(module interface{})`
+ `func (m *ModuleContainer) ProvideByIntf(interfacePtr interface{}, moduleImpl interface{})`
+ `func (m *ModuleContainer) RemoveByName(name ModuleName) (removed bool)`
+ `func (m *ModuleContainer) RemoveByType(moduleType interface{}) (removed bool)`
+ `func (m *ModuleContainer) RemoveByIntf(interfacePtr interface{}) (removed bool)`
+ `func (m *ModuleContainer) GetByName(name ModuleName) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByName(name ModuleName) interface{}`
+ `func (m *ModuleContainer) GetByType(moduleType interface{}) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByType(moduleType interface{}) interface{}`
+ `func (m *ModuleContainer) GetByIntf(interfacePtr interface{}) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByIntf(interfacePtr interface{}) interface{}`
+ `func (m *ModuleContainer) Inject(injectee interface{}) (allInjected bool)`
+ `func (m *ModuleContainer) MustInject(injectee interface{})`
