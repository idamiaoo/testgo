package leetcode

import (
	"testing"
)

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	var max, start int
	for r := 0; r < len(s); r++ {
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l < 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > max {
					max = r - l + 1
					start = l
				}
			}
		}
	}
	return s[start : start+max]
}

func Test5(t *testing.T) {
	var s = "abcdef"
	v := "#"
	for i := 0; i < len(s); i++ {
		v += string(s[i]) + "#"
	}
	v += "#"
	t.Log(v)
}
