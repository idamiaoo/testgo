package sword2

import (
	"testing"
)

func largestRectangleArea(heights []int) int {
	heights = append(heights, -1)
	var res int
	s := newStack()
	s.Put(-1)
	var i int
	for i = range heights {
		for s.Top() != -1 && heights[i] < heights[s.Top()] {
			height := heights[s.Top()]
			s.Pop()
			v := height * (i - s.Top() - 1)
			if v > res {
				res = v
			}
		}
		s.Put(i)
	}
	return res
}

func Test39(t *testing.T) {
	t.Log(largestRectangleArea([]int{2, 1, 2}))
}
