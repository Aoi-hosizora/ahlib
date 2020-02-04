package xset

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
)

type Set struct {
	_a []interface{}
}

func NewSet() *Set {
	return &Set{
		_a: make([]interface{}, 0),
	}
}

func FromSlice(slice []interface{}) *Set {
	set := NewSet()
	for _, item := range slice {
		set.Add(item)
	}
	return set
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		if xslice.IndexOf(s._a, item) == -1 {
			s._a = append(s._a, item)
		}
	}
}

func (s *Set) Remove(items ...interface{}) {
	for _, item := range items {
		s._a = xslice.Delete(xslice.Sti(s._a), item, 1)
	}
}

func (s *Set) Clear() {
	s._a = make([]interface{}, 0)
}

func (s *Set) Contains(item interface{}) bool {
	return xslice.IndexOf(s._a, item) != -1
}

func (s *Set) Slice() []interface{} {
	return s._a
}

func (s *Set) Size() int {
	return len(s._a)
}

func (s *Set) String() string {
	return fmt.Sprintf("%v", s._a)
}

func (s *Set) Equal(other *Set) bool {
	if other.Size() != s.Size() {
		return false
	}
	for _, item := range s.Slice() {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

func (s *Set) Union(other *Set) *Set {
	ret := FromSlice(s._a)
	for _, item := range other.Slice() {
		ret.Add(item)
	}
	return ret
}

func (s *Set) Intersect(other *Set) *Set {
	ret := NewSet()
	for _, item := range other.Slice() {
		if s.Contains(item) {
			ret.Add(item)
		}
	}
	return ret
}

func (s *Set) Diff(other *Set) *Set {
	ret := FromSlice(s._a)
	for _, item := range other.Slice() {
		if s.Contains(item) {
			ret.Remove(item)
		}
	}
	return ret
}
