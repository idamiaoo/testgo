package sword2

import (
	"strconv"
	"testing"
)

func findMinDifference(timePoints []string) int {
	timeSort(timePoints, 0, len(timePoints)-1)
	var min = -1
	var h1, m1, h2, m2 int
	for i := 0; i < len(timePoints); i++ {
		if i == len(timePoints)-1 {
			h1, _ = strconv.Atoi(timePoints[i][:2])
			m1, _ = strconv.Atoi(timePoints[i][3:])
			h2, _ = strconv.Atoi(timePoints[0][:2])
			h2 += 24
			m2, _ = strconv.Atoi(timePoints[0][3:])
		} else {
			h1, _ = strconv.Atoi(timePoints[i][:2])
			m1, _ = strconv.Atoi(timePoints[i][3:])
			h2, _ = strconv.Atoi(timePoints[i+1][:2])
			m2, _ = strconv.Atoi(timePoints[i+1][3:])
		}
		if (h2*60+m2)-(h1*60+m1) < min || min == -1 {
			min = (h2*60 + m2) - (h1*60 + m1)
		}
	}
	return min
}

func timeSort(times []string, left, right int) {
	if left < right {
		p := timePartition(times, left, right)
		timeSort(times, left, p-1)
		timeSort(times, p+1, right)
	}
}

func timePartition(times []string, left, right int) int {
	b := times[right]
	i := -1
	for j := 0; j < right; j++ {
		if times[j] <= b {
			i++
			times[i], times[j] = times[j], times[i]
		}
	}
	i++
	times[i], times[right] = times[right], times[i]
	return i
}

func Test35(t *testing.T) {
	timePoints := []string{"00:00", "23:59", "00:00"}
	t.Log(findMinDifference(timePoints))
}
