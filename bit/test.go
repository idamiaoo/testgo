package main

import (
	"fmt"
	"go/cmkj_server/util"
	"time"
)

type combinations struct {
	comb    []int
	n       int
	k       int
	didnull bool
}

func newCombainations(n, k int) *combinations {
	com := new(combinations)
	com.n = n
	com.k = k
	com.didnull = true
	com.comb = make([]int, k)
	for i := 0; i < k; i++ {
		com.comb[i] = i
	}
	return com
}

func (c *combinations) next() bool {
	if c.k == 0 {
		return !c.didnull
	}
	ii := c.k - 1
	for ii >= 0 && c.comb[ii]+c.k+1 > c.n+ii {
		ii--
	}
	if ii < 0 {
		return false
	}
	c.comb[ii]++
	ii++
	for ii < c.k {
		c.comb[ii] = c.comb[ii-1] + 1
		ii++
	}
	return true
}

func print(data []int) {
	util.Log.Debug(data)
}

func main() {
	var base uint64 = 0xffffffffffffffff

	var mark uint64
	mark = ^(base << 52)
	fmt.Println(mark)

	v := mark
	var c int
	for c = 0; v != 0; c++ {
		v &= v - 1
	}
	fmt.Println(c)
	start := time.Now()
	var hub [][]int
	com := newCombainations(5, 3)
	tmp := make([]int, 2)
	print(com.comb)
	copy(tmp, com.comb)
	hub = append(hub, tmp)
	for com.next() {
		print(com.comb)
		copy(tmp, com.comb)
		hub = append(hub, tmp)
	}
	util.Log.Debug(len(hub))
	util.Log.Debug(time.Now().Sub(start))
}
