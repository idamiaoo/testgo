package sword2

import (
	"strconv"
	"strings"
	"testing"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	var path []string
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == 4 {
			if index == len(s) {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}
		for i := index; i < len(s); i++ {
			if isValidIP(s[index : i+1]) {
				path = append(path, s[index:i+1])
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

func isValidIP(s string) bool {
	v, _ := strconv.Atoi(s)
	if v <= 255 && (s == "0" || s[0] != '0') {
		return true
	}
	return false
}

func Test87(t *testing.T) {
	// var s = "25525511135"
	var s = "10203040"
	t.Log(restoreIpAddresses(s))
}
