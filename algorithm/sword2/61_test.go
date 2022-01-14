package sword2

import (
	"container/heap"
	"testing"
)

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct {
	i, j int
}
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int {
	return len(h.data)
}

func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

type hp1 struct {
	data []int
}

func (h hp1) Len() int {
	return len(h.data)
}

func (h hp1) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}
func (h hp1) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *hp1) Push(v interface{}) {
	h.data = append(h.data, v.(int))
}
func (h *hp1) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

func Test61(t *testing.T) {
	n1 := []int{1, 1, 2}
	n2 := []int{1, 2, 3}
	t.Log(kSmallestPairs(n2, n1, 10))

	h := &hp1{}

	heap.Push(h, 1)
	heap.Push(h, 2)
	heap.Push(h, 3)
	t.Log(heap.Pop(h))
	t.Log(heap.Pop(h))
	t.Log(heap.Pop(h))

}
