package sword2

import (
	"fmt"
	"testing"
)

func lenLongestFibSubseq(arr []int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	var ans int
	for j := 1; j < len(arr); j++ {
		for i := 0; i < j; i++ {
			temp := arr[j] - arr[i]
			if v, ok := m[temp]; ok && v < i {
				dp[i][j] = dp[v][i] + 1
			} else {
				dp[i][j] = 2
			}
			if dp[i][j] > 2 && dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func sumV(arr []int, v int) int {
	var ans int
	var backtracking func(int, int, int)

	temp := make([]int, 100)
	backtracking = func(index, fCount, score int) {
		if fCount == v {
			fmt.Println(score)
			temp[score]++
			return
		}

		if index >= len(arr) {
			return
		}

		backtracking(index+1, fCount, score+arr[index])
		backtracking(index+1, fCount+1, score)

	}
	backtracking(0, 0, 0)

	for _, v := range temp {
		if v > 0 {
			ans++
		}
	}
	return ans
}

func TestSumV(t *testing.T) {
	var arr = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8}
	t.Log(sumV(arr, 3))
}
