package leetcode

func exist(board [][]byte, word string) bool {
	var ans bool
	var backtracking func(int, int, int)
	backtracking = func(i, j, index int) {
		if board[i][j] != word[index] {
			return
		}
		if index == len(word)-1 {
			ans = true
			return
		}

		if i > 0 {
			board[i][j] = '0'
			backtracking(i-1, j, index+1)
			board[i][j] = word[index]
		}
		if j > 0 {
			board[i][j] = '0'
			backtracking(i, j-1, index+1)
			board[i][j] = word[index]
		}
		if i < len(board)-1 {
			board[i][j] = '0'
			backtracking(i+1, j, index+1)
			board[i][j] = word[index]

		}
		if j < len(board[0])-1 {
			board[i][j] = '0'
			backtracking(i, j+1, index+1)
			board[i][j] = word[index]
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtracking(i, j, 0)
		}
	}
	return ans
}
