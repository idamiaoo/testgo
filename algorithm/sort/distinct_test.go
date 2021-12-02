package sort

import (
	"testing"
)

func distinct(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	var i, j int
	for i, j = 0, 1; j < len(nums); {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return nums[:i+1]
}

func TestDistinct(t *testing.T) {
	a := []int{1, 1, 2, 3, 4, 4, 6, 7, 7, 7, 8, 10}
	t.Log(distinct(a))
}
