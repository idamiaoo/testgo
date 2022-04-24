package leetcode

type Solution struct {
	m, n   int
	matrix [][]int
}

func Constructor(m int, n int) Solution {
	s := Solution{
		m: m,
		n: n,
	}
	s.matrix = make([][]int, n)
	for i := range s.matrix {
		s.matrix[i] = make([]int, m)
	}
	return s
}

func (this *Solution) Flip() []int {
	return nil
}

func (this *Solution) Reset() {
	for i := range this.matrix {
		for j := range this.matrix[i] {
			this.matrix[i][j] = 0
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
