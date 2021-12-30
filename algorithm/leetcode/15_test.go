package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	for l, r := 0, len(nums)-1; l < r; {
		v := nums[l] + nums[r]
		if v > 0 {

		} else {

		}
	}
	return nil
}
