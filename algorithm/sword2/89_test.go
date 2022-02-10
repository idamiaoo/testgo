package sword2

import (
	"testing"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := nums[0]
	s := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f, s = s, max(s, f+nums[i])
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test89(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	t.Log(rob(nums))
}
