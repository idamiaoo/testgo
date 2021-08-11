package main

import (
	"fmt"
)

func singleNumbers(nums []int) []int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	var n int
	for (res>>n)&1 == 0 {
		n++
	}
	var res1, res2 int
	for _, v := range nums {
		if v&(1<<n) == 0 {
			res1 ^= v
		} else {
			res2 ^= v
		}
	}

	return []int{res1, res2}
}

func main() {
	nums := []int{9, 1, 7, 9, 7, 5}
	fmt.Println(singleNumbers(nums))
}
