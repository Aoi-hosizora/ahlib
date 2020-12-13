# xproperty

## References

+ xtesting*

## Documents

### Types

+ `type PropertyMappers struct`
+ `type PropertyMapper struct`
+ `type PropertyDict map[string]*PropertyMapperValue`
+ `type PropertyMapperValue struct`

### Variables

+ None

### Constants

+ None

### Functions

+ `func New() *PropertyMappers`
+ `func NewMapper(src interface{}, dest interface{}, dict PropertyDict) *PropertyMapper`
+ `func NewValue(revert bool, destinations ...string) *PropertyMapperValue`
+ `func NewValueCompletely(id int, revert bool, arg interface{}, destinations ...string) *PropertyMapperValue`
+ `func AddMapper(mapper *PropertyMapper)`
+ `func AddMappers(mappers ...*PropertyMapper)`
+ `func GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error)`
+ `func GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper`

### Methods

+ `func (p *PropertyMapperValue) GetId() int`
+ `func (p *PropertyMapperValue) GetRevert() bool`
+ `func (p *PropertyMapperValue) GetArg() interface{}`
+ `func (p *PropertyMapperValue) GetDestinations() []string`
+ `func (p *PropertyMapper) GetDict() PropertyDict`
+ `func (p *PropertyMappers) AddMapper(mapper *PropertyMapper)`
+ `func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper)`
+ `func (p *PropertyMappers) GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error)`
+ `func (p *PropertyMappers) GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper`
