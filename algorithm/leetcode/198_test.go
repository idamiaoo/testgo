package leetcode

import "testing"

func rob198(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max198(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max198(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test198(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	t.Log(rob198(nums))
}
