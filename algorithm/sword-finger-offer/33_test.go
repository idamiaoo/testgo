package sword_finger_offer

import (
	"testing"
)

func verifyPostorder(postorder []int) bool {
	return false
}

func qSort(nums []int, left, right int) {
	if left <= right {
		p := partition(nums, left, right)
		qSort(nums, left, p-1)
		qSort(nums, p+1, right)
	}
}

func partition(nums []int, left, right int) int {
	s := nums[right]
	p := left - 1
	for i := left; i < right; i++ {
		if nums[i] < s {
			p++
			nums[p], nums[i] = nums[i], nums[p]
		}
	}
	p++
	nums[p], nums[right] = nums[right], nums[p]
	return p
}

func bSort(nums []int, left, right int) {
	for i := left; i < right; i++ {
		for j := left; j <= right-(i-left)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func TestSort(t *testing.T) {
	nums := []int{9, 3, 6, 4, 0, 5, 8, 2}
	bSort(nums, 0, len(nums)-1)
	t.Log(nums)
}
