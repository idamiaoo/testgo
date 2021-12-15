package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rows, cols := make([]int, len(grid)), make([]int, len(grid))
	l := len(grid)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] > rows[i] {
				rows[i] = grid[i][j]
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[j][i] > cols[i] {
				cols[i] = grid[j][i]
			}
		}
	}

	var sum int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if rows[i] < cols[j] {
				sum += rows[i] - grid[i][j]
			} else {
				sum += cols[j] - grid[i][j]
			}
		}
	}

	return sum
}
