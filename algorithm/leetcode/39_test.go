package leetcode

import (
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var sum int
	var path []int
	var ans [][]int
	var backtracking func(index int)
	backtracking = func(index int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

func Test39(t *testing.T) {
	var candidates = []int{2, 3, 6, 7}
	var target = 7
	t.Log(combinationSum(candidates, target))
}
