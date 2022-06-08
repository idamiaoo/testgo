package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func largestCombination(candidates []int) int {
	var (
		ans          = 1
		backtracking func(int, int, int, int) bool
	)

	backtracking = func(index, l, cur, sum int) bool {
		fmt.Println(index, l, cur)
		if index == len(candidates) || sum == 0 {
			if cur == l && sum > 0 {
				return true
			}
			return false
		}

		if backtracking(index+1, l, cur+1, sum&candidates[index]) {
			return true
		}

		if backtracking(index+1, l, cur, sum) {
			return true
		}
		return false
	}

	for l := len(candidates); l > 1; l-- {
		if backtracking(0, l, 0, math.MaxInt) {
			ans = l
			break
		}

	}
	return ans
}

func and(nuns []int) int {
	ans := nuns[0]
	for i := 1; i < len(nuns); i++ {
		ans &= nuns[i]
	}
	return ans
}

func Test2275(t *testing.T) {
	t.Log(largestCombination([]int{33, 93, 31, 99, 74, 37, 3, 4, 2, 94, 77, 10, 75, 54, 24, 95, 65, 100, 41, 82, 35, 65, 38, 49, 85, 72, 67, 21, 20, 31}))
}
