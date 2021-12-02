package sword_finger_offer

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if nodeExist(root.Left, p) {
		if nodeExist(root.Left, q) {
			return lowestCommonAncestor(root.Left, p, q)
		}
	} else {
		if !nodeExist(root.Left, q) {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}
	return root
}

func nodeExist(root, n *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == n.Val {
		return true
	}

	return nodeExist(root.Left, n) || nodeExist(root.Right, n)
}
