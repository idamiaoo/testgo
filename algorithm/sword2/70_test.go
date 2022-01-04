package sword2

func singleNonDuplicate(nums []int) int {
	var left, right = 0, len(nums) - 1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 2
			} else {
				right = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nums[right]
}
