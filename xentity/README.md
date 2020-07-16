# xentity

### Functions

+ `type EntityMappers struct {}`
+ `type MapFunc func(from interface{}, to interface{}) error`
+ `type EntityMapper struct {}`
+ `NewEntityMappers() *EntityMappers`
+ `NewEntityMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper`
+ `(e *EntityMappers) AddMapper(mapper *EntityMapper)`
+ `(e *EntityMappers) AddMappers(mappers ...*EntityMapper)`
+ `(e *EntityMappers) GetMapper(from interface{}, to interface{}) (*EntityMapper, error)`
+ `(e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error`
+ `(e *EntityMappers) Map(from interface{}, toModel interface{}, options ...MapFunc) (interface{}, error)`
+ `(e *EntityMappers) MapSlice(from []interface{}, toModel interface{}, options ...MapFunc) (interface{}, error)`
+ `(e *EntityMappers) MustMapProp(from interface{}, to interface{}, options ...MapFunc)`
+ `(e *EntityMappers) MustMap(from interface{}, toModel interface{}, options ...MapFunc) interface{}`
+ `(e *EntityMappers) MustMapSlice(from []interface{}, toModel interface{}, options ...MapFunc) interface{}`
+ `AddMapper(mapper *EntityMapper)`
+ `AddMappers(mappers ...*EntityMapper)`
+ `GetMapper(from interface{}, to interface{}) (*EntityMapper, error)`
+ `MapProp(from interface{}, to interface{}, options ...MapFunc) error`
+ `Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `MapSlice(from []interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `MustMapProp(from interface{}, to interface{}, options ...MapFunc)`
+ `MustMap(from interface{}, to interface{}, options ...MapFunc) interface{}`
+ `MustMapSlice(from []interface{}, to interface{}, options ...MapFunc) interface{}`
