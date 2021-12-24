package sword2

func isPalindromeList(head *ListNode) bool {
	var n1, n2 = head, head
	for n2 != nil && n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	if n2 == nil {
		n2 = reverseList(n1)
	} else {
		n2 = reverseList(n1.Next)
	}
	for n2 != nil {
		if n2.Val != head.Val {
			return false
		}
		n2 = n2.Next
		head = head.Next
	}
	return true
}
