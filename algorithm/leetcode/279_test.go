package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minNum = min279(minNum, dp[i-j*j])
		}
		dp[i] = minNum + 1
	}
	return dp[n]
}

func min279(a, b int) int {
	if a < b {
		return a
	}
	return b
}
