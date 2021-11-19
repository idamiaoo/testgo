package sort

import (
	"fmt"
	"strconv"
	"testing"
)

func quickSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p-1)
		quickSort(nums, p+1, right)
	}
}

func partition(nums []int, left, right int) int {
	p := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] < p {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func TestQuickSort(t *testing.T) {
	a := []int{3, 9, 6, 8, 4, 7, 9, 3, 4, 7, 5}
	quickSort(a, 0, 10)
	t.Log(a)
}

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) String() string {
	var s string
	l := n
	for l != nil {
		if s != "" {
			s += "->" + strconv.Itoa(l.Val)
		} else {
			s = strconv.Itoa(l.Val)
		}
		l = l.Next
	}
	return s
}

func ReverseList(list *Node) *Node {
	var newList *Node
	var tmp *Node
	for list != nil {
		tmp = list.Next
		list.Next = newList
		newList = list
		list = tmp
	}
	return newList
}

func TestReverseList(t *testing.T) {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val: 2,
	}
	n3 := &Node{
		Val: 3,
	}
	n4 := &Node{
		Val: 4,
	}
	n5 := &Node{
		Val: 5,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	t.Log(n1)
	t.Log(ReverseList(n1))
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

/*     ^ 10
0 1    & 01
1 1    | 11
*/

func Swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
