package leetcode

import (
	"testing"
)

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var first, second int
	if nums[0] > nums[1] {
		first = 0
		second = 1
	} else {
		first = 1
		second = 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[first] {
			second = first
			first = i
		} else if nums[i] > nums[second] {
			second = i
		}
	}
	if nums[second]*2 <= nums[first] {
		return first
	}
	return -1
}

func Test747(t *testing.T) {
	t.Log(dominantIndex([]int{1, 0}))
}
