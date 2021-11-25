package main

import "fmt"

func validateStackSequences(pushed, popped []int) {

}

func quickSort(array []int, left, right int) {
	if left < right {
		p := partition(array, left, right)
		quickSort(array, left, p-1)
		quickSort(array, p+1, right)
	}
}

func partition(array []int, left, right int) int {
	v := array[right]
	p := left - 1
	for i := left; i < right; i++ {
		if array[i] < v {
			p++
			array[p], array[i] = array[i], array[p]
		}
	}
	array[p+1], array[right] = array[right], array[p+1]
	return p + 1
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	nums := []int{3, 9, 6, 8, 4, 7, 9, 3, 4, 7, 5}
	quickSort(nums, 0, 10)
	fmt.Println(nums)

	a, b := 11, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
