# xdi

## References

+ xcolor
+ xreflect
+ xtesting*

## Documents

### Types

+ `type ServiceName string`
+ `type DiContainer struct`
+ `type LogLevel int8`
+ `type Logger interface`

### Constants

+ `const LogName LogLevel`
+ `const LogType LogLevel`
+ `const LogImpl LogLevel`
+ `const LogInject LogLevel`
+ `const LogAll LogLevel`
+ `const LogSilent LogLevel`

### Functions

+ `func NewDiContainer() *DiContainer`
+ `func SetLogger(logger Logger)`
+ `func ProvideName(name ServiceName, service interface{})`
+ `func ProvideType(service interface{})`
+ `func ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `func GetByName(name ServiceName) (service interface{}, exist bool)`
+ `func GetByNameForce(name ServiceName) interface{}`
+ `func GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `func GetByTypeForce(serviceType interface{}) interface{}`
+ `func GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `func GetByImplForce(interfacePtr interface{}) interface{}`
+ `func Inject(ctrl interface{}) (allInjected bool)`
+ `func MustInject(ctrl interface{})`
+ `func DefaultLogger(level LogLevel) Logger`

### Methods

+ `func (s ServiceName) String() string`
+ `func (d *DiContainer) SetLogger(logger Logger)`
+ `func (d *DiContainer) ProvideName(name ServiceName, service interface{})`
+ `func (d *DiContainer) ProvideType(service interface{})`
+ `func (d *DiContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `func (d *DiContainer) GetByName(name ServiceName) (service interface{}, exist bool)`
+ `func (d *DiContainer) GetByNameForce(name ServiceName) interface{}`
+ `func (d *DiContainer) GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `func (d *DiContainer) GetByTypeForce(serviceType interface{}) interface{}`
+ `func (d *DiContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `func (d *DiContainer) GetByImplForce(interfacePtr interface{}) interface{}`
+ `func (d *DiContainer) Inject(ctrl interface{}) (allInjected bool)`
+ `func (d *DiContainer) MustInject(ctrl interface{})`
+ `func (d *defaultLogger) LogName(name, typ string)`
+ `func (d *defaultLogger) LogType(typ string)`
+ `func (d *defaultLogger) LogImpl(itfTyp, srvTyp string)`
+ `func (d *defaultLogger) LogInject(parentTyp, fieldTyp, fieldName string)`
