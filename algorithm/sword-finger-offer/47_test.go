package sword_finger_offer

import (
	"testing"
)

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxValue(t *testing.T) {
	tests := []struct {
		cases [][]int
		value int
	}{
		{
			cases: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			value: 12,
		},
	}

	for _, c := range tests {
		if maxValue(c.cases) != c.value {
			t.FailNow()
		}
	}
}
