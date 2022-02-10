package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nh := &ListNode{Next: head}
	slow, fast := nh, nh
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return nh.Next
}
