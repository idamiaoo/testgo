package leetcode

func findDisappearedNumbers(nums []int) []int {
	n := len(nums) + 1
	for _, v := range nums {
		m := v % n
		nums[m-1] += n
	}
	var ans []int
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
