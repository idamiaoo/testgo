package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	mid := searchMid(head)
	start2 := reverseList(mid.Next)

	for start2 != nil {
		if head.Val != start2.Val {
			return false
		}
		head = head.Next
		start2 = start2.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var temp, newHead *ListNode
	for head != nil {
		temp = head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

func searchMid(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
