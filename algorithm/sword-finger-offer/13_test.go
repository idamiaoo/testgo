package sword_finger_offer

import (
	"testing"
)

func movingCount(m int, n int, k int) int {
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
	}
	var count int
	dfs(board, 0, 0, k, &count)
	return count
}

func dfs(board [][]int, i, j, k int, count *int) {
	if i < 0 || j < 0 || i >= len(board) || j >= (len(board[0])) || dSum(i, j) > k || board[i][j] == -1 {
		return
	}
	*count++
	board[i][j] = -1
	dfs(board, i-1, j, k, count)
	dfs(board, i+1, j, k, count)
	dfs(board, i, j-1, k, count)
	dfs(board, i, j+1, k, count)
}

func dSum(a, b int) int {
	var sum int
	for a > 0 {
		sum += a % 10
		a = a / 10
	}
	for b > 0 {
		sum += b % 10
		b = b / 10
	}
	return sum
}

func TestMovingCount(t *testing.T) {
	t.Log(movingCount(3, 2, 17))
}
