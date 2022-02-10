package sword2

import (
	"testing"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	sub := len(a) - len(b)

	bytesA := []byte(a)
	bytesB := []byte(b)

	var carry byte
	for i := len(bytesB) - 1; i >= 0; i-- {
		v := (bytesA[i+sub] - '0') + (bytesB[i] - '0') + carry
		if v < 2 {
			carry = 0
			bytesA[i+sub] = '0' + v
		} else {
			carry = 1
			bytesA[i+sub] = '0' + (v - 2)
		}
	}
	for carry > 0 && sub > 0 {
		v := bytesA[sub-1] - '0' + carry
		if v < 2 {
			carry = 0
			bytesA[sub-1] = '0' + v
		} else {
			carry = 1
			bytesA[sub-1] = '0' + (v - 2)
		}
		sub--
	}

	if carry > 0 {
		return "1" + string(bytesA)
	}
	return string(bytesA)
}

func Test2(t *testing.T) {
	a, b := "1", "111"
	t.Log(addBinary(a, b))
}
