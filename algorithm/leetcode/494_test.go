package leetcode

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j > nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[n][neg]
}
