package leetcode

import (
	"fmt"
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
		if len(path) == len(s)-lRemove-rRemove {
			if isValid(string(path)) {
				ans = append(ans, string(path))
			}
			return
		}
		if index > 0 && s[index] == s[index-1] {
			return
		}
		if l > 0 && s[index] == '(' {
			v := index + 1
			for s[v] == s[index] {
				v++
			}
			path = append(path, s[index+1:v]...)
			backtracking(l-1, r, v)
			path = path[:len(path)-(v-index-1)]
		}

		if r > 0 && s[index] == ')' {
			v := index + 1
			for s[v] == s[index] {
				v++
			}
			path = append(path, s[index+1:v]...)
			fmt.Println(string(path), v, index)
			backtracking(l, r-1, v)
			path = path[:len(path)-(v-index-1)]
		}
		path = append(path, s[index])
		backtracking(l, r, index+1)
		path = path[:len(path)-1]
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
			cnt--
			if cnt < 0 {
				return false
			}
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
