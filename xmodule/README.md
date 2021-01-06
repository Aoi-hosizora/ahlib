# xmodule

## Dependencies

+ xcolor
+ xtesting*

## Documents

### Types

+ `type ModuleName string`
+ `type ModuleContainer struct`
+ `type LogLevel int8`
+ `type Logger interface`

### Variables

+ `var LogLeftArrow func(arg1, arg2, arg3 string)`
+ `var LogRightArrow func(arg1, arg2, arg3 string)`

### Constants

+ `const LogName LogLevel`
+ `const LogType LogLevel`
+ `const LogImpl LogLevel`
+ `const LogInject LogLevel`
+ `const LogAll LogLevel`
+ `const LogSilent LogLevel`

### Functions

+ `func NewModuleContainer() *ModuleContainer`
+ `func SetLogger(logger Logger)`
+ `func ProvideName(name ModuleName, module interface{})`
+ `func ProvideType(module interface{})`
+ `func ProvideImpl(interfacePtr interface{}, moduleImpl interface{})`
+ `func GetByName(name ModuleName) (module interface{}, exist bool)`
+ `func MustGetByName(name ModuleName) interface{}`
+ `func GetByType(moduleType interface{}) (module interface{}, exist bool)`
+ `func MustGetByType(moduleType interface{}) interface{}`
+ `func GetByImpl(interfacePtr interface{}) (module interface{}, exist bool)`
+ `func MustGetByImpl(interfacePtr interface{}) interface{}`
+ `func Inject(ctrl interface{}) (allInjected bool)`
+ `func MustInject(ctrl interface{})`
+ `func DefaultLogger(level LogLevel) Logger`

### Methods

+ `func (s ModuleName) String() string`
+ `func (m *ModuleContainer) SetLogger(logger Logger)`
+ `func (m *ModuleContainer) ProvideName(name ModuleName, module interface{})`
+ `func (m *ModuleContainer) ProvideType(module interface{})`
+ `func (m *ModuleContainer) ProvideImpl(interfacePtr interface{}, moduleImpl interface{})`
+ `func (m *ModuleContainer) GetByName(name ModuleName) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByName(name ModuleName) interface{}`
+ `func (m *ModuleContainer) GetByType(moduleType interface{}) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByType(moduleType interface{}) interface{}`
+ `func (m *ModuleContainer) GetByImpl(interfacePtr interface{}) (module interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByImpl(interfacePtr interface{}) interface{}`
+ `func (m *ModuleContainer) Inject(ctrl interface{}) (allInjected bool)`
+ `func (m *ModuleContainer) MustInject(ctrl interface{})`
+ `func (d *defaultLogger) LogName(name, typ string)`
+ `func (d *defaultLogger) LogType(typ string)`
+ `func (d *defaultLogger) LogImpl(interfaceTyp, moduleTyp string)`
+ `func (d *defaultLogger) LogInject(parentTyp, fieldTyp, fieldName string)`
