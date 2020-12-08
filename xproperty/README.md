# xproperty

## References

+ xtesting*

## Documents

### Types

+ `type PropertyMappers struct`
+ `type PropertyMapper struct`
+ `type PropertyMapperValue struct`
+ `type PropertyDict map[string]*PropertyMapperValue`
+ `type VariableDict map[string]int`

### Constants

+ None

### Functions

+ `func New() *PropertyMappers`
+ `func NewMapper(from interface{}, to interface{}, dict PropertyDict) *PropertyMapper`
+ `func NewValue(revert bool, destinations ...string) *PropertyMapperValue`
+ `func AddMapper(mapper *PropertyMapper)`
+ `func AddMappers(mappers ...*PropertyMapper)`
+ `func GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `func GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper`

### Methods

+ `func (p *PropertyMapper) GetDict() PropertyDict`
+ `func (p *PropertyMappers) AddMapper(m *PropertyMapper)`
+ `func (p *PropertyMappers) AddMappers(mappers ...*PropertyMapper)`
+ `func (p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `func (p *PropertyMappers) GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper`
