package sword2

import (
	"testing"
)

func subarraySum(nums []int, k int) int {
	var res int
	var sum int
	for i, j := 0, 0; j < len(nums); {
		sum += nums[j]
		for sum > k && i < j {
			sum -= nums[i]
			i++
		}
		if sum == k {
			res++
		}
		j++
	}
	return res
}

func Test10(t *testing.T) {
	num := []int{1, 1, 1}
	k := 2
	t.Log(subarraySum(num, k))
}
