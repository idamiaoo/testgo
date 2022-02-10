package sword2

import (
	"testing"
)

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min99(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min99(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test99(t *testing.T) {
	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	v := minPathSum(grid)
	t.Log(v)
	t.Log(grid)
}
