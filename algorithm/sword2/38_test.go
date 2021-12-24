package sword2

import (
	"fmt"
	"testing"
)

func dailyTemperatures(temperatures []int) []int {
	s := newStack()
	var res = make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		fmt.Println(s.Range())
		for s.Size() > 0 && temperatures[s.Top()] < temperatures[i] {
			v := s.Pop()
			res[v] = i - v
		}
		s.Put(i)
	}
	return res
}

func Test38(t *testing.T) {
	var temp = []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Log(dailyTemperatures(temp))
}
