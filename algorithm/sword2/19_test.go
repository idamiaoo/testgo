package sword2

import (
	"testing"
)

func validPalindrome(s string) bool {
	var vf func(string, int) bool
	vf = func(s string, n int) bool {
		for i, j := 0, len(s)-1; i < j; {
			if s[i] == s[j] {
				i++
				j--
				continue
			}
			if n > 0 {
				return vf(s[i+1:j+1], n-1) || vf(s[i:j], n-1)
			}
			return false
		}
		return true
	}
	return vf(s, 1)
}

func Test19(t *testing.T) {
	var s = "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"

	t.Log(validPalindrome(s))
}
