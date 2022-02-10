package leetcode

import (
	"testing"
)

func dailyTemperatures(temperatures []int) []int {
	var ans = make([]int, len(temperatures))
	s := &stack739{}
	for i := 0; i < len(temperatures); i++ {
		for !s.Empty() && temperatures[s.Top()] < temperatures[i] {
			ans[s.Top()] = i - s.Top()
			s.Pop()
		}
		s.Put(i)
	}
	return ans
}

type stack739 struct {
	size int
	data []int
}

func (s *stack739) Top() int {
	return s.data[s.size-1]
}

func (s *stack739) Empty() bool {
	return s.size == 0
}

func (s *stack739) Put(v int) {
	if s.size >= len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.size] = v
	}
	s.size++
}

func (s *stack739) Pop() int {
	s.size--
	return s.data[s.size]
}

func Test739(t *testing.T) {
	var temperatures = []int{30, 40, 50, 60}
	t.Log(dailyTemperatures(temperatures))
}
