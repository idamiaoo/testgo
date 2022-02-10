package sword2

import (
	"testing"
)

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := NumMatrix{
		prefixSum: matrix,
	}
	var sum int
	for i, nums := range matrix {
		sum = 0
		for j := 0; j < len(nums); j++ {
			sum += nums[j]
			if i == 0 {
				m.prefixSum[i][j] = sum
			} else {
				m.prefixSum[i][j] = sum + m.prefixSum[i-1][j]
			}
		}
	}
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.prefixSum[row2][col2]
	}
	if row1 == 0 {
		return this.prefixSum[row2][col2] - this.prefixSum[row2][col1-1]
	}
	if col1 == 0 {
		return this.prefixSum[row2][col2] - this.prefixSum[row1-1][col2]
	}
	return this.prefixSum[row2][col2] - this.prefixSum[row1-1][col2] - this.prefixSum[row2][col1-1] + this.prefixSum[row1-1][col1-1]
}

func Test13(t *testing.T) {
	var nums = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	var nums2 = [][]int{
		{-4, -5},
	}
	m := Constructor(nums)
	t.Log(m.prefixSum)
	t.Log(m.SumRegion(1, 1, 2, 2))

	m2 := Constructor(nums2)
	t.Log(m2.prefixSum)
	t.Log(m2.SumRegion(0, 1, 0, 1))
}

// [0,0,0,0],[0,0,0,1],[0,1,0,1]
