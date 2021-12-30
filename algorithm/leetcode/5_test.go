package leetcode

import (
	"testing"
)

func longestPalindrome(s string) string {
	return ""
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
