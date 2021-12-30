package leetcode

import (
	"testing"
)

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

func Test114(t *testing.T) {
	codec := Constructor48()
	t1 := codec.deserialize("1,2,5,3,4,null,6")
	flatten(t1)
	t.Log(codec.serialize(t1))
}
