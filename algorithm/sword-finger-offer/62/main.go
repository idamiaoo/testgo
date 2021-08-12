package main

import (
	"fmt"
	"time"
)

func lastRemaining(n int, m int) int {
	nums := make([]int, n)
	var i, x, c int
	for {
		if nums[i] == 0 {
			c += 1
			if c == m {
				fmt.Println(i)
				nums[i] = 1
				x += 1
				if x == n {
					return i
				}
				c = 0
			}
		}

		i += 1
		if i == n {
			i = 0
		}
	}
}

type node struct {
	v    int
	next *node
}

func lastRemaining2(n int, m int) int {

	nums := make([]int, n)
	var i, x, c int
	for {
		if nums[i] == 0 {
			c += 1
			if c == m {
				fmt.Println(i)
				nums[i] = 1
				x += 1
				if x == n {
					return i
				}
				c = 0
			}
		}

		i += 1
		if i == n {
			i = 0
		}
	}
}

func main() {
	begin := time.Now()
	fmt.Println(lastRemaining(10, 17))
	fmt.Println("cost: ", time.Since(begin))
}
