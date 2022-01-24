package sword2

import (
	"testing"
)

func findTargetSumWays1(nums []int, target int) int {
	var backtracking func(index int)
	var temp, ans int
	backtracking = func(index int) {
		if index == len(nums) {
			if temp == target {
				ans++
			}
			return
		}
		temp += -nums[index]
		backtracking(index + 1)
		temp -= -nums[index]

		temp += nums[index]
		backtracking(index + 1)
		temp -= nums[index]
	}
	backtracking(0)
	return ans
}

func findTargetSumWays(nums []int, target int) int {
	return 0
}

func Test102(t *testing.T) {
	var nums = []int{1}
	var target = 1
	t.Log(findTargetSumWays(nums, target))
}
