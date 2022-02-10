package sword_finger_offer

import (
	"testing"
)

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	return 0
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	return pow(a*a, b/2) * pow(a, b%2)
}

func TestPow1(t *testing.T) {
	t.Log(pow(2, 5))
}
