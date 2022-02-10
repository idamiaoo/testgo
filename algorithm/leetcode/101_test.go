package leetcode

func isSymmetric(root *TreeNode) bool {
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(r1 *TreeNode, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		}
		if r1 == nil || r2 == nil {
			return false
		}
		return r1.Val == r2.Val && compare(r1.Left, r2.Right) && compare(r1.Right, r2.Left)
	}
	return compare(root, root)
}
