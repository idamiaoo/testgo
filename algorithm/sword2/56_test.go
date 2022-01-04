package sword2

func findTarget(root *TreeNode, k int) bool {
	values := make(map[int]bool)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		values[root.Val] = true
		helper(root.Right)
	}
	helper(root)
	for v, _ := range values {
		if k-v != v && values[k-v] {
			return true
		}
	}
	return false
}
