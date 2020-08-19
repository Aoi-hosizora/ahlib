# xproperty

### References

+ None

### Functions

+ `type PropertyMappers struct {}`
+ `type PropertyMapper struct {}`
+ `type PropertyDict map[string]*PropertyMapperValue`
+ `type PropertyMapperValue struct {}`
+ `New() *PropertyMappers`
+ `NewMapper(from interface{}, to interface{}, dict map[string]*PropertyMapperValue) *PropertyMapper`
+ `NewValue(revert bool, destProps ...string) *PropertyMapperValue`
+ `(p *PropertyMappers) AddMapper(mapper *PropertyMapper)`
+ `(p *PropertyMappers) AddMappers(mappers ...*PropertyMapper)`
+ `(p *PropertyMappers) GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `(p *PropertyMappers) GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper`
+ `AddMapper(mapper *PropertyMapper)`
+ `AddMappers(mappers ...*PropertyMapper)`
+ `GetMapper(from interface{}, to interface{}) (*PropertyMapper, error)`
+ `GetDefaultMapper(from interface{}, to interface{}) *PropertyMapper`
+ `(p PropertyDict) ApplyOrderBy(source string) string`
+ `(p *PropertyMapper) ApplyOrderBy(source string) string`
+ `(p PropertyDict) ApplyCypherOrderBy(parent, source string) string`
+ `(p *PropertyMapper) ApplyCypherOrderBy(parent, source string) string`
