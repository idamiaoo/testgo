package leetcode

func maxSubArray(nums []int) int {
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func f55(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if i+nums[i] >= len(nums)-1 {
			return true
		}

	}
	return false
}
