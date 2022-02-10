package sword2

import (
	"testing"
)

func insert(aNode *Node, x int) *Node {
	n := &Node{Val: x}
	if aNode == nil {
		n.Next = n
		return n
	}

	var h = aNode

	for h.Next != aNode {
		if h.Val > h.Next.Val {
			if x < h.Next.Val || x > h.Val {
				break
			}
		}
		if h.Val <= x && x <= h.Next.Val {
			break
		}
		h = h.Next
	}
	h.Next = &Node{
		Val:  x,
		Next: h.Next,
	}
	return aNode
}

func Test29(t *testing.T) {
	n1 := &Node{Val: 3}
	n2 := &Node{Val: 5}
	n3 := &Node{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	insert(n1, 6)
}
