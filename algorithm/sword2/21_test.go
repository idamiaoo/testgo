package sword2

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var h1, h2 = head, head
	for h2 != nil {
		if n >= 0 {
			h2 = h2.Next
			n--
		} else {
			h1 = h1.Next
			h2 = h2.Next
		}
	}

	if n == 0 {
		head = head.Next
	} else {
		h1.Next = h1.Next.Next
	}
	return head
}
