package leetcode

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	var ans []string
	if digits == "" {
		return ans
	}
	var path []byte
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == len(digits) {
			ans = append(ans, string(path))
			return
		}
		var start, end int

		start = int(digits[index] - '2')
		switch start {
		case 5:
			start *= 3
			end = start + 4
		case 6:
			start = 5*3 + 4
			end = start + 3
		case 7:
			start = 5*3 + 4 + 3
			end = start + 4
		default:
			start *= 3
			end = start + 3
		}

		for j := start; j < end; j++ {
			path = append(path, 'a'+byte(j))
			backtracking(index + 1)
			path = path[:len(path)-1]
		}

	}
	backtracking(0)
	return ans
}

func Test17(t *testing.T) {
	fmt.Println(string('a' + 1))
	t.Log(letterCombinations("9"))
}
