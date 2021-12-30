package sword_finger_offer

import (
	"testing"
)

func levelOrder32(root *TreeNode) [][]int {
	var ans [][]int
	var nodes []*TreeNode
	nodes = append(nodes, root)
	var floor int
	for len(nodes) > 0 {
		floor++
		l := len(nodes)
		var temp []int
		if floor%2 == 1 {
			for i := 0; i < l; i++ {
				if nodes[i] != nil {
					temp = append(temp, nodes[i].Val)
					if nodes[i].Left != nil {
						nodes = append(nodes, nodes[i].Left)
					}
					if nodes[i].Right != nil {
						nodes = append(nodes, nodes[i].Right)
					}
				}
			}
		} else {
			for i := 0; i < l; i++ {
				if nodes[i] != nil {
					if nodes[i].Left != nil {
						nodes = append(nodes, nodes[i].Left)
					}
					if nodes[i].Right != nil {
						nodes = append(nodes, nodes[i].Right)
					}
				}
			}
			for i := l - 1; i >= 0; i-- {
				if nodes[i] != nil {
					temp = append(temp, nodes[i].Val)
				}
			}
		}
		nodes = nodes[l:]
		if len(temp) > 0 {
			ans = append(ans, temp)
		}
	}
	return ans
}

func Test32(t *testing.T) {
	var a = 0.7
	t.Log(int(a))
}
