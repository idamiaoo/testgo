package heap

type Heap interface {
	Put(int)
	Pop() int
	Size() int
}

type head struct {
	data []int
}

func (h *head) Put(v int) {
	h.data = append(h.data, v)
	h.up(len(h.data) - 1)
}

func (h *head) Pop() int {
	if len(h.data) == 0 {
		return 0
	}
	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return v
}

func (h *head) Size() int {
	return len(h.data)
}

func (h *head) up(i int) {

}

func (h *head) down(i int) {
	for i > 0 {
		if h.data[i] < h.data[(i-1)/2] {
			h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
}
