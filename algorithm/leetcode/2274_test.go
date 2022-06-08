package leetcode

import (
	"testing"
)

func maxConsecutive(bottom int, top int, special []int) int {
	quickSort2274(special, 0, len(special)-1)
	var ans = max2274(special[0]-bottom, top-special[len(special)-1])
	if len(special) > 1 {
		for i := 1; i < len(special); i++ {
			ans = max2274(ans, special[i]-special[i-1]-1)
		}
	}
	return ans
}

func max2274(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func quickSort2274(nums []int, left, right int) {
	if left < right {
		p := partition2274(nums, left, right)
		quickSort2274(nums, left, p-1)
		quickSort2274(nums, p+1, right)
	}
}

func partition2274(nums []int, left, right int) int {
	b := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] < b {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func Test2274(t *testing.T) {
	var cases = []struct {
		bottom  int
		top     int
		special []int
		result  int
	}{
		{2, 9, []int{4, 6}, 3},
	}
	t.Run("test 2274", func(t *testing.T) {
		for _, v := range cases {
			if a := maxConsecutive(v.bottom, v.top, v.special); a != v.result {
				t.Fatal(v.bottom, v.top, v.special, v.result, a)
			}
		}
	})
}
