package test

import (
	"fmt"
	"strings"
	"testing"
)

/*
1. 按快排的方式处理,但是找到 中间点后,比较 k 和 i+1 的大小
	k == i+1 return
	k < i+1  走快排左分支
	k > i+1 走快排右分支

	o (n)
*/

func stringToInt(s string) (error, int) {
	if s == "" {
		return fmt.Errorf("s is not number"), 0
	}
	var value int
	firstZero := true
	for _, v := range []byte(s) {
		if v < '0' || v > '9' {
			return fmt.Errorf("s is not number"), 0
		}

		if v == '0' && firstZero {
			continue
		}

		if firstZero == false {
			firstZero = true
		}
		value = 10*value + int(v-'0')
	}
	return nil, value
}

func printString(s string) {
	var result []string
	backtracking(s, "", &result)
	fmt.Println(result)
}

func backtracking(s, pick string, result *[]string) {
	if len(pick) == len(s) {
		*result = append(*result, pick)
		return
	}
	for _, c := range s {
		if strings.Contains(pick, string(c)) {
			continue
		}
		pick = pick + string(c)
		backtracking(s, pick, result)
		pick = pick[:len(pick)-1]
	}
}

func TestPrintString(t *testing.T) {
	printString("abdc")
}
