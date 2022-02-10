package sword2

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var a, b = headA, headB
	for a != nil && b != nil {
		if a == b {
			return a
		}
		if a.Next == nil {
			a = headB
			headB = nil
		} else {
			a = a.Next
		}
		if b.Next == nil {
			b = headA
			headA = nil
		} else {
			b = b.Next
		}
	}
	return nil
}
