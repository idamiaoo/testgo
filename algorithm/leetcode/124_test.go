package leetcode

import (
	"math"
)

func maxPathSum(root *TreeNode) int {
	var ans = math.MinInt32
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		right := helper(root.Right)
		left := helper(root.Left)

		if right+left+root.Val > ans {
			ans = right + left + root.Val
		}

		if right+root.Val > ans {
			ans = right + root.Val
		}

		if left+root.Val > ans {
			ans = left + root.Val
		}

		if root.Val > ans {
			ans = root.Val
		}

		if left < 0 && right < 0 {
			return root.Val
		}

		if left > right {
			return left + root.Val
		}
		return right + root.Val
	}
	helper(root)
	return ans
}
