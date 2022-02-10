package leetcode

import (
	"testing"
)

func eatenApples(apples []int, days []int) int {
	store := &minHeap{}
	var i int
	var x int
	for i = 0; ; i++ {
		if i < len(apples) {
			if apples[i] != 0 {
				a := &apple{
					num:    apples[i],
					expire: i + days[i],
				}
				store.add(a)
			}
		}

		for !store.empty() {
			if store.first().expire > i {
				break
			}
			store.del()
		}

		if store.empty() {
			if i >= len(apples) {
				return i - x
			}
			x++
		} else {
			if store.first().num > 1 {
				store.first().num--
			} else {
				store.del()
			}
		}
	}
}

type apple struct {
	num    int
	expire int
}

type minHeap struct {
	data []*apple
}

func (h *minHeap) add(v *apple) {
	h.data = append(h.data, v)
	if len(h.data) > 1 {
		h.up(len(h.data) - 1)
	}
}

func (h *minHeap) first() *apple {
	return h.data[0]
}

func (h *minHeap) empty() bool {
	return len(h.data) <= 0
}

func (h *minHeap) del() *apple {
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
		if i <= p || h.data[i].expire >= h.data[p].expire {
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
				if h.data[c+1].expire < h.data[c].expire {
					c = c + 1
				}
			}
			if h.data[i].expire <= h.data[c].expire {
				break
			}
			h.swap(i, c)
			i = c
		}
	}
}

func Test1705(t *testing.T) {
	apples := []int{3, 0, 0, 0, 0, 2}
	days := []int{3, 0, 0, 0, 0, 2}
	t.Log(eatenApples(apples, days))
}
