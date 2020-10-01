# xdi

### References

+ xcolor
+ xreflect
+ xtesting*

### Functions

+ `type ServiceName string`
+ `(s *ServiceName) String()`
+ `type DiContainer struct {}`
+ `NewDiContainer() *DiContainer`
+ `(d *DiContainer) SetLogger(logger Logger)`
+ `(d *DiContainer) ProvideName(name ServiceName, service interface{})`
+ `(d *DiContainer) ProvideType(service interface{})`
+ `(d *DiContainer) ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `(d *DiContainer) GetByName(name ServiceName) (service interface{}, exist bool)`
+ `(d *DiContainer) GetByNameForce(name ServiceName) interface{}`
+ `(d *DiContainer) GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `(d *DiContainer) GetByTypeForce(serviceType interface{}) interface{}`
+ `(d *DiContainer) GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `(d *DiContainer) GetByImplForce(interfacePtr interface{}) interface{}`
+ `(d *DiContainer) Inject(ctrl interface{}) (allInjected bool)`
+ `(d *DiContainer) MustInject(ctrl interface{})`
+ `SetLogger(logger Logger)`
+ `ProvideName(name ServiceName, service interface{})`
+ `ProvideType(service interface{})`
+ `ProvideImpl(interfacePtr interface{}, serviceImpl interface{})`
+ `GetByName(name ServiceName) (service interface{}, exist bool)`
+ `GetByNameForce(name ServiceName) interface{}`
+ `GetByType(serviceType interface{}) (service interface{}, exist bool)`
+ `GetByTypeForce(serviceType interface{}) interface{}`
+ `GetByImpl(interfacePtr interface{}) (service interface{}, exist bool)`
+ `GetByImplForce(interfacePtr interface{}) interface{}`
+ `Inject(ctrl interface{}) (allInjected bool)`
+ `MustInject(ctrl interface{})`
+ `type Logger interface {}`
+ `DefaultLogger() Logger`
+ `SilentLogger() Logger`
