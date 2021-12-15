package sword_finger_offer

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	var maxLen int
	var i, j int
	for j < len(s) {
		if _, ok := m[s[j]]; ok {
			if j-i >= maxLen {
				maxLen = j - i
			}
			delete(m, s[i])
			i++
		} else {
			m[s[j]] = struct{}{}
			j++
		}
	}
	if j-i > maxLen {
		maxLen = j - i
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("pwwkew"))
}
