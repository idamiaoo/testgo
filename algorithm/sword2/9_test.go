package sword2

import (
	"testing"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	var p = 1
	var res int
	for i, j := 0, 0; j < len(nums); {
		p *= nums[j]
		for p >= k && i <= j {
			p /= nums[i]
			i++
		}
		if i <= j {
			res += j - i + 1
		}
		j++
	}
	return res
}

// 10 ,2 , 5 , 6

// 11

func Test9(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100
	t.Log(numSubarrayProductLessThanK(nums, k))
}
