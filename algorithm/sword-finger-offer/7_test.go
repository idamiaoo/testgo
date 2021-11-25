package sword_finger_offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}

	var inorderRootIndex int
	for i, v := range inorder {
		if v == preorder[0] {
			inorderRootIndex = i
		}
	}
	if inorderRootIndex == 0 {
		node.Right = buildTree(preorder[1:], inorder[1:])
	} else {
		node.Left = buildTree(preorder[1:1+inorderRootIndex], inorder[0:inorderRootIndex])
		node.Right = buildTree(preorder[1+inorderRootIndex:], inorder[inorderRootIndex+1:])
	}
	return node
}
