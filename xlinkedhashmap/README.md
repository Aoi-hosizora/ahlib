# xlinkedhashmap

## References

+ xreflect
+ xtesting*

## Documents

### Types

+ `type LinkedHashMap struct`

### Constants

+ None

### Functions

+ `func New() *LinkedHashMap`
+ `FromInterface(object interface{}) *LinkedHashMap`

### Methods

+ `func (l *LinkedHashMap) Keys() []string`
+ `func (l *LinkedHashMap) Values() []interface{}`
+ `func (l *LinkedHashMap) Len() int`
+ `func (l *LinkedHashMap) Set(key string, value interface{})`
+ `func (l *LinkedHashMap) Has(key string) bool`
+ `func (l *LinkedHashMap) Get(key string) (interface{}, bool)`
+ `func (l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) interface{}`
+ `func (l *LinkedHashMap) GetForce(key string) interface{}`
+ `func (l *LinkedHashMap) Remove(key string) (interface{}, bool)`
+ `func (l *LinkedHashMap) Clear()`
+ `func (l *LinkedHashMap) MarshalJSON() ([]byte, error)`
+ `func (l *LinkedHashMap) MarshalYAML() (interface{}, error)`
+ `func (l *LinkedHashMap) String() string`
