package sword2

import (
	"testing"
)

func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		if root.Val == target {
			ans++
		}
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
	}

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root, targetSum)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return ans
}

func Test50(t *testing.T) {
	codec := Constructor48()
	s := "-8,6,8,null,null,8,2,null,null,null,-2"
	target := -2
	t1 := codec.deserialize(s)
	t.Log(pathSum(t1, target))
}
