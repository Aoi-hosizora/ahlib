# xorderedmap

## Dependencies

+ xreflect
+ xtesting*

## Documents

### Types

+ `type OrderedMap struct`

### Variables

+ None

### Constants

+ None

### Functions

+ `func New() *OrderedMap`
+ `func FromInterface(object interface{}) *OrderedMap`

### Methods

+ `func (l *OrderedMap) Keys() []string`
+ `func (l *OrderedMap) Values() []interface{}`
+ `func (l *OrderedMap) Len() int`
+ `func (l *OrderedMap) Set(key string, value interface{})`
+ `func (l *OrderedMap) Has(key string) bool`
+ `func (l *OrderedMap) Get(key string) (interface{}, bool)`
+ `func (l *OrderedMap) GetOr(key string, defaultValue interface{}) interface{}`
+ `func (l *OrderedMap) MustGet(key string) interface{}`
+ `func (l *OrderedMap) Remove(key string) (interface{}, bool)`
+ `func (l *OrderedMap) Clear()`
+ `func (l *OrderedMap) MarshalJSON() ([]byte, error)`
+ `func (l *OrderedMap) MarshalYAML() (interface{}, error)`
+ `func (l *OrderedMap) String() string`
