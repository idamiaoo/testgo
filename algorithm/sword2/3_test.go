package sword2

import (
	"testing"
)

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	base := 1
	for i := 2; i <= n; i++ {
		if i == 2*base {
			res[i] = 1
			base *= 2
		} else {
			res[i] = 1 + res[i-base]
		}
	}
	return res
}

func Test3(t *testing.T) {
	t.Log(countBits(5))
}
