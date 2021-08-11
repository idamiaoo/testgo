package mike

import (
	"testing"
)

func tree1() *Node {
	root := &Node{}
	root.Right = &Node{}
	return root
}

func tree2() *Node {
	root := &Node{}
	root.Left = &Node{}
	root.Right = &Node{}
	return root
}

func tree3() *Node {
	root := &Node{}
	root.Left = &Node{}
	root.Right = &Node{}
	root.Right.Left = &Node{}
	root.Right.Left.Left = &Node{}
	root.Right.Left.Left.Left = &Node{}
	root.Right.Left.Left.Left.Left = &Node{}
	root.Right.Right = &Node{}
	root.Right.Right.Right = &Node{}

	return root
}
func TestMaxDistance(t *testing.T) {
	t.Log(MaxDistance(tree3()))
}
