# xlinkedhashmap

### References

+ xslice
+ xreflect

### Functions

+ `type LinkedHashMap struct {}`
+ `NewLinkedHashMap() *LinkedHashMap`
+ `(l *LinkedHashMap) Set(key string, value interface{})`
+ `(l *LinkedHashMap) Get(key string) (value interface{}, exist bool)`
+ `(l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) (value interface{})`
+ `(l *LinkedHashMap) Remove(key string) (value interface{}, exist bool)`
+ `(l *LinkedHashMap) Clear()`
+ `(l *LinkedHashMap) MarshalJSON() ([]byte, error)`
+ `(l *LinkedHashMap) String() string`
+ `ObjectToLinkedHashMap(object interface{}) *LinkedHashMap`
+ `FromInterface(object interface{}) *LinkedHashMap`
