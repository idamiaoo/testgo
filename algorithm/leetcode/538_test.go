package leetcode

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		root.Val += sum
		sum = root.Val
		helper(root.Left)
	}
	helper(root)
	return root
}
