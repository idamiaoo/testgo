package sword2

import (
	"testing"
)

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[len(nums)-1]-nums[i] == 0 {
				return i
			}
			continue
		}
		if i == len(nums)-1 {
			if nums[len(nums)-2] == 0 {
				return i
			}
			continue
		}
		if nums[i-1] == nums[len(nums)-1]-nums[i] {
			return i
		}
	}
	return -1
}

func Test12(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	nums1 := []int{1, 2, 3}
	t.Log(pivotIndex(nums), pivotIndex(nums1))
}
