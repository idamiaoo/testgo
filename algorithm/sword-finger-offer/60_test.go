package sword_finger_offer

import (
	"testing"
)

func dicesProbability(n int) []float64 {
	return nil
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swap1(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b ^ *b
	*a = *a ^ *b
}

func TestSwap(t *testing.T) {
	a, b := -1, 9
	swap(&a, &b)
	t.Log(a, b)
}
