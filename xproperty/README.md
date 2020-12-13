# xproperty

## References

+ xtesting*

## Documents

### Types

+ `type PropertyMappers struct`
+ `type PropertyMapper struct`
+ `type PropertyDict map[string]*PropertyMapperValue`
+ `type VariableDict map[string]int`
+ `type PropertyMapperValue struct`

### Constants

+ None

### Functions

+ `func New() *PropertyMappers`
+ `func NewMapper(src interface{}, dest interface{}, dict PropertyDict) *PropertyMapper`
+ `func NewValue(revert bool, destinations ...string) *PropertyMapperValue`
+ `func AddMapper(mapper *PropertyMapper)`
+ `func AddMappers(mappers ...*PropertyMapper)`
+ `func GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error)`
+ `func GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper`

### Methods

+ `func (p *PropertyMapper) GetDict() PropertyDict`
+ `func (p *PropertyMappers) AddMapper(mapper *PropertyMapper)`
+ `func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper)`
+ `func (p *PropertyMappers) GetMapper(src interface{}, dest interface{}) (*PropertyMapper, error)`
+ `func (p *PropertyMappers) GetDefaultMapper(src interface{}, dest interface{}) *PropertyMapper`
