package leetcode

import (
	"math"
	"testing"
)

func isValidBST(root *TreeNode) bool {
	var helper func(*TreeNode, int, int) bool
	helper = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
	}
	return helper(root, math.MaxInt64, math.MaxInt)
}

func Test98(t *testing.T) {

}
