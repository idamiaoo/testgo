package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	var newHead *ListNode
	var newTail *ListNode
	var tempHead *ListNode
	var tempTail *ListNode
	var temp *ListNode
	var i int
	for head != nil {
		if tempTail == nil {
			tempTail = head
		}
		temp = head.Next
		head.Next = tempHead
		tempHead = head
		head = temp
		i++
		if i == k {
			if newHead == nil {
				newHead = tempHead
			} else {
				newTail.Next = tempHead
			}
			newTail = tempTail
			tempHead = nil
			tempTail = nil
			i = 0
		}
	}

	if i > 0 {
		for tempHead != nil {
			temp := tempHead.Next
			tempHead.Next = newTail.Next
			newTail.Next = tempHead
			tempHead = temp
		}
	}

	return newHead
}

var (
	n1 = &ListNode{
		Val: 1,
	}
	n2 = &ListNode{
		Val: 2,
	}
	n3 = &ListNode{
		Val: 3,
	}
	n4 = &ListNode{
		Val: 4,
	}
	n5 = &ListNode{
		Val: 5,
	}
)

func main() {
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	reverseKGroup(n1, 2)
}
