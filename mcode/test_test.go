package testgo

import (
	"fmt"
	"testing"
)

type stack struct {
	size int
	ele  []int
}

func (s *stack) Pop() int {
	v := s.ele[s.size-1]
	s.ele = s.ele[:s.size-1]
	s.size--
	return v
}

func (s *stack) Push(v int) {
	s.size++
	s.ele = append(s.ele, v)
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) Top() int {
	return s.ele[s.size-1]
}

func (s *stack) Value() []int {
	return s.ele
}

func TestStack(t *testing.T) {
	s := &stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()

	t1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t2 := t1[1:7]
	fmt.Println(t2, cap(t2))

	// t.Log(nextExceed([]int{5, 3, 1, 2, 4}))
}

func nextExceed(input []int) []int {
	result := make([]int, len(input), len(input))
	for i := range result {
		result[i] = -1
	}
	iStack := &stack{}

	for i := 0; i < len(input); i++ {
		for iStack.Size() != 0 && input[iStack.Top()] < input[i] {
			fmt.Printf("i=%d, iStack.Top()=%d\n", i, iStack.Top())
			result[iStack.Top()] = i - iStack.Top()
			iStack.Pop()
			fmt.Println("result", result)
		}
		iStack.Push(i)
		fmt.Printf("i=%d, stack=%v\n", i, iStack.Value())
	}
	return result
}
