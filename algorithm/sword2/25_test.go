package sword2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2 []int
	for l1 != nil {
		v1 = append(v1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		v2 = append(v2, l2.Val)
		l2 = l2.Next
	}
	n := len(v1)
	if len(v2) > len(v1) {
		n = len(v2)
	}
	var sum, carry int
	var l3 *ListNode
	for i := 0; i < n; i++ {
		sum += carry
		if i < len(v1) {
			sum += v1[len(v1)-1-i]
		}
		if i < len(v2) {
			sum += v2[len(v2)-1-i]
		}
		carry = sum / 10
		sum = sum % 10

		n := &ListNode{Val: sum}
		n.Next = l3
		l3 = n
		sum = 0
	}
	if carry > 0 {
		n := &ListNode{Val: carry}
		n.Next = l3
		l3 = n
	}
	return l3
}
