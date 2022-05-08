package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var head, p *Node
	var temp = root
	for temp.Left != nil {
		p = temp
		head = temp.Left
		for p != nil {
			if p.Left != head {
				head.Next = p.Left
				head = head.Next
			}
			head.Next = p.Right
			head = head.Next
			p = p.Next
		}
		temp = temp.Left
	}
	return root
}
