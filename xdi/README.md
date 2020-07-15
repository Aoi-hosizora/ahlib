# xdi

## Functions

### For DiContainer

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

### For Global

+ `LogFunc`
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
