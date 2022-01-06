package stack

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) Push(val T) {
	s.s = append(s.s, val)
}

func (s *Stack[T]) Size() int {
	return len(s.s)
}

func (s *Stack[T]) Pop() T {
	last := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return last
}

func (s *Stack[T]) Peek() T {
	last := s.s[len(s.s)-1]
	return last
}
