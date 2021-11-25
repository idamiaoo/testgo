package sword_finger_offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if isSub(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if A.Val == B.Val {
		if B.Left == nil && B.Right == nil {
			return true
		} else if B.Left != nil && B.Right != nil {
			return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
		} else if B.Left != nil {
			return isSub(A.Left, B.Left)
		} else if B.Right != nil {
			return isSub(A.Right, B.Right)
		}
	}
	return false
}
