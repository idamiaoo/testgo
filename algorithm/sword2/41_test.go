package sword2_test

import (
	"testing"
)

type MovingAverage struct {
	sum        int
	size       int
	head, tail int
	data       []int
}

/** Initialize your data structure here. */
func Constructor41(size int) MovingAverage {
	return MovingAverage{
		tail: -1,
		size: size,
		data: make([]int, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.tail == -1 || (this.tail >= this.head && this.tail < this.size-1) {
		this.tail++
		this.data[this.tail] = val
		this.sum += val
	} else {
		this.sum -= this.data[this.head]
		this.head++
		if this.head == this.size {
			this.head = 0
		}
		this.tail++
		if this.tail == this.size {
			this.tail = 0
		}
		this.data[this.tail] = val
		this.sum += val
	}
	if this.tail < this.head {
		return float64(this.sum) / float64(this.tail-this.head+1+this.size)
	}
	return float64(this.sum) / float64(this.tail-this.head+1)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func Test41(t *testing.T) {
	m := Constructor41(3)
	t.Log(m.Next(1))
	t.Log(m.Next(10))
	t.Log(m.Next(3))
	t.Log(m.Next(5))

}
