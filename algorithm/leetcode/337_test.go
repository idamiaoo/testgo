package leetcode

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		s := node.Val + left[1] + right[1]
		ns := max337(left[0], left[1]) + max337(right[0], right[1])
		return []int{s, ns}
	}
	val := dfs(root)
	return max337(val[0], val[1])
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}
