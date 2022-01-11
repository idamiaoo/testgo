package sword2

func rob90(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob89(nums[1:]), rob89(nums[0:len(nums)-1]))
}

func rob89(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := nums[0]
	s := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f, s = s, max(s, f+nums[i])
	}
	return s
}
