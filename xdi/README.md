# xdi

### References

+ xcolor
+ xreflect

### Functions

#### For DiContainer

+ `(d *DiContainer) SetLogMode(provideLog bool, injectLog bool)`
+ `(d *DiContainer) SetLogFunc(logFunc LogFunc)`
+ `(d *DiContainer) ProvideType(service interface{})`
+ `(d *DiContainer) ProvideImpl(itfNilPtr interface{}, impl interface{})`
+ `(d *DiContainer) ProvideName(name string, service interface{})`
+ `(d *DiContainer) GetByType(srvType interface{}) (service interface{}, exist bool)`
+ `(d *DiContainer) GetByName(name string) (service interface{}, exist bool)`
+ `(d *DiContainer) GetByTypeForce(srvType interface{}) interface{}`
+ `(d *DiContainer) GetByNameForce(name string) interface{}`
+ `(d *DiContainer) Inject(ctrl interface{}) (allInjected bool)`
+ `(d *DiContainer) MustInject(ctrl interface{})`

#### For Global

+ `type LogFunc func(kind, parentType, fieldName, fieldType string)`
+ `DefaultLogFunc() LogFunc`
+ `type ServiceName string`
+ `(s *ServiceName) String() string`
+ `NewDiContainer() *DiContainer`
+ `SetLogMode(provideLog bool, injectLog bool)`
+ `SetLogFunc(logFunc LogFunc)`
+ `ProvideType(service interface{})`
+ `ProvideImpl(itfNilPtr interface{}, impl interface{})`
+ `ProvideName(name string, service interface{})`
+ `GetByType(srvType interface{}) (service interface{}, exist bool)`
+ `GetByName(name string) (service interface{}, exist bool)`
+ `GetByTypeForce(srvType interface{}) interface{}`
+ `GetByNameForce(name string) interface{}`
+ `Inject(ctrl interface{}) (allInjected bool)`
+ `MustInject(ctrl interface{})`
