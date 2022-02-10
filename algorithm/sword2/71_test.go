package sword2

import (
	"math/rand"
	"sort"
)

type Solution struct {
	pre []int
}

func Constructor71(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{
		pre: w,
	}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
