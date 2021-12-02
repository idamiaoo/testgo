package sword_finger_offer

import (
	"testing"
)

func constructArr(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	b := make([]int, len(a))
	b[0] = 1

	temp := 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}

	for i := len(a) - 2; i >= 0; i-- {
		temp = temp * a[i+1]
		b[i] *= temp
	}
	return b
}

func TestConstructArr(t *testing.T) {
	t.Log(constructArr([]int{1, 2, 3, 4, 5}))
}
