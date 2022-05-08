package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	index := searchNum(inorder, preorder[0])
	node.Left = buildTree(preorder[1:1+index], inorder[:index])
	node.Right = buildTree(preorder[1+index:], inorder[index+1:])
	return node
}

func searchNum(nums []int, num int) int {
	for i, v := range nums {
		if v == num {
			return i
		}
	}
	return -1
}
