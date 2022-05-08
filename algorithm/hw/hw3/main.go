package main

import (
	"fmt"
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

	for right < len(s) {
		if (yuan[s[right]] && right < len(s)-1 && yuan[s[right+1]]) || !yuan[s[right]] {
			right++
			continue
		}

		for !yuan[s[left]] {
			left++
		}
		var sum int
		for i := left + 1; i < right; i++ {
			if !yuan[s[i]] {
				sum++
			}
		}

		if sum == target {
			ans = max(ans, right-left+1)
			for left < right && yuan[s[left]] {
				left++
			}
			right++
		} else if sum > target {
			for left < right && yuan[s[left]] {
				left++
			}
			if left == right {
				right++
			}
		} else {
			right++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
