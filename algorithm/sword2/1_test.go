package sword2

import (
	"fmt"
	"testing"
)

func divide(a int, b int) int {
	var symbol int
	if ((a>>31)&1)^((b>>31)&1) == 0 {
		symbol = 1
	} else {
		symbol = -1
	}

	if a == 0 {
		return 0
	}
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	var s, sum int
	for {
		if sum == 0 {
			if sum+b < a {
				return 0
			}
			sum = b
			s = 1
		} else {
			if sum+sum < a {
				if symbol > 0 {
					v := s + divide(a-sum, b)
					if v == 1<<31 {
						fmt.Println(v)
						return (1 << 31) - 1
					}
					return v
				}
				return -(s + divide(a-sum, b))
			}
			sum += sum
			s += s
		}
	}
}

func Test1(t *testing.T) {
	// v2 := int(2147483647)
	// t.Log((v1 >> 31) & 1)
	// t.Log((v2 >> 31) & 1)
	// t.Logf("%b", v1)
	t.Log(divide(-(1 << 31), -1))
}
