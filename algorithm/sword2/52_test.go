package sword2

import (
	"testing"
)

func increasingBST(root *TreeNode) *TreeNode {
	var nodes []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if len(nodes) > 0 {
			nodes[len(nodes)-1].Right = root
		}
		root.Left = nil
		nodes = append(nodes, root)
		dfs(root.Right)
	}
	dfs(root)
	if len(nodes) > 0 {
		return nodes[0]
	}
	return nil
}

func Test52(t *testing.T) {
	codec := Constructor48()
	t1 := codec.deserialize("5,3,6,2,4,null,8,1,null,null,null,7,9")
	t.Log(codec.serialize(increasingBST(t1)))
}
