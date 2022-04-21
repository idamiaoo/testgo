package leetcode

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = nums[i] * ans[i-1]
	}
	temp := 1
	for i := len(nums) - 1; i >= 1; i-- {
		ans[i] = ans[i-1] * temp
		temp *= nums[i]
	}
	ans[0] = temp
	return ans
}
