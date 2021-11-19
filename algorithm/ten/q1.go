package ten

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) String() string {
	var data []string
	for n != nil {
		data = append(data, strconv.Itoa(n.Val))
		n = n.Next
	}
	return strings.Join(data, "->")
}

func UpList(list *Node) *Node {
	var l1, l2 *Node
	var i int
	for list != nil {
		i++
		if i%2 != 0 {
			if l1 == nil {
				l1 = list
			} else {
				l1.Next = list
			}
		} else {
			if l2 == nil {
				l2 = list
			} else {
				l2.Next = list
			}
		}
	}
	l3 := revList(l2)

	return mergeList(l1, l3)
}

func revList(list *Node) *Node {
	var newHead *Node
	for list != nil {
		temp := list.Next
		list.Next = newHead
		newHead = list
		list = temp
	}
	return newHead
}

func mergeList(l1, l2 *Node) *Node {
	var newHead, newTail *Node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if newHead == nil {
				newHead = l1
			} else {
				newHead.Next = l1
			}
			newTail = l1
			l1 = l1.Next
		} else {
			if newHead == nil {
				newHead = l2
			} else {
				newHead.Next = l2
			}
			newTail = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		newTail.Next = l1
	}
	if l2 != nil {
		newTail.Next = l2
	}
	return newHead
}
