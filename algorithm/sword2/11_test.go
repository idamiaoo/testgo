package sword2

import (
	"testing"
)

func findMaxLength(nums []int) int {
	if len(nums) > 1 {
		for i := 1; i < len(nums); i++ {
			nums[i] += nums[i-1]
		}
		for l := (len(nums) / 2) * 2; l >= 2; l -= 2 {
			for i := 0; i+l-1 < len(nums); i++ {
				if i == 0 {
					if nums[i+l-1] == l/2 {
						return l
					}
				} else {
					if nums[i+l-1]-nums[i-1] == l/2 {
						return l
					}
				}
			}
		}
	}
	return 0
}

func Test11(t *testing.T) {
	nums1 := []int{0, 1, 0}
	t.Log(findMaxLength(nums1))
}
