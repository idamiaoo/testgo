package string

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := -1

	dict := make([]int, 256)
	for i := 0; i < len(dict); i++ {
		dict[i] = -1
	}

	chs := []byte(s)
	for i := 0; i < len(chs); i++ {
		if dict[chs[i]] > start {
			start = dict[chs[i]]
		}
		fmt.Println("start:", start)
		dict[chs[i]] = i
		if i-start > maxLen {
			maxLen = i - start
		}
	}
	return maxLen
}

func TestLen(t *testing.T) {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	a := int16(0xFFFF)
	fmt.Printf("%d\n", a)
}

/*
	0000 0000 0000 0001
	0 1 2 1 0 1 2 1
	3
	0 1 2 3 4 5 6 7

	0 1 2 3 2 1 0 1 2 3 2 1
	4
	0 1 2 3 4 5 6 7 8 9 10 11

    0 1 2 3 4 3 2 1

	0 1 0 1 0 1
    2
	1 2 3 4 5 6

								   1 2 3 4 5
	1                              1 2 4 6 8
*/
