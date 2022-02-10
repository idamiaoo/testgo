package sword2

import (
	"fmt"
	"testing"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			i++
			continue
		}
		if !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z') || (s[j] >= '0' && s[j] <= '9')) {
			j--
			continue
		}
		fmt.Println(string(s[i]), string(s[j]))
		if s[i] >= '0' && s[i] <= '9' {
			if s[i] == s[j] {
				i++
				j--
				continue
			}

		} else {
			if s[i]-'a' == s[j]-'a' || s[i]-'a' == s[j]-'A' || s[i]-'A' == s[j]-'a' || s[i]-'A' == s[j]-'A' {
				i++
				j--
				continue
			}
		}
		return false

	}
	return true
}

func Test18(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama"))
}
