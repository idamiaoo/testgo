package sword_finger_offer

func levelOrder(root *TreeNode) []int {
	var ans []int

	var printTree func(root *TreeNode)

	printTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		printTree(root.Left)
	}

	if root == nil {
		return nil
	}
	return nil
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+deep(root.Left), 1+deep(root.Right))
}
