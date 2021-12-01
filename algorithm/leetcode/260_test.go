package leetcode

import (
	"testing"
)

func singleNumber(nums []int) []int {
	var x int
	for _, v := range nums {
		x ^= v
	}
	l := x & (-x)
	var n1, n2 int
	for _, v := range nums {
		if l&v > 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}
	return []int{n1, n2}
}

func TestSingleNumber(t *testing.T) {
	t.Log(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

// 0000 0101
// 0000 0011
// 0000 0110
// 1111 1010
