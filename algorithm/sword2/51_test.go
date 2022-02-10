package sword2

import (
	"math"
	"testing"
)

func maxPathSum(root *TreeNode) int {
	var max = math.MinInt32

	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lv := dfs(root.Left)
		rv := dfs(root.Right)

		if lv+rv+root.Val > max {
			max = lv + rv + root.Val
		}
		if lv+root.Val > max {
			max = lv + root.Val
		}
		if rv+root.Val > max {
			max = rv + root.Val
		}
		if root.Val > max {
			max = root.Val
		}

		if lv < 0 && rv < 0 {
			return root.Val
		}

		if lv > rv {
			return lv + root.Val
		}
		return rv + root.Val
	}
	dfs(root)
	return max
}

func Test51(t *testing.T) {
	codec := Constructor48()
	t.Log(maxPathSum(codec.deserialize("2,-1,-2")))
}
