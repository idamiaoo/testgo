package sword2

import (
	"testing"
)

func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})

	var backtracking func(nums []int, i int, path []int)

	backtracking = func(nums []int, i int, path []int) {
		if i < 0 || i >= len(nums) || nums[i] == -99 {
			return
		}
		path = append(path, nums[i])
		l := len(path)
		var temp []int
		for _, v := range path {
			temp = append(temp, v)
		}
		ans = append(ans, temp)
		nums[i] = -99
		for j := i + 1; j < len(nums); j++ {
			backtracking(nums, j, path)
		}
		nums[i] = path[l-1]
		path = path[:l-1]
	}
	for j := 0; j < len(nums); j++ {
		backtracking(nums, j, nil)
	}
	return ans
}

/*
[9,0,3,5,7]

[[],[9],[9,0],[9,0,3],[9,0,3,7],[9,0,3,5,7],[9,0,3,7],[9,0,5],[9,0,5,7],[9,0,7],[9,3],[9,3,5],[9,3,5,7],[9,3,7],[9,5],
[9,5,7],[9,7],[0],[0,3],[0,3,5],[0,3,5,7],[0,3,7],[0,5],[0,5,7],[0,7],[3],[3,5],[3,5,7],[3,7],[5],[5,7],[7]]

[[],[0],[0,3],[0,3,5],[0,3,5,9],[0,3,5,7,9],[0,3,5,9],[0,3,7],[0,3,7,9],[0,3,9],[0,5],[0,5,7],[0,5,7,9],[0,5,9],[0,7],
[0,7,9],[0,9],[3],[3,5],[3,5,7],[3,5,7,9],[3,5,9],[3,7],[3,7,9],[3,9],[5],[5,7],[5,7,9],[5,9],[7],[7,9],[9]]

[[],[9],[0],[0,9],[3],[3,9],[0,3],[0,3,9],[5],[5,9],[0,5],[0,5,9],[3,5],[3,5,9],[0,3,5],[0,3,5,9],[7],[7,9],[0,7],
[0,7,9],[3,7],[3,7,9],[0,3,7],[0,3,7,9],[5,7],[5,7,9],[0,5,7],[0,5,7,9],[3,5,7],[3,5,7,9],[0,3,5,7],[0,3,5,7,9]]
*/

func Test79(t *testing.T) {
	nums := []int{0, 3, 5, 7, 9}
	t.Log(subsets(nums))
}
