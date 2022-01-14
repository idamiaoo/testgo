package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var nodes []*TreeNode
	var ans [][]int
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		var temp []int
		l := len(nodes)
		for _, v := range nodes {
			temp = append(temp, v.Val)
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}
		ans = append(ans, temp)
		nodes = nodes[l:]
	}
	return ans
}
