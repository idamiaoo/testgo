package leetcode

func singleNumber136(nums []int) int {
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
