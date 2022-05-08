package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func restoreIpAddresses(s string) []string {
	var (
		ans          []string
		path         []string
		backtracking func(index int)
	)
	backtracking = func(index int) {
		if len(path) > 4 || index > len(s) {
			return
		}
		if index == len(s) {
			if len(path) == 4 {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}
		if isValidIP(s[index : index+1]) {
			path = append(path, s[index:index+1])
			backtracking(index + 1)
			path = path[:len(path)-1]
		}
		if index+2 <= len(s) && isValidIP(s[index:index+2]) {
			path = append(path, s[index:index+2])
			backtracking(index + 2)
			path = path[:len(path)-1]
		}
		if index+3 <= len(s) && isValidIP(s[index:index+3]) {
			path = append(path, s[index:index+3])
			backtracking(index + 3)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

func isValidIP(s string) bool {
	v, _ := strconv.Atoi(s)
	if (len(s) > 1 && s[0] == '0') || v > 255 {
		return false
	}
	return true
}

func Test93(t *testing.T) {
	var s = "101023"
	t.Log(restoreIpAddresses(s))
}
