package sword2

import (
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var helper func(start, sum int, path []int)
	helper = func(start, sum int, path []int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			helper(i, sum, path)
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	helper(0, 0, nil)
	return ans
}

func Test81(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	t.Log(combinationSum(candidates, target))
}
