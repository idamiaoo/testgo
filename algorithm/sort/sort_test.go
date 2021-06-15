package sort

import (
	"fmt"
	"testing"
)


// [1,4] [2,3]
func mergeSort(m, n []int) []int {
	var temp []int
	i, j := 0, 0
	for i < len(m) && j < len(n) {
		if m[i] < n[j] {
			temp = append(temp, m[i])
			i++
		} else {
			temp = append(temp, n[j])
			j++
		}
	}
	if i < len(m) {
		temp = append(temp, m[i:]...)
	}

	if j < len(n) {
		temp = append(temp, n[j:]...)
	}
	return temp
}

func TestMerge(t *testing.T) {
	m := []int{1, 4}
	n := []int{2, 3}

	fmt.Println(mergeSort(m, n))
}


func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

func TestInsert(t *testing.T) {
	a := []int{5, 4, 2, 6, 9, 87, 3, 46, 10}
	insertionSort(a)
	fmt.Println(a)
}

func TestQuick(t *testing.T) {
	a := []int{4, 5, 2, 6, 9, 87, 3, 46, 10, 11, 1}
	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}

func TestQuickV2(t *testing.T) {
	a := []int{4, 5, 2, 6, 9, 87, 3, 46, 10, 11, 1}
	fmt.Println(a)
	quickSortV2(a, 0, len(a)-1)
	fmt.Println(a)
}

/*
4  5  2  6, 9, 87, 3, 46, 10,11
4  2  5  6, 9, 87, 3, 46, 10,11
4  2  3  6, 9, 87, 5, 46, 10,11
*/

func swap(a, b *int) {
	*a, *b = *b, *a
}

func TestPlus(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
