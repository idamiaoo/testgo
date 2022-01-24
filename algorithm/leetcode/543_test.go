package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		if left+right+1 > max {
			max = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	helper(root)
	return max
}
