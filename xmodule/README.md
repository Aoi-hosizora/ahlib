# xmodule

## References

+ xcolor
+ xtesting*

## Documents

### Types

+ `type ServiceName string`
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
+ `func ProvideName(name ServiceName, service interface{})`
+ `func ProvideType(service interface{})`
+ `func ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `func GetByName(name ServiceName) (service interface{}, exist bool)`
+ `func MustGetByName(name ServiceName) interface{}`
+ `func GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `func MustGetByType(serviceType interface{}) interface{}`
+ `func GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `func MustGetByImpl(interfacePtr interface{}) interface{}`
+ `func Inject(ctrl interface{}) (allInjected bool)`
+ `func MustInject(ctrl interface{})`
+ `func DefaultLogger(level LogLevel) Logger`

### Methods

+ `func (s ServiceName) String() string`
+ `func (m *ModuleContainer) SetLogger(logger Logger)`
+ `func (m *ModuleContainer) ProvideName(name ServiceName, service interface{})`
+ `func (m *ModuleContainer) ProvideType(service interface{})`
+ `func (m *ModuleContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `func (m *ModuleContainer) GetByName(name ServiceName) (service interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByName(name ServiceName) interface{}`
+ `func (m *ModuleContainer) GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByType(serviceType interface{}) interface{}`
+ `func (m *ModuleContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `func (m *ModuleContainer) MustGetByImpl(interfacePtr interface{}) interface{}`
+ `func (m *ModuleContainer) Inject(ctrl interface{}) (allInjected bool)`
+ `func (m *ModuleContainer) MustInject(ctrl interface{})`
+ `func (d *defaultLogger) LogName(name, typ string)`
+ `func (d *defaultLogger) LogType(typ string)`
+ `func (d *defaultLogger) LogImpl(itfTyp, srvTyp string)`
+ `func (d *defaultLogger) LogInject(parentTyp, fieldTyp, fieldName string)`
