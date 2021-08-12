package main

import "fmt"

type node struct {
	v    int
	next *node
}

type MaxQueue struct {
	head, tail *node
	mHead      *node
}

func Constructor() MaxQueue {
	return MaxQueue{
		mHead: &node{},
	}
}

func (q *MaxQueue) MaxValue() int {
	if q.head == nil {
		return 0
	}

	if q.mHead.next == nil {
		return q.head.v
	}

	if q.mHead.next.v > q.head.v {
		return q.mHead.next.v
	}
	return q.head.v
}

func (q *MaxQueue) PushBack(value int) {
	nd := &node{v: value}
	mnd := &node{v: value}
	if q.tail == nil {
		q.tail = nd
		q.head = nd
		q.mHead.next = mnd
	} else {
		q.tail.next = nd
		q.tail = nd

		m := q.mHead
		for {
			if m.next == nil {
				m.next = mnd
				break
			}
			if mnd.v > m.next.v {
				m.next = mnd
				break
			}
			m = m.next
		}
	}
}

func (q *MaxQueue) PopFront() int {
	if q.head == nil {
		return 0
	}
	nd := q.head
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	if q.mHead.next != nil && nd.v == q.mHead.next.v {
		q.mHead.next = q.mHead.next.next
	}
	return nd.v
}

func main() {
	queue := Constructor()
	fmt.Println(queue.MaxValue())
	queue.PushBack(1)
	fmt.Println(queue.MaxValue())
	queue.PushBack(2)
	fmt.Println(queue.MaxValue())
	fmt.Println(queue.MaxValue())
	queue.PushBack(1)
	fmt.Println(queue.MaxValue())
	queue.PopFront()
	fmt.Println(queue.MaxValue())
	fmt.Println(queue.MaxValue())
	queue.PopFront()
	fmt.Println(queue.MaxValue())
	queue.PopFront()
	fmt.Println(queue.MaxValue())
}
