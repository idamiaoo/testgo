package sword2

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	h := l
	var s string
	for h != nil {
		if s == "" {
			s = fmt.Sprintf("%d", h.Val)
		} else {
			s = fmt.Sprintf("%s,%d", s, h.Val)
		}
		h = h.Next
	}
	return s
}

func buildList(nums []int) *ListNode {
	var head, tail *ListNode
	for _, v := range nums {
		n := &ListNode{Val: v}
		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}
	return head
}

func buildLists(nums [][]int) []*ListNode {
	var lists []*ListNode
	for _, v := range nums {
		lists = append(lists, buildList(v))
	}
	return lists
}
