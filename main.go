package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		target int
		s      string
	)
	fmt.Scan(&target)
	fmt.Scan(&s)
	fmt.Println(maxSubStringLength(s, target))
}

func maxSubStringLength(s string, target int) int {
	yuan := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	var left, right int
	var ans int
	var sum int

	for right < len(s) {
		if yuan[s[right]] && ((right < len(s)-1 && !yuan[s[right+1]]) || right == len(s)-1) {
			for !yuan[s[left]] {
				sum--
				left++
			}
			if sum >= target {
				if sum == target {
					ans = max(ans, right-left+1)
				}
				for left < right && yuan[s[left]] {
					left++
				}
				if left < right {
					continue
				}
			}
		}
		if !yuan[s[right]] {
			sum++
		}
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// aavaaavaaaa

func minSender(address []int) int {
	var (
		ans          = math.MaxInt64
		left         int
		backtracting func(int)
		choice       = make(map[int]bool)
	)
	backtracting = func(cnt int) {
		if cnt == len(address) || left == 0 {
			if cnt < ans {
				ans = cnt
			}
			return
		}
		if cnt >= ans {
			return
		}
		for i := 0; i < len(address); i++ {
			if choice[i] {
				continue
			}
			choice[i] = true
			var pre, cur, next bool
			if address[i] == 1 {
				address[i] = 0
				cur = true
				left--
			}
			if i > 0 && address[i-1] == 1 {
				address[i-1] = 0
				pre = true
				left--
			}
			if i < len(address)-1 && address[i+1] == 1 {
				address[i+1] = 0
				next = true
				left--
			}
			backtracting(cnt + 1)
			if cur {
				address[i] = 1
				left++
			}
			if pre {
				address[i-1] = 1
				left++
			}
			if next {
				address[i+1] = 1
				left++
			}
			choice[i] = false
		}

	}
	for _, v := range address {
		if v == 1 {
			left++
		}
	}
	backtracting(0)
	return ans

}

func buy() {

}
