package leetcode

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		if nums[first] > 0 {
			break
		}
		second, third := first+1, len(nums)-1
		for second < third {
			s := nums[first] + nums[second] + nums[third]
			if s < 0 {
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}
			} else if s > 0 {
				third--
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				third--
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			}
		}
	}
	return ans
}

func Test15(t *testing.T) {
	var nums = []int{-2, 0, 0, 2, 2}
	t.Log(threeSum(nums))
}
