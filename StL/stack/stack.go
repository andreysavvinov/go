package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if len(s.items) == 0 {
		return zero, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T

	if len(s.items) == 0 {
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
