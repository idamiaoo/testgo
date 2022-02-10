package sword2

import (
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		l := len(nodes)
		max := math.MinInt32
		for i := 0; i < l; i++ {
			if nodes[i].Val > max {
				max = nodes[i].Val
			}
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		res = append(res, max)
		nodes = nodes[l:]
	}
	return res
}
