package sword2

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	var temp *ListNode

	for head != nil {
		temp = head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sub := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return sub
}
