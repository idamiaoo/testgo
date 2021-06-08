package sort

import (
	"fmt"
	"testing"
)

func quickSort(data []int) {
	base := data[0]
	p, head := 0, 0
	tail := len(data) - 1

	for head <= tail {
		for tail >= p && data[tail] >= base {
			tail--
		}
		if tail >= p {
			data[p] = data[tail]
			p = tail
		}

		for head <= p && data[head] <= base {
			head++
		}
		if head <= p {
			data[p] = data[head]
			p = head
		}
	}
	data[p] = base
	if p > 1 {
		quickSort(data[:p])
	}
	if p < len(data)-1 {
		quickSort(data[p+1:])
	}
}

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

func quickSortV2(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSortV2(values, left, p-1)
	}
	if right-p > 1 {
		quickSortV2(values, p+1, right)
	}
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
