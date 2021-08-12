package slice

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, s=%v\n", len(s), cap(s), s)
}

func Append() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2)
	printSlice(s)
	s = append(s, 3)
	printSlice(s)
	s = append(s, 4)
	printSlice(s)

	printSlice(s[4:])

	ss := make([]int, 6)
	_aa := ss[2:]
	printSlice(ss)
	printSlice(_aa)
	ss[2] = 3
	printSlice(ss)
	printSlice(_aa)
	/*
		s = append(s, 2, 3, 4)
		printSlice(s)
		s = append(s, 5)
		printSlice(s)
		s = append(s, 6)
		printSlice(s)
	*/
	printSlice(ss[:3])
}
