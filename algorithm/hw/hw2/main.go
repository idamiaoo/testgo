package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var (
		s    string
		nums []string
	)
	fmt.Scan(&s)
	nums = append(nums, strings.Split(s, ",")...)
	fmt.Println(minNumber(nums, 3))
}

func minNumber(nums []string, n int) int {
	var (
		ans         = math.MaxInt64
		path        []string
		backtracing func(int)
		choice      = make(map[int]bool)
	)

	backtracing = func(count int) {
		if count == n || len(path) == len(nums) {
			var s string
			for _, v := range path {
				s += v
			}
			num, _ := strconv.Atoi(s)
			if num < ans {
				ans = num
			}
			return
		}
		for i := range nums {
			if choice[i] {
				continue
			}
			choice[i] = true
			path = append(path, nums[i])
			backtracing(count + 1)
			path = path[:len(path)-1]
			choice[i] = false
		}
	}

	backtracing(0)
	return ans
}
