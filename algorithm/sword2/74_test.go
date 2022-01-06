package sword2

import (
	"testing"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	var ans [][]int
	var start, end int
	sort74(intervals, 0, len(intervals)-1)
	for j := 1; j < len(intervals); j++ {
		if intervals[end][1] < intervals[j][0] {
			ans = append(ans, []int{intervals[start][0], intervals[end][1]})
			start = j
			end = j
			continue
		}
		if intervals[end][1] > intervals[j][1] {
			continue
		}
		end = j
	}
	ans = append(ans, []int{intervals[start][0], intervals[end][1]})
	return ans
}

func sort74(intervals [][]int, left, right int) {
	if left < right {
		p := partition74(intervals, left, right)
		sort74(intervals, left, p-1)
		sort74(intervals, p+1, right)
	}
}

func partition74(intervals [][]int, left, right int) int {
	b := intervals[right]
	i := left - 1
	for j := left; j < right; j++ {
		if intervals[j][0] < b[0] || (intervals[j][0] == b[0] && intervals[j][1] < b[1]) {
			i++
			intervals[i], intervals[j] = intervals[j], intervals[i]
		}
	}
	i++
	intervals[i], intervals[right] = intervals[right], intervals[i]
	return i
}

func Test74(t *testing.T) {
	intervals := [][]int{{1, 4}, {2, 3}}
	t.Log(merge(intervals))
}
