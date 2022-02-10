package sword2

import (
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var backtracking func(start, sum int)
	var path []int

	backtracking = func(start, sum int) {
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
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtracking(i+1, sum)
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return ans
}

/*
[2,5,2,1,2]
5
*/

func Test82(t *testing.T) {
	candidates := []int{1, 2, 2, 2, 5}
	target := 5
	t.Log(combinationSum2(candidates, target))
}
