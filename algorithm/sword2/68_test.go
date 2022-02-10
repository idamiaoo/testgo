package sword2

import (
	"testing"
)

func searchInsert(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	var mid int
	for left < right {
		if nums[left] >= target {
			return left
		}
		if nums[right] < target {
			return right + 1
		}

		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func Test68(t *testing.T) {
	nums := []int{1, 3}
	t.Log(searchInsert(nums, 3))
}
