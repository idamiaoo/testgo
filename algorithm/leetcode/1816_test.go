package leetcode

import (
	"testing"
)

func truncateSentence(s string, k int) string {
	var index int
	var v byte
	for index, v = range []byte(s) {
		if v == ' ' {
			k--
			if k == 0 {
				break
			}
		}
	}
	if k != 0 {
		return s
	}
	return s[0:index]
}

func TestTruncateSentence(t *testing.T) {
	t.Log(truncateSentence("Hello moto", 1))
}
