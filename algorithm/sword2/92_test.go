package sword2

import (
	"testing"
)

func minFlipsMonoIncr(s string) int {
	var one, zero int
	if s[0] == '1' {
		zero = 1
	} else {
		one = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			one, zero = min(one, zero), zero+1
		} else {
			one = min(one, zero) + 1
		}
	}
	return min(one, zero)
}

func Test92(t *testing.T) {
	t.Log(minFlipsMonoIncr("00011000"))
}
