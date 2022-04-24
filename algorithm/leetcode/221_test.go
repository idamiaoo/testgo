package leetcode

func maximalSquare(matrix [][]byte) int {
	var ans int
	m, n := len(matrix), len(matrix[0])
	temp := make([]byte, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				temp[j] = matrix[i][j]
			} else {
				if matrix[i][j] == 0 {
					temp[j] = 0
				} else {
					temp[j] += 1
				}
			}
		}
	}

	return ans
}
