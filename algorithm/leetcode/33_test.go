package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
