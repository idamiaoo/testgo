package sword2

type CBTInserter struct {
	root  *TreeNode
	nodes []*TreeNode
}

func Constructor43(root *TreeNode) CBTInserter {
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for {
		if nodes[0].Left == nil || nodes[0].Right == nil {
			break
		}
		nodes = append(nodes, nodes[0].Left)
		nodes = append(nodes, nodes[0].Right)
		nodes = nodes[1:]
	}
	return CBTInserter{
		root:  root,
		nodes: nodes,
	}
}

func (this *CBTInserter) Insert(v int) int {
	for {
		if this.nodes[0].Left == nil || this.nodes[0].Right == nil {
			break
		}
		this.nodes = append(this.nodes, this.nodes[0].Left)
		this.nodes = append(this.nodes, this.nodes[0].Right)
		this.nodes = this.nodes[1:]
	}

	if this.nodes[0].Left == nil {
		this.nodes[0].Left = &TreeNode{Val: v}
	} else {
		this.nodes[0].Right = &TreeNode{Val: v}
	}
	return this.nodes[0].Val

}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
