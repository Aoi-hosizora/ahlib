# xslice

### Common Functions

+ `Shuffle(slice []interface{}, source rand.Source)`
+ `Reverse(slice []interface{}) []interface{}`
+ `Map(slice []interface{}, mapFunc func(interface{}) interface{}) []interface{}`
+ `IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int`
+ `IndexOf(slice []interface{}, value interface{}) int`
+ `ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool`
+ `Contains(slice []interface{}, value interface{}) bool`
+ `DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{}`
+ `Delete(slice []interface{}, value interface{}, n int) []interface{}`
+ `DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{}`
+ `DeleteAll(slice []interface{}, value interface{}) []interface{}`
+ `DiffWith(s1 []interface{}, s2 []interface{}, equaller Equaller) []interface{}`
+ `Diff(s1 []interface{}, s2 []interface{}) []interface{}`
+ `EqualWith(s1 []interface{}, s2 []interface{}, equaller Equaller) bool`
+ `Equal(s1 []interface{}, s2 []interface{}) bool`

### Helper Functions

+ `Sti(slice interface{}) []interface{}`
+ `Its(slice []interface{}, model interface{}) interface{}`
+ `ItsToString(slice []interface{}) []string`
+ `ItsOfString(slice []interface{}) []string`
+ `ItsOfByte(slice []interface{}) []byte`
+ `ItsOfRune(slice []interface{}) []rune`
+ `ItsOfInt(slice []interface{}) []int`
+ `ItsOfUint(slice []interface{}) []uint`
+ `ItsOfInt8(slice []interface{}) []int8`
+ `ItsOfUint8(slice []interface{}) []uint8`
+ `ItsOfInt16(slice []interface{}) []int16`
+ `ItsOfUint16(slice []interface{}) []uint16`
+ `ItsOfInt32(slice []interface{}) []int32`
+ `ItsOfUint32(slice []interface{}) []uint32`
+ `ItsOfInt64(slice []interface{}) []int64`
+ `ItsOfUint64(slice []interface{}) []uint64`
+ `ItsOfFloat32(slice []interface{}) []float32`
+ `ItsOfFloat64(slice []interface{}) []float64`
