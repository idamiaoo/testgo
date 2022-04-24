package leetcode

import "testing"

func moveZeroes(nums []int) {
	var zeroIndex int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i == zeroIndex {
				zeroIndex++
			} else {
				nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
				for ; zeroIndex <= i; zeroIndex++ {
					if nums[zeroIndex] == 0 {
						break
					}
				}
			}
		}
	}
}

func Test283(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
