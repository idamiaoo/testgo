package sword_finger_offer

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left = treeToDoublyList(root.Left)
	root.Right = treeToDoublyList(root.Right)

	return root
}
