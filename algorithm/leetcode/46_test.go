package leetcode

import (
	"testing"
)

func permute(nums []int) [][]int {
	var result [][]int
	for i := range nums {
		dfs(nums, nil, i, &result)
	}
	return result
}

func dfs(s, p []int, i int, result *[][]int) {
	if s[i] == -11 {
		return
	}
	p = append(p, s[i])
	if len(p) == len(s) {
		*result = append(*result, p)
		return
	}
	s[i] = -11
	for ii := range s {
		dfs(s, p, ii, result)
	}
	s[i] = p[len(p)-1]
}

func TestPermute(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}
