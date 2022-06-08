package leetcode

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	newNums := nums[:len(nums)/2]
	for i := range newNums {
		if i%2 == 0 {
			newNums[i] = min2293(nums[2*i], nums[2*i+1])
		} else {
			newNums[i] = max2293(nums[2*i], nums[2*i+1])
		}
	}
	return minMaxGame(newNums)
}

func min2293(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2293(a, b int) int {
	if a > b {
		return a
	}
	return b
}
