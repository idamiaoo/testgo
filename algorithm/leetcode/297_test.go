package leetcode

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var nodes []*TreeNode
	var ans []string
	nodes = append(nodes, root)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			ans = append(ans, "null")
		} else {
			ans = append(ans, strconv.Itoa(nodes[i].Val))
			nodes = append(nodes, nodes[i].Left)
			nodes = append(nodes, nodes[i].Right)
		}
	}

	var i int
	for i = len(ans) - 1; i >= 1; i-- {
		if ans[i] != "null" {
			break
		}
	}
	return strings.Join(ans[:i+1], ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "null" {
		return nil
	}
	res := strings.Split(data, ",")
	var nodes []*TreeNode
	for _, v := range res {
		if v == "null" {
			nodes = append(nodes, nil)
		} else {
			n, _ := strconv.Atoi(v)
			nodes = append(nodes, &TreeNode{Val: n})
		}
	}

	if len(nodes)%2 == 0 {
		nodes = append(nodes, nil)
	}
	for p, c := 0, 1; c < len(nodes); p++ {
		if nodes[p] != nil {
			nodes[p].Left = nodes[c]
			nodes[p].Right = nodes[c+1]
			c += 2
		}
	}
	return nodes[0]
}
