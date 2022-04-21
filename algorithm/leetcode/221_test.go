package leetcode

import (
	"testing"
)

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	s := &stack{}
	heights := make([]byte, n+1)
	var ans int
	for i := 0; i < m; i++ {
		s.Put(-1)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		for i, v := range heights {
			for s.Top() != -1 && heights[s.Top()] > v {
				height := int(heights[s.Pop()])
				width := i - s.Top() - 1
				if height >= width {
					if width*width > ans {
						ans = width * width
					}
				} else {
					if height*height > ans {
						ans = height * height
					}
				}
			}
			s.Put(i)
		}
		s.Reset()
	}
	return ans
}

type stack struct {
	data []int
}

func (s *stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *stack) Put(v int) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *stack) Reset() {
	s.data = nil
}

func Test221(t *testing.T) {
	var matrix = [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	t.Log(maximalSquare(matrix))
}
