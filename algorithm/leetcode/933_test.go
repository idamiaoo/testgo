package leetcode

type RecentCounter struct {
	min int
	d   []int
}

func Constructor933() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.d = append(this.d, t)
	for this.d[this.min] < t-3000 {
		this.min++
	}
	return len(this.d) - this.min
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);

["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]

*/
