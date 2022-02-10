package sword2

import (
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	var f, s int
	var t int
	for i := 2; i <= len(cost); i++ {
		t = min(s+cost[i-1], f+cost[i-2])
		f = s
		s = t
	}
	return s
}

func min88(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test88(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	t.Log(minCostClimbingStairs(cost))
}

// 10 ,15 ,20
