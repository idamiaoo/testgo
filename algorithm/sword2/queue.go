package sword2

type queue struct {
	cap  int
	size int
	head int
	tail int
	data []interface{}
}

func newQueue(cap int) *queue {
	return &queue{
		cap:  cap,
		data: make([]interface{}, cap),
	}
}

func (q *queue) Push(val interface{}) {
	if q.size < q.cap {
		q.data[q.tail] = val
		q.tail++
		q.size++
	} else {

		if q.tail == q.cap {
			q.tail = 0
		}
	}
}

func (q *queue) Pop() interface{} {
	return nil
}
