package xslice

import (
	"sort"
)

type SortSlice struct {
	s    []interface{}
	less func(i, j int) bool
	swap func(i, j int)
}

func (s SortSlice) Len() int {
	return len(s.s)
}

func (s SortSlice) Less(i, j int) bool {
	return s.less(i, j)
}

func (s SortSlice) Swap(i, j int) {
	s.swap(i, j)
}

func NewSortSlice(s []interface{}, less func(i, j int) bool) SortSlice {
	return SortSlice{
		s:    s,
		less: less,
		swap: func(i, j int) { s[i], s[j] = s[j], s[i] },
	}
}

func ReverseSortSlice(s SortSlice) SortSlice {
	return SortSlice{
		s: s.s,
		less: func(i, j int) bool {
			return !s.less(i, j)
		},
		swap: s.swap,
	}
}

func Sort(s []interface{}, less func(i, j int) bool) {
	interfaceSlice := NewSortSlice(s, less)
	sort.Sort(interfaceSlice)
}

func StableSort(s []interface{}, less func(i, j int) bool) {
	interfaceSlice := NewSortSlice(s, less)
	sort.Stable(interfaceSlice)
}
