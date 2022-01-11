package sword2

import (
	"testing"
)

func combine(n int, k int) [][]int {
	var ans [][]int
	var helper func(n int, cur int, path []int)
	helper = func(n int, cur int, path []int) {
		if len(path) == k {
			var temp []int
			for _, v := range path {
				temp = append(temp, v)
			}
			ans = append(ans, temp)
			return
		}

		if cur > n {
			return
		}

		helper(n, cur+1, path)
		path = append(path, cur)
		l := len(path)
		helper(n, cur+1, path)
		path = path[:l-1]

	}
	helper(n, 1, nil)
	return ans
}

func Test80(t *testing.T) {
	t.Log(combine(4, 2))
}

func quanpailie(str string) []string {
	var ans []string
	var backtracking func(string, []byte)
	backtracking = func(str string, path []byte) {
		if len(path) == len(str) {
			ans = append(ans, string(path))
		}

		for i := 0; i < len(str); i++ {
			if search(path, str[i]) {
				continue
			}
			path = append(path, str[i])
			l := len(path)
			backtracking(str, path)
			path = path[:l-1]
		}
	}
	backtracking(str, nil)
	return ans
}

func search(arr []byte, b byte) bool {
	for _, v := range arr {
		if v == b {
			return true
		}
	}
	return false
}

func TestQuanpailie(t *testing.T) {
	t.Log(quanpailie("abc"))
}
