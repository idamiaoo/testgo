package leetcode

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[len(nums)-1] {
			temp := nums[len(nums)-1]
			for j := i + 1; j < len(nums); j++ {
				nums[j] = nums[j-1]
			}
			nums[i] = temp
			return
		}
	}
}
