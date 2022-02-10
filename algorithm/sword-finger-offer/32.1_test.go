package sword_finger_offer

func levelOrder(root *TreeNode) []int {
	var ans []int
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			ans = append(ans, nodes[i].Val)
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
	}
	return ans
}
