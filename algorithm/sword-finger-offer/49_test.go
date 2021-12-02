package sword_finger_offer

import (
	"testing"
)

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	p2, p3, p5 := 1, 1, 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestNthUglyNumber(t *testing.T) {
	t.Log(nthUglyNumber(10))
}
