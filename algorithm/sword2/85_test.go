package sword2

import (
	"testing"
)

func generateParenthesis(n int) []string {
	var ans []string
	var path string
	var backtracking func(right, left int)
	backtracking = func(right, left int) {
		if left == n && right == n {
			ans = append(ans, path)
			return
		}

		if right < n {
			path += "("
			right += 1
			backtracking(right, left)
			right -= 1
			path = path[:len(path)-1]
		}
		if left < right {
			path += ")"
			left += 1
			backtracking(right, left)
			left -= 1
			path = path[:len(path)-1]
		}

	}
	backtracking(0, 0)
	return ans
}

func Test85(t *testing.T) {
	t.Log(generateParenthesis(3))
}
