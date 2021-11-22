package sword_finger_offer

import (
	"testing"
)

func myPow2(x float64, n int) float64 {
	if n >= 0 {
		var v = 1.0
		for n > 0 {
			v *= x
			n--
		}
		return v
	}
	return 1 / myPow2(x, -n)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n > 0 {
		return myPow(x*x, n/2) * myPow(x, n%2)
	}
	return 1 / myPow(x, -n)
}

func myPow3(x float64, n int) float64 {
	if n >= 0 {
		var v = 1.0
		for n > 0 {
			v *= x
			n--
		}
		return v
	}
	return 1 / myPow3(x, -n)
}
func TestPow(t *testing.T) {
	t.Log(myPow2(2, 2))
}

//2147483647

func BenchmarkPow2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myPow2(0.00001, 2147483647)
	}
	b.StopTimer()
}

func BenchmarkPow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myPow(0.00001, 2147483647)
	}
	b.StopTimer()
}
