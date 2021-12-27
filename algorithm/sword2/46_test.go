package sword2

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		res = append(res, nodes[l-1].Val)
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
