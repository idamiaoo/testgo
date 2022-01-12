package sword2

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
