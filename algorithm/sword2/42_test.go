package sword2

import (
	"testing"
)

type RecentCounter struct {
	sum  int
	data []int
}

func Constructor42() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	this.sum++

	for len(this.data) > 0 && this.data[0] < t-3000 {
		this.data = this.data[1:]
		this.sum--
	}
	return this.sum
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func Test42(t *testing.T) {
	c := Constructor42()
	t.Log(c.Ping(1))
	t.Log(c.Ping(100))
	t.Log(c.Ping(3001))
	t.Log(c.Ping(3002))

}
