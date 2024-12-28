package ds

import (
	"iter"
	"maps"
	"slices"
)

type empty struct{}

type Set[T comparable] map[T]empty

func NewSet[T comparable]() Set[T] {
	return make(map[T]empty)
}

// Add tries to add v to the set s. Returns false if v is already in the set.
func (s Set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}
	s[v] = empty{}
	return true
}

// All returns an iterator over the values in the set s. The iteration order is
// unspecified and may differ between invokations.
func (s Set[T]) All() iter.Seq[T] {
	return maps.Keys(s)
}

// Clone creates a clone of set s with shallowly copied values.
func (s Set[T]) Clone() Set[T] {
	return maps.Clone(s)
}

// Contains checks if the item v is present in the set s.
func (s Set[T]) Contains(v T) bool {
	_, found := s[v]
	return found
}

// Difference creates a new set of all items in s which are not present in the
// other set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := make(map[T]empty)
	for item := range s {
		if !other.Contains(item) {
			diff[item] = empty{}
		}
	}
	return diff
}

// Intersection creates a new set of the items present in both set s and other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	items := make(map[T]empty)
	for item := range s {
		if other.Contains(item) {
			items[item] = empty{}
		}
	}
	return items
}

// Items returns an unordered slice of all items in the set s.
func (s Set[T]) Items() []T {
	return slices.Collect(s.All())
}

// Union creates a new set of all items present in either set s or other.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := maps.Clone(s)
	maps.Copy(union, other)
	return union
}
