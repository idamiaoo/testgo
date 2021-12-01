package leetcode

import (
	"testing"
)

func maxPower(s string) int {
	if len(s) == 1 {
		return 1
	}
	max, cur := 1, 1
	for i, j := 0, 1; i < j && j < len(s); {
		if s[i] == s[j] {
			cur++
			j++
		} else {
			if cur > max {
				max = cur
			}
			cur = 1
			i = j
			j = j + 1
		}
	}
	if cur > max {
		max = cur
	}
	return max
}

func TestMaxPower(t *testing.T) {
	t.Log(maxPower("cc"))
}
