package leetcode

import (
	"testing"
)

func removeInvalidParentheses(s string) []string {
	var lRemove, rRemove int
	for _, v := range []byte(s) {
		if v == '(' {
			lRemove++
		} else if v == ')' {
			if lRemove == 0 {
				rRemove++
			} else {
				lRemove--
			}
		}
	}
	var path []byte
	var ans []string
	var backtracking func(int, int, int)
	backtracking = func(l, r int, index int) {
		if l == 0 && r == 0 && len(path) == len(s)-lRemove-rRemove {
			if isValid(string(path)) {
				ans = append(ans, string(path))
			}
			return
		}
		if index >= len(s) {
			return
		}

		switch s[index] {
		case '(':
			if l > 0 {
				v := index + 1
				for v < len(s)-1 && s[index] == s[v] {
					path = append(path, s[index])
					v++
				}
				backtracking(l-1, r, v)
				path = path[:len(path)-(v-index-1)]
			}
			path = append(path, s[index])
			backtracking(l, r, index+1)
			path = path[:len(path)-1]
		case ')':
			if r > 0 {
				v := index + 1
				for v < len(s)-1 && s[index] == s[v] {
					path = append(path, s[index])
					v++
				}
				backtracking(l-1, r, v)
				path = path[:len(path)-(v-index-1)]
			}
			path = append(path, s[index])
			backtracking(l, r, index+1)
			path = path[:len(path)-1]
		default:
			path = append(path, s[index])
			backtracking(l, r, index+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(lRemove, rRemove, 0)
	return ans
}

func isValid(s string) bool {
	var cnt int
	for _, v := range []byte(s) {
		if v == '(' {
			cnt++
		} else if v == ')' {
			if cnt == 0 {
				return false
			}
			cnt--
		}
	}
	return cnt == 0
}

func Test301(t *testing.T) {
	var s = "(a)())()"
	var s1 = "()())()"
	t.Log(removeInvalidParentheses(s))
	t.Log(removeInvalidParentheses(s1))

}
