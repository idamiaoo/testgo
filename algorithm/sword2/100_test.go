package sword2

import (
	"testing"
)

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min100(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min100(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test100(t *testing.T) {
	var triangle = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	t.Log(minimumTotal(triangle))
}
