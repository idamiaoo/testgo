package sword_finger_offer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	h1 := head
	var h2, nHead *Node
	for h1 != nil {
		if h2 == nil {
			h2 = &Node{
				Val: h1.Val,
			}
			nHead = h2
			h1 = h1.Next
		} else {
			h2.Next = &Node{
				Val: h1.Val,
			}
			h1 = h1.Next
			h2 = h2.Next
		}
	}

	h3 := head
	h4 := nHead
	for h3 != nil {
		if h3.Random != nil {
			searchNode(head, h3.Random, nHead, h4)
		}
		h3 = h3.Next
		h4 = h4.Next
	}
	return nHead
}

func searchNode(h1, n1, h2, n2 *Node) {
	for h1 != nil {
		if h1 == n1 {
			n2.Random = h2
			return
		} else {
			h1 = h1.Next
			h2 = h2.Next
		}
	}
}
