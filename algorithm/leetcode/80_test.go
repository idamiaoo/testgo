package leetcode

import (
	"testing"
)

func l80(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	l := 1
	r := 2
	for r < len(nums) {
		if nums[l] == nums[r] {
			if nums[l-1] != nums[l] {
				l++
				nums[l] = nums[r]
			}
		} else {
			l++
			nums[l] = nums[r]
		}
		r++
	}
	return l + 1
}

func Test80(t *testing.T) {
	nums := []int{2, 2, 2, 3, 3, 3}
	if l80(nums) != 4 {
		t.Error("l80 failed")
	}
}
