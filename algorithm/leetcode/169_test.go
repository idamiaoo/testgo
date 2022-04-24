package leetcode

func majorityElement(nums []int) int {
	var ans, vote int
	for i := 0; i < len(nums); i++ {
		if vote == 0 {
			ans = nums[i]
			vote = 1
			continue
		}
		if nums[i] == ans {
			vote++
		} else {
			vote--
		}
	}
	return ans
}
