# xentity

## References

+ xslice
+ xcondition*
+ xtesting*

## Documents

### Types

+ `type EntityMappers struct`
+ `type EntityMapper struct`
+ `type MapFunc func`

### Constants

+ None

### Functions

+ `func New() *EntityMappers`
+ `func NewMapper(from interface{}, ctor func() interface{}, mapFunc MapFunc) *EntityMapper`
+ `func AddMapper(mapper *EntityMapper)`
+ `func AddMappers(mappers ...*EntityMapper)`
+ `func GetMapper(from interface{}, to interface{}) (*EntityMapper, error)`
+ `func MapProp(from interface{}, to interface{}, options ...MapFunc) error`
+ `func Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `func MapSlice(from interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `func MustMapProp(from interface{}, to interface{}, options ...MapFunc)`
+ `func MustMap(from interface{}, to interface{}, options ...MapFunc) interface{}`
+ `func MustMapSlice(from interface{}, to interface{}, options ...MapFunc) interface{}`

### Methods

+ `func (e *EntityMapper) GetMapFunc() MapFunc`
+ `func (e *EntityMappers) AddMapper(m *EntityMapper)`
+ `func (e *EntityMappers) AddMappers(ms ...*EntityMapper)`
+ `func (e *EntityMappers) GetMapper(from interface{}, to interface{}) (*EntityMapper, error)`
+ `func (e *EntityMappers) MapProp(from interface{}, to interface{}, options ...MapFunc) error`
+ `func (e *EntityMappers) Map(from interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `func (e *EntityMappers) MapSlice(from interface{}, to interface{}, options ...MapFunc) (interface{}, error)`
+ `func (e *EntityMappers) MustMapProp(from interface{}, to interface{}, options ...MapFunc)`
+ `func (e *EntityMappers) MustMap(from interface{}, to interface{}, options ...MapFunc) interface{}`
+ `func (e *EntityMappers) MustMapSlice(from interface{}, to interface{}, options ...MapFunc) interface{}`
