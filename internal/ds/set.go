package ds

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

// Contains checks if the item v is present in the set s.
func (s Set[T]) Contains(v T) bool {
	_, found := s[v]
	return found
}
