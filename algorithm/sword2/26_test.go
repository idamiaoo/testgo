package sword2

import (
	"testing"
)

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	h1 := head
	for h1.Next != nil {
		if h1.Next.Next == nil {
			temp := h1.Next
			h1.Next = nil
			h1 = temp
			break
		}
		h1 = h1.Next
	}
	temp := head.Next
	head.Next = h1
	head.Next.Next = temp
	reorderList(temp)
}

func Test26(t *testing.T) {
	l := buildList([]int{1, 2, 3})
	t.Log(l.String())
	reorderList(l)
	t.Log(l.String())
}
