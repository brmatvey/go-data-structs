package stack

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, 10),
	}
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Len() int    { return len(s.items) }
func (s *Stack[T]) Peek() T     { return s.items[len(s.items)-1] }
func (s *Stack[T]) Pop()        { s.items = s.items[:len(s.items)-1] }
func (s *Stack[T]) Push(val T)  { s.items = append(s.items, val) }
func (s *Stack[T]) Empty() bool { return len(s.items) == 0 }
