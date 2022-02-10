package sword2

import (
	"testing"
)

func permute(nums []int) [][]int {
	var ans [][]int
	var path []int
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if nums[i] == -99 {
				continue
			}
			path = append(path, nums[i])
			nums[i] = -99
			backtracking()
			nums[i] = path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return ans
}

func Test83(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(permute(nums))
}
