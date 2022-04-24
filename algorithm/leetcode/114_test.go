package leetcode

func flatten(root *TreeNode) {
	var head = &TreeNode{}
	var tail = head
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		tail.Left = root
		tail = root
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	for head != nil {
		head.Right = head.Left
		head.Left = nil
		head = head.Right
	}
}
