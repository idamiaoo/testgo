package sword2

type stack struct {
	top    int
	values []int
}

func newStack() *stack {
	return &stack{
		top:    -1,
		values: make([]int, 0, 1),
	}
}

func (s *stack) Pop() int {
	v := s.values[s.top]
	s.top--
	return v
}

func (s *stack) Range() []int {
	return s.values[:s.top+1]
}

func (s *stack) Top() int {
	return s.values[s.top]
}

func (s *stack) Size() int {
	return s.top + 1
}

func (s *stack) Put(v int) {
	if s.top == len(s.values)-1 {
		s.values = append(s.values, v)
		s.top++
	} else {
		s.top++
		s.values[s.top] = v
	}
}
