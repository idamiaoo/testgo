package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	var swapA, swapB bool
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
		if a == nil && !swapA {
			a = headB
			swapA = true
		}
		if b == nil && !swapB {
			b = headA
			swapB = true
		}

	}
	return nil
}
