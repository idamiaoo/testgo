package sword2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodes *NodesTack
}

func Constructor55(root *TreeNode) BSTIterator {
	nodes := NewNodesTack()
	for root != nil {
		nodes.Put(root)
		root = root.Left
	}
	return BSTIterator{nodes: nodes}
}

func (this *BSTIterator) Next() int {
	node := this.nodes.Pop()
	v := node.Val
	node = node.Right
	for node != nil {
		this.nodes.Put(node)
		node = node.Left
	}
	return v
}

func (this *BSTIterator) HasNext() bool {
	return !this.nodes.Empty()
}

type NodesTack struct {
	top   int
	nodes []*TreeNode
}

func NewNodesTack() *NodesTack {
	return &NodesTack{
		top: -1,
	}
}

func (s *NodesTack) Pop() *TreeNode {
	node := s.nodes[s.top]
	s.top--
	return node
}

func (s *NodesTack) Put(node *TreeNode) {
	s.top++
	if len(s.nodes) <= s.top {
		s.nodes = append(s.nodes, node)
	} else {
		s.nodes[s.top] = node
	}
}

func (s *NodesTack) Empty() bool {
	return s.top == -1
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
