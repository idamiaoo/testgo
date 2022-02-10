package leetcode

import (
	"testing"
)

func longestConsecutive(nums []int) int {
	var ans, cur int
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	for _, v := range nums {
		if cnt[v-1] == 0 {
			cur = 0
			for cnt[v] != 0 {
				cur++
				v++
			}
			if cur > ans {
				ans = cur
			}
		}
	}
	return ans
}

func Test128(t *testing.T) {
	t.Log(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
