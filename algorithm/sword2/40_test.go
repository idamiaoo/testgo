package sword2

import (
	"testing"
)

func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	var res int
	for _, row := range matrix {
		for j := range row {
			if row[j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		if v := largestRectangleArea(heights); v > res {
			res = v
		}
	}
	return res
}

func Test40(t *testing.T) {

}
