package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Right, root.Left = root.Left, root.Right
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
