package sword2

type KthLargest struct {
	heap *Heap
}

func Constructor59(k int, nums []int) KthLargest {
	h := Heap{cap: k}
	for _, v := range nums {
		h.Add(v)
	}
	return KthLargest{heap: &h}
}

func (this *KthLargest) Add(val int) int {
	this.heap.Add(val)
	return this.heap.data[0]
}

type Heap struct {
	cap  int
	size int
	data []int
}

func (h *Heap) Add(v int) {
	if h.size < h.cap {
		if len(h.data) >= h.size {
			h.data = append(h.data, v)
		} else {
			h.data[h.size] = v
		}
		h.size++
		h.up(h.size - 1)
	} else {
		if v > h.data[0] {
			h.data[0] = v
			h.down(0)
		}
	}

}

func (h *Heap) Del() int {
	v := h.data[0]
	h.swap(0, h.size-1)
	h.size--
	h.down(0)
	return v
}

func (h *Heap) up(i int) {
	var x int
	for i > 0 {
		x = (i - 1) / 2
		if i <= x || h.data[i] >= h.data[x] {
			return
		}
		h.swap(i, x)
		i = x
	}
}

func (h *Heap) down(i int) {
	var x int
	for {
		x = 2*i + 1
		if x >= h.size {
			return
		}
		if x < h.size-1 {
			if h.data[x+1] < h.data[x] {
				x = x + 1
			}
		}
		if h.data[i] <= h.data[x] {
			return
		}
		h.swap(i, x)
		i = x
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

/*
0 12
1 34
2 56
3 78
*/
