package sword2

import (
	"testing"
)

func mySqrt(x int) int {
	var a, b = 1, x
	var mid int
	for a <= b {
		mid = a + ((b - a) >> 2)
		if mid <= x/mid {
			if (mid + 1) > x/(mid+1) {
				return mid
			}
			a = mid + 1
		} else {
			b = mid - 1
		}
	}
	return 0
}

func Test72(t *testing.T) {
	t.Log(mySqrt(8))
}
