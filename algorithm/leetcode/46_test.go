package leetcode

import (
	"testing"
)

func permute(nums []int) [][]int {
	var path []int
	var ans [][]int
	choice := make(map[int]bool)

	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if choice[i] {
				continue
			}
			path = append(path, nums[i])
			choice[i] = true
			backtracking()
			choice[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return ans
}

func Test46(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}
