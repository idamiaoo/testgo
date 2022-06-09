package leetcode

func minimumObstacles(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i==0 {
				if j== 0 {
					dp[i][j]= grid[i][j]
				} else if j == len(grid)-1{
					dp[i][j]= grid[i][j] + dp[]
				} else {

				}
			}
				continue
			}

			if j == len(grid[0])-1 {

			}

			if i == len(grid) - 1 {

			}

			if j == 0 {

			}

		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min2290(a, b int) int {
	if a < b {
		return a
	}
	return b
}

[0,1,0,0,0],
[0,1,0,1,0],
[0,0,0,1,0]