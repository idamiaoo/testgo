package sword2

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	var h = root
	for h != nil {
		if h.Child != nil {
			child := flatten(h.Child)
			h.Child = nil
			next := h.Next
			h.Next = child
			child.Prev = h

			if next != nil {
				for child.Next != nil {
					child = child.Next
				}
				child.Next = next
				next.Prev = child
			}
			h = next
		} else {
			h = h.Next
		}
	}
	return root
}
