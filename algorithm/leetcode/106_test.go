package leetcode

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	search := func(nums []int, num int) int {
		for i, v := range nums {
			if v == num {
				return i
			}
		}
		return -1
	}
	index := search(inorder, postorder[len(postorder)-1])
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
