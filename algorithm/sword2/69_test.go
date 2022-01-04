package sword2

import (
	"testing"
)

func peakIndexInMountainArray(arr []int) int {
	var left, right = 0, len(arr) - 1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func Test69(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}
	t.Log(peakIndexInMountainArray(nums))
}
