package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var res int
	for n := 0; n < 31; n++ {
		var sub int
		for _, v := range nums {
			if v&(1<<n) > 0 {
				sub++
			}
		}
		if sub%3 > 0 {
			res |= 1 << n
		}
	}
	return res
}

func main() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber(nums))
}
