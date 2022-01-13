package leetcode

import (
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var i int
	for {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || i >= len(strs[j-1]) || strs[j][i] != strs[j-1][i] {
				return strs[0][:i]

			}
		}
		i++
	}
}

func Test14(t *testing.T) {
	var strs = []string{"", "b"}
	t.Log(longestCommonPrefix(strs))
}
