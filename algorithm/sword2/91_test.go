package sword2

import (
	"testing"
)

func minCost(costs [][]int) int {
	r, b, g := costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		r, b, g = min(b+costs[i][0], g+costs[i][0]), min(r+costs[i][1], g+costs[i][1]), min(r+costs[i][2], b+costs[i][2])
	}
	return min(r, min(b, g))
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test91(t *testing.T) {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	t.Log(minCost(costs))
}
