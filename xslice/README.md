# xslice

### Common Functions

+ `Shuffle(slice []interface{}, source rand.Source)`
+ `Reverse(slice []interface{}) []interface{}`
+ `IndexOf(slice []interface{}, value interface{}) (index int)`
+ `Contains(slice []interface{}, value interface{}) bool`
+ `Map(slice []interface{}, mapFunc func(interface{}) interface{}) []interface{}`
+ `Delete(slice []interface{}, value interface{}, n int) []interface{}`
+ `DeleteAll(slice []interface{}, value interface{}) []interface{}`
+ `SliceDiff(s1 []interface{}, s2 []interface{}) []interface{}`
+ `Equal(s1 []interface{}, s2 []interface{}) bool`

### Helper Functions

+ `Sti(slice interface{}) []interface{}`
+ `Its(slice []interface{}, model interface{}) interface{}`
+ `ItsToString(slice []interface{}) []string`
+ `ItsOfInt(slice []interface{}) []int`
+ `ItsOfUint(slice []interface{}) []uint`
+ `ItsOfUint32(slice []interface{}) []uint32`
+ `ItsOfUint64(slice []interface{}) []uint64`
+ `ItsOfInt32(slice []interface{}) []int32`
+ `ItsOfInt64(slice []interface{}) []int64`
+ `ItsOfFloat32(slice []interface{}) []float32`
+ `ItsOfFloat64(slice []interface{}) []float64`
+ `ItsOfComplex64(slice []interface{}) []complex64`
+ `ItsOfComplex128(slice []interface{}) []complex128`
+ `ItsOfString(slice []interface{}) []string`
+ `ItsOfByte(slice []interface{}) []byte`
+ `ItsOfRune(slice []interface{}) []rune`
