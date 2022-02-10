package leetcode

import (
	"testing"
)

func maxDepth(s string) int {
	var depth, size int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			size++
			if size > depth {
				depth = size
			}
		} else if s[i] == ')' {
			size--
		}
	}
	return depth
}

func Test1614(t *testing.T) {
	var s = "(1+(2*3)+((8)/4))+1"

	t.Log(maxDepth(s))
}
