package sword2

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if !dp[i][j] && j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]
}
