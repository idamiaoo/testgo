package leetcode

import "testing"

func findSubsequences(nums []int) [][]int {
	var (
		ans          [][]int
		path         []int
		backtracking func(index int)
	)

	backtracking = func(index int) {
		if len(path) > 1 {
			ans = append(ans, append([]int{}, path...))
		}
		if index == len(nums) {
			return
		}

		if len(path) == 0 {
			path = append(path, nums[index])
			backtracking(index + 1)
			path = path[:len(path)-1]
			backtracking(index + 1)
		} else {
			if nums[index] >= path[len(path)-1] {
				path = append(path, nums[index])
				backtracking(index + 1)
				path = path[:len(path)-1]
				backtracking(index + 1)
			}
		}
	}
	backtracking(0)
	return ans
}

func Test491(t *testing.T) {
	t.Log(findSubsequences([]int{4, 6, 7, 7}))
}
