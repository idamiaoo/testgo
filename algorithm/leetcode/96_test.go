package leetcode

func numTrees(n int) int {
	dp := make([]int, n)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j >= i; j++ {
			dp[i] = dp[i-1] * dp[i-j]
		}
	}
	return dp[n]
}
