package main

type stack struct {
	s []rune
}

func (s *stack) push(r rune) {
	s.s = append(s.s, r)
}

func (s *stack) pop() (rune, bool) {
	if len(s.s) == 0 {
		return 0, true
	}

	last := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return last, false
}
