package leetcode

import (
	"sort"
	"testing"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var path []int
	var backtracking func()
	choice := make(map[int]bool)
	backtracking = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if choice[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && choice[i-1] == false {
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

func Test47(t *testing.T) {

}
