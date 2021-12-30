package sword2

import (
	"testing"
)

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(node *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func Test54(t *testing.T) {
	codec := Constructor48()

	root := codec.deserialize("4,1,6,0,2,5,7,null,null,null,3,null,null,null,8")
	t.Log(codec.serialize(convertBST(root)))
}
