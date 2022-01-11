package sword2

import (
	"sort"
	"testing"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var path []int
	vis := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if vis[i] {
				continue
			}

			if i > 0 && nums[i-1] == nums[i] && vis[i-1] == false {
				continue
			}
			path = append(path, nums[i])
			vis[i] = true
			backtracking()
			vis[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return ans
}

/*
1,1,2,2,3,3
*/

func Test84(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	t.Log(permuteUnique(nums))
}
