# xproperty

### Functions

+ `type PropertyMappers struct {}`
+ `type PropertyMapper struct {}`
+ `type PropertyMapperValue struct {}`
+ `NewPropertyMappers() *PropertyMappers`
+ `NewPropertyMapper(from interface{}, to interface{}, dict map[string]*PropertyMapperValue) *PropertyMapper`
+ `NewPropertyMapperValue(destProps []string, revert bool) *PropertyMapperValue`
+ `(p *PropertyMappers) AddMapper(mapper *PropertyMapper)`
+ `(p *PropertyMappers) AddMappers(mappers ...*PropertyMapper)`
+ `(p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `(p *PropertyMappers) GetMapperDefault(from interface{}, to interface{}) *PropertyMapper`
+ `AddMapper(mapper *PropertyMapper)`
+ `AddMappers(mappers ...*PropertyMapper)`
+ `GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `GetMapperDefault(from interface{}, to interface{}) *PropertyMapper`

### Extension Functions

+ `(p *PropertyMapper) ApplyOrderBy(source string) string`
