package test

import (
	"fmt"
	"testing"
)

func TestMoney(t *testing.T) {
	all := 10000
	arr := [6]int{1, 5, 10, 20, 50, 100}
	_len := len(arr)
	res := make([][]int, _len+1, _len+1)
	for i := range res {
		res[i] = make([]int, all+1, all+1)
	}
	for i := 0; i <= _len; i++ {
		res[i][0] = 1
	}
	for j := 1; j <= all; j++ {
		res[0][j] = 0
	}
	for i := 1; i <= _len; i++ {
		for j := 1; j <= all; j++ {
			res[i][j] = 0
			for k := 0; k <= j/arr[i-1]; k++ {
				res[i][j] += res[i-1][j-k*arr[i-1]]
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
