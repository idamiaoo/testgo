package sword2

import (
	"testing"
)

func findKthLargest(nums []int, k int) int {
	var left, right = 0, len(nums) - 1

	p := partition76(nums, left, right)
	for p != k-1 {
		if p > k-1 {
			right = p - 1
		} else {
			left = p + 1
		}
		p = partition76(nums, left, right)
	}
	return nums[k-1]
}

func partition76(nums []int, left, right int) int {
	b := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] > b {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

/*
[5,2,4,1,3,6,0]
4
[3,2,3,1,2,4,5,5,6]
9
*/

func Test76(t *testing.T) {
	var nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	var k = 9
	t.Log(findKthLargest(nums, k))
}
