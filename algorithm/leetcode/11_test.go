package leetcode

import (
	"testing"
)

func maxArea(height []int) int {
	var ans int
	for l, r := 0, len(height)-1; l < r; {
		if height[l] >= height[r] {
			if height[r]*(r-l) > ans {
				ans = height[r] * (r - l)
			}
			r--
		} else {
			if height[l]*(r-l) > ans {
				ans = height[l] * (r - l)
			}
			l++
		}
	}
	return ans
}

func Test11(t *testing.T) {
	height := []int{4, 3, 2, 1, 4}
	t.Log(maxArea(height))
}
