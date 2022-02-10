package sword2

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var res int
	for i, j := 0, 0; j < len(s); {
		if _, ok := m[s[j]]; ok {
			delete(m, s[i])
			i++
		} else {
			m[s[j]] = struct{}{}
			if len(m) > res {
				res = len(m)
			}
			j++
		}
	}
	return res
}

func Test16(t *testing.T) {
	t.Log(lengthOfLongestSubstring("pwwkew"))
}
