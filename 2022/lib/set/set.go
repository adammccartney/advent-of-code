package set

type Set struct {
	m map[rune]bool
}

func NewSet() *Set {
	return &Set{m: make(map[rune]bool)}
}

func (s *Set) Add(r rune) {
	s.m[r] = true
}

func (s *Set) Remove(r rune) {
	delete(s.m, r)
}

func (s *Set) Contains(r rune) bool {
	_, ok := s.m[r]
	return ok
}
