package leetcode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var helper func(root1 *TreeNode, root2 *TreeNode) *TreeNode
	helper = func(root1 *TreeNode, root2 *TreeNode) *TreeNode {
		if root1 == nil && root2 == nil {
			return nil
		}
		if root1 != nil && root2 != nil {
			root1.Val += root2.Val
			root1.Left = mergeTrees(root1.Left, root2.Left)
			root1.Right = mergeTrees(root1.Right, root2.Right)
			return root1
		}
		if root1 != nil {
			return root1
		}
		return root2
	}
	return helper(root1, root2)
}
