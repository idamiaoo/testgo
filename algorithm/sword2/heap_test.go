package sword2

import (
	"testing"
)

type minHeap struct {
	data []int
}

func (h *minHeap) add(v int) {
	h.data = append(h.data, v)
	if len(h.data) > 1 {
		h.up(len(h.data) - 1)
	}
}

func (h *minHeap) del() int {
	v := h.data[0]
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return v
}

func (h *minHeap) up(i int) {
	var p int
	for {
		p = (i - 1) / 2
		if i <= p || h.data[i] >= h.data[p] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *minHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeap) down(i int) {
	for {
		c := 2*i + 1
		if c > len(h.data)-1 {
			break
		} else {
			if c < len(h.data)-1 {
				if h.data[c+1] < h.data[c] {
					c = c + 1
				}
			}
			if h.data[i] <= h.data[c] {
				break
			}
			h.swap(i, c)
			i = c
		}
	}
}

func TestMinHeap(t *testing.T) {
	nums := []int{20, 1, 34, 4, 9, 79, 3}
	heap := &minHeap{}

	for _, v := range nums {
		heap.add(v)
	}
	t.Log(heap.data)
	for i := 0; i < len(nums); i++ {
		t.Log(heap.del())
	}
}
