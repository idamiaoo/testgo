package sword2

import (
	"testing"
)

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	var res, sum int
	for _, v := range nums {
		sum += v
		res += m[sum-k]
		m[sum]++
	}
	return res
}

func Test10(t *testing.T) {
	num := []int{1, 1, 1}
	k := 2
	t.Log(subarraySum(num, k))
}
