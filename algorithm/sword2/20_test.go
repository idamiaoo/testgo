package sword2

import (
	"testing"
)

func countSubstrings(s string) int {
	if len(s) == 1 {
		return 1
	}
	var res = 1

	var check func(string, int, int) int

	check = func(s string, l, r int) int {
		if l > r {
			return 0
		}
		for i, j := l, r; i <= j; {
			if s[i] != s[j] {
				return check(s, l+1, r)
			}
			i++
			j--
		}
		return 1 + check(s, l+1, r)
	}

	for i := 1; i < len(s); i++ {
		res = res + check(s, 0, i)
	}
	return res
}

func Test20(t *testing.T) {
	t.Log(countSubstrings("aaa"))
}
