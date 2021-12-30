package sword2

import (
	"testing"
)

func sumNumbers(root *TreeNode) int {
	var path []int
	var sum int

	var backtrack func(node *TreeNode)

	backtrack = func(root *TreeNode) {
		path = append(path, root.Val)

		if root.Left == nil && root.Right == nil {
			var temp int
			for i := 0; i < len(path); i++ {
				temp = temp*10 + path[i]
			}
			sum += temp
			return
		}

		l := len(path)
		if root.Left != nil {
			backtrack(root.Left)
		}
		if root.Right != nil {
			path = path[:l]
			backtrack(root.Right)
		}
		path = path[:l-1]
	}
	backtrack(root)
	return sum
}

func Test49(t *testing.T) {
	codec := Constructor48()

	t1 := codec.deserialize("1,2,3")
	t.Log(sumNumbers(t1))
}
