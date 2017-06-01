package silice

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

	ss := make([]int, 0, 6)
	printSlice(ss)

	/*
		s = append(s, 2, 3, 4)
		printSlice(s)
		s = append(s, 5)
		printSlice(s)
		s = append(s, 6)
		printSlice(s)
	*/
}
