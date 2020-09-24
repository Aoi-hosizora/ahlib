# xlinkedhashmap

### References

+ xreflect
+ xtesting*

### Functions

+ `type LinkedHashMap struct {}`
+ `New() *LinkedHashMap`
+ `(l *LinkedHashMap) Keys() []string`
+ `(l *LinkedHashMap) Values() []interface{}`
+ `(l *LinkedHashMap) Len() int`
+ `(l *LinkedHashMap) Set(key string, value interface{})`
+ `(l *LinkedHashMap) Has(key string) bool`
+ `(l *LinkedHashMap) Get(key string) (interface{}, bool)`
+ `(l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) interface{}`
+ `(l *LinkedHashMap) GetForce(key string) interface{}`
+ `(l *LinkedHashMap) Remove(key string) (interface{}, bool)`
+ `(l *LinkedHashMap) Clear()`
+ `(l *LinkedHashMap) MarshalJSON() ([]byte, error)`
+ `(l *LinkedHashMap) MarshalYAML() (interface{}, error)`
+ `(l *LinkedHashMap) String() string`
+ `FromInterface(object interface{}) *LinkedHashMap`
