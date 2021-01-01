# xentity

## Dependencies

+ xtesting*

## Documents

### Types

+ `type EntityMappers struct`
+ `type EntityMapper struct`
+ `type MapFunc func`

### Variables

+ None

### Constants

+ None

### Functions

+ `func New() *EntityMappers`
+ `NewMapper(src interface{}, destCtor func() interface{}, mapFunc MapFunc) *EntityMapper`
+ `func AddMapper(mapper *EntityMapper)`
+ `func AddMappers(mappers ...*EntityMapper)`
+ `func GetMapper(src interface{}, dest interface{}) (*EntityMapper, error)`
+ `func MapProp(src interface{}, dest interface{}, opts ...MapFunc) error`
+ `func Map(src interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error)`
+ `func MapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error)`
+ `func MustMapProp(src interface{}, dest interface{}, opts ...MapFunc)`
+ `func MustMap(src interface{}, destModel interface{}, opts ...MapFunc) interface{}`
+ `func MustMapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) interface{}`

### Methods

+ `func (e *EntityMapper) GetMapFunc() MapFunc`
+ `func (e *EntityMappers) AddMapper(mapper *EntityMapper)`
+ `func (e *EntityMappers) AddMappers(mappers ...*EntityMapper)`
+ `func (e *EntityMappers) GetMapper(src interface{}, dest interface{}) (*EntityMapper, error)`
+ `func (e *EntityMappers) MapProp(src interface{}, dest interface{}, opts ...MapFunc) error`
+ `func (e *EntityMappers) Map(src interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error)`
+ `func (e *EntityMappers) MapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) (interface{}, error)`
+ `func (e *EntityMappers) MustMapProp(src interface{}, dest interface{}, opts ...MapFunc)`
+ `func (e *EntityMappers) MustMap(src interface{}, destModel interface{}, opts ...MapFunc) interface{}`
+ `func (e *EntityMappers) MustMapSlice(srcSlice interface{}, destModel interface{}, opts ...MapFunc) interface{}`
