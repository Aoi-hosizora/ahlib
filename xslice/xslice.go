package xslice

import (
	"math/rand"
	"time"
)

// Equaller represents how two data equal, used for XXXWith methods.
type Equaller func(i, j interface{}) bool

// defaultEqualler represents a default Equaller, it just checks equality directly.
var defaultEqualler Equaller = func(i, j interface{}) bool {
	return i == j
}

func ShuffleSelf(slice []interface{}) {
	coreShuffle(checkSliceParam(slice))
}

func Shuffle(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreShuffle(checkSliceParam(newSlice))
	return newSlice
}

func ShuffleSelfG(slice interface{}) {
	coreShuffle(checkInterfaceParam(slice))
}

func ShuffleG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreShuffle(checkInterfaceParam(newSlice))
	return newSlice
}

func coreShuffle(slice innerSlice) {
	rand.Seed(time.Now().UnixNano())
	for i := slice.length() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		itemJ := slice.get(j)
		itemI := slice.get(i)
		slice.set(i, itemJ)
		slice.set(j, itemI)
	}
}

func ReverseSelf(slice []interface{}) {
	coreReverse(checkSliceParam(slice))
}

func Reverse(slice []interface{}) []interface{} {
	newSlice := cloneInterfaceSlice(slice)
	coreReverse(checkSliceParam(newSlice))
	return newSlice
}

func ReverseSelfG(slice interface{}) {
	coreReverse(checkInterfaceParam(slice))
}

func ReverseG(slice interface{}) interface{} {
	newSlice := cloneSliceInterface(slice)
	coreReverse(checkInterfaceParam(newSlice))
	return newSlice
}

func coreReverse(slice innerSlice) {
	for i, j := 0, slice.length()-1; i < j; i, j = i+1, j-1 {
		itemJ := slice.get(j)
		itemI := slice.get(i)
		slice.set(i, itemJ)
		slice.set(j, itemI)
	}
}

func IndexOf(slice []interface{}, value interface{}) int {
	return coreIndexOf(checkSliceParam(slice), value, defaultEqualler)
}

func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreIndexOf(checkSliceParam(slice), value, equaller)
}

func IndexOfG(slice interface{}, value interface{}) int {
	return coreIndexOf(checkInterfaceParam(slice), value, defaultEqualler)
}

func IndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int {
	return coreIndexOf(checkInterfaceParam(slice), value, equaller)
}

func coreIndexOf(slice innerSlice, value interface{}, equaller Equaller) int {
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(value, val) {
			return idx
		}
	}
	return -1
}

func Contains(slice []interface{}, value interface{}) bool {
	return coreContains(checkSliceParam(slice), value, defaultEqualler)
}

func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool {
	return coreContains(checkSliceParam(slice), value, equaller)
}

func ContainsG(slice interface{}, value interface{}) bool {
	return coreContains(checkInterfaceParam(slice), value, defaultEqualler)
}

func ContainsWithG(slice interface{}, value interface{}, equaller Equaller) bool {
	return coreContains(checkInterfaceParam(slice), value, equaller)
}

func coreContains(slice innerSlice, value interface{}, equaller Equaller) bool {
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(value, val) {
			return true
		}
	}
	return false
}

func Count(slice []interface{}, value interface{}) int {
	return coreCount(checkSliceParam(slice), value, defaultEqualler)
}

func CountWith(slice []interface{}, value interface{}, equaller Equaller) int {
	return coreCount(checkSliceParam(slice), value, equaller)
}

func CountG(slice interface{}, value interface{}) int {
	return coreCount(checkInterfaceParam(slice), value, defaultEqualler)
}

func CountWithG(slice interface{}, value interface{}, equaller Equaller) int {
	return coreCount(checkInterfaceParam(slice), value, equaller)
}

func coreCount(slice innerSlice, value interface{}, equaller Equaller) int {
	cnt := 0
	for idx := 0; idx < slice.length(); idx++ {
		val := slice.get(idx)
		if equaller(value, val) {
			cnt++
		}
	}
	return cnt
}

func Delete(slice []interface{}, value interface{}, n int) []interface{} {
	return coreDelete(checkSliceParam(slice), value, n, defaultEqualler).actual().([]interface{})
}

func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{} {
	return coreDelete(checkSliceParam(slice), value, n, equaller).actual().([]interface{})
}

func DeleteG(slice interface{}, value interface{}, n int) interface{} {
	return coreDelete(checkInterfaceParam(slice), value, n, defaultEqualler).actual()
}

func DeleteWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{} {
	return coreDelete(checkInterfaceParam(slice), value, n, equaller).actual()
}

func DeleteAll(slice []interface{}, value interface{}) []interface{} {
	return coreDelete(checkSliceParam(slice), value, 0, defaultEqualler).actual().([]interface{})
}

func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{} {
	return coreDelete(checkSliceParam(slice), value, 0, equaller).actual().([]interface{})
}

func DeleteAllG(slice []interface{}, value interface{}) interface{} {
	return coreDelete(checkInterfaceParam(slice), value, 0, defaultEqualler).actual()
}

func DeleteAllWithG(slice []interface{}, value interface{}, equaller Equaller) interface{} {
	return coreDelete(checkInterfaceParam(slice), value, 0, equaller).actual()
}

func coreDelete(slice innerSlice, value interface{}, n int, equaller Equaller) innerSlice {
	if n <= 0 {
		n = slice.length()
	}
	cnt := 0
	idx := coreIndexOf(slice, value, equaller)
	for idx != -1 && cnt < n {
		slice.remove(idx)
		cnt++
		idx = coreIndexOf(slice, value, equaller)
	}
	return slice
}

func Diff(slice1, slice2 []interface{}) []interface{} {
	return coreDiff(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

func DiffWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreDiff(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

func DiffG(slice1, slice2 interface{}) interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreDiff(newSlice1, newSlice2, defaultEqualler).actual()
}

func DiffWithG(slice1, slice2 interface{}, equaller Equaller) interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreDiff(newSlice1, newSlice2, equaller).actual()
}

func coreDiff(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, 0)
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		exist := false
		for i2 := 0; i2 < slice2.length(); i2++ {
			item2 := slice2.get(i2)
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result.append(item1)
		}
	}
	return result
}

func Union(slice1, slice2 []interface{}) []interface{} {
	return coreUnion(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

func UnionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreUnion(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

func UnionG(slice1, slice2 interface{}) []interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreUnion(newSlice1, newSlice2, defaultEqualler).actual().([]interface{})
}

func UnionWithG(slice1, slice2 interface{}, equaller Equaller) []interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreUnion(newSlice1, newSlice2, equaller).actual().([]interface{})
}

func coreUnion(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, slice1.length())
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		result.append(item1)
	}
	for i2 := 0; i2 < slice2.length(); i2++ {
		item2 := slice2.get(i2)
		exist := false
		for i1 := 0; i1 < slice1.length(); i1++ {
			item1 := slice1.get(i1)
			if equaller(item1, item2) {
				exist = true
				break
			}
		}
		if !exist {
			result.append(item2)
		}
	}
	return result
}

func Intersection(slice1, slice2 []interface{}) []interface{} {
	return coreIntersection(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler).actual().([]interface{})
}

func IntersectionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{} {
	return coreIntersection(checkSliceParam(slice1), checkSliceParam(slice2), equaller).actual().([]interface{})
}

func IntersectionG(slice1, slice2 interface{}) []interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreIntersection(newSlice1, newSlice2, defaultEqualler).actual().([]interface{})
}

func IntersectionWithG(slice1, slice2 interface{}, equaller Equaller) []interface{} {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreIntersection(newSlice1, newSlice2, equaller).actual().([]interface{})
}

func coreIntersection(slice1, slice2 innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice1, 0, 0)
	for i1 := 0; i1 < slice1.length(); i1++ {
		item1 := slice1.get(i1)
		for i2 := 0; i2 < slice2.length(); i2++ {
			item2 := slice2.get(i2)
			if equaller(item1, item2) {
				result.append(item1)
				break
			}
		}
	}
	return result
}

func ToSet(slice []interface{}) []interface{} {
	return coreToSet(checkSliceParam(slice), defaultEqualler).actual().([]interface{})
}

func ToSetWith(slice []interface{}, equaller Equaller) []interface{} {
	return coreToSet(checkSliceParam(slice), equaller).actual().([]interface{})
}

func ToSetG(slice interface{}) []interface{} {
	return coreToSet(checkInterfaceParam(slice), defaultEqualler).actual().([]interface{})
}

func ToSetWithG(slice interface{}, equaller Equaller) []interface{} {
	return coreToSet(checkInterfaceParam(slice), equaller).actual().([]interface{})
}

func coreToSet(slice innerSlice, equaller Equaller) innerSlice {
	result := makeInnerSlice(slice, 0, 0)
	for idx := 0; idx < slice.length(); idx++ {
		item := slice.get(idx)
		if coreCount(result, item, equaller) == 0 {
			result.append(item)
		}
	}
	return result
}

func ElementMatch(slice1, slice2 []interface{}) bool {
	return coreElementMatch(checkSliceParam(slice1), checkSliceParam(slice2), defaultEqualler)
}

func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool {
	return coreElementMatch(checkSliceParam(slice1), checkSliceParam(slice2), equaller)
}

func ElementMatchG(slice1, slice2 interface{}) bool {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreElementMatch(newSlice1, newSlice2, defaultEqualler)
}

func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool {
	newSlice1, newSlice2 := checkSameInterfaceParam(slice1, slice2)
	return coreElementMatch(newSlice1, newSlice2, equaller)
}

func coreElementMatch(slice1, slice2 innerSlice, equaller Equaller) bool {
	extra1 := makeInnerSlice(slice1, 0, 0)
	extra2 := makeInnerSlice(slice2, 0, 0)

	visited := make([]bool, slice2.length())
	for idx1 := 0; idx1 < slice1.length(); idx1++ {
		item1 := slice1.get(idx1)
		exist := false
		for idx2 := 0; idx2 < slice2.length(); idx2++ {
			if visited[idx2] {
				continue
			}
			item2 := slice2.get(idx2)
			if equaller(item1, item2) {
				visited[idx2] = true
				exist = true
				break
			}
		}
		if !exist {
			extra1.append(item1)
		}
	}

	for item2 := 0; item2 < slice2.length(); item2++ {
		if !visited[item2] {
			extra2.append(item2)
		}
	}

	return extra1.length() == 0 && extra2.length() == 0
}

// Range generates an integer origin from small to large with step.
func Range(min, max, step int) []int {
	if min >= max {
		panic("min should less then max")
	} else if step <= 0 {
		panic("step should larger than 0")
	}

	out := make([]int, 0)
	for idx := min; idx <= max; idx += step {
		out = append(out, idx)
	}
	return out
}

// ReverseRange generates an reverse integer origin from small to large with step.
func ReverseRange(min, max, step int) []int {
	if min >= max {
		panic("min should less then max")
	} else if step <= 0 {
		panic("step should larger than 0")
	}

	out := make([]int, 0)
	for idx := max; idx >= min; idx -= step {
		out = append(out, idx)
	}
	return out
}

// // GenerateByIndex generates a origin by indies and a generate function.
// func GenerateByIndex(indies []int, fn func(i int) interface{}) []interface{} {
// 	out := make([]interface{}, len(indies))
// 	for idx, num := range indies {
// 		out[idx] = fn(num)
// 	}
// 	return out
// }
