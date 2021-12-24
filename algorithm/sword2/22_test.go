package sword2

func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
