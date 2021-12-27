package sword2

func findBottomLeftValue(root *TreeNode) int {
	var res int
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		res = nodes[0].Val
		l := len(nodes)
		for i := 0; i < l; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		nodes = nodes[l:]
	}
	return res
}
