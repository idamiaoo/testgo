package sword2

import (
	"testing"
)

/*
[[-8,-7,-7,-5,1,1,3,4],[-2],[-10,-10,-7,0,1,3],[2]]

[-10,-10,-7,-2,0,-8,-7,-7,-5,1,1,1,2,3,3,4]
*/
func mergeKLists(lists []*ListNode) *ListNode {
	heap := &NodeHeap{}
	for _, h := range lists {
		heap.add(h)
	}

	var head = &ListNode{}
	cur := head
	for {
		cur.Next = heap.del()
		if cur.Next == nil {
			break
		}
		heap.add(cur.Next.Next)
		cur = cur.Next
	}
	return head.Next
}

type NodeHeap struct {
	cap   int
	size  int
	nodes []*ListNode
}

func (h *NodeHeap) add(n *ListNode) {
	if n == nil {
		return
	}
	if h.size >= len(h.nodes) {
		h.nodes = append(h.nodes, n)
	} else {
		h.nodes[h.size] = n
	}
	h.size++
	h.up(h.size - 1)
}

func (h *NodeHeap) del() *ListNode {
	if h.size == 0 {
		return nil
	}
	node := h.nodes[0]
	h.swap(0, h.size-1)
	h.size--
	h.down(0)
	return node
}

func (h *NodeHeap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// 0 1 2 3 4 5
func (h *NodeHeap) up(i int) {
	for {
		x := (i - 1) / 2
		if x >= i || h.nodes[i].Val >= h.nodes[x].Val {
			return
		}
		h.swap(i, x)
		i = x
	}
}

func (h *NodeHeap) down(i int) {
	for {
		x := 2*i + 1
		if x >= h.size {
			return
		}
		if x+1 <= h.size-1 && h.nodes[x+1].Val < h.nodes[x].Val {
			x = x + 1
		}
		if h.nodes[i].Val < h.nodes[x].Val {
			return
		}
		h.swap(i, x)
		i = x
	}
}

func Test78(t *testing.T) {
	nums := [][]int{{-8, -7, -7, -5, 1, 1, 3, 4}, {-2}, {-10, -10, -7, 0, 1, 3}, {2}}
	t.Log(mergeKLists(buildLists(nums)))
}
