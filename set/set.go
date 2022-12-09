package set

type set struct {
	m map[rune]bool
}

func NewSet() *set {
	return &set{m: make(map[rune]bool)}
}

func (s *set) Add(r rune) {
	s.m[r] = true
}

func (s *set) Remove(r rune) {
	delete(s.m, r)
}

func (s *set) Contains(r rune) bool {
	_, ok := s.m[r]
	return ok
}
