package main

import "fmt"

func main() {
	sli1 := []int{5, 4, 3, 2, 1}
	fmt.Println(len(sli1), cap(sli1), sli1)
	sli2 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(sli2), cap(sli2), sli2)
	sli3 := make([]byte, 10)
	str := "hello lunar"
	copy(sli1[2:], sli2)
	copy(sli3, str)
	fmt.Println(cap(sli1), sli1)
	fmt.Println(cap(sli3), sli3)
}
