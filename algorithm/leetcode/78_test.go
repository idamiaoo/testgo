package leetcode

import (
	"testing"
)

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int
	var backtracking func(int)
	backtracking = func(index int) {
		ans = append(ans, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

func Test78(t *testing.T) {
	var nums = []int{1, 2, 3}
	t.Log(subsets(nums))
}
