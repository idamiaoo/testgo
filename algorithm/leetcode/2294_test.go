package leetcode

import (
	"math"
	"sort"
	"testing"
)

func partitionArray(nums []int, k int) int {
	var (
		ans int
		min = math.MinInt
	)
	sort.Ints(nums)
	for i := range nums {
		if min == math.MinInt || nums[i]-min > k {
			ans++
			min = nums[i]
		}
	}
	return ans
}

func Test2294(t *testing.T) {
	var cases = []struct {
		nums   []int
		k      int
		result int
	}{
		{[]int{3, 6, 1, 2, 5}, 2, 2},
		{[]int{1, 2, 3}, 1, 2},
		{[]int{2, 2, 4, 5}, 0, 3},
		{[]int{16, 8, 17, 0, 3, 17, 8, 20}, 10, 2},
	}
	for _, v := range cases {
		if result := partitionArray(v.nums, v.k); result != v.result {
			t.Fatalf("nums %v expect result is %d, but got %d", v.nums, v.result, result)
		}
	}
}
