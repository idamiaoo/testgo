package leetcode

import (
	"testing"
)

// 9 * 10
// step1: 找出区间
// step2: 找出具体数字，以及在该数字的第几位
// step3: 按位找数字
func findNthDigit(n int) int {
	var x = 1
	var s = 0
	for {
		if x*9*pow(10, x-1) >= n {
			break
		} else {
			s += x * 9 * pow(10, x-1)
			x++
		}
	}
	if x == 1 {
		return n
	}

	d := (n - s) / x
	m := (n - s) % x
	if m == 0 {
		d = d - 1
		m = x
	}
	p := pow(10, x-1) + d

	var i = 1
	for {
		if i == m {
			return p / pow(10, x-i)
		}
		p = p % pow(10, x-i)
		i++
	}
}

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}
	if n > 1 {
		return pow(x*x, n/2) * pow(x, n%2)
	}
	return 1 / pow(x, -n)
}

func TestFindNthDigit(t *testing.T) {
	t.Log(findNthDigit(17))
}
