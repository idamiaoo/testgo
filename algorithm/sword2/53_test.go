package sword2

import (
	"testing"
)

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	for root != nil {
		if root.Val > p.Val {
			ans = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return ans
}

func Test53(t *testing.T) {
	codec := Constructor48()
	t1 := codec.deserialize("5,3,6,2,4,null,null,1")
	t.Log(inorderSuccessor(t1, &TreeNode{Val: 1}))
}
