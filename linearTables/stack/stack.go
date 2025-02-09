package stack

type number interface {
	int | float64
}

type Stack[T number] struct {
	Value []T
}

func NewStack() *Stack[int] {
	return &Stack[int]{}
}

func (s *Stack[T]) push(element T) {
	s.Value = append(s.Value, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.Value) == 0 {
		var zero T
		return zero, false
	} else {
		element := s.Value[len(s.Value)-1]
		s.Value = s.Value[:len(s.Value)-1]
		return element, true
	}
}

func (s *Stack[T]) top() (T, bool) {
	if len(s.Value) == 0 {
		var zero T
		return zero, false
	}

	return s.Value[len(s.Value)-1], true
}

func (s *Stack[T]) size() uint {
	return uint(len(s.Value))
}
